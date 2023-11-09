// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: managementpb/node/node.proto

package nodev1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "github.com/percona/pmm/api/managementpb/agent"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Node status.
type UniversalNode_Status int32

const (
	// Invalid status.
	UniversalNode_STATUS_UNSPECIFIED UniversalNode_Status = 0
	// The node is up.
	UniversalNode_UP UniversalNode_Status = 1
	// The node is down.
	UniversalNode_DOWN UniversalNode_Status = 2
	// The node's status cannot be known (e.g. there are no metrics yet).
	UniversalNode_UNKNOWN UniversalNode_Status = 3
)

// Enum value maps for UniversalNode_Status.
var (
	UniversalNode_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "UP",
		2: "DOWN",
		3: "UNKNOWN",
	}
	UniversalNode_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"UP":                 1,
		"DOWN":               2,
		"UNKNOWN":            3,
	}
)

func (x UniversalNode_Status) Enum() *UniversalNode_Status {
	p := new(UniversalNode_Status)
	*p = x
	return p
}

func (x UniversalNode_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UniversalNode_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_node_node_proto_enumTypes[0].Descriptor()
}

func (UniversalNode_Status) Type() protoreflect.EnumType {
	return &file_managementpb_node_node_proto_enumTypes[0]
}

func (x UniversalNode_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UniversalNode_Status.Descriptor instead.
func (UniversalNode_Status) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{0, 0}
}

type UniversalNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Node identifier.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node type.
	NodeType string `protobuf:"bytes,2,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	// User-defined node name.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Linux machine-id.
	MachineId string `protobuf:"bytes,4,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	// Linux distribution name and version.
	Distro string `protobuf:"bytes,5,opt,name=distro,proto3" json:"distro,omitempty"`
	// Node model.
	NodeModel string `protobuf:"bytes,6,opt,name=node_model,json=nodeModel,proto3" json:"node_model,omitempty"`
	// A node's unique docker container identifier.
	ContainerId string `protobuf:"bytes,7,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// Container name.
	ContainerName string `protobuf:"bytes,8,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// Node address (DNS name or IP).
	Address string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	// Node region.
	Region string `protobuf:"bytes,10,opt,name=region,proto3" json:"region,omitempty"`
	// Node availability zone.
	Az string `protobuf:"bytes,11,opt,name=az,proto3" json:"az,omitempty"`
	// Custom user-assigned labels for Node.
	CustomLabels map[string]string `protobuf:"bytes,12,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The health status of the node.
	Status UniversalNode_Status `protobuf:"varint,15,opt,name=status,proto3,enum=node.v1beta1.UniversalNode_Status" json:"status,omitempty"`
	// List of agents related to this node.
	Agents []*UniversalNode_Agent `protobuf:"bytes,16,rep,name=agents,proto3" json:"agents,omitempty"`
	// List of services running on this node.
	Services []*UniversalNode_Service `protobuf:"bytes,17,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *UniversalNode) Reset() {
	*x = UniversalNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalNode) ProtoMessage() {}

func (x *UniversalNode) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalNode.ProtoReflect.Descriptor instead.
func (*UniversalNode) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{0}
}

func (x *UniversalNode) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *UniversalNode) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *UniversalNode) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *UniversalNode) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *UniversalNode) GetDistro() string {
	if x != nil {
		return x.Distro
	}
	return ""
}

func (x *UniversalNode) GetNodeModel() string {
	if x != nil {
		return x.NodeModel
	}
	return ""
}

func (x *UniversalNode) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *UniversalNode) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *UniversalNode) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UniversalNode) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UniversalNode) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *UniversalNode) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *UniversalNode) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UniversalNode) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UniversalNode) GetStatus() UniversalNode_Status {
	if x != nil {
		return x.Status
	}
	return UniversalNode_STATUS_UNSPECIFIED
}

func (x *UniversalNode) GetAgents() []*UniversalNode_Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

func (x *UniversalNode) GetServices() []*UniversalNode_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node type to be filtered out.
	NodeType inventorypb.NodeType `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=inventory.NodeType" json:"node_type,omitempty"`
}

func (x *ListNodeRequest) Reset() {
	*x = ListNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeRequest) ProtoMessage() {}

func (x *ListNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeRequest.ProtoReflect.Descriptor instead.
func (*ListNodeRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{1}
}

func (x *ListNodeRequest) GetNodeType() inventorypb.NodeType {
	if x != nil {
		return x.NodeType
	}
	return inventorypb.NodeType(0)
}

type ListNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*UniversalNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ListNodeResponse) Reset() {
	*x = ListNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeResponse) ProtoMessage() {}

func (x *ListNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeResponse.ProtoReflect.Descriptor instead.
func (*ListNodeResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{2}
}

func (x *ListNodeResponse) GetNodes() []*UniversalNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Node identifier.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{3}
}

func (x *GetNodeRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *UniversalNode `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{4}
}

func (x *GetNodeResponse) GetNode() *UniversalNode {
	if x != nil {
		return x.Node
	}
	return nil
}

// Service represents a service running on a node.
type UniversalNode_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Service identifier.
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Service type.
	ServiceType string `protobuf:"bytes,2,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// Service name.
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *UniversalNode_Service) Reset() {
	*x = UniversalNode_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalNode_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalNode_Service) ProtoMessage() {}

func (x *UniversalNode_Service) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalNode_Service.ProtoReflect.Descriptor instead.
func (*UniversalNode_Service) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UniversalNode_Service) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *UniversalNode_Service) GetServiceType() string {
	if x != nil {
		return x.ServiceType
	}
	return ""
}

