// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/dataproc/v1/operations.proto

package dataprocpb

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Operation type for Batch resources
type BatchOperationMetadata_BatchOperationType int32

const (
	// Batch operation type is unknown.
	BatchOperationMetadata_BATCH_OPERATION_TYPE_UNSPECIFIED BatchOperationMetadata_BatchOperationType = 0
	// Batch operation type.
	BatchOperationMetadata_BATCH BatchOperationMetadata_BatchOperationType = 1
)

// Enum value maps for BatchOperationMetadata_BatchOperationType.
var (
	BatchOperationMetadata_BatchOperationType_name = map[int32]string{
		0: "BATCH_OPERATION_TYPE_UNSPECIFIED",
		1: "BATCH",
	}
	BatchOperationMetadata_BatchOperationType_value = map[string]int32{
		"BATCH_OPERATION_TYPE_UNSPECIFIED": 0,
		"BATCH":                            1,
	}
)

func (x BatchOperationMetadata_BatchOperationType) Enum() *BatchOperationMetadata_BatchOperationType {
	p := new(BatchOperationMetadata_BatchOperationType)
	*p = x
	return p
}

func (x BatchOperationMetadata_BatchOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatchOperationMetadata_BatchOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataproc_v1_operations_proto_enumTypes[0].Descriptor()
}

func (BatchOperationMetadata_BatchOperationType) Type() protoreflect.EnumType {
	return &file_google_cloud_dataproc_v1_operations_proto_enumTypes[0]
}

func (x BatchOperationMetadata_BatchOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchOperationMetadata_BatchOperationType.Descriptor instead.
func (BatchOperationMetadata_BatchOperationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP(), []int{0, 0}
}

// The operation state.
type ClusterOperationStatus_State int32

const (
	// Unused.
	ClusterOperationStatus_UNKNOWN ClusterOperationStatus_State = 0
	// The operation has been created.
	ClusterOperationStatus_PENDING ClusterOperationStatus_State = 1
	// The operation is running.
	ClusterOperationStatus_RUNNING ClusterOperationStatus_State = 2
	// The operation is done; either cancelled or completed.
	ClusterOperationStatus_DONE ClusterOperationStatus_State = 3
)

// Enum value maps for ClusterOperationStatus_State.
var (
	ClusterOperationStatus_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "RUNNING",
		3: "DONE",
	}
	ClusterOperationStatus_State_value = map[string]int32{
		"UNKNOWN": 0,
		"PENDING": 1,
		"RUNNING": 2,
		"DONE":    3,
	}
)

func (x ClusterOperationStatus_State) Enum() *ClusterOperationStatus_State {
	p := new(ClusterOperationStatus_State)
	*p = x
	return p
}

func (x ClusterOperationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterOperationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataproc_v1_operations_proto_enumTypes[1].Descriptor()
}

func (ClusterOperationStatus_State) Type() protoreflect.EnumType {
	return &file_google_cloud_dataproc_v1_operations_proto_enumTypes[1]
}

