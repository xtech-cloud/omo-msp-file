// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/file/bucket.proto

package file

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

// 创建存储桶的请求
type BucketMakeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             uint64   `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Engine               Engine   `protobuf:"varint,3,opt,name=engine,proto3,enum=file.Engine" json:"engine,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Scope                string   `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	AccessKey            string   `protobuf:"bytes,6,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	AccessSecret         string   `protobuf:"bytes,7,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketMakeRequest) Reset()         { *m = BucketMakeRequest{} }
func (m *BucketMakeRequest) String() string { return proto.CompactTextString(m) }
func (*BucketMakeRequest) ProtoMessage()    {}
func (*BucketMakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{0}
}

func (m *BucketMakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketMakeRequest.Unmarshal(m, b)
}
func (m *BucketMakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketMakeRequest.Marshal(b, m, deterministic)
}
func (m *BucketMakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketMakeRequest.Merge(m, src)
}
func (m *BucketMakeRequest) XXX_Size() int {
	return xxx_messageInfo_BucketMakeRequest.Size(m)
}
func (m *BucketMakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketMakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketMakeRequest proto.InternalMessageInfo

func (m *BucketMakeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BucketMakeRequest) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *BucketMakeRequest) GetEngine() Engine {
	if m != nil {
		return m.Engine
	}
	return Engine_ENGINE_INVALID
}

func (m *BucketMakeRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BucketMakeRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *BucketMakeRequest) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *BucketMakeRequest) GetAccessSecret() string {
	if m != nil {
		return m.AccessSecret
	}
	return ""
}

// 列举存储桶的回复
type BucketListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketListRequest) Reset()         { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()    {}
func (*BucketListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{1}
}

func (m *BucketListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketListRequest.Unmarshal(m, b)
}
func (m *BucketListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketListRequest.Marshal(b, m, deterministic)
}
func (m *BucketListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketListRequest.Merge(m, src)
}
func (m *BucketListRequest) XXX_Size() int {
	return xxx_messageInfo_BucketListRequest.Size(m)
}
func (m *BucketListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketListRequest proto.InternalMessageInfo

func (m *BucketListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BucketListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 列举存储桶的回复
type BucketListResponse struct {
	Status               *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*BucketEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BucketListResponse) Reset()         { *m = BucketListResponse{} }
func (m *BucketListResponse) String() string { return proto.CompactTextString(m) }
func (*BucketListResponse) ProtoMessage()    {}
func (*BucketListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{2}
}

func (m *BucketListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketListResponse.Unmarshal(m, b)
}
func (m *BucketListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketListResponse.Marshal(b, m, deterministic)
}
func (m *BucketListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketListResponse.Merge(m, src)
}
func (m *BucketListResponse) XXX_Size() int {
	return xxx_messageInfo_BucketListResponse.Size(m)
}
func (m *BucketListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BucketListResponse proto.InternalMessageInfo

func (m *BucketListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BucketListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *BucketListResponse) GetEntity() []*BucketEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 删除存储桶的请求
type BucketRemoveRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketRemoveRequest) Reset()         { *m = BucketRemoveRequest{} }
func (m *BucketRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*BucketRemoveRequest) ProtoMessage()    {}
func (*BucketRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{3}
}

func (m *BucketRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketRemoveRequest.Unmarshal(m, b)
}
func (m *BucketRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketRemoveRequest.Marshal(b, m, deterministic)
}
func (m *BucketRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketRemoveRequest.Merge(m, src)
}
func (m *BucketRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_BucketRemoveRequest.Size(m)
}
func (m *BucketRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketRemoveRequest proto.InternalMessageInfo

func (m *BucketRemoveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取存储桶信息的请求
type BucketGetRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketGetRequest) Reset()         { *m = BucketGetRequest{} }
func (m *BucketGetRequest) String() string { return proto.CompactTextString(m) }
func (*BucketGetRequest) ProtoMessage()    {}
func (*BucketGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{4}
}

func (m *BucketGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketGetRequest.Unmarshal(m, b)
}
func (m *BucketGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketGetRequest.Marshal(b, m, deterministic)
}
func (m *BucketGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketGetRequest.Merge(m, src)
}
func (m *BucketGetRequest) XXX_Size() int {
	return xxx_messageInfo_BucketGetRequest.Size(m)
}
func (m *BucketGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketGetRequest proto.InternalMessageInfo

func (m *BucketGetRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取存储桶信息的回复
type BucketGetResponse struct {
	Status               *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *BucketEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BucketGetResponse) Reset()         { *m = BucketGetResponse{} }
func (m *BucketGetResponse) String() string { return proto.CompactTextString(m) }
func (*BucketGetResponse) ProtoMessage()    {}
func (*BucketGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{5}
}

func (m *BucketGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketGetResponse.Unmarshal(m, b)
}
func (m *BucketGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketGetResponse.Marshal(b, m, deterministic)
}
func (m *BucketGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketGetResponse.Merge(m, src)
}
func (m *BucketGetResponse) XXX_Size() int {
	return xxx_messageInfo_BucketGetResponse.Size(m)
}
func (m *BucketGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BucketGetResponse proto.InternalMessageInfo

func (m *BucketGetResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BucketGetResponse) GetEntity() *BucketEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 查找一个存储桶信息的请求
type BucketFindRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketFindRequest) Reset()         { *m = BucketFindRequest{} }
func (m *BucketFindRequest) String() string { return proto.CompactTextString(m) }
func (*BucketFindRequest) ProtoMessage()    {}
func (*BucketFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{6}
}

func (m *BucketFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketFindRequest.Unmarshal(m, b)
}
func (m *BucketFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketFindRequest.Marshal(b, m, deterministic)
}
func (m *BucketFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketFindRequest.Merge(m, src)
}
func (m *BucketFindRequest) XXX_Size() int {
	return xxx_messageInfo_BucketFindRequest.Size(m)
}
func (m *BucketFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketFindRequest proto.InternalMessageInfo

func (m *BucketFindRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 查找一个存储桶信息的回复
type BucketFindResponse struct {
	Status               *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *BucketEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BucketFindResponse) Reset()         { *m = BucketFindResponse{} }
func (m *BucketFindResponse) String() string { return proto.CompactTextString(m) }
func (*BucketFindResponse) ProtoMessage()    {}
func (*BucketFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{7}
}

func (m *BucketFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketFindResponse.Unmarshal(m, b)
}
func (m *BucketFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketFindResponse.Marshal(b, m, deterministic)
}
func (m *BucketFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketFindResponse.Merge(m, src)
}
func (m *BucketFindResponse) XXX_Size() int {
	return xxx_messageInfo_BucketFindResponse.Size(m)
}
func (m *BucketFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BucketFindResponse proto.InternalMessageInfo

func (m *BucketFindResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BucketFindResponse) GetEntity() *BucketEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 更新存储桶引擎的请求
type BucketUpdateEngineRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Engine               Engine   `protobuf:"varint,2,opt,name=engine,proto3,enum=file.Engine" json:"engine,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Scope                string   `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	AccessKey            string   `protobuf:"bytes,5,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	AccessSecret         string   `protobuf:"bytes,6,opt,name=accessSecret,proto3" json:"accessSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketUpdateEngineRequest) Reset()         { *m = BucketUpdateEngineRequest{} }
func (m *BucketUpdateEngineRequest) String() string { return proto.CompactTextString(m) }
func (*BucketUpdateEngineRequest) ProtoMessage()    {}
func (*BucketUpdateEngineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{8}
}

func (m *BucketUpdateEngineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketUpdateEngineRequest.Unmarshal(m, b)
}
func (m *BucketUpdateEngineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketUpdateEngineRequest.Marshal(b, m, deterministic)
}
func (m *BucketUpdateEngineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketUpdateEngineRequest.Merge(m, src)
}
func (m *BucketUpdateEngineRequest) XXX_Size() int {
	return xxx_messageInfo_BucketUpdateEngineRequest.Size(m)
}
func (m *BucketUpdateEngineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketUpdateEngineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketUpdateEngineRequest proto.InternalMessageInfo

func (m *BucketUpdateEngineRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BucketUpdateEngineRequest) GetEngine() Engine {
	if m != nil {
		return m.Engine
	}
	return Engine_ENGINE_INVALID
}

func (m *BucketUpdateEngineRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BucketUpdateEngineRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *BucketUpdateEngineRequest) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *BucketUpdateEngineRequest) GetAccessSecret() string {
	if m != nil {
		return m.AccessSecret
	}
	return ""
}

// 更新存储桶的容量的请求
type BucketUpdateCapacityRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Capacity             uint64   `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketUpdateCapacityRequest) Reset()         { *m = BucketUpdateCapacityRequest{} }
func (m *BucketUpdateCapacityRequest) String() string { return proto.CompactTextString(m) }
func (*BucketUpdateCapacityRequest) ProtoMessage()    {}
func (*BucketUpdateCapacityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{9}
}

