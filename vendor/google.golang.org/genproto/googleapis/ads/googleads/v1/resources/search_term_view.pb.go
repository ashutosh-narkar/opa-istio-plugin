// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/search_term_view.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A search term view with metrics aggregated by search term at the ad group
// level.
type SearchTermView struct {
	// The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/searchTermViews/{campaign_id}~{ad_group_id}~
	// {URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The search term.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// The ad group the search term served in.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates whether the search term is currently one of your
	// targeted or excluded keywords.
	Status               enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                      `json:"-"`
	XXX_unrecognized     []byte                                                        `json:"-"`
	XXX_sizecache        int32                                                         `json:"-"`
}

func (m *SearchTermView) Reset()         { *m = SearchTermView{} }
func (m *SearchTermView) String() string { return proto.CompactTextString(m) }
func (*SearchTermView) ProtoMessage()    {}
func (*SearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_term_view_733d815a42669c80, []int{0}
}
func (m *SearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermView.Unmarshal(m, b)
}
func (m *SearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermView.Marshal(b, m, deterministic)
}
func (dst *SearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermView.Merge(dst, src)
}
func (m *SearchTermView) XXX_Size() int {
	return xxx_messageInfo_SearchTermView.Size(m)
}
func (m *SearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermView proto.InternalMessageInfo

func (m *SearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *SearchTermView) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *SearchTermView) GetStatus() enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus {
	if m != nil {
		return m.Status
	}
	return enums.SearchTermTargetingStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*SearchTermView)(nil), "google.ads.googleads.v1.resources.SearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/search_term_view.proto", fileDescriptor_search_term_view_733d815a42669c80)
}

var fileDescriptor_search_term_view_733d815a42669c80 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0x72, 0x71, 0xdb, 0x75, 0xeb, 0x83, 0x7a, 0x11, 0xc6, 0x14, 0xbb, 0xc5, 0xe0, 0xd3,
	0x0a, 0xb9, 0x87, 0x96, 0x2d, 0x85, 0xca, 0x10, 0x0c, 0x39, 0x04, 0x23, 0x1b, 0x1d, 0x82, 0x40,
	0xac, 0xad, 0xcd, 0x46, 0x60, 0xed, 0x8a, 0xdd, 0x95, 0x7d, 0xcd, 0x1f, 0xc9, 0x25, 0xc7, 0xfc,
	0x94, 0xfc, 0x94, 0xfc, 0x8a, 0xa0, 0x8f, 0xdd, 0xc4, 0x07, 0x27, 0xb9, 0x3d, 0x69, 0xde, 0x9b,
	0x37, 0x6f, 0x66, 0xc1, 0x1f, 0xca, 0x39, 0xdd, 0x11, 0x0f, 0xa7, 0xd2, 0x6b, 0x60, 0x85, 0xf6,
	0xbe, 0x27, 0x88, 0xe4, 0xa5, 0xd8, 0x12, 0xe9, 0x49, 0x82, 0xc5, 0xf6, 0x3a, 0x51, 0x44, 0xe4,
	0xc9, 0x3e, 0x23, 0x07, 0x58, 0x08, 0xae, 0xb8, 0x33, 0x6e, 0xe8, 0x10, 0xa7, 0x12, 0x1a, 0x25,
	0xdc, 0xfb, 0xd0, 0x28, 0x07, 0xff, 0x4f, 0x35, 0x27, 0xac, 0xcc, 0x8f, 0x1b, 0x2b, 0x2c, 0x28,
	0x51, 0x19, 0xa3, 0x89, 0x54, 0x58, 0x95, 0xb2, 0x31, 0x19, 0x7c, 0x6f, 0x3b, 0xd4, 0x5f, 0x9b,
	0xf2, 0xca, 0x3b, 0x08, 0x5c, 0x14, 0x44, 0xe8, 0xfa, 0x50, 0x3b, 0x14, 0x99, 0x87, 0x19, 0xe3,
	0x0a, 0xab, 0x8c, 0xb3, 0xb6, 0xfa, 0xe3, 0xd6, 0x06, 0xfd, 0x55, 0x6d, 0xb2, 0x26, 0x22, 0x8f,
	0x32, 0x72, 0x70, 0x7e, 0x82, 0xaf, 0x7a, 0xbe, 0x84, 0xe1, 0x9c, 0xb8, 0xd6, 0xc8, 0x9a, 0x7e,
	0x0e, 0xbf, 0xe8, 0x9f, 0x17, 0x38, 0x27, 0xce, 0x3f, 0xd0, 0x7b, 0x31, 0x9b, 0x6b, 0x8f, 0xac,
	0x69, 0x6f, 0x36, 0x6c, 0x53, 0x42, 0x3d, 0x0b, 0x5c, 0x29, 0x91, 0x31, 0x1a, 0xe1, 0x5d, 0x49,
	0x42, 0x20, 0x8d, 0x8f, 0xf3, 0x1b, 0x7c, 0xc2, 0x69, 0x42, 0x05, 0x2f, 0x0b, 0xb7, 0xf3, 0x0e,
	0xed, 0x47, 0x9c, 0x2e, 0x2a, 0xb2, 0xc3, 0x40, 0xb7, 0x49, 0xef, 0x7e, 0x18, 0x59, 0xd3, 0xfe,
	0x2c, 0x82, 0xa7, 0x76, 0x5c, 0x2f, 0x10, 0x3e, 0x67, 0x5b, 0xeb, 0xf5, 0xad, 0x6a, 0xfd, 0x19,
	0x2b, 0xf3, 0xd3, 0xd5, 0xb0, 0x75, 0x99, 0xdf, 0xd8, 0x60, 0xb2, 0xe5, 0x39, 0x7c, 0xf3, 0x92,
	0xf3, 0x6f, 0xc7, 0x6b, 0x5c, 0x56, 0x31, 0x96, 0xd6, 0xe5, 0x79, 0xab, 0xa4, 0x7c, 0x87, 0x19,
	0x85, 0x5c, 0x50, 0x8f, 0x12, 0x56, 0x87, 0xd4, 0x07, 0x2f, 0x32, 0xf9, 0xca, 0xe3, 0xfa, 0x6b,
	0xd0, 0x9d, 0xdd, 0x59, 0x04, 0xc1, 0xbd, 0x3d, 0x5e, 0x34, 0x2d, 0x83, 0x54, 0xc2, 0x06, 0x56,
	0x28, 0xf2, 0x61, 0xa8, 0x99, 0x0f, 0x9a, 0x13, 0x07, 0xa9, 0x8c, 0x0d, 0x27, 0x8e, 0xfc, 0xd8,
	0x70, 0x1e, 0xed, 0x49, 0x53, 0x40, 0x28, 0x48, 0x25, 0x42, 0x86, 0x85, 0x50, 0xe4, 0x23, 0x64,
	0x78, 0x9b, 0x6e, 0x3d, 0xec, 0xaf, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xb3, 0xb8, 0xe9,
	0x08, 0x03, 0x00, 0x00,
}
