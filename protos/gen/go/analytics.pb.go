// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.4
// source: analytics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LastDangerousSignalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LastDangerousSignalRequest) Reset() {
	*x = LastDangerousSignalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastDangerousSignalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastDangerousSignalRequest) ProtoMessage() {}

func (x *LastDangerousSignalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastDangerousSignalRequest.ProtoReflect.Descriptor instead.
func (*LastDangerousSignalRequest) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

type LastDangerousSignalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signal *Signal `protobuf:"bytes,1,opt,name=signal,proto3" json:"signal,omitempty"`
}

func (x *LastDangerousSignalResponse) Reset() {
	*x = LastDangerousSignalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastDangerousSignalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastDangerousSignalResponse) ProtoMessage() {}

func (x *LastDangerousSignalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastDangerousSignalResponse.ProtoReflect.Descriptor instead.
func (*LastDangerousSignalResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *LastDangerousSignalResponse) GetSignal() *Signal {
	if x != nil {
		return x.Signal
	}
	return nil
}

type LastDangerousSignalStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval *durationpb.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *LastDangerousSignalStreamRequest) Reset() {
	*x = LastDangerousSignalStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastDangerousSignalStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastDangerousSignalStreamRequest) ProtoMessage() {}

func (x *LastDangerousSignalStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastDangerousSignalStreamRequest.ProtoReflect.Descriptor instead.
func (*LastDangerousSignalStreamRequest) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *LastDangerousSignalStreamRequest) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

var File_analytics_proto protoreflect.FileDescriptor

var file_analytics_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x4c, 0x61,
	0x73, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x1b, 0x4c, 0x61, 0x73, 0x74,
	0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22,
	0x59, 0x0a, 0x20, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x32, 0xe9, 0x01, 0x0a, 0x09, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x66, 0x0a, 0x13, 0x4c, 0x61, 0x73, 0x74,
	0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12,
	0x25, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x61, 0x73, 0x74,
	0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x74, 0x0a, 0x19, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75,
	0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61,
	0x6e, 0x67, 0x65, 0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65,
	0x72, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x6c, 0x69, 0x61, 0x6e, 0x37, 0x37, 0x36, 0x2f, 0x62,
	0x61, 0x62, 0x79, 0x2d, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_proto_rawDescOnce sync.Once
	file_analytics_proto_rawDescData = file_analytics_proto_rawDesc
)

func file_analytics_proto_rawDescGZIP() []byte {
	file_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_proto_rawDescData)
	})
	return file_analytics_proto_rawDescData
}

var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_analytics_proto_goTypes = []any{
	(*LastDangerousSignalRequest)(nil),       // 0: analytics.LastDangerousSignalRequest
	(*LastDangerousSignalResponse)(nil),      // 1: analytics.LastDangerousSignalResponse
	(*LastDangerousSignalStreamRequest)(nil), // 2: analytics.LastDangerousSignalStreamRequest
	(*Signal)(nil),                           // 3: signal.Signal
	(*durationpb.Duration)(nil),              // 4: google.protobuf.Duration
}
var file_analytics_proto_depIdxs = []int32{
	3, // 0: analytics.LastDangerousSignalResponse.signal:type_name -> signal.Signal
	4, // 1: analytics.LastDangerousSignalStreamRequest.interval:type_name -> google.protobuf.Duration
	0, // 2: analytics.Analytics.LastDangerousSignal:input_type -> analytics.LastDangerousSignalRequest
	2, // 3: analytics.Analytics.LastDangerousSignalStream:input_type -> analytics.LastDangerousSignalStreamRequest
	1, // 4: analytics.Analytics.LastDangerousSignal:output_type -> analytics.LastDangerousSignalResponse
	1, // 5: analytics.Analytics.LastDangerousSignalStream:output_type -> analytics.LastDangerousSignalResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_analytics_proto_init() }
func file_analytics_proto_init() {
	if File_analytics_proto != nil {
		return
	}
	file_signal_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_analytics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LastDangerousSignalRequest); i {
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
		file_analytics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LastDangerousSignalResponse); i {
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
		file_analytics_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LastDangerousSignalStreamRequest); i {
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
			RawDescriptor: file_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_proto_depIdxs,
		MessageInfos:      file_analytics_proto_msgTypes,
	}.Build()
	File_analytics_proto = out.File
	file_analytics_proto_rawDesc = nil
	file_analytics_proto_goTypes = nil
	file_analytics_proto_depIdxs = nil
}
