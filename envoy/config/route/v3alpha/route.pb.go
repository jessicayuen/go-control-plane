// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/route/v3alpha/route.proto

package envoy_config_route_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
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

type RouteConfiguration struct {
	Name                            string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VirtualHosts                    []*VirtualHost               `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Vhds                            *Vhds                        `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	InternalOnlyHeaders             []string                     `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	ResponseHeadersToAdd            []*v3alpha.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	ResponseHeadersToRemove         []string                     `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	RequestHeadersToAdd             []*v3alpha.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove          []string                     `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	MostSpecificHeaderMutationsWins bool                         `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	ValidateClusters                *wrappers.BoolValue          `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                     `json:"-"`
	XXX_unrecognized                []byte                       `json:"-"`
	XXX_sizecache                   int32                        `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_368555498583d647, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*v3alpha.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*v3alpha.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if m != nil {
		return m.MostSpecificHeaderMutationsWins
	}
	return false
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	ConfigSource         *v3alpha.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_368555498583d647, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *v3alpha.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.route.v3alpha.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.config.route.v3alpha.Vhds")
}

func init() {
	proto.RegisterFile("envoy/config/route/v3alpha/route.proto", fileDescriptor_368555498583d647)
}

var fileDescriptor_368555498583d647 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x52, 0xd4, 0x4c,
	0x14, 0xad, 0x30, 0xc3, 0xcf, 0x34, 0x50, 0x05, 0xcd, 0xf7, 0x41, 0x9c, 0x05, 0x04, 0xb4, 0x30,
	0x0b, 0x4d, 0xca, 0xc1, 0x85, 0xe2, 0xca, 0xb0, 0x90, 0x05, 0x16, 0x54, 0xb0, 0x70, 0x99, 0x6a,
	0x92, 0x66, 0xa6, 0xab, 0x32, 0x7d, 0x63, 0x77, 0x27, 0x30, 0x6f, 0x60, 0xb9, 0x74, 0xe9, 0xfb,
	0xf8, 0x2e, 0xbe, 0x80, 0x1b, 0x56, 0x56, 0xba, 0x3b, 0x62, 0x18, 0xc0, 0x85, 0xbb, 0x74, 0xdf,
	0x73, 0xcf, 0x39, 0xf7, 0xf6, 0x09, 0xda, 0xa5, 0xbc, 0x82, 0x49, 0x98, 0x02, 0xbf, 0x60, 0xc3,
	0x50, 0x40, 0xa9, 0x68, 0x58, 0xed, 0x91, 0xbc, 0x18, 0x11, 0x73, 0x0a, 0x0a, 0x01, 0x0a, 0x70,
	0x5f, 0xe3, 0x02, 0x83, 0x0b, 0x4c, 0xc5, 0xe2, 0xfa, 0x4f, 0x5a, 0x1c, 0x29, 0x88, 0x1b, 0x8a,
	0x73, 0x22, 0x2d, 0x43, 0xff, 0xf9, 0xfd, 0x28, 0x73, 0x97, 0x48, 0x28, 0x45, 0xda, 0xc0, 0x5f,
	0xfc, 0xcd, 0x58, 0x92, 0xc2, 0xb8, 0x00, 0x4e, 0xb9, 0x92, 0xb6, 0x65, 0x73, 0x08, 0x30, 0xcc,
	0x69, 0xa8, 0x4f, 0xe7, 0xe5, 0x45, 0x78, 0x29, 0x48, 0x51, 0x50, 0xd1, 0xd4, 0xb7, 0xcb, 0xac,
	0x20, 0x21, 0xe1, 0x1c, 0x14, 0x51, 0x0c, 0xb8, 0x0c, 0x2b, 0x2a, 0x24, 0x03, 0xce, 0xf8, 0xd0,
	0x42, 0x36, 0x2a, 0x92, 0xb3, 0x8c, 0xd4, 0x5a, 0xf6, 0xc3, 0x14, 0x76, 0x7e, 0xce, 0x22, 0x1c,
	0xd7, 0xb2, 0x07, 0xda, 0x50, 0x29, 0x34, 0x03, 0xc6, 0xa8, 0xcb, 0xc9, 0x98, 0xba, 0x8e, 0xe7,
	0xf8, 0xbd, 0x58, 0x7f, 0xe3, 0x23, 0xb4, 0x5c, 0x31, 0xa1, 0x4a, 0x92, 0x27, 0x23, 0x90, 0x4a,
	0xba, 0x33, 0x5e, 0xc7, 0x5f, 0x1c, 0x3c, 0x0d, 0xee, 0x5f, 0x61, 0x70, 0x66, 0x1a, 0x0e, 0x41,
	0xaa, 0x78, 0xa9, 0xba, 0x39, 0x48, 0xfc, 0x12, 0x75, 0xab, 0x51, 0x26, 0xdd, 0x9e, 0xe7, 0xf8,
	0x8b, 0x03, 0xef, 0x41, 0x92, 0x51, 0x26, 0x63, 0x8d, 0xc6, 0x03, 0xf4, 0x3f, 0xe3, 0x8a, 0x0a,
	0x4e, 0xf2, 0x04, 0x78, 0x3e, 0x49, 0x46, 0x94, 0x64, 0x54, 0x48, 0xb7, 0xe3, 0x75, 0xfc, 0x5e,
	0xbc, 0xd6, 0x14, 0x8f, 0x79, 0x3e, 0x39, 0x34, 0x25, 0xcc, 0xd1, 0x86, 0xa0, 0xb2, 0x00, 0x2e,
	0x69, 0x03, 0x4f, 0x14, 0x24, 0x24, 0xcb, 0xdc, 0xae, 0x9e, 0xe0, 0x59, 0x5b, 0xbc, 0x7e, 0xc2,
	0xdf, 0xda, 0x86, 0xe4, 0x8c, 0xe4, 0x25, 0x3d, 0x2e, 0xea, 0xd5, 0x44, 0xbd, 0xeb, 0x68, 0xee,
	0xab, 0xd3, 0x59, 0xf9, 0x31, 0x1f, 0xff, 0xd7, 0xf0, 0x5a, 0xa9, 0x0f, 0xf0, 0x36, 0xcb, 0xf0,
	0x1b, 0xd4, 0xbf, 0x4b, 0x4f, 0xd0, 0x31, 0x54, 0xd4, 0x9d, 0xd5, 0x46, 0x37, 0xa6, 0x3a, 0x63,
	0x5d, 0xc6, 0x39, 0x5a, 0x17, 0xf4, 0x53, 0x49, 0xa5, 0xba, 0xed, 0x75, 0xee, 0xdf, 0xbc, 0xae,
	0x59, 0xda, 0x96, 0xd5, 0xd7, 0xe8, 0xd1, 0x1d, 0x6a, 0xd6, 0xe9, 0x82, 0x76, 0xba, 0x7e, 0xbb,
	0xcf, 0x1a, 0x3d, 0x42, 0x8f, 0xc7, 0x20, 0x55, 0x22, 0x0b, 0x9a, 0xb2, 0x0b, 0x96, 0x5a, 0x82,
	0x64, 0x5c, 0xda, 0x10, 0x26, 0x97, 0x8c, 0x4b, 0x17, 0x79, 0x8e, 0xbf, 0x10, 0x6f, 0xd5, 0xd0,
	0x53, 0x8b, 0x34, 0x4c, 0xef, 0x1b, 0xdc, 0x47, 0xc6, 0x25, 0x7e, 0x87, 0x56, 0x9b, 0x60, 0x26,
	0x69, 0x5e, 0x4a, 0x55, 0xbf, 0xe9, 0xbc, 0x8e, 0x46, 0x3f, 0x30, 0xf1, 0x0f, 0x9a, 0xf8, 0x07,
	0x11, 0x40, 0xae, 0xa7, 0x8c, 0x57, 0x9a, 0xa6, 0x03, 0xdb, 0xb3, 0xbf, 0xfb, 0xed, 0xfb, 0xe7,
	0xcd, 0x6d, 0xb4, 0x65, 0xb6, 0x44, 0x0a, 0x16, 0x54, 0x83, 0x60, 0x3a, 0xe0, 0x3b, 0x57, 0xa8,
	0x5b, 0xc7, 0x0a, 0x9f, 0xa1, 0xe5, 0xd6, 0x5f, 0xaa, 0x13, 0x3f, 0x15, 0xea, 0xd6, 0x9a, 0x0d,
	0xd1, 0xa9, 0x86, 0x47, 0x0b, 0xd7, 0xd1, 0xec, 0x17, 0x67, 0x66, 0xc5, 0x89, 0x97, 0xd2, 0x3f,
	0xee, 0xf7, 0xdd, 0xda, 0xc7, 0x1a, 0x5a, 0x6d, 0xf9, 0xa8, 0x15, 0xa3, 0x57, 0xc8, 0x67, 0x60,
	0xe8, 0x0b, 0x01, 0x57, 0x93, 0x07, 0x92, 0x1f, 0x21, 0xed, 0xfc, 0xa4, 0x1e, 0xfc, 0xc4, 0x39,
	0x9f, 0xd3, 0x1b, 0xd8, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x85, 0x51, 0x52, 0xdc, 0x04,
	0x00, 0x00,
}
