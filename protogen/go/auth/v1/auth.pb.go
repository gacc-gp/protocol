// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: auth/v1/auth.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TunnelType int32

const (
	TunnelType_TUNNEL_TYPE_UNSPECIFIED TunnelType = 0
	TunnelType_TUNNEL_TYPE_UDP         TunnelType = 1
	TunnelType_TUNNEL_TYPE_TCP         TunnelType = 2
)

// Enum value maps for TunnelType.
var (
	TunnelType_name = map[int32]string{
		0: "TUNNEL_TYPE_UNSPECIFIED",
		1: "TUNNEL_TYPE_UDP",
		2: "TUNNEL_TYPE_TCP",
	}
	TunnelType_value = map[string]int32{
		"TUNNEL_TYPE_UNSPECIFIED": 0,
		"TUNNEL_TYPE_UDP":         1,
		"TUNNEL_TYPE_TCP":         2,
	}
)

func (x TunnelType) Enum() *TunnelType {
	p := new(TunnelType)
	*p = x
	return p
}

func (x TunnelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TunnelType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_v1_auth_proto_enumTypes[0].Descriptor()
}

func (TunnelType) Type() protoreflect.EnumType {
	return &file_auth_v1_auth_proto_enumTypes[0]
}

func (x TunnelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TunnelType.Descriptor instead.
func (TunnelType) EnumDescriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

type MessageType int32

const (
	MessageType_MESSAGE_TYPE_UNSPECIFIED MessageType = 0
	MessageType_MESSAGE_TYPE_AUTH_REQ    MessageType = 1
	MessageType_MESSAGE_TYPE_AUTH_RESP   MessageType = 2
	MessageType_MESSAGE_TYPE_TUNNEL_REQ  MessageType = 3
	MessageType_MESSAGE_TYPE_TUNNEL_RESP MessageType = 4
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MESSAGE_TYPE_UNSPECIFIED",
		1: "MESSAGE_TYPE_AUTH_REQ",
		2: "MESSAGE_TYPE_AUTH_RESP",
		3: "MESSAGE_TYPE_TUNNEL_REQ",
		4: "MESSAGE_TYPE_TUNNEL_RESP",
	}
	MessageType_value = map[string]int32{
		"MESSAGE_TYPE_UNSPECIFIED": 0,
		"MESSAGE_TYPE_AUTH_REQ":    1,
		"MESSAGE_TYPE_AUTH_RESP":   2,
		"MESSAGE_TYPE_TUNNEL_REQ":  3,
		"MESSAGE_TYPE_TUNNEL_RESP": 4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_v1_auth_proto_enumTypes[1].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_auth_v1_auth_proto_enumTypes[1]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

// Message
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  MessageType       `protobuf:"varint,1,opt,name=type,proto3,enum=auth.v1.MessageType" json:"type,omitempty"`
	Metas map[string]string `protobuf:"bytes,3,rep,name=metas,proto3" json:"metas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_MESSAGE_TYPE_UNSPECIFIED
}

func (x *Header) GetMetas() map[string]string {
	if x != nil {
		return x.Metas
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body   *anypb.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

// Auth
type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TunIp    string `protobuf:"bytes,2,opt,name=tun_ip,json=tunIp,proto3" json:"tun_ip,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AuthReq) GetTunIp() string {
	if x != nil {
		return x.TunIp
	}
	return ""
}

func (x *AuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  int32  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthResp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *AuthResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Tunnel
type TunnelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TunnelType TunnelType `protobuf:"varint,1,opt,name=tunnel_type,json=tunnelType,proto3,enum=auth.v1.TunnelType" json:"tunnel_type,omitempty"`
	// If tunnel_type is UDP, udp port must be provided.
	UdpPort int32 `protobuf:"varint,2,opt,name=udp_port,json=udpPort,proto3" json:"udp_port,omitempty"`
}

func (x *TunnelReq) Reset() {
	*x = TunnelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelReq) ProtoMessage() {}

func (x *TunnelReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelReq.ProtoReflect.Descriptor instead.
func (*TunnelReq) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *TunnelReq) GetTunnelType() TunnelType {
	if x != nil {
		return x.TunnelType
	}
	return TunnelType_TUNNEL_TYPE_UNSPECIFIED
}

func (x *TunnelReq) GetUdpPort() int32 {
	if x != nil {
		return x.UdpPort
	}
	return 0
}

type TunnelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TunnelType TunnelType `protobuf:"varint,1,opt,name=tunnel_type,json=tunnelType,proto3,enum=auth.v1.TunnelType" json:"tunnel_type,omitempty"`
	// If tunnel_type is UDP, udp port must be provided.
	UdpPort int32 `protobuf:"varint,2,opt,name=udp_port,json=udpPort,proto3" json:"udp_port,omitempty"`
}