func (x ClusterOperationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterOperationStatus_State.Descriptor instead.
func (ClusterOperationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP(), []int{1, 0}
}

// Metadata describing the Batch operation.
type BatchOperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the batch for the operation.
	Batch string `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
	// Batch UUID for the operation.
	BatchUuid string `protobuf:"bytes,2,opt,name=batch_uuid,json=batchUuid,proto3" json:"batch_uuid,omitempty"`
	// The time when the operation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the operation finished.
	DoneTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=done_time,json=doneTime,proto3" json:"done_time,omitempty"`
	// The operation type.
	OperationType BatchOperationMetadata_BatchOperationType `protobuf:"varint,6,opt,name=operation_type,json=operationType,proto3,enum=google.cloud.dataproc.v1.BatchOperationMetadata_BatchOperationType" json:"operation_type,omitempty"`
	// Short description of the operation.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Labels associated with the operation.
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Warnings encountered during operation execution.
	Warnings []string `protobuf:"bytes,9,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *BatchOperationMetadata) Reset() {
	*x = BatchOperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationMetadata) ProtoMessage() {}

func (x *BatchOperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationMetadata.ProtoReflect.Descriptor instead.
func (*BatchOperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP(), []int{0}
}

func (x *BatchOperationMetadata) GetBatch() string {
	if x != nil {
		return x.Batch
	}
	return ""
}

func (x *BatchOperationMetadata) GetBatchUuid() string {
	if x != nil {
		return x.BatchUuid
	}
	return ""
}

func (x *BatchOperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *BatchOperationMetadata) GetDoneTime() *timestamp.Timestamp {
	if x != nil {
		return x.DoneTime
	}
	return nil
}

func (x *BatchOperationMetadata) GetOperationType() BatchOperationMetadata_BatchOperationType {
	if x != nil {
		return x.OperationType
	}
	return BatchOperationMetadata_BATCH_OPERATION_TYPE_UNSPECIFIED
}

func (x *BatchOperationMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BatchOperationMetadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *BatchOperationMetadata) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

// The status of the operation.
type ClusterOperationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. A message containing the operation state.
	State ClusterOperationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.dataproc.v1.ClusterOperationStatus_State" json:"state,omitempty"`
	// Output only. A message containing the detailed operation state.
	InnerState string `protobuf:"bytes,2,opt,name=inner_state,json=innerState,proto3" json:"inner_state,omitempty"`
	// Output only. A message containing any operation metadata details.
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	// Output only. The time this state was entered.
	StateStartTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=state_start_time,json=stateStartTime,proto3" json:"state_start_time,omitempty"`
}

func (x *ClusterOperationStatus) Reset() {
	*x = ClusterOperationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterOperationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterOperationStatus) ProtoMessage() {}

func (x *ClusterOperationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterOperationStatus.ProtoReflect.Descriptor instead.
func (*ClusterOperationStatus) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterOperationStatus) GetState() ClusterOperationStatus_State {
	if x != nil {
		return x.State
	}
	return ClusterOperationStatus_UNKNOWN
}

func (x *ClusterOperationStatus) GetInnerState() string {
	if x != nil {
		return x.InnerState
	}
	return ""
}

func (x *ClusterOperationStatus) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *ClusterOperationStatus) GetStateStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StateStartTime
	}
	return nil
}

