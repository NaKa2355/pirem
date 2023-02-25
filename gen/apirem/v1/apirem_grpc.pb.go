// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: apirem/v1/apirem.proto

package v1

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

// PiRemServiceClient is the client API for PiRemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PiRemServiceClient interface {
	SendRawIr(ctx context.Context, in *SendRawIrRequest, opts ...grpc.CallOption) (*SendRawIrResponse, error)
	ReceiveRawIr(ctx context.Context, in *ReceiveRawIrRequest, opts ...grpc.CallOption) (*ReceiveRawIrResponse, error)
	GetAllDeviceInfo(ctx context.Context, in *GetAllDeviceInfoRequest, opts ...grpc.CallOption) (*GetAllDeviceInfoResponse, error)
	GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error)
	GetDeviceStatus(ctx context.Context, in *GetDeviceStatusRequest, opts ...grpc.CallOption) (*GetDeviceStatusResponse, error)
	IsBusy(ctx context.Context, in *IsBusyRequest, opts ...grpc.CallOption) (*IsBusyResponse, error)
}

type piRemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPiRemServiceClient(cc grpc.ClientConnInterface) PiRemServiceClient {
	return &piRemServiceClient{cc}
}

func (c *piRemServiceClient) SendRawIr(ctx context.Context, in *SendRawIrRequest, opts ...grpc.CallOption) (*SendRawIrResponse, error) {
	out := new(SendRawIrResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/SendRawIr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piRemServiceClient) ReceiveRawIr(ctx context.Context, in *ReceiveRawIrRequest, opts ...grpc.CallOption) (*ReceiveRawIrResponse, error) {
	out := new(ReceiveRawIrResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/ReceiveRawIr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piRemServiceClient) GetAllDeviceInfo(ctx context.Context, in *GetAllDeviceInfoRequest, opts ...grpc.CallOption) (*GetAllDeviceInfoResponse, error) {
	out := new(GetAllDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/GetAllDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piRemServiceClient) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoRequest, opts ...grpc.CallOption) (*GetDeviceInfoResponse, error) {
	out := new(GetDeviceInfoResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piRemServiceClient) GetDeviceStatus(ctx context.Context, in *GetDeviceStatusRequest, opts ...grpc.CallOption) (*GetDeviceStatusResponse, error) {
	out := new(GetDeviceStatusResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/GetDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piRemServiceClient) IsBusy(ctx context.Context, in *IsBusyRequest, opts ...grpc.CallOption) (*IsBusyResponse, error) {
	out := new(IsBusyResponse)
	err := c.cc.Invoke(ctx, "/apirem.PiRemService/IsBusy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PiRemServiceServer is the server API for PiRemService service.
// All implementations must embed UnimplementedPiRemServiceServer
// for forward compatibility
type PiRemServiceServer interface {
	SendRawIr(context.Context, *SendRawIrRequest) (*SendRawIrResponse, error)
	ReceiveRawIr(context.Context, *ReceiveRawIrRequest) (*ReceiveRawIrResponse, error)
	GetAllDeviceInfo(context.Context, *GetAllDeviceInfoRequest) (*GetAllDeviceInfoResponse, error)
	GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error)
	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)
	IsBusy(context.Context, *IsBusyRequest) (*IsBusyResponse, error)
	mustEmbedUnimplementedPiRemServiceServer()
}

// UnimplementedPiRemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPiRemServiceServer struct {
}

func (UnimplementedPiRemServiceServer) SendRawIr(context.Context, *SendRawIrRequest) (*SendRawIrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRawIr not implemented")
}
func (UnimplementedPiRemServiceServer) ReceiveRawIr(context.Context, *ReceiveRawIrRequest) (*ReceiveRawIrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveRawIr not implemented")
}
func (UnimplementedPiRemServiceServer) GetAllDeviceInfo(context.Context, *GetAllDeviceInfoRequest) (*GetAllDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDeviceInfo not implemented")
}
func (UnimplementedPiRemServiceServer) GetDeviceInfo(context.Context, *GetDeviceInfoRequest) (*GetDeviceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceInfo not implemented")
}
func (UnimplementedPiRemServiceServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceStatus not implemented")
}
func (UnimplementedPiRemServiceServer) IsBusy(context.Context, *IsBusyRequest) (*IsBusyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBusy not implemented")
}
func (UnimplementedPiRemServiceServer) mustEmbedUnimplementedPiRemServiceServer() {}

// UnsafePiRemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PiRemServiceServer will
// result in compilation errors.
type UnsafePiRemServiceServer interface {
	mustEmbedUnimplementedPiRemServiceServer()
}

func RegisterPiRemServiceServer(s grpc.ServiceRegistrar, srv PiRemServiceServer) {
	s.RegisterService(&PiRemService_ServiceDesc, srv)
}

func _PiRemService_SendRawIr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRawIrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).SendRawIr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/SendRawIr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).SendRawIr(ctx, req.(*SendRawIrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiRemService_ReceiveRawIr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRawIrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).ReceiveRawIr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/ReceiveRawIr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).ReceiveRawIr(ctx, req.(*ReceiveRawIrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiRemService_GetAllDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).GetAllDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/GetAllDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).GetAllDeviceInfo(ctx, req.(*GetAllDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiRemService_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).GetDeviceInfo(ctx, req.(*GetDeviceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiRemService_GetDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).GetDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/GetDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).GetDeviceStatus(ctx, req.(*GetDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiRemService_IsBusy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBusyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiRemServiceServer).IsBusy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apirem.PiRemService/IsBusy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiRemServiceServer).IsBusy(ctx, req.(*IsBusyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PiRemService_ServiceDesc is the grpc.ServiceDesc for PiRemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PiRemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apirem.PiRemService",
	HandlerType: (*PiRemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRawIr",
			Handler:    _PiRemService_SendRawIr_Handler,
		},
		{
			MethodName: "ReceiveRawIr",
			Handler:    _PiRemService_ReceiveRawIr_Handler,
		},
		{
			MethodName: "GetAllDeviceInfo",
			Handler:    _PiRemService_GetAllDeviceInfo_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _PiRemService_GetDeviceInfo_Handler,
		},
		{
			MethodName: "GetDeviceStatus",
			Handler:    _PiRemService_GetDeviceStatus_Handler,
		},
		{
			MethodName: "IsBusy",
			Handler:    _PiRemService_IsBusy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apirem/v1/apirem.proto",
}