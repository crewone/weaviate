// Code generated by protoc-gen-go. DO NOT EDIT.

package protocol

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v1_weaviate_proto protoreflect.FileDescriptor

var file_v1_weaviate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x12, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x8a, 0x03, 0x0a, 0x08, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x09, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x6a, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x57, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x69,
	0x61, 0x74, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_v1_weaviate_proto_goTypes = []any{
	(*SearchRequest)(nil),       // 0: weaviate.v1.SearchRequest
	(*BatchObjectsRequest)(nil), // 1: weaviate.v1.BatchObjectsRequest
	(*BatchDeleteRequest)(nil),  // 2: weaviate.v1.BatchDeleteRequest
	(*TenantsGetRequest)(nil),   // 3: weaviate.v1.TenantsGetRequest
	(*AggregateRequest)(nil),    // 4: weaviate.v1.AggregateRequest
	(*SearchReply)(nil),         // 5: weaviate.v1.SearchReply
	(*BatchObjectsReply)(nil),   // 6: weaviate.v1.BatchObjectsReply
	(*BatchDeleteReply)(nil),    // 7: weaviate.v1.BatchDeleteReply
	(*TenantsGetReply)(nil),     // 8: weaviate.v1.TenantsGetReply
	(*AggregateReply)(nil),      // 9: weaviate.v1.AggregateReply
}
var file_v1_weaviate_proto_depIdxs = []int32{
	0, // 0: weaviate.v1.Weaviate.Search:input_type -> weaviate.v1.SearchRequest
	1, // 1: weaviate.v1.Weaviate.BatchObjects:input_type -> weaviate.v1.BatchObjectsRequest
	2, // 2: weaviate.v1.Weaviate.BatchDelete:input_type -> weaviate.v1.BatchDeleteRequest
	3, // 3: weaviate.v1.Weaviate.TenantsGet:input_type -> weaviate.v1.TenantsGetRequest
	4, // 4: weaviate.v1.Weaviate.Aggregate:input_type -> weaviate.v1.AggregateRequest
	5, // 5: weaviate.v1.Weaviate.Search:output_type -> weaviate.v1.SearchReply
	6, // 6: weaviate.v1.Weaviate.BatchObjects:output_type -> weaviate.v1.BatchObjectsReply
	7, // 7: weaviate.v1.Weaviate.BatchDelete:output_type -> weaviate.v1.BatchDeleteReply
	8, // 8: weaviate.v1.Weaviate.TenantsGet:output_type -> weaviate.v1.TenantsGetReply
	9, // 9: weaviate.v1.Weaviate.Aggregate:output_type -> weaviate.v1.AggregateReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_weaviate_proto_init() }
func file_v1_weaviate_proto_init() {
	if File_v1_weaviate_proto != nil {
		return
	}
	file_v1_aggregate_proto_init()
	file_v1_batch_proto_init()
	file_v1_batch_delete_proto_init()
	file_v1_search_get_proto_init()
	file_v1_tenants_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_weaviate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_weaviate_proto_goTypes,
		DependencyIndexes: file_v1_weaviate_proto_depIdxs,
	}.Build()
	File_v1_weaviate_proto = out.File
	file_v1_weaviate_proto_rawDesc = nil
	file_v1_weaviate_proto_goTypes = nil
	file_v1_weaviate_proto_depIdxs = nil
}
