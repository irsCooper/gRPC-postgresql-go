// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notepad/notepad.proto

package ssov1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Article struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SetRequest struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags                 []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{1}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *SetRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SetRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SetRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SetResponse struct {
	ArticleId            int64    `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{2}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetArticleId() int64 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

type GetOneRequest struct {
	ArticleId            int64    `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneRequest) Reset()         { *m = GetOneRequest{} }
func (m *GetOneRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneRequest) ProtoMessage()    {}
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{3}
}

func (m *GetOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneRequest.Unmarshal(m, b)
}
func (m *GetOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneRequest.Marshal(b, m, deterministic)
}
func (m *GetOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneRequest.Merge(m, src)
}
func (m *GetOneRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneRequest.Size(m)
}
func (m *GetOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneRequest proto.InternalMessageInfo

func (m *GetOneRequest) GetArticleId() int64 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

type GetOneResponse struct {
	Article              *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneResponse) Reset()         { *m = GetOneResponse{} }
func (m *GetOneResponse) String() string { return proto.CompactTextString(m) }
func (*GetOneResponse) ProtoMessage()    {}
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{4}
}

func (m *GetOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneResponse.Unmarshal(m, b)
}
func (m *GetOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneResponse.Marshal(b, m, deterministic)
}
func (m *GetOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneResponse.Merge(m, src)
}
func (m *GetOneResponse) XXX_Size() int {
	return xxx_messageInfo_GetOneResponse.Size(m)
}
func (m *GetOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneResponse proto.InternalMessageInfo

func (m *GetOneResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type GetAllRequest struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{5}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type GetAllResponse struct {
	Article              []*Article `protobuf:"bytes,1,rep,name=article,proto3" json:"article,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{6}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetArticle() []*Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type IsOwnerRequest struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsOwnerRequest) Reset()         { *m = IsOwnerRequest{} }
func (m *IsOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*IsOwnerRequest) ProtoMessage()    {}
func (*IsOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{7}
}

func (m *IsOwnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsOwnerRequest.Unmarshal(m, b)
}
func (m *IsOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsOwnerRequest.Marshal(b, m, deterministic)
}
func (m *IsOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsOwnerRequest.Merge(m, src)
}
func (m *IsOwnerRequest) XXX_Size() int {
	return xxx_messageInfo_IsOwnerRequest.Size(m)
}
func (m *IsOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsOwnerRequest proto.InternalMessageInfo

func (m *IsOwnerRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type IsOwnerResponse struct {
	IsOwner              bool     `protobuf:"varint,1,opt,name=is_owner,json=isOwner,proto3" json:"is_owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsOwnerResponse) Reset()         { *m = IsOwnerResponse{} }
func (m *IsOwnerResponse) String() string { return proto.CompactTextString(m) }
func (*IsOwnerResponse) ProtoMessage()    {}
func (*IsOwnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{8}
}

func (m *IsOwnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsOwnerResponse.Unmarshal(m, b)
}
func (m *IsOwnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsOwnerResponse.Marshal(b, m, deterministic)
}
func (m *IsOwnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsOwnerResponse.Merge(m, src)
}
func (m *IsOwnerResponse) XXX_Size() int {
	return xxx_messageInfo_IsOwnerResponse.Size(m)
}
func (m *IsOwnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsOwnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsOwnerResponse proto.InternalMessageInfo

func (m *IsOwnerResponse) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

type DeleteRequest struct {
	ArticleId            int64    `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetArticleId() int64 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

type DeleteResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9de4b64e70d8b5e5, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Article)(nil), "notepad.Article")
	proto.RegisterType((*SetRequest)(nil), "notepad.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "notepad.SetResponse")
	proto.RegisterType((*GetOneRequest)(nil), "notepad.GetOneRequest")
	proto.RegisterType((*GetOneResponse)(nil), "notepad.GetOneResponse")
	proto.RegisterType((*GetAllRequest)(nil), "notepad.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "notepad.GetAllResponse")
	proto.RegisterType((*IsOwnerRequest)(nil), "notepad.IsOwnerRequest")
	proto.RegisterType((*IsOwnerResponse)(nil), "notepad.IsOwnerResponse")
	proto.RegisterType((*DeleteRequest)(nil), "notepad.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "notepad.DeleteResponse")
}

func init() {
	proto.RegisterFile("notepad/notepad.proto", fileDescriptor_9de4b64e70d8b5e5)
}

var fileDescriptor_9de4b64e70d8b5e5 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0x03, 0xe5, 0xb1, 0x30, 0xe4, 0xf1, 0x5e, 0x56, 0xc4, 0x42, 0x62, 0x42, 0x7a, 0x22,
	0x48, 0x50, 0xf0, 0x64, 0xe4, 0x82, 0x31, 0x31, 0x5c, 0x24, 0x29, 0x37, 0x2f, 0xa4, 0xd2, 0x8d,
	0x36, 0x36, 0x2d, 0x76, 0x07, 0xfc, 0x4e, 0x7e, 0x4a, 0xc3, 0xee, 0xb4, 0x65, 0x83, 0x11, 0x2e,
	0x9e, 0xda, 0x99, 0xf9, 0xef, 0xfc, 0x76, 0xfe, 0xd3, 0xc2, 0x69, 0x14, 0xa3, 0x58, 0x79, 0xfe,
	0x25, 0x3d, 0x07, 0xab, 0x24, 0xc6, 0x98, 0x33, 0x0a, 0x9d, 0x57, 0x60, 0x93, 0x04, 0x83, 0x65,
	0x28, 0x78, 0x0b, 0x2a, 0xf1, 0x47, 0x24, 0x92, 0x45, 0xe0, 0xdb, 0x85, 0x4e, 0xa1, 0x5b, 0x75,
	0x99, 0x8a, 0xa7, 0x3e, 0x6f, 0xc0, 0x1f, 0x0c, 0x30, 0x14, 0x76, 0x51, 0xe5, 0x75, 0xc0, 0x6d,
	0x60, 0xcb, 0x38, 0x42, 0x11, 0xa1, 0x6d, 0x69, 0x3d, 0x85, 0x9c, 0x43, 0x09, 0xbd, 0x17, 0x69,
	0x97, 0x3a, 0x56, 0xb7, 0xea, 0xaa, 0x77, 0xe7, 0x0d, 0x60, 0x2e, 0xd0, 0x15, 0xef, 0x6b, 0x21,
	0xf1, 0xb7, 0x61, 0x7d, 0xa8, 0x29, 0x98, 0x5c, 0xc5, 0x91, 0x14, 0xfc, 0x1c, 0xc0, 0xd3, 0x53,
	0xa6, 0x3c, 0xcb, 0xad, 0x52, 0x66, 0xea, 0x3b, 0x03, 0xf8, 0xfb, 0x20, 0x70, 0x16, 0x89, 0xf4,
	0x76, 0x07, 0xf4, 0x63, 0xa8, 0xa7, 0x7a, 0x02, 0xf4, 0x80, 0x51, 0x59, 0xa9, 0x6b, 0xa3, 0xff,
	0x83, 0xd4, 0x70, 0xb2, 0xd7, 0x4d, 0x05, 0x4e, 0x4f, 0xd1, 0x26, 0x61, 0x78, 0xd8, 0x0b, 0x22,
	0x29, 0xed, 0x77, 0x24, 0xeb, 0x67, 0xd2, 0x05, 0xd4, 0xa7, 0x72, 0xb6, 0x6d, 0x75, 0x04, 0xaa,
	0x0f, 0xff, 0x32, 0x31, 0xb1, 0x5a, 0x50, 0x09, 0xe4, 0x42, 0x09, 0x94, 0xba, 0xe2, 0xb2, 0x40,
	0x4b, 0xb6, 0x96, 0xdd, 0x8b, 0x50, 0xe0, 0xb1, 0x96, 0x75, 0xa1, 0x9e, 0xea, 0xa9, 0x79, 0x13,
	0xca, 0x12, 0x3d, 0x5c, 0x4b, 0xba, 0x08, 0x45, 0xa3, 0xcf, 0x22, 0xb0, 0x47, 0x3d, 0x11, 0xbf,
	0x02, 0x6b, 0x2e, 0x90, 0x9f, 0x64, 0x23, 0xe6, 0x5f, 0x50, 0xbb, 0x61, 0x26, 0xa9, 0xeb, 0x0d,
	0x94, 0xf5, 0x6a, 0x78, 0x33, 0xab, 0x1b, 0xbb, 0x6d, 0x9f, 0xed, 0xe5, 0x8d, 0xa3, 0x93, 0x30,
	0x34, 0x8f, 0xe6, 0x8b, 0x32, 0x8f, 0xee, 0x2e, 0x65, 0x0c, 0x8c, 0xbc, 0xe3, 0xb9, 0xc6, 0xb4,
	0xbe, 0x6d, 0xef, 0x17, 0x72, 0xb0, 0xf6, 0x66, 0x07, 0x6c, 0x98, 0xbb, 0x03, 0x36, 0x4d, 0xbc,
	0x6b, 0x3e, 0x35, 0x82, 0x44, 0x66, 0xd5, 0xcd, 0xf0, 0x56, 0xca, 0x78, 0x33, 0x7c, 0x2e, 0xab,
	0xdf, 0xfc, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa6, 0xfb, 0x39, 0xff, 0x03, 0x00, 0x00,
}
