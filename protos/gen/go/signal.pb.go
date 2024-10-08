// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: signal.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_AUDIO       Type = 0
	Type_TEMPERATURE Type = 1
	Type_HEART_RATE  Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "AUDIO",
		1: "TEMPERATURE",
		2: "HEART_RATE",
	}
	Type_value = map[string]int32{
		"AUDIO":       0,
		"TEMPERATURE": 1,
		"HEART_RATE":  2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_signal_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_signal_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      Type                   `protobuf:"varint,1,opt,name=type,proto3,enum=signal.Type" json:"type,omitempty"`
	Value     float64                `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_signal_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_AUDIO
}

func (x *Signal) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Signal) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_signal_proto protoreflect.FileDescriptor

var file_signal_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2a, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x55, 0x44, 0x49, 0x4f, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x4d, 0x50, 0x45, 0x52,
	0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x41, 0x52, 0x54,
	0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x6c, 0x69, 0x61, 0x6e, 0x37, 0x37, 0x36, 0x2f,
	0x62, 0x61, 0x62, 0x79, 0x2d, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signal_proto_rawDescOnce sync.Once
	file_signal_proto_rawDescData = file_signal_proto_rawDesc
)

func file_signal_proto_rawDescGZIP() []byte {
	file_signal_proto_rawDescOnce.Do(func() {
		file_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_signal_proto_rawDescData)
	})
	return file_signal_proto_rawDescData
}

var file_signal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_signal_proto_goTypes = []any{
	(Type)(0),                     // 0: signal.Type
	(*Signal)(nil),                // 1: signal.Signal
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_signal_proto_depIdxs = []int32{
	0, // 0: signal.Signal.type:type_name -> signal.Type
	2, // 1: signal.Signal.timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_signal_proto_init() }
func file_signal_proto_init() {
	if File_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Signal); i {
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
			RawDescriptor: file_signal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signal_proto_goTypes,
		DependencyIndexes: file_signal_proto_depIdxs,
		EnumInfos:         file_signal_proto_enumTypes,
		MessageInfos:      file_signal_proto_msgTypes,
	}.Build()
	File_signal_proto = out.File
	file_signal_proto_rawDesc = nil
	file_signal_proto_goTypes = nil
	file_signal_proto_depIdxs = nil
}
