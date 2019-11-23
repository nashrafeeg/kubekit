// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete.proto

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

type DeleteClusterConfigStatus int32

const (
	DeleteClusterConfigStatus_DELETED   DeleteClusterConfigStatus = 0
	DeleteClusterConfigStatus_NOT_FOUND DeleteClusterConfigStatus = 1
)

var DeleteClusterConfigStatus_name = map[int32]string{
	0: "DELETED",
	1: "NOT_FOUND",
}

var DeleteClusterConfigStatus_value = map[string]int32{
	"DELETED":   0,
	"NOT_FOUND": 1,
}

func (x DeleteClusterConfigStatus) String() string {
	return proto.EnumName(DeleteClusterConfigStatus_name, int32(x))
}

func (DeleteClusterConfigStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}

type DeleteRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ClusterName          string   `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	DestroyAll           bool     `protobuf:"varint,3,opt,name=destroy_all,json=destroyAll,proto3" json:"destroy_all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{0}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *DeleteRequest) GetDestroyAll() bool {
	if m != nil {
		return m.DestroyAll
	}
	return false
}

type DeleteResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{1}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteClusterConfigRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ClusterName          string   `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteClusterConfigRequest) Reset()         { *m = DeleteClusterConfigRequest{} }
func (m *DeleteClusterConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterConfigRequest) ProtoMessage()    {}
func (*DeleteClusterConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{2}
}

func (m *DeleteClusterConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterConfigRequest.Unmarshal(m, b)
}
func (m *DeleteClusterConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterConfigRequest.Marshal(b, m, deterministic)
}
func (m *DeleteClusterConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterConfigRequest.Merge(m, src)
}
func (m *DeleteClusterConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterConfigRequest.Size(m)
}
func (m *DeleteClusterConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterConfigRequest proto.InternalMessageInfo

func (m *DeleteClusterConfigRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteClusterConfigRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type DeleteClusterConfigResponse struct {
	Api                  string                    `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ClusterName          string                    `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Status               DeleteClusterConfigStatus `protobuf:"varint,3,opt,name=status,proto3,enum=kubekit.v1.DeleteClusterConfigStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DeleteClusterConfigResponse) Reset()         { *m = DeleteClusterConfigResponse{} }
func (m *DeleteClusterConfigResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterConfigResponse) ProtoMessage()    {}
func (*DeleteClusterConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_600d681a62b3a9a7, []int{3}
}

func (m *DeleteClusterConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterConfigResponse.Unmarshal(m, b)
}
func (m *DeleteClusterConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterConfigResponse.Marshal(b, m, deterministic)
}
func (m *DeleteClusterConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterConfigResponse.Merge(m, src)
}
func (m *DeleteClusterConfigResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterConfigResponse.Size(m)
}
func (m *DeleteClusterConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterConfigResponse proto.InternalMessageInfo

func (m *DeleteClusterConfigResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteClusterConfigResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *DeleteClusterConfigResponse) GetStatus() DeleteClusterConfigStatus {
	if m != nil {
		return m.Status
	}
	return DeleteClusterConfigStatus_DELETED
}

func init() {
	proto.RegisterEnum("kubekit.v1.DeleteClusterConfigStatus", DeleteClusterConfigStatus_name, DeleteClusterConfigStatus_value)
	proto.RegisterType((*DeleteRequest)(nil), "kubekit.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "kubekit.v1.DeleteResponse")
	proto.RegisterType((*DeleteClusterConfigRequest)(nil), "kubekit.v1.DeleteClusterConfigRequest")
	proto.RegisterType((*DeleteClusterConfigResponse)(nil), "kubekit.v1.DeleteClusterConfigResponse")
}

func init() { proto.RegisterFile("delete.proto", fileDescriptor_600d681a62b3a9a7) }

var fileDescriptor_600d681a62b3a9a7 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x2e, 0x4d, 0x4a, 0xcd, 0xce,
	0x2c, 0xd1, 0x2b, 0x33, 0x54, 0x4a, 0xe5, 0xe2, 0x75, 0x01, 0xcb, 0x05, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x09, 0x70, 0x31, 0x27, 0x16, 0x64, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x81, 0x98, 0x42, 0x8a, 0x5c, 0x3c, 0xc9, 0x39, 0xa5, 0xc5, 0x25, 0xa9, 0x45, 0xf1, 0x79, 0x89,
	0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x29, 0x6e, 0xa8, 0x98, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x3c, 0x17,
	0x77, 0x4a, 0x6a, 0x71, 0x49, 0x51, 0x7e, 0x65, 0x7c, 0x62, 0x4e, 0x8e, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x47, 0x10, 0x17, 0x54, 0xc8, 0x31, 0x27, 0x47, 0xc9, 0x8a, 0x8b, 0x0f, 0x66, 0x4d, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x16, 0x7b, 0xc4, 0xb8, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b,
	0xa1, 0x36, 0x40, 0x79, 0x4a, 0x81, 0x5c, 0x52, 0x10, 0xbd, 0xce, 0x10, 0x1b, 0x9d, 0xf3, 0xf3,
	0xd2, 0x32, 0xd3, 0x29, 0x71, 0xaf, 0xd2, 0x44, 0x46, 0x2e, 0x69, 0xac, 0x66, 0xe2, 0x74, 0x1c,
	0x11, 0x81, 0x60, 0x0b, 0x77, 0x3f, 0xc8, 0xff, 0x7c, 0x46, 0xaa, 0x7a, 0x88, 0x70, 0xd6, 0xc3,
	0x62, 0x5b, 0x30, 0x58, 0x31, 0xcc, 0x9b, 0x5a, 0xe6, 0x5c, 0x92, 0x38, 0x15, 0x09, 0x71, 0x73,
	0xb1, 0xbb, 0xb8, 0xfa, 0xb8, 0x86, 0xb8, 0xba, 0x08, 0x30, 0x08, 0xf1, 0x72, 0x71, 0xfa, 0xf9,
	0x87, 0xc4, 0xbb, 0xf9, 0x87, 0xfa, 0xb9, 0x08, 0x30, 0x3a, 0xb1, 0x44, 0x31, 0x95, 0x19, 0x26,
	0xb1, 0x81, 0xe3, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xba, 0x46, 0x2f, 0xf5, 0xeb, 0x01,
	0x00, 0x00,
}