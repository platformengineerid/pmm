// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: agentlocalpb/agentlocal.proto

package agentlocalpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"

	inventorypb "github.com/percona/pmm/api/inventorypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServerInfo contains information about the PMM Server.
type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PMM Server URL in a form https://HOST:PORT/.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// PMM Server's TLS certificate validation should be skipped if true.
	InsecureTls bool `protobuf:"varint,2,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty"`
	// True if pmm-agent is currently connected to the server.
	Connected bool `protobuf:"varint,3,opt,name=connected,proto3" json:"connected,omitempty"`
	// PMM Server version (if agent is connected).
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Ping time from pmm-agent to pmm-managed (if agent is connected).
	Latency *durationpb.Duration `protobuf:"bytes,5,opt,name=latency,proto3" json:"latency,omitempty"`
	// Clock drift from PMM Server (if agent is connected).
	ClockDrift *durationpb.Duration `protobuf:"bytes,6,opt,name=clock_drift,json=clockDrift,proto3" json:"clock_drift,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{0}
}

func (x *ServerInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ServerInfo) GetInsecureTls() bool {
	if x != nil {
		return x.InsecureTls
	}
	return false
}

func (x *ServerInfo) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *ServerInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerInfo) GetLatency() *durationpb.Duration {
	if x != nil {
		return x.Latency
	}
	return nil
}

func (x *ServerInfo) GetClockDrift() *durationpb.Duration {
	if x != nil {
		return x.ClockDrift
	}
	return nil
}

// AgentInfo contains information about Agent managed by pmm-agent.
type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId   string                  `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	AgentType inventorypb.AgentType   `protobuf:"varint,2,opt,name=agent_type,json=agentType,proto3,enum=inventory.AgentType" json:"agent_type,omitempty"`
	Status    inventorypb.AgentStatus `protobuf:"varint,3,opt,name=status,proto3,enum=inventory.AgentStatus" json:"status,omitempty"`
	// The current listen port of this Agent (exporter or vmagent).
	// Zero for other Agent types, or if unknown or not yet supported.
	ListenPort      uint32 `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	ProcessExecPath string `protobuf:"bytes,5,opt,name=process_exec_path,json=processExecPath,proto3" json:"process_exec_path,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{1}
}

func (x *AgentInfo) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentInfo) GetAgentType() inventorypb.AgentType {
	if x != nil {
		return x.AgentType
	}
	return inventorypb.AgentType(0)
}

func (x *AgentInfo) GetStatus() inventorypb.AgentStatus {
	if x != nil {
		return x.Status
	}
	return inventorypb.AgentStatus(0)
}

