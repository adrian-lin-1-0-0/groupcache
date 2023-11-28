//
//Copyright 2012 Google Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: test.proto

package testpb_v1

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

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	City *string `protobuf:"bytes,2,opt,name=city,proto3,oneof" json:"city,omitempty"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *TestMessage) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lower       string `protobuf:"bytes,1,opt,name=lower,proto3" json:"lower,omitempty"`                                       // to be returned upper case
	RepeatCount *int32 `protobuf:"varint,2,opt,name=repeat_count,json=repeatCount,proto3,oneof" json:"repeat_count,omitempty"` // .. this many times
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestRequest) GetLower() string {
	if x != nil {
		return x.Lower
	}
	return ""
}

func (x *TestRequest) GetRepeatCount() int32 {
	if x != nil && x.RepeatCount != nil {
		return *x.RepeatCount
	}
	return 0
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *string `protobuf:"bytes,1,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestResponse) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type CacheStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  *int64 `protobuf:"varint,1,opt,name=items,proto3,oneof" json:"items,omitempty"`
	Bytes  *int64 `protobuf:"varint,2,opt,name=bytes,proto3,oneof" json:"bytes,omitempty"`
	Gets   *int64 `protobuf:"varint,3,opt,name=gets,proto3,oneof" json:"gets,omitempty"`
	Hits   *int64 `protobuf:"varint,4,opt,name=hits,proto3,oneof" json:"hits,omitempty"`
	Evicts *int64 `protobuf:"varint,5,opt,name=evicts,proto3,oneof" json:"evicts,omitempty"`
}

func (x *CacheStats) Reset() {
	*x = CacheStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheStats) ProtoMessage() {}

func (x *CacheStats) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheStats.ProtoReflect.Descriptor instead.
func (*CacheStats) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *CacheStats) GetItems() int64 {
	if x != nil && x.Items != nil {
		return *x.Items
	}
	return 0
}

func (x *CacheStats) GetBytes() int64 {
	if x != nil && x.Bytes != nil {
		return *x.Bytes
	}
	return 0
}

func (x *CacheStats) GetGets() int64 {
	if x != nil && x.Gets != nil {
		return *x.Gets
	}
	return 0
}

func (x *CacheStats) GetHits() int64 {
	if x != nil && x.Hits != nil {
		return *x.Hits
	}
	return 0
}

