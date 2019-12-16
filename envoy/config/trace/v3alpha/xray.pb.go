// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v3alpha/xray.proto

package envoy_config_trace_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type XRayConfig struct {
	DaemonEndpoint       *core.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	SegmentName          string              `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	SamplingRuleManifest *core.DataSource    `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb0fe92938f4aa2, []int{0}
}

func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRayConfig.Unmarshal(m, b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return xxx_messageInfo_XRayConfig.Size(m)
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() *core.SocketAddress {
	if m != nil {
		return m.DaemonEndpoint
	}
	return nil
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *core.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
}

func init() {
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v3alpha.XRayConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v3alpha/xray.proto", fileDescriptor_fdb0fe92938f4aa2)
}

var fileDescriptor_fdb0fe92938f4aa2 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0x7f, 0xf8, 0xa1, 0xa9, 0x28, 0x04, 0xd1, 0xd2, 0x85, 0xb4, 0xc5, 0x62, 0x17,
	0x32, 0x03, 0xed, 0x42, 0x70, 0x67, 0xd5, 0xa5, 0xa5, 0xa4, 0x9b, 0xee, 0xc2, 0x6d, 0x72, 0x1b,
	0x07, 0x93, 0xb9, 0xc3, 0xcc, 0x24, 0x34, 0x4f, 0xa0, 0xcf, 0xe0, 0xfb, 0xf8, 0x5e, 0xd2, 0x4c,
	0xa2, 0x20, 0x74, 0x17, 0x32, 0xdf, 0x39, 0xf7, 0x9e, 0x7b, 0xfc, 0x09, 0xca, 0x92, 0x2a, 0x1e,
	0x93, 0xdc, 0x89, 0x94, 0x5b, 0x0d, 0x31, 0xf2, 0x72, 0x0e, 0x99, 0x7a, 0x05, 0xbe, 0xd7, 0x50,
	0x31, 0xa5, 0xc9, 0x52, 0x30, 0xa8, 0x31, 0xe6, 0x30, 0x56, 0x63, 0xac, 0xc1, 0x06, 0xd7, 0xce,
	0x02, 0x94, 0xf8, 0x51, 0xc6, 0xa4, 0x91, 0x43, 0x92, 0x68, 0x34, 0xc6, 0x39, 0x0c, 0x46, 0x47,
	0xa8, 0x2d, 0x18, 0x6c, 0x91, 0x22, 0x51, 0xc0, 0x41, 0x4a, 0xb2, 0x60, 0x05, 0x49, 0xc3, 0x4b,
	0xd4, 0x46, 0x90, 0x14, 0x32, 0x6d, 0x90, 0xcb, 0x12, 0x32, 0x91, 0x80, 0x45, 0xde, 0x7e, 0xb8,
	0x87, 0xf1, 0x7b, 0xc7, 0xf7, 0x37, 0x21, 0x54, 0x8f, 0xf5, 0x86, 0xc1, 0xd2, 0x3f, 0x4b, 0x00,
	0x73, 0x92, 0x11, 0xca, 0x44, 0x91, 0x90, 0xb6, 0xef, 0x0d, 0xbd, 0x69, 0x6f, 0x36, 0x61, 0x2e,
	0x09, 0x28, 0xd1, 0x06, 0x60, 0x87, 0x3d, 0xd8, 0x9a, 0xe2, 0x37, 0xb4, 0x0f, 0x6e, 0xe7, 0xf0,
	0xd4, 0xa9, 0x9f, 0x1b, 0x71, 0x30, 0xf2, 0x4f, 0x0c, 0xa6, 0x39, 0x4a, 0x1b, 0x49, 0xc8, 0xb1,
	0xdf, 0x19, 0x7a, 0xd3, 0x6e, 0xd8, 0x6b, 0xfe, 0x2d, 0x21, 0xc7, 0x60, 0xe3, 0x5f, 0x18, 0xc8,
	0x55, 0x26, 0x64, 0x1a, 0xe9, 0x22, 0xc3, 0x28, 0x07, 0x29, 0x76, 0x68, 0x6c, 0xff, 0x5f, 0x3d,
	0x79, 0x7c, 0x6c, 0xf2, 0x13, 0x58, 0x58, 0x53, 0xa1, 0x63, 0x0c, 0xcf, 0x5b, 0x87, 0xb0, 0xc8,
	0xf0, 0xa5, 0xd1, 0xdf, 0xdf, 0x7e, 0x7e, 0x7d, 0x5c, 0xdd, 0x34, 0x55, 0xfd, 0xe9, 0x60, 0xe6,
	0x8c, 0x7e, 0xa3, 0x2f, 0xee, 0xfc, 0xa9, 0x20, 0x37, 0x4b, 0x69, 0xda, 0x57, 0xec, 0x78, 0x75,
	0x8b, 0xee, 0x46, 0x43, 0xb5, 0x3a, 0x1c, 0x70, 0xe5, 0x6d, 0xff, 0xd7, 0x97, 0x9c, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x18, 0x78, 0xc0, 0x13, 0x02, 0x00, 0x00,
}