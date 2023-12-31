// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/trail/trail.proto

package trail

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataTrail_CreateTrail_FullMethodName = "/trail.DataTrail/CreateTrail"
)

// DataTrailClient is the client API for DataTrail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataTrailClient interface {
	CreateTrail(ctx context.Context, in *Trail, opts ...grpc.CallOption) (*Trail, error)
}

type dataTrailClient struct {
	cc grpc.ClientConnInterface
}

func NewDataTrailClient(cc grpc.ClientConnInterface) DataTrailClient {
	return &dataTrailClient{cc}
}

func (c *dataTrailClient) CreateTrail(ctx context.Context, in *Trail, opts ...grpc.CallOption) (*Trail, error) {
	out := new(Trail)
	err := c.cc.Invoke(ctx, DataTrail_CreateTrail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataTrailServer is the server API for DataTrail service.
// All implementations must embed UnimplementedDataTrailServer
// for forward compatibility
type DataTrailServer interface {
	CreateTrail(context.Context, *Trail) (*Trail, error)
	mustEmbedUnimplementedDataTrailServer()
}

// UnimplementedDataTrailServer must be embedded to have forward compatible implementations.
type UnimplementedDataTrailServer struct {
}

func (UnimplementedDataTrailServer) CreateTrail(context.Context, *Trail) (*Trail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrail not implemented")
}
func (UnimplementedDataTrailServer) mustEmbedUnimplementedDataTrailServer() {}

// UnsafeDataTrailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataTrailServer will
// result in compilation errors.
type UnsafeDataTrailServer interface {
	mustEmbedUnimplementedDataTrailServer()
}

func RegisterDataTrailServer(s grpc.ServiceRegistrar, srv DataTrailServer) {
	s.RegisterService(&DataTrail_ServiceDesc, srv)
}

func _DataTrail_CreateTrail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataTrailServer).CreateTrail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataTrail_CreateTrail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataTrailServer).CreateTrail(ctx, req.(*Trail))
	}
	return interceptor(ctx, in, info, handler)
}

// DataTrail_ServiceDesc is the grpc.ServiceDesc for DataTrail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataTrail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trail.DataTrail",
	HandlerType: (*DataTrailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrail",
			Handler:    _DataTrail_CreateTrail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/trail/trail.proto",
}
