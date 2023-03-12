// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: client.proto

package iop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OracleNodeClient is the client API for OracleNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OracleNodeClient interface {
	SendDeal(ctx context.Context, in *SendDealRequest, opts ...grpc.CallOption) (*SendDealResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Chanllenge(ctx context.Context, in *ChanllengeRequest, opts ...grpc.CallOption) (*ChanllengeResponse, error)
}

type oracleNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewOracleNodeClient(cc grpc.ClientConnInterface) OracleNodeClient {
	return &oracleNodeClient{cc}
}

func (c *oracleNodeClient) SendDeal(ctx context.Context, in *SendDealRequest, opts ...grpc.CallOption) (*SendDealResponse, error) {
	out := new(SendDealResponse)
	err := c.cc.Invoke(ctx, "/iop.OracleNode/SendDeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleNodeClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/iop.OracleNode/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleNodeClient) Chanllenge(ctx context.Context, in *ChanllengeRequest, opts ...grpc.CallOption) (*ChanllengeResponse, error) {
	out := new(ChanllengeResponse)
	err := c.cc.Invoke(ctx, "/iop.OracleNode/Chanllenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OracleNodeServer is the server API for OracleNode service.
// All implementations must embed UnimplementedOracleNodeServer
// for forward compatibility
type OracleNodeServer interface {
	SendDeal(context.Context, *SendDealRequest) (*SendDealResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Chanllenge(context.Context, *ChanllengeRequest) (*ChanllengeResponse, error)
	mustEmbedUnimplementedOracleNodeServer()
}

// UnimplementedOracleNodeServer must be embedded to have forward compatible implementations.
type UnimplementedOracleNodeServer struct {
}

func (UnimplementedOracleNodeServer) SendDeal(context.Context, *SendDealRequest) (*SendDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDeal not implemented")
}
func (UnimplementedOracleNodeServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedOracleNodeServer) Chanllenge(context.Context, *ChanllengeRequest) (*ChanllengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chanllenge not implemented")
}
func (UnimplementedOracleNodeServer) mustEmbedUnimplementedOracleNodeServer() {}

// UnsafeOracleNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OracleNodeServer will
// result in compilation errors.
type UnsafeOracleNodeServer interface {
	mustEmbedUnimplementedOracleNodeServer()
}

func RegisterOracleNodeServer(s grpc.ServiceRegistrar, srv OracleNodeServer) {
	s.RegisterService(&OracleNode_ServiceDesc, srv)
}

func _OracleNode_SendDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleNodeServer).SendDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iop.OracleNode/SendDeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleNodeServer).SendDeal(ctx, req.(*SendDealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OracleNode_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleNodeServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iop.OracleNode/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleNodeServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OracleNode_Chanllenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChanllengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleNodeServer).Chanllenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iop.OracleNode/Chanllenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleNodeServer).Chanllenge(ctx, req.(*ChanllengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OracleNode_ServiceDesc is the grpc.ServiceDesc for OracleNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OracleNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iop.OracleNode",
	HandlerType: (*OracleNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendDeal",
			Handler:    _OracleNode_SendDeal_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _OracleNode_Validate_Handler,
		},
		{
			MethodName: "Chanllenge",
			Handler:    _OracleNode_Chanllenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
