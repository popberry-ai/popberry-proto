// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: bff/bridgr/config/v1alpha1/service.proto

package configv1alpha1

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
	ConfigService_RegisterDevice_FullMethodName        = "/bff.bridgr.config.v1alpha1.ConfigService/RegisterDevice"
	ConfigService_VerifyDevice_FullMethodName          = "/bff.bridgr.config.v1alpha1.ConfigService/VerifyDevice"
	ConfigService_AssignDeviceGroup_FullMethodName     = "/bff.bridgr.config.v1alpha1.ConfigService/AssignDeviceGroup"
	ConfigService_UpdateForwardingRules_FullMethodName = "/bff.bridgr.config.v1alpha1.ConfigService/UpdateForwardingRules"
	ConfigService_GetForwardingRules_FullMethodName    = "/bff.bridgr.config.v1alpha1.ConfigService/GetForwardingRules"
	ConfigService_GetConfig_FullMethodName             = "/bff.bridgr.config.v1alpha1.ConfigService/GetConfig"
	ConfigService_UpdateConfig_FullMethodName          = "/bff.bridgr.config.v1alpha1.ConfigService/UpdateConfig"
)

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error)
	VerifyDevice(ctx context.Context, in *VerifyDeviceRequest, opts ...grpc.CallOption) (*VerifyDeviceResponse, error)
	AssignDeviceGroup(ctx context.Context, in *AssignDeviceGroupRequest, opts ...grpc.CallOption) (*AssignDeviceGroupResponse, error)
	UpdateForwardingRules(ctx context.Context, in *UpdateForwardingRulesRequest, opts ...grpc.CallOption) (*UpdateForwardingRulesResponse, error)
	GetForwardingRules(ctx context.Context, in *GetForwardingRulesRequest, opts ...grpc.CallOption) (*GetForwardingRulesResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetConfigResponse], error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) RegisterDevice(ctx context.Context, in *RegisterDeviceRequest, opts ...grpc.CallOption) (*RegisterDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterDeviceResponse)
	err := c.cc.Invoke(ctx, ConfigService_RegisterDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) VerifyDevice(ctx context.Context, in *VerifyDeviceRequest, opts ...grpc.CallOption) (*VerifyDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyDeviceResponse)
	err := c.cc.Invoke(ctx, ConfigService_VerifyDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) AssignDeviceGroup(ctx context.Context, in *AssignDeviceGroupRequest, opts ...grpc.CallOption) (*AssignDeviceGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignDeviceGroupResponse)
	err := c.cc.Invoke(ctx, ConfigService_AssignDeviceGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateForwardingRules(ctx context.Context, in *UpdateForwardingRulesRequest, opts ...grpc.CallOption) (*UpdateForwardingRulesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateForwardingRulesResponse)
	err := c.cc.Invoke(ctx, ConfigService_UpdateForwardingRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetForwardingRules(ctx context.Context, in *GetForwardingRulesRequest, opts ...grpc.CallOption) (*GetForwardingRulesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetForwardingRulesResponse)
	err := c.cc.Invoke(ctx, ConfigService_GetForwardingRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetConfigResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ConfigService_ServiceDesc.Streams[0], ConfigService_GetConfig_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetConfigRequest, GetConfigResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ConfigService_GetConfigClient = grpc.ServerStreamingClient[GetConfigResponse]

func (c *configServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, ConfigService_UpdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility.
type ConfigServiceServer interface {
	RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error)
	VerifyDevice(context.Context, *VerifyDeviceRequest) (*VerifyDeviceResponse, error)
	AssignDeviceGroup(context.Context, *AssignDeviceGroupRequest) (*AssignDeviceGroupResponse, error)
	UpdateForwardingRules(context.Context, *UpdateForwardingRulesRequest) (*UpdateForwardingRulesResponse, error)
	GetForwardingRules(context.Context, *GetForwardingRulesRequest) (*GetForwardingRulesResponse, error)
	GetConfig(*GetConfigRequest, grpc.ServerStreamingServer[GetConfigResponse]) error
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigServiceServer struct{}

func (UnimplementedConfigServiceServer) RegisterDevice(context.Context, *RegisterDeviceRequest) (*RegisterDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedConfigServiceServer) VerifyDevice(context.Context, *VerifyDeviceRequest) (*VerifyDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyDevice not implemented")
}
func (UnimplementedConfigServiceServer) AssignDeviceGroup(context.Context, *AssignDeviceGroupRequest) (*AssignDeviceGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignDeviceGroup not implemented")
}
func (UnimplementedConfigServiceServer) UpdateForwardingRules(context.Context, *UpdateForwardingRulesRequest) (*UpdateForwardingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateForwardingRules not implemented")
}
func (UnimplementedConfigServiceServer) GetForwardingRules(context.Context, *GetForwardingRulesRequest) (*GetForwardingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForwardingRules not implemented")
}
func (UnimplementedConfigServiceServer) GetConfig(*GetConfigRequest, grpc.ServerStreamingServer[GetConfigResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}
func (UnimplementedConfigServiceServer) testEmbeddedByValue()                       {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	// If the following call pancis, it indicates UnimplementedConfigServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_RegisterDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).RegisterDevice(ctx, req.(*RegisterDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_VerifyDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).VerifyDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_VerifyDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).VerifyDevice(ctx, req.(*VerifyDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_AssignDeviceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignDeviceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).AssignDeviceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_AssignDeviceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).AssignDeviceGroup(ctx, req.(*AssignDeviceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateForwardingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateForwardingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateForwardingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_UpdateForwardingRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateForwardingRules(ctx, req.(*UpdateForwardingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetForwardingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForwardingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetForwardingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_GetForwardingRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetForwardingRules(ctx, req.(*GetForwardingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigServiceServer).GetConfig(m, &grpc.GenericServerStream[GetConfigRequest, GetConfigResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ConfigService_GetConfigServer = grpc.ServerStreamingServer[GetConfigResponse]

func _ConfigService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.bridgr.config.v1alpha1.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _ConfigService_RegisterDevice_Handler,
		},
		{
			MethodName: "VerifyDevice",
			Handler:    _ConfigService_VerifyDevice_Handler,
		},
		{
			MethodName: "AssignDeviceGroup",
			Handler:    _ConfigService_AssignDeviceGroup_Handler,
		},
		{
			MethodName: "UpdateForwardingRules",
			Handler:    _ConfigService_UpdateForwardingRules_Handler,
		},
		{
			MethodName: "GetForwardingRules",
			Handler:    _ConfigService_GetForwardingRules_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigService_UpdateConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfig",
			Handler:       _ConfigService_GetConfig_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bff/bridgr/config/v1alpha1/service.proto",
}
