// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qanpb/object_details.proto

package qanv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *MetricsRequest) Validate() error {
	if this.PeriodStartFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartFrom", err)
		}
	}
	if this.PeriodStartTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartTo", err)
		}
	}
	for _, item := range this.Labels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Labels", err)
			}
		}
	}
	return nil
}

func (this *MetricsReply) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	for _, item := range this.Sparkline {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sparkline", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *MetricValues) Validate() error {
	return nil
}

func (this *Labels) Validate() error {
	return nil
}

func (this *QueryExampleRequest) Validate() error {
	if this.PeriodStartFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartFrom", err)
		}
	}
	if this.PeriodStartTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartTo", err)
		}
	}
	for _, item := range this.Labels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Labels", err)
			}
		}
	}
	return nil
}

func (this *QueryExampleReply) Validate() error {
	for _, item := range this.QueryExamples {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("QueryExamples", err)
			}
		}
	}
	return nil
}

func (this *QueryExample) Validate() error {
	return nil
}

func (this *ObjectDetailsLabelsRequest) Validate() error {
	if this.PeriodStartFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartFrom", err)
		}
	}
	if this.PeriodStartTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartTo", err)
		}
	}
	return nil
}

func (this *ObjectDetailsLabelsReply) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

func (this *ListLabelValues) Validate() error {
	return nil
}

func (this *QueryPlanRequest) Validate() error {
	return nil
}

func (this *QueryPlanReply) Validate() error {
	return nil
}

func (this *HistogramRequest) Validate() error {
	if this.PeriodStartFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartFrom", err)
		}
	}
	if this.PeriodStartTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeriodStartTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeriodStartTo", err)
		}
	}
	for _, item := range this.Labels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Labels", err)
			}
		}
	}
	return nil
}

func (this *HistogramReply) Validate() error {
	for _, item := range this.HistogramItems {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HistogramItems", err)
			}
		}
	}
	return nil
}

func (this *HistogramItem) Validate() error {
	return nil
}

func (this *QueryExistsRequest) Validate() error {
	return nil
}
