// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: platformpb/platform.proto

package platformpbv1

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
	PlatformService_Connect_FullMethodName                        = "/platformpb.v1.PlatformService/Connect"
	PlatformService_Disconnect_FullMethodName                     = "/platformpb.v1.PlatformService/Disconnect"
	PlatformService_SearchOrganizationTickets_FullMethodName      = "/platformpb.v1.PlatformService/SearchOrganizationTickets"
	PlatformService_SearchOrganizationEntitlements_FullMethodName = "/platformpb.v1.PlatformService/SearchOrganizationEntitlements"
	PlatformService_GetContactInformation_FullMethodName          = "/platformpb.v1.PlatformService/GetContactInformation"
	PlatformService_ServerInfo_FullMethodName                     = "/platformpb.v1.PlatformService/ServerInfo"
	PlatformService_UserStatus_FullMethodName                     = "/platformpb.v1.PlatformService/UserStatus"
)

// PlatformServiceClient is the client API for PlatformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlatformServiceClient interface {
	// Connect a PMM server to the organization created on Percona Portal. That allows the user to sign in to the PMM server with their Percona Account.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// Disconnect a PMM server from the organization created on Percona Portal.
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// SearchOrganizationTickets searches support tickets belonging to the Percona Portal Organization that the PMM server is connected to.
	SearchOrganizationTickets(ctx context.Context, in *SearchOrganizationTicketsRequest, opts ...grpc.CallOption) (*SearchOrganizationTicketsResponse, error)
	// SearchOrganizationEntitlements fetches details of the entitlement's available to the Portal organization that the PMM server is connected to.
	SearchOrganizationEntitlements(ctx context.Context, in *SearchOrganizationEntitlementsRequest, opts ...grpc.CallOption) (*SearchOrganizationEntitlementsResponse, error)
	// GetContactInformation fetches the contact details of the customer success employee handling the Percona customer account from Percona Platform.
	GetContactInformation(ctx context.Context, in *GetContactInformationRequest, opts ...grpc.CallOption) (*GetContactInformationResponse, error)
	// ServerInfo returns PMM server ID and name.
	ServerInfo(ctx context.Context, in *ServerInfoRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error)
	// UserStatus returns a boolean indicating whether the current user is logged in with their Percona Account or not.
	UserStatus(ctx context.Context, in *UserStatusRequest, opts ...grpc.CallOption) (*UserStatusResponse, error)
}

type platformServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformServiceClient(cc grpc.ClientConnInterface) PlatformServiceClient {
	return &platformServiceClient{cc}
}

