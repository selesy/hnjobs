// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hnjobs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ImportClient is the client API for Import service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImportClient interface {
	AddWhoIsHiring(ctx context.Context, in *ItemReference, opts ...grpc.CallOption) (*ImportStatus, error)
	UpsertJob(ctx context.Context, in *ItemReference, opts ...grpc.CallOption) (*JobStatus, error)
}

type importClient struct {
	cc grpc.ClientConnInterface
}

func NewImportClient(cc grpc.ClientConnInterface) ImportClient {
	return &importClient{cc}
}

func (c *importClient) AddWhoIsHiring(ctx context.Context, in *ItemReference, opts ...grpc.CallOption) (*ImportStatus, error) {
	out := new(ImportStatus)
	err := c.cc.Invoke(ctx, "/hnjobs.Import/AddWhoIsHiring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *importClient) UpsertJob(ctx context.Context, in *ItemReference, opts ...grpc.CallOption) (*JobStatus, error) {
	out := new(JobStatus)
	err := c.cc.Invoke(ctx, "/hnjobs.Import/UpsertJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImportServer is the server API for Import service.
// All implementations must embed UnimplementedImportServer
// for forward compatibility
type ImportServer interface {
	AddWhoIsHiring(context.Context, *ItemReference) (*ImportStatus, error)
	UpsertJob(context.Context, *ItemReference) (*JobStatus, error)
	mustEmbedUnimplementedImportServer()
}

// UnimplementedImportServer must be embedded to have forward compatible implementations.
type UnimplementedImportServer struct {
}

func (UnimplementedImportServer) AddWhoIsHiring(context.Context, *ItemReference) (*ImportStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWhoIsHiring not implemented")
}
func (UnimplementedImportServer) UpsertJob(context.Context, *ItemReference) (*JobStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertJob not implemented")
}
func (UnimplementedImportServer) mustEmbedUnimplementedImportServer() {}

// UnsafeImportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImportServer will
// result in compilation errors.
type UnsafeImportServer interface {
	mustEmbedUnimplementedImportServer()
}

func RegisterImportServer(s grpc.ServiceRegistrar, srv ImportServer) {
	s.RegisterService(&_Import_serviceDesc, srv)
}

func _Import_AddWhoIsHiring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).AddWhoIsHiring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hnjobs.Import/AddWhoIsHiring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).AddWhoIsHiring(ctx, req.(*ItemReference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Import_UpsertJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemReference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServer).UpsertJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hnjobs.Import/UpsertJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServer).UpsertJob(ctx, req.(*ItemReference))
	}
	return interceptor(ctx, in, info, handler)
}

var _Import_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hnjobs.Import",
	HandlerType: (*ImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddWhoIsHiring",
			Handler:    _Import_AddWhoIsHiring_Handler,
		},
		{
			MethodName: "UpsertJob",
			Handler:    _Import_UpsertJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hnjobs.proto",
}

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostClient interface {
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

// PostServer is the server API for Post service.
// All implementations must embed UnimplementedPostServer
// for forward compatibility
type PostServer interface {
	mustEmbedUnimplementedPostServer()
}

// UnimplementedPostServer must be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (UnimplementedPostServer) mustEmbedUnimplementedPostServer() {}

// UnsafePostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServer will
// result in compilation errors.
type UnsafePostServer interface {
	mustEmbedUnimplementedPostServer()
}

func RegisterPostServer(s grpc.ServiceRegistrar, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hnjobs.Post",
	HandlerType: (*PostServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "hnjobs.proto",
}