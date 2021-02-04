// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: tensorflow/core/framework/step_stats.proto

package step_stats_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	allocation_description_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/allocation_description_go_proto"
	tensor_description_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_description_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// An allocation/de-allocation operation performed by the allocator.
type AllocationRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp of the operation.
	AllocMicros int64 `protobuf:"varint,1,opt,name=alloc_micros,json=allocMicros,proto3" json:"alloc_micros,omitempty"`
	// Number of bytes allocated, or de-allocated if negative.
	AllocBytes int64 `protobuf:"varint,2,opt,name=alloc_bytes,json=allocBytes,proto3" json:"alloc_bytes,omitempty"`
}

func (x *AllocationRecord) Reset() {
	*x = AllocationRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationRecord) ProtoMessage() {}

func (x *AllocationRecord) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationRecord.ProtoReflect.Descriptor instead.
func (*AllocationRecord) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{0}
}

func (x *AllocationRecord) GetAllocMicros() int64 {
	if x != nil {
		return x.AllocMicros
	}
	return 0
}

func (x *AllocationRecord) GetAllocBytes() int64 {
	if x != nil {
		return x.AllocBytes
	}
	return 0
}

type AllocatorMemoryUsed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllocatorName string `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// These are per-node allocator memory stats.
	TotalBytes int64 `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	PeakBytes  int64 `protobuf:"varint,3,opt,name=peak_bytes,json=peakBytes,proto3" json:"peak_bytes,omitempty"`
	// The bytes that are not deallocated.
	LiveBytes int64 `protobuf:"varint,4,opt,name=live_bytes,json=liveBytes,proto3" json:"live_bytes,omitempty"`
	// The allocation and deallocation timeline.
	AllocationRecords []*AllocationRecord `protobuf:"bytes,6,rep,name=allocation_records,json=allocationRecords,proto3" json:"allocation_records,omitempty"`
	// These are snapshots of the overall allocator memory stats.
	// The number of live bytes currently allocated by the allocator.
	AllocatorBytesInUse int64 `protobuf:"varint,5,opt,name=allocator_bytes_in_use,json=allocatorBytesInUse,proto3" json:"allocator_bytes_in_use,omitempty"`
}

func (x *AllocatorMemoryUsed) Reset() {
	*x = AllocatorMemoryUsed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocatorMemoryUsed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocatorMemoryUsed) ProtoMessage() {}

func (x *AllocatorMemoryUsed) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocatorMemoryUsed.ProtoReflect.Descriptor instead.
func (*AllocatorMemoryUsed) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{1}
}

func (x *AllocatorMemoryUsed) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

func (x *AllocatorMemoryUsed) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *AllocatorMemoryUsed) GetPeakBytes() int64 {
	if x != nil {
		return x.PeakBytes
	}
	return 0
}

func (x *AllocatorMemoryUsed) GetLiveBytes() int64 {
	if x != nil {
		return x.LiveBytes
	}
	return 0
}

func (x *AllocatorMemoryUsed) GetAllocationRecords() []*AllocationRecord {
	if x != nil {
		return x.AllocationRecords
	}
	return nil
}

func (x *AllocatorMemoryUsed) GetAllocatorBytesInUse() int64 {
	if x != nil {
		return x.AllocatorBytesInUse
	}
	return 0
}

