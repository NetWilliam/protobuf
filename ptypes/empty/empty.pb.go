// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/golang/protobuf/ptypes/empty/empty.proto

package empty

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	known "github.com/golang/protobuf/v2/types/known"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// Symbols defined in public import of google/protobuf/empty.proto

type Empty = known.Empty

var File_github_com_golang_protobuf_ptypes_empty_empty_proto protoreflect.FileDescriptor

var file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x3b, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescOnce sync.Once
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescData = file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDesc
)

func file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescGZIP() []byte {
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescOnce.Do(func() {
		file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescData)
	})
	return file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDescData
}

var file_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes = []interface{}{}
var file_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs = []int32{}

func init() { file_github_com_golang_protobuf_ptypes_empty_empty_proto_init() }
func file_github_com_golang_protobuf_ptypes_empty_empty_proto_init() {
	if File_github_com_golang_protobuf_ptypes_empty_empty_proto != nil {
		return
	}
	File_github_com_golang_protobuf_ptypes_empty_empty_proto = protoimpl.FileBuilder{
		RawDescriptor:     file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDesc,
		GoTypes:           file_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes,
		DependencyIndexes: file_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs,
		FilesRegistry:     protoregistry.GlobalFiles,
		TypesRegistry:     protoregistry.GlobalTypes,
	}.Init()
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_rawDesc = nil
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_goTypes = nil
	file_github_com_golang_protobuf_ptypes_empty_empty_proto_depIdxs = nil
}
