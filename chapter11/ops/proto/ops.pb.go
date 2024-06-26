// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.18.0
// source: ops.proto

package proto

import (
	model "github.com/ibiscum/Go-for-DevOps/chapter11/ops/proto/jaeger/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SamplerType int32

const (
	SamplerType_STUnknown SamplerType = 0
	SamplerType_STNever   SamplerType = 1
	SamplerType_STAlways  SamplerType = 2
	SamplerType_STFloat   SamplerType = 3
)

// Enum value maps for SamplerType.
var (
	SamplerType_name = map[int32]string{
		0: "STUnknown",
		1: "STNever",
		2: "STAlways",
		3: "STFloat",
	}
	SamplerType_value = map[string]int32{
		"STUnknown": 0,
		"STNever":   1,
		"STAlways":  2,
		"STFloat":   3,
	}
)

func (x SamplerType) Enum() *SamplerType {
	p := new(SamplerType)
	*p = x
	return p
}

func (x SamplerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamplerType) Descriptor() protoreflect.EnumDescriptor {
	return file_ops_proto_enumTypes[0].Descriptor()
}

func (SamplerType) Type() protoreflect.EnumType {
	return &file_ops_proto_enumTypes[0]
}

func (x SamplerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamplerType.Descriptor instead.
func (SamplerType) EnumDescriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{0}
}

// The request to get traces from Jaeger.
type ListTracesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service to find traces for.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Filter the traces for this operation.
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Filter the traces for only traces with these all these tags.
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	// Traces must start after the time in unix nanoseconds.
	Start int64 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	// Traces must end before this time in unix nanoseconds.
	End int64 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	// The minimum duration of a matched trace.
	DurationMin int64 `protobuf:"varint,6,opt,name=duration_min,json=durationMin,proto3" json:"duration_min,omitempty"`
	// The maximum duration of a matched trace.
	DurationMax int64 `protobuf:"varint,7,opt,name=duration_max,json=durationMax,proto3" json:"duration_max,omitempty"`
	// The number of traces to return.
	SearchDepth int32 `protobuf:"varint,8,opt,name=search_depth,json=searchDepth,proto3" json:"search_depth,omitempty"`
}

func (x *ListTracesReq) Reset() {
	*x = ListTracesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTracesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTracesReq) ProtoMessage() {}

func (x *ListTracesReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTracesReq.ProtoReflect.Descriptor instead.
func (*ListTracesReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{0}
}

func (x *ListTracesReq) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ListTracesReq) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *ListTracesReq) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ListTracesReq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ListTracesReq) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ListTracesReq) GetDurationMin() int64 {
	if x != nil {
		return x.DurationMin
	}
	return 0
}

func (x *ListTracesReq) GetDurationMax() int64 {
	if x != nil {
		return x.DurationMax
	}
	return 0
}

func (x *ListTracesReq) GetSearchDepth() int32 {
	if x != nil {
		return x.SearchDepth
	}
	return 0
}

// This represents a trace identifier and when the trace started.
type TraceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The trace identifier in hex form.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the trace started in unix nanosecods.
	Start int64 `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *TraceItem) Reset() {
	*x = TraceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceItem) ProtoMessage() {}

func (x *TraceItem) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceItem.ProtoReflect.Descriptor instead.
func (*TraceItem) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{1}
}

func (x *TraceItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TraceItem) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

// A response to ListTraces().
type ListTracesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of traces that met the search criteria.
	Traces []*TraceItem `protobuf:"bytes,1,rep,name=traces,proto3" json:"traces,omitempty"`
}

func (x *ListTracesResp) Reset() {
	*x = ListTracesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTracesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTracesResp) ProtoMessage() {}

func (x *ListTracesResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTracesResp.ProtoReflect.Descriptor instead.
func (*ListTracesResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{2}
}

func (x *ListTracesResp) GetTraces() []*TraceItem {
	if x != nil {
		return x.Traces
	}
	return nil
}

// The request to get a URL showing the trace information for a trace id.
type ShowTraceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the trace in hex.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowTraceReq) Reset() {
	*x = ShowTraceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowTraceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowTraceReq) ProtoMessage() {}

func (x *ShowTraceReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowTraceReq.ProtoReflect.Descriptor instead.
func (*ShowTraceReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{3}
}

func (x *ShowTraceReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The resonse to ShowTrace().
type ShowTraceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URL to view the trace.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the operations being performed.
	Operations []string `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// A list of tag values in spans labelled "error".
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	// A list of all tags in the spans.
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// The longest duration found in any span.
	Duration *durationpb.Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ShowTraceResp) Reset() {
	*x = ShowTraceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowTraceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowTraceResp) ProtoMessage() {}

func (x *ShowTraceResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowTraceResp.ProtoReflect.Descriptor instead.
func (*ShowTraceResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{4}
}

func (x *ShowTraceResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShowTraceResp) GetOperations() []string {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *ShowTraceResp) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ShowTraceResp) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ShowTraceResp) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type ShowLogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hex ID of the trace.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowLogsReq) Reset() {
	*x = ShowLogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowLogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowLogsReq) ProtoMessage() {}

