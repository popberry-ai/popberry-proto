// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: wallet/v1alpha1/wallet.proto

package walletv1alpha1

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
	WalletService_UpdateBalance_FullMethodName             = "/wallet.v1alpha1.WalletService/UpdateBalance"
	WalletService_IdentifyUser_FullMethodName              = "/wallet.v1alpha1.WalletService/IdentifyUser"
	WalletService_UpdateTransactionResponse_FullMethodName = "/wallet.v1alpha1.WalletService/UpdateTransactionResponse"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the Wallet service
type WalletServiceClient interface {
	// Method for updating wallet balance using transaction records
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error)
	// Method for identifying the existence of a user with various identification methods
	IdentifyUser(ctx context.Context, in *IdentifyUserRequest, opts ...grpc.CallOption) (*IdentifyUserResponse, error)
	// Method for updating the transaction response after another service has processed the request
	UpdateTransactionResponse(ctx context.Context, in *UpdateTransactionResponseRequest, opts ...grpc.CallOption) (*UpdateTransactionResponseResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*UpdateBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBalanceResponse)
	err := c.cc.Invoke(ctx, WalletService_UpdateBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) IdentifyUser(ctx context.Context, in *IdentifyUserRequest, opts ...grpc.CallOption) (*IdentifyUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IdentifyUserResponse)
	err := c.cc.Invoke(ctx, WalletService_IdentifyUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) UpdateTransactionResponse(ctx context.Context, in *UpdateTransactionResponseRequest, opts ...grpc.CallOption) (*UpdateTransactionResponseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTransactionResponseResponse)
	err := c.cc.Invoke(ctx, WalletService_UpdateTransactionResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility.
//
// Define the Wallet service
type WalletServiceServer interface {
	// Method for updating wallet balance using transaction records
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error)
	// Method for identifying the existence of a user with various identification methods
	IdentifyUser(context.Context, *IdentifyUserRequest) (*IdentifyUserResponse, error)
	// Method for updating the transaction response after another service has processed the request
	UpdateTransactionResponse(context.Context, *UpdateTransactionResponseRequest) (*UpdateTransactionResponseResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*UpdateBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedWalletServiceServer) IdentifyUser(context.Context, *IdentifyUserRequest) (*IdentifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentifyUser not implemented")
}
func (UnimplementedWalletServiceServer) UpdateTransactionResponse(context.Context, *UpdateTransactionResponseRequest) (*UpdateTransactionResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactionResponse not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}
func (UnimplementedWalletServiceServer) testEmbeddedByValue()                       {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_UpdateBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_IdentifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).IdentifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_IdentifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).IdentifyUser(ctx, req.(*IdentifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_UpdateTransactionResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).UpdateTransactionResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_UpdateTransactionResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).UpdateTransactionResponse(ctx, req.(*UpdateTransactionResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.v1alpha1.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBalance",
			Handler:    _WalletService_UpdateBalance_Handler,
		},
		{
			MethodName: "IdentifyUser",
			Handler:    _WalletService_IdentifyUser_Handler,
		},
		{
			MethodName: "UpdateTransactionResponse",
			Handler:    _WalletService_UpdateTransactionResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/v1alpha1/wallet.proto",
}
