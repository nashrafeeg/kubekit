// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

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

type PlatformName int32

const (
	PlatformName_UNKNOWN   PlatformName = 0
	PlatformName_AWS       PlatformName = 1
	PlatformName_EKS       PlatformName = 2
	PlatformName_AZURE     PlatformName = 3
	PlatformName_AKS       PlatformName = 4
	PlatformName_OPENSTACK PlatformName = 5
	PlatformName_VSPHERE   PlatformName = 6
	PlatformName_VRA       PlatformName = 7
	PlatformName_STACKI    PlatformName = 8
	PlatformName_RAW       PlatformName = 9
)

var PlatformName_name = map[int32]string{
	0: "UNKNOWN",
	1: "AWS",
	2: "EKS",
	3: "AZURE",
	4: "AKS",
	5: "OPENSTACK",
	6: "VSPHERE",
	7: "VRA",
	8: "STACKI",
	9: "RAW",
}

var PlatformName_value = map[string]int32{
	"UNKNOWN":   0,
	"AWS":       1,
	"EKS":       2,
	"AZURE":     3,
	"AKS":       4,
	"OPENSTACK": 5,
	"VSPHERE":   6,
	"VRA":       7,
	"STACKI":    8,
	"RAW":       9,
}

func (x PlatformName) String() string {
	return proto.EnumName(PlatformName_name, int32(x))
}

func (PlatformName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{0}
}

type Status int32

const (
	Status_UNKNOWN_STATUS       Status = 0
	Status_ABSENT               Status = 1
	Status_CREATING             Status = 2
	Status_PROVISIONED          Status = 3
	Status_FAILED_PROVISIONING  Status = 4
	Status_FAILED_CONFIGURATION Status = 5
	Status_CREATED              Status = 6
	Status_FAILED_CREATION      Status = 7
	Status_RUNNING              Status = 8
	Status_STOPPED              Status = 9
	Status_TERMINATING          Status = 10
	Status_TERMINATED           Status = 11
	Status_FAILED_TERMINATION   Status = 12
)

var Status_name = map[int32]string{
	0:  "UNKNOWN_STATUS",
	1:  "ABSENT",
	2:  "CREATING",
	3:  "PROVISIONED",
	4:  "FAILED_PROVISIONING",
	5:  "FAILED_CONFIGURATION",
	6:  "CREATED",
	7:  "FAILED_CREATION",
	8:  "RUNNING",
	9:  "STOPPED",
	10: "TERMINATING",
	11: "TERMINATED",
	12: "FAILED_TERMINATION",
}

var Status_value = map[string]int32{
	"UNKNOWN_STATUS":       0,
	"ABSENT":               1,
	"CREATING":             2,
	"PROVISIONED":          3,
	"FAILED_PROVISIONING":  4,
	"FAILED_CONFIGURATION": 5,
	"CREATED":              6,
	"FAILED_CREATION":      7,
	"RUNNING":              8,
	"STOPPED":              9,
	"TERMINATING":          10,
	"TERMINATED":           11,
	"FAILED_TERMINATION":   12,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{1}
}

