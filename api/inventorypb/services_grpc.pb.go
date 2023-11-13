// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: inventorypb/services.proto

package inventorypb

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
	ServicesService_ListServices_FullMethodName           = "/inventory.ServicesService/ListServices"
	ServicesService_ListActiveServiceTypes_FullMethodName = "/inventory.ServicesService/ListActiveServiceTypes"
	ServicesService_GetService_FullMethodName             = "/inventory.ServicesService/GetService"
	ServicesService_AddMySQLService_FullMethodName        = "/inventory.ServicesService/AddMySQLService"
	ServicesService_AddMongoDBService_FullMethodName      = "/inventory.ServicesService/AddMongoDBService"
	ServicesService_AddPostgreSQLService_FullMethodName   = "/inventory.ServicesService/AddPostgreSQLService"
	ServicesService_AddProxySQLService_FullMethodName     = "/inventory.ServicesService/AddProxySQLService"
	ServicesService_AddHAProxyService_FullMethodName      = "/inventory.ServicesService/AddHAProxyService"
	ServicesService_AddExternalService_FullMethodName     = "/inventory.ServicesService/AddExternalService"
	ServicesService_RemoveService_FullMethodName          = "/inventory.ServicesService/RemoveService"
	ServicesService_AddCustomLabels_FullMethodName        = "/inventory.ServicesService/AddCustomLabels"
	ServicesService_RemoveCustomLabels_FullMethodName     = "/inventory.ServicesService/RemoveCustomLabels"
	ServicesService_ChangeService_FullMethodName          = "/inventory.ServicesService/ChangeService"
)

// ServicesServiceClient is the client API for ServicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesServiceClient interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// ListActiveServiceTypes returns a list of active Services.
	ListActiveServiceTypes(ctx context.Context, in *ListActiveServiceTypesRequest, opts ...grpc.CallOption) (*ListActiveServiceTypesResponse, error)
	// GetService returns a single Service by ID.
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error)
	// AddMongoDBService adds MongoDB Service.
	AddMongoDBService(ctx context.Context, in *AddMongoDBServiceRequest, opts ...grpc.CallOption) (*AddMongoDBServiceResponse, error)
	// AddPostgreSQLService adds PostgreSQL Service.
	AddPostgreSQLService(ctx context.Context, in *AddPostgreSQLServiceRequest, opts ...grpc.CallOption) (*AddPostgreSQLServiceResponse, error)
	// AddProxySQLService adds ProxySQL Service.
	AddProxySQLService(ctx context.Context, in *AddProxySQLServiceRequest, opts ...grpc.CallOption) (*AddProxySQLServiceResponse, error)
	// AddHAProxyService adds HAProxy Service.
	AddHAProxyService(ctx context.Context, in *AddHAProxyServiceRequest, opts ...grpc.CallOption) (*AddHAProxyServiceResponse, error)
	// AddExternalService adds External Service.
	AddExternalService(ctx context.Context, in *AddExternalServiceRequest, opts ...grpc.CallOption) (*AddExternalServiceResponse, error)
	// RemoveService removes Service.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
	// AddCustomLabels adds custom labels to a Service.
	AddCustomLabels(ctx context.Context, in *AddCustomLabelsRequest, opts ...grpc.CallOption) (*AddCustomLabelsResponse, error)
	// RemoveCustomLabels removes custom labels from a Service.
	RemoveCustomLabels(ctx context.Context, in *RemoveCustomLabelsRequest, opts ...grpc.CallOption) (*RemoveCustomLabelsResponse, error)
	// ChangeService allows changing configuration of a service.
	ChangeService(ctx context.Context, in *ChangeServiceRequest, opts ...grpc.CallOption) (*ChangeServiceResponse, error)
}

type servicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesServiceClient(cc grpc.ClientConnInterface) ServicesServiceClient {
	return &servicesServiceClient{cc}
}

