// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v2/metrics_service.proto

package envoy_config_metrics_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
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

type MetricsServiceConfig struct {
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricsServiceConfig) Reset()         { *m = MetricsServiceConfig{} }
func (m *MetricsServiceConfig) String() string { return proto.CompactTextString(m) }
func (*MetricsServiceConfig) ProtoMessage()    {}
func (*MetricsServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_81ac893a597f6d53, []int{0}
}

func (m *MetricsServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsServiceConfig.Unmarshal(m, b)
}
func (m *MetricsServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsServiceConfig.Marshal(b, m, deterministic)
}
func (m *MetricsServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsServiceConfig.Merge(m, src)
}
func (m *MetricsServiceConfig) XXX_Size() int {
	return xxx_messageInfo_MetricsServiceConfig.Size(m)
}
func (m *MetricsServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsServiceConfig proto.InternalMessageInfo

func (m *MetricsServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsServiceConfig)(nil), "envoy.config.metrics.v2.MetricsServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v2/metrics_service.proto", fileDescriptor_81ac893a597f6d53)
}

var fileDescriptor_81ac893a597f6d53 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e,
	0xd6, 0x2f, 0x33, 0x82, 0x31, 0xe3, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xc1, 0xca, 0xf5, 0x20, 0xca, 0xf5, 0xa0, 0x6a, 0xf4, 0xca, 0x8c,
	0xa4, 0x54, 0x20, 0xe6, 0x24, 0x16, 0x64, 0x82, 0x34, 0x27, 0xe7, 0x17, 0xa5, 0xea, 0xa7, 0x17,
	0x15, 0x24, 0xa3, 0x6a, 0x97, 0x92, 0x2d, 0x4d, 0x29, 0x48, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x86, 0x4a, 0x8b,
	0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x64, 0x2e,
	0x11, 0x5f, 0x88, 0x5d, 0xc1, 0x10, 0xf3, 0x9c, 0xc1, 0x0e, 0x10, 0xf2, 0xe6, 0xe2, 0x41, 0xb6,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x4e, 0x0f, 0xe2, 0xca, 0xc4, 0x82, 0x4c, 0xbd,
	0x32, 0x23, 0x3d, 0x90, 0x63, 0xf4, 0xdc, 0x8b, 0x0a, 0x92, 0xa1, 0x7a, 0x9d, 0x38, 0x7e, 0x39,
	0xb1, 0x76, 0x31, 0x32, 0x09, 0x30, 0x06, 0x71, 0xa7, 0x23, 0x09, 0xbb, 0xed, 0x6a, 0x38, 0x71,
	0x91, 0x8d, 0x49, 0x80, 0x91, 0x4b, 0x35, 0x33, 0x1f, 0x62, 0x44, 0x41, 0x51, 0x7e, 0x45, 0xa5,
	0x1e, 0x0e, 0x3f, 0x3b, 0x09, 0xa3, 0xba, 0x29, 0x00, 0xe4, 0xd4, 0x00, 0xc6, 0x24, 0x36, 0xb0,
	0x9b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x07, 0xfa, 0x60, 0x5b, 0x01, 0x00, 0x00,
}