func (x *CacheStats) GetEvicts() int64 {
	if x != nil && x.Evicts != nil {
		return *x.Evicts
	}
	return 0
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gets       *int64      `protobuf:"varint,1,opt,name=gets,proto3,oneof" json:"gets,omitempty"`
	CacheHits  *int64      `protobuf:"varint,12,opt,name=cache_hits,json=cacheHits,proto3,oneof" json:"cache_hits,omitempty"`
	Fills      *int64      `protobuf:"varint,2,opt,name=fills,proto3,oneof" json:"fills,omitempty"`
	TotalAlloc *uint64     `protobuf:"varint,3,opt,name=total_alloc,json=totalAlloc,proto3,oneof" json:"total_alloc,omitempty"`
	MainCache  *CacheStats `protobuf:"bytes,4,opt,name=main_cache,json=mainCache,proto3,oneof" json:"main_cache,omitempty"`
	HotCache   *CacheStats `protobuf:"bytes,5,opt,name=hot_cache,json=hotCache,proto3,oneof" json:"hot_cache,omitempty"`
	ServerIn   *int64      `protobuf:"varint,6,opt,name=server_in,json=serverIn,proto3,oneof" json:"server_in,omitempty"`
	Loads      *int64      `protobuf:"varint,8,opt,name=loads,proto3,oneof" json:"loads,omitempty"`
	PeerLoads  *int64      `protobuf:"varint,9,opt,name=peer_loads,json=peerLoads,proto3,oneof" json:"peer_loads,omitempty"`
	PeerErrors *int64      `protobuf:"varint,10,opt,name=peer_errors,json=peerErrors,proto3,oneof" json:"peer_errors,omitempty"`
	LocalLoads *int64      `protobuf:"varint,11,opt,name=local_loads,json=localLoads,proto3,oneof" json:"local_loads,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *StatsResponse) GetGets() int64 {
	if x != nil && x.Gets != nil {
		return *x.Gets
	}
	return 0
}

func (x *StatsResponse) GetCacheHits() int64 {
	if x != nil && x.CacheHits != nil {
		return *x.CacheHits
	}
	return 0
}

func (x *StatsResponse) GetFills() int64 {
	if x != nil && x.Fills != nil {
		return *x.Fills
	}
	return 0
}

func (x *StatsResponse) GetTotalAlloc() uint64 {
	if x != nil && x.TotalAlloc != nil {
		return *x.TotalAlloc
	}
	return 0
}

func (x *StatsResponse) GetMainCache() *CacheStats {
	if x != nil {
		return x.MainCache
	}
	return nil
}

func (x *StatsResponse) GetHotCache() *CacheStats {
	if x != nil {
		return x.HotCache
	}
	return nil
}

func (x *StatsResponse) GetServerIn() int64 {
	if x != nil && x.ServerIn != nil {
		return *x.ServerIn
	}
	return 0
}

func (x *StatsResponse) GetLoads() int64 {
	if x != nil && x.Loads != nil {
		return *x.Loads
	}
	return 0
}

func (x *StatsResponse) GetPeerLoads() int64 {
	if x != nil && x.PeerLoads != nil {
		return *x.PeerLoads
	}
	return 0
}

func (x *StatsResponse) GetPeerErrors() int64 {
	if x != nil && x.PeerErrors != nil {
		return *x.PeerErrors
	}
	return 0
}

func (x *StatsResponse) GetLocalLoads() int64 {
	if x != nil && x.LocalLoads != nil {
		return *x.LocalLoads
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x22,
	0x5c, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x67, 0x65, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x04, 0x67, 0x65, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03,
	0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x65, 0x76, 0x69,
	0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x06, 0x65, 0x76, 0x69,
	0x63, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x67,
	0x65, 0x74, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x68, 0x69, 0x74, 0x73, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x73, 0x22, 0xb0, 0x04, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04, 0x67, 0x65, 0x74, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x48,
	0x69, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x6c, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x48, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x69, 0x6e,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x09, 0x68, 0x6f, 0x74, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x48, 0x05, 0x52, 0x08, 0x68, 0x6f, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x09, 0x70, 0x65, 0x65, 0x72,
	0x4c, 0x6f, 0x61, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x09, 0x52,
	0x0a, 0x70, 0x65, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x0a, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4c, 0x6f, 0x61, 0x64,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x67, 0x65, 0x74, 0x73, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x68, 0x69, 0x74, 0x73, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x68, 0x6f, 0x74, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x70, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x7b, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_test_proto_goTypes = []any{
	(*TestMessage)(nil),   // 0: TestMessage
	(*TestRequest)(nil),   // 1: TestRequest
	(*TestResponse)(nil),  // 2: TestResponse
	(*CacheStats)(nil),    // 3: CacheStats
	(*StatsResponse)(nil), // 4: StatsResponse
	(*Empty)(nil),         // 5: Empty
}
var file_test_proto_depIdxs = []int32{
	3, // 0: StatsResponse.main_cache:type_name -> CacheStats
	3, // 1: StatsResponse.hot_cache:type_name -> CacheStats
	5, // 2: GroupCacheTest.InitPeers:input_type -> Empty
	1, // 3: GroupCacheTest.Get:input_type -> TestRequest
	5, // 4: GroupCacheTest.GetStats:input_type -> Empty
	5, // 5: GroupCacheTest.InitPeers:output_type -> Empty
	2, // 6: GroupCacheTest.Get:output_type -> TestResponse
	4, // 7: GroupCacheTest.GetStats:output_type -> StatsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TestMessage); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TestRequest); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TestResponse); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CacheStats); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StatsResponse); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
	file_test_proto_msgTypes[0].OneofWrappers = []any{}
	file_test_proto_msgTypes[1].OneofWrappers = []any{}
	file_test_proto_msgTypes[2].OneofWrappers = []any{}
	file_test_proto_msgTypes[3].OneofWrappers = []any{}
	file_test_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
