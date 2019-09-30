// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/header_to_metadata/v3alpha/header_to_metadata.proto

package envoy_config_filter_http_header_to_metadata_v3alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Config_ValueType int32

const (
	Config_STRING Config_ValueType = 0
	Config_NUMBER Config_ValueType = 1
	// The value is a serialized `protobuf.Value
	// <https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/struct.proto#L62>`_.
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
	return fileDescriptor_592f77e24734e10b, []int{0, 0}
}

// ValueEncode defines the encoding algorithm.
type Config_ValueEncode int32

const (
	// The value is not encoded.
	Config_NONE Config_ValueEncode = 0
	// The value is encoded in `Base64 <https://tools.ietf.org/html/rfc4648#section-4>`_.
	// Note: this is mostly used for STRING and PROTOBUF_VALUE to escape the
	// non-ASCII characters in the header.
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
	return fileDescriptor_592f77e24734e10b, []int{0, 1}
}

type Config struct {
	// The list of rules to apply to requests.
	RequestRules []*Config_Rule `protobuf:"bytes,1,rep,name=request_rules,json=requestRules,proto3" json:"request_rules,omitempty"`
	// The list of rules to apply to responses.
	ResponseRules        []*Config_Rule `protobuf:"bytes,2,rep,name=response_rules,json=responseRules,proto3" json:"response_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_592f77e24734e10b, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
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
	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The value to pair with the given key.
	//
	// When used for a `on_header_present` case, if value is non-empty it'll be used
	// instead of the header value. If both are empty, no metadata is added.
	//
	// When used for a `on_header_missing` case, a non-empty value must be provided
	// otherwise no metadata is added.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// The value's type — defaults to string.
	Type Config_ValueType `protobuf:"varint,4,opt,name=type,proto3,enum=envoy.config.filter.http.header_to_metadata.v3alpha.Config_ValueType" json:"type,omitempty"`
	// How is the value encoded, default is NONE (not encoded).
	// The value will be decoded accordingly before storing to metadata.
	Encode               Config_ValueEncode `protobuf:"varint,5,opt,name=encode,proto3,enum=envoy.config.filter.http.header_to_metadata.v3alpha.Config_ValueEncode" json:"encode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config_KeyValuePair) Reset()         { *m = Config_KeyValuePair{} }
func (m *Config_KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*Config_KeyValuePair) ProtoMessage()    {}
func (*Config_KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_592f77e24734e10b, []int{0, 0}
}
func (m *Config_KeyValuePair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config_KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config_KeyValuePair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config_KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_KeyValuePair.Merge(m, src)
}
func (m *Config_KeyValuePair) XXX_Size() int {
	return m.Size()
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

// A Rule defines what metadata to apply when a header is present or missing.
type Config_Rule struct {
	// The header that triggers this rule — required.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// If the header is present, apply this metadata KeyValuePair.
	//
	// If the value in the KeyValuePair is non-empty, it'll be used instead
	// of the header value.
	OnHeaderPresent *Config_KeyValuePair `protobuf:"bytes,2,opt,name=on_header_present,json=onHeaderPresent,proto3" json:"on_header_present,omitempty"`
	// If the header is not present, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set, since it'll be used in lieu
	// of the missing header value.
	OnHeaderMissing *Config_KeyValuePair `protobuf:"bytes,3,opt,name=on_header_missing,json=onHeaderMissing,proto3" json:"on_header_missing,omitempty"`
	// Whether or not to remove the header after a rule is applied.
	//
	// This prevents headers from leaking.
	Remove               bool     `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config_Rule) Reset()         { *m = Config_Rule{} }
func (m *Config_Rule) String() string { return proto.CompactTextString(m) }
func (*Config_Rule) ProtoMessage()    {}
func (*Config_Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_592f77e24734e10b, []int{0, 1}
}
func (m *Config_Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config_Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config_Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Rule.Merge(m, src)
}
func (m *Config_Rule) XXX_Size() int {
	return m.Size()
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
	proto.RegisterEnum("envoy.config.filter.http.header_to_metadata.v3alpha.Config_ValueType", Config_ValueType_name, Config_ValueType_value)
	proto.RegisterEnum("envoy.config.filter.http.header_to_metadata.v3alpha.Config_ValueEncode", Config_ValueEncode_name, Config_ValueEncode_value)
	proto.RegisterType((*Config)(nil), "envoy.config.filter.http.header_to_metadata.v3alpha.Config")
	proto.RegisterType((*Config_KeyValuePair)(nil), "envoy.config.filter.http.header_to_metadata.v3alpha.Config.KeyValuePair")
	proto.RegisterType((*Config_Rule)(nil), "envoy.config.filter.http.header_to_metadata.v3alpha.Config.Rule")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/header_to_metadata/v3alpha/header_to_metadata.proto", fileDescriptor_592f77e24734e10b)
}

var fileDescriptor_592f77e24734e10b = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x19, 0x27, 0x35, 0xcd, 0xa4, 0x0d, 0xe9, 0x88, 0x1f, 0x2b, 0x48, 0x51, 0x08, 0x9b,
	0x6c, 0xb0, 0xa5, 0x04, 0xc1, 0x96, 0x18, 0x99, 0x16, 0xd1, 0xfc, 0x68, 0x9a, 0x54, 0x62, 0x65,
	0x0d, 0xc9, 0x6d, 0x62, 0xe1, 0x78, 0xcc, 0x78, 0x12, 0xe1, 0x07, 0xe1, 0x2d, 0xe0, 0x05, 0x58,
	0xb1, 0x64, 0xc9, 0x23, 0xa0, 0xec, 0x78, 0x0b, 0xe4, 0x19, 0x1b, 0x05, 0xa9, 0xab, 0x06, 0x76,
	0x73, 0x7d, 0x74, 0xbf, 0x33, 0xf7, 0x5c, 0x79, 0xf0, 0x39, 0x44, 0x1b, 0x9e, 0x3a, 0x33, 0x1e,
	0x5d, 0x05, 0x0b, 0xe7, 0x2a, 0x08, 0x25, 0x08, 0x67, 0x29, 0x65, 0xec, 0x2c, 0x81, 0xcd, 0x41,
	0xf8, 0x92, 0xfb, 0x2b, 0x90, 0x6c, 0xce, 0x24, 0x73, 0x36, 0x3d, 0x16, 0xc6, 0x4b, 0x76, 0x8d,
	0x64, 0xc7, 0x82, 0x4b, 0x4e, 0x7a, 0x8a, 0x66, 0x6b, 0x9a, 0xad, 0x69, 0x76, 0x46, 0xb3, 0xaf,
	0x69, 0xc9, 0x69, 0x8d, 0x07, 0x1b, 0x16, 0x06, 0x73, 0x26, 0xc1, 0x29, 0x0e, 0x9a, 0xd6, 0xfe,
	0x74, 0x1b, 0x9b, 0x2f, 0x15, 0x8a, 0x00, 0x3e, 0x16, 0xf0, 0x61, 0x0d, 0x89, 0xf4, 0xc5, 0x3a,
	0x84, 0xc4, 0x42, 0xad, 0x52, 0xa7, 0xda, 0x7d, 0x61, 0xdf, 0xc0, 0xd0, 0xd6, 0x4c, 0x9b, 0xae,
	0x43, 0xa0, 0x47, 0x39, 0x36, 0x2b, 0x12, 0xb2, 0xc0, 0x35, 0x01, 0x49, 0xcc, 0xa3, 0x04, 0x72,
	0x1f, 0xe3, 0x1f, 0xf9, 0x1c, 0x17, 0x5c, 0x65, 0xd4, 0xf8, 0x6c, 0xe0, 0xa3, 0x37, 0x90, 0x5e,
	0xb2, 0x70, 0x0d, 0x63, 0x16, 0x08, 0xf2, 0x04, 0x93, 0xa2, 0xdf, 0x8f, 0xd8, 0x0a, 0x92, 0x98,
	0xcd, 0xc0, 0x42, 0x2d, 0xd4, 0xa9, 0xd0, 0x93, 0x42, 0x19, 0x16, 0x02, 0x79, 0x88, 0x4b, 0xef,
	0x21, 0xb5, 0x8c, 0x4c, 0x77, 0x2b, 0x5f, 0x7f, 0x7d, 0x2b, 0x95, 0x85, 0xd1, 0x42, 0x34, 0xfb,
	0x4a, 0xee, 0xe2, 0x83, 0x4d, 0x06, 0xb6, 0x4a, 0xaa, 0x5d, 0x17, 0xe4, 0x2d, 0x2e, 0xcb, 0x34,
	0x06, 0xab, 0xdc, 0x42, 0x9d, 0x5a, 0xd7, 0xdb, 0x67, 0x22, 0x75, 0xed, 0x49, 0x1a, 0x03, 0x55,
	0x48, 0xe2, 0x63, 0x13, 0xa2, 0x19, 0x9f, 0x83, 0x75, 0xa0, 0xe0, 0xa7, 0x7b, 0xc3, 0x3d, 0x85,
	0xa3, 0x39, 0xb6, 0xf1, 0xc5, 0xc0, 0xe5, 0x2c, 0x38, 0xf2, 0x08, 0x9b, 0x9a, 0xa0, 0xa3, 0xd9,
	0x1d, 0x3d, 0x17, 0x88, 0xc4, 0x27, 0x3c, 0xf2, 0x73, 0x9f, 0x58, 0x40, 0x02, 0x91, 0x54, 0x41,
	0x55, 0xbb, 0x67, 0xfb, 0xdc, 0x6b, 0x77, 0x5d, 0xf4, 0x0e, 0x8f, 0xce, 0x54, 0xc7, 0x58, 0x1b,
	0xfc, 0xed, 0xba, 0x0a, 0x92, 0x24, 0x88, 0x16, 0x2a, 0xff, 0xff, 0xe2, 0x3a, 0xd0, 0x06, 0xe4,
	0x3e, 0x36, 0x05, 0xac, 0xf8, 0x46, 0x6f, 0xf5, 0x90, 0xe6, 0x55, 0xfb, 0x39, 0xae, 0xfc, 0xd9,
	0x11, 0xc1, 0xd8, 0xbc, 0x98, 0xd0, 0xd7, 0xc3, 0xd3, 0xfa, 0xad, 0xec, 0x3c, 0x9c, 0x0e, 0x5c,
	0x8f, 0xd6, 0x11, 0x21, 0xb8, 0x36, 0xa6, 0xa3, 0xc9, 0xc8, 0x9d, 0xbe, 0xf2, 0x2f, 0xfb, 0xe7,
	0x53, 0xaf, 0x6e, 0xb4, 0x1f, 0xe3, 0xea, 0x4e, 0xfe, 0xe4, 0x10, 0x97, 0x87, 0xa3, 0xa1, 0xa7,
	0x1b, 0xdd, 0xfe, 0x85, 0xf7, 0xec, 0x69, 0x1d, 0xb9, 0xcb, 0xef, 0xdb, 0x26, 0xfa, 0xb1, 0x6d,
	0xa2, 0x9f, 0xdb, 0x26, 0xc2, 0xfd, 0x80, 0xeb, 0x01, 0x63, 0xc1, 0x3f, 0xa6, 0x37, 0x99, 0xd5,
	0xbd, 0xa7, 0xa7, 0x9a, 0xf0, 0x41, 0xae, 0x8c, 0xb3, 0xff, 0x7f, 0x8c, 0xde, 0x99, 0xea, 0x21,
	0xe8, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x08, 0x42, 0x5f, 0xa8, 0xa6, 0x04, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResponseRules) > 0 {
		for iNdEx := len(m.ResponseRules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResponseRules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHeaderToMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RequestRules) > 0 {
		for iNdEx := len(m.RequestRules) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestRules[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHeaderToMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Config_KeyValuePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config_KeyValuePair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config_KeyValuePair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Encode != 0 {
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(m.Encode))
		i--
		dAtA[i] = 0x28
	}
	if m.Type != 0 {
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MetadataNamespace) > 0 {
		i -= len(m.MetadataNamespace)
		copy(dAtA[i:], m.MetadataNamespace)
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.MetadataNamespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Config_Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config_Rule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config_Rule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Remove {
		i--
		if m.Remove {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.OnHeaderMissing != nil {
		{
			size, err := m.OnHeaderMissing.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHeaderToMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.OnHeaderPresent != nil {
		{
			size, err := m.OnHeaderPresent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHeaderToMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Header) > 0 {
		i -= len(m.Header)
		copy(dAtA[i:], m.Header)
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Header)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeaderToMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeaderToMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestRules) > 0 {
		for _, e := range m.RequestRules {
			l = e.Size()
			n += 1 + l + sovHeaderToMetadata(uint64(l))
		}
	}
	if len(m.ResponseRules) > 0 {
		for _, e := range m.ResponseRules {
			l = e.Size()
			n += 1 + l + sovHeaderToMetadata(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Config_KeyValuePair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetadataNamespace)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovHeaderToMetadata(uint64(m.Type))
	}
	if m.Encode != 0 {
		n += 1 + sovHeaderToMetadata(uint64(m.Encode))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Config_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.OnHeaderPresent != nil {
		l = m.OnHeaderPresent.Size()
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.OnHeaderMissing != nil {
		l = m.OnHeaderMissing.Size()
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.Remove {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHeaderToMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeaderToMetadata(x uint64) (n int) {
	return sovHeaderToMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestRules = append(m.RequestRules, &Config_Rule{})
			if err := m.RequestRules[len(m.RequestRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseRules = append(m.ResponseRules, &Config_Rule{})
			if err := m.ResponseRules[len(m.ResponseRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Config_KeyValuePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyValuePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValuePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Config_ValueType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encode", wireType)
			}
			m.Encode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encode |= Config_ValueEncode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Config_Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Header = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnHeaderPresent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnHeaderPresent == nil {
				m.OnHeaderPresent = &Config_KeyValuePair{}
			}
			if err := m.OnHeaderPresent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnHeaderMissing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnHeaderMissing == nil {
				m.OnHeaderMissing = &Config_KeyValuePair{}
			}
			if err := m.OnHeaderMissing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remove", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Remove = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaderToMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHeaderToMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaderToMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHeaderToMetadata
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHeaderToMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeaderToMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHeaderToMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHeaderToMetadata
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHeaderToMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaderToMetadata   = fmt.Errorf("proto: integer overflow")
)