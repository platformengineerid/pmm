// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: managementpb/dbaas/pxc_clusters.proto

package dbaasv1beta1

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
	PXCClusters_GetPXCClusterCredentials_FullMethodName = "/dbaas.v1beta1.PXCClusters/GetPXCClusterCredentials"
	PXCClusters_CreatePXCCluster_FullMethodName         = "/dbaas.v1beta1.PXCClusters/CreatePXCCluster"
	PXCClusters_UpdatePXCCluster_FullMethodName         = "/dbaas.v1beta1.PXCClusters/UpdatePXCCluster"
	PXCClusters_GetPXCClusterResources_FullMethodName   = "/dbaas.v1beta1.PXCClusters/GetPXCClusterResources"
)

// PXCClustersClient is the client API for PXCClusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PXCClustersClient interface {
	// GetPXCClusterCredentials returns a PXC cluster credentials by cluster name.
	GetPXCClusterCredentials(ctx context.Context, in *GetPXCClusterCredentialsRequest, opts ...grpc.CallOption) (*GetPXCClusterCredentialsResponse, error)
	// CreatePXCCluster creates a new PXC cluster.
	CreatePXCCluster(ctx context.Context, in *CreatePXCClusterRequest, opts ...grpc.CallOption) (*CreatePXCClusterResponse, error)
	// UpdatePXCCluster updates existing PXC cluster.
	UpdatePXCCluster(ctx context.Context, in *UpdatePXCClusterRequest, opts ...grpc.CallOption) (*UpdatePXCClusterResponse, error)
	// GetPXCClusterResources returns expected resources to be consumed by the cluster.
	GetPXCClusterResources(ctx context.Context, in *GetPXCClusterResourcesRequest, opts ...grpc.CallOption) (*GetPXCClusterResourcesResponse, error)
}

type pXCClustersClient struct {
	cc grpc.ClientConnInterface
}

func NewPXCClustersClient(cc grpc.ClientConnInterface) PXCClustersClient {
	return &pXCClustersClient{cc}
}

func (c *pXCClustersClient) GetPXCClusterCredentials(ctx context.Context, in *GetPXCClusterCredentialsRequest, opts ...grpc.CallOption) (*GetPXCClusterCredentialsResponse, error) {
	out := new(GetPXCClusterCredentialsResponse)
	err := c.cc.Invoke(ctx, PXCClusters_GetPXCClusterCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pXCClustersClient) CreatePXCCluster(ctx context.Context, in *CreatePXCClusterRequest, opts ...grpc.CallOption) (*CreatePXCClusterResponse, error) {
	out := new(CreatePXCClusterResponse)
	err := c.cc.Invoke(ctx, PXCClusters_CreatePXCCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pXCClustersClient) UpdatePXCCluster(ctx context.Context, in *UpdatePXCClusterRequest, opts ...grpc.CallOption) (*UpdatePXCClusterResponse, error) {
	out := new(UpdatePXCClusterResponse)
	err := c.cc.Invoke(ctx, PXCClusters_UpdatePXCCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pXCClustersClient) GetPXCClusterResources(ctx context.Context, in *GetPXCClusterResourcesRequest, opts ...grpc.CallOption) (*GetPXCClusterResourcesResponse, error) {
	out := new(GetPXCClusterResourcesResponse)
	err := c.cc.Invoke(ctx, PXCClusters_GetPXCClusterResources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PXCClustersServer is the server API for PXCClusters service.
// All implementations must embed UnimplementedPXCClustersServer
// for forward compatibility
type PXCClustersServer interface {
	// GetPXCClusterCredentials returns a PXC cluster credentials by cluster name.
	GetPXCClusterCredentials(context.Context, *GetPXCClusterCredentialsRequest) (*GetPXCClusterCredentialsResponse, error)
	// CreatePXCCluster creates a new PXC cluster.
	CreatePXCCluster(context.Context, *CreatePXCClusterRequest) (*CreatePXCClusterResponse, error)
	// UpdatePXCCluster updates existing PXC cluster.
	UpdatePXCCluster(context.Context, *UpdatePXCClusterRequest) (*UpdatePXCClusterResponse, error)
	// GetPXCClusterResources returns expected resources to be consumed by the cluster.
	GetPXCClusterResources(context.Context, *GetPXCClusterResourcesRequest) (*GetPXCClusterResourcesResponse, error)
	mustEmbedUnimplementedPXCClustersServer()
}

// UnimplementedPXCClustersServer must be embedded to have forward compatible implementations.
type UnimplementedPXCClustersServer struct{}

func (UnimplementedPXCClustersServer) GetPXCClusterCredentials(context.Context, *GetPXCClusterCredentialsRequest) (*GetPXCClusterCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPXCClusterCredentials not implemented")
}

func (UnimplementedPXCClustersServer) CreatePXCCluster(context.Context, *CreatePXCClusterRequest) (*CreatePXCClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePXCCluster not implemented")
}

func (UnimplementedPXCClustersServer) UpdatePXCCluster(context.Context, *UpdatePXCClusterRequest) (*UpdatePXCClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePXCCluster not implemented")
}

func (UnimplementedPXCClustersServer) GetPXCClusterResources(context.Context, *GetPXCClusterResourcesRequest) (*GetPXCClusterResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPXCClusterResources not implemented")
}
func (UnimplementedPXCClustersServer) mustEmbedUnimplementedPXCClustersServer() {}

// UnsafePXCClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PXCClustersServer will
// result in compilation errors.
type UnsafePXCClustersServer interface {
	mustEmbedUnimplementedPXCClustersServer()
}

func RegisterPXCClustersServer(s grpc.ServiceRegistrar, srv PXCClustersServer) {
	s.RegisterService(&PXCClusters_ServiceDesc, srv)
}

func _PXCClusters_GetPXCClusterCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPXCClusterCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PXCClustersServer).GetPXCClusterCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PXCClusters_GetPXCClusterCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PXCClustersServer).GetPXCClusterCredentials(ctx, req.(*GetPXCClusterCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PXCClusters_CreatePXCCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePXCClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PXCClustersServer).CreatePXCCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PXCClusters_CreatePXCCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PXCClustersServer).CreatePXCCluster(ctx, req.(*CreatePXCClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PXCClusters_UpdatePXCCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePXCClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PXCClustersServer).UpdatePXCCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PXCClusters_UpdatePXCCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PXCClustersServer).UpdatePXCCluster(ctx, req.(*UpdatePXCClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PXCClusters_GetPXCClusterResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPXCClusterResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PXCClustersServer).GetPXCClusterResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PXCClusters_GetPXCClusterResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PXCClustersServer).GetPXCClusterResources(ctx, req.(*GetPXCClusterResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PXCClusters_ServiceDesc is the grpc.ServiceDesc for PXCClusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PXCClusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dbaas.v1beta1.PXCClusters",
	HandlerType: (*PXCClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPXCClusterCredentials",
			Handler:    _PXCClusters_GetPXCClusterCredentials_Handler,
		},
		{
			MethodName: "CreatePXCCluster",
			Handler:    _PXCClusters_CreatePXCCluster_Handler,
		},
		{
			MethodName: "UpdatePXCCluster",
			Handler:    _PXCClusters_UpdatePXCCluster_Handler,
		},
		{
			MethodName: "GetPXCClusterResources",
			Handler:    _PXCClusters_GetPXCClusterResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/dbaas/pxc_clusters.proto",
}
