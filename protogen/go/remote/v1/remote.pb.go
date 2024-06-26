// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: remote/v1/remote.proto

package remotev1

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

type AuthResult int32

const (
	AuthResult_AUTH_RESULT_SUCC          AuthResult = 0
	AuthResult_AUTH_RESULT_TOKEN_INVALID AuthResult = 1
)

// Enum value maps for AuthResult.
var (
	AuthResult_name = map[int32]string{
		0: "AUTH_RESULT_SUCC",
		1: "AUTH_RESULT_TOKEN_INVALID",
	}
	AuthResult_value = map[string]int32{
		"AUTH_RESULT_SUCC":          0,
		"AUTH_RESULT_TOKEN_INVALID": 1,
	}
)

func (x AuthResult) Enum() *AuthResult {
	p := new(AuthResult)
	*p = x
	return p
}

func (x AuthResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthResult) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_v1_remote_proto_enumTypes[0].Descriptor()
}

func (AuthResult) Type() protoreflect.EnumType {
	return &file_remote_v1_remote_proto_enumTypes[0]
}

func (x AuthResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthResult.Descriptor instead.
func (AuthResult) EnumDescriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{0}
}

// Auth
type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TunIp    string `protobuf:"bytes,2,opt,name=tun_ip,json=tunIp,proto3" json:"tun_ip,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Retry    int32  `protobuf:"varint,4,opt,name=retry,proto3" json:"retry,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[0]
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
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{0}
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

func (x *AuthReq) GetRetry() int32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

// AuthRsp
type AuthRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  AuthResult `protobuf:"varint,1,opt,name=result,proto3,enum=remote.v1.AuthResult" json:"result,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthRsp) Reset() {
	*x = AuthRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRsp) ProtoMessage() {}

func (x *AuthRsp) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRsp.ProtoReflect.Descriptor instead.
func (*AuthRsp) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRsp) GetResult() AuthResult {
	if x != nil {
		return x.Result
	}
	return AuthResult_AUTH_RESULT_SUCC
}

func (x *AuthRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Frame
type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame []byte `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{2}
}

func (x *Frame) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

// Ping/Pong/ConnectionClose/BadFrame
// Currently these are empty messages.
type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{3}
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{4}
}

type ConnectionClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionClose) Reset() {
	*x = ConnectionClose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionClose) ProtoMessage() {}

func (x *ConnectionClose) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionClose.ProtoReflect.Descriptor instead.
func (*ConnectionClose) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{5}
}

type BadFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BadFrame) Reset() {
	*x = BadFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_v1_remote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadFrame) ProtoMessage() {}

func (x *BadFrame) ProtoReflect() protoreflect.Message {
	mi := &file_remote_v1_remote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadFrame.ProtoReflect.Descriptor instead.
func (*BadFrame) Descriptor() ([]byte, []int) {
	return file_remote_v1_remote_proto_rawDescGZIP(), []int{6}
}

var File_remote_v1_remote_proto protoreflect.FileDescriptor

var file_remote_v1_remote_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x69, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x75, 0x6e, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x75, 0x6e,
	0x49, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x22, 0x52,
	0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x6e,
	0x67, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x42, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x2a, 0x41, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x01, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x61, 0x63, 0x63, 0x2d, 0x67, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x15, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_v1_remote_proto_rawDescOnce sync.Once
	file_remote_v1_remote_proto_rawDescData = file_remote_v1_remote_proto_rawDesc
)

func file_remote_v1_remote_proto_rawDescGZIP() []byte {
	file_remote_v1_remote_proto_rawDescOnce.Do(func() {
		file_remote_v1_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_v1_remote_proto_rawDescData)
	})
	return file_remote_v1_remote_proto_rawDescData
}

var file_remote_v1_remote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_v1_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_remote_v1_remote_proto_goTypes = []interface{}{
	(AuthResult)(0),         // 0: remote.v1.AuthResult
	(*AuthReq)(nil),         // 1: remote.v1.AuthReq
	(*AuthRsp)(nil),         // 2: remote.v1.AuthRsp
	(*Frame)(nil),           // 3: remote.v1.Frame
	(*Ping)(nil),            // 4: remote.v1.Ping
	(*Pong)(nil),            // 5: remote.v1.Pong
	(*ConnectionClose)(nil), // 6: remote.v1.ConnectionClose
	(*BadFrame)(nil),        // 7: remote.v1.BadFrame
}
var file_remote_v1_remote_proto_depIdxs = []int32{
	0, // 0: remote.v1.AuthRsp.result:type_name -> remote.v1.AuthResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_remote_v1_remote_proto_init() }
func file_remote_v1_remote_proto_init() {
	if File_remote_v1_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_v1_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_remote_v1_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRsp); i {
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
		file_remote_v1_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_remote_v1_remote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_remote_v1_remote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_remote_v1_remote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionClose); i {
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
		file_remote_v1_remote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadFrame); i {
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
			RawDescriptor: file_remote_v1_remote_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_v1_remote_proto_goTypes,
		DependencyIndexes: file_remote_v1_remote_proto_depIdxs,
		EnumInfos:         file_remote_v1_remote_proto_enumTypes,
		MessageInfos:      file_remote_v1_remote_proto_msgTypes,
	}.Build()
	File_remote_v1_remote_proto = out.File
	file_remote_v1_remote_proto_rawDesc = nil
	file_remote_v1_remote_proto_goTypes = nil
	file_remote_v1_remote_proto_depIdxs = nil
}
