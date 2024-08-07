// Copyright 2024 The Jumpstarter Authors

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: jumpstarter/v1/router.proto

package jumpstarterv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RouterService_Stream_FullMethodName = "/jumpstarter.v1.RouterService/Stream"
)

// RouterServiceClient is the client API for RouterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// StreamService
// Claims:
// iss: jumpstarter controller
// aud: jumpstarter router
// sub: jumpstarter client/exporter
// stream: stream id
type RouterServiceClient interface {
	// Stream connects caller to another caller of the same stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamRequest, StreamResponse], error)
}

type routerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterServiceClient(cc grpc.ClientConnInterface) RouterServiceClient {
	return &routerServiceClient{cc}
}

func (c *routerServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamRequest, StreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RouterService_ServiceDesc.Streams[0], RouterService_Stream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRequest, StreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RouterService_StreamClient = grpc.BidiStreamingClient[StreamRequest, StreamResponse]

// RouterServiceServer is the server API for RouterService service.
// All implementations must embed UnimplementedRouterServiceServer
// for forward compatibility.
//
// StreamService
// Claims:
// iss: jumpstarter controller
// aud: jumpstarter router
// sub: jumpstarter client/exporter
// stream: stream id
type RouterServiceServer interface {
	// Stream connects caller to another caller of the same stream
	Stream(grpc.BidiStreamingServer[StreamRequest, StreamResponse]) error
	mustEmbedUnimplementedRouterServiceServer()
}

// UnimplementedRouterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRouterServiceServer struct{}

func (UnimplementedRouterServiceServer) Stream(grpc.BidiStreamingServer[StreamRequest, StreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedRouterServiceServer) mustEmbedUnimplementedRouterServiceServer() {}
func (UnimplementedRouterServiceServer) testEmbeddedByValue()                       {}

// UnsafeRouterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServiceServer will
// result in compilation errors.
type UnsafeRouterServiceServer interface {
	mustEmbedUnimplementedRouterServiceServer()
}

func RegisterRouterServiceServer(s grpc.ServiceRegistrar, srv RouterServiceServer) {
	// If the following call pancis, it indicates UnimplementedRouterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RouterService_ServiceDesc, srv)
}

func _RouterService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServiceServer).Stream(&grpc.GenericServerStream[StreamRequest, StreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RouterService_StreamServer = grpc.BidiStreamingServer[StreamRequest, StreamResponse]

// RouterService_ServiceDesc is the grpc.ServiceDesc for RouterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jumpstarter.v1.RouterService",
	HandlerType: (*RouterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _RouterService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "jumpstarter/v1/router.proto",
}
