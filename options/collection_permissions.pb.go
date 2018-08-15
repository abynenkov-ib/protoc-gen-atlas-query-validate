// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: options/collection_permissions.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	options/collection_permissions.proto

It has these top-level messages:
	CollectionPermissions
	FilterOptions
	FieldPermissions
	MethodOptions
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CollectionPermissions struct {
	DisableSorting   *bool          `protobuf:"varint,1,opt,name=disable_sorting,json=disableSorting" json:"disable_sorting,omitempty"`
	Filters          *FilterOptions `protobuf:"bytes,2,opt,name=filters" json:"filters,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *CollectionPermissions) Reset()         { *m = CollectionPermissions{} }
func (m *CollectionPermissions) String() string { return proto.CompactTextString(m) }
func (*CollectionPermissions) ProtoMessage()    {}
func (*CollectionPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{0}
}

func (m *CollectionPermissions) GetDisableSorting() bool {
	if m != nil && m.DisableSorting != nil {
		return *m.DisableSorting
	}
	return false
}

func (m *CollectionPermissions) GetFilters() *FilterOptions {
	if m != nil {
		return m.Filters
	}
	return nil
}

type FilterOptions struct {
	// Types that are valid to be assigned to Operators:
	//	*FilterOptions_Allow
	//	*FilterOptions_Deny
	Operators        isFilterOptions_Operators `protobuf_oneof:"operators"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *FilterOptions) Reset()         { *m = FilterOptions{} }
func (m *FilterOptions) String() string { return proto.CompactTextString(m) }
func (*FilterOptions) ProtoMessage()    {}
func (*FilterOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{1}
}

type isFilterOptions_Operators interface {
	isFilterOptions_Operators()
}

type FilterOptions_Allow struct {
	Allow string `protobuf:"bytes,1,opt,name=allow,oneof"`
}
type FilterOptions_Deny struct {
	Deny string `protobuf:"bytes,2,opt,name=deny,oneof"`
}

func (*FilterOptions_Allow) isFilterOptions_Operators() {}
func (*FilterOptions_Deny) isFilterOptions_Operators()  {}

func (m *FilterOptions) GetOperators() isFilterOptions_Operators {
	if m != nil {
		return m.Operators
	}
	return nil
}

func (m *FilterOptions) GetAllow() string {
	if x, ok := m.GetOperators().(*FilterOptions_Allow); ok {
		return x.Allow
	}
	return ""
}

func (m *FilterOptions) GetDeny() string {
	if x, ok := m.GetOperators().(*FilterOptions_Deny); ok {
		return x.Deny
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FilterOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FilterOptions_OneofMarshaler, _FilterOptions_OneofUnmarshaler, _FilterOptions_OneofSizer, []interface{}{
		(*FilterOptions_Allow)(nil),
		(*FilterOptions_Deny)(nil),
	}
}

func _FilterOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FilterOptions)
	// operators
	switch x := m.Operators.(type) {
	case *FilterOptions_Allow:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Allow)
	case *FilterOptions_Deny:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Deny)
	case nil:
	default:
		return fmt.Errorf("FilterOptions.Operators has unexpected type %T", x)
	}
	return nil
}

func _FilterOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FilterOptions)
	switch tag {
	case 1: // operators.allow
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operators = &FilterOptions_Allow{x}
		return true, err
	case 2: // operators.deny
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operators = &FilterOptions_Deny{x}
		return true, err
	default:
		return false, nil
	}
}

func _FilterOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FilterOptions)
	// operators
	switch x := m.Operators.(type) {
	case *FilterOptions_Allow:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Allow)))
		n += len(x.Allow)
	case *FilterOptions_Deny:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Deny)))
		n += len(x.Deny)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FieldPermissions struct {
	Name             *string        `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	DisableSorting   *bool          `protobuf:"varint,2,opt,name=disable_sorting,json=disableSorting" json:"disable_sorting,omitempty"`
	Filters          *FilterOptions `protobuf:"bytes,3,opt,name=filters" json:"filters,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FieldPermissions) Reset()         { *m = FieldPermissions{} }
func (m *FieldPermissions) String() string { return proto.CompactTextString(m) }
func (*FieldPermissions) ProtoMessage()    {}
func (*FieldPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{2}
}

func (m *FieldPermissions) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FieldPermissions) GetDisableSorting() bool {
	if m != nil && m.DisableSorting != nil {
		return *m.DisableSorting
	}
	return false
}

