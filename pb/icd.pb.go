// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: icd.proto

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

type ICD10 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType            string `protobuf:"bytes,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	Icd10                 string `protobuf:"bytes,2,opt,name=icd10,proto3" json:"icd10,omitempty"`
	Icd10Desc             string `protobuf:"bytes,3,opt,name=icd10_desc,json=icd10Desc,proto3" json:"icd10_desc,omitempty"`
	Icd9                  string `protobuf:"bytes,4,opt,name=icd9,proto3" json:"icd9,omitempty"`
	Icd9Desc              string `protobuf:"bytes,5,opt,name=icd9_desc,json=icd9Desc,proto3" json:"icd9_desc,omitempty"`
	RecognizedMedicalText string `protobuf:"bytes,6,opt,name=recognized_medical_text,json=recognizedMedicalText,proto3" json:"recognized_medical_text,omitempty"`
}

func (x *ICD10) Reset() {
	*x = ICD10{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ICD10) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICD10) ProtoMessage() {}

func (x *ICD10) ProtoReflect() protoreflect.Message {
	mi := &file_icd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICD10.ProtoReflect.Descriptor instead.
func (*ICD10) Descriptor() ([]byte, []int) {
	return file_icd_proto_rawDescGZIP(), []int{0}
}

func (x *ICD10) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *ICD10) GetIcd10() string {
	if x != nil {
		return x.Icd10
	}
	return ""
}

func (x *ICD10) GetIcd10Desc() string {
	if x != nil {
		return x.Icd10Desc
	}
	return ""
}

func (x *ICD10) GetIcd9() string {
	if x != nil {
		return x.Icd9
	}
	return ""
}

func (x *ICD10) GetIcd9Desc() string {
	if x != nil {
		return x.Icd9Desc
	}
	return ""
}

func (x *ICD10) GetRecognizedMedicalText() string {
	if x != nil {
		return x.RecognizedMedicalText
	}
	return ""
}

var File_icd_proto protoreflect.FileDescriptor

var file_icd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0xc6, 0x01, 0x0a, 0x05, 0x49, 0x43, 0x44, 0x31, 0x30, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x63,
	0x64, 0x31, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x63, 0x64, 0x31, 0x30,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x63, 0x64, 0x31, 0x30, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x63, 0x64, 0x31, 0x30, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x64, 0x39, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x64, 0x39, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x63, 0x64, 0x39, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x63, 0x64, 0x39, 0x44, 0x65, 0x73, 0x63,
	0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x61, 0x69, 0x32, 0x30, 0x32,
	0x32, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_icd_proto_rawDescOnce sync.Once
	file_icd_proto_rawDescData = file_icd_proto_rawDesc
)

func file_icd_proto_rawDescGZIP() []byte {
	file_icd_proto_rawDescOnce.Do(func() {
		file_icd_proto_rawDescData = protoimpl.X.CompressGZIP(file_icd_proto_rawDescData)
	})
	return file_icd_proto_rawDescData
}

var file_icd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_icd_proto_goTypes = []interface{}{
	(*ICD10)(nil), // 0: pb.ICD10
}
var file_icd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_icd_proto_init() }
func file_icd_proto_init() {
	if File_icd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_icd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ICD10); i {
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
			RawDescriptor: file_icd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_icd_proto_goTypes,
		DependencyIndexes: file_icd_proto_depIdxs,
		MessageInfos:      file_icd_proto_msgTypes,
	}.Build()
	File_icd_proto = out.File
	file_icd_proto_rawDesc = nil
	file_icd_proto_goTypes = nil
	file_icd_proto_depIdxs = nil
}