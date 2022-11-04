// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/node.proto

package managementpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"

	_ "github.com/percona/pmm/api/inventorypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *RegisterNodeRequest) Validate() error {
	if this.NodeName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeName", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeName))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *RegisterNodeResponse) Validate() error {
	if this.GenericNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.GenericNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("GenericNode", err)
		}
	}
	if this.ContainerNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContainerNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContainerNode", err)
		}
	}
	if this.PmmAgent != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PmmAgent); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PmmAgent", err)
		}
	}
	return nil
}