func (x *TunnelResp) Reset() {
	*x = TunnelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelResp) ProtoMessage() {}

func (x *TunnelResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelResp.ProtoReflect.Descriptor instead.
func (*TunnelResp) Descriptor() ([]byte, []int) {
	return file_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *TunnelResp) GetTunnelType() TunnelType {
	if x != nil {
		return x.TunnelType
	}
	return TunnelType_TUNNEL_TYPE_UNSPECIFIED
}

func (x *TunnelResp) GetUdpPort() int32 {
	if x != nil {
		return x.UdpPort
	}
	return 0
}

var File_auth_v1_auth_proto protoreflect.FileDescriptor

var file_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x1a,
	0x38, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x53, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x75, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x75, 0x6e, 0x49, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x09, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x75, 0x64, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x5d, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x64, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x75, 0x64, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x53, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x44, 0x50, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x55, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x43, 0x50, 0x10, 0x02, 0x2a, 0x9d, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10,
	0x02, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x03, 0x12, 0x1c,
	0x0a, 0x18, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x04, 0x42, 0x8d, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x63, 0x63, 0x2d, 0x67, 0x70, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x41, 0x75,
	0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_auth_proto_rawDescOnce sync.Once
	file_auth_v1_auth_proto_rawDescData = file_auth_v1_auth_proto_rawDesc
)

func file_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_auth_proto_rawDescData)
	})
	return file_auth_v1_auth_proto_rawDescData
}

var file_auth_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_auth_v1_auth_proto_goTypes = []interface{}{
	(TunnelType)(0),    // 0: auth.v1.TunnelType
	(MessageType)(0),   // 1: auth.v1.MessageType
	(*Header)(nil),     // 2: auth.v1.Header
	(*Message)(nil),    // 3: auth.v1.Message
	(*AuthReq)(nil),    // 4: auth.v1.AuthReq
	(*AuthResp)(nil),   // 5: auth.v1.AuthResp
	(*TunnelReq)(nil),  // 6: auth.v1.TunnelReq
	(*TunnelResp)(nil), // 7: auth.v1.TunnelResp
	nil,                // 8: auth.v1.Header.MetasEntry
	(*anypb.Any)(nil),  // 9: google.protobuf.Any
}
var file_auth_v1_auth_proto_depIdxs = []int32{
	1, // 0: auth.v1.Header.type:type_name -> auth.v1.MessageType
	8, // 1: auth.v1.Header.metas:type_name -> auth.v1.Header.MetasEntry
	2, // 2: auth.v1.Message.header:type_name -> auth.v1.Header
	9, // 3: auth.v1.Message.body:type_name -> google.protobuf.Any
	0, // 4: auth.v1.TunnelReq.tunnel_type:type_name -> auth.v1.TunnelType
	0, // 5: auth.v1.TunnelResp.tunnel_type:type_name -> auth.v1.TunnelType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_auth_v1_auth_proto_init() }
func file_auth_v1_auth_proto_init() {
	if File_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
		file_auth_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelReq); i {
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
		file_auth_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelResp); i {
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
			RawDescriptor: file_auth_v1_auth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_auth_v1_auth_proto_depIdxs,
		EnumInfos:         file_auth_v1_auth_proto_enumTypes,
		MessageInfos:      file_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_auth_v1_auth_proto = out.File
	file_auth_v1_auth_proto_rawDesc = nil
	file_auth_v1_auth_proto_goTypes = nil
	file_auth_v1_auth_proto_depIdxs = nil
}