func (c *servicesServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ServicesService_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) ListActiveServiceTypes(ctx context.Context, in *ListActiveServiceTypesRequest, opts ...grpc.CallOption) (*ListActiveServiceTypesResponse, error) {
	out := new(ListActiveServiceTypesResponse)
	err := c.cc.Invoke(ctx, ServicesService_ListActiveServiceTypes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_GetService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error) {
	out := new(AddMySQLServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddMySQLService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddMongoDBService(ctx context.Context, in *AddMongoDBServiceRequest, opts ...grpc.CallOption) (*AddMongoDBServiceResponse, error) {
	out := new(AddMongoDBServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddMongoDBService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddPostgreSQLService(ctx context.Context, in *AddPostgreSQLServiceRequest, opts ...grpc.CallOption) (*AddPostgreSQLServiceResponse, error) {
	out := new(AddPostgreSQLServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddPostgreSQLService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddProxySQLService(ctx context.Context, in *AddProxySQLServiceRequest, opts ...grpc.CallOption) (*AddProxySQLServiceResponse, error) {
	out := new(AddProxySQLServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddProxySQLService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddHAProxyService(ctx context.Context, in *AddHAProxyServiceRequest, opts ...grpc.CallOption) (*AddHAProxyServiceResponse, error) {
	out := new(AddHAProxyServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddHAProxyService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddExternalService(ctx context.Context, in *AddExternalServiceRequest, opts ...grpc.CallOption) (*AddExternalServiceResponse, error) {
	out := new(AddExternalServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddExternalService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_RemoveService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) AddCustomLabels(ctx context.Context, in *AddCustomLabelsRequest, opts ...grpc.CallOption) (*AddCustomLabelsResponse, error) {
	out := new(AddCustomLabelsResponse)
	err := c.cc.Invoke(ctx, ServicesService_AddCustomLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) RemoveCustomLabels(ctx context.Context, in *RemoveCustomLabelsRequest, opts ...grpc.CallOption) (*RemoveCustomLabelsResponse, error) {
	out := new(RemoveCustomLabelsResponse)
	err := c.cc.Invoke(ctx, ServicesService_RemoveCustomLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesServiceClient) ChangeService(ctx context.Context, in *ChangeServiceRequest, opts ...grpc.CallOption) (*ChangeServiceResponse, error) {
	out := new(ChangeServiceResponse)
	err := c.cc.Invoke(ctx, ServicesService_ChangeService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServiceServer is the server API for ServicesService service.
// All implementations must embed UnimplementedServicesServiceServer
// for forward compatibility
type ServicesServiceServer interface {
	// ListServices returns a list of Services filtered by type.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// ListActiveServiceTypes returns a list of active Services.
	ListActiveServiceTypes(context.Context, *ListActiveServiceTypesRequest) (*ListActiveServiceTypesResponse, error)
	// GetService returns a single Service by ID.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error)
	// AddMongoDBService adds MongoDB Service.
	AddMongoDBService(context.Context, *AddMongoDBServiceRequest) (*AddMongoDBServiceResponse, error)
	// AddPostgreSQLService adds PostgreSQL Service.
	AddPostgreSQLService(context.Context, *AddPostgreSQLServiceRequest) (*AddPostgreSQLServiceResponse, error)
	// AddProxySQLService adds ProxySQL Service.
	AddProxySQLService(context.Context, *AddProxySQLServiceRequest) (*AddProxySQLServiceResponse, error)
	// AddHAProxyService adds HAProxy Service.
	AddHAProxyService(context.Context, *AddHAProxyServiceRequest) (*AddHAProxyServiceResponse, error)
	// AddExternalService adds External Service.
	AddExternalService(context.Context, *AddExternalServiceRequest) (*AddExternalServiceResponse, error)
	// RemoveService removes Service.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
	// AddCustomLabels adds custom labels to a Service.
	AddCustomLabels(context.Context, *AddCustomLabelsRequest) (*AddCustomLabelsResponse, error)
	// RemoveCustomLabels removes custom labels from a Service.
	RemoveCustomLabels(context.Context, *RemoveCustomLabelsRequest) (*RemoveCustomLabelsResponse, error)
	// ChangeService allows changing configuration of a service.
	ChangeService(context.Context, *ChangeServiceRequest) (*ChangeServiceResponse, error)
	mustEmbedUnimplementedServicesServiceServer()
}

// UnimplementedServicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServiceServer struct{}

func (UnimplementedServicesServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}

func (UnimplementedServicesServiceServer) ListActiveServiceTypes(context.Context, *ListActiveServiceTypesRequest) (*ListActiveServiceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActiveServiceTypes not implemented")
}

func (UnimplementedServicesServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}

func (UnimplementedServicesServiceServer) AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMySQLService not implemented")
}

func (UnimplementedServicesServiceServer) AddMongoDBService(context.Context, *AddMongoDBServiceRequest) (*AddMongoDBServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMongoDBService not implemented")
}

func (UnimplementedServicesServiceServer) AddPostgreSQLService(context.Context, *AddPostgreSQLServiceRequest) (*AddPostgreSQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostgreSQLService not implemented")
}

func (UnimplementedServicesServiceServer) AddProxySQLService(context.Context, *AddProxySQLServiceRequest) (*AddProxySQLServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProxySQLService not implemented")
}

func (UnimplementedServicesServiceServer) AddHAProxyService(context.Context, *AddHAProxyServiceRequest) (*AddHAProxyServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHAProxyService not implemented")
}

func (UnimplementedServicesServiceServer) AddExternalService(context.Context, *AddExternalServiceRequest) (*AddExternalServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternalService not implemented")
}

func (UnimplementedServicesServiceServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}

func (UnimplementedServicesServiceServer) AddCustomLabels(context.Context, *AddCustomLabelsRequest) (*AddCustomLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomLabels not implemented")
}

func (UnimplementedServicesServiceServer) RemoveCustomLabels(context.Context, *RemoveCustomLabelsRequest) (*RemoveCustomLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustomLabels not implemented")
}

func (UnimplementedServicesServiceServer) ChangeService(context.Context, *ChangeServiceRequest) (*ChangeServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeService not implemented")
}
func (UnimplementedServicesServiceServer) mustEmbedUnimplementedServicesServiceServer() {}

// UnsafeServicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServiceServer will
// result in compilation errors.
type UnsafeServicesServiceServer interface {
	mustEmbedUnimplementedServicesServiceServer()
}

func RegisterServicesServiceServer(s grpc.ServiceRegistrar, srv ServicesServiceServer) {
	s.RegisterService(&ServicesService_ServiceDesc, srv)
}

func _ServicesService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_ListActiveServiceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActiveServiceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ListActiveServiceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ListActiveServiceTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ListActiveServiceTypes(ctx, req.(*ListActiveServiceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddMySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddMySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddMySQLService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddMySQLService(ctx, req.(*AddMySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddMongoDBService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddMongoDBService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddMongoDBService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddMongoDBService(ctx, req.(*AddMongoDBServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddPostgreSQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostgreSQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddPostgreSQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddPostgreSQLService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddPostgreSQLService(ctx, req.(*AddPostgreSQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddProxySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProxySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddProxySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddProxySQLService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddProxySQLService(ctx, req.(*AddProxySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddHAProxyService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHAProxyServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddHAProxyService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddHAProxyService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddHAProxyService(ctx, req.(*AddHAProxyServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddExternalService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExternalServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddExternalService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddExternalService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddExternalService(ctx, req.(*AddExternalServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_RemoveService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_AddCustomLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).AddCustomLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_AddCustomLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).AddCustomLabels(ctx, req.(*AddCustomLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_RemoveCustomLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustomLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).RemoveCustomLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_RemoveCustomLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).RemoveCustomLabels(ctx, req.(*RemoveCustomLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesService_ChangeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServiceServer).ChangeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesService_ChangeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServiceServer).ChangeService(ctx, req.(*ChangeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServicesService_ServiceDesc is the grpc.ServiceDesc for ServicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.ServicesService",
	HandlerType: (*ServicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _ServicesService_ListServices_Handler,
		},
		{
			MethodName: "ListActiveServiceTypes",
			Handler:    _ServicesService_ListActiveServiceTypes_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _ServicesService_GetService_Handler,
		},
		{
			MethodName: "AddMySQLService",
			Handler:    _ServicesService_AddMySQLService_Handler,
		},
		{
			MethodName: "AddMongoDBService",
			Handler:    _ServicesService_AddMongoDBService_Handler,
		},
		{
			MethodName: "AddPostgreSQLService",
			Handler:    _ServicesService_AddPostgreSQLService_Handler,
		},
		{
			MethodName: "AddProxySQLService",
			Handler:    _ServicesService_AddProxySQLService_Handler,
		},
		{
			MethodName: "AddHAProxyService",
			Handler:    _ServicesService_AddHAProxyService_Handler,
		},
		{
			MethodName: "AddExternalService",
			Handler:    _ServicesService_AddExternalService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _ServicesService_RemoveService_Handler,
		},
		{
			MethodName: "AddCustomLabels",
			Handler:    _ServicesService_AddCustomLabels_Handler,
		},
		{
			MethodName: "RemoveCustomLabels",
			Handler:    _ServicesService_RemoveCustomLabels_Handler,
		},
		{
			MethodName: "ChangeService",
			Handler:    _ServicesService_ChangeService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventorypb/services.proto",
}
