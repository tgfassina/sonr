// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: node/highway/v1/highway.proto

#include "node/highway/v1/highway.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG
namespace node {
namespace highway {
namespace v1 {
}  // namespace v1
}  // namespace highway
}  // namespace node
static constexpr ::PROTOBUF_NAMESPACE_ID::EnumDescriptor const** file_level_enum_descriptors_node_2fhighway_2fv1_2fhighway_2eproto = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::ServiceDescriptor const** file_level_service_descriptors_node_2fhighway_2fv1_2fhighway_2eproto = nullptr;
const uint32_t TableStruct_node_2fhighway_2fv1_2fhighway_2eproto::offsets[1] = {};
static constexpr ::PROTOBUF_NAMESPACE_ID::internal::MigrationSchema* schemas = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::Message* const* file_default_instances = nullptr;

const char descriptor_table_protodef_node_2fhighway_2fv1_2fhighway_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\035node/highway/v1/highway.proto\022\017node.hi"
  "ghway.v1\032\035node/highway/v1/request.proto\032"
  "\036node/highway/v1/response.proto2\362\005\n\016High"
  "wayService\022V\n\tListPeers\022!.node.highway.v"
  "1.ListPeersRequest\032\".node.highway.v1.Lis"
  "tPeersResponse\"\0000\001\022c\n\016DecideExchange\022&.n"
  "ode.highway.v1.DecideExchangeRequest\032\'.n"
  "ode.highway.v1.DecideExchangeResponse\"\000\022"
  "]\n\014SendExchange\022$.node.highway.v1.SendEx"
  "changeRequest\032%.node.highway.v1.SendExch"
  "angeResponse\"\000\022Z\n\013CacheRecord\022#.node.hig"
  "hway.v1.CacheRecordRequest\032$.node.highwa"
  "y.v1.CacheRecordResponse\"\000\022T\n\tGetRecord\022"
  "!.node.highway.v1.GetRecordRequest\032\".nod"
  "e.highway.v1.GetRecordResponse\"\000\022Z\n\013Stor"
  "eRecord\022#.node.highway.v1.StoreRecordReq"
  "uest\032$.node.highway.v1.StoreRecordRespon"
  "se\"\000\022]\n\014RegisterName\022$.node.highway.v1.R"
  "egisterNameRequest\032%.node.highway.v1.Reg"
  "isterNameResponse\"\000\022W\n\nVerifyName\022\".node"
  ".highway.v1.VerifyNameRequest\032#.node.hig"
  "hway.v1.VerifyNameResponse\"\000B&Z$github.c"
  "om/sonr-io/core/node/highwayb\006proto3"
  ;
static const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable*const descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto_deps[2] = {
  &::descriptor_table_node_2fhighway_2fv1_2frequest_2eproto,
  &::descriptor_table_node_2fhighway_2fv1_2fresponse_2eproto,
};
static ::PROTOBUF_NAMESPACE_ID::internal::once_flag descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto_once;
const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto = {
  false, false, 916, descriptor_table_protodef_node_2fhighway_2fv1_2fhighway_2eproto, "node/highway/v1/highway.proto", 
  &descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto_once, descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto_deps, 2, 0,
  schemas, file_default_instances, TableStruct_node_2fhighway_2fv1_2fhighway_2eproto::offsets,
  nullptr, file_level_enum_descriptors_node_2fhighway_2fv1_2fhighway_2eproto, file_level_service_descriptors_node_2fhighway_2fv1_2fhighway_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable* descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto_getter() {
  return &descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY static ::PROTOBUF_NAMESPACE_ID::internal::AddDescriptorsRunner dynamic_init_dummy_node_2fhighway_2fv1_2fhighway_2eproto(&descriptor_table_node_2fhighway_2fv1_2fhighway_2eproto);
namespace node {
namespace highway {
namespace v1 {

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace highway
}  // namespace node
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
