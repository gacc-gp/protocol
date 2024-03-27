// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tunnel/v1/tunnel.proto

package tunnelv1

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

// Message
type MessageType int32

const (
	MessageType_MESSAGE_TYPE_UNSPECIFIED MessageType = 0
	MessageType_MESSAGE_TYPE_CHECKIN     MessageType = 1
	MessageType_MESSAGE_TYPE_CHECKIN_ACK MessageType = 2
	MessageType_MESSAGE_TYPE_FLOW        MessageType = 101
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:   "MESSAGE_TYPE_UNSPECIFIED",
		1:   "MESSAGE_TYPE_CHECKIN",
		2:   "MESSAGE_TYPE_CHECKIN_ACK",
		101: "MESSAGE_TYPE_FLOW",
	}
	MessageType_value = map[string]int32{
		"MESSAGE_TYPE_UNSPECIFIED": 0,
		"MESSAGE_TYPE_CHECKIN":     1,
		"MESSAGE_TYPE_CHECKIN_ACK": 2,
		"MESSAGE_TYPE_FLOW":        101,
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
	return file_tunnel_v1_tunnel_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_tunnel_v1_tunnel_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{0}
}

type MessageHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MessageType       `protobuf:"varint,1,opt,name=type,proto3,enum=tunnel.v1.MessageType" json:"type,omitempty"`
	Meta map[string]string `protobuf:"bytes,3,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MessageHeader) Reset() {
	*x = MessageHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_v1_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageHeader) ProtoMessage() {}

func (x *MessageHeader) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageHeader.ProtoReflect.Descriptor instead.
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *MessageHeader) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_MESSAGE_TYPE_UNSPECIFIED
}

func (x *MessageHeader) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *MessageHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Body   *anypb.Any     `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_v1_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[1]
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
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeader() *MessageHeader {
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

// CheckIn
type MessageCheckIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *MessageCheckIn) Reset() {
	*x = MessageCheckIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_v1_tunnel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCheckIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCheckIn) ProtoMessage() {}

func (x *MessageCheckIn) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCheckIn.ProtoReflect.Descriptor instead.
func (*MessageCheckIn) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{2}
}

func (x *MessageCheckIn) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *MessageCheckIn) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type MessageCheckInAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *MessageCheckInAck) Reset() {
	*x = MessageCheckInAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_v1_tunnel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCheckInAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCheckInAck) ProtoMessage() {}

func (x *MessageCheckInAck) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCheckInAck.ProtoReflect.Descriptor instead.
func (*MessageCheckInAck) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{3}
}

func (x *MessageCheckInAck) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *MessageCheckInAck) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

// Flow
type Flow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packet []byte `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
}

func (x *Flow) Reset() {
	*x = Flow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_v1_tunnel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flow) ProtoMessage() {}

func (x *Flow) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flow.ProtoReflect.Descriptor instead.
func (*Flow) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{4}
}

func (x *Flow) GetPacket() []byte {
	if x != nil {
		return x.Packet
	}
	return nil
}

var File_tunnel_v1_tunnel_proto protoreflect.FileDescriptor

var file_tunnel_v1_tunnel_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x65, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x43, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x11, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x1e,
	0x0a, 0x04, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2a, 0x7a,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x18, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43,
	0x4b, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x49, 0x4e, 0x5f, 0x41, 0x43,
	0x4b, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x65, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x63, 0x63, 0x2d, 0x67, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tunnel_v1_tunnel_proto_rawDescOnce sync.Once
	file_tunnel_v1_tunnel_proto_rawDescData = file_tunnel_v1_tunnel_proto_rawDesc
)

func file_tunnel_v1_tunnel_proto_rawDescGZIP() []byte {
	file_tunnel_v1_tunnel_proto_rawDescOnce.Do(func() {
		file_tunnel_v1_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunnel_v1_tunnel_proto_rawDescData)
	})
	return file_tunnel_v1_tunnel_proto_rawDescData
}

var file_tunnel_v1_tunnel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tunnel_v1_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tunnel_v1_tunnel_proto_goTypes = []interface{}{
	(MessageType)(0),          // 0: tunnel.v1.MessageType
	(*MessageHeader)(nil),     // 1: tunnel.v1.MessageHeader
	(*Message)(nil),           // 2: tunnel.v1.Message
	(*MessageCheckIn)(nil),    // 3: tunnel.v1.MessageCheckIn
	(*MessageCheckInAck)(nil), // 4: tunnel.v1.MessageCheckInAck
	(*Flow)(nil),              // 5: tunnel.v1.Flow
	nil,                       // 6: tunnel.v1.MessageHeader.MetaEntry
	(*anypb.Any)(nil),         // 7: google.protobuf.Any
}
var file_tunnel_v1_tunnel_proto_depIdxs = []int32{
	0, // 0: tunnel.v1.MessageHeader.type:type_name -> tunnel.v1.MessageType
	6, // 1: tunnel.v1.MessageHeader.meta:type_name -> tunnel.v1.MessageHeader.MetaEntry
	1, // 2: tunnel.v1.Message.header:type_name -> tunnel.v1.MessageHeader
	7, // 3: tunnel.v1.Message.body:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tunnel_v1_tunnel_proto_init() }
func file_tunnel_v1_tunnel_proto_init() {
	if File_tunnel_v1_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tunnel_v1_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageHeader); i {
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
		file_tunnel_v1_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_tunnel_v1_tunnel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCheckIn); i {
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
		file_tunnel_v1_tunnel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCheckInAck); i {
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
		file_tunnel_v1_tunnel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flow); i {
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
			RawDescriptor: file_tunnel_v1_tunnel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tunnel_v1_tunnel_proto_goTypes,
		DependencyIndexes: file_tunnel_v1_tunnel_proto_depIdxs,
		EnumInfos:         file_tunnel_v1_tunnel_proto_enumTypes,
		MessageInfos:      file_tunnel_v1_tunnel_proto_msgTypes,
	}.Build()
	File_tunnel_v1_tunnel_proto = out.File
	file_tunnel_v1_tunnel_proto_rawDesc = nil
	file_tunnel_v1_tunnel_proto_goTypes = nil
	file_tunnel_v1_tunnel_proto_depIdxs = nil
}