type Cluster struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Platform             PlatformName `protobuf:"varint,4,opt,name=platform,proto3,enum=kubekit.v1.PlatformName" json:"platform,omitempty"`
	Nodes                int32        `protobuf:"varint,5,opt,name=nodes,proto3" json:"nodes,omitempty"`
	Status               Status       `protobuf:"varint,6,opt,name=status,proto3,enum=kubekit.v1.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{0}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetPlatform() PlatformName {
	if m != nil {
		return m.Platform
	}
	return PlatformName_UNKNOWN
}

func (m *Cluster) GetNodes() int32 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

func (m *Cluster) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN_STATUS
}

func init() {
	proto.RegisterEnum("kubekit.v1.PlatformName", PlatformName_name, PlatformName_value)
	proto.RegisterEnum("kubekit.v1.Status", Status_name, Status_value)
	proto.RegisterType((*Cluster)(nil), "kubekit.v1.Cluster")
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_3cfb3b8ec240c376) }

var fileDescriptor_3cfb3b8ec240c376 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0x87, 0x63, 0xf0, 0x1f, 0x3c, 0x10, 0x32, 0x9a, 0x44, 0xad, 0x8f, 0xa8, 0x27, 0xc4, 0x01,
	0x29, 0x6d, 0x5f, 0x60, 0x83, 0x37, 0xe9, 0x8a, 0x76, 0xd7, 0xda, 0x5d, 0x83, 0x94, 0x4b, 0x44,
	0x5a, 0x57, 0xaa, 0x12, 0x42, 0x84, 0x4d, 0x2e, 0x7d, 0x8e, 0x3e, 0x67, 0x5f, 0xa1, 0xda, 0xb5,
	0x43, 0x73, 0xdb, 0x99, 0xdf, 0x37, 0xdf, 0xcc, 0x61, 0xe1, 0xf4, 0xfb, 0xe3, 0xa1, 0x6e, 0xaa,
	0xfd, 0xfc, 0x79, 0xbf, 0x6b, 0x76, 0x04, 0x0f, 0x87, 0xfb, 0xea, 0xe1, 0x57, 0x33, 0x7f, 0xb9,
	0xfc, 0xf0, 0x27, 0x80, 0x64, 0xd1, 0xa6, 0x44, 0x10, 0x3e, 0x6d, 0xb6, 0x55, 0x16, 0x4c, 0x82,
	0x69, 0xaa, 0xfd, 0x9b, 0x3e, 0xc3, 0xe0, 0xf9, 0x71, 0xd3, 0xfc, 0xdc, 0xed, 0xb7, 0x59, 0x38,
	0x09, 0xa6, 0xe3, 0x8f, 0xd9, 0xfc, 0xff, 0xf8, 0xbc, 0xe8, 0x32, 0xb9, 0xd9, 0x56, 0xfa, 0x48,
	0xd2, 0x05, 0x44, 0x4f, 0xbb, 0x1f, 0x55, 0x9d, 0x45, 0x93, 0x60, 0x1a, 0xe9, 0xb6, 0xa0, 0x19,
	0xc4, 0x75, 0xb3, 0x69, 0x0e, 0x75, 0x16, 0x7b, 0x13, 0xbd, 0x35, 0x19, 0x9f, 0xe8, 0x8e, 0x98,
	0xfd, 0x86, 0xd1, 0x5b, 0x37, 0x0d, 0x21, 0x29, 0xe5, 0x52, 0xaa, 0xb5, 0xc4, 0x13, 0x4a, 0xa0,
	0xcf, 0xd6, 0x06, 0x03, 0xf7, 0xe0, 0x4b, 0x83, 0x3d, 0x4a, 0x21, 0x62, 0xb7, 0xa5, 0xe6, 0xd8,
	0xf7, 0xe1, 0xd2, 0x60, 0x48, 0xa7, 0x90, 0xaa, 0x82, 0x4b, 0x63, 0xd9, 0x62, 0x89, 0x91, 0x33,
	0xac, 0x4c, 0xf1, 0x85, 0x6b, 0x8e, 0xb1, 0x83, 0x56, 0x9a, 0x61, 0x42, 0x00, 0xb1, 0x07, 0x04,
	0x0e, 0x5c, 0x53, 0xb3, 0x35, 0xa6, 0xb3, 0xbf, 0x01, 0xc4, 0xed, 0x3d, 0x44, 0x30, 0xee, 0xf6,
	0xde, 0x19, 0xcb, 0x6c, 0x69, 0xf0, 0xc4, 0xcd, 0xb0, 0x2b, 0xc3, 0xa5, 0xc5, 0x80, 0x46, 0x30,
	0x58, 0x68, 0xce, 0xac, 0x90, 0x37, 0xd8, 0xa3, 0x33, 0x18, 0x16, 0x5a, 0xad, 0x84, 0x11, 0x4a,
	0xf2, 0x1c, 0xfb, 0xf4, 0x1e, 0xce, 0xaf, 0x99, 0xf8, 0xca, 0xf3, 0xbb, 0x63, 0xdf, 0x91, 0x21,
	0x65, 0x70, 0xd1, 0x05, 0x0b, 0x25, 0xaf, 0xc5, 0x4d, 0xa9, 0x99, 0x15, 0x4a, 0xb6, 0x77, 0x7a,
	0x23, 0xcf, 0x31, 0xa6, 0x73, 0x38, 0x7b, 0xc5, 0xfc, 0x16, 0x25, 0x31, 0x71, 0x84, 0x2e, 0xa5,
	0x17, 0x0d, 0x5c, 0x61, 0xac, 0x2a, 0x0a, 0x9e, 0x63, 0xea, 0xf6, 0x5b, 0xae, 0xbf, 0x09, 0xd9,
	0x1e, 0x04, 0x34, 0x06, 0x78, 0x6d, 0xf0, 0x1c, 0x87, 0xf4, 0x0e, 0xa8, 0xf3, 0x1d, 0x39, 0x25,
	0x71, 0x74, 0x15, 0xde, 0xf6, 0x5e, 0x2e, 0xef, 0x63, 0xff, 0x3f, 0x3e, 0xfd, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x23, 0xcc, 0xbb, 0x04, 0x30, 0x02, 0x00, 0x00,
}