func (m *BucketUpdateCapacityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketUpdateCapacityRequest.Unmarshal(m, b)
}
func (m *BucketUpdateCapacityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketUpdateCapacityRequest.Marshal(b, m, deterministic)
}
func (m *BucketUpdateCapacityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketUpdateCapacityRequest.Merge(m, src)
}
func (m *BucketUpdateCapacityRequest) XXX_Size() int {
	return xxx_messageInfo_BucketUpdateCapacityRequest.Size(m)
}
func (m *BucketUpdateCapacityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketUpdateCapacityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketUpdateCapacityRequest proto.InternalMessageInfo

func (m *BucketUpdateCapacityRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BucketUpdateCapacityRequest) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

// 重置存储桶的访问令牌
type BucketResetTokenRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketResetTokenRequest) Reset()         { *m = BucketResetTokenRequest{} }
func (m *BucketResetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*BucketResetTokenRequest) ProtoMessage()    {}
func (*BucketResetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74316b9dda746edf, []int{10}
}

func (m *BucketResetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketResetTokenRequest.Unmarshal(m, b)
}
func (m *BucketResetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketResetTokenRequest.Marshal(b, m, deterministic)
}
func (m *BucketResetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketResetTokenRequest.Merge(m, src)
}
func (m *BucketResetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_BucketResetTokenRequest.Size(m)
}
func (m *BucketResetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketResetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BucketResetTokenRequest proto.InternalMessageInfo

func (m *BucketResetTokenRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*BucketMakeRequest)(nil), "file.BucketMakeRequest")
	proto.RegisterType((*BucketListRequest)(nil), "file.BucketListRequest")
	proto.RegisterType((*BucketListResponse)(nil), "file.BucketListResponse")
	proto.RegisterType((*BucketRemoveRequest)(nil), "file.BucketRemoveRequest")
	proto.RegisterType((*BucketGetRequest)(nil), "file.BucketGetRequest")
	proto.RegisterType((*BucketGetResponse)(nil), "file.BucketGetResponse")
	proto.RegisterType((*BucketFindRequest)(nil), "file.BucketFindRequest")
	proto.RegisterType((*BucketFindResponse)(nil), "file.BucketFindResponse")
	proto.RegisterType((*BucketUpdateEngineRequest)(nil), "file.BucketUpdateEngineRequest")
	proto.RegisterType((*BucketUpdateCapacityRequest)(nil), "file.BucketUpdateCapacityRequest")
	proto.RegisterType((*BucketResetTokenRequest)(nil), "file.BucketResetTokenRequest")
}

