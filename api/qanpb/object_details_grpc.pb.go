// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: qanpb/object_details.proto

package qanv1beta1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ObjectDetailsClient is the client API for ObjectDetails service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectDetailsClient interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsReply, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(ctx context.Context, in *QueryExampleRequest, opts ...grpc.CallOption) (*QueryExampleReply, error)
	// GetLabels gets list of labels for object details.
	GetLabels(ctx context.Context, in *ObjectDetailsLabelsRequest, opts ...grpc.CallOption) (*ObjectDetailsLabelsReply, error)
	// GetQueryPlan gets query plan and plan id for specific filtering.
	GetQueryPlan(ctx context.Context, in *QueryPlanRequest, opts ...grpc.CallOption) (*QueryPlanReply, error)
	// GetHistogram gets histogram items for specific filtering.
	GetHistogram(ctx context.Context, in *HistogramRequest, opts ...grpc.CallOption) (*HistogramReply, error)
	// QueryExists check if query exists in clickhouse.
	QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	// QueryByQueryID get query for given query ID.
	QueryByQueryID(ctx context.Context, in *QueryByQueryIDRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
}

type objectDetailsClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectDetailsClient(cc grpc.ClientConnInterface) ObjectDetailsClient {
	return &objectDetailsClient{cc}
}

func (c *objectDetailsClient) GetMetrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsReply, error) {
	out := new(MetricsReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetQueryExample(ctx context.Context, in *QueryExampleRequest, opts ...grpc.CallOption) (*QueryExampleReply, error) {
	out := new(QueryExampleReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetQueryExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetLabels(ctx context.Context, in *ObjectDetailsLabelsRequest, opts ...grpc.CallOption) (*ObjectDetailsLabelsReply, error) {
	out := new(ObjectDetailsLabelsReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetQueryPlan(ctx context.Context, in *QueryPlanRequest, opts ...grpc.CallOption) (*QueryPlanReply, error) {
	out := new(QueryPlanReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetQueryPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) GetHistogram(ctx context.Context, in *HistogramRequest, opts ...grpc.CallOption) (*HistogramReply, error) {
	out := new(HistogramReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/GetHistogram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) QueryExists(ctx context.Context, in *QueryExistsRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/QueryExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectDetailsClient) QueryByQueryID(ctx context.Context, in *QueryByQueryIDRequest, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.ObjectDetails/QueryByQueryID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectDetailsServer is the server API for ObjectDetails service.
// All implementations must embed UnimplementedObjectDetailsServer
// for forward compatibility
type ObjectDetailsServer interface {
	// GetMetrics gets map of metrics for specific filtering.
	GetMetrics(context.Context, *MetricsRequest) (*MetricsReply, error)
	// GetQueryExample gets list of query examples.
	GetQueryExample(context.Context, *QueryExampleRequest) (*QueryExampleReply, error)
	// GetLabels gets list of labels for object details.
	GetLabels(context.Context, *ObjectDetailsLabelsRequest) (*ObjectDetailsLabelsReply, error)
	// GetQueryPlan gets query plan and plan id for specific filtering.
	GetQueryPlan(context.Context, *QueryPlanRequest) (*QueryPlanReply, error)
	// GetHistogram gets histogram items for specific filtering.
	GetHistogram(context.Context, *HistogramRequest) (*HistogramReply, error)
	// QueryExists check if query exists in clickhouse.
	QueryExists(context.Context, *QueryExistsRequest) (*wrapperspb.BoolValue, error)
	// QueryByQueryID get query for given query ID.
	QueryByQueryID(context.Context, *QueryByQueryIDRequest) (*wrapperspb.StringValue, error)
	mustEmbedUnimplementedObjectDetailsServer()
}

// UnimplementedObjectDetailsServer must be embedded to have forward compatible implementations.
type UnimplementedObjectDetailsServer struct{}

func (UnimplementedObjectDetailsServer) GetMetrics(context.Context, *MetricsRequest) (*MetricsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}

func (UnimplementedObjectDetailsServer) GetQueryExample(context.Context, *QueryExampleRequest) (*QueryExampleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryExample not implemented")
}

func (UnimplementedObjectDetailsServer) GetLabels(context.Context, *ObjectDetailsLabelsRequest) (*ObjectDetailsLabelsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}

func (UnimplementedObjectDetailsServer) GetQueryPlan(context.Context, *QueryPlanRequest) (*QueryPlanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryPlan not implemented")
}

func (UnimplementedObjectDetailsServer) GetHistogram(context.Context, *HistogramRequest) (*HistogramReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistogram not implemented")
}

func (UnimplementedObjectDetailsServer) QueryExists(context.Context, *QueryExistsRequest) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryExists not implemented")
}

func (UnimplementedObjectDetailsServer) QueryByQueryID(context.Context, *QueryByQueryIDRequest) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryByQueryID not implemented")
}
func (UnimplementedObjectDetailsServer) mustEmbedUnimplementedObjectDetailsServer() {}

// UnsafeObjectDetailsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectDetailsServer will
// result in compilation errors.
type UnsafeObjectDetailsServer interface {
	mustEmbedUnimplementedObjectDetailsServer()
}

func RegisterObjectDetailsServer(s grpc.ServiceRegistrar, srv ObjectDetailsServer) {
	s.RegisterService(&ObjectDetails_ServiceDesc, srv)
}

func _ObjectDetails_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetMetrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetQueryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetQueryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetQueryExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetQueryExample(ctx, req.(*QueryExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectDetailsLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetLabels(ctx, req.(*ObjectDetailsLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetQueryPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetQueryPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetQueryPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetQueryPlan(ctx, req.(*QueryPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_GetHistogram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistogramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).GetHistogram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/GetHistogram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).GetHistogram(ctx, req.(*HistogramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_QueryExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).QueryExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/QueryExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).QueryExists(ctx, req.(*QueryExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectDetails_QueryByQueryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByQueryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetailsServer).QueryByQueryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.ObjectDetails/QueryByQueryID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetailsServer).QueryByQueryID(ctx, req.(*QueryByQueryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectDetails_ServiceDesc is the grpc.ServiceDesc for ObjectDetails service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectDetails_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.ObjectDetails",
	HandlerType: (*ObjectDetailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetrics",
			Handler:    _ObjectDetails_GetMetrics_Handler,
		},
		{
			MethodName: "GetQueryExample",
			Handler:    _ObjectDetails_GetQueryExample_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _ObjectDetails_GetLabels_Handler,
		},
		{
			MethodName: "GetQueryPlan",
			Handler:    _ObjectDetails_GetQueryPlan_Handler,
		},
		{
			MethodName: "GetHistogram",
			Handler:    _ObjectDetails_GetHistogram_Handler,
		},
		{
			MethodName: "QueryExists",
			Handler:    _ObjectDetails_QueryExists_Handler,
		},
		{
			MethodName: "QueryByQueryID",
			Handler:    _ObjectDetails_QueryByQueryID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/object_details.proto",
}
