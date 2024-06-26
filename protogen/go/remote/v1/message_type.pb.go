// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: remote/v1/message_type.proto

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

type MessageType int32

const (
	MessageType_MESSAGE_TYPE_UNSPECIFIED      MessageType = 0
	MessageType_MESSAGE_TYPE_AUTH_REQ         MessageType = 1
	MessageType_MESSAGE_TYPE_AUTH_RSP         MessageType = 2
	MessageType_MESSAGE_TYPE_PING             MessageType = 3
	MessageType_MESSAGE_TYPE_PONG             MessageType = 4
	MessageType_MESSAGE_TYPE_FRAME            MessageType = 5
	MessageType_MESSAGE_TYPE_CONNECTION_CLOSE MessageType = 6
	MessageType_MESSAGE_TYPE_BAD_FRAME        MessageType = 7
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MESSAGE_TYPE_UNSPECIFIED",
		1: "MESSAGE_TYPE_AUTH_REQ",
		2: "MESSAGE_TYPE_AUTH_RSP",
		3: "MESSAGE_TYPE_PING",
		4: "MESSAGE_TYPE_PONG",
		5: "MESSAGE_TYPE_FRAME",
		6: "MESSAGE_TYPE_CONNECTION_CLOSE",
		7: "MESSAGE_TYPE_BAD_FRAME",
	}
	MessageType_value = map[string]int32{
		"MESSAGE_TYPE_UNSPECIFIED":      0,
		"MESSAGE_TYPE_AUTH_REQ":         1,
		"MESSAGE_TYPE_AUTH_RSP":         2,
		"MESSAGE_TYPE_PING":             3,
		"MESSAGE_TYPE_PONG":             4,
		"MESSAGE_TYPE_FRAME":            5,
		"MESSAGE_TYPE_CONNECTION_CLOSE": 6,
		"MESSAGE_TYPE_BAD_FRAME":        7,
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
	return file_remote_v1_message_type_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_remote_v1_message_type_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_remote_v1_message_type_proto_rawDescGZIP(), []int{0}
}

var File_remote_v1_message_type_proto protoreflect.FileDescriptor

var file_remote_v1_message_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0xe6, 0x01, 0x0a, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x51,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x53, 0x50, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x41, 0x4d,
	0x45, 0x10, 0x05, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x44, 0x5f, 0x46, 0x52, 0x41, 0x4d, 0x45,
	0x10, 0x07, 0x42, 0xa2, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x63, 0x63, 0x2d, 0x67, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_v1_message_type_proto_rawDescOnce sync.Once
	file_remote_v1_message_type_proto_rawDescData = file_remote_v1_message_type_proto_rawDesc
)

func file_remote_v1_message_type_proto_rawDescGZIP() []byte {
	file_remote_v1_message_type_proto_rawDescOnce.Do(func() {
		file_remote_v1_message_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_v1_message_type_proto_rawDescData)
	})
	return file_remote_v1_message_type_proto_rawDescData
}

var file_remote_v1_message_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_v1_message_type_proto_goTypes = []interface{}{
	(MessageType)(0), // 0: remote.v1.MessageType
}
var file_remote_v1_message_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_remote_v1_message_type_proto_init() }
func file_remote_v1_message_type_proto_init() {
	if File_remote_v1_message_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_remote_v1_message_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_v1_message_type_proto_goTypes,
		DependencyIndexes: file_remote_v1_message_type_proto_depIdxs,
		EnumInfos:         file_remote_v1_message_type_proto_enumTypes,
	}.Build()
	File_remote_v1_message_type_proto = out.File
	file_remote_v1_message_type_proto_rawDesc = nil
	file_remote_v1_message_type_proto_goTypes = nil
	file_remote_v1_message_type_proto_depIdxs = nil
}
