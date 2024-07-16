// Copyright 2024 The Jumpstarter Authors

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: jumpstarter/v1/jumpstarter.proto

package jumpstarterv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ControllerService_Register_FullMethodName        = "/jumpstarter.v1.ControllerService/Register"
	ControllerService_Bye_FullMethodName             = "/jumpstarter.v1.ControllerService/Bye"
	ControllerService_Listen_FullMethodName          = "/jumpstarter.v1.ControllerService/Listen"
	ControllerService_Dial_FullMethodName            = "/jumpstarter.v1.ControllerService/Dial"
	ControllerService_AuditStream_FullMethodName     = "/jumpstarter.v1.ControllerService/AuditStream"
	ControllerService_ListExporters_FullMethodName   = "/jumpstarter.v1.ControllerService/ListExporters"
	ControllerService_GetExporter_FullMethodName     = "/jumpstarter.v1.ControllerService/GetExporter"
	ControllerService_LeaseExporter_FullMethodName   = "/jumpstarter.v1.ControllerService/LeaseExporter"
	ControllerService_ReleaseExporter_FullMethodName = "/jumpstarter.v1.ControllerService/ReleaseExporter"
)

// ControllerServiceClient is the client API for ControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service where a exporter can connect to make itself available
type ControllerServiceClient interface {
	// Exporter registration
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Exporter disconnection
	// Disconnecting with bye will invalidate any existing router tokens
	// we will eventually have a mechanism to tell the router this token
	// has been invalidated
	Bye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Exporter listening
	// Returns stream tokens for accepting incoming client connections
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ControllerService_ListenClient, error)
	// Client connecting
	// Returns stream token for connecting to the desired exporter
	// Leases are checked before token issuance
	Dial(ctx context.Context, in *DialRequest, opts ...grpc.CallOption) (*DialResponse, error)
	// Audit events from the exporters
	// audit events are used to track the exporter's activity
	AuditStream(ctx context.Context, opts ...grpc.CallOption) (ControllerService_AuditStreamClient, error)
	// List exporters
	// Returns all exporters matching filter
	ListExporters(ctx context.Context, in *ListExportersRequest, opts ...grpc.CallOption) (*ListExportersResponse, error)
	// Get exporter
	// Get information of the exporter of the given uuid
	GetExporter(ctx context.Context, in *GetExporterRequest, opts ...grpc.CallOption) (*GetExporterResponse, error)
	// Lease exporter
	// Lease the exporter of the given uuid for a given amount of time
	LeaseExporter(ctx context.Context, in *LeaseExporterRequest, opts ...grpc.CallOption) (*LeaseExporterResponse, error)
	// Release exporter
	// Return the exporter of the given uuid early
	ReleaseExporter(ctx context.Context, in *ReleaseExporterRequest, opts ...grpc.CallOption) (*ReleaseExporterResponse, error)
}

type controllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerServiceClient(cc grpc.ClientConnInterface) ControllerServiceClient {
	return &controllerServiceClient{cc}
}

func (c *controllerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, ControllerService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) Bye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ControllerService_Bye_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ControllerService_ListenClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ControllerService_ServiceDesc.Streams[0], ControllerService_Listen_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &controllerServiceListenClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControllerService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type controllerServiceListenClient struct {
	grpc.ClientStream
}

func (x *controllerServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerServiceClient) Dial(ctx context.Context, in *DialRequest, opts ...grpc.CallOption) (*DialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DialResponse)
	err := c.cc.Invoke(ctx, ControllerService_Dial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) AuditStream(ctx context.Context, opts ...grpc.CallOption) (ControllerService_AuditStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ControllerService_ServiceDesc.Streams[1], ControllerService_AuditStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &controllerServiceAuditStreamClient{ClientStream: stream}
	return x, nil
}

