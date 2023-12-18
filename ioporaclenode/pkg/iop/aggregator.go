package iop

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"go.dedis.ch/kyber/v3/util/random"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Aggregator struct {
	suite             pairing.Suite
	ethClient         *ethclient.Client
	connectionManager *ConnectionManager
	oracleContract    *OracleContractWrapper
	registryContract  *RegistryContractWrapper
	account           common.Address
	ecdsaPrivateKey   *ecdsa.PrivateKey
	chainId           *big.Int
	size              int64 // 总的
	enrollNodes       []int64
	minRank           int64 // 每次节点的最小
	currentSize       int64
	reputation        bool
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	oracleContract *OracleContractWrapper,
	registryContract *RegistryContractWrapper,
	account common.Address,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	chainId *big.Int,
	enrollNodes []int64,
) *Aggregator {
	return &Aggregator{
		suite:             suite,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		account:           account,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		chainId:           chainId,
		enrollNodes:       enrollNodes,
	}
}

func (a *Aggregator) WatchAndHandleValidationRequestsLog(ctx context.Context, o *OracleNode) error {
	sink := make(chan *OracleContractValidationRequest)
	defer close(sink)

	sub, err := a.oracleContract.WatchValidationRequest(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			typ := ValidateRequest_Type(event.Typ)
			log.Infof("Received ValidationRequest event for %s type with hash %s", typ, common.Hash(event.Hash))
			isAggregator, err := a.registryContract.IsAggregator(nil, a.account)
			o.isAggregator = isAggregator
			if err != nil {
				log.Errorf("Is aggregator: %v", err)
				continue
			}
			o.aggregator.size = event.Size.Int64()
			o.aggregator.minRank = event.MinRank.Int64()
			o.aggregator.currentSize = 0

			if !isAggregator {
				a.ValidatorEnroll(o)
				continue
			}

			if err := a.HandleValidationRequest(ctx, event, typ); err != nil {
				log.Errorf("Handle ValidationRequest log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (a *Aggregator) ValidatorEnroll(o *OracleNode) {
	if o.reputation < o.aggregator.minRank {
		return
	}

	// 此时，该节点参与，但是需要先向聚合器报名，此时需要发送自己的信誉值
	node, _ := o.registryContract.FindOracleNodeByAddress(nil, a.account)

	aggregator, _ := o.registryContract.GetAggregator(nil)
	conn, err := o.connectionManager.FindByAddress(aggregator)
	if err != nil {
		log.Errorf("Find connection by address: %v", err)
	}
	client := NewOracleNodeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	request := &SendEnrollRequest{
		Enroll: &EnrollDeal{Reputation: o.reputation, Index: node.Index.Bytes()},
	}

	log.Infof("Sending EnrollRequest to Aggregator %d", node.Index)
	result, err := client.Enroll(ctx, request)
	if err != nil {
		log.Errorf("Send EnrollRequest: %v", err)
	}
	cancel()

	if !result.EnrollSuccess {
		log.Infof("node enroll fail %d", node.Index)
		return
	}
}

// 报名函数

func (a *Aggregator) Enroll(deal *EnrollDeal) bool {
	index := new(big.Int).SetBytes(deal.Index).Int64()
	isEnroll := a.isEnroll(index)

	// 验证器向聚合器节点进行报名
	if !isEnroll && a.minRank <= deal.Reputation && a.currentSize < a.size {
		a.enrollNodes = append(a.enrollNodes, index)
		a.currentSize += deal.Reputation

		return true
	}
	return false
}

// 判断是否报名
func (a *Aggregator) isEnroll(index int64) bool {
	for _, enrollNode := range a.enrollNodes {
		if index == enrollNode {
			return true
		}
	}
	return false
}

// 获取报名节点

func (a *Aggregator) getEnrollNodes(getNode bool) ([]int64, bool) {
	if !getNode {
		return nil, false
	}
	if a.currentSize >= a.size {
		return a.enrollNodes, true
	}
	return nil, false
}

func (a *Aggregator) HandleValidationRequest(ctx context.Context, event *OracleContractValidationRequest, typ ValidateRequest_Type) error {
	result, MulSig, MulR, _hash, MulY, nodes, err := a.AggregateValidationResults(ctx, event.Hash, typ)

	pk, err := PointToBig(MulY)
	fmt.Println(pk)
	if err != nil {
		return fmt.Errorf("aggregate validation results: %w", err)
	}
	if err != nil {
		return fmt.Errorf("signature to big int: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(a.ecdsaPrivateKey, a.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}

	sig, err := ScalarToBig(MulSig)
	if err != nil {
		return fmt.Errorf("signature tranform to big int: %w", err)
	}
	if err != nil {
		return fmt.Errorf("public key tranform to big int: %w", err)
	}
	R, err := PointToBig(MulR)

	if err != nil {
		return fmt.Errorf("multi R tranform to big int: %w", err)
	}
	hash, err := ScalarToBig(_hash)
	if err != nil {
		return fmt.Errorf("hash tranform to big int: %w", err)
	}
	switch typ {
	case ValidateRequest_block:
		_, err = a.oracleContract.SubmitBlockValidationResult(auth, result, event.Hash, sig, R[0], R[1], hash, nodes)
	case ValidateRequest_transaction:
		_, err = a.oracleContract.SubmitTransactionValidationResult(auth, result, event.Hash, sig, R[0], R[1], hash, nodes)
	default:
		return fmt.Errorf("unknown validation request type %s", typ)
	}

	if err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}

	resultStr := "valid"
	if !result {
		resultStr = "invalid"
	}
	log.Infof("Submitted validation result (%s) for hash %s of type %s", resultStr, common.Hash(event.Hash), typ)

	return nil
}

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash, typ ValidateRequest_Type) (bool, kyber.Scalar, kyber.Point, kyber.Scalar, kyber.Point, []common.Address, error) {

	Signatures := make([][]kyber.Scalar, 0)
	Rs := make([][]kyber.Point, 0)
	PK := make([][][2]*big.Int, 0)
	nodes := make([]common.Address, 0)
	totalRank := int64(0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	timeout := time.After(Timeout)
loop:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop
		default:
			if a.currentSize >= a.size {
				break loop
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}

	rand.Seed(time.Now().Unix())
	for _, enrollNodeIndex := range a.enrollNodes {
		enrollNode, err := a.registryContract.FindOracleNodeByIndex(nil, big.NewInt(enrollNodeIndex))

		node, err := a.registryContract.FindOracleNodeByAddress(nil, enrollNode.Addr)
		conn, err := a.connectionManager.FindByAddress(node.Addr)
		if err != nil {
			log.Errorf("Find connection by address: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			client := NewOracleNodeClient(conn)
			ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)

			result, err := client.Validate(ctxTimeout, &ValidateRequest{
				Type:    typ,
				Hash:    txHash[:],
				Size:    a.size,
				MinRank: a.minRank,
			})

			cancel()
			if err != nil {
				log.Errorf("Validate %s: %v", typ, err)
				return
			}

			mutex.Lock()
			if result.Valid {
				totalRank += result.Reputation
				sI := make([]kyber.Scalar, 0)
				RI := make([]kyber.Point, 0)
				nodes = append(nodes, enrollNode.Addr)
				tmpScalar := a.suite.G1().Scalar().Pick(random.New())

				tmpPoint := a.suite.G1().Point().Mul(tmpScalar, nil)

				scalarSize := tmpScalar.MarshalSize()
				PointSize := tmpPoint.MarshalSize()

				for i := int64(0); i < result.Reputation; i++ {
					sSlice := result.Signature[i*int64(scalarSize) : (i+1)*int64(scalarSize)]

					sI = append(sI, a.suite.G1().Scalar().SetBytes(sSlice))

					RSliceBytes := result.R[i*int64(PointSize) : (i+1)*int64(PointSize)]
					RSlice := a.suite.G1().Point().Base()
					err := RSlice.UnmarshalBinary(RSliceBytes)
					if err != nil {
						fmt.Println("UnmarshalBinary R ,", err)
					}
					RI = append(RI, RSlice)
				}

				Signatures = append(Signatures, sI) //获取到所有的签名
				Rs = append(Rs, RI)

				PK = append(PK, enrollNode.PubKeys)

			}

		}()
	}

	wg.Wait()

	index := 64
	S := make([]byte, (totalRank+1)*64)
	for i := 0; i < len(a.enrollNodes); i++ {
		for j := 0; j < len(PK[i]); j++ {
			for k := 0; k < 2; k++ {
				tmp := PK[i][j][k].Bytes()
				for _, byteTmp := range tmp {
					S[index] = byteTmp
					index++
				}
			}
		}

	}

	MulSignature := a.suite.G1().Scalar().Zero()
	MulR := a.suite.G1().Point().Null()
	MulY := a.suite.G1().Point().Null()

	R := a.suite.G1().Point().Null()

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(PK[i]); j++ {

			tmpX := PK[i][j][0]
			tmpY := PK[i][j][1]
			tmpXByte := tmpX.Bytes()
			XByte := make([]byte, 32)

			for k := 31; k >= 0; k-- {
				if len(tmpXByte)-(len(XByte)-k) >= 0 {
					XByte[k] = tmpXByte[len(tmpXByte)-(len(XByte)-k)]
				} else {
					XByte[k] = 0
				}

			}
			for k := 0; k < 32; k++ {
				S[k] = XByte[k]
			}
			tmpYByte := tmpY.Bytes()

			YByte := make([]byte, 32)
			for k := 31; k >= 0; k-- {
				if len(tmpYByte)-(len(YByte)-k) >= 0 {
					YByte[k] = tmpYByte[len(tmpYByte)-(len(YByte)-k)]
				} else {
					YByte[k] = 0
				}

			}
			for k := 0; k < 32; k++ {
				S[k+32] = YByte[k]
			}
			pkbytes := S[0:64]
			pk := a.suite.G1().Point().Null()
			err := pk.UnmarshalBinary(pkbytes)

			if err != nil {
				fmt.Println("translate pk ", err)
			}

			hash1 := sha256.New()

			hash1.Write(S)
			aI := hash1.Sum(nil)

			aScalar := a.suite.G1().Scalar().SetBytes(aI)
			MulSignature.Add(MulSignature, a.suite.G1().Scalar().Mul(aScalar, Signatures[i][j]))
			MulY.Add(MulY, a.suite.G1().Point().Mul(aScalar, pk))
			MulR.Add(MulR, a.suite.G1().Point().Mul(aScalar, Rs[i][j]))
			R.Add(R, Rs[i][j])
		}
	}

	message, _ := encodeValidateResult(txHash, true, typ)

	m := make([][]byte, 2)
	m[0] = message
	m[1], _ = R.MarshalBinary()
	hash := sha256.New()
	e := hash.Sum(bytes.Join(m, []byte("")))
	_hash := a.suite.G1().Scalar().SetBytes(e)

	left := a.suite.G1().Point().Mul(MulSignature, nil)
	right := MulR.Clone()

	right.Add(right, a.suite.G1().Point().Mul(_hash, MulY))
	fmt.Println("435", right.Equal(left))
	a.enrollNodes = []int64{}

	return true, MulSignature, MulR, _hash, MulY, nodes, nil

}

//func (a *Aggregator) SetThreshold(threshold int) {
//	a.t = threshold
//}
