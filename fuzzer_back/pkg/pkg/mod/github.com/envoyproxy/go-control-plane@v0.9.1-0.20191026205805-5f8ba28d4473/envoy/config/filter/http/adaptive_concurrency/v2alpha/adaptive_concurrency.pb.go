// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto

package envoy_config_filter_http_adaptive_concurrency_v2alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GradientControllerConfig struct {
	SampleAggregatePercentile *_type.Percent                                              `protobuf:"bytes,1,opt,name=sample_aggregate_percentile,json=sampleAggregatePercentile,proto3" json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    *GradientControllerConfig_ConcurrencyLimitCalculationParams `protobuf:"bytes,2,opt,name=concurrency_limit_params,json=concurrencyLimitParams,proto3" json:"concurrency_limit_params,omitempty"`
	MinRttCalcParams          *GradientControllerConfig_MinimumRTTCalculationParams       `protobuf:"bytes,3,opt,name=min_rtt_calc_params,json=minRttCalcParams,proto3" json:"min_rtt_calc_params,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                                    `json:"-"`
	XXX_unrecognized          []byte                                                      `json:"-"`
	XXX_sizecache             int32                                                       `json:"-"`
}

func (m *GradientControllerConfig) Reset()         { *m = GradientControllerConfig{} }
func (m *GradientControllerConfig) String() string { return proto.CompactTextString(m) }
func (*GradientControllerConfig) ProtoMessage()    {}
func (*GradientControllerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0}
}