type ControllerService_AuditStreamClient interface {
	Send(*AuditStreamRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type controllerServiceAuditStreamClient struct {
	grpc.ClientStream
}

func (x *controllerServiceAuditStreamClient) Send(m *AuditStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerServiceAuditStreamClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controllerServiceClient) ListExporters(ctx context.Context, in *ListExportersRequest, opts ...grpc.CallOption) (*ListExportersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExportersResponse)
	err := c.cc.Invoke(ctx, ControllerService_ListExporters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) GetExporter(ctx context.Context, in *GetExporterRequest, opts ...grpc.CallOption) (*GetExporterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExporterResponse)
	err := c.cc.Invoke(ctx, ControllerService_GetExporter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) LeaseExporter(ctx context.Context, in *LeaseExporterRequest, opts ...grpc.CallOption) (*LeaseExporterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaseExporterResponse)
	err := c.cc.Invoke(ctx, ControllerService_LeaseExporter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerServiceClient) ReleaseExporter(ctx context.Context, in *ReleaseExporterRequest, opts ...grpc.CallOption) (*ReleaseExporterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReleaseExporterResponse)
	err := c.cc.Invoke(ctx, ControllerService_ReleaseExporter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerServiceServer is the server API for ControllerService service.
// All implementations must embed UnimplementedControllerServiceServer
// for forward compatibility
//
// A service where a exporter can connect to make itself available
type ControllerServiceServer interface {
	// Exporter registration
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Exporter disconnection
	// Disconnecting with bye will invalidate any existing router tokens
	// we will eventually have a mechanism to tell the router this token
	// has been invalidated
	Bye(context.Context, *ByeRequest) (*emptypb.Empty, error)
	// Exporter listening
	// Returns stream tokens for accepting incoming client connections
	Listen(*ListenRequest, ControllerService_ListenServer) error
	// Client connecting
	// Returns stream token for connecting to the desired exporter
	// Leases are checked before token issuance
	Dial(context.Context, *DialRequest) (*DialResponse, error)
	// Audit events from the exporters
	// audit events are used to track the exporter's activity
	AuditStream(ControllerService_AuditStreamServer) error
	// List exporters
	// Returns all exporters matching filter
	ListExporters(context.Context, *ListExportersRequest) (*ListExportersResponse, error)
	// Get exporter
	// Get information of the exporter of the given uuid
	GetExporter(context.Context, *GetExporterRequest) (*GetExporterResponse, error)
	// Lease exporter
	// Lease the exporter of the given uuid for a given amount of time
	LeaseExporter(context.Context, *LeaseExporterRequest) (*LeaseExporterResponse, error)
	// Release exporter
	// Return the exporter of the given uuid early
	ReleaseExporter(context.Context, *ReleaseExporterRequest) (*ReleaseExporterResponse, error)
	mustEmbedUnimplementedControllerServiceServer()
}

// UnimplementedControllerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedControllerServiceServer struct {
}

func (UnimplementedControllerServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedControllerServiceServer) Bye(context.Context, *ByeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bye not implemented")
}
func (UnimplementedControllerServiceServer) Listen(*ListenRequest, ControllerService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedControllerServiceServer) Dial(context.Context, *DialRequest) (*DialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dial not implemented")
}
func (UnimplementedControllerServiceServer) AuditStream(ControllerService_AuditStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AuditStream not implemented")
}
func (UnimplementedControllerServiceServer) ListExporters(context.Context, *ListExportersRequest) (*ListExportersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExporters not implemented")
}
func (UnimplementedControllerServiceServer) GetExporter(context.Context, *GetExporterRequest) (*GetExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExporter not implemented")
}
func (UnimplementedControllerServiceServer) LeaseExporter(context.Context, *LeaseExporterRequest) (*LeaseExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseExporter not implemented")
}
func (UnimplementedControllerServiceServer) ReleaseExporter(context.Context, *ReleaseExporterRequest) (*ReleaseExporterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseExporter not implemented")
}
func (UnimplementedControllerServiceServer) mustEmbedUnimplementedControllerServiceServer() {}

// UnsafeControllerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServiceServer will
// result in compilation errors.
type UnsafeControllerServiceServer interface {
	mustEmbedUnimplementedControllerServiceServer()
}

func RegisterControllerServiceServer(s grpc.ServiceRegistrar, srv ControllerServiceServer) {
	s.RegisterService(&ControllerService_ServiceDesc, srv)
}

func _ControllerService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_Bye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).Bye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_Bye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).Bye(ctx, req.(*ByeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServiceServer).Listen(m, &controllerServiceListenServer{ServerStream: stream})
}

