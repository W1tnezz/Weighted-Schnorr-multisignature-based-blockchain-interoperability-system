package iop

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *OracleNode) Enroll(_ context.Context, request *SendEnrollRequest) (*SendEnrollResponse, error) {
	//	此时接收到报名请求
	success := n.aggregator.Enroll(request.Enroll)
	return &SendEnrollResponse{EnrollSuccess: success}, nil
}

func (n *OracleNode) SendR(_ context.Context, request *SendRRequest) (*SendRResponse, error) {
	// 这里接收到了传递过来的参数R
	n.validator.HandleR(request.R)
	return &SendRResponse{}, nil
}

// 这个函数的功能是验证器来验证的过程，以及构造出应答
func (n *OracleNode) Validate(ctx context.Context, request *ValidateRequest) (*ValidateResponse, error) {

	var result *ValidateResult
	var err error

	switch request.Type {
	case ValidateRequest_block:
		result, err = n.validator.ValidateBlock(
			ctx,
			common.BytesToHash(request.Hash),
		)
	case ValidateRequest_transaction:
		result, err = n.validator.ValidateTransaction(
			ctx,
			common.BytesToHash(request.Hash),
			request.minRank,
		)
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate %s: %v", request.Type, err)
	}

	resultStr := "valid"
	if !result.valid {
		resultStr = "invalid"
	}
	log.Infof("Validated hash %s of type %s with result: %s", common.BytesToHash(request.Hash), request.Type, resultStr)

	return ValidateResultToResponse(result), nil
}

func ValidateResultToResponse(result *ValidateResult) *ValidateResponse {
	resp := &ValidateResponse{
		Hash:      result.hash[:],
		Valid:     result.valid,
		Signature: result.signature,
		R:         result.R,
	}

	if result.blockNumber != nil {
		resp.BlockNumber = result.blockNumber.Int64()
	}

	return resp
}
