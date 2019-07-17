// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync/conflict.proto

package sync

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

type Conflict struct {
	AlphaChanges         []*Change `protobuf:"bytes,1,rep,name=alphaChanges,proto3" json:"alphaChanges,omitempty"`
	BetaChanges          []*Change `protobuf:"bytes,2,rep,name=betaChanges,proto3" json:"betaChanges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Conflict) Reset()         { *m = Conflict{} }
func (m *Conflict) String() string { return proto.CompactTextString(m) }
func (*Conflict) ProtoMessage()    {}
func (*Conflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_280f6db0ca354ed7, []int{0}
}

func (m *Conflict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conflict.Unmarshal(m, b)
}
func (m *Conflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conflict.Marshal(b, m, deterministic)
}
func (m *Conflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conflict.Merge(m, src)
}
func (m *Conflict) XXX_Size() int {
	return xxx_messageInfo_Conflict.Size(m)
}
func (m *Conflict) XXX_DiscardUnknown() {
	xxx_messageInfo_Conflict.DiscardUnknown(m)
}

var xxx_messageInfo_Conflict proto.InternalMessageInfo

func (m *Conflict) GetAlphaChanges() []*Change {
	if m != nil {
		return m.AlphaChanges
	}
	return nil
}

func (m *Conflict) GetBetaChanges() []*Change {
	if m != nil {
		return m.BetaChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*Conflict)(nil), "sync.Conflict")
}

func init() { proto.RegisterFile("sync/conflict.proto", fileDescriptor_280f6db0ca354ed7) }

var fileDescriptor_280f6db0ca354ed7 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xae, 0xcc, 0x4b,
	0xd6, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0xc9, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x01, 0x09, 0x4a, 0x09, 0x42, 0xa4, 0x32, 0x12, 0xf3, 0xd2, 0x53, 0x21, 0x12, 0x4a, 0x39,
	0x5c, 0x1c, 0xce, 0x50, 0xa5, 0x42, 0x06, 0x5c, 0x3c, 0x89, 0x39, 0x05, 0x19, 0x89, 0xce, 0x60,
	0x05, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x20, 0x5d, 0x7a, 0x10, 0xc1,
	0x20, 0x14, 0x15, 0x42, 0x7a, 0x5c, 0xdc, 0x49, 0xa9, 0x25, 0x70, 0x0d, 0x4c, 0x58, 0x34, 0x20,
	0x2b, 0x70, 0x52, 0x8b, 0x52, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0xcf, 0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7, 0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd,
	0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x07, 0xe9, 0x4d, 0x62, 0x03, 0x3b, 0xce, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0x81, 0xc9, 0x55, 0xcc, 0x00, 0x00, 0x00,
}