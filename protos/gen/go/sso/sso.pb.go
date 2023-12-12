// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sso/sso.proto

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

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	AppId                int32    `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetAppId() int32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type IsAdminRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAdminRequest) Reset()         { *m = IsAdminRequest{} }
func (m *IsAdminRequest) String() string { return proto.CompactTextString(m) }
func (*IsAdminRequest) ProtoMessage()    {}
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{4}
}

func (m *IsAdminRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAdminRequest.Unmarshal(m, b)
}
func (m *IsAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAdminRequest.Marshal(b, m, deterministic)
}
func (m *IsAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAdminRequest.Merge(m, src)
}
func (m *IsAdminRequest) XXX_Size() int {
	return xxx_messageInfo_IsAdminRequest.Size(m)
}
func (m *IsAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAdminRequest proto.InternalMessageInfo

func (m *IsAdminRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type IsAdminResponse struct {
	IsAdmin              bool     `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAdminResponse) Reset()         { *m = IsAdminResponse{} }
func (m *IsAdminResponse) String() string { return proto.CompactTextString(m) }
func (*IsAdminResponse) ProtoMessage()    {}
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_402072c85cc63057, []int{5}
}

func (m *IsAdminResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAdminResponse.Unmarshal(m, b)
}
func (m *IsAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAdminResponse.Marshal(b, m, deterministic)
}
func (m *IsAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAdminResponse.Merge(m, src)
}
func (m *IsAdminResponse) XXX_Size() int {
	return xxx_messageInfo_IsAdminResponse.Size(m)
}
func (m *IsAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAdminResponse proto.InternalMessageInfo

func (m *IsAdminResponse) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "auth.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "auth.RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*IsAdminRequest)(nil), "auth.IsAdminRequest")
	proto.RegisterType((*IsAdminResponse)(nil), "auth.IsAdminResponse")
}

func init() {
	proto.RegisterFile("sso/sso.proto", fileDescriptor_402072c85cc63057)
}

var fileDescriptor_402072c85cc63057 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x24, 0xb6, 0xf9, 0xf0, 0x61, 0x6d, 0x59, 0x13, 0xad, 0x39, 0x95, 0x80, 0x50, 0x51, 0xa2,
	0x55, 0x10, 0xc4, 0x53, 0xf5, 0x14, 0xf0, 0xb4, 0x17, 0xc1, 0x4b, 0x59, 0xc9, 0xd2, 0x2e, 0xda,
	0xec, 0x9a, 0xb7, 0xa9, 0x3f, 0xca, 0x3f, 0x29, 0xd9, 0x6c, 0x5b, 0x1b, 0xf0, 0xe2, 0x71, 0x66,
	0xe7, 0xcd, 0xce, 0x1b, 0x1e, 0xf4, 0x10, 0xe5, 0x15, 0xa2, 0x4c, 0x55, 0x29, 0xb5, 0x24, 0x5d,
	0x56, 0xe9, 0x45, 0xf2, 0x04, 0x7d, 0xca, 0xe7, 0x02, 0x35, 0x2f, 0x29, 0xff, 0xac, 0x38, 0x6a,
	0x12, 0x82, 0xcb, 0x97, 0x4c, 0x7c, 0x0c, 0x9d, 0x91, 0x33, 0xde, 0xa7, 0x0d, 0x20, 0x31, 0x04,
	0x8a, 0x21, 0x7e, 0xc9, 0x32, 0x1f, 0xee, 0x99, 0x87, 0x0d, 0x4e, 0x2e, 0x60, 0xb0, 0x35, 0x41,
	0x25, 0x0b, 0xe4, 0xe4, 0x04, 0xfc, 0x0a, 0x79, 0x39, 0x13, 0xb9, 0xf1, 0xe9, 0x50, 0xaf, 0x86,
	0x59, 0x9e, 0xbc, 0xc0, 0xc1, 0xb3, 0x9c, 0x8b, 0xe2, 0xdf, 0xdf, 0x91, 0x08, 0x3c, 0xa6, 0x54,
	0xed, 0xdc, 0x19, 0x39, 0x63, 0x97, 0xba, 0x4c, 0xa9, 0x2c, 0x4f, 0xce, 0xa0, 0x67, 0x8d, 0x6d,
	0x84, 0x10, 0x5c, 0x2d, 0xdf, 0x79, 0xb1, 0x76, 0x36, 0x20, 0x39, 0x87, 0xc3, 0x0c, 0xa7, 0xf9,
	0x72, 0x9b, 0xe0, 0xcf, 0xa8, 0x97, 0xd0, 0xdf, 0x48, 0xad, 0xe7, 0x29, 0x04, 0x02, 0x67, 0xac,
	0xe6, 0x8c, 0x38, 0xa0, 0xbe, 0x68, 0x24, 0x37, 0xdf, 0x0e, 0x74, 0xa7, 0x95, 0x5e, 0x90, 0x7b,
	0x08, 0xd6, 0x75, 0x90, 0x28, 0xad, 0x6b, 0x4e, 0x5b, 0x1d, 0xc7, 0xc7, 0x6d, 0xda, 0xda, 0x5f,
	0x83, 0x6b, 0x76, 0x20, 0xa4, 0x11, 0xfc, 0x6e, 0x2a, 0x3e, 0xda, 0xe1, 0xec, 0xc4, 0x1d, 0xf8,
	0x36, 0x23, 0x09, 0x9b, 0xf7, 0xdd, 0xed, 0xe2, 0xa8, 0xc5, 0x36, 0x73, 0x8f, 0xe4, 0x75, 0x20,
	0x4a, 0x4c, 0xeb, 0x7b, 0x58, 0x4d, 0x1e, 0x10, 0xe5, 0x6a, 0xf2, 0xe6, 0x99, 0xcb, 0xb8, 0xfd,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x50, 0x69, 0xbb, 0x7d, 0x2a, 0x02, 0x00, 0x00,
}
