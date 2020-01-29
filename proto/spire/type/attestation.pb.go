// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attestation.proto

package spire_type

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

type AttestationData struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a47e68d91e2b64e, []int{0}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestationData) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*AttestationData)(nil), "spire.type.AttestationData")
}

func init() { proto.RegisterFile("attestation.proto", fileDescriptor_3a47e68d91e2b64e) }

var fileDescriptor_3a47e68d91e2b64e = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0x29, 0x49,
	0x2d, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a,
	0x2e, 0xc8, 0x2c, 0x4a, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x55, 0xb2, 0xe7, 0xe2, 0x77, 0x44, 0x28,
	0x70, 0x49, 0x2c, 0x49, 0x14, 0x12, 0xe2, 0x62, 0x01, 0x49, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x12, 0x4c,
	0x60, 0x61, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x2b,
	0x1d, 0x44, 0x68, 0x00, 0x00, 0x00,
}