// Output sizes recorded for a single execution of a graph node.
type NodeOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot              int32                                          `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	TensorDescription *tensor_description_go_proto.TensorDescription `protobuf:"bytes,3,opt,name=tensor_description,json=tensorDescription,proto3" json:"tensor_description,omitempty"`
}

func (x *NodeOutput) Reset() {
	*x = NodeOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeOutput) ProtoMessage() {}

func (x *NodeOutput) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeOutput.ProtoReflect.Descriptor instead.
func (*NodeOutput) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{2}
}

func (x *NodeOutput) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *NodeOutput) GetTensorDescription() *tensor_description_go_proto.TensorDescription {
	if x != nil {
		return x.TensorDescription
	}
	return nil
}

// For memory tracking.
type MemoryStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempMemorySize           int64   `protobuf:"varint,1,opt,name=temp_memory_size,json=tempMemorySize,proto3" json:"temp_memory_size,omitempty"`
	PersistentMemorySize     int64   `protobuf:"varint,3,opt,name=persistent_memory_size,json=persistentMemorySize,proto3" json:"persistent_memory_size,omitempty"`
	PersistentTensorAllocIds []int64 `protobuf:"varint,5,rep,packed,name=persistent_tensor_alloc_ids,json=persistentTensorAllocIds,proto3" json:"persistent_tensor_alloc_ids,omitempty"`
	// Deprecated: Do not use.
	DeviceTempMemorySize int64 `protobuf:"varint,2,opt,name=device_temp_memory_size,json=deviceTempMemorySize,proto3" json:"device_temp_memory_size,omitempty"`
	// Deprecated: Do not use.
	DevicePersistentMemorySize int64 `protobuf:"varint,4,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize,proto3" json:"device_persistent_memory_size,omitempty"`
	// Deprecated: Do not use.
	DevicePersistentTensorAllocIds []int64 `protobuf:"varint,6,rep,packed,name=device_persistent_tensor_alloc_ids,json=devicePersistentTensorAllocIds,proto3" json:"device_persistent_tensor_alloc_ids,omitempty"`
}

func (x *MemoryStats) Reset() {
	*x = MemoryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStats) ProtoMessage() {}

func (x *MemoryStats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStats.ProtoReflect.Descriptor instead.
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{3}
}

func (x *MemoryStats) GetTempMemorySize() int64 {
	if x != nil {
		return x.TempMemorySize
	}
	return 0
}

func (x *MemoryStats) GetPersistentMemorySize() int64 {
	if x != nil {
		return x.PersistentMemorySize
	}
	return 0
}

func (x *MemoryStats) GetPersistentTensorAllocIds() []int64 {
	if x != nil {
		return x.PersistentTensorAllocIds
	}
	return nil
}

// Deprecated: Do not use.
func (x *MemoryStats) GetDeviceTempMemorySize() int64 {
	if x != nil {
		return x.DeviceTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (x *MemoryStats) GetDevicePersistentMemorySize() int64 {
	if x != nil {
		return x.DevicePersistentMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (x *MemoryStats) GetDevicePersistentTensorAllocIds() []int64 {
	if x != nil {
		return x.DevicePersistentTensorAllocIds
	}
	return nil
}

// Time/size stats recorded for a single execution of a graph node.
type NodeExecStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO(tucker): Use some more compact form of node identity than
	// the full string name.  Either all processes should agree on a
	// global id (cost_id?) for each node, or we should use a hash of
	// the name.
	NodeName         string                                                   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AllStartMicros   int64                                                    `protobuf:"varint,2,opt,name=all_start_micros,json=allStartMicros,proto3" json:"all_start_micros,omitempty"`
	OpStartRelMicros int64                                                    `protobuf:"varint,3,opt,name=op_start_rel_micros,json=opStartRelMicros,proto3" json:"op_start_rel_micros,omitempty"`
	OpEndRelMicros   int64                                                    `protobuf:"varint,4,opt,name=op_end_rel_micros,json=opEndRelMicros,proto3" json:"op_end_rel_micros,omitempty"`
	AllEndRelMicros  int64                                                    `protobuf:"varint,5,opt,name=all_end_rel_micros,json=allEndRelMicros,proto3" json:"all_end_rel_micros,omitempty"`
	Memory           []*AllocatorMemoryUsed                                   `protobuf:"bytes,6,rep,name=memory,proto3" json:"memory,omitempty"`
	Output           []*NodeOutput                                            `protobuf:"bytes,7,rep,name=output,proto3" json:"output,omitempty"`
	TimelineLabel    string                                                   `protobuf:"bytes,8,opt,name=timeline_label,json=timelineLabel,proto3" json:"timeline_label,omitempty"`
	ScheduledMicros  int64                                                    `protobuf:"varint,9,opt,name=scheduled_micros,json=scheduledMicros,proto3" json:"scheduled_micros,omitempty"`
	ThreadId         uint32                                                   `protobuf:"varint,10,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	ReferencedTensor []*allocation_description_go_proto.AllocationDescription `protobuf:"bytes,11,rep,name=referenced_tensor,json=referencedTensor,proto3" json:"referenced_tensor,omitempty"`
	MemoryStats      *MemoryStats                                             `protobuf:"bytes,12,opt,name=memory_stats,json=memoryStats,proto3" json:"memory_stats,omitempty"`
	AllStartNanos    int64                                                    `protobuf:"varint,13,opt,name=all_start_nanos,json=allStartNanos,proto3" json:"all_start_nanos,omitempty"`
	OpStartRelNanos  int64                                                    `protobuf:"varint,14,opt,name=op_start_rel_nanos,json=opStartRelNanos,proto3" json:"op_start_rel_nanos,omitempty"`
	OpEndRelNanos    int64                                                    `protobuf:"varint,15,opt,name=op_end_rel_nanos,json=opEndRelNanos,proto3" json:"op_end_rel_nanos,omitempty"`
	AllEndRelNanos   int64                                                    `protobuf:"varint,16,opt,name=all_end_rel_nanos,json=allEndRelNanos,proto3" json:"all_end_rel_nanos,omitempty"`
	ScheduledNanos   int64                                                    `protobuf:"varint,17,opt,name=scheduled_nanos,json=scheduledNanos,proto3" json:"scheduled_nanos,omitempty"`
}