func (m *FieldPermissions) GetFilters() *FilterOptions {
	if m != nil {
		return m.Filters
	}
	return nil
}

type MethodOptions struct {
	ForMessage         *string             `protobuf:"bytes,1,opt,name=for_message,json=forMessage" json:"for_message,omitempty"`
	Fields             []*FieldPermissions `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	AllowMissingFields *bool               `protobuf:"varint,3,opt,name=allow_missing_fields,json=allowMissingFields" json:"allow_missing_fields,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorCollectionPermissions, []int{3}
}

func (m *MethodOptions) GetForMessage() string {
	if m != nil && m.ForMessage != nil {
		return *m.ForMessage
	}
	return ""
}

func (m *MethodOptions) GetFields() []*FieldPermissions {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *MethodOptions) GetAllowMissingFields() bool {
	if m != nil && m.AllowMissingFields != nil {
		return *m.AllowMissingFields
	}
	return false
}

var E_Permissions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*CollectionPermissions)(nil),
	Field:         52121,
	Name:          "perm.permissions",
	Tag:           "bytes,52121,opt,name=permissions",
	Filename:      "options/collection_permissions.proto",
}

var E_MethodPermissions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*MethodOptions)(nil),
	Field:         52121,
	Name:          "perm.method_permissions",
	Tag:           "bytes,52121,opt,name=method_permissions,json=methodPermissions",
	Filename:      "options/collection_permissions.proto",
}

func init() {
	proto.RegisterType((*CollectionPermissions)(nil), "perm.CollectionPermissions")
	proto.RegisterType((*FilterOptions)(nil), "perm.FilterOptions")
	proto.RegisterType((*FieldPermissions)(nil), "perm.FieldPermissions")
	proto.RegisterType((*MethodOptions)(nil), "perm.MethodOptions")
	proto.RegisterExtension(E_Permissions)
	proto.RegisterExtension(E_MethodPermissions)
}

func init() {
	proto.RegisterFile("options/collection_permissions.proto", fileDescriptorCollectionPermissions)
}

