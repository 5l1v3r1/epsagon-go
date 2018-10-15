// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trace.proto

package protocol

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

type Trace struct {
	AppName              string       `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Token                string       `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Events               []*Event     `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	Exceptions           []*Exception `protobuf:"bytes,4,rep,name=exceptions,proto3" json:"exceptions,omitempty"`
	Version              string       `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Platform             string       `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_0571941a1d628a80, []int{0}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Trace) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Trace) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Trace) GetExceptions() []*Exception {
	if m != nil {
		return m.Exceptions
	}
	return nil
}

func (m *Trace) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Trace) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func init() {
	proto.RegisterType((*Trace)(nil), "protocol.Trace")
}

func init() { proto.RegisterFile("trace.proto", fileDescriptor_0571941a1d628a80) }

var fileDescriptor_0571941a1d628a80 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x3f, 0xcf, 0x82, 0x30,
	0x10, 0xc6, 0xc3, 0xcb, 0xcb, 0x1f, 0x8f, 0x81, 0xa4, 0x3a, 0x54, 0x26, 0xe2, 0x22, 0x13, 0x83,
	0x7c, 0x06, 0x57, 0x07, 0xe2, 0x6e, 0x2a, 0x39, 0x13, 0x22, 0xb4, 0x97, 0xd2, 0x10, 0x3f, 0xa4,
	0x1f, 0xca, 0xd0, 0x02, 0x61, 0x6a, 0x7e, 0xf7, 0x7b, 0x7a, 0xf7, 0x40, 0x62, 0xb4, 0x68, 0xb0,
	0x24, 0xad, 0x8c, 0x62, 0xb1, 0x7d, 0x1a, 0xd5, 0x65, 0x29, 0x7e, 0x1a, 0x24, 0xd3, 0x2a, 0xe9,
	0x54, 0x96, 0xe0, 0x88, 0xd2, 0x38, 0x38, 0x7d, 0x3d, 0x08, 0xee, 0xd3, 0x3f, 0x76, 0x84, 0x58,
	0x10, 0x3d, 0xa4, 0xe8, 0x91, 0x7b, 0xb9, 0x57, 0xec, 0xea, 0x48, 0x10, 0xdd, 0x44, 0x8f, 0xec,
	0x00, 0x81, 0x51, 0x6f, 0x94, 0xfc, 0xcf, 0xce, 0x1d, 0xb0, 0x33, 0x84, 0x76, 0xd3, 0xc0, 0xfd,
	0xdc, 0x2f, 0x92, 0x4b, 0x5a, 0x2e, 0x37, 0xcb, 0xeb, 0x34, 0xaf, 0x67, 0xcd, 0x2a, 0x80, 0xb5,
	0xc3, 0xc0, 0xff, 0x6d, 0x78, 0xbf, 0x09, 0x2f, 0xae, 0xde, 0xc4, 0x18, 0x87, 0x68, 0x44, 0x3d,
	0xb4, 0x4a, 0xf2, 0xc0, 0xb5, 0x99, 0x91, 0x65, 0x10, 0x53, 0x27, 0xcc, 0x4b, 0xe9, 0x9e, 0x87,
	0x56, 0xad, 0xfc, 0x0c, 0xed, 0xd6, 0xea, 0x17, 0x00, 0x00, 0xff, 0xff, 0x98, 0xec, 0x38, 0x25,
	0x0c, 0x01, 0x00, 0x00,
}