func (x *UniversalNode_Service) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type UniversalNode_Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Agent identifier.
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// Agent type.
	AgentType string `protobuf:"bytes,2,opt,name=agent_type,json=agentType,proto3" json:"agent_type,omitempty"`
	// Actual Agent status.
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// True if Agent is running and connected to pmm-managed.
	IsConnected bool `protobuf:"varint,4,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
}

func (x *UniversalNode_Agent) Reset() {
	*x = UniversalNode_Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_node_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalNode_Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalNode_Agent) ProtoMessage() {}

func (x *UniversalNode_Agent) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_node_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalNode_Agent.ProtoReflect.Descriptor instead.
func (*UniversalNode_Agent) Descriptor() ([]byte, []int) {
	return file_managementpb_node_node_proto_rawDescGZIP(), []int{0, 1}
}

func (x *UniversalNode_Agent) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *UniversalNode_Agent) GetAgentType() string {
	if x != nil {
		return x.AgentType
	}
	return ""
}

func (x *UniversalNode_Agent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UniversalNode_Agent) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

var File_managementpb_node_node_proto protoreflect.FileDescriptor

var file_managementpb_node_node_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x08, 0x0a, 0x0d, 0x55, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x61, 0x7a, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x7a, 0x12, 0x52,
	0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x3f, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x6e, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x7c, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x1a, 0x3f,
	0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57,
	0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03,
	0x22, 0x43, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x32, 0xc6, 0x02, 0x0a, 0x08,
	0x4d, 0x67, 0x6d, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x2f, 0x12, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x21, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20,
	0x61, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a,
	0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x4e, 0x6f, 0x64, 0x65, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x95, 0x01, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92, 0x41, 0x28, 0x12, 0x08, 0x47, 0x65, 0x74, 0x20,
	0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x20,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x62, 0x79, 0x20, 0x49,
	0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x4e, 0x6f, 0x64, 0x65,
	0x2f, 0x47, 0x65, 0x74, 0x42, 0xa8, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x18, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_node_node_proto_rawDescOnce sync.Once
	file_managementpb_node_node_proto_rawDescData = file_managementpb_node_node_proto_rawDesc
)

func file_managementpb_node_node_proto_rawDescGZIP() []byte {
	file_managementpb_node_node_proto_rawDescOnce.Do(func() {
		file_managementpb_node_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_node_node_proto_rawDescData)
	})
	return file_managementpb_node_node_proto_rawDescData
}

var (
	file_managementpb_node_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_node_node_proto_msgTypes  = make([]protoimpl.MessageInfo, 8)
	file_managementpb_node_node_proto_goTypes   = []interface{}{
		(UniversalNode_Status)(0),     // 0: node.v1beta1.UniversalNode.Status
		(*UniversalNode)(nil),         // 1: node.v1beta1.UniversalNode
		(*ListNodeRequest)(nil),       // 2: node.v1beta1.ListNodeRequest
		(*ListNodeResponse)(nil),      // 3: node.v1beta1.ListNodeResponse
		(*GetNodeRequest)(nil),        // 4: node.v1beta1.GetNodeRequest
		(*GetNodeResponse)(nil),       // 5: node.v1beta1.GetNodeResponse
		(*UniversalNode_Service)(nil), // 6: node.v1beta1.UniversalNode.Service
		(*UniversalNode_Agent)(nil),   // 7: node.v1beta1.UniversalNode.Agent
		nil,                           // 8: node.v1beta1.UniversalNode.CustomLabelsEntry
		(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
		(inventorypb.NodeType)(0),     // 10: inventory.NodeType
	}
)

var file_managementpb_node_node_proto_depIdxs = []int32{
	8,  // 0: node.v1beta1.UniversalNode.custom_labels:type_name -> node.v1beta1.UniversalNode.CustomLabelsEntry
	9,  // 1: node.v1beta1.UniversalNode.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: node.v1beta1.UniversalNode.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 3: node.v1beta1.UniversalNode.status:type_name -> node.v1beta1.UniversalNode.Status
	7,  // 4: node.v1beta1.UniversalNode.agents:type_name -> node.v1beta1.UniversalNode.Agent
	6,  // 5: node.v1beta1.UniversalNode.services:type_name -> node.v1beta1.UniversalNode.Service
	10, // 6: node.v1beta1.ListNodeRequest.node_type:type_name -> inventory.NodeType
	1,  // 7: node.v1beta1.ListNodeResponse.nodes:type_name -> node.v1beta1.UniversalNode
	1,  // 8: node.v1beta1.GetNodeResponse.node:type_name -> node.v1beta1.UniversalNode
	2,  // 9: node.v1beta1.MgmtNode.ListNodes:input_type -> node.v1beta1.ListNodeRequest
	4,  // 10: node.v1beta1.MgmtNode.GetNode:input_type -> node.v1beta1.GetNodeRequest
	3,  // 11: node.v1beta1.MgmtNode.ListNodes:output_type -> node.v1beta1.ListNodeResponse
	5,  // 12: node.v1beta1.MgmtNode.GetNode:output_type -> node.v1beta1.GetNodeResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_managementpb_node_node_proto_init() }
func file_managementpb_node_node_proto_init() {
	if File_managementpb_node_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_node_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalNode_Service); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_node_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalNode_Agent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_node_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_node_node_proto_goTypes,
		DependencyIndexes: file_managementpb_node_node_proto_depIdxs,
		EnumInfos:         file_managementpb_node_node_proto_enumTypes,
		MessageInfos:      file_managementpb_node_node_proto_msgTypes,
	}.Build()
	File_managementpb_node_node_proto = out.File
	file_managementpb_node_node_proto_rawDesc = nil
	file_managementpb_node_node_proto_goTypes = nil
	file_managementpb_node_node_proto_depIdxs = nil
}
