// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auction/category/category/v1/category.proto

package categoryv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Категория.
type Category struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ModifyTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	DeleteTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a884d8571fba727, []int{0}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Category) GetModifyTime() *timestamp.Timestamp {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *Category) GetDeleteTime() *timestamp.Timestamp {
	if m != nil {
		return m.DeleteTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Category)(nil), "auction.category.category.v1.Category")
}

func init() {
	proto.RegisterFile("auction/category/category/v1/category.proto", fileDescriptor_5a884d8571fba727)
}

var fileDescriptor_5a884d8571fba727 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0x2c, 0x4d, 0x2e,
	0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0x44, 0x30, 0xca, 0x0c,
	0xe1, 0x6c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x19, 0xa8, 0x62, 0x3d, 0xb8, 0x38, 0x9c,
	0x51, 0x66, 0x28, 0x25, 0x9f, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x0f, 0x56, 0x9b, 0x54, 0x9a,
	0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b, 0x00, 0xd1, 0xae, 0xf4, 0x94, 0x91,
	0x8b, 0xc3, 0x19, 0xaa, 0x41, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x35, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0xb2, 0xe6, 0xe2, 0x4e, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x8d,
	0x07, 0x19, 0x25, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0xb1, 0x47, 0x0f, 0x66,
	0x8f, 0x5e, 0x08, 0xcc, 0x9e, 0x20, 0x2e, 0x88, 0x72, 0x90, 0x00, 0x48, 0x73, 0x6e, 0x7e, 0x4a,
	0x66, 0x5a, 0x25, 0x44, 0x33, 0x0b, 0x61, 0xcd, 0x10, 0xe5, 0x30, 0xcd, 0x29, 0xa9, 0x39, 0xa9,
	0x30, 0x9b, 0x59, 0x09, 0x6b, 0x86, 0x28, 0x07, 0x09, 0x38, 0x35, 0x33, 0x72, 0x29, 0x24, 0xe7,
	0xe7, 0xea, 0xe1, 0x0b, 0x2d, 0x27, 0x5e, 0x58, 0x48, 0x04, 0x80, 0x0c, 0x0b, 0x60, 0x8c, 0xe2,
	0x82, 0xc9, 0x96, 0x19, 0x2e, 0x62, 0x62, 0x76, 0x74, 0x76, 0x5e, 0xc5, 0x24, 0xe3, 0x08, 0x35,
	0x01, 0xa6, 0x16, 0xc1, 0x08, 0x33, 0x3c, 0x05, 0x97, 0x8e, 0x81, 0x89, 0x22, 0x18, 0x61, 0x86,
	0x49, 0x6c, 0x60, 0x57, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xb9, 0x2b, 0x57, 0xe2,
	0x01, 0x00, 0x00,
}
