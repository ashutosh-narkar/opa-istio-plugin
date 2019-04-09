// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta1/location.proto

package admin // import "google.golang.org/genproto/googleapis/firestore/admin/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The metadata message for [google.cloud.location.Location.metadata][google.cloud.location.Location.metadata].
type LocationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationMetadata) Reset()         { *m = LocationMetadata{} }
func (m *LocationMetadata) String() string { return proto.CompactTextString(m) }
func (*LocationMetadata) ProtoMessage()    {}
func (*LocationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_location_ccd18b97671f76b4, []int{0}
}
func (m *LocationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationMetadata.Unmarshal(m, b)
}
func (m *LocationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationMetadata.Marshal(b, m, deterministic)
}
func (dst *LocationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationMetadata.Merge(dst, src)
}
func (m *LocationMetadata) XXX_Size() int {
	return xxx_messageInfo_LocationMetadata.Size(m)
}
func (m *LocationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_LocationMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LocationMetadata)(nil), "google.firestore.admin.v1beta1.LocationMetadata")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta1/location.proto", fileDescriptor_location_ccd18b97671f76b4)
}

var fileDescriptor_location_ccd18b97671f76b4 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4b, 0xc0, 0x30,
	0x10, 0x85, 0x69, 0x11, 0x87, 0x82, 0x20, 0x9d, 0xa4, 0x88, 0x43, 0x71, 0x35, 0x47, 0x71, 0x74,
	0xb2, 0x85, 0x76, 0x51, 0x28, 0x08, 0x0e, 0x6e, 0xd7, 0x36, 0x86, 0x40, 0x9a, 0x0b, 0xe9, 0x29,
	0xf8, 0x77, 0x9c, 0xc4, 0x5f, 0x29, 0x4d, 0x62, 0x37, 0x1d, 0xc3, 0x7b, 0xdf, 0xf7, 0x2e, 0xc5,
	0x8d, 0x22, 0x52, 0x46, 0xc2, 0xab, 0xf6, 0x72, 0x63, 0xf2, 0x12, 0x70, 0x59, 0xb5, 0x85, 0xf7,
	0x66, 0x92, 0x8c, 0x0d, 0x18, 0x9a, 0x91, 0x35, 0x59, 0xe1, 0x3c, 0x31, 0x95, 0x57, 0xb1, 0x2e,
	0x8e, 0xba, 0x08, 0x75, 0x91, 0xea, 0xd5, 0x45, 0xd2, 0xf1, 0x87, 0x93, 0x60, 0x90, 0x8d, 0x55,
	0x91, 0xac, 0x2e, 0x53, 0x82, 0x4e, 0x03, 0x5a, 0x4b, 0x1c, 0xb4, 0x5b, 0x4c, 0xeb, 0xb2, 0x38,
	0x7f, 0x48, 0x4b, 0x8f, 0x92, 0x71, 0x41, 0xc6, 0xf6, 0x2b, 0x2b, 0xea, 0x99, 0x56, 0xf1, 0xff,
	0x64, 0x7b, 0xf6, 0x0b, 0x8e, 0xbb, 0x69, 0xcc, 0x5e, 0xba, 0x04, 0x28, 0x32, 0x68, 0x95, 0x20,
	0xaf, 0x40, 0x49, 0x1b, 0x76, 0x20, 0x46, 0xe8, 0xf4, 0xf6, 0xd7, 0x8f, 0xef, 0xc2, 0xeb, 0x33,
	0x3f, 0x19, 0xba, 0xfe, 0xe9, 0x3b, 0xbf, 0x1e, 0xa2, 0xac, 0x33, 0xf4, 0xb6, 0x88, 0xfe, 0xb8,
	0xe1, 0x3e, 0xdc, 0xf0, 0xdc, 0xb4, 0x3b, 0x33, 0x9d, 0x06, 0xfb, 0xed, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x76, 0x88, 0x2e, 0x88, 0x4e, 0x01, 0x00, 0x00,
}
