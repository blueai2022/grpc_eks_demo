// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: rpc_recognize_icd.proto

package pb

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

type RecognizeICD10Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedicalText string `protobuf:"bytes,1,opt,name=medical_text,json=medicalText,proto3" json:"medical_text,omitempty"`
}

func (x *RecognizeICD10Request) Reset() {
	*x = RecognizeICD10Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recognize_icd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeICD10Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeICD10Request) ProtoMessage() {}

func (x *RecognizeICD10Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recognize_icd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeICD10Request.ProtoReflect.Descriptor instead.
func (*RecognizeICD10Request) Descriptor() ([]byte, []int) {
	return file_rpc_recognize_icd_proto_rawDescGZIP(), []int{0}
}

func (x *RecognizeICD10Request) GetMedicalText() string {
	if x != nil {
		return x.MedicalText
	}
	return ""
}

type RecognizeICD10Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Result  *ICD10 `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RecognizeICD10Response) Reset() {
	*x = RecognizeICD10Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_recognize_icd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeICD10Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeICD10Response) ProtoMessage() {}

func (x *RecognizeICD10Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_recognize_icd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeICD10Response.ProtoReflect.Descriptor instead.
func (*RecognizeICD10Response) Descriptor() ([]byte, []int) {
	return file_rpc_recognize_icd_proto_rawDescGZIP(), []int{1}
}

func (x *RecognizeICD10Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RecognizeICD10Response) GetResult() *ICD10 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_rpc_recognize_icd_proto protoreflect.FileDescriptor

var file_rpc_recognize_icd_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x5f,
	0x69, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x09, 0x69,
	0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f,
	0x67, 0x6e, 0x69, 0x7a, 0x65, 0x49, 0x43, 0x44, 0x31, 0x30, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x6c,
	0x54, 0x65, 0x78, 0x74, 0x22, 0x55, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a,
	0x65, 0x49, 0x43, 0x44, 0x31, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x43,
	0x44, 0x31, 0x30, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x69,
	0x32, 0x30, 0x32, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_recognize_icd_proto_rawDescOnce sync.Once
	file_rpc_recognize_icd_proto_rawDescData = file_rpc_recognize_icd_proto_rawDesc
)

func file_rpc_recognize_icd_proto_rawDescGZIP() []byte {
	file_rpc_recognize_icd_proto_rawDescOnce.Do(func() {
		file_rpc_recognize_icd_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_recognize_icd_proto_rawDescData)
	})
	return file_rpc_recognize_icd_proto_rawDescData
}

var file_rpc_recognize_icd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_recognize_icd_proto_goTypes = []interface{}{
	(*RecognizeICD10Request)(nil),  // 0: pb.RecognizeICD10Request
	(*RecognizeICD10Response)(nil), // 1: pb.RecognizeICD10Response
	(*ICD10)(nil),                  // 2: pb.ICD10
}
var file_rpc_recognize_icd_proto_depIdxs = []int32{
	2, // 0: pb.RecognizeICD10Response.result:type_name -> pb.ICD10
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_recognize_icd_proto_init() }
func file_rpc_recognize_icd_proto_init() {
	if File_rpc_recognize_icd_proto != nil {
		return
	}
	file_icd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_recognize_icd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeICD10Request); i {
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
		file_rpc_recognize_icd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeICD10Response); i {
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
			RawDescriptor: file_rpc_recognize_icd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_recognize_icd_proto_goTypes,
		DependencyIndexes: file_rpc_recognize_icd_proto_depIdxs,
		MessageInfos:      file_rpc_recognize_icd_proto_msgTypes,
	}.Build()
	File_rpc_recognize_icd_proto = out.File
	file_rpc_recognize_icd_proto_rawDesc = nil
	file_rpc_recognize_icd_proto_goTypes = nil
	file_rpc_recognize_icd_proto_depIdxs = nil
}