func (x *ShowLogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowLogsReq.ProtoReflect.Descriptor instead.
func (*ShowLogsReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{5}
}

func (x *ShowLogsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ShowLogsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Logs []*model.Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ShowLogsResp) Reset() {
	*x = ShowLogsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowLogsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowLogsResp) ProtoMessage() {}

func (x *ShowLogsResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowLogsResp.ProtoReflect.Descriptor instead.
func (*ShowLogsResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{6}
}

func (x *ShowLogsResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShowLogsResp) GetLogs() []*model.Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

// Used to request we change the OTEL sampling.
type ChangeSamplingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of sampling to change to.
	Type SamplerType `protobuf:"varint,1,opt,name=type,proto3,enum=ops.SamplerType" json:"type,omitempty"`
	// This is the sampling rate if type == STFloat. Values must be
	// > 0 and <= 1.0 .
	FloatValue float64 `protobuf:"fixed64,2,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
}

func (x *ChangeSamplingReq) Reset() {
	*x = ChangeSamplingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSamplingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSamplingReq) ProtoMessage() {}

func (x *ChangeSamplingReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSamplingReq.ProtoReflect.Descriptor instead.
func (*ChangeSamplingReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeSamplingReq) GetType() SamplerType {
	if x != nil {
		return x.Type
	}
	return SamplerType_STUnknown
}

func (x *ChangeSamplingReq) GetFloatValue() float64 {
	if x != nil {
		return x.FloatValue
	}
	return 0
}

// The response to a sampling change.
type ChangeSamplingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeSamplingResp) Reset() {
	*x = ChangeSamplingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSamplingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSamplingResp) ProtoMessage() {}

func (x *ChangeSamplingResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSamplingResp.ProtoReflect.Descriptor instead.
func (*ChangeSamplingResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{8}
}

// The request to get the deployed version of the service.
type DeployedVersionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeployedVersionReq) Reset() {
	*x = DeployedVersionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedVersionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedVersionReq) ProtoMessage() {}

func (x *DeployedVersionReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedVersionReq.ProtoReflect.Descriptor instead.
func (*DeployedVersionReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{9}
}

// The response to DeployedVersion().
type DeployedVersionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version running according to prometheus metrics.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DeployedVersionResp) Reset() {
	*x = DeployedVersionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedVersionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedVersionResp) ProtoMessage() {}

func (x *DeployedVersionResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedVersionResp.ProtoReflect.Descriptor instead.
func (*DeployedVersionResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{10}
}

func (x *DeployedVersionResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Alert describes a Prometheus alert.
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the current state of the alert.
	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// This is the current value of the alert.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// This is how long it has been active in unix nanoseconds.
	ActiveAt int64 `protobuf:"varint,3,opt,name=active_at,json=activeAt,proto3" json:"active_at,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{11}
}

func (x *Alert) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Alert) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Alert) GetActiveAt() int64 {
	if x != nil {
		return x.ActiveAt
	}
	return 0
}