func (x *NodeExecStats) Reset() {
	*x = NodeExecStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExecStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExecStats) ProtoMessage() {}

func (x *NodeExecStats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExecStats.ProtoReflect.Descriptor instead.
func (*NodeExecStats) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{4}
}

func (x *NodeExecStats) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeExecStats) GetAllStartMicros() int64 {
	if x != nil {
		return x.AllStartMicros
	}
	return 0
}

func (x *NodeExecStats) GetOpStartRelMicros() int64 {
	if x != nil {
		return x.OpStartRelMicros
	}
	return 0
}

func (x *NodeExecStats) GetOpEndRelMicros() int64 {
	if x != nil {
		return x.OpEndRelMicros
	}
	return 0
}

func (x *NodeExecStats) GetAllEndRelMicros() int64 {
	if x != nil {
		return x.AllEndRelMicros
	}
	return 0
}

func (x *NodeExecStats) GetMemory() []*AllocatorMemoryUsed {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *NodeExecStats) GetOutput() []*NodeOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *NodeExecStats) GetTimelineLabel() string {
	if x != nil {
		return x.TimelineLabel
	}
	return ""
}

func (x *NodeExecStats) GetScheduledMicros() int64 {
	if x != nil {
		return x.ScheduledMicros
	}
	return 0
}

func (x *NodeExecStats) GetThreadId() uint32 {
	if x != nil {
		return x.ThreadId
	}
	return 0
}

func (x *NodeExecStats) GetReferencedTensor() []*allocation_description_go_proto.AllocationDescription {
	if x != nil {
		return x.ReferencedTensor
	}
	return nil
}

func (x *NodeExecStats) GetMemoryStats() *MemoryStats {
	if x != nil {
		return x.MemoryStats
	}
	return nil
}

func (x *NodeExecStats) GetAllStartNanos() int64 {
	if x != nil {
		return x.AllStartNanos
	}
	return 0
}

func (x *NodeExecStats) GetOpStartRelNanos() int64 {
	if x != nil {
		return x.OpStartRelNanos
	}
	return 0
}

func (x *NodeExecStats) GetOpEndRelNanos() int64 {
	if x != nil {
		return x.OpEndRelNanos
	}
	return 0
}

func (x *NodeExecStats) GetAllEndRelNanos() int64 {
	if x != nil {
		return x.AllEndRelNanos
	}
	return 0
}

func (x *NodeExecStats) GetScheduledNanos() int64 {
	if x != nil {
		return x.ScheduledNanos
	}
	return 0
}

type DeviceStepStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device    string           `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	NodeStats []*NodeExecStats `protobuf:"bytes,2,rep,name=node_stats,json=nodeStats,proto3" json:"node_stats,omitempty"`
	// Its key is thread id.
	ThreadNames map[uint32]string `protobuf:"bytes,3,rep,name=thread_names,json=threadNames,proto3" json:"thread_names,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeviceStepStats) Reset() {
	*x = DeviceStepStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStepStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStepStats) ProtoMessage() {}

func (x *DeviceStepStats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStepStats.ProtoReflect.Descriptor instead.
func (*DeviceStepStats) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceStepStats) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *DeviceStepStats) GetNodeStats() []*NodeExecStats {
	if x != nil {
		return x.NodeStats
	}
	return nil
}

func (x *DeviceStepStats) GetThreadNames() map[uint32]string {
	if x != nil {
		return x.ThreadNames
	}
	return nil
}

type StepStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DevStats []*DeviceStepStats `protobuf:"bytes,1,rep,name=dev_stats,json=devStats,proto3" json:"dev_stats,omitempty"`
}

func (x *StepStats) Reset() {
	*x = StepStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepStats) ProtoMessage() {}

func (x *StepStats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_step_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepStats.ProtoReflect.Descriptor instead.
func (*StepStats) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_step_stats_proto_rawDescGZIP(), []int{6}
}

func (x *StepStats) GetDevStats() []*DeviceStepStats {
	if x != nil {
		return x.DevStats
	}
	return nil
}

