// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: managementpb/agent/agent.proto

package agentv1beta1

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
	AgentService_ListAgents_FullMethodName = "/agent.v1beta1.AgentService/ListAgents"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	// ListAgents returns a list of Agents filtered by service_id.
	ListAgents(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) ListAgents(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentResponse, error) {
	out := new(ListAgentResponse)
	err := c.cc.Invoke(ctx, AgentService_ListAgents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	// ListAgents returns a list of Agents filtered by service_id.
	ListAgents(context.Context, *ListAgentRequest) (*ListAgentResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct{}

func (UnimplementedAgentServiceServer) ListAgents(context.Context, *ListAgentRequest) (*ListAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_ListAgents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).ListAgents(ctx, req.(*ListAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.v1beta1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAgents",
			Handler:    _AgentService_ListAgents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/agent/agent.proto",
}
