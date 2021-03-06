// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/cors/v2/cors.proto

package envoy_config_filter_http_cors_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Cors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cors) Reset()         { *m = Cors{} }
func (m *Cors) String() string { return proto.CompactTextString(m) }
func (*Cors) ProtoMessage()    {}
func (*Cors) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba80e325986c17c, []int{0}
}

func (m *Cors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cors.Unmarshal(m, b)
}
func (m *Cors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cors.Marshal(b, m, deterministic)
}
func (m *Cors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cors.Merge(m, src)
}
func (m *Cors) XXX_Size() int {
	return xxx_messageInfo_Cors.Size(m)
}
func (m *Cors) XXX_DiscardUnknown() {
	xxx_messageInfo_Cors.DiscardUnknown(m)
}

var xxx_messageInfo_Cors proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cors)(nil), "envoy.config.filter.http.cors.v2.Cors")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/cors/v2/cors.proto", fileDescriptor_8ba80e325986c17c)
}

var fileDescriptor_8ba80e325986c17c = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xce, 0x31, 0x0a, 0xc2, 0x30,
	0x14, 0xc6, 0x71, 0x04, 0x2d, 0xd8, 0xb1, 0xa3, 0x83, 0x88, 0xa3, 0xca, 0x7b, 0xd0, 0xde, 0xa0,
	0x5e, 0xa0, 0x57, 0x88, 0x35, 0x6d, 0x03, 0x35, 0x2f, 0x24, 0xcf, 0xd0, 0x9e, 0xc2, 0xd5, 0x73,
	0x7a, 0x00, 0x91, 0x24, 0x0e, 0x6e, 0x4e, 0x6f, 0xf9, 0xf1, 0xff, 0x5e, 0x7e, 0x94, 0xda, 0xd3,
	0x8c, 0x2d, 0xe9, 0x4e, 0xf5, 0xd8, 0xa9, 0x91, 0xa5, 0xc5, 0x81, 0xd9, 0x60, 0x4b, 0xd6, 0xa1,
	0x2f, 0xe3, 0x05, 0x63, 0x89, 0xa9, 0xd8, 0x45, 0x0c, 0x09, 0x43, 0xc2, 0x10, 0x30, 0x44, 0xe4,
	0xcb, 0xcd, 0xf6, 0x7e, 0x35, 0x02, 0x85, 0xd6, 0xc4, 0x82, 0x15, 0x69, 0x87, 0x37, 0xd5, 0x5b,
	0xc1, 0x32, 0x15, 0xf6, 0x59, 0xbe, 0x3c, 0x93, 0x75, 0x35, 0xbd, 0x9e, 0xef, 0xc7, 0xea, 0x54,
	0x1c, 0x52, 0x51, 0x4e, 0x2c, 0xb5, 0x0b, 0xfe, 0x5b, 0x75, 0xbf, 0xd9, 0x4a, 0x8c, 0x66, 0x10,
	0x39, 0x28, 0x82, 0xc8, 0x8d, 0xa5, 0x69, 0x86, 0x7f, 0xbf, 0xd4, 0xeb, 0xb0, 0xd4, 0x84, 0xd9,
	0x66, 0x71, 0xc9, 0xe2, 0x7e, 0xf5, 0x09, 0x00, 0x00, 0xff, 0xff, 0x70, 0x93, 0x6d, 0x01, 0xf0,
	0x00, 0x00, 0x00,
}
