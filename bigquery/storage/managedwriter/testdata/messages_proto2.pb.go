// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.1
// source: messages_proto2.proto

package testdata

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// SimpleMessage represents a simple message that transmits a string and int64 value.
type SimpleMessageProto2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is a simple scalar string.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// value is a simple int64 value.
	Value *int64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (x *SimpleMessageProto2) Reset() {
	*x = SimpleMessageProto2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessageProto2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessageProto2) ProtoMessage() {}

func (x *SimpleMessageProto2) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessageProto2.ProtoReflect.Descriptor instead.
func (*SimpleMessageProto2) Descriptor() ([]byte, []int) {
	return file_messages_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMessageProto2) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SimpleMessageProto2) GetValue() int64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

type GithubArchiveEntityProto2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Login      *string `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
	GravatarId *string `protobuf:"bytes,3,opt,name=gravatar_id,json=gravatarId" json:"gravatar_id,omitempty"`
	AvatarUrl  *string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
	Url        *string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
}

func (x *GithubArchiveEntityProto2) Reset() {
	*x = GithubArchiveEntityProto2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubArchiveEntityProto2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubArchiveEntityProto2) ProtoMessage() {}

func (x *GithubArchiveEntityProto2) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubArchiveEntityProto2.ProtoReflect.Descriptor instead.
func (*GithubArchiveEntityProto2) Descriptor() ([]byte, []int) {
	return file_messages_proto2_proto_rawDescGZIP(), []int{1}
}

func (x *GithubArchiveEntityProto2) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GithubArchiveEntityProto2) GetLogin() string {
	if x != nil && x.Login != nil {
		return *x.Login
	}
	return ""
}

func (x *GithubArchiveEntityProto2) GetGravatarId() string {
	if x != nil && x.GravatarId != nil {
		return *x.GravatarId
	}
	return ""
}

func (x *GithubArchiveEntityProto2) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

func (x *GithubArchiveEntityProto2) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type GithubArchiveRepoProto2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url  *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (x *GithubArchiveRepoProto2) Reset() {
	*x = GithubArchiveRepoProto2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubArchiveRepoProto2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubArchiveRepoProto2) ProtoMessage() {}

func (x *GithubArchiveRepoProto2) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubArchiveRepoProto2.ProtoReflect.Descriptor instead.
func (*GithubArchiveRepoProto2) Descriptor() ([]byte, []int) {
	return file_messages_proto2_proto_rawDescGZIP(), []int{2}
}

func (x *GithubArchiveRepoProto2) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GithubArchiveRepoProto2) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GithubArchiveRepoProto2) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

// GithubArchiveMessageProto2 is the proto2 version of github archive row.
type GithubArchiveMessageProto2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      *string                    `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Public    *bool                      `protobuf:"varint,2,opt,name=public" json:"public,omitempty"`
	Payload   *string                    `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	Repo      *GithubArchiveRepoProto2   `protobuf:"bytes,4,opt,name=repo" json:"repo,omitempty"`
	Actor     *GithubArchiveEntityProto2 `protobuf:"bytes,5,opt,name=actor" json:"actor,omitempty"`
	Org       *GithubArchiveEntityProto2 `protobuf:"bytes,6,opt,name=org" json:"org,omitempty"`
	CreatedAt *int64                     `protobuf:"varint,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Id        *string                    `protobuf:"bytes,8,opt,name=id" json:"id,omitempty"`
	Other     *string                    `protobuf:"bytes,9,opt,name=other" json:"other,omitempty"`
}

func (x *GithubArchiveMessageProto2) Reset() {
	*x = GithubArchiveMessageProto2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubArchiveMessageProto2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubArchiveMessageProto2) ProtoMessage() {}

func (x *GithubArchiveMessageProto2) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubArchiveMessageProto2.ProtoReflect.Descriptor instead.
func (*GithubArchiveMessageProto2) Descriptor() ([]byte, []int) {
	return file_messages_proto2_proto_rawDescGZIP(), []int{3}
}

func (x *GithubArchiveMessageProto2) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *GithubArchiveMessageProto2) GetPublic() bool {
	if x != nil && x.Public != nil {
		return *x.Public
	}
	return false
}

func (x *GithubArchiveMessageProto2) GetPayload() string {
	if x != nil && x.Payload != nil {
		return *x.Payload
	}
	return ""
}

func (x *GithubArchiveMessageProto2) GetRepo() *GithubArchiveRepoProto2 {
	if x != nil {
		return x.Repo
	}
	return nil
}

func (x *GithubArchiveMessageProto2) GetActor() *GithubArchiveEntityProto2 {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *GithubArchiveMessageProto2) GetOrg() *GithubArchiveEntityProto2 {
	if x != nil {
		return x.Org
	}
	return nil
}

func (x *GithubArchiveMessageProto2) GetCreatedAt() int64 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *GithubArchiveMessageProto2) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *GithubArchiveMessageProto2) GetOther() string {
	if x != nil && x.Other != nil {
		return *x.Other
	}
	return ""
}

var File_messages_proto2_proto protoreflect.FileDescriptor

var file_messages_proto2_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x3f, 0x0a, 0x13, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x19, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x4f, 0x0a, 0x17, 0x47, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0xd0, 0x02, 0x0a, 0x1a, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x35,
	0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x52,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x35, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x3d, 0x5a, 0x3b,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_messages_proto2_proto_rawDescOnce sync.Once
	file_messages_proto2_proto_rawDescData = file_messages_proto2_proto_rawDesc
)

func file_messages_proto2_proto_rawDescGZIP() []byte {
	file_messages_proto2_proto_rawDescOnce.Do(func() {
		file_messages_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto2_proto_rawDescData)
	})
	return file_messages_proto2_proto_rawDescData
}

var file_messages_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto2_proto_goTypes = []interface{}{
	(*SimpleMessageProto2)(nil),        // 0: testdata.SimpleMessageProto2
	(*GithubArchiveEntityProto2)(nil),  // 1: testdata.GithubArchiveEntityProto2
	(*GithubArchiveRepoProto2)(nil),    // 2: testdata.GithubArchiveRepoProto2
	(*GithubArchiveMessageProto2)(nil), // 3: testdata.GithubArchiveMessageProto2
}
var file_messages_proto2_proto_depIdxs = []int32{
	2, // 0: testdata.GithubArchiveMessageProto2.repo:type_name -> testdata.GithubArchiveRepoProto2
	1, // 1: testdata.GithubArchiveMessageProto2.actor:type_name -> testdata.GithubArchiveEntityProto2
	1, // 2: testdata.GithubArchiveMessageProto2.org:type_name -> testdata.GithubArchiveEntityProto2
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_messages_proto2_proto_init() }
func file_messages_proto2_proto_init() {
	if File_messages_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessageProto2); i {
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
		file_messages_proto2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubArchiveEntityProto2); i {
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
		file_messages_proto2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubArchiveRepoProto2); i {
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
		file_messages_proto2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubArchiveMessageProto2); i {
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
			RawDescriptor: file_messages_proto2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto2_proto_goTypes,
		DependencyIndexes: file_messages_proto2_proto_depIdxs,
		MessageInfos:      file_messages_proto2_proto_msgTypes,
	}.Build()
	File_messages_proto2_proto = out.File
	file_messages_proto2_proto_rawDesc = nil
	file_messages_proto2_proto_goTypes = nil
	file_messages_proto2_proto_depIdxs = nil
}