func (c *platformServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, PlatformService_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, PlatformService_Disconnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) SearchOrganizationTickets(ctx context.Context, in *SearchOrganizationTicketsRequest, opts ...grpc.CallOption) (*SearchOrganizationTicketsResponse, error) {
	out := new(SearchOrganizationTicketsResponse)
	err := c.cc.Invoke(ctx, PlatformService_SearchOrganizationTickets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) SearchOrganizationEntitlements(ctx context.Context, in *SearchOrganizationEntitlementsRequest, opts ...grpc.CallOption) (*SearchOrganizationEntitlementsResponse, error) {
	out := new(SearchOrganizationEntitlementsResponse)
	err := c.cc.Invoke(ctx, PlatformService_SearchOrganizationEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) GetContactInformation(ctx context.Context, in *GetContactInformationRequest, opts ...grpc.CallOption) (*GetContactInformationResponse, error) {
	out := new(GetContactInformationResponse)
	err := c.cc.Invoke(ctx, PlatformService_GetContactInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) ServerInfo(ctx context.Context, in *ServerInfoRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := c.cc.Invoke(ctx, PlatformService_ServerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformServiceClient) UserStatus(ctx context.Context, in *UserStatusRequest, opts ...grpc.CallOption) (*UserStatusResponse, error) {
	out := new(UserStatusResponse)
	err := c.cc.Invoke(ctx, PlatformService_UserStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServiceServer is the server API for PlatformService service.
// All implementations must embed UnimplementedPlatformServiceServer
// for forward compatibility
type PlatformServiceServer interface {
	// Connect a PMM server to the organization created on Percona Portal. That allows the user to sign in to the PMM server with their Percona Account.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Disconnect a PMM server from the organization created on Percona Portal.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// SearchOrganizationTickets searches support tickets belonging to the Percona Portal Organization that the PMM server is connected to.
	SearchOrganizationTickets(context.Context, *SearchOrganizationTicketsRequest) (*SearchOrganizationTicketsResponse, error)
	// SearchOrganizationEntitlements fetches details of the entitlement's available to the Portal organization that the PMM server is connected to.
	SearchOrganizationEntitlements(context.Context, *SearchOrganizationEntitlementsRequest) (*SearchOrganizationEntitlementsResponse, error)
	// GetContactInformation fetches the contact details of the customer success employee handling the Percona customer account from Percona Platform.
	GetContactInformation(context.Context, *GetContactInformationRequest) (*GetContactInformationResponse, error)
	// ServerInfo returns PMM server ID and name.
	ServerInfo(context.Context, *ServerInfoRequest) (*ServerInfoResponse, error)
	// UserStatus returns a boolean indicating whether the current user is logged in with their Percona Account or not.
	UserStatus(context.Context, *UserStatusRequest) (*UserStatusResponse, error)
	mustEmbedUnimplementedPlatformServiceServer()
}

// UnimplementedPlatformServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlatformServiceServer struct{}

func (UnimplementedPlatformServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func (UnimplementedPlatformServiceServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func (UnimplementedPlatformServiceServer) SearchOrganizationTickets(context.Context, *SearchOrganizationTicketsRequest) (*SearchOrganizationTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOrganizationTickets not implemented")
}

func (UnimplementedPlatformServiceServer) SearchOrganizationEntitlements(context.Context, *SearchOrganizationEntitlementsRequest) (*SearchOrganizationEntitlementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOrganizationEntitlements not implemented")
}

func (UnimplementedPlatformServiceServer) GetContactInformation(context.Context, *GetContactInformationRequest) (*GetContactInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactInformation not implemented")
}

func (UnimplementedPlatformServiceServer) ServerInfo(context.Context, *ServerInfoRequest) (*ServerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerInfo not implemented")
}

func (UnimplementedPlatformServiceServer) UserStatus(context.Context, *UserStatusRequest) (*UserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserStatus not implemented")
}
func (UnimplementedPlatformServiceServer) mustEmbedUnimplementedPlatformServiceServer() {}

// UnsafePlatformServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformServiceServer will
// result in compilation errors.
type UnsafePlatformServiceServer interface {
	mustEmbedUnimplementedPlatformServiceServer()
}

func RegisterPlatformServiceServer(s grpc.ServiceRegistrar, srv PlatformServiceServer) {
	s.RegisterService(&PlatformService_ServiceDesc, srv)
}

func _PlatformService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_SearchOrganizationTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOrganizationTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).SearchOrganizationTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_SearchOrganizationTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).SearchOrganizationTickets(ctx, req.(*SearchOrganizationTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_SearchOrganizationEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOrganizationEntitlementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).SearchOrganizationEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_SearchOrganizationEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).SearchOrganizationEntitlements(ctx, req.(*SearchOrganizationEntitlementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_GetContactInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).GetContactInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_GetContactInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).GetContactInformation(ctx, req.(*GetContactInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_ServerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).ServerInfo(ctx, req.(*ServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformService_UserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServiceServer).UserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlatformService_UserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServiceServer).UserStatus(ctx, req.(*UserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlatformService_ServiceDesc is the grpc.ServiceDesc for PlatformService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlatformService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platformpb.v1.PlatformService",
	HandlerType: (*PlatformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _PlatformService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _PlatformService_Disconnect_Handler,
		},
		{
			MethodName: "SearchOrganizationTickets",
			Handler:    _PlatformService_SearchOrganizationTickets_Handler,
		},
		{
			MethodName: "SearchOrganizationEntitlements",
			Handler:    _PlatformService_SearchOrganizationEntitlements_Handler,
		},
		{
			MethodName: "GetContactInformation",
			Handler:    _PlatformService_GetContactInformation_Handler,
		},
		{
			MethodName: "ServerInfo",
			Handler:    _PlatformService_ServerInfo_Handler,
		},
		{
			MethodName: "UserStatus",
			Handler:    _PlatformService_UserStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platformpb/platform.proto",
}
