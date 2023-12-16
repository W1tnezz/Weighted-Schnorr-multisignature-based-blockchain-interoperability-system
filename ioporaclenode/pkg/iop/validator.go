package iop

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"

	"crypto/ecdsa"
	"crypto/sha256"

	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotaledger/hive.go/serializer"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
)

const CONFIRMATIONS uint64 = 5

type ValidateResult struct {
	hash        common.Hash
	valid       bool
	blockNumber *big.Int
	signature   []byte
	R           []byte
}

type Validator struct {
	sync.RWMutex
	suite             pairing.Suite
	registryContract  *RegistryContractWrapper
	oracleContract    *OracleContractWrapper
	ecdsaPrivateKey   *ecdsa.PrivateKey
	ethClient         *ethclient.Client
	connectionManager *ConnectionManager
	RAll              map[uint64]kyber.Point
	account           common.Address
	mqttClient        mqtt.Client
	mqttTopic         []byte
	iotaClient        *iota.NodeHTTPAPIClient
	schnorrPrivateKey []kyber.Scalar
	reputation        int
}

func NewValidator(
	suite pairing.Suite,
	registryContract *RegistryContractWrapper,
	oracleContract *OracleContractWrapper,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	account common.Address,
	mqttClient mqtt.Client,
	mqttTopic []byte,
	iotaClient *iota.NodeHTTPAPIClient,
	schnorrPrivateKey []kyber.Scalar,
	reputation int,
) *Validator {
	return &Validator{
		suite:             suite,
		registryContract:  registryContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		oracleContract:    oracleContract,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		account:           account,
		mqttClient:        mqttClient,
		mqttTopic:         mqttTopic,
		iotaClient:        iotaClient,
		schnorrPrivateKey: schnorrPrivateKey,
	}
}

func (v *Validator) Sign(message []byte) ([][]byte, error) {
	v.RAll = make(map[uint64]kyber.Point)

	//此时要获取所有的报名节点，要考虑是否达到阈值，循环质询

	//enrollNodes, err := v.oracleContract.FindEnrollNodes()

	// 先产生自己的R，然后在等待一段时间，随后广播, 构造R序列
	R_i := make([]kyber.Point, 0)
	for i := 0; i < v.reputation; i++ {
		r := v.suite.G1().Scalar().Pick(random.New())
		R_i = append(R_i, v.suite.G1().Point().Mul(r, nil))
	}

	R_Pi := v.suite.G1().Point().Null()

	for _, R := range R_i {
		R_Pi.Add(R_Pi, R)
	}

	Rbytes, err := R_Pi.MarshalBinary()

	if err != nil {
		fmt.Errorf("marshal R_Pi error : %v", err)
	}
	time.Sleep(5 * time.Second)

	v.sendR(enrollNodes, Rbytes)

	// 此时需要获取到其他人的R,此时需要等待其他人广播完成，获取完全足够的R
	timeout := time.After(Timeout)
	count, err := v.oracleContract.CountEnrollNodes(nil)
loop:
	for {
		select {
		case <-timeout:
			fmt.Errorf("Timeout")
			break loop
		default:
			if count.Int64() == int64(len(v.RAll)) {
				break loop
			}
			time.Sleep(50 * time.Millisecond)
		}
	}

	hash_1 := sha256.New()
	gt := hash_1.Sum(bytes.Join(PK, []byte("")))
	R := v.suite.G1().Point().Null()
	for key := range v.RAll {
		R.Add(R, v.RAll[key])
	}
	m := make([][]byte, 3)
	m[0] = message
	m[1], err = R.MarshalBinary()
	m[2] = gt
	hash := sha256.New()
	e := hash.Sum(bytes.Join(m, []byte("")))
	s := v.suite.G1().Scalar().Add(r, v.suite.G1().Scalar().Mul(v.suite.G1().Scalar().SetBytes(e), v.schnorrPrivateKey))
	signature := make([][]byte, 2)
	signature[0], err = s.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("s transform to Bytes: %w", err)
	}
	signature[1] = Rbytes
	return signature, nil
}

func (v *Validator) ListenAndProcess(o *OracleNode) error {

	// 启动协程监听并处理DKG过程中其他节点发送的Deal
	go func() {
		if err := v.ListenAndProcessResponse(o); err != nil {
			log.Errorf("Listen and process response: %v", err)
		}
	}()
	return nil
}

