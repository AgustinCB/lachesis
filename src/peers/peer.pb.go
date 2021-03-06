// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer.proto

package peers

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Peer struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NetAddr              string   `protobuf:"bytes,2,opt,name=NetAddr,proto3" json:"NetAddr,omitempty"`
	PubKeyHex            string   `protobuf:"bytes,3,opt,name=PubKeyHex,proto3" json:"PubKeyHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_055ae5a865fc1c9e, []int{0}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Peer) GetNetAddr() string {
	if m != nil {
		return m.NetAddr
	}
	return ""
}

func (m *Peer) GetPubKeyHex() string {
	if m != nil {
		return m.PubKeyHex
	}
	return ""
}

func init() {
	proto.RegisterType((*Peer)(nil), "poset.Peer")
}

func init() { proto.RegisterFile("peer.proto", fileDescriptor_055ae5a865fc1c9e) }

var fileDescriptor_055ae5a865fc1c9e = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x4d, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xc8, 0x2f, 0x4e, 0x2d, 0x51, 0xf2, 0xe3,
	0x62, 0x09, 0x48, 0x4d, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0e, 0x62, 0xf2, 0x74, 0x11, 0x92, 0xe0, 0x62, 0xf7, 0x4b, 0x2d, 0x71, 0x4c, 0x49, 0x29,
	0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x64, 0xb8, 0x38, 0x03, 0x4a, 0x93,
	0xbc, 0x53, 0x2b, 0x3d, 0x52, 0x2b, 0x24, 0x98, 0xc1, 0x72, 0x08, 0x81, 0x24, 0x36, 0xb0, 0xe9,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xad, 0x7c, 0x6d, 0x6b, 0x00, 0x00, 0x00,
}
