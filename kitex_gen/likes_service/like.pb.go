// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: idl/like.proto

package likes_service

import (
	context "context"
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

// 用户点赞，http操作
type AddLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *AddLikesRequest) Reset() {
	*x = AddLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLikesRequest) ProtoMessage() {}

func (x *AddLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLikesRequest.ProtoReflect.Descriptor instead.
func (*AddLikesRequest) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{0}
}

func (x *AddLikesRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddLikesRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type AddLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddLikesResponse) Reset() {
	*x = AddLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLikesResponse) ProtoMessage() {}

func (x *AddLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLikesResponse.ProtoReflect.Descriptor instead.
func (*AddLikesResponse) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{1}
}

func (x *AddLikesResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddLikesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 用户取消点赞，http操作
type DelLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *DelLikesRequest) Reset() {
	*x = DelLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLikesRequest) ProtoMessage() {}

func (x *DelLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLikesRequest.ProtoReflect.Descriptor instead.
func (*DelLikesRequest) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{2}
}

func (x *DelLikesRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DelLikesRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type DelLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DelLikesResponse) Reset() {
	*x = DelLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelLikesResponse) ProtoMessage() {}

func (x *DelLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelLikesResponse.ProtoReflect.Descriptor instead.
func (*DelLikesResponse) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{3}
}

func (x *DelLikesResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DelLikesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 查询帖子的时候使用GetLikes
type GetLikesCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetLikesCountRequest) Reset() {
	*x = GetLikesCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikesCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikesCountRequest) ProtoMessage() {}

func (x *GetLikesCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikesCountRequest.ProtoReflect.Descriptor instead.
func (*GetLikesCountRequest) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{4}
}

func (x *GetLikesCountRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type GetLikesCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Count  int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetLikesCountResponse) Reset() {
	*x = GetLikesCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikesCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikesCountResponse) ProtoMessage() {}

func (x *GetLikesCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikesCountResponse.ProtoReflect.Descriptor instead.
func (*GetLikesCountResponse) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{5}
}

func (x *GetLikesCountResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetLikesCountResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetLikesCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type IsLikedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *IsLikedRequest) Reset() {
	*x = IsLikedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikedRequest) ProtoMessage() {}

func (x *IsLikedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikedRequest.ProtoReflect.Descriptor instead.
func (*IsLikedRequest) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{6}
}

func (x *IsLikedRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *IsLikedRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type IsLikedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	IsLiked bool   `protobuf:"varint,3,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
}

func (x *IsLikedResponse) Reset() {
	*x = IsLikedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLikedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLikedResponse) ProtoMessage() {}

func (x *IsLikedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLikedResponse.ProtoReflect.Descriptor instead.
func (*IsLikedResponse) Descriptor() ([]byte, []int) {
	return file_idl_like_proto_rawDescGZIP(), []int{7}
}

func (x *IsLikedResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *IsLikedResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IsLikedResponse) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

var File_idl_like_proto protoreflect.FileDescriptor

var file_idl_like_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x3c, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x0e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x0f, 0x49, 0x73, 0x4c,
	0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x6b,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x32, 0x86, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x15,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x44, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x49, 0x73, 0x4c, 0x69,
	0x6b, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x49, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x54, 0x75, 0x2d, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x57, 0x65, 0x47, 0x6f, 0x2f, 0x69, 0x74, 0x75, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x6c,
	0x69, 0x6b, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_idl_like_proto_rawDescOnce sync.Once
	file_idl_like_proto_rawDescData = file_idl_like_proto_rawDesc
)

func file_idl_like_proto_rawDescGZIP() []byte {
	file_idl_like_proto_rawDescOnce.Do(func() {
		file_idl_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_like_proto_rawDescData)
	})
	return file_idl_like_proto_rawDescData
}

var file_idl_like_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_idl_like_proto_goTypes = []interface{}{
	(*AddLikesRequest)(nil),       // 0: like.AddLikesRequest
	(*AddLikesResponse)(nil),      // 1: like.AddLikesResponse
	(*DelLikesRequest)(nil),       // 2: like.DelLikesRequest
	(*DelLikesResponse)(nil),      // 3: like.DelLikesResponse
	(*GetLikesCountRequest)(nil),  // 4: like.GetLikesCountRequest
	(*GetLikesCountResponse)(nil), // 5: like.GetLikesCountResponse
	(*IsLikedRequest)(nil),        // 6: like.IsLikedRequest
	(*IsLikedResponse)(nil),       // 7: like.IsLikedResponse
}
var file_idl_like_proto_depIdxs = []int32{
	0, // 0: like.LikesService.AddLikes:input_type -> like.AddLikesRequest
	2, // 1: like.LikesService.DelLikes:input_type -> like.DelLikesRequest
	6, // 2: like.LikesService.IsLiked:input_type -> like.IsLikedRequest
	4, // 3: like.LikesService.GetLikesCount:input_type -> like.GetLikesCountRequest
	1, // 4: like.LikesService.AddLikes:output_type -> like.AddLikesResponse
	3, // 5: like.LikesService.DelLikes:output_type -> like.DelLikesResponse
	7, // 6: like.LikesService.IsLiked:output_type -> like.IsLikedResponse
	5, // 7: like.LikesService.GetLikesCount:output_type -> like.GetLikesCountResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_like_proto_init() }
func file_idl_like_proto_init() {
	if File_idl_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLikesRequest); i {
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
		file_idl_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLikesResponse); i {
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
		file_idl_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelLikesRequest); i {
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
		file_idl_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelLikesResponse); i {
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
		file_idl_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikesCountRequest); i {
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
		file_idl_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikesCountResponse); i {
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
		file_idl_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikedRequest); i {
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
		file_idl_like_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLikedResponse); i {
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
			RawDescriptor: file_idl_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_like_proto_goTypes,
		DependencyIndexes: file_idl_like_proto_depIdxs,
		MessageInfos:      file_idl_like_proto_msgTypes,
	}.Build()
	File_idl_like_proto = out.File
	file_idl_like_proto_rawDesc = nil
	file_idl_like_proto_goTypes = nil
	file_idl_like_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.3. DO NOT EDIT.

type LikesService interface {
	AddLikes(ctx context.Context, req *AddLikesRequest) (res *AddLikesResponse, err error)
	DelLikes(ctx context.Context, req *DelLikesRequest) (res *DelLikesResponse, err error)
	IsLiked(ctx context.Context, req *IsLikedRequest) (res *IsLikedResponse, err error)
	GetLikesCount(ctx context.Context, req *GetLikesCountRequest) (res *GetLikesCountResponse, err error)
}