var fileDescriptorCollectionPermissions = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8b, 0xd4, 0x30,
	0x18, 0xb6, 0x33, 0xe3, 0xc7, 0xbc, 0x65, 0xfd, 0x88, 0xeb, 0x52, 0x14, 0xb5, 0x0c, 0x82, 0xbd,
	0x4c, 0x2b, 0x7b, 0x92, 0xf1, 0xb6, 0xc2, 0x22, 0xc2, 0xa0, 0xc4, 0x9b, 0x07, 0x4b, 0x3f, 0xd2,
	0x6e, 0x20, 0xcd, 0x5b, 0x92, 0x2c, 0xea, 0xc5, 0x1f, 0xe1, 0x45, 0xfc, 0xb7, 0x92, 0xa4, 0x5d,
	0x5b, 0x1c, 0x70, 0x6f, 0xc9, 0xf3, 0xa6, 0xcf, 0xc7, 0xcb, 0x53, 0x78, 0x81, 0xbd, 0xe1, 0x28,
	0x75, 0x56, 0xa1, 0x10, 0xac, 0xb2, 0xe7, 0xbc, 0x67, 0xaa, 0xe3, 0x5a, 0x5b, 0x38, 0xed, 0x15,
	0x1a, 0x24, 0x2b, 0x0b, 0x3d, 0x8e, 0x5b, 0xc4, 0x56, 0xb0, 0xcc, 0x61, 0xe5, 0x65, 0x93, 0xd5,
	0x4c, 0x57, 0x8a, 0xf7, 0x06, 0x95, 0x7f, 0xb7, 0x41, 0x78, 0xf4, 0xf6, 0x8a, 0xe7, 0xe3, 0x5f,
	0x1a, 0xf2, 0x12, 0xee, 0xd5, 0x5c, 0x17, 0xa5, 0x60, 0xb9, 0x46, 0x65, 0xb8, 0x6c, 0xa3, 0x20,
	0x0e, 0x92, 0x3b, 0xf4, 0xee, 0x00, 0x7f, 0xf2, 0x28, 0xd9, 0xc2, 0xed, 0x86, 0x0b, 0xc3, 0x94,
	0x8e, 0x16, 0x71, 0x90, 0x84, 0xa7, 0x0f, 0x53, 0xab, 0x9d, 0x9e, 0x3b, 0xf0, 0x83, 0x37, 0x4b,
	0xc7, 0x37, 0x9b, 0xf7, 0x70, 0x34, 0x9b, 0x90, 0x13, 0xb8, 0x59, 0x08, 0x81, 0x5f, 0x1d, 0xfd,
	0xfa, 0xdd, 0x0d, 0xea, 0xaf, 0xe4, 0x18, 0x56, 0x35, 0x93, 0xdf, 0x1d, 0xa9, 0x85, 0xdd, 0xed,
	0x2c, 0x84, 0x35, 0xf6, 0x4c, 0x15, 0x06, 0x95, 0xde, 0xfc, 0x80, 0xfb, 0xe7, 0x9c, 0x89, 0x7a,
	0xea, 0x9b, 0xc0, 0x4a, 0x16, 0x1d, 0x8b, 0x82, 0x78, 0x91, 0xac, 0xa9, 0x3b, 0x1f, 0xca, 0xb2,
	0xf8, 0x5f, 0x96, 0xe5, 0x35, 0xb2, 0xfc, 0x0c, 0xe0, 0x68, 0xcf, 0xcc, 0x05, 0xd6, 0x63, 0x98,
	0xe7, 0x10, 0x36, 0xa8, 0xf2, 0x8e, 0x69, 0x5d, 0xb4, 0xcc, 0x47, 0xa2, 0xd0, 0xa0, 0xda, 0x7b,
	0x84, 0xa4, 0x70, 0xab, 0xb1, 0x96, 0xed, 0xb2, 0x96, 0x49, 0x78, 0x7a, 0x32, 0x0a, 0xcc, 0x63,
	0xd0, 0xe1, 0x15, 0x79, 0x05, 0xc7, 0x6e, 0x1d, 0xb9, 0x9b, 0xc8, 0x36, 0x1f, 0xbe, 0x5e, 0x3a,
	0xff, 0xc4, 0xcd, 0xf6, 0x7e, 0xe4, 0x38, 0xf4, 0xee, 0x0b, 0x84, 0x93, 0x3a, 0x90, 0xa7, 0xa9,
	0xef, 0x40, 0x3a, 0x76, 0xc0, 0x6b, 0x0d, 0x86, 0xa3, 0xdf, 0xbf, 0x7c, 0xd0, 0x27, 0xde, 0xc7,
	0xc1, 0x2e, 0xd0, 0x29, 0xe1, 0xae, 0x06, 0xd2, 0xb9, 0xcc, 0xd3, 0xd6, 0x91, 0x67, 0xff, 0xc8,
	0xcc, 0x16, 0x73, 0xa5, 0x33, 0x2c, 0x74, 0x36, 0xa4, 0x0f, 0x3c, 0xe1, 0x44, 0xf2, 0x6c, 0xf7,
	0xf9, 0x75, 0xcb, 0xcd, 0xc5, 0x65, 0x99, 0x56, 0xd8, 0x65, 0x5c, 0x36, 0x58, 0x0a, 0xfc, 0x86,
	0x3d, 0x93, 0xbe, 0xcc, 0xd5, 0xb6, 0x65, 0x72, 0x6b, 0xb9, 0xb2, 0xf1, 0x7f, 0x78, 0x33, 0x1c,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x49, 0x40, 0xe8, 0x22, 0x03, 0x00, 0x00,
}