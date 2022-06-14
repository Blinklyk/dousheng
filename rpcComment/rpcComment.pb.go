// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rpcComment/rpcComment.proto

package rpcComment

import (
	rpcUser "github.com/RaymondCode/simple-demo/pb/rpcUser"
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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User       *rpcUser.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Content    string        `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate string        `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcComment_rpcComment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_rpcComment_rpcComment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_rpcComment_rpcComment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUser() *rpcUser.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type CommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VideoId     string `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	ActionType  string `protobuf:"bytes,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	CommentText string `protobuf:"bytes,4,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"`
	CommentId   string `protobuf:"bytes,5,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId      int64  `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentRequest) Reset() {
	*x = CommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcComment_rpcComment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentRequest) ProtoMessage() {}

func (x *CommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpcComment_rpcComment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentRequest.ProtoReflect.Descriptor instead.
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return file_rpcComment_rpcComment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *CommentRequest) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *CommentRequest) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *CommentRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *CommentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	Comment    *Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentResponse) Reset() {
	*x = CommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcComment_rpcComment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponse) ProtoMessage() {}

func (x *CommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpcComment_rpcComment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponse.ProtoReflect.Descriptor instead.
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return file_rpcComment_rpcComment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommentResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VideoId string `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	UserId  int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CommentListReq) Reset() {
	*x = CommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcComment_rpcComment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListReq) ProtoMessage() {}

func (x *CommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpcComment_rpcComment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListReq.ProtoReflect.Descriptor instead.
func (*CommentListReq) Descriptor() ([]byte, []int) {
	return file_rpcComment_rpcComment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentListReq) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *CommentListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CommentListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg   string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	CommentList []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"`
}

func (x *CommentListResp) Reset() {
	*x = CommentListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcComment_rpcComment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResp) ProtoMessage() {}

func (x *CommentListResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpcComment_rpcComment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResp.ProtoReflect.Descriptor instead.
func (*CommentListResp) Descriptor() ([]byte, []int) {
	return file_rpcComment_rpcComment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommentListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentListResp) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

var File_rpcComment_rpcComment_proto protoreflect.FileDescriptor

var file_rpcComment_rpcComment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72,
	0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x15, 0x72, 0x70, 0x63, 0x55, 0x73,
	0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x77, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xa6, 0x01, 0x0a, 0x11, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x79, 0x6d,
	0x6f, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x64,
	0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x3b, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcComment_rpcComment_proto_rawDescOnce sync.Once
	file_rpcComment_rpcComment_proto_rawDescData = file_rpcComment_rpcComment_proto_rawDesc
)

func file_rpcComment_rpcComment_proto_rawDescGZIP() []byte {
	file_rpcComment_rpcComment_proto_rawDescOnce.Do(func() {
		file_rpcComment_rpcComment_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcComment_rpcComment_proto_rawDescData)
	})
	return file_rpcComment_rpcComment_proto_rawDescData
}

var file_rpcComment_rpcComment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpcComment_rpcComment_proto_goTypes = []interface{}{
	(*Comment)(nil),         // 0: rpcComment.Comment
	(*CommentRequest)(nil),  // 1: rpcComment.CommentRequest
	(*CommentResponse)(nil), // 2: rpcComment.CommentResponse
	(*CommentListReq)(nil),  // 3: rpcComment.CommentListReq
	(*CommentListResp)(nil), // 4: rpcComment.CommentListResp
	(*rpcUser.User)(nil),    // 5: rpcUser.User
}
var file_rpcComment_rpcComment_proto_depIdxs = []int32{
	5, // 0: rpcComment.Comment.user:type_name -> rpcUser.User
	0, // 1: rpcComment.CommentResponse.comment:type_name -> rpcComment.Comment
	0, // 2: rpcComment.CommentListResp.comment_list:type_name -> rpcComment.Comment
	1, // 3: rpcComment.RPCCommentService.Comment:input_type -> rpcComment.CommentRequest
	3, // 4: rpcComment.RPCCommentService.GetCommentList:input_type -> rpcComment.CommentListReq
	2, // 5: rpcComment.RPCCommentService.Comment:output_type -> rpcComment.CommentResponse
	4, // 6: rpcComment.RPCCommentService.GetCommentList:output_type -> rpcComment.CommentListResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpcComment_rpcComment_proto_init() }
func file_rpcComment_rpcComment_proto_init() {
	if File_rpcComment_rpcComment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcComment_rpcComment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_rpcComment_rpcComment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentRequest); i {
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
		file_rpcComment_rpcComment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponse); i {
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
		file_rpcComment_rpcComment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListReq); i {
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
		file_rpcComment_rpcComment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResp); i {
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
			RawDescriptor: file_rpcComment_rpcComment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpcComment_rpcComment_proto_goTypes,
		DependencyIndexes: file_rpcComment_rpcComment_proto_depIdxs,
		MessageInfos:      file_rpcComment_rpcComment_proto_msgTypes,
	}.Build()
	File_rpcComment_rpcComment_proto = out.File
	file_rpcComment_rpcComment_proto_rawDesc = nil
	file_rpcComment_rpcComment_proto_goTypes = nil
	file_rpcComment_rpcComment_proto_depIdxs = nil
}