func init() {
	proto.RegisterFile("proto/file/bucket.proto", fileDescriptor_74316b9dda746edf)
}

var fileDescriptor_74316b9dda746edf = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0x5e, 0x96, 0x34, 0x63, 0xb7, 0x6a, 0x82, 0x1b, 0x5a, 0xb3, 0x02, 0xa2, 0x58, 0x08, 0x0a,
	0x12, 0x9d, 0x54, 0x24, 0x84, 0xe0, 0x17, 0x43, 0xdb, 0x24, 0x60, 0x7f, 0x32, 0x78, 0x00, 0x2f,
	0xb9, 0x42, 0xd4, 0x2e, 0x29, 0xb5, 0x83, 0x34, 0x89, 0x07, 0xe0, 0xcd, 0x78, 0x08, 0x5e, 0x06,
	0xd9, 0x8e, 0x57, 0x67, 0x5d, 0x02, 0xfc, 0xe0, 0x5f, 0xee, 0xee, 0xf3, 0xf9, 0xbe, 0xbb, 0xef,
	0x1c, 0xe8, 0xcd, 0x17, 0x85, 0x2c, 0xf6, 0x27, 0xd9, 0x8c, 0xf6, 0xcf, 0xca, 0x64, 0x4a, 0x72,
	0xa4, 0x3d, 0x18, 0x28, 0x57, 0xdf, 0x0d, 0x8b, 0x2f, 0x7c, 0x41, 0xa9, 0x09, 0xb3, 0x5f, 0x1e,
	0xdc, 0x3a, 0xd0, 0xf8, 0x13, 0x3e, 0xa5, 0x98, 0xbe, 0x96, 0x24, 0x24, 0x22, 0x04, 0x39, 0x3f,
	0xa7, 0xc8, 0x1b, 0x78, 0xc3, 0xcd, 0x58, 0x7f, 0x63, 0x1f, 0x6e, 0x24, 0x7c, 0xce, 0x93, 0x4c,
	0x5e, 0x44, 0xeb, 0x03, 0x6f, 0x18, 0xc4, 0x97, 0x36, 0x3e, 0x84, 0x90, 0xf2, 0xcf, 0x59, 0x4e,
	0x91, 0x3f, 0xf0, 0x86, 0xdb, 0xe3, 0xee, 0x48, 0xdd, 0x34, 0x3a, 0xd4, 0xbe, 0xb8, 0x8a, 0x61,
	0x04, 0x1b, 0x3c, 0x4d, 0x17, 0x24, 0x44, 0x14, 0xe8, 0xc4, 0xd6, 0xc4, 0xdb, 0xd0, 0x11, 0x49,
	0x31, 0xa7, 0xa8, 0xa3, 0xfd, 0xc6, 0xc0, 0xbb, 0xb0, 0xc9, 0x93, 0x84, 0x84, 0x78, 0x4f, 0x17,
	0x51, 0xa8, 0x23, 0x4b, 0x07, 0x32, 0xe8, 0x1a, 0xe3, 0x94, 0x92, 0x05, 0xc9, 0x68, 0x43, 0x03,
	0x6a, 0x3e, 0xf6, 0xc6, 0x92, 0xfb, 0x90, 0x09, 0x69, 0xc9, 0xed, 0x42, 0x58, 0x4c, 0x26, 0x82,
	0xa4, 0xa6, 0xe7, 0xc7, 0x95, 0xa5, 0x8a, 0x48, 0x8a, 0x32, 0x97, 0x9a, 0x9d, 0x1f, 0x1b, 0x83,
	0x7d, 0x07, 0x74, 0x53, 0x88, 0x79, 0x91, 0x0b, 0x52, 0x84, 0x85, 0xe4, 0xb2, 0x14, 0x3a, 0xc7,
	0x96, 0x25, 0x7c, 0xaa, 0x7d, 0x71, 0x15, 0x53, 0x19, 0x65, 0x21, 0xf9, 0xac, 0xea, 0x97, 0x31,
	0xf0, 0xa9, 0x6a, 0x96, 0x54, 0x6d, 0xf4, 0x07, 0xfe, 0x70, 0x6b, 0x8c, 0xe6, 0xac, 0xb9, 0xe5,
	0x50, 0x47, 0xe2, 0x0a, 0xc1, 0x9e, 0xc0, 0x8e, 0xf1, 0xc7, 0x74, 0x5e, 0x7c, 0x73, 0xe7, 0x53,
	0x96, 0x59, 0x6a, 0xe7, 0xa3, 0xbe, 0xd9, 0x23, 0xb8, 0x69, 0xa0, 0xc7, 0x24, 0xdb, 0x70, 0x64,
	0x7b, 0xa2, 0x71, 0xff, 0xc4, 0x67, 0x59, 0xf9, 0xba, 0x46, 0xb5, 0x55, 0xfe, 0xd8, 0x5e, 0x73,
	0x94, 0xe5, 0x69, 0x8b, 0xae, 0xd8, 0xc4, 0x36, 0xd8, 0x00, 0xff, 0x5b, 0x41, 0x3f, 0x3d, 0xd8,
	0x33, 0x81, 0x4f, 0xf3, 0x94, 0x4b, 0xaa, 0xc4, 0xd9, 0xdc, 0x29, 0x47, 0xd5, 0xeb, 0x7f, 0xa7,
	0x6a, 0xbf, 0x41, 0xd5, 0x41, 0xa3, 0xaa, 0x3b, 0x7f, 0x52, 0x75, 0x78, 0x8d, 0xaa, 0x4f, 0xe0,
	0x8e, 0x4b, 0xe4, 0x6d, 0xb5, 0x85, 0x6d, 0x54, 0x5a, 0x96, 0x97, 0x3d, 0x83, 0x9e, 0xd5, 0x98,
	0x20, 0xf9, 0xb1, 0x98, 0x52, 0xde, 0x92, 0x6a, 0xfc, 0x23, 0x80, 0xd0, 0xe0, 0xf1, 0x05, 0x04,
	0xea, 0xd5, 0xc0, 0x9e, 0xdb, 0x76, 0xe7, 0x1d, 0xe9, 0xef, 0x54, 0x81, 0x19, 0xcf, 0xa7, 0x76,
	0xb4, 0x6c, 0x0d, 0x5f, 0x43, 0xa0, 0xb6, 0xa9, 0x7e, 0xce, 0x59, 0xd1, 0x7e, 0xb4, 0x1a, 0xb8,
	0x3c, 0xfc, 0x0a, 0x42, 0xb3, 0x0c, 0xb8, 0xe7, 0xa2, 0x6a, 0x0b, 0xd2, 0x74, 0xf1, 0x4b, 0xf0,
	0x8f, 0x49, 0xe2, 0xae, 0x7b, 0x70, 0xb9, 0x2e, 0xfd, 0xde, 0x8a, 0xdf, 0x2d, 0x59, 0xe9, 0xb3,
	0x5e, 0xb2, 0x23, 0xed, 0x7a, 0xc9, 0xae, 0x94, 0xd9, 0x1a, 0x1e, 0x41, 0xd7, 0xd5, 0x1c, 0xde,
	0x77, 0xb1, 0xd7, 0xa8, 0xb1, 0xa9, 0xfc, 0x77, 0xb0, 0x5d, 0x1f, 0x39, 0x3e, 0x58, 0xcd, 0x74,
	0x45, 0x0e, 0x4d, 0xb9, 0x0e, 0x00, 0x96, 0xf3, 0xc6, 0x7b, 0xf5, 0x56, 0x5e, 0xd1, 0x41, 0x43,
	0x8e, 0xb3, 0x50, 0xff, 0x43, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x13, 0x58, 0xe4, 0x55,
	0x7d, 0x06, 0x00, 0x00,
}
