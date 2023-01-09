// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: substreams/sink/kv/v1/read.proto

package kvv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key to fetch
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keys to fetch
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *GetManyRequest) Reset() {
	*x = GetManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyRequest) ProtoMessage() {}

func (x *GetManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyRequest.ProtoReflect.Descriptor instead.
func (*GetManyRequest) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{1}
}

func (x *GetManyRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type GetByPrefixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// server may impose a hard limit, trying to go above it would return grpc_error: INVALID_ARGUMENT
	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// requested prefix
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *GetByPrefixRequest) Reset() {
	*x = GetByPrefixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPrefixRequest) ProtoMessage() {}

func (x *GetByPrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPrefixRequest.ProtoReflect.Descriptor instead.
func (*GetByPrefixRequest) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{2}
}

func (x *GetByPrefixRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetByPrefixRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// server may impose a hard limit, trying to go above it would return grpc_error: INVALID_ARGUMENT
	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// scanning will start at this point, lexicographically
	BeginPrefix string `protobuf:"bytes,2,opt,name=begin_prefix,json=beginPrefix,proto3" json:"begin_prefix,omitempty"`
	// If set, scanning will stop when it reaches this point or above, excluding this exact key
	ExclusiveEndPrefix *string `protobuf:"bytes,3,opt,name=exclusive_end_prefix,json=exclusiveEndPrefix,proto3,oneof" json:"exclusive_end_prefix,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{3}
}

func (x *ScanRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ScanRequest) GetBeginPrefix() string {
	if x != nil {
		return x.BeginPrefix
	}
	return ""
}

func (x *ScanRequest) GetExclusiveEndPrefix() string {
	if x != nil && x.ExclusiveEndPrefix != nil {
		return *x.ExclusiveEndPrefix
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value that was found for the requested key
	Value *anypb.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type GetManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Values that were found for the requested keys
	Value []*anypb.Any `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *GetManyResponse) Reset() {
	*x = GetManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyResponse) ProtoMessage() {}

func (x *GetManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyResponse.ProtoReflect.Descriptor instead.
func (*GetManyResponse) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{5}
}

func (x *GetManyResponse) GetValue() []*anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type GetByPrefixResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KV are the key/value pairs that were found with the given prefix
	KeyValues []*KV `protobuf:"bytes,1,rep,name=key_values,json=keyValues,proto3" json:"key_values,omitempty"`
	// limit_reached is true if there is at least ONE MORE result than the requested limit
	LimitReached bool `protobuf:"varint,2,opt,name=limit_reached,json=limitReached,proto3" json:"limit_reached,omitempty"`
}

func (x *GetByPrefixResponse) Reset() {
	*x = GetByPrefixResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPrefixResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPrefixResponse) ProtoMessage() {}

func (x *GetByPrefixResponse) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPrefixResponse.ProtoReflect.Descriptor instead.
func (*GetByPrefixResponse) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{6}
}

func (x *GetByPrefixResponse) GetKeyValues() []*KV {
	if x != nil {
		return x.KeyValues
	}
	return nil
}

func (x *GetByPrefixResponse) GetLimitReached() bool {
	if x != nil {
		return x.LimitReached
	}
	return false
}

