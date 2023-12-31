// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: test/proto3pb/test_import.proto

package proto3pb

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

type ImportedGlobalEnum int32

const (
	ImportedGlobalEnum_IMPORT_FOO ImportedGlobalEnum = 0
	ImportedGlobalEnum_IMPORT_BAR ImportedGlobalEnum = 1
	ImportedGlobalEnum_IMPORT_BAZ ImportedGlobalEnum = 2
)

// Enum value maps for ImportedGlobalEnum.
var (
	ImportedGlobalEnum_name = map[int32]string{
		0: "IMPORT_FOO",
		1: "IMPORT_BAR",
		2: "IMPORT_BAZ",
	}
	ImportedGlobalEnum_value = map[string]int32{
		"IMPORT_FOO": 0,
		"IMPORT_BAR": 1,
		"IMPORT_BAZ": 2,
	}
)

func (x ImportedGlobalEnum) Enum() *ImportedGlobalEnum {
	p := new(ImportedGlobalEnum)
	*p = x
	return p
}

func (x ImportedGlobalEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportedGlobalEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto3pb_test_import_proto_enumTypes[0].Descriptor()
}

func (ImportedGlobalEnum) Type() protoreflect.EnumType {
	return &file_test_proto3pb_test_import_proto_enumTypes[0]
}

func (x ImportedGlobalEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportedGlobalEnum.Descriptor instead.
func (ImportedGlobalEnum) EnumDescriptor() ([]byte, []int) {
	return file_test_proto3pb_test_import_proto_rawDescGZIP(), []int{0}
}

var File_test_proto3pb_test_import_proto protoreflect.FileDescriptor

var file_test_proto3pb_test_import_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x70, 0x62, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2a, 0x44, 0x0a, 0x12, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x46, 0x4f, 0x4f, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x42, 0x41, 0x5a, 0x10, 0x02,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x65, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_test_proto3pb_test_import_proto_rawDescOnce sync.Once
	file_test_proto3pb_test_import_proto_rawDescData = file_test_proto3pb_test_import_proto_rawDesc
)

func file_test_proto3pb_test_import_proto_rawDescGZIP() []byte {
	file_test_proto3pb_test_import_proto_rawDescOnce.Do(func() {
		file_test_proto3pb_test_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto3pb_test_import_proto_rawDescData)
	})
	return file_test_proto3pb_test_import_proto_rawDescData
}

var file_test_proto3pb_test_import_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto3pb_test_import_proto_goTypes = []interface{}{
	(ImportedGlobalEnum)(0), // 0: google.expr.proto3.test.ImportedGlobalEnum
}
var file_test_proto3pb_test_import_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto3pb_test_import_proto_init() }
func file_test_proto3pb_test_import_proto_init() {
	if File_test_proto3pb_test_import_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto3pb_test_import_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto3pb_test_import_proto_goTypes,
		DependencyIndexes: file_test_proto3pb_test_import_proto_depIdxs,
		EnumInfos:         file_test_proto3pb_test_import_proto_enumTypes,
	}.Build()
	File_test_proto3pb_test_import_proto = out.File
	file_test_proto3pb_test_import_proto_rawDesc = nil
	file_test_proto3pb_test_import_proto_goTypes = nil
	file_test_proto3pb_test_import_proto_depIdxs = nil
}