func (v *Validator) ListenAndProcessResponse(o *OracleNode) error {
	if token := v.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("connect to broker: %w", token.Error())
	}

	topic := fmt.Sprintf("messages/indexation/%s", hex.EncodeToString(v.mqttTopic))
	if token := v.mqttClient.Subscribe(topic, 1, func(c mqtt.Client, m mqtt.Message) {
		v.publishHandler(m, o)
	}); token.Wait() && token.Error() != nil {
		return fmt.Errorf("subscribe to topic: %w", token.Error())
	}

	return nil
}

// 当发布的消息到达该订阅的时候，执行该函数
func (v *Validator) publishHandler(msg mqtt.Message, o *OracleNode) {
	iotaMsg := &iota.Message{}
	if _, err := iotaMsg.Deserialize(msg.Payload(), serializer.DeSeriModeNoValidation); err != nil {
		log.Errorf("Malformed mqtt message: %w", err)
		return
	}
	var response *RDeal
	if err := json.Unmarshal(iotaMsg.Payload.(*iota.Indexation).Data, &response); err != nil {
		log.Errorf("Unmarshal response: %w", err)
	}
	isAggregator := o.isAggregator
	if isAggregator {
		return
	}
	go func() {
		RPoint := v.suite.G1().Point()
		err := RPoint.UnmarshalBinary(response.R)
		if err != nil {
			fmt.Errorf("R transform to Point: %w", err)
		}
		v.RAll[new(big.Int).SetBytes(response.Index).Uint64()] = RPoint
	}()
}

// 这个是发出去的东西进行处理
func (v *Validator) HandleR(R *RDeal) error {
	v.Lock()
	defer v.Unlock()

	if err := v.BroadcastResponse(R); err != nil {
		return fmt.Errorf("broadcast response: %w", err)
	}
	return nil
}

func (v *Validator) BroadcastResponse(R *RDeal) error {
	log.Infof("205 Broadcasting response for R from %d", new(big.Int).SetBytes(R.Index).Uint64())

	b, err := json.Marshal(R)
	payload := &iota.Indexation{
		Index: v.mqttTopic,
		Data:  b,
	}

	msg, err := iota.NewMessageBuilder().
		Payload(payload).
		Build()
	if err != nil {
		return fmt.Errorf("build iota message: %w", err)
	}
	if _, err := v.iotaClient.SubmitMessage(context.Background(), msg); err != nil {
		return fmt.Errorf("submit message: %w", err)
	}
	log.Infof("Broadcast for deal %d completed", new(big.Int).SetBytes(R.Index).Uint64())

	return nil
}

func (v *Validator) sendR(nodes []common.Address, R []byte) {
	node, _ := v.registryContract.FindOracleNodeByAddress(nil, v.account)
	for i := range nodes {
		if nodes[i] == v.account {
			continue
		}
		conn, err := v.connectionManager.FindByAddress(nodes[i])
		if err != nil {
			log.Errorf("Find connection by address: %v", err)
			continue
		}
		client := NewOracleNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

		request := &SendRRequest{
			R: &RDeal{R: R, Index: node.Index.Bytes()},
		}
		log.Infof("Sending R to node %d", i)
		if _, err := client.SendR(ctx, request); err != nil {
			log.Errorf("Send deal: %v", err)
		}
		cancel()
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, hash common.Hash, minRank int) (*ValidateResult, error) {
	if v.reputation < minRank {
		return nil, fmt.Errorf("该节点不参与")
	}

	// 此时，该节点参与，但是需要先向聚合器报名，此时需要发送自己的信誉值

	aggregator, _ := v.registryContract.GetAggregator()

	receipt, err := v.ethClient.TransactionReceipt(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil {
		return nil, fmt.Errorf("transaction receipt: %w", err)
	}

	blockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	valid := false
	if found {
		confirmed := blockNumber - receipt.BlockNumber.Uint64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid, ValidateRequest_transaction)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("dist key share: %w", err)
	}

	// 以下是进行签名，

	sig, err := v.Sign(message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %w", err)
	}

	return &ValidateResult{
		hash,
		valid,
		receipt.BlockNumber,
		sig[0],
		sig[1],
	}, nil
}

func (v *Validator) ValidateBlock(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
	block, err := v.ethClient.BlockByHash(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil && found {
		return nil, fmt.Errorf("block: %w", err)
	}

	latestBlockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	var blockNumber *big.Int
	valid := false
	if found {
		blockNumber = block.Number()
		confirmed := latestBlockNumber - block.NumberU64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid, ValidateRequest_block)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}

	// distKey, err := v.dkg.DistKeyShare()
	if err != nil {
		return nil, fmt.Errorf("dist key share: %w", err)
	}

	sig, err := v.Sign(message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %w", err)
	}

	return &ValidateResult{
		hash,
		valid,
		blockNumber,
		sig[0],
		sig[1],
	}, nil
}
