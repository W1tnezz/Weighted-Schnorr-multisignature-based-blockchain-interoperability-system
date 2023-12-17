package iop

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
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
			if !isAggregator {

				continue
			}
			a.size = event.Size.Int64()
			a.minRank = event.MinRank.Int64()
			a.currentSize = 0
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
	result, MulSig, MulR, _hash, nodes, err := a.AggregateValidationResults(ctx, event.Hash, typ)

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

func (a *Aggregator) AggregateValidationResults(ctx context.Context, txHash common.Hash, typ ValidateRequest_Type) (bool, kyber.Scalar, kyber.Point, kyber.Scalar, []common.Address, error) {

	Signatures := make([][]kyber.Scalar, 0)
	Rs := make([][]kyber.Point, 0)
	PK := make([][][2]*big.Int, 0)
	nodes := make([]common.Address, 0)
	totalRank := int64(0)

	var wg sync.WaitGroup
	var mutex sync.Mutex
	// 获取到了报名的节点数
	time.Sleep(time.Duration(10) * time.Second)

	rand.Seed(time.Now().Unix())

	for _, enrollNodeIndex := range a.enrollNodes {
		enrollNode, err := a.registryContract.FindOracleNodeByIndex(nil, big.NewInt(enrollNodeIndex))
		nodes = append(nodes, enrollNode.Addr)
		node, err := a.registryContract.FindOracleNodeByAddress(nil, enrollNode.Addr)
		conn, err := a.connectionManager.FindByAddress(node.Addr)
		if err != nil {
			log.Errorf("Find connection by address: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
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

			if err != nil {
				log.Errorf("Validate %s: %v", typ, err)
				return
			}

			mutex.Lock()
			if result.Valid {
				totalRank += result.Reputation
				sI := make([]kyber.Scalar, result.Reputation)
				RI := make([]kyber.Point, result.Reputation)
				err := json.Unmarshal(result.Signature, &sI)
				if err != nil {
					return
				}
				errR := json.Unmarshal(result.R, &RI)
				if errR != nil {
					return
				}
				Signatures = append(Signatures, sI) //获取到所有的签名
				Rs = append(Rs, RI)
				PK = append(PK, enrollNode.PubKeys)
			}
			mutex.Unlock()
		}()
	}

	wg.Wait()

	TMPS := make([][]byte, 0)
	for i := 0; i < len(PK); i++ {
		for j := 0; j < len(PK[i]); j++ {
			for k := 0; k < len(PK[i][j]); k++ {
				TMPS = append(TMPS, PK[i][j][k].Bytes())
			}
		}
	}

	MulSignature := a.suite.G1().Scalar().Zero()
	MulR := a.suite.G1().Point().Null()
	R := a.suite.G1().Point().Null()
	for i := 0; i < len(PK); i++ {
		for j := 0; j < len(PK[i]); j++ {
			var S = TMPS
			for k := 0; k < 2; k++ {
				S = append(S, PK[i][j][k].Bytes())
			}
			hash1 := sha256.New()
			SBytes, _ := json.Marshal(S)
			aI := hash1.Sum(SBytes)
			aScalar := a.suite.G1().Scalar().SetBytes(aI)
			MulSignature.Add(MulSignature, a.suite.G1().Scalar().Mul(aScalar, Signatures[i][j]))
			MulR.Add(MulR, a.suite.G1().Point().Mul(aScalar, Rs[i][j]))
			R.Add(R, Rs[i][j])
		}
	}

	message, _ := encodeValidateResult(txHash, true, typ)
	//
	//gt := hash_1.Sum(bytes.Join(PK, []byte("")))
	//
	//m := make([][]byte, 3)
	//m[0] = message
	//m[1], err = R.MarshalBinary()
	//m[2] = gt
	//hash := sha256.New()
	//e := hash.Sum(bytes.Join(m, []byte("")))
	//MulSignature := a.suite.G1().Scalar().Zero()
	//MulR := a.suite.G1().Point().Null()
	//apk := a.suite.G1().Point().Null()
	//
	//for i := 0; i < len(J); i++ {
	//	pub := a.suite.G1().Point()
	//	err = pub.UnmarshalBinary(J[i])
	//	verify_R := Rs[i].Clone()
	//	verify_R.Add(verify_R, a.suite.G1().Point().Mul(a.suite.G1().Scalar().SetBytes(e), pub))
	//	S2 := a.suite.G1().Point().Mul(Signatures[i], nil)
	//	if !verify_R.Equal(S2) {
	//		return false, nil, nil, nil, nil, fmt.Errorf("签名验证失败 ，该签名的公钥为：", J[i])
	//	}
	//	hash := sha256.New()
	//	h := make([][]byte, 3)
	//	h[0] = J[i]
	//	h[1] = bytes.Join(J, []byte(""))
	//	h[2] = bytes.Join(PK, []byte(""))
	//	a_j := hash.Sum(bytes.Join(h, []byte("")))
	//	aScalar := a.suite.G1().Scalar().SetBytes(a_j)
	//
	//	MulSignature.Add(MulSignature, a.suite.G1().Scalar().Mul(aScalar, Signatures[i]))
	//	MulR.Add(MulR, a.suite.G1().Point().Mul(aScalar, Rs[i]))
	//	apk.Add(apk, a.suite.G1().Point().Mul(aScalar, pub))
	//
	//}
	m := make([][]byte, 2)
	m[0] = message
	m[1], _ = R.MarshalBinary()
	hash := sha256.New()
	e := hash.Sum(bytes.Join(m, []byte("")))
	_hash := a.suite.G1().Scalar().SetBytes(e)

	return true, MulSignature, MulR, _hash, nodes, nil

}

//func (a *Aggregator) SetThreshold(threshold int) {
//	a.t = threshold
//}