func (x *AgentInfo) GetListenPort() uint32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *AgentInfo) GetProcessExecPath() string {
	if x != nil {
		return x.ProcessExecPath
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returns network info (latency and clock_drift) if true.
	GetNetworkInfo bool `protobuf:"varint,1,opt,name=get_network_info,json=getNetworkInfo,proto3" json:"get_network_info,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{2}
}

func (x *StatusRequest) GetGetNetworkInfo() bool {
	if x != nil {
		return x.GetNetworkInfo
	}
	return false
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId      string       `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	RunsOnNodeId string       `protobuf:"bytes,2,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	ServerInfo   *ServerInfo  `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	AgentsInfo   []*AgentInfo `protobuf:"bytes,4,rep,name=agents_info,json=agentsInfo,proto3" json:"agents_info,omitempty"`
	// Config file path if pmm-agent was started with one.
	ConfigFilepath string `protobuf:"bytes,5,opt,name=config_filepath,json=configFilepath,proto3" json:"config_filepath,omitempty"`
	// PMM Agent version.
	AgentVersion string `protobuf:"bytes,6,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
	NodeName     string `protobuf:"bytes,7,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Shows connection uptime in percentage between agent and server
	ConnectionUptime float32 `protobuf:"fixed32,8,opt,name=connection_uptime,json=connectionUptime,proto3" json:"connection_uptime,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{3}
}

func (x *StatusResponse) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *StatusResponse) GetRunsOnNodeId() string {
	if x != nil {
		return x.RunsOnNodeId
	}
	return ""
}

func (x *StatusResponse) GetServerInfo() *ServerInfo {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

func (x *StatusResponse) GetAgentsInfo() []*AgentInfo {
	if x != nil {
		return x.AgentsInfo
	}
	return nil
}

func (x *StatusResponse) GetConfigFilepath() string {
	if x != nil {
		return x.ConfigFilepath
	}
	return ""
}

func (x *StatusResponse) GetAgentVersion() string {
	if x != nil {
		return x.AgentVersion
	}
	return ""
}

func (x *StatusResponse) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *StatusResponse) GetConnectionUptime() float32 {
	if x != nil {
		return x.ConnectionUptime
	}
	return 0
}

type ReloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReloadRequest) Reset() {
	*x = ReloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadRequest) ProtoMessage() {}

func (x *ReloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadRequest.ProtoReflect.Descriptor instead.
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{4}
}

// ReloadRequest may not be received by the client due to pmm-agent restart.
type ReloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReloadResponse) Reset() {
	*x = ReloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentlocalpb_agentlocal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadResponse) ProtoMessage() {}

func (x *ReloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agentlocalpb_agentlocal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadResponse.ProtoReflect.Descriptor instead.
func (*ReloadResponse) Descriptor() ([]byte, []int) {
	return file_agentlocalpb_agentlocal_proto_rawDescGZIP(), []int{5}
}

var File_agentlocalpb_agentlocal_proto protoreflect.FileDescriptor

var file_agentlocalpb_agentlocal_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x70, 0x62, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x33, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x72,
	0x69, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x72, 0x69, 0x66, 0x74,
	0x22, 0xd8, 0x01, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x78, 0x65, 0x63, 0x50, 0x61, 0x74, 0x68, 0x22, 0x39, 0x0a, 0x0d, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10,
	0x67, 0x65, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xdb, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x73, 0x5f, 0x6f, 0x6e, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x75, 0x6e, 0x73, 0x4f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xda, 0x01, 0x0a, 0x11, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x5a, 0x0f, 0x12, 0x0d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x2f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x59, 0x0a, 0x06, 0x52, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x52, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x52, 0x65,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x92, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x42, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70,
	0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0xca, 0x02, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0xe2, 0x02, 0x16, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agentlocalpb_agentlocal_proto_rawDescOnce sync.Once
	file_agentlocalpb_agentlocal_proto_rawDescData = file_agentlocalpb_agentlocal_proto_rawDesc
)

func file_agentlocalpb_agentlocal_proto_rawDescGZIP() []byte {
	file_agentlocalpb_agentlocal_proto_rawDescOnce.Do(func() {
		file_agentlocalpb_agentlocal_proto_rawDescData = protoimpl.X.CompressGZIP(file_agentlocalpb_agentlocal_proto_rawDescData)
	})
	return file_agentlocalpb_agentlocal_proto_rawDescData
}

var (
	file_agentlocalpb_agentlocal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_agentlocalpb_agentlocal_proto_goTypes  = []interface{}{
		(*ServerInfo)(nil),           // 0: agentlocal.ServerInfo
		(*AgentInfo)(nil),            // 1: agentlocal.AgentInfo
		(*StatusRequest)(nil),        // 2: agentlocal.StatusRequest
		(*StatusResponse)(nil),       // 3: agentlocal.StatusResponse
		(*ReloadRequest)(nil),        // 4: agentlocal.ReloadRequest
		(*ReloadResponse)(nil),       // 5: agentlocal.ReloadResponse
		(*durationpb.Duration)(nil),  // 6: google.protobuf.Duration
		(inventorypb.AgentType)(0),   // 7: inventory.AgentType
		(inventorypb.AgentStatus)(0), // 8: inventory.AgentStatus
	}
)

var file_agentlocalpb_agentlocal_proto_depIdxs = []int32{
	6, // 0: agentlocal.ServerInfo.latency:type_name -> google.protobuf.Duration
	6, // 1: agentlocal.ServerInfo.clock_drift:type_name -> google.protobuf.Duration
	7, // 2: agentlocal.AgentInfo.agent_type:type_name -> inventory.AgentType
	8, // 3: agentlocal.AgentInfo.status:type_name -> inventory.AgentStatus
	0, // 4: agentlocal.StatusResponse.server_info:type_name -> agentlocal.ServerInfo
	1, // 5: agentlocal.StatusResponse.agents_info:type_name -> agentlocal.AgentInfo
	2, // 6: agentlocal.AgentLocalService.Status:input_type -> agentlocal.StatusRequest
	4, // 7: agentlocal.AgentLocalService.Reload:input_type -> agentlocal.ReloadRequest
	3, // 8: agentlocal.AgentLocalService.Status:output_type -> agentlocal.StatusResponse
	5, // 9: agentlocal.AgentLocalService.Reload:output_type -> agentlocal.ReloadResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_agentlocalpb_agentlocal_proto_init() }
func file_agentlocalpb_agentlocal_proto_init() {
	if File_agentlocalpb_agentlocal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agentlocalpb_agentlocal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_agentlocalpb_agentlocal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
		file_agentlocalpb_agentlocal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_agentlocalpb_agentlocal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_agentlocalpb_agentlocal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadRequest); i {
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
		file_agentlocalpb_agentlocal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadResponse); i {
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
			RawDescriptor: file_agentlocalpb_agentlocal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agentlocalpb_agentlocal_proto_goTypes,
		DependencyIndexes: file_agentlocalpb_agentlocal_proto_depIdxs,
		MessageInfos:      file_agentlocalpb_agentlocal_proto_msgTypes,
	}.Build()
	File_agentlocalpb_agentlocal_proto = out.File
	file_agentlocalpb_agentlocal_proto_rawDesc = nil
	file_agentlocalpb_agentlocal_proto_goTypes = nil
	file_agentlocalpb_agentlocal_proto_depIdxs = nil
}
