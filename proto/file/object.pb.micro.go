// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/file/object.proto

package file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Object service

type ObjectService interface {
	// 打开一个对象的存储通道
	Open(ctx context.Context, in *ObjectOpenRequest, opts ...client.CallOption) (*ObjectOpenResponse, error)
	// 上传一个对象
	Upload(ctx context.Context, opts ...client.CallOption) (Object_UploadService, error)
	// 下载一个对象
	Download(ctx context.Context, in *ObjectDownloadRequest, opts ...client.CallOption) (Object_DownloadService, error)
	// 链接一个外部对象
	Link(ctx context.Context, in *ObjectLinkRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 获取一个对象信息
	Get(ctx context.Context, in *ObjectGetRequest, opts ...client.CallOption) (*ObjectGetResponse, error)
	// 删除一个对象
	Remove(ctx context.Context, in *ObjectRemoveRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 列举一个存储桶中的所有对象
	List(ctx context.Context, in *ObjectListRequest, opts ...client.CallOption) (*ObjectListResponse, error)
}

type objectService struct {
	c    client.Client
	name string
}

func NewObjectService(name string, c client.Client) ObjectService {
	return &objectService{
		c:    c,
		name: name,
	}
}

func (c *objectService) Open(ctx context.Context, in *ObjectOpenRequest, opts ...client.CallOption) (*ObjectOpenResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Open", in)
	out := new(ObjectOpenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Upload(ctx context.Context, opts ...client.CallOption) (Object_UploadService, error) {
	req := c.c.NewRequest(c.name, "Object.Upload", &ObjectUploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &objectServiceUpload{stream}, nil
}

type Object_UploadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ObjectUploadRequest) error
}

type objectServiceUpload struct {
	stream client.Stream
}

func (x *objectServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *objectServiceUpload) Context() context.Context {
	return x.stream.Context()
}

func (x *objectServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *objectServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *objectServiceUpload) Send(m *ObjectUploadRequest) error {
	return x.stream.Send(m)
}

func (c *objectService) Download(ctx context.Context, in *ObjectDownloadRequest, opts ...client.CallOption) (Object_DownloadService, error) {
	req := c.c.NewRequest(c.name, "Object.Download", &ObjectDownloadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &objectServiceDownload{stream}, nil
}

type Object_DownloadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ObjectDownloadResponse, error)
}

type objectServiceDownload struct {
	stream client.Stream
}

func (x *objectServiceDownload) Close() error {
	return x.stream.Close()
}

func (x *objectServiceDownload) Context() context.Context {
	return x.stream.Context()
}

func (x *objectServiceDownload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *objectServiceDownload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *objectServiceDownload) Recv() (*ObjectDownloadResponse, error) {
	m := new(ObjectDownloadResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectService) Link(ctx context.Context, in *ObjectLinkRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Link", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Get(ctx context.Context, in *ObjectGetRequest, opts ...client.CallOption) (*ObjectGetResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Get", in)
	out := new(ObjectGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) Remove(ctx context.Context, in *ObjectRemoveRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Object.Remove", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectService) List(ctx context.Context, in *ObjectListRequest, opts ...client.CallOption) (*ObjectListResponse, error) {
	req := c.c.NewRequest(c.name, "Object.List", in)
	out := new(ObjectListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Object service

type ObjectHandler interface {
	// 打开一个对象的存储通道
	Open(context.Context, *ObjectOpenRequest, *ObjectOpenResponse) error
	// 上传一个对象
	Upload(context.Context, Object_UploadStream) error
	// 下载一个对象
	Download(context.Context, *ObjectDownloadRequest, Object_DownloadStream) error
	// 链接一个外部对象
	Link(context.Context, *ObjectLinkRequest, *BlankResponse) error
	// 获取一个对象信息
	Get(context.Context, *ObjectGetRequest, *ObjectGetResponse) error
	// 删除一个对象
	Remove(context.Context, *ObjectRemoveRequest, *BlankResponse) error
	// 列举一个存储桶中的所有对象
	List(context.Context, *ObjectListRequest, *ObjectListResponse) error
}

func RegisterObjectHandler(s server.Server, hdlr ObjectHandler, opts ...server.HandlerOption) error {
	type object interface {
		Open(ctx context.Context, in *ObjectOpenRequest, out *ObjectOpenResponse) error
		Upload(ctx context.Context, stream server.Stream) error
		Download(ctx context.Context, stream server.Stream) error
		Link(ctx context.Context, in *ObjectLinkRequest, out *BlankResponse) error
		Get(ctx context.Context, in *ObjectGetRequest, out *ObjectGetResponse) error
		Remove(ctx context.Context, in *ObjectRemoveRequest, out *BlankResponse) error
		List(ctx context.Context, in *ObjectListRequest, out *ObjectListResponse) error
	}
	type Object struct {
		object
	}
	h := &objectHandler{hdlr}
	return s.Handle(s.NewHandler(&Object{h}, opts...))
}

type objectHandler struct {
	ObjectHandler
}

func (h *objectHandler) Open(ctx context.Context, in *ObjectOpenRequest, out *ObjectOpenResponse) error {
	return h.ObjectHandler.Open(ctx, in, out)
}

func (h *objectHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.ObjectHandler.Upload(ctx, &objectUploadStream{stream})
}

type Object_UploadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ObjectUploadRequest, error)
}

type objectUploadStream struct {
	stream server.Stream
}

func (x *objectUploadStream) Close() error {
	return x.stream.Close()
}

func (x *objectUploadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *objectUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *objectUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *objectUploadStream) Recv() (*ObjectUploadRequest, error) {
	m := new(ObjectUploadRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *objectHandler) Download(ctx context.Context, stream server.Stream) error {
	m := new(ObjectDownloadRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ObjectHandler.Download(ctx, m, &objectDownloadStream{stream})
}

type Object_DownloadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ObjectDownloadResponse) error
}

type objectDownloadStream struct {
	stream server.Stream
}

func (x *objectDownloadStream) Close() error {
	return x.stream.Close()
}

func (x *objectDownloadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *objectDownloadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *objectDownloadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *objectDownloadStream) Send(m *ObjectDownloadResponse) error {
	return x.stream.Send(m)
}

func (h *objectHandler) Link(ctx context.Context, in *ObjectLinkRequest, out *BlankResponse) error {
	return h.ObjectHandler.Link(ctx, in, out)
}

func (h *objectHandler) Get(ctx context.Context, in *ObjectGetRequest, out *ObjectGetResponse) error {
	return h.ObjectHandler.Get(ctx, in, out)
}

func (h *objectHandler) Remove(ctx context.Context, in *ObjectRemoveRequest, out *BlankResponse) error {
	return h.ObjectHandler.Remove(ctx, in, out)
}

func (h *objectHandler) List(ctx context.Context, in *ObjectListRequest, out *ObjectListResponse) error {
	return h.ObjectHandler.List(ctx, in, out)
}
