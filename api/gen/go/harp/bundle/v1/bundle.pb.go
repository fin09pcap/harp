// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: harp/bundle/v1/bundle.proto

package bundlev1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Bundle is a concrete secret bundle.
type Bundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata.
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Version of the file
	Version uint32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// Secret package collection
	Packages []*Package `protobuf:"bytes,4,rep,name=packages,proto3" json:"packages,omitempty"`
	// Bundle template object
	Template *Template `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
	// Associated values
	Values *wrapperspb.BytesValue `protobuf:"bytes,6,opt,name=values,proto3" json:"values,omitempty"`
	// Merkle Tree root
	MerkleTreeRoot []byte `protobuf:"bytes,7,opt,name=merkle_tree_root,json=merkleTreeRoot,proto3" json:"merkle_tree_root,omitempty"`
	// User data storage
	UserData map[string]*anypb.Any `protobuf:"bytes,99,rep,name=user_data,json=userData,proto3" json:"user_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Bundle) Reset() {
	*x = Bundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bundle) ProtoMessage() {}

func (x *Bundle) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bundle.ProtoReflect.Descriptor instead.
func (*Bundle) Descriptor() ([]byte, []int) {
	return file_harp_bundle_v1_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *Bundle) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Bundle) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Bundle) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Bundle) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *Bundle) GetTemplate() *Template {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *Bundle) GetValues() *wrapperspb.BytesValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Bundle) GetMerkleTreeRoot() []byte {
	if x != nil {
		return x.MerkleTreeRoot
	}
	return nil
}

func (x *Bundle) GetUserData() map[string]*anypb.Any {
	if x != nil {
		return x.UserData
	}
	return nil
}

// Package is a secret organizational unit.
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata.
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Package name as a complete secret path (CSO compliance recommended)
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Active secret version
	Secrets *SecretChain `protobuf:"bytes,4,opt,name=secrets,proto3" json:"secrets,omitempty"`
	// SecretChain versions
	Versions map[uint32]*SecretChain `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty" protobuf_key:"fixed32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User data storage
	UserData map[string]*anypb.Any `protobuf:"bytes,99,rep,name=user_data,json=userData,proto3" json:"user_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_harp_bundle_v1_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *Package) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Package) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetSecrets() *SecretChain {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *Package) GetVersions() map[uint32]*SecretChain {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *Package) GetUserData() map[string]*anypb.Any {
	if x != nil {
		return x.UserData
	}
	return nil
}

