// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/proto/auth.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth/proto/auth.proto

It has these top-level messages:
	LoginReq
	LoginRes
	CheckAuthReq
	CheckAuthRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginReq struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto1.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRes struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *LoginRes) Reset()                    { *m = LoginRes{} }
func (m *LoginRes) String() string            { return proto1.CompactTextString(m) }
func (*LoginRes) ProtoMessage()               {}
func (*LoginRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginRes) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *LoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckAuthReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *CheckAuthReq) Reset()                    { *m = CheckAuthReq{} }
func (m *CheckAuthReq) String() string            { return proto1.CompactTextString(m) }
func (*CheckAuthReq) ProtoMessage()               {}
func (*CheckAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckAuthReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckAuthRes struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	ErrMsg  string `protobuf:"bytes,3,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CheckAuthRes) Reset()                    { *m = CheckAuthRes{} }
func (m *CheckAuthRes) String() string            { return proto1.CompactTextString(m) }
func (*CheckAuthRes) ProtoMessage()               {}
func (*CheckAuthRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckAuthRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CheckAuthRes) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CheckAuthRes) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto1.RegisterType((*LoginReq)(nil), "proto.LoginReq")
	proto1.RegisterType((*LoginRes)(nil), "proto.LoginRes")
	proto1.RegisterType((*CheckAuthReq)(nil), "proto.CheckAuthReq")
	proto1.RegisterType((*CheckAuthRes)(nil), "proto.CheckAuthRes")
}

func init() { proto1.RegisterFile("auth/proto/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x4e, 0x80, 0x30,
	0x10, 0x85, 0x83, 0x58, 0xc4, 0x89, 0x89, 0x66, 0xfc, 0x49, 0xc3, 0xca, 0x10, 0x17, 0xba, 0x81,
	0xa8, 0x27, 0x50, 0xb7, 0xba, 0xe9, 0x0d, 0x10, 0x1a, 0x20, 0x44, 0x8a, 0x9d, 0x36, 0x5e, 0xdf,
	0xb4, 0x50, 0x52, 0x8d, 0x71, 0xd5, 0x79, 0x6f, 0xa6, 0xaf, 0x5f, 0x07, 0x2e, 0x1b, 0x6b, 0x86,
	0x7a, 0xd1, 0xca, 0xa8, 0xda, 0x95, 0x95, 0x2f, 0x91, 0xf9, 0xa3, 0x7c, 0x86, 0xfc, 0x55, 0xf5,
	0xe3, 0x2c, 0xe4, 0x27, 0x16, 0x90, 0x5b, 0x92, 0x7a, 0x6e, 0x3e, 0x24, 0x4f, 0xae, 0x93, 0xdb,
	0x63, 0xb1, 0x6b, 0xd7, 0x5b, 0x1a, 0xa2, 0x2f, 0xa5, 0x3b, 0x7e, 0xb0, 0xf6, 0x82, 0x2e, 0xc5,
	0x9e, 0x41, 0xc8, 0xe1, 0x88, 0x6c, 0xdb, 0x4a, 0x22, 0x1f, 0x91, 0x8b, 0x20, 0xf1, 0x0a, 0x32,
	0xa9, 0xf5, 0x1b, 0xf5, 0xdb, 0xfd, 0x4d, 0xe1, 0x05, 0x30, 0xa3, 0x26, 0x39, 0xf3, 0xd4, 0xdb,
	0xab, 0x28, 0x6f, 0xe0, 0xe4, 0x65, 0x90, 0xed, 0xf4, 0x64, 0xcd, 0xe0, 0xd8, 0xf6, 0xa9, 0x24,
	0x9e, 0x12, 0x3f, 0xa6, 0xfe, 0x7b, 0xfd, 0x0c, 0x52, 0x3b, 0x06, 0x74, 0x57, 0x46, 0x3c, 0x69,
	0xcc, 0xf3, 0xd0, 0xc1, 0xa1, 0x8b, 0xc3, 0x3b, 0x60, 0xfe, 0x57, 0x78, 0xba, 0x6e, 0xac, 0x0a,
	0x7b, 0x2a, 0x7e, 0x19, 0x84, 0xf7, 0xc0, 0x3c, 0x06, 0x9e, 0x6f, 0x9d, 0x18, 0xbd, 0xf8, 0xc3,
	0xa4, 0xf7, 0xcc, 0x7b, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x6a, 0x59, 0xcc, 0x9e,
	0x01, 0x00, 0x00,
}