// This requests an set of active alerts in the system.
type AlertsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Labels that the alert must match. Must have all labels. None indicates all alerts.
	Labels []string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// It must be an alert that is active since this time in unix nanoseconds. 0 indicates any alive alert.
	ActiveAt int64 `protobuf:"varint,2,opt,name=active_at,json=activeAt,proto3" json:"active_at,omitempty"`
	// It must have one of these states. None indicates all states.
	States []string `protobuf:"bytes,3,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *AlertsReq) Reset() {
	*x = AlertsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsReq) ProtoMessage() {}

func (x *AlertsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsReq.ProtoReflect.Descriptor instead.
func (*AlertsReq) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{12}
}

func (x *AlertsReq) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *AlertsReq) GetActiveAt() int64 {
	if x != nil {
		return x.ActiveAt
	}
	return 0
}

func (x *AlertsReq) GetStates() []string {
	if x != nil {
		return x.States
	}
	return nil
}

// The response to Alerts().
type AlertsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of alerts that matched the filter.
	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AlertsResp) Reset() {
	*x = AlertsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsResp) ProtoMessage() {}

func (x *AlertsResp) ProtoReflect() protoreflect.Message {
	mi := &file_ops_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsResp.ProtoReflect.Descriptor instead.
func (*AlertsResp) Descriptor() ([]byte, []int) {
	return file_ops_proto_rawDescGZIP(), []int{13}
}

func (x *AlertsResp) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_ops_proto protoreflect.FileDescriptor

var file_ops_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x70, 0x73,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x70, 0x74, 0x68, 0x22, 0x31, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x38, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26,
	0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x77, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x0b, 0x53,
	0x68, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x68,
	0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x05, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x74, 0x22, 0x58, 0x0a,
	0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2a, 0x44, 0x0a, 0x0b, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4e, 0x65, 0x76,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x03, 0x32,
	0xe1, 0x02, 0x0a, 0x03, 0x4f, 0x70, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6f, 0x70, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x11, 0x2e,
	0x6f, 0x70, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x77, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x10, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x6f, 0x70,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6f, 0x70, 0x73,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x12, 0x0e, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x0f, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x50, 0x61, 0x63, 0x6b, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x47, 0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x44, 0x65, 0x76, 0x4f, 0x70, 0x73, 0x2f,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x31, 0x30, 0x2f, 0x6f, 0x70, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ops_proto_rawDescOnce sync.Once
	file_ops_proto_rawDescData = file_ops_proto_rawDesc
)

func file_ops_proto_rawDescGZIP() []byte {
	file_ops_proto_rawDescOnce.Do(func() {
		file_ops_proto_rawDescData = protoimpl.X.CompressGZIP(file_ops_proto_rawDescData)
	})
	return file_ops_proto_rawDescData
}

var file_ops_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ops_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_ops_proto_goTypes = []interface{}{
	(SamplerType)(0),            // 0: ops.SamplerType
	(*ListTracesReq)(nil),       // 1: ops.ListTracesReq
	(*TraceItem)(nil),           // 2: ops.TraceItem
	(*ListTracesResp)(nil),      // 3: ops.ListTracesResp
	(*ShowTraceReq)(nil),        // 4: ops.ShowTraceReq
	(*ShowTraceResp)(nil),       // 5: ops.ShowTraceResp
	(*ShowLogsReq)(nil),         // 6: ops.ShowLogsReq
	(*ShowLogsResp)(nil),        // 7: ops.ShowLogsResp
	(*ChangeSamplingReq)(nil),   // 8: ops.ChangeSamplingReq
	(*ChangeSamplingResp)(nil),  // 9: ops.ChangeSamplingResp
	(*DeployedVersionReq)(nil),  // 10: ops.DeployedVersionReq
	(*DeployedVersionResp)(nil), // 11: ops.DeployedVersionResp
	(*Alert)(nil),               // 12: ops.Alert
	(*AlertsReq)(nil),           // 13: ops.AlertsReq
	(*AlertsResp)(nil),          // 14: ops.AlertsResp
	(*durationpb.Duration)(nil), // 15: google.protobuf.Duration
	(*model.Log)(nil),           // 16: jaeger.api_v2.Log
}
var file_ops_proto_depIdxs = []int32{
	2,  // 0: ops.ListTracesResp.traces:type_name -> ops.TraceItem
	15, // 1: ops.ShowTraceResp.duration:type_name -> google.protobuf.Duration
	16, // 2: ops.ShowLogsResp.logs:type_name -> jaeger.api_v2.Log
	0,  // 3: ops.ChangeSamplingReq.type:type_name -> ops.SamplerType
	12, // 4: ops.AlertsResp.alerts:type_name -> ops.Alert
	1,  // 5: ops.Ops.ListTraces:input_type -> ops.ListTracesReq
	4,  // 6: ops.Ops.ShowTrace:input_type -> ops.ShowTraceReq
	6,  // 7: ops.Ops.ShowLogs:input_type -> ops.ShowLogsReq
	8,  // 8: ops.Ops.ChangeSampling:input_type -> ops.ChangeSamplingReq
	10, // 9: ops.Ops.DeployedVersion:input_type -> ops.DeployedVersionReq
	13, // 10: ops.Ops.Alerts:input_type -> ops.AlertsReq
	3,  // 11: ops.Ops.ListTraces:output_type -> ops.ListTracesResp
	5,  // 12: ops.Ops.ShowTrace:output_type -> ops.ShowTraceResp
	7,  // 13: ops.Ops.ShowLogs:output_type -> ops.ShowLogsResp
	9,  // 14: ops.Ops.ChangeSampling:output_type -> ops.ChangeSamplingResp
	11, // 15: ops.Ops.DeployedVersion:output_type -> ops.DeployedVersionResp
	14, // 16: ops.Ops.Alerts:output_type -> ops.AlertsResp
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ops_proto_init() }
func file_ops_proto_init() {
	if File_ops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTracesReq); i {
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
		file_ops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceItem); i {
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
		file_ops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTracesResp); i {
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
		file_ops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowTraceReq); i {
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
		file_ops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowTraceResp); i {
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
		file_ops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowLogsReq); i {
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
		file_ops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowLogsResp); i {
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
		file_ops_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSamplingReq); i {
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
		file_ops_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSamplingResp); i {
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
		file_ops_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedVersionReq); i {
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
		file_ops_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedVersionResp); i {
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
		file_ops_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_ops_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsReq); i {
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
		file_ops_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsResp); i {
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
			RawDescriptor: file_ops_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ops_proto_goTypes,
		DependencyIndexes: file_ops_proto_depIdxs,
		EnumInfos:         file_ops_proto_enumTypes,
		MessageInfos:      file_ops_proto_msgTypes,
	}.Build()
	File_ops_proto = out.File
	file_ops_proto_rawDesc = nil
	file_ops_proto_goTypes = nil
	file_ops_proto_depIdxs = nil
}