// SecretChain describe a secret version chain.
type SecretChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata.
	Annotations map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Version identifier
	Version uint32 `protobuf:"fixed32,3,opt,name=version,proto3" json:"version,omitempty"`
	// Secret K/V collection
	Data []*KV `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	// Link to previous version
	PreviousVersion *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=previous_version,json=previousVersion,proto3" json:"previous_version,omitempty"`
	// Link to next version
	NextVersion *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	// Locked buffer when encryption is enabled
	Locked *wrapperspb.BytesValue `protobuf:"bytes,7,opt,name=locked,proto3" json:"locked,omitempty"`
	// User data storage
	UserData map[string]*anypb.Any `protobuf:"bytes,99,rep,name=user_data,json=userData,proto3" json:"user_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecretChain) Reset() {
	*x = SecretChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretChain) ProtoMessage() {}

func (x *SecretChain) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretChain.ProtoReflect.Descriptor instead.
func (*SecretChain) Descriptor() ([]byte, []int) {
	return file_harp_bundle_v1_bundle_proto_rawDescGZIP(), []int{2}
}

func (x *SecretChain) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SecretChain) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *SecretChain) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SecretChain) GetData() []*KV {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SecretChain) GetPreviousVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PreviousVersion
	}
	return nil
}

func (x *SecretChain) GetNextVersion() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextVersion
	}
	return nil
}

func (x *SecretChain) GetLocked() *wrapperspb.BytesValue {
	if x != nil {
		return x.Locked
	}
	return nil
}

func (x *SecretChain) GetUserData() map[string]*anypb.Any {
	if x != nil {
		return x.UserData
	}
	return nil
}

// KV contains the key, the value and the type of the value.
type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Golang type of initial value before packing
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Value must be encoded using secret.Pack method
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harp_bundle_v1_bundle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_harp_bundle_v1_bundle_proto_msgTypes[3]
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
	return file_harp_bundle_v1_bundle_proto_rawDescGZIP(), []int{3}
}

func (x *KV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KV) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *KV) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_harp_bundle_v1_bundle_proto protoreflect.FileDescriptor

var file_harp_bundle_v1_bundle_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68,
	0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x68,
	0x61, 0x72, 0x70, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x05, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x49,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x61, 0x72,
	0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x74, 0x72,
	0x65, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x41, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x63, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c,
	0x05, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x61, 0x72,
	0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68,
	0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x41,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x42, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x63,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x58, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x05,
	0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3f, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4e,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x47, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6e,
	0x65, 0x78, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x46, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x63, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x9f, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x65, 0x63, 0x2e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x61, 0x72, 0x70, 0x2f, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x53, 0x42, 0x58, 0xaa, 0x02, 0x0e, 0x68, 0x61, 0x72, 0x70, 0x2e, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x68, 0x61, 0x72, 0x70, 0x5c,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_harp_bundle_v1_bundle_proto_rawDescOnce sync.Once
	file_harp_bundle_v1_bundle_proto_rawDescData = file_harp_bundle_v1_bundle_proto_rawDesc
)

func file_harp_bundle_v1_bundle_proto_rawDescGZIP() []byte {
	file_harp_bundle_v1_bundle_proto_rawDescOnce.Do(func() {
		file_harp_bundle_v1_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_harp_bundle_v1_bundle_proto_rawDescData)
	})
	return file_harp_bundle_v1_bundle_proto_rawDescData
}

var (
	file_harp_bundle_v1_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
	file_harp_bundle_v1_bundle_proto_goTypes  = []interface{}{
		(*Bundle)(nil),                 // 0: harp.bundle.v1.Bundle
		(*Package)(nil),                // 1: harp.bundle.v1.Package
		(*SecretChain)(nil),            // 2: harp.bundle.v1.SecretChain
		(*KV)(nil),                     // 3: harp.bundle.v1.KV
		nil,                            // 4: harp.bundle.v1.Bundle.LabelsEntry
		nil,                            // 5: harp.bundle.v1.Bundle.AnnotationsEntry
		nil,                            // 6: harp.bundle.v1.Bundle.UserDataEntry
		nil,                            // 7: harp.bundle.v1.Package.LabelsEntry
		nil,                            // 8: harp.bundle.v1.Package.AnnotationsEntry
		nil,                            // 9: harp.bundle.v1.Package.VersionsEntry
		nil,                            // 10: harp.bundle.v1.Package.UserDataEntry
		nil,                            // 11: harp.bundle.v1.SecretChain.LabelsEntry
		nil,                            // 12: harp.bundle.v1.SecretChain.AnnotationsEntry
		nil,                            // 13: harp.bundle.v1.SecretChain.UserDataEntry
		(*Template)(nil),               // 14: harp.bundle.v1.Template
		(*wrapperspb.BytesValue)(nil),  // 15: google.protobuf.BytesValue
		(*wrapperspb.UInt32Value)(nil), // 16: google.protobuf.UInt32Value
		(*anypb.Any)(nil),              // 17: google.protobuf.Any
	}
)
var file_harp_bundle_v1_bundle_proto_depIdxs = []int32{
	4,  // 0: harp.bundle.v1.Bundle.labels:type_name -> harp.bundle.v1.Bundle.LabelsEntry
	5,  // 1: harp.bundle.v1.Bundle.annotations:type_name -> harp.bundle.v1.Bundle.AnnotationsEntry
	1,  // 2: harp.bundle.v1.Bundle.packages:type_name -> harp.bundle.v1.Package
	14, // 3: harp.bundle.v1.Bundle.template:type_name -> harp.bundle.v1.Template
	15, // 4: harp.bundle.v1.Bundle.values:type_name -> google.protobuf.BytesValue
	6,  // 5: harp.bundle.v1.Bundle.user_data:type_name -> harp.bundle.v1.Bundle.UserDataEntry
	7,  // 6: harp.bundle.v1.Package.labels:type_name -> harp.bundle.v1.Package.LabelsEntry
	8,  // 7: harp.bundle.v1.Package.annotations:type_name -> harp.bundle.v1.Package.AnnotationsEntry
	2,  // 8: harp.bundle.v1.Package.secrets:type_name -> harp.bundle.v1.SecretChain
	9,  // 9: harp.bundle.v1.Package.versions:type_name -> harp.bundle.v1.Package.VersionsEntry
	10, // 10: harp.bundle.v1.Package.user_data:type_name -> harp.bundle.v1.Package.UserDataEntry
	11, // 11: harp.bundle.v1.SecretChain.labels:type_name -> harp.bundle.v1.SecretChain.LabelsEntry
	12, // 12: harp.bundle.v1.SecretChain.annotations:type_name -> harp.bundle.v1.SecretChain.AnnotationsEntry
	3,  // 13: harp.bundle.v1.SecretChain.data:type_name -> harp.bundle.v1.KV
	16, // 14: harp.bundle.v1.SecretChain.previous_version:type_name -> google.protobuf.UInt32Value
	16, // 15: harp.bundle.v1.SecretChain.next_version:type_name -> google.protobuf.UInt32Value
	15, // 16: harp.bundle.v1.SecretChain.locked:type_name -> google.protobuf.BytesValue
	13, // 17: harp.bundle.v1.SecretChain.user_data:type_name -> harp.bundle.v1.SecretChain.UserDataEntry
	17, // 18: harp.bundle.v1.Bundle.UserDataEntry.value:type_name -> google.protobuf.Any
	2,  // 19: harp.bundle.v1.Package.VersionsEntry.value:type_name -> harp.bundle.v1.SecretChain
	17, // 20: harp.bundle.v1.Package.UserDataEntry.value:type_name -> google.protobuf.Any
	17, // 21: harp.bundle.v1.SecretChain.UserDataEntry.value:type_name -> google.protobuf.Any
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_harp_bundle_v1_bundle_proto_init() }
func file_harp_bundle_v1_bundle_proto_init() {
	if File_harp_bundle_v1_bundle_proto != nil {
		return
	}
	file_harp_bundle_v1_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_harp_bundle_v1_bundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bundle); i {
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
		file_harp_bundle_v1_bundle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_harp_bundle_v1_bundle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretChain); i {
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
		file_harp_bundle_v1_bundle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_harp_bundle_v1_bundle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_harp_bundle_v1_bundle_proto_goTypes,
		DependencyIndexes: file_harp_bundle_v1_bundle_proto_depIdxs,
		MessageInfos:      file_harp_bundle_v1_bundle_proto_msgTypes,
	}.Build()
	File_harp_bundle_v1_bundle_proto = out.File
	file_harp_bundle_v1_bundle_proto_rawDesc = nil
	file_harp_bundle_v1_bundle_proto_goTypes = nil
	file_harp_bundle_v1_bundle_proto_depIdxs = nil
}