func (m *GradientControllerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig.Unmarshal(m, b)
}
func (m *GradientControllerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig.Merge(m, src)
}
func (m *GradientControllerConfig) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig.Size(m)
}
func (m *GradientControllerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig proto.InternalMessageInfo

func (m *GradientControllerConfig) GetSampleAggregatePercentile() *_type.Percent {
	if m != nil {
		return m.SampleAggregatePercentile
	}
	return nil
}

func (m *GradientControllerConfig) GetConcurrencyLimitParams() *GradientControllerConfig_ConcurrencyLimitCalculationParams {
	if m != nil {
		return m.ConcurrencyLimitParams
	}
	return nil
}

func (m *GradientControllerConfig) GetMinRttCalcParams() *GradientControllerConfig_MinimumRTTCalculationParams {
	if m != nil {
		return m.MinRttCalcParams
	}
	return nil
}

type GradientControllerConfig_ConcurrencyLimitCalculationParams struct {
	MaxGradient               *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=max_gradient,json=maxGradient,proto3" json:"max_gradient,omitempty"`
	MaxConcurrencyLimit       *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrency_limit,json=maxConcurrencyLimit,proto3" json:"max_concurrency_limit,omitempty"`
	ConcurrencyUpdateInterval *duration.Duration    `protobuf:"bytes,3,opt,name=concurrency_update_interval,json=concurrencyUpdateInterval,proto3" json:"concurrency_update_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-"`
	XXX_unrecognized          []byte                `json:"-"`
	XXX_sizecache             int32                 `json:"-"`
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Reset() {
	*m = GradientControllerConfig_ConcurrencyLimitCalculationParams{}
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0, 0}
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Size(m)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxGradient() *wrappers.DoubleValue {
	if m != nil {
		return m.MaxGradient
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxConcurrencyLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrencyLimit
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetConcurrencyUpdateInterval() *duration.Duration {
	if m != nil {
		return m.ConcurrencyUpdateInterval
	}
	return nil
}

type GradientControllerConfig_MinimumRTTCalculationParams struct {
	Interval             *duration.Duration    `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	RequestCount         *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	Jitter               *_type.Percent        `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) Reset() {
	*m = GradientControllerConfig_MinimumRTTCalculationParams{}
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_MinimumRTTCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_MinimumRTTCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{0, 1}
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Size(m)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetRequestCount() *wrappers.UInt32Value {
	if m != nil {
		return m.RequestCount
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetJitter() *_type.Percent {
	if m != nil {
		return m.Jitter
	}
	return nil
}

type AdaptiveConcurrency struct {
	// Types that are valid to be assigned to ConcurrencyControllerConfig:
	//	*AdaptiveConcurrency_GradientControllerConfig
	ConcurrencyControllerConfig isAdaptiveConcurrency_ConcurrencyControllerConfig `protobuf_oneof:"concurrency_controller_config"`
	Enabled                     *core.RuntimeFeatureFlag                          `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                                          `json:"-"`
	XXX_unrecognized            []byte                                            `json:"-"`
	XXX_sizecache               int32                                             `json:"-"`
}

func (m *AdaptiveConcurrency) Reset()         { *m = AdaptiveConcurrency{} }
func (m *AdaptiveConcurrency) String() string { return proto.CompactTextString(m) }
func (*AdaptiveConcurrency) ProtoMessage()    {}
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_c58a0beecb0ec580, []int{1}
}

func (m *AdaptiveConcurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdaptiveConcurrency.Unmarshal(m, b)
}
func (m *AdaptiveConcurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdaptiveConcurrency.Marshal(b, m, deterministic)
}
func (m *AdaptiveConcurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdaptiveConcurrency.Merge(m, src)
}
func (m *AdaptiveConcurrency) XXX_Size() int {
	return xxx_messageInfo_AdaptiveConcurrency.Size(m)
}
func (m *AdaptiveConcurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AdaptiveConcurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AdaptiveConcurrency proto.InternalMessageInfo

type isAdaptiveConcurrency_ConcurrencyControllerConfig interface {
	isAdaptiveConcurrency_ConcurrencyControllerConfig()
}

type AdaptiveConcurrency_GradientControllerConfig struct {
	GradientControllerConfig *GradientControllerConfig `protobuf:"bytes,1,opt,name=gradient_controller_config,json=gradientControllerConfig,proto3,oneof"`
}

func (*AdaptiveConcurrency_GradientControllerConfig) isAdaptiveConcurrency_ConcurrencyControllerConfig() {
}

func (m *AdaptiveConcurrency) GetConcurrencyControllerConfig() isAdaptiveConcurrency_ConcurrencyControllerConfig {
	if m != nil {
		return m.ConcurrencyControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetGradientControllerConfig() *GradientControllerConfig {
	if x, ok := m.GetConcurrencyControllerConfig().(*AdaptiveConcurrency_GradientControllerConfig); ok {
		return x.GradientControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.Enabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdaptiveConcurrency) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdaptiveConcurrency_GradientControllerConfig)(nil),
	}
}

func init() {
	proto.RegisterType((*GradientControllerConfig)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig")
	proto.RegisterType((*GradientControllerConfig_ConcurrencyLimitCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig.ConcurrencyLimitCalculationParams")
	proto.RegisterType((*GradientControllerConfig_MinimumRTTCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.GradientControllerConfig.MinimumRTTCalculationParams")
	proto.RegisterType((*AdaptiveConcurrency)(nil), "envoy.config.filter.http.adaptive_concurrency.v2alpha.AdaptiveConcurrency")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto", fileDescriptor_c58a0beecb0ec580)
}

var fileDescriptor_c58a0beecb0ec580 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0x8e, 0x03, 0x4b, 0xd8, 0x81, 0x5d, 0x21, 0x47, 0xbb, 0x1b, 0x02, 0xbb, 0x02, 0xb4, 0x95,
	0x2a, 0x2a, 0x8d, 0xa5, 0xa0, 0x9e, 0x11, 0x4e, 0x45, 0x4b, 0xd5, 0x8a, 0xc8, 0x85, 0x4a, 0x3d,
	0x59, 0x2f, 0xce, 0xc3, 0x4c, 0x3b, 0x9e, 0x31, 0x93, 0x71, 0x9a, 0xfc, 0x85, 0xfe, 0x82, 0xf6,
	0xde, 0x53, 0xef, 0xed, 0x7f, 0xe9, 0xb9, 0x7f, 0xa0, 0xc7, 0x2a, 0xa7, 0xca, 0x9e, 0x31, 0x44,
	0x24, 0x50, 0x81, 0xc8, 0x29, 0xd6, 0x7b, 0xdf, 0xf7, 0xbe, 0xf7, 0xbd, 0xcf, 0x26, 0x1d, 0x14,
	0x03, 0x39, 0xf2, 0x22, 0x29, 0x4e, 0x58, 0xec, 0x9d, 0x30, 0xae, 0x51, 0x79, 0xa7, 0x5a, 0xa7,
	0x1e, 0xf4, 0x20, 0xd5, 0x6c, 0x80, 0x61, 0x24, 0x45, 0x94, 0x29, 0x85, 0x22, 0x1a, 0x79, 0x83,
	0x16, 0xf0, 0xf4, 0x14, 0x66, 0x16, 0x69, 0xaa, 0xa4, 0x96, 0xee, 0xc3, 0x82, 0x91, 0x1a, 0x46,
	0x6a, 0x18, 0x69, 0xce, 0x48, 0x67, 0x82, 0x2c, 0x63, 0x73, 0xdd, 0x08, 0x81, 0x94, 0x79, 0x83,
	0x96, 0x17, 0x49, 0x85, 0x5e, 0x17, 0xfa, 0x68, 0x48, 0x9b, 0x0d, 0x53, 0xd5, 0xa3, 0x14, 0xbd,
	0x14, 0x55, 0x84, 0x42, 0xdb, 0xca, 0x7a, 0x2c, 0x65, 0xcc, 0xb1, 0x00, 0x82, 0x10, 0x52, 0x83,
	0x66, 0x52, 0xf4, 0x6d, 0xf5, 0x3f, 0x5b, 0x2d, 0x9e, 0xba, 0xd9, 0x89, 0xd7, 0xcb, 0x54, 0xd1,
	0x70, 0x55, 0xfd, 0xad, 0x82, 0x34, 0x45, 0x55, 0xe2, 0xff, 0x19, 0x00, 0x67, 0x3d, 0xd0, 0xe8,
	0x95, 0x7f, 0x4c, 0x61, 0xeb, 0x5b, 0x8d, 0x34, 0x1e, 0x2b, 0xe8, 0x31, 0x14, 0xba, 0x2d, 0x85,
	0x56, 0x92, 0x73, 0x54, 0xed, 0x62, 0x69, 0xf7, 0x05, 0x59, 0xeb, 0x43, 0x92, 0x72, 0x0c, 0x21,
	0x8e, 0x15, 0xc6, 0xa0, 0x31, 0xb4, 0xaa, 0x19, 0xc7, 0x86, 0xb3, 0xe1, 0xdc, 0x5f, 0x6a, 0xd5,
	0xa9, 0x31, 0x2a, 0xdf, 0x89, 0x76, 0x4c, 0x35, 0x58, 0x35, 0xb8, 0xbd, 0x12, 0xd6, 0x39, 0x47,
	0xb9, 0x9f, 0x1d, 0xd2, 0x98, 0x30, 0x2e, 0xe4, 0x2c, 0x61, 0x3a, 0x4c, 0x41, 0x41, 0xd2, 0x6f,
	0x54, 0x0b, 0xca, 0x33, 0x7a, 0x2b, 0xef, 0xe9, 0x55, 0x8b, 0xd0, 0xf6, 0x45, 0xf3, 0xb3, 0x7c,
	0x5c, 0x1b, 0x78, 0x94, 0xf1, 0xc2, 0xc0, 0x4e, 0x31, 0xd8, 0x5f, 0x1c, 0xfb, 0xbf, 0xbd, 0x73,
	0xaa, 0x2b, 0x4e, 0xf0, 0x77, 0x74, 0xa9, 0xd9, 0x74, 0xb8, 0x1f, 0x1d, 0x52, 0x4f, 0x98, 0x08,
	0x95, 0xd6, 0x61, 0x04, 0x3c, 0x2a, 0x25, 0xcf, 0x15, 0x92, 0xdf, 0xdc, 0xb5, 0xe4, 0xe7, 0x4c,
	0xb0, 0x24, 0x4b, 0x82, 0xa3, 0xa3, 0xeb, 0xc4, 0xae, 0x24, 0x4c, 0x04, 0xba, 0xd8, 0xc7, 0xd4,
	0x9a, 0x5f, 0xaa, 0x64, 0xf3, 0x97, 0xeb, 0xba, 0x87, 0x64, 0x39, 0x81, 0x61, 0x18, 0xdb, 0xe9,
	0xf6, 0x94, 0xeb, 0xd4, 0xc4, 0x88, 0x96, 0x31, 0xa2, 0x8f, 0x64, 0xd6, 0xe5, 0xf8, 0x12, 0x78,
	0x86, 0xfe, 0x9f, 0x63, 0x7f, 0xc9, 0xfd, 0x7d, 0xb3, 0x52, 0xfc, 0xbe, 0xef, 0x06, 0x4b, 0x09,
	0x0c, 0x4b, 0xf9, 0xee, 0x2b, 0xf2, 0x57, 0x4e, 0x38, 0x75, 0x58, 0x7b, 0xd1, 0x69, 0xe6, 0xe3,
	0x03, 0xa1, 0x77, 0x5a, 0x86, 0xb9, 0x36, 0xf6, 0xe7, 0xb7, 0xab, 0x1b, 0x95, 0xa0, 0x9e, 0xc0,
	0xf0, 0xb2, 0x78, 0x17, 0xc9, 0xda, 0x24, 0x6d, 0x96, 0xe6, 0xf1, 0x0d, 0x99, 0xd0, 0xa8, 0x06,
	0xc0, 0xad, 0xff, 0xab, 0xd3, 0xd2, 0xed, 0x1b, 0xe2, 0x93, 0xb1, 0x5f, 0xfb, 0xe4, 0xcc, 0x2f,
	0x3a, 0xdb, 0x95, 0x60, 0x75, 0x82, 0xe9, 0xb8, 0x20, 0x3a, 0xb0, 0x3c, 0xcd, 0xaf, 0x0e, 0x59,
	0xbb, 0xc6, 0x74, 0x77, 0x8f, 0x2c, 0x9e, 0xcf, 0x74, 0x6e, 0x32, 0xf3, 0x1c, 0xe6, 0x3e, 0x25,
	0x7f, 0x28, 0x3c, 0xcb, 0xb0, 0xaf, 0xc3, 0x48, 0x66, 0xe2, 0x86, 0xe6, 0x2c, 0x5b, 0x6c, 0x3b,
	0x87, 0xba, 0x0f, 0xc8, 0xc2, 0x6b, 0xa6, 0x35, 0x2a, 0x6b, 0xc0, 0xcc, 0xd7, 0xd0, 0xb6, 0x6c,
	0xbd, 0xaf, 0x92, 0xfa, 0x9e, 0x8d, 0xe1, 0x84, 0xbf, 0xee, 0x07, 0x87, 0x34, 0xcb, 0x0c, 0xe4,
	0xb7, 0xb3, 0x11, 0x0c, 0x4d, 0x8a, 0xed, 0x9a, 0x87, 0x77, 0x1c, 0xed, 0x8b, 0xf8, 0x3e, 0xa9,
	0x04, 0x8d, 0xf8, 0xaa, 0x8f, 0xcf, 0x2e, 0xa9, 0xa1, 0x80, 0x2e, 0xc7, 0x9e, 0xb5, 0xe9, 0x9e,
	0xd5, 0x01, 0x29, 0xa3, 0x83, 0x16, 0xcd, 0x3f, 0xad, 0x34, 0xc8, 0x84, 0x66, 0x09, 0xee, 0x23,
	0xe8, 0x4c, 0xe1, 0x3e, 0x87, 0x38, 0x28, 0x51, 0xfe, 0xff, 0xe4, 0xdf, 0xc9, 0xdc, 0x4c, 0xad,
	0xe7, 0xce, 0xfd, 0xf0, 0x1d, 0x1f, 0x48, 0x9b, 0x49, 0xc3, 0x9c, 0x2a, 0x39, 0x1c, 0xdd, 0x6e,
	0x59, 0xbf, 0x31, 0xc3, 0xde, 0x4e, 0x7e, 0xce, 0x8e, 0xd3, 0x5d, 0x28, 0xee, 0xba, 0xf3, 0x33,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x9f, 0x15, 0x62, 0xa4, 0x06, 0x00, 0x00,
}
