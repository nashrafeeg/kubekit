// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apply.proto

package v1

import (
	fmt "fmt"
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

type ApplyAction int32

const (
	ApplyAction_ALL       ApplyAction = 0
	ApplyAction_PROVISION ApplyAction = 1
	ApplyAction_CONFIGURE ApplyAction = 2
)

var ApplyAction_name = map[int32]string{
	0: "ALL",
	1: "PROVISION",
	2: "CONFIGURE",
}

var ApplyAction_value = map[string]int32{
	"ALL":       0,
	"PROVISION": 1,
	"CONFIGURE": 2,
}

func (x ApplyAction) String() string {
	return proto.EnumName(ApplyAction_name, int32(x))
}

func (ApplyAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_993661bab0ce9d1e, []int{0}
}

type ApplyRequest struct {
	Api                  string            `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ClusterName          string            `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Action               ApplyAction       `protobuf:"varint,3,opt,name=action,proto3,enum=kubekit.v1.ApplyAction" json:"action,omitempty"`
	PackageUrl           string            `protobuf:"bytes,4,opt,name=package_url,json=packageUrl,proto3" json:"package_url,omitempty"`
	ForcePackage         bool              `protobuf:"varint,5,opt,name=force_package,json=forcePackage,proto3" json:"force_package,omitempty"`
	CaCerts              map[string]string `protobuf:"bytes,6,rep,name=ca_certs,json=caCerts,proto3" json:"ca_certs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyRequest) Reset()         { *m = ApplyRequest{} }
func (m *ApplyRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRequest) ProtoMessage()    {}
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_993661bab0ce9d1e, []int{0}
}

func (m *ApplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRequest.Unmarshal(m, b)
}
func (m *ApplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRequest.Marshal(b, m, deterministic)
}
func (m *ApplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRequest.Merge(m, src)
}
func (m *ApplyRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRequest.Size(m)
}
func (m *ApplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRequest proto.InternalMessageInfo

func (m *ApplyRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ApplyRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ApplyRequest) GetAction() ApplyAction {
	if m != nil {
		return m.Action
	}
	return ApplyAction_ALL
}

func (m *ApplyRequest) GetPackageUrl() string {
	if m != nil {
		return m.PackageUrl
	}
	return ""
}

func (m *ApplyRequest) GetForcePackage() bool {
	if m != nil {
		return m.ForcePackage
	}
	return false
}

func (m *ApplyRequest) GetCaCerts() map[string]string {
	if m != nil {
		return m.CaCerts
	}
	return nil
}

type ApplyResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyResponse) Reset()         { *m = ApplyResponse{} }
func (m *ApplyResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyResponse) ProtoMessage()    {}
func (*ApplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_993661bab0ce9d1e, []int{1}
}

func (m *ApplyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyResponse.Unmarshal(m, b)
}
func (m *ApplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyResponse.Marshal(b, m, deterministic)
}
func (m *ApplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyResponse.Merge(m, src)
}
func (m *ApplyResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyResponse.Size(m)
}
func (m *ApplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyResponse proto.InternalMessageInfo

func (m *ApplyResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ApplyResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterEnum("kubekit.v1.ApplyAction", ApplyAction_name, ApplyAction_value)
	proto.RegisterType((*ApplyRequest)(nil), "kubekit.v1.ApplyRequest")
	proto.RegisterMapType((map[string]string)(nil), "kubekit.v1.ApplyRequest.CaCertsEntry")
	proto.RegisterType((*ApplyResponse)(nil), "kubekit.v1.ApplyResponse")
}

func init() { proto.RegisterFile("apply.proto", fileDescriptor_993661bab0ce9d1e) }

var fileDescriptor_993661bab0ce9d1e = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xb6, 0xed, 0xd6, 0x6d, 0xaf, 0xad, 0x94, 0x20, 0x5a, 0x76, 0xb1, 0x4e, 0x84, 0xe2, 0xa1,
	0xb2, 0xe9, 0x41, 0x77, 0x72, 0x8e, 0x29, 0x83, 0xb1, 0x8d, 0xca, 0x3c, 0x78, 0x29, 0x59, 0x88,
	0x32, 0xda, 0xb5, 0x35, 0x49, 0x07, 0xfd, 0xcf, 0xfc, 0xf3, 0xa4, 0x6d, 0x86, 0x03, 0xbd, 0xe5,
	0xfb, 0xf1, 0xde, 0xc7, 0xf7, 0x02, 0x06, 0xce, 0xb2, 0xb8, 0xf0, 0x33, 0x96, 0x8a, 0x14, 0x41,
	0x94, 0xaf, 0x69, 0xb4, 0x11, 0xfe, 0xae, 0xdf, 0xb5, 0x48, 0x9c, 0x73, 0x41, 0x59, 0x2d, 0xf5,
	0xbe, 0x55, 0x30, 0x47, 0xa5, 0x35, 0xa0, 0x5f, 0x39, 0xe5, 0x02, 0xd9, 0xa0, 0xe1, 0x6c, 0xe3,
	0x28, 0xae, 0xe2, 0x75, 0x82, 0xf2, 0x89, 0x2e, 0xc0, 0x94, 0x33, 0x61, 0x82, 0xb7, 0xd4, 0x51,
	0x2b, 0xc9, 0x90, 0xdc, 0x1c, 0x6f, 0x29, 0xba, 0x01, 0x1d, 0x13, 0xb1, 0x49, 0x13, 0x47, 0x73,
	0x15, 0xef, 0x78, 0x70, 0xe6, 0xff, 0x26, 0xfa, 0xd5, 0xfa, 0x51, 0x25, 0x07, 0xd2, 0x86, 0xce,
	0xc1, 0xc8, 0x30, 0x89, 0xf0, 0x27, 0x0d, 0x73, 0x16, 0x3b, 0x8d, 0x6a, 0x25, 0x48, 0x6a, 0xc5,
	0x62, 0x74, 0x09, 0xd6, 0x47, 0xca, 0x08, 0x0d, 0x25, 0xe7, 0x34, 0x5d, 0xc5, 0x6b, 0x07, 0x66,
	0x45, 0x2e, 0x6b, 0x0e, 0x3d, 0x42, 0x9b, 0xe0, 0x90, 0x50, 0x26, 0xb8, 0xa3, 0xbb, 0x9a, 0x67,
	0x0c, 0xae, 0xfe, 0x04, 0xcb, 0x5e, 0xfe, 0x18, 0x8f, 0x4b, 0xdf, 0x24, 0x11, 0xac, 0x08, 0x5a,
	0xa4, 0x46, 0xdd, 0x21, 0x98, 0x87, 0x42, 0xd9, 0x3e, 0xa2, 0xc5, 0xbe, 0x7d, 0x44, 0x0b, 0x74,
	0x02, 0xcd, 0x1d, 0x8e, 0xf3, 0x7d, 0xed, 0x1a, 0x0c, 0xd5, 0x7b, 0xa5, 0xf7, 0x00, 0x96, 0x4c,
	0xe0, 0x59, 0x9a, 0x70, 0xfa, 0xcf, 0xe9, 0x4e, 0x41, 0xe7, 0x02, 0x8b, 0x9c, 0xcb, 0x69, 0x89,
	0xae, 0xef, 0xc0, 0x38, 0xb8, 0x0a, 0x6a, 0x81, 0x36, 0x9a, 0xcd, 0xec, 0x23, 0x64, 0x41, 0x67,
	0x19, 0x2c, 0xde, 0xa6, 0xaf, 0xd3, 0xc5, 0xdc, 0x56, 0x4a, 0x38, 0x5e, 0xcc, 0x9f, 0xa7, 0x2f,
	0xab, 0x60, 0x62, 0xab, 0x4f, 0x8d, 0x77, 0x75, 0xd7, 0x5f, 0xeb, 0xd5, 0xc7, 0xdd, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x3a, 0xe6, 0x27, 0xe2, 0x01, 0x00, 0x00,
}
