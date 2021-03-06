// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/tap/v2alpha/common.proto

package envoy_config_common_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/service/tap/v2alpha"
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

type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0}
}

func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig.Size(m)
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v2alpha.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

type CommonExtensionConfig_TapDSConfig struct {
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0, 0}
}

func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Size(m)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AdminConfig struct {
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{1}
}

func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminConfig.Unmarshal(m, b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return xxx_messageInfo_AdminConfig.Size(m)
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.config.common.tap.v2alpha.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/tap/v2alpha/common.proto", fileDescriptor_79cf139d98a2fe3f)
}

var fileDescriptor_79cf139d98a2fe3f = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xca, 0xd3, 0x30,
	0x14, 0xc7, 0xed, 0xf6, 0x7d, 0xfa, 0x2d, 0xed, 0x40, 0x02, 0xa2, 0x54, 0x70, 0x63, 0x6e, 0xe2,
	0xc5, 0x4c, 0xa0, 0x7b, 0x02, 0x33, 0x05, 0x45, 0x90, 0xb9, 0x79, 0x3f, 0x8e, 0x6d, 0x9c, 0x81,
	0x35, 0x09, 0x6d, 0x56, 0x36, 0x1f, 0xc1, 0x1b, 0x6f, 0x7d, 0x0c, 0x9f, 0xcd, 0x2b, 0xd9, 0x85,
	0x48, 0x93, 0xd4, 0x75, 0xe0, 0xc7, 0xee, 0xda, 0x9c, 0xff, 0xf9, 0x9d, 0xff, 0xf9, 0x27, 0x68,
	0xca, 0x65, 0xa5, 0x0e, 0x34, 0x55, 0xf2, 0xb3, 0xd8, 0xd0, 0x54, 0xe5, 0xb9, 0x92, 0xd4, 0x80,
	0xa6, 0x55, 0x02, 0x5b, 0xfd, 0x05, 0xfc, 0x11, 0xd1, 0x85, 0x32, 0x0a, 0x0f, 0xac, 0x9a, 0x38,
	0x35, 0xf1, 0x25, 0x03, 0x9a, 0x78, 0x75, 0x3c, 0x71, 0x38, 0xd0, 0x82, 0x56, 0x09, 0x4d, 0x55,
	0xc1, 0x3d, 0x7a, 0x5d, 0xaa, 0x5d, 0x91, 0x72, 0xc7, 0x89, 0x9f, 0x39, 0x59, 0xc9, 0x8b, 0x4a,
	0xa4, 0xfc, 0xd6, 0x79, 0xf1, 0x93, 0x5d, 0xa6, 0x81, 0x82, 0x94, 0xca, 0x80, 0x11, 0x4a, 0x96,
	0x34, 0x17, 0x9b, 0x02, 0x4c, 0xc3, 0x79, 0x58, 0xc1, 0x56, 0x64, 0x60, 0x38, 0x6d, 0x3e, 0x5c,
	0x61, 0xf4, 0xb3, 0x8b, 0x1e, 0xcc, 0x2d, 0xe9, 0xf5, 0xde, 0x70, 0x59, 0x0a, 0x25, 0xe7, 0xd6,
	0x07, 0xfe, 0x80, 0x22, 0xc8, 0x72, 0x21, 0xd7, 0xce, 0xd7, 0xa3, 0x60, 0x18, 0x3c, 0x0f, 0x93,
	0x29, 0xb9, 0xb0, 0x19, 0x79, 0x59, 0x37, 0x39, 0xc6, 0x9b, 0x3b, 0xcb, 0x10, 0x4e, 0xbf, 0xf8,
	0x1d, 0xea, 0x97, 0xb5, 0xbf, 0xb4, 0x61, 0x76, 0x2c, 0x73, 0xec, 0x99, 0x7e, 0xcb, 0x33, 0xda,
	0x47, 0xd0, 0xff, 0x58, 0x91, 0x6b, 0xf6, 0xb0, 0x0d, 0x8a, 0x0c, 0xe8, 0xac, 0x6c, 0x58, 0x5d,
	0xcb, 0x62, 0x17, 0xfd, 0xfd, 0x77, 0xdb, 0x7a, 0xce, 0xab, 0xd5, 0xc9, 0xb5, 0x25, 0xbb, 0xdf,
	0xf8, 0x2b, 0x0a, 0x5b, 0x55, 0xfc, 0x1e, 0xf5, 0xcf, 0x6e, 0xca, 0x07, 0x33, 0xf0, 0x83, 0x41,
	0x0b, 0x52, 0x25, 0xa4, 0xbe, 0x51, 0xe2, 0x3a, 0x56, 0x56, 0xc6, 0x6e, 0x8e, 0xec, 0xfa, 0x5b,
	0xd0, 0xb9, 0x1f, 0x2c, 0xa3, 0xb4, 0x75, 0x8e, 0x1f, 0xa3, 0x2b, 0x09, 0x39, 0xb7, 0x59, 0xf4,
	0xd8, 0xbd, 0x23, 0xbb, 0x2a, 0x3a, 0xc3, 0x60, 0x69, 0x0f, 0x19, 0x46, 0xa1, 0x1f, 0x66, 0x0e,
	0x9a, 0xe3, 0xee, 0x6f, 0x16, 0x8c, 0x66, 0x28, 0x6c, 0x65, 0x8c, 0xc7, 0xa8, 0xe7, 0x25, 0x22,
	0xb3, 0x5e, 0x5a, 0x90, 0x1b, 0x57, 0x79, 0x9b, 0x31, 0xf1, 0xeb, 0xc7, 0x9f, 0xef, 0xd7, 0x13,
	0xfc, 0xd4, 0xb9, 0xe4, 0xcd, 0xfa, 0xe5, 0x59, 0x44, 0x33, 0x1b, 0x11, 0x7a, 0x21, 0x94, 0xdb,
	0x46, 0x17, 0x6a, 0x7f, 0xb8, 0x94, 0x28, 0x0b, 0x5d, 0xa4, 0x8b, 0xfa, 0x41, 0x2d, 0x82, 0x4f,
	0x77, 0xed, 0xcb, 0x9a, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x49, 0x0c, 0xba, 0x52, 0x32, 0x03,
	0x00, 0x00,
}
