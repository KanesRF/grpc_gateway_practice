// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/simple_calc.proto

package grpc_calc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First     uint64 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	Second    uint64 `protobuf:"varint,3,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_proto_simple_calc_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetFirst() uint64 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *Operation) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Operation) GetSecond() uint64 {
	if x != nil {
		return x.Second
	}
	return 0
}

type OperationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *OperationResult) Reset() {
	*x = OperationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_simple_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResult) ProtoMessage() {}

func (x *OperationResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_simple_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResult.ProtoReflect.Descriptor instead.
func (*OperationResult) Descriptor() ([]byte, []int) {
	return file_proto_simple_calc_proto_rawDescGZIP(), []int{1}
}

func (x *OperationResult) GetResult() uint64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *OperationResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_simple_calc_proto protoreflect.FileDescriptor

var file_proto_simple_calc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x61, 0x6c, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x57, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x75, 0x0a, 0x0d,
	0x43, 0x61, 0x6c, 0x63, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x64, 0x0a,
	0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x7b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x7d, 0x2f, 0x7b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x7d, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x6c,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_simple_calc_proto_rawDescOnce sync.Once
	file_proto_simple_calc_proto_rawDescData = file_proto_simple_calc_proto_rawDesc
)

func file_proto_simple_calc_proto_rawDescGZIP() []byte {
	file_proto_simple_calc_proto_rawDescOnce.Do(func() {
		file_proto_simple_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_simple_calc_proto_rawDescData)
	})
	return file_proto_simple_calc_proto_rawDescData
}

var file_proto_simple_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_simple_calc_proto_goTypes = []interface{}{
	(*Operation)(nil),       // 0: grpc_calc.Operation
	(*OperationResult)(nil), // 1: grpc_calc.OperationResult
}
var file_proto_simple_calc_proto_depIdxs = []int32{
	0, // 0: grpc_calc.CalcPerformer.Calculate:input_type -> grpc_calc.Operation
	1, // 1: grpc_calc.CalcPerformer.Calculate:output_type -> grpc_calc.OperationResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_simple_calc_proto_init() }
func file_proto_simple_calc_proto_init() {
	if File_proto_simple_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_simple_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_proto_simple_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResult); i {
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
			RawDescriptor: file_proto_simple_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_simple_calc_proto_goTypes,
		DependencyIndexes: file_proto_simple_calc_proto_depIdxs,
		MessageInfos:      file_proto_simple_calc_proto_msgTypes,
	}.Build()
	File_proto_simple_calc_proto = out.File
	file_proto_simple_calc_proto_rawDesc = nil
	file_proto_simple_calc_proto_goTypes = nil
	file_proto_simple_calc_proto_depIdxs = nil
}