type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// KV are the key/value pairs that were found during scan
	KeyValues []*KV `protobuf:"bytes,1,rep,name=key_values,json=keyValues,proto3" json:"key_values,omitempty"`
	// limit_reached is true if there is at least ONE MORE result than the requested limit
	LimitReached bool `protobuf:"varint,2,opt,name=limit_reached,json=limitReached,proto3" json:"limit_reached,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{7}
}

func (x *ScanResponse) GetKeyValues() []*KV {
	if x != nil {
		return x.KeyValues
	}
	return nil
}

func (x *ScanResponse) GetLimitReached() bool {
	if x != nil {
		return x.LimitReached
	}
	return false
}

type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_substreams_sink_kv_v1_read_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KV.ProtoReflect.Descriptor instead.
func (*KV) Descriptor() ([]byte, []int) {
	return file_substreams_sink_kv_v1_read_proto_rawDescGZIP(), []int{8}
}

func (x *KV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KV) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_substreams_sink_kv_v1_read_proto protoreflect.FileDescriptor

var file_substreams_sink_kv_v1_read_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6e,
	0x6b, 0x2f, 0x6b, 0x76, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x96,
	0x01, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x35, 0x0a, 0x14, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x76, 0x65, 0x45, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x17,
	0x0a, 0x15, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b,
	0x76, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x22, 0x6d, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xe3, 0x02, 0x0a, 0x02, 0x4b,
	0x76, 0x12, 0x4c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73,
	0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x29, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x22, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xe8, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x76, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x2d, 0x73, 0x69, 0x6e, 0x6b, 0x2d, 0x6b, 0x76, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x2f, 0x6b, 0x76, 0x2f, 0x76,
	0x31, 0x3b, 0x6b, 0x76, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x4b, 0xaa, 0x02, 0x15, 0x53,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x4b,
	0x76, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x5c, 0x53, 0x69, 0x6e, 0x6b, 0x5c, 0x4b, 0x76, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x53,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x5c, 0x53, 0x69, 0x6e, 0x6b, 0x5c, 0x4b,
	0x76, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x18, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x3a, 0x3a, 0x53,
	0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x4b, 0x76, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_substreams_sink_kv_v1_read_proto_rawDescOnce sync.Once
	file_substreams_sink_kv_v1_read_proto_rawDescData = file_substreams_sink_kv_v1_read_proto_rawDesc
)

func file_substreams_sink_kv_v1_read_proto_rawDescGZIP() []byte {
	file_substreams_sink_kv_v1_read_proto_rawDescOnce.Do(func() {
		file_substreams_sink_kv_v1_read_proto_rawDescData = protoimpl.X.CompressGZIP(file_substreams_sink_kv_v1_read_proto_rawDescData)
	})
	return file_substreams_sink_kv_v1_read_proto_rawDescData
}

var file_substreams_sink_kv_v1_read_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_substreams_sink_kv_v1_read_proto_goTypes = []interface{}{
	(*GetRequest)(nil),          // 0: substreams.sink.kv.v1.GetRequest
	(*GetManyRequest)(nil),      // 1: substreams.sink.kv.v1.GetManyRequest
	(*GetByPrefixRequest)(nil),  // 2: substreams.sink.kv.v1.GetByPrefixRequest
	(*ScanRequest)(nil),         // 3: substreams.sink.kv.v1.ScanRequest
	(*GetResponse)(nil),         // 4: substreams.sink.kv.v1.GetResponse
	(*GetManyResponse)(nil),     // 5: substreams.sink.kv.v1.GetManyResponse
	(*GetByPrefixResponse)(nil), // 6: substreams.sink.kv.v1.GetByPrefixResponse
	(*ScanResponse)(nil),        // 7: substreams.sink.kv.v1.ScanResponse
	(*KV)(nil),                  // 8: substreams.sink.kv.v1.KV
	(*anypb.Any)(nil),           // 9: google.protobuf.Any
}
var file_substreams_sink_kv_v1_read_proto_depIdxs = []int32{
	9, // 0: substreams.sink.kv.v1.GetResponse.value:type_name -> google.protobuf.Any
	9, // 1: substreams.sink.kv.v1.GetManyResponse.value:type_name -> google.protobuf.Any
	8, // 2: substreams.sink.kv.v1.GetByPrefixResponse.key_values:type_name -> substreams.sink.kv.v1.KV
	8, // 3: substreams.sink.kv.v1.ScanResponse.key_values:type_name -> substreams.sink.kv.v1.KV
	9, // 4: substreams.sink.kv.v1.KV.value:type_name -> google.protobuf.Any
	0, // 5: substreams.sink.kv.v1.Kv.Get:input_type -> substreams.sink.kv.v1.GetRequest
	1, // 6: substreams.sink.kv.v1.Kv.GetMany:input_type -> substreams.sink.kv.v1.GetManyRequest
	2, // 7: substreams.sink.kv.v1.Kv.GetByPrefix:input_type -> substreams.sink.kv.v1.GetByPrefixRequest
	3, // 8: substreams.sink.kv.v1.Kv.Scan:input_type -> substreams.sink.kv.v1.ScanRequest
	4, // 9: substreams.sink.kv.v1.Kv.Get:output_type -> substreams.sink.kv.v1.GetResponse
	5, // 10: substreams.sink.kv.v1.Kv.GetMany:output_type -> substreams.sink.kv.v1.GetManyResponse
	6, // 11: substreams.sink.kv.v1.Kv.GetByPrefix:output_type -> substreams.sink.kv.v1.GetByPrefixResponse
	7, // 12: substreams.sink.kv.v1.Kv.Scan:output_type -> substreams.sink.kv.v1.ScanResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_substreams_sink_kv_v1_read_proto_init() }
func file_substreams_sink_kv_v1_read_proto_init() {
	if File_substreams_sink_kv_v1_read_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_substreams_sink_kv_v1_read_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyRequest); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPrefixRequest); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyResponse); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPrefixResponse); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
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
		file_substreams_sink_kv_v1_read_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KV); i {
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
	file_substreams_sink_kv_v1_read_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_substreams_sink_kv_v1_read_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_substreams_sink_kv_v1_read_proto_goTypes,
		DependencyIndexes: file_substreams_sink_kv_v1_read_proto_depIdxs,
		MessageInfos:      file_substreams_sink_kv_v1_read_proto_msgTypes,
	}.Build()
	File_substreams_sink_kv_v1_read_proto = out.File
	file_substreams_sink_kv_v1_read_proto_rawDesc = nil
	file_substreams_sink_kv_v1_read_proto_goTypes = nil
	file_substreams_sink_kv_v1_read_proto_depIdxs = nil
}