type ControllerService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type controllerServiceListenServer struct {
	grpc.ServerStream
}

func (x *controllerServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ControllerService_Dial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).Dial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_Dial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).Dial(ctx, req.(*DialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_AuditStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServiceServer).AuditStream(&controllerServiceAuditStreamServer{ServerStream: stream})
}

type ControllerService_AuditStreamServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AuditStreamRequest, error)
	grpc.ServerStream
}

type controllerServiceAuditStreamServer struct {
	grpc.ServerStream
}

func (x *controllerServiceAuditStreamServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerServiceAuditStreamServer) Recv() (*AuditStreamRequest, error) {
	m := new(AuditStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ControllerService_ListExporters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExportersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).ListExporters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_ListExporters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).ListExporters(ctx, req.(*ListExportersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_GetExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).GetExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_GetExporter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).GetExporter(ctx, req.(*GetExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_LeaseExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).LeaseExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_LeaseExporter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).LeaseExporter(ctx, req.(*LeaseExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControllerService_ReleaseExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServiceServer).ReleaseExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControllerService_ReleaseExporter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServiceServer).ReleaseExporter(ctx, req.(*ReleaseExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControllerService_ServiceDesc is the grpc.ServiceDesc for ControllerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControllerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jumpstarter.v1.ControllerService",
	HandlerType: (*ControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ControllerService_Register_Handler,
		},
		{
			MethodName: "Bye",
			Handler:    _ControllerService_Bye_Handler,
		},
		{
			MethodName: "Dial",
			Handler:    _ControllerService_Dial_Handler,
		},
		{
			MethodName: "ListExporters",
			Handler:    _ControllerService_ListExporters_Handler,
		},
		{
			MethodName: "GetExporter",
			Handler:    _ControllerService_GetExporter_Handler,
		},
		{
			MethodName: "LeaseExporter",
			Handler:    _ControllerService_LeaseExporter_Handler,
		},
		{
			MethodName: "ReleaseExporter",
			Handler:    _ControllerService_ReleaseExporter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _ControllerService_Listen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AuditStream",
			Handler:       _ControllerService_AuditStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "jumpstarter/v1/jumpstarter.proto",
}

const (
	ExporterService_GetReport_FullMethodName           = "/jumpstarter.v1.ExporterService/GetReport"
	ExporterService_DriverCall_FullMethodName          = "/jumpstarter.v1.ExporterService/DriverCall"
	ExporterService_StreamingDriverCall_FullMethodName = "/jumpstarter.v1.ExporterService/StreamingDriverCall"
	ExporterService_LogStream_FullMethodName           = "/jumpstarter.v1.ExporterService/LogStream"
)

// ExporterServiceClient is the client API for ExporterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service a exporter can share locally to be used without a server
// Channel/Call credentials are used to authenticate the client, and routing to the right exporter
type ExporterServiceClient interface {
	// Exporter registration
	GetReport(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReportResponse, error)
	DriverCall(ctx context.Context, in *DriverCallRequest, opts ...grpc.CallOption) (*DriverCallResponse, error)
	StreamingDriverCall(ctx context.Context, in *StreamingDriverCallRequest, opts ...grpc.CallOption) (ExporterService_StreamingDriverCallClient, error)
	LogStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ExporterService_LogStreamClient, error)
}

type exporterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExporterServiceClient(cc grpc.ClientConnInterface) ExporterServiceClient {
	return &exporterServiceClient{cc}
}

func (c *exporterServiceClient) GetReport(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, ExporterService_GetReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exporterServiceClient) DriverCall(ctx context.Context, in *DriverCallRequest, opts ...grpc.CallOption) (*DriverCallResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DriverCallResponse)
	err := c.cc.Invoke(ctx, ExporterService_DriverCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exporterServiceClient) StreamingDriverCall(ctx context.Context, in *StreamingDriverCallRequest, opts ...grpc.CallOption) (ExporterService_StreamingDriverCallClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExporterService_ServiceDesc.Streams[0], ExporterService_StreamingDriverCall_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &exporterServiceStreamingDriverCallClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExporterService_StreamingDriverCallClient interface {
	Recv() (*StreamingDriverCallResponse, error)
	grpc.ClientStream
}

type exporterServiceStreamingDriverCallClient struct {
	grpc.ClientStream
}

func (x *exporterServiceStreamingDriverCallClient) Recv() (*StreamingDriverCallResponse, error) {
	m := new(StreamingDriverCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exporterServiceClient) LogStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ExporterService_LogStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExporterService_ServiceDesc.Streams[1], ExporterService_LogStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &exporterServiceLogStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExporterService_LogStreamClient interface {
	Recv() (*LogStreamResponse, error)
	grpc.ClientStream
}

type exporterServiceLogStreamClient struct {
	grpc.ClientStream
}

func (x *exporterServiceLogStreamClient) Recv() (*LogStreamResponse, error) {
	m := new(LogStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExporterServiceServer is the server API for ExporterService service.
// All implementations must embed UnimplementedExporterServiceServer
// for forward compatibility
//
// A service a exporter can share locally to be used without a server
// Channel/Call credentials are used to authenticate the client, and routing to the right exporter
type ExporterServiceServer interface {
	// Exporter registration
	GetReport(context.Context, *emptypb.Empty) (*GetReportResponse, error)
	DriverCall(context.Context, *DriverCallRequest) (*DriverCallResponse, error)
	StreamingDriverCall(*StreamingDriverCallRequest, ExporterService_StreamingDriverCallServer) error
	LogStream(*emptypb.Empty, ExporterService_LogStreamServer) error
	mustEmbedUnimplementedExporterServiceServer()
}

// UnimplementedExporterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExporterServiceServer struct {
}

func (UnimplementedExporterServiceServer) GetReport(context.Context, *emptypb.Empty) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (UnimplementedExporterServiceServer) DriverCall(context.Context, *DriverCallRequest) (*DriverCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverCall not implemented")
}
func (UnimplementedExporterServiceServer) StreamingDriverCall(*StreamingDriverCallRequest, ExporterService_StreamingDriverCallServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingDriverCall not implemented")
}
func (UnimplementedExporterServiceServer) LogStream(*emptypb.Empty, ExporterService_LogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method LogStream not implemented")
}
func (UnimplementedExporterServiceServer) mustEmbedUnimplementedExporterServiceServer() {}

// UnsafeExporterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExporterServiceServer will
// result in compilation errors.
type UnsafeExporterServiceServer interface {
	mustEmbedUnimplementedExporterServiceServer()
}

func RegisterExporterServiceServer(s grpc.ServiceRegistrar, srv ExporterServiceServer) {
	s.RegisterService(&ExporterService_ServiceDesc, srv)
}

func _ExporterService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExporterService_GetReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServiceServer).GetReport(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExporterService_DriverCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServiceServer).DriverCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExporterService_DriverCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServiceServer).DriverCall(ctx, req.(*DriverCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExporterService_StreamingDriverCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingDriverCallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExporterServiceServer).StreamingDriverCall(m, &exporterServiceStreamingDriverCallServer{ServerStream: stream})
}

type ExporterService_StreamingDriverCallServer interface {
	Send(*StreamingDriverCallResponse) error
	grpc.ServerStream
}

type exporterServiceStreamingDriverCallServer struct {
	grpc.ServerStream
}

func (x *exporterServiceStreamingDriverCallServer) Send(m *StreamingDriverCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExporterService_LogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExporterServiceServer).LogStream(m, &exporterServiceLogStreamServer{ServerStream: stream})
}

type ExporterService_LogStreamServer interface {
	Send(*LogStreamResponse) error
	grpc.ServerStream
}

type exporterServiceLogStreamServer struct {
	grpc.ServerStream
}

func (x *exporterServiceLogStreamServer) Send(m *LogStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ExporterService_ServiceDesc is the grpc.ServiceDesc for ExporterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExporterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jumpstarter.v1.ExporterService",
	HandlerType: (*ExporterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _ExporterService_GetReport_Handler,
		},
		{
			MethodName: "DriverCall",
			Handler:    _ExporterService_DriverCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingDriverCall",
			Handler:       _ExporterService_StreamingDriverCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogStream",
			Handler:       _ExporterService_LogStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jumpstarter/v1/jumpstarter.proto",
}