// Metadata describing the operation.
type ClusterOperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the cluster for the operation.
	ClusterName string `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Output only. Cluster UUID for the operation.
	ClusterUuid string `protobuf:"bytes,8,opt,name=cluster_uuid,json=clusterUuid,proto3" json:"cluster_uuid,omitempty"`
	// Output only. Current operation status.
	Status *ClusterOperationStatus `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The previous operation status.
	StatusHistory []*ClusterOperationStatus `protobuf:"bytes,10,rep,name=status_history,json=statusHistory,proto3" json:"status_history,omitempty"`
	// Output only. The operation type.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// Output only. Short description of operation.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Labels associated with the operation
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Errors encountered during operation execution.
	Warnings []string `protobuf:"bytes,14,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *ClusterOperationMetadata) Reset() {
	*x = ClusterOperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterOperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterOperationMetadata) ProtoMessage() {}

func (x *ClusterOperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataproc_v1_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterOperationMetadata.ProtoReflect.Descriptor instead.
func (*ClusterOperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterOperationMetadata) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ClusterOperationMetadata) GetClusterUuid() string {
	if x != nil {
		return x.ClusterUuid
	}
	return ""
}

func (x *ClusterOperationMetadata) GetStatus() *ClusterOperationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ClusterOperationMetadata) GetStatusHistory() []*ClusterOperationStatus {
	if x != nil {
		return x.StatusHistory
	}
	return nil
}

func (x *ClusterOperationMetadata) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *ClusterOperationMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ClusterOperationMetadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ClusterOperationMetadata) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

var File_google_cloud_dataproc_v1_operations_proto protoreflect.FileDescriptor

var file_google_cloud_dataproc_v1_operations_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x55, 0x75, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x6f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6a, 0x0a,
	0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24,
	0x0a, 0x20, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x22,
	0xb5, 0x02, 0x0a, 0x16, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x51, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a,
	0x0b, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x22, 0xa3, 0x04, 0x0a, 0x18, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x5c, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x2a, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x12, 0x1f, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x73, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataproc_v1_operations_proto_rawDescOnce sync.Once
	file_google_cloud_dataproc_v1_operations_proto_rawDescData = file_google_cloud_dataproc_v1_operations_proto_rawDesc
)

func file_google_cloud_dataproc_v1_operations_proto_rawDescGZIP() []byte {
	file_google_cloud_dataproc_v1_operations_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataproc_v1_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataproc_v1_operations_proto_rawDescData)
	})
	return file_google_cloud_dataproc_v1_operations_proto_rawDescData
}

var file_google_cloud_dataproc_v1_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_dataproc_v1_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_dataproc_v1_operations_proto_goTypes = []interface{}{
	(BatchOperationMetadata_BatchOperationType)(0), // 0: google.cloud.dataproc.v1.BatchOperationMetadata.BatchOperationType
	(ClusterOperationStatus_State)(0),              // 1: google.cloud.dataproc.v1.ClusterOperationStatus.State
	(*BatchOperationMetadata)(nil),                 // 2: google.cloud.dataproc.v1.BatchOperationMetadata
	(*ClusterOperationStatus)(nil),                 // 3: google.cloud.dataproc.v1.ClusterOperationStatus
	(*ClusterOperationMetadata)(nil),               // 4: google.cloud.dataproc.v1.ClusterOperationMetadata
	nil,                                            // 5: google.cloud.dataproc.v1.BatchOperationMetadata.LabelsEntry
	nil,                                            // 6: google.cloud.dataproc.v1.ClusterOperationMetadata.LabelsEntry
	(*timestamp.Timestamp)(nil),                    // 7: google.protobuf.Timestamp
}
var file_google_cloud_dataproc_v1_operations_proto_depIdxs = []int32{
	7, // 0: google.cloud.dataproc.v1.BatchOperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: google.cloud.dataproc.v1.BatchOperationMetadata.done_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.cloud.dataproc.v1.BatchOperationMetadata.operation_type:type_name -> google.cloud.dataproc.v1.BatchOperationMetadata.BatchOperationType
	5, // 3: google.cloud.dataproc.v1.BatchOperationMetadata.labels:type_name -> google.cloud.dataproc.v1.BatchOperationMetadata.LabelsEntry
	1, // 4: google.cloud.dataproc.v1.ClusterOperationStatus.state:type_name -> google.cloud.dataproc.v1.ClusterOperationStatus.State
	7, // 5: google.cloud.dataproc.v1.ClusterOperationStatus.state_start_time:type_name -> google.protobuf.Timestamp
	3, // 6: google.cloud.dataproc.v1.ClusterOperationMetadata.status:type_name -> google.cloud.dataproc.v1.ClusterOperationStatus
	3, // 7: google.cloud.dataproc.v1.ClusterOperationMetadata.status_history:type_name -> google.cloud.dataproc.v1.ClusterOperationStatus
	6, // 8: google.cloud.dataproc.v1.ClusterOperationMetadata.labels:type_name -> google.cloud.dataproc.v1.ClusterOperationMetadata.LabelsEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_dataproc_v1_operations_proto_init() }
func file_google_cloud_dataproc_v1_operations_proto_init() {
	if File_google_cloud_dataproc_v1_operations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dataproc_v1_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperationMetadata); i {
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
		file_google_cloud_dataproc_v1_operations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterOperationStatus); i {
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
		file_google_cloud_dataproc_v1_operations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterOperationMetadata); i {
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
			RawDescriptor: file_google_cloud_dataproc_v1_operations_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_dataproc_v1_operations_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataproc_v1_operations_proto_depIdxs,
		EnumInfos:         file_google_cloud_dataproc_v1_operations_proto_enumTypes,
		MessageInfos:      file_google_cloud_dataproc_v1_operations_proto_msgTypes,
	}.Build()
	File_google_cloud_dataproc_v1_operations_proto = out.File
	file_google_cloud_dataproc_v1_operations_proto_rawDesc = nil
	file_google_cloud_dataproc_v1_operations_proto_goTypes = nil
	file_google_cloud_dataproc_v1_operations_proto_depIdxs = nil
}
