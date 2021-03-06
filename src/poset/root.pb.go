// Code generated by protoc-gen-go. DO NOT EDIT.
// source: root.proto

package poset

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

type RootEvent struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	CreatorID            int64    `protobuf:"varint,2,opt,name=CreatorID,proto3" json:"CreatorID,omitempty"`
	Index                int64    `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
	LamportTimestamp     int64    `protobuf:"varint,4,opt,name=LamportTimestamp,proto3" json:"LamportTimestamp,omitempty"`
	Round                int64    `protobuf:"varint,5,opt,name=Round,proto3" json:"Round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RootEvent) Reset()         { *m = RootEvent{} }
func (m *RootEvent) String() string { return proto.CompactTextString(m) }
func (*RootEvent) ProtoMessage()    {}
func (*RootEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{0}
}

func (m *RootEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootEvent.Unmarshal(m, b)
}
func (m *RootEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootEvent.Marshal(b, m, deterministic)
}
func (m *RootEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootEvent.Merge(m, src)
}
func (m *RootEvent) XXX_Size() int {
	return xxx_messageInfo_RootEvent.Size(m)
}
func (m *RootEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RootEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RootEvent proto.InternalMessageInfo

func (m *RootEvent) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *RootEvent) GetCreatorID() int64 {
	if m != nil {
		return m.CreatorID
	}
	return 0
}

func (m *RootEvent) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RootEvent) GetLamportTimestamp() int64 {
	if m != nil {
		return m.LamportTimestamp
	}
	return 0
}

func (m *RootEvent) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type Root struct {
	NextRound            int64                 `protobuf:"varint,1,opt,name=NextRound,proto3" json:"NextRound,omitempty"`
	SelfParent           *RootEvent            `protobuf:"bytes,2,opt,name=SelfParent,proto3" json:"SelfParent,omitempty"`
	Others               map[string]*RootEvent `protobuf:"bytes,3,rep,name=Others,proto3" json:"Others,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Root) Reset()         { *m = Root{} }
func (m *Root) String() string { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()    {}
func (*Root) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{1}
}

func (m *Root) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Root.Unmarshal(m, b)
}
func (m *Root) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Root.Marshal(b, m, deterministic)
}
func (m *Root) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Root.Merge(m, src)
}
func (m *Root) XXX_Size() int {
	return xxx_messageInfo_Root.Size(m)
}
func (m *Root) XXX_DiscardUnknown() {
	xxx_messageInfo_Root.DiscardUnknown(m)
}

var xxx_messageInfo_Root proto.InternalMessageInfo

func (m *Root) GetNextRound() int64 {
	if m != nil {
		return m.NextRound
	}
	return 0
}

func (m *Root) GetSelfParent() *RootEvent {
	if m != nil {
		return m.SelfParent
	}
	return nil
}

func (m *Root) GetOthers() map[string]*RootEvent {
	if m != nil {
		return m.Others
	}
	return nil
}

func init() {
	proto.RegisterType((*RootEvent)(nil), "poset.RootEvent")
	proto.RegisterType((*Root)(nil), "poset.Root")
	proto.RegisterMapType((map[string]*RootEvent)(nil), "poset.Root.OthersEntry")
}

func init() { proto.RegisterFile("root.proto", fileDescriptor_08a043f6ee9336a8) }

var fileDescriptor_08a043f6ee9336a8 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0x26, 0x29, 0x64, 0x72, 0x09, 0xc3, 0x1f, 0xfe, 0x8b, 0x78, 0x08, 0x3d, 0x48,
	0xf0, 0x10, 0xa5, 0x5e, 0xc4, 0xab, 0x16, 0x2c, 0x8a, 0xca, 0xea, 0x0b, 0xac, 0x74, 0xa4, 0x62,
	0x93, 0x09, 0x9b, 0x69, 0x69, 0x1f, 0xc4, 0x37, 0xf3, 0x81, 0x24, 0xbb, 0xa5, 0x06, 0xc4, 0xdb,
	0xcc, 0xf7, 0xfd, 0x98, 0x6f, 0x66, 0x00, 0x1c, 0xb3, 0x54, 0xad, 0x63, 0x61, 0x4c, 0x5a, 0xee,
	0x48, 0x26, 0x9f, 0x0a, 0x52, 0xc3, 0x2c, 0xb3, 0x0d, 0x35, 0x82, 0x08, 0xf1, 0xad, 0xed, 0x96,
	0x5a, 0x15, 0xaa, 0x4c, 0x8d, 0xaf, 0xf1, 0x18, 0xd2, 0x6b, 0x47, 0x56, 0xd8, 0xcd, 0x6f, 0xf4,
	0xa8, 0x50, 0x65, 0x64, 0x7e, 0x04, 0xfc, 0x07, 0xc9, 0xbc, 0x59, 0xd0, 0x56, 0x47, 0xde, 0x09,
	0x0d, 0x9e, 0x42, 0x7e, 0x6f, 0xeb, 0x96, 0x9d, 0xbc, 0xbc, 0xd7, 0xd4, 0x89, 0xad, 0x5b, 0x1d,
	0x7b, 0xe0, 0x97, 0xde, 0x4f, 0x30, 0xbc, 0x6e, 0x16, 0x3a, 0x09, 0x13, 0x7c, 0x33, 0xf9, 0x52,
	0x10, 0xf7, 0x7b, 0xf5, 0xf1, 0x0f, 0xb4, 0x95, 0x80, 0xa8, 0x10, 0x7f, 0x10, 0xf0, 0x1c, 0xe0,
	0x99, 0x56, 0x6f, 0x4f, 0xd6, 0x51, 0x23, 0x7e, 0xbb, 0x6c, 0x9a, 0x57, 0xfe, 0xb4, 0xea, 0x70,
	0x96, 0x19, 0x30, 0x78, 0x06, 0xe3, 0x47, 0x59, 0x92, 0xeb, 0x74, 0x54, 0x44, 0x65, 0x36, 0xfd,
	0x3f, 0xa0, 0xab, 0xe0, 0xcc, 0x1a, 0x71, 0x3b, 0xb3, 0xc7, 0x8e, 0xee, 0x20, 0x1b, 0xc8, 0x98,
	0x43, 0xf4, 0x41, 0xbb, 0xfd, 0x87, 0xfa, 0x12, 0x4f, 0x20, 0xd9, 0xd8, 0xd5, 0x9a, 0xfe, 0x8c,
	0x0f, 0xf6, 0xd5, 0xe8, 0x52, 0xbd, 0x8e, 0xfd, 0xf3, 0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0xec, 0x23, 0x14, 0x8a, 0x01, 0x00, 0x00,
}