var File_tensorflow_core_framework_step_stats_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_step_stats_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x74, 0x65, 0x70,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x36, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x9d, 0x02, 0x0a,
	0x13, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x65, 0x61, 0x6b, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x65, 0x61, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x69, 0x76, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x12, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x75, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x42, 0x79, 0x74, 0x65, 0x73, 0x49, 0x6e, 0x55, 0x73, 0x65, 0x22, 0x6e, 0x0a, 0x0a,
	0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x4c,
	0x0a, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfe, 0x02, 0x0a,
	0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x1b,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x18, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x17, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x45, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x1a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4e, 0x0a,
	0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x49, 0x64, 0x73, 0x22, 0x93, 0x06,
	0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x6f, 0x70, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x6c, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x72, 0x65, 0x6c, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6f, 0x70, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x12, 0x2b, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x6c,
	0x6c, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x37, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x64, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x70, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x6c, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x27, 0x0a, 0x10, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x72, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6f, 0x70, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12,
	0x29, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x45,
	0x6e, 0x64, 0x52, 0x65, 0x6c, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4e, 0x61,
	0x6e, 0x6f, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09,
	0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x09, 0x53, 0x74,
	0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x08, 0x64, 0x65, 0x76, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x42, 0x83, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0f,
	0x53, 0x74, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50,
	0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_step_stats_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_step_stats_proto_rawDescData = file_tensorflow_core_framework_step_stats_proto_rawDesc
)

func file_tensorflow_core_framework_step_stats_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_step_stats_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_step_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_step_stats_proto_rawDescData)
	})
	return file_tensorflow_core_framework_step_stats_proto_rawDescData
}

var file_tensorflow_core_framework_step_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tensorflow_core_framework_step_stats_proto_goTypes = []interface{}{
	(*AllocationRecord)(nil),    // 0: tensorflow.AllocationRecord
	(*AllocatorMemoryUsed)(nil), // 1: tensorflow.AllocatorMemoryUsed
	(*NodeOutput)(nil),          // 2: tensorflow.NodeOutput
	(*MemoryStats)(nil),         // 3: tensorflow.MemoryStats
	(*NodeExecStats)(nil),       // 4: tensorflow.NodeExecStats
	(*DeviceStepStats)(nil),     // 5: tensorflow.DeviceStepStats
	(*StepStats)(nil),           // 6: tensorflow.StepStats
	nil,                         // 7: tensorflow.DeviceStepStats.ThreadNamesEntry
	(*tensor_description_go_proto.TensorDescription)(nil),         // 8: tensorflow.TensorDescription
	(*allocation_description_go_proto.AllocationDescription)(nil), // 9: tensorflow.AllocationDescription
}
var file_tensorflow_core_framework_step_stats_proto_depIdxs = []int32{
	0, // 0: tensorflow.AllocatorMemoryUsed.allocation_records:type_name -> tensorflow.AllocationRecord
	8, // 1: tensorflow.NodeOutput.tensor_description:type_name -> tensorflow.TensorDescription
	1, // 2: tensorflow.NodeExecStats.memory:type_name -> tensorflow.AllocatorMemoryUsed
	2, // 3: tensorflow.NodeExecStats.output:type_name -> tensorflow.NodeOutput
	9, // 4: tensorflow.NodeExecStats.referenced_tensor:type_name -> tensorflow.AllocationDescription
	3, // 5: tensorflow.NodeExecStats.memory_stats:type_name -> tensorflow.MemoryStats
	4, // 6: tensorflow.DeviceStepStats.node_stats:type_name -> tensorflow.NodeExecStats
	7, // 7: tensorflow.DeviceStepStats.thread_names:type_name -> tensorflow.DeviceStepStats.ThreadNamesEntry
	5, // 8: tensorflow.StepStats.dev_stats:type_name -> tensorflow.DeviceStepStats
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_step_stats_proto_init() }
func file_tensorflow_core_framework_step_stats_proto_init() {
	if File_tensorflow_core_framework_step_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_step_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationRecord); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocatorMemoryUsed); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeOutput); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryStats); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExecStats); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceStepStats); i {
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
		file_tensorflow_core_framework_step_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepStats); i {
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
			RawDescriptor: file_tensorflow_core_framework_step_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_step_stats_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_step_stats_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_step_stats_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_step_stats_proto = out.File
	file_tensorflow_core_framework_step_stats_proto_rawDesc = nil
	file_tensorflow_core_framework_step_stats_proto_goTypes = nil
	file_tensorflow_core_framework_step_stats_proto_depIdxs = nil
}
