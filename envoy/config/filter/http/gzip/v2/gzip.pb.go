// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/gzip/v2/gzip.proto

package envoy_config_filter_http_gzip_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2b7bd5745a19167, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2b7bd5745a19167, []int{0, 0, 0}
}

type Gzip struct {
	MemoryLevel                *wrappers.UInt32Value      `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	ContentLength              *wrappers.UInt32Value      `protobuf:"bytes,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	CompressionLevel           Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	CompressionStrategy        Gzip_CompressionStrategy   `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	ContentType                []string                   `protobuf:"bytes,6,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	DisableOnEtagHeader        bool                       `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	RemoveAcceptEncodingHeader bool                       `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	WindowBits                 *wrappers.UInt32Value      `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                   `json:"-"`
	XXX_unrecognized           []byte                     `json:"-"`
	XXX_sizecache              int32                      `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2b7bd5745a19167, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2b7bd5745a19167, []int{0, 0}
}

func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.config.filter.http.gzip.v2.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/gzip/v2/gzip.proto", fileDescriptor_e2b7bd5745a19167)
}

var fileDescriptor_e2b7bd5745a19167 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xdb, 0xb5, 0x8d, 0x3b, 0x46, 0xf0, 0x10, 0x44, 0x15, 0x4c, 0xd5, 0x4e, 0xd1,
	0x86, 0x1c, 0xa9, 0xbd, 0xa1, 0x5d, 0x1a, 0x96, 0xb2, 0xa1, 0x02, 0x53, 0xd6, 0x71, 0x8d, 0xdc,
	0xf4, 0x2d, 0xb5, 0x94, 0xda, 0x56, 0xe2, 0xb6, 0xcb, 0x8e, 0xf0, 0x07, 0x20, 0x71, 0xe2, 0xef,
	0xe4, 0xb8, 0x03, 0x42, 0x71, 0x32, 0x18, 0x03, 0x69, 0xe3, 0xd4, 0x4a, 0xdf, 0xf7, 0x7b, 0x9f,
	0xf3, 0xbd, 0x87, 0xf6, 0x81, 0xaf, 0x44, 0xee, 0x46, 0x82, 0x9f, 0xb3, 0xd8, 0x3d, 0x67, 0x89,
	0x82, 0xd4, 0x9d, 0x2b, 0x25, 0xdd, 0xf8, 0x92, 0x49, 0x77, 0xd5, 0xd7, 0xbf, 0x44, 0xa6, 0x42,
	0x09, 0xdc, 0xd3, 0x66, 0x52, 0x9a, 0x49, 0x69, 0x26, 0x85, 0x99, 0x68, 0xd3, 0xaa, 0xdf, 0xdd,
	0x89, 0x85, 0x88, 0x13, 0x70, 0xb5, 0x7f, 0xba, 0x3c, 0x77, 0xd7, 0x29, 0x95, 0x12, 0xd2, 0xac,
	0x9c, 0xd0, 0xdd, 0x59, 0xce, 0x24, 0x75, 0x29, 0xe7, 0x42, 0x51, 0xc5, 0x04, 0xcf, 0xdc, 0x05,
	0x8b, 0x53, 0xaa, 0xa0, 0xd2, 0x9f, 0xad, 0x68, 0xc2, 0x66, 0x54, 0x81, 0x7b, 0xfd, 0xa7, 0x14,
	0x76, 0x3f, 0x37, 0x51, 0xe3, 0xcd, 0x25, 0x93, 0xf8, 0x2d, 0xda, 0x5c, 0xc0, 0x42, 0xa4, 0x79,
	0x98, 0xc0, 0x0a, 0x12, 0xdb, 0xe8, 0x19, 0x4e, 0xa7, 0xff, 0x9c, 0x94, 0xc1, 0xe4, 0x3a, 0x98,
	0x9c, 0x1d, 0x73, 0x35, 0xe8, 0x7f, 0xa4, 0xc9, 0x12, 0x3c, 0xf3, 0xca, 0x6b, 0xee, 0x35, 0x6c,
	0xd3, 0x31, 0x82, 0x4e, 0x09, 0x8f, 0x0b, 0x16, 0x8f, 0xd1, 0x56, 0x24, 0xb8, 0x02, 0xae, 0xc2,
	0x04, 0x78, 0xac, 0xe6, 0x76, 0xed, 0x1e, 0xd3, 0x5a, 0x57, 0x5e, 0x63, 0xaf, 0xe6, 0xec, 0x04,
	0x0f, 0x2b, 0x78, 0xac, 0x59, 0x9c, 0xa1, 0xc7, 0x91, 0x58, 0xc8, 0x14, 0xb2, 0x8c, 0x09, 0x5e,
	0x3d, 0xaf, 0xde, 0x33, 0x9c, 0xad, 0xfe, 0x01, 0xb9, 0xab, 0x39, 0x52, 0x7c, 0x1c, 0x79, 0xfd,
	0x9b, 0xd7, 0x2f, 0x24, 0x3e, 0x5f, 0x2e, 0xbc, 0xf6, 0x95, 0xb7, 0xf1, 0xc9, 0xa8, 0x59, 0x46,
	0x60, 0x45, 0xb7, 0x0c, 0x78, 0x8d, 0x9e, 0xdc, 0x0c, 0xcd, 0x54, 0x51, 0x66, 0x9c, 0xdb, 0x0d,
	0x9d, 0xfb, 0xea, 0xff, 0x73, 0x4f, 0xab, 0x09, 0x37, 0x52, 0xb7, 0xa3, 0xbf, 0x65, 0xbc, 0x8f,
	0x36, 0xaf, 0xbb, 0x53, 0xb9, 0x04, 0xbb, 0xd9, 0xab, 0x3b, 0xa6, 0x86, 0xbe, 0x1a, 0x35, 0xab,
	0x1f, 0x74, 0x2a, 0x75, 0x92, 0x4b, 0xc0, 0x03, 0xf4, 0x74, 0xc6, 0x32, 0x3a, 0x4d, 0x20, 0x14,
	0x3c, 0x04, 0x45, 0xe3, 0x70, 0x0e, 0x74, 0x06, 0xa9, 0xdd, 0xea, 0x19, 0x4e, 0x3b, 0xd8, 0xae,
	0xd4, 0x0f, 0xdc, 0x57, 0x34, 0x3e, 0xd2, 0x12, 0x1e, 0xa2, 0x17, 0x29, 0x2c, 0xc4, 0x0a, 0x42,
	0x1a, 0x45, 0x20, 0x55, 0x08, 0x3c, 0x12, 0x33, 0xc6, 0x7f, 0xb1, 0x6d, 0xcd, 0x76, 0x4b, 0xd3,
	0x50, 0x7b, 0xfc, 0xca, 0x52, 0x8d, 0x38, 0x42, 0x9d, 0x35, 0xe3, 0x33, 0xb1, 0x0e, 0xa7, 0x4c,
	0x65, 0xb6, 0x79, 0xff, 0x5b, 0x79, 0xe4, 0x98, 0x01, 0x2a, 0x59, 0x8f, 0xa9, 0xac, 0x7b, 0x80,
	0xac, 0xdb, 0xcb, 0xd9, 0x75, 0x50, 0xa3, 0xd8, 0x0f, 0xee, 0xa0, 0xd6, 0xa1, 0x3f, 0x1a, 0x9e,
	0x8d, 0x27, 0xd6, 0x03, 0xdc, 0x46, 0x0d, 0xcf, 0x3f, 0x9d, 0x58, 0x06, 0x36, 0xd1, 0xc6, 0xe9,
	0x89, 0xef, 0x1f, 0x5a, 0xb5, 0xdd, 0x11, 0xda, 0xfe, 0x47, 0xc5, 0x7f, 0x82, 0x9b, 0xa8, 0x3d,
	0x3a, 0x1e, 0x4f, 0xfc, 0xc0, 0x3f, 0xb4, 0x8c, 0x42, 0x3a, 0x3a, 0x1b, 0x8d, 0xde, 0x0d, 0xdf,
	0x5b, 0x35, 0xdc, 0x42, 0xf5, 0x60, 0xec, 0x5b, 0x75, 0x4f, 0x7c, 0xff, 0xf6, 0xe3, 0xcb, 0xc6,
	0x4b, 0xbc, 0x57, 0xae, 0x15, 0x2e, 0x14, 0xf0, 0x62, 0x62, 0x56, 0xad, 0x36, 0xbb, 0xb9, 0xdb,
	0x01, 0x4d, 0xe4, 0x9c, 0x22, 0xc2, 0x44, 0x79, 0x05, 0x32, 0x15, 0x17, 0xf9, 0x9d, 0x07, 0xe1,
	0x99, 0xc5, 0x45, 0x9c, 0x14, 0xd5, 0x9c, 0x18, 0xd3, 0xa6, 0xee, 0x68, 0xf0, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x25, 0x17, 0xf0, 0x7d, 0x27, 0x04, 0x00, 0x00,
}
