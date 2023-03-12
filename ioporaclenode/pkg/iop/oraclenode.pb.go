// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: client.proto

package iop

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateRequest_Type int32

const (
	ValidateRequest_unknown     ValidateRequest_Type = 0
	ValidateRequest_block       ValidateRequest_Type = 1
	ValidateRequest_transaction ValidateRequest_Type = 2
)

// Enum value maps for ValidateRequest_Type.
var (
	ValidateRequest_Type_name = map[int32]string{
		0: "unknown",
		1: "block",
		2: "transaction",
	}
	ValidateRequest_Type_value = map[string]int32{
		"unknown":     0,
		"block":       1,
		"transaction": 2,
	}
)

func (x ValidateRequest_Type) Enum() *ValidateRequest_Type {
	p := new(ValidateRequest_Type)
	*p = x
	return p
}

func (x ValidateRequest_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidateRequest_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[0].Descriptor()
}

func (ValidateRequest_Type) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[0]
}

func (x ValidateRequest_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidateRequest_Type.Descriptor instead.
func (ValidateRequest_Type) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2, 0}
}

type SendDealRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deal *Deal `protobuf:"bytes,1,opt,name=deal,proto3" json:"deal,omitempty"`
}

func (x *SendDealRequest) Reset() {
	*x = SendDealRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDealRequest) ProtoMessage() {}

func (x *SendDealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDealRequest.ProtoReflect.Descriptor instead.
func (*SendDealRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *SendDealRequest) GetDeal() *Deal {
	if x != nil {
		return x.Deal
	}
	return nil
}

type SendDealResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendDealResponse) Reset() {
	*x = SendDealResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDealResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDealResponse) ProtoMessage() {}

func (x *SendDealResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDealResponse.ProtoReflect.Descriptor instead.
func (*SendDealResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ValidateRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=iop.ValidateRequest_Type" json:"type,omitempty"`
	Hash []byte               `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateRequest) GetType() ValidateRequest_Type {
	if x != nil {
		return x.Type
	}
	return ValidateRequest_unknown
}

func (x *ValidateRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Deal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Deal      *EncryptedDeal `protobuf:"bytes,2,opt,name=deal,proto3" json:"deal,omitempty"`
	Signature []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Deal) Reset() {
	*x = Deal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deal) ProtoMessage() {}

func (x *Deal) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deal.ProtoReflect.Descriptor instead.
func (*Deal) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *Deal) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Deal) GetDeal() *EncryptedDeal {
	if x != nil {
		return x.Deal
	}
	return nil
}

func (x *Deal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EncryptedDeal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DhKey     []byte `protobuf:"bytes,1,opt,name=dhKey,proto3" json:"dhKey,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce     []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Cipher    []byte `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (x *EncryptedDeal) Reset() {
	*x = EncryptedDeal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedDeal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedDeal) ProtoMessage() {}

func (x *EncryptedDeal) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedDeal.ProtoReflect.Descriptor instead.
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptedDeal) GetDhKey() []byte {
	if x != nil {
		return x.DhKey
	}
	return nil
}

func (x *EncryptedDeal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *EncryptedDeal) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EncryptedDeal) GetCipher() []byte {
	if x != nil {
		return x.Cipher
	}
	return nil
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash        []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Valid       bool   `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	BlockNumber int64  `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateResponse) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ValidateResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateResponse) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ValidateResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ChanllengeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       []byte `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Chanllenge []byte `protobuf:"bytes,2,opt,name=chanllenge,proto3" json:"chanllenge,omitempty"`
}

func (x *ChanllengeRequest) Reset() {
	*x = ChanllengeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChanllengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChanllengeRequest) ProtoMessage() {}

func (x *ChanllengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChanllengeRequest.ProtoReflect.Descriptor instead.
func (*ChanllengeRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{6}
}

func (x *ChanllengeRequest) GetType() []byte {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ChanllengeRequest) GetChanllenge() []byte {
	if x != nil {
		return x.Chanllenge
	}
	return nil
}

type ChanllengeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *ChanllengeResponse) Reset() {
	*x = ChanllengeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChanllengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChanllengeResponse) ProtoMessage() {}

func (x *ChanllengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChanllengeResponse.ProtoReflect.Descriptor instead.
func (*ChanllengeResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{7}
}

func (x *ChanllengeResponse) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x69, 0x6f, 0x70, 0x22, 0x30, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x52,
	0x04, 0x64, 0x65, 0x61, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x69, 0x6f,
	0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x02, 0x22, 0x62, 0x0a, 0x04, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x69, 0x6f, 0x70, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x61,
	0x6c, 0x52, 0x04, 0x64, 0x65, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x71, 0x0a, 0x0d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x64, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x22, 0x7c, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22,
	0x2a, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0xbd, 0x01, 0x0a, 0x0a,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69,
	0x6f, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x69, 0x6f, 0x70,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6f, 0x70, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x3b, 0x69, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_client_proto_goTypes = []interface{}{
	(ValidateRequest_Type)(0),  // 0: iop.ValidateRequest.Type
	(*SendDealRequest)(nil),    // 1: iop.SendDealRequest
	(*SendDealResponse)(nil),   // 2: iop.SendDealResponse
	(*ValidateRequest)(nil),    // 3: iop.ValidateRequest
	(*Deal)(nil),               // 4: iop.Deal
	(*EncryptedDeal)(nil),      // 5: iop.EncryptedDeal
	(*ValidateResponse)(nil),   // 6: iop.ValidateResponse
	(*ChanllengeRequest)(nil),  // 7: iop.ChanllengeRequest
	(*ChanllengeResponse)(nil), // 8: iop.ChanllengeResponse
}
var file_client_proto_depIdxs = []int32{
	4, // 0: iop.SendDealRequest.deal:type_name -> iop.Deal
	0, // 1: iop.ValidateRequest.type:type_name -> iop.ValidateRequest.Type
	5, // 2: iop.Deal.deal:type_name -> iop.EncryptedDeal
	1, // 3: iop.OracleNode.SendDeal:input_type -> iop.SendDealRequest
	3, // 4: iop.OracleNode.Validate:input_type -> iop.ValidateRequest
	7, // 5: iop.OracleNode.Chanllenge:input_type -> iop.ChanllengeRequest
	2, // 6: iop.OracleNode.SendDeal:output_type -> iop.SendDealResponse
	6, // 7: iop.OracleNode.Validate:output_type -> iop.ValidateResponse
	8, // 8: iop.OracleNode.Chanllenge:output_type -> iop.ChanllengeResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDealRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDealResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedDeal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChanllengeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChanllengeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		EnumInfos:         file_client_proto_enumTypes,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
