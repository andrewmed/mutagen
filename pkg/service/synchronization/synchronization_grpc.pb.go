// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package synchronization

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

// SynchronizationClient is the client API for Synchronization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SynchronizationClient interface {
	// Create creates a new session.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Flush flushes sessions.
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
	// Pause pauses sessions.
	Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error)
	// Reset resets sessions' histories.
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	// Terminate terminates sessions.
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type synchronizationClient struct {
	cc grpc.ClientConnInterface
}

func NewSynchronizationClient(cc grpc.ClientConnInterface) SynchronizationClient {
	return &synchronizationClient{cc}
}

func (c *synchronizationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error) {
	out := new(PauseResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error) {
	out := new(ResumeResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizationClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/synchronization.Synchronization/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SynchronizationServer is the server API for Synchronization service.
// All implementations must embed UnimplementedSynchronizationServer
// for forward compatibility
type SynchronizationServer interface {
	// Create creates a new session.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Flush flushes sessions.
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	// Pause pauses sessions.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)
	// Reset resets sessions' histories.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	// Terminate terminates sessions.
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	mustEmbedUnimplementedSynchronizationServer()
}

// UnimplementedSynchronizationServer must be embedded to have forward compatible implementations.
type UnimplementedSynchronizationServer struct {
}

func (UnimplementedSynchronizationServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSynchronizationServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSynchronizationServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedSynchronizationServer) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedSynchronizationServer) Resume(context.Context, *ResumeRequest) (*ResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedSynchronizationServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedSynchronizationServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (UnimplementedSynchronizationServer) mustEmbedUnimplementedSynchronizationServer() {}

// UnsafeSynchronizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SynchronizationServer will
// result in compilation errors.
type UnsafeSynchronizationServer interface {
	mustEmbedUnimplementedSynchronizationServer()
}

func RegisterSynchronizationServer(s grpc.ServiceRegistrar, srv SynchronizationServer) {
	s.RegisterService(&Synchronization_ServiceDesc, srv)
}

func _Synchronization_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Pause(ctx, req.(*PauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Resume(ctx, req.(*ResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronization_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizationServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synchronization.Synchronization/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizationServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Synchronization_ServiceDesc is the grpc.ServiceDesc for Synchronization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Synchronization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "synchronization.Synchronization",
	HandlerType: (*SynchronizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Synchronization_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Synchronization_List_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _Synchronization_Flush_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Synchronization_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _Synchronization_Resume_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Synchronization_Reset_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _Synchronization_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/synchronization/synchronization.proto",
}