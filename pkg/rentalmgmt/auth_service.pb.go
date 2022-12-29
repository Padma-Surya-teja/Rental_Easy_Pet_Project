// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/rentalmgmt/auth_service.proto

package rentalmgmt

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49dea76826d17116, []int{0}
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

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49dea76826d17116, []int{1}
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

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "rentalmgmt.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "rentalmgmt.LoginResponse")
}

func init() { proto.RegisterFile("pkg/rentalmgmt/auth_service.proto", fileDescriptor_49dea76826d17116) }

var fileDescriptor_49dea76826d17116 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4f, 0x85, 0x30,
	0x14, 0x85, 0xc5, 0x44, 0xa3, 0x17, 0x5c, 0x3a, 0x21, 0xc6, 0x44, 0x98, 0x9c, 0x4a, 0x82, 0xbb,
	0x89, 0x0e, 0x4e, 0xba, 0xa0, 0x93, 0x0b, 0xa9, 0x78, 0x03, 0x04, 0xfb, 0xc3, 0xde, 0x56, 0xe3,
	0x7f, 0x6f, 0xa0, 0xbc, 0x07, 0xc3, 0x1b, 0xcf, 0xf9, 0x92, 0xd3, 0xaf, 0x17, 0x72, 0x33, 0x76,
	0xa5, 0x45, 0xe5, 0xc4, 0x97, 0xec, 0xa4, 0x2b, 0x85, 0x77, 0x7d, 0x43, 0x68, 0x7f, 0x86, 0x16,
	0xb9, 0xb1, 0xda, 0x69, 0x06, 0x2b, 0x2e, 0x9e, 0x20, 0x79, 0xd6, 0xdd, 0xa0, 0x6a, 0xfc, 0xf6,
	0x48, 0x8e, 0x65, 0x70, 0xe6, 0x09, 0xad, 0x12, 0x12, 0xd3, 0xe8, 0x26, 0xba, 0x3d, 0xaf, 0xf7,
	0x79, 0x62, 0x46, 0x10, 0xfd, 0x6a, 0xfb, 0x99, 0x1e, 0x07, 0xb6, 0xcb, 0x45, 0x05, 0x17, 0xcb,
	0x0e, 0x19, 0xad, 0x08, 0x59, 0x0e, 0x89, 0x68, 0x5b, 0x24, 0x6a, 0x9c, 0x1e, 0x51, 0x2d, 0x63,
	0x71, 0xe8, 0xde, 0xa6, 0xaa, 0x7a, 0x81, 0xf8, 0xc1, 0xbb, 0xfe, 0x35, 0xc8, 0xb1, 0x7b, 0x38,
	0x99, 0x27, 0x58, 0xca, 0x57, 0x41, 0xbe, 0xb5, 0xcb, 0x2e, 0x0f, 0x90, 0xf0, 0x5e, 0x71, 0xf4,
	0x78, 0xfd, 0x7e, 0x15, 0x68, 0x83, 0x82, 0xfe, 0xf8, 0xa0, 0x4a, 0xb9, 0x39, 0xc4, 0xc7, 0xe9,
	0xfc, 0xf9, 0xbb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xd1, 0xbf, 0xf7, 0x21, 0x01, 0x00,
	0x00,
}
