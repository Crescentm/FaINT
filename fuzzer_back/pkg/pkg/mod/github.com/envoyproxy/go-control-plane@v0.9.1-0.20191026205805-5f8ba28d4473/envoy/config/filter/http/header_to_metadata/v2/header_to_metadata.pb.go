// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto

package envoy_config_filter_http_header_to_metadata_v2

import (
	fmt "fmt"
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

type Config_ValueType int32

const (
	Config_STRING         Config_ValueType = 0
	Config_NUMBER         Config_ValueType = 1
	Config_PROTOBUF_VALUE Config_ValueType = 2
)

var Config_ValueType_name = map[int32]string{
	0: "STRING",
	1: "NUMBER",
	2: "PROTOBUF_VALUE",
}

var Config_ValueType_value = map[string]int32{
	"STRING":         0,
	"NUMBER":         1,
	"PROTOBUF_VALUE": 2,
}

func (x Config_ValueType) String() string {
	return proto.EnumName(Config_ValueType_name, int32(x))
}

func (Config_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5891b90c2f7c1ebe, []int{0, 0}
}

type Config_ValueEncode int32

const (
	Config_NONE   Config_ValueEncode = 0
	Config_BASE64 Config_ValueEncode = 1
)

var Config_ValueEncode_name = map[int32]string{
	0: "NONE",
	1: "BASE64",
}

var Config_ValueEncode_value = map[string]int32{
	"NONE":   0,
	"BASE64": 1,
}

func (x Config_ValueEncode) String() string {
	return proto.EnumName(Config_ValueEncode_name, int32(x))
}

func (Config_ValueEncode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5891b90c2f7c1ebe, []int{0, 1}
}

type Config struct {
	RequestRules         []*Config_Rule `protobuf:"bytes,1,rep,name=request_rules,json=requestRules,proto3" json:"request_rules,omitempty"`
	ResponseRules        []*Config_Rule `protobuf:"bytes,2,rep,name=response_rules,json=responseRules,proto3" json:"response_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_5891b90c2f7c1ebe, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetRequestRules() []*Config_Rule {
	if m != nil {
		return m.RequestRules
	}
	return nil
}

func (m *Config) GetResponseRules() []*Config_Rule {
	if m != nil {
		return m.ResponseRules
	}
	return nil
}

type Config_KeyValuePair struct {
	MetadataNamespace    string             `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	Key                  string             `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string             `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Type                 Config_ValueType   `protobuf:"varint,4,opt,name=type,proto3,enum=envoy.config.filter.http.header_to_metadata.v2.Config_ValueType" json:"type,omitempty"`
	Encode               Config_ValueEncode `protobuf:"varint,5,opt,name=encode,proto3,enum=envoy.config.filter.http.header_to_metadata.v2.Config_ValueEncode" json:"encode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config_KeyValuePair) Reset()         { *m = Config_KeyValuePair{} }
func (m *Config_KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*Config_KeyValuePair) ProtoMessage()    {}
func (*Config_KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5891b90c2f7c1ebe, []int{0, 0}
}

func (m *Config_KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_KeyValuePair.Unmarshal(m, b)
}
func (m *Config_KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *Config_KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_KeyValuePair.Merge(m, src)
}
func (m *Config_KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_Config_KeyValuePair.Size(m)
}
func (m *Config_KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_Config_KeyValuePair proto.InternalMessageInfo

func (m *Config_KeyValuePair) GetMetadataNamespace() string {
	if m != nil {
		return m.MetadataNamespace
	}
	return ""
}

func (m *Config_KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Config_KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Config_KeyValuePair) GetType() Config_ValueType {
	if m != nil {
		return m.Type
	}
	return Config_STRING
}

func (m *Config_KeyValuePair) GetEncode() Config_ValueEncode {
	if m != nil {
		return m.Encode
	}
	return Config_NONE
}

type Config_Rule struct {
	Header               string               `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	OnHeaderPresent      *Config_KeyValuePair `protobuf:"bytes,2,opt,name=on_header_present,json=onHeaderPresent,proto3" json:"on_header_present,omitempty"`
	OnHeaderMissing      *Config_KeyValuePair `protobuf:"bytes,3,opt,name=on_header_missing,json=onHeaderMissing,proto3" json:"on_header_missing,omitempty"`
	Remove               bool                 `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Config_Rule) Reset()         { *m = Config_Rule{} }
func (m *Config_Rule) String() string { return proto.CompactTextString(m) }
func (*Config_Rule) ProtoMessage()    {}
func (*Config_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_5891b90c2f7c1ebe, []int{0, 1}
}

func (m *Config_Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Rule.Unmarshal(m, b)
}
func (m *Config_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Rule.Marshal(b, m, deterministic)
}
func (m *Config_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Rule.Merge(m, src)
}
func (m *Config_Rule) XXX_Size() int {
	return xxx_messageInfo_Config_Rule.Size(m)
}
func (m *Config_Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Rule proto.InternalMessageInfo

func (m *Config_Rule) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Config_Rule) GetOnHeaderPresent() *Config_KeyValuePair {
	if m != nil {
		return m.OnHeaderPresent
	}
	return nil
}

func (m *Config_Rule) GetOnHeaderMissing() *Config_KeyValuePair {
	if m != nil {
		return m.OnHeaderMissing
	}
	return nil
}

func (m *Config_Rule) GetRemove() bool {
	if m != nil {
		return m.Remove
	}
	return false
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.header_to_metadata.v2.Config_ValueType", Config_ValueType_name, Config_ValueType_value)
	proto.RegisterEnum("envoy.config.filter.http.header_to_metadata.v2.Config_ValueEncode", Config_ValueEncode_name, Config_ValueEncode_value)
	proto.RegisterType((*Config)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config")
	proto.RegisterType((*Config_KeyValuePair)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config.KeyValuePair")
	proto.RegisterType((*Config_Rule)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config.Rule")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto", fileDescriptor_5891b90c2f7c1ebe)
}

var fileDescriptor_5891b90c2f7c1ebe = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0x9b, 0x9a, 0x76, 0xd2, 0x86, 0x74, 0xc5, 0x1f, 0x93, 0x0b, 0x51, 0xb8, 0xe4,
	0x82, 0x2d, 0x05, 0x04, 0x07, 0x38, 0x50, 0x57, 0xa6, 0x20, 0x88, 0x63, 0x6d, 0x9d, 0x22, 0x71,
	0x31, 0xdb, 0x64, 0xda, 0x5a, 0x38, 0x5e, 0xb3, 0xde, 0x58, 0xf8, 0x6d, 0xe0, 0x19, 0xb9, 0x71,
	0x42, 0xde, 0xb5, 0xa1, 0x48, 0xbd, 0xb4, 0xe5, 0xb6, 0x33, 0xb3, 0xf3, 0xfb, 0x76, 0xbe, 0xd1,
	0xc2, 0x21, 0x66, 0x25, 0xaf, 0xdc, 0x05, 0xcf, 0x4e, 0x93, 0x33, 0xf7, 0x34, 0x49, 0x25, 0x0a,
	0xf7, 0x5c, 0xca, 0xdc, 0x3d, 0x47, 0xb6, 0x44, 0x11, 0x4b, 0x1e, 0xaf, 0x50, 0xb2, 0x25, 0x93,
	0xcc, 0x2d, 0x27, 0x97, 0x64, 0x9d, 0x5c, 0x70, 0xc9, 0x89, 0xa3, 0x40, 0x8e, 0x06, 0x39, 0x1a,
	0xe4, 0xd4, 0x20, 0xe7, 0x92, 0x96, 0x72, 0x32, 0x78, 0x50, 0xb2, 0x34, 0x59, 0x32, 0x89, 0x6e,
	0x7b, 0xd0, 0xa0, 0xd1, 0x4f, 0x0b, 0xac, 0x03, 0x45, 0x21, 0x9f, 0x61, 0x57, 0xe0, 0xd7, 0x35,
	0x16, 0x32, 0x16, 0xeb, 0x14, 0x0b, 0xdb, 0x18, 0x6e, 0x8c, 0xbb, 0x93, 0x97, 0x57, 0xd4, 0x72,
	0x34, 0xce, 0xa1, 0xeb, 0x14, 0xe9, 0x4e, 0x43, 0xac, 0x83, 0x82, 0x9c, 0x40, 0x4f, 0x60, 0x91,
	0xf3, 0xac, 0xc0, 0x46, 0xc2, 0xbc, 0xb9, 0xc4, 0x6e, 0x8b, 0x54, 0x1a, 0x83, 0xef, 0x26, 0xec,
	0xbc, 0xc7, 0xea, 0x98, 0xa5, 0x6b, 0x0c, 0x59, 0x22, 0xc8, 0x13, 0x20, 0x6d, 0x6b, 0x9c, 0xb1,
	0x15, 0x16, 0x39, 0x5b, 0xa0, 0x6d, 0x0c, 0x8d, 0xf1, 0x36, 0xdd, 0x6b, 0x2b, 0x41, 0x5b, 0x20,
	0x0f, 0x61, 0xe3, 0x0b, 0x56, 0xb6, 0x59, 0xd7, 0xbd, 0xdb, 0xbf, 0xbc, 0x8e, 0x30, 0x87, 0x06,
	0xad, 0x73, 0xe4, 0x2e, 0x6c, 0x96, 0x35, 0xd6, 0xde, 0x50, 0xcd, 0x3a, 0x20, 0x11, 0x74, 0x64,
	0x95, 0xa3, 0xdd, 0x19, 0x1a, 0xe3, 0xde, 0xe4, 0xf5, 0x35, 0x47, 0x51, 0xef, 0x8d, 0xaa, 0x1c,
	0xa9, 0xa2, 0x91, 0x4f, 0x60, 0x61, 0xb6, 0xe0, 0x4b, 0xb4, 0x37, 0x15, 0xd7, 0xbb, 0x09, 0xd7,
	0x57, 0x24, 0xda, 0x10, 0x07, 0x3f, 0x4c, 0xe8, 0xd4, 0x66, 0x91, 0x47, 0x60, 0xe9, 0x66, 0x6d,
	0xc7, 0xdf, 0x71, 0x9b, 0x34, 0xe1, 0xb0, 0xc7, 0xb3, 0xb8, 0x11, 0xc8, 0x05, 0x16, 0x98, 0x49,
	0x65, 0x4d, 0x77, 0x72, 0x70, 0xcd, 0x07, 0x5d, 0xdc, 0x0d, 0xbd, 0xc3, 0xb3, 0xb7, 0xea, 0x72,
	0xa8, 0xd9, 0xff, 0x0a, 0xae, 0x92, 0xa2, 0x48, 0xb2, 0x33, 0x65, 0xf7, 0xff, 0x16, 0x9c, 0x6a,
	0x36, 0xb9, 0x0f, 0x96, 0xc0, 0x15, 0x2f, 0xf5, 0xfe, 0xb6, 0x68, 0x13, 0x8d, 0x5e, 0xc0, 0xf6,
	0x9f, 0x95, 0x10, 0x00, 0xeb, 0x28, 0xa2, 0xef, 0x82, 0xc3, 0xfe, 0xad, 0xfa, 0x1c, 0xcc, 0xa7,
	0x9e, 0x4f, 0xfb, 0x06, 0x21, 0xd0, 0x0b, 0xe9, 0x2c, 0x9a, 0x79, 0xf3, 0x37, 0xf1, 0xf1, 0xfe,
	0x87, 0xb9, 0xdf, 0x37, 0x47, 0x8f, 0xa1, 0x7b, 0xc1, 0x73, 0xb2, 0x05, 0x9d, 0x60, 0x16, 0xf8,
	0xba, 0xd1, 0xdb, 0x3f, 0xf2, 0x9f, 0x3f, 0xeb, 0x1b, 0xde, 0x47, 0x78, 0x95, 0x70, 0x3d, 0x4f,
	0x2e, 0xf8, 0xb7, 0xea, 0x8a, 0xa3, 0x79, 0xf7, 0xf4, 0x10, 0x11, 0x9f, 0x36, 0xc9, 0xb0, 0xfe,
	0xcc, 0xa1, 0x71, 0x62, 0xa9, 0x5f, 0xfd, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x72,
	0x8e, 0x2f, 0x69, 0x04, 0x00, 0x00,
}
