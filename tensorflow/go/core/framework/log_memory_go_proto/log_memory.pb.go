// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: tensorflow/core/framework/log_memory.proto

package log_memory_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

type MemoryLogStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Handle describing the feeds and fetches of the step.
	Handle string `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
}

func (x *MemoryLogStep) Reset() {
	*x = MemoryLogStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogStep) ProtoMessage() {}

func (x *MemoryLogStep) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogStep.ProtoReflect.Descriptor instead.
func (*MemoryLogStep) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{0}
}

func (x *MemoryLogStep) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

func (x *MemoryLogStep) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

type MemoryLogTensorAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the kernel making the allocation as set in GraphDef,
	// e.g., "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName,proto3" json:"kernel_name,omitempty"`
	// Allocated tensor details.
	Tensor *tensor_description_go_proto.TensorDescription `protobuf:"bytes,3,opt,name=tensor,proto3" json:"tensor,omitempty"`
}

func (x *MemoryLogTensorAllocation) Reset() {
	*x = MemoryLogTensorAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogTensorAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogTensorAllocation) ProtoMessage() {}

func (x *MemoryLogTensorAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogTensorAllocation.ProtoReflect.Descriptor instead.
func (*MemoryLogTensorAllocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{1}
}

func (x *MemoryLogTensorAllocation) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

func (x *MemoryLogTensorAllocation) GetKernelName() string {
	if x != nil {
		return x.KernelName
	}
	return ""
}

func (x *MemoryLogTensorAllocation) GetTensor() *tensor_description_go_proto.TensorDescription {
	if x != nil {
		return x.Tensor
	}
	return nil
}

type MemoryLogTensorDeallocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,1,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,2,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
}

func (x *MemoryLogTensorDeallocation) Reset() {
	*x = MemoryLogTensorDeallocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogTensorDeallocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogTensorDeallocation) ProtoMessage() {}

func (x *MemoryLogTensorDeallocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogTensorDeallocation.ProtoReflect.Descriptor instead.
func (*MemoryLogTensorDeallocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{2}
}

func (x *MemoryLogTensorDeallocation) GetAllocationId() int64 {
	if x != nil {
		return x.AllocationId
	}
	return 0
}

func (x *MemoryLogTensorDeallocation) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

type MemoryLogTensorOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the kernel producing an output as set in GraphDef, e.g.,
	// "affine2/weights/Assign".
	KernelName string `protobuf:"bytes,2,opt,name=kernel_name,json=kernelName,proto3" json:"kernel_name,omitempty"`
	// Index of the output being set.
	Index int32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// Output tensor details.
	Tensor *tensor_description_go_proto.TensorDescription `protobuf:"bytes,4,opt,name=tensor,proto3" json:"tensor,omitempty"`
}

func (x *MemoryLogTensorOutput) Reset() {
	*x = MemoryLogTensorOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogTensorOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogTensorOutput) ProtoMessage() {}

func (x *MemoryLogTensorOutput) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogTensorOutput.ProtoReflect.Descriptor instead.
func (*MemoryLogTensorOutput) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{3}
}

func (x *MemoryLogTensorOutput) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

func (x *MemoryLogTensorOutput) GetKernelName() string {
	if x != nil {
		return x.KernelName
	}
	return ""
}

func (x *MemoryLogTensorOutput) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *MemoryLogTensorOutput) GetTensor() *tensor_description_go_proto.TensorDescription {
	if x != nil {
		return x.Tensor
	}
	return nil
}

type MemoryLogRawAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the operation making the allocation.
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Number of bytes in the allocation.
	NumBytes int64 `protobuf:"varint,3,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,4,opt,name=ptr,proto3" json:"ptr,omitempty"`
	// Id of the tensor buffer being allocated, used to match to a
	// corresponding deallocation.
	AllocationId int64 `protobuf:"varint,5,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,6,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
}

func (x *MemoryLogRawAllocation) Reset() {
	*x = MemoryLogRawAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogRawAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogRawAllocation) ProtoMessage() {}

func (x *MemoryLogRawAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogRawAllocation.ProtoReflect.Descriptor instead.
func (*MemoryLogRawAllocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{4}
}

func (x *MemoryLogRawAllocation) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

func (x *MemoryLogRawAllocation) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *MemoryLogRawAllocation) GetNumBytes() int64 {
	if x != nil {
		return x.NumBytes
	}
	return 0
}

func (x *MemoryLogRawAllocation) GetPtr() uint64 {
	if x != nil {
		return x.Ptr
	}
	return 0
}

func (x *MemoryLogRawAllocation) GetAllocationId() int64 {
	if x != nil {
		return x.AllocationId
	}
	return 0
}

func (x *MemoryLogRawAllocation) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

type MemoryLogRawDeallocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Process-unique step id.
	StepId int64 `protobuf:"varint,1,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	// Name of the operation making the deallocation.
	Operation string `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// Id of the tensor buffer being deallocated, used to match to a
	// corresponding allocation.
	AllocationId int64 `protobuf:"varint,3,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Name of the allocator used.
	AllocatorName string `protobuf:"bytes,4,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// True if the deallocation is queued and will be performed later,
	// e.g. for GPU lazy freeing of buffers.
	Deferred bool `protobuf:"varint,5,opt,name=deferred,proto3" json:"deferred,omitempty"`
}

func (x *MemoryLogRawDeallocation) Reset() {
	*x = MemoryLogRawDeallocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryLogRawDeallocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryLogRawDeallocation) ProtoMessage() {}

func (x *MemoryLogRawDeallocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_log_memory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryLogRawDeallocation.ProtoReflect.Descriptor instead.
func (*MemoryLogRawDeallocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_log_memory_proto_rawDescGZIP(), []int{5}
}

func (x *MemoryLogRawDeallocation) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

func (x *MemoryLogRawDeallocation) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *MemoryLogRawDeallocation) GetAllocationId() int64 {
	if x != nil {
		return x.AllocationId
	}
	return 0
}

func (x *MemoryLogRawDeallocation) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

func (x *MemoryLogRawDeallocation) GetDeferred() bool {
	if x != nil {
		return x.Deferred
	}
	return false
}

var File_tensorflow_core_framework_log_memory_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_log_memory_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x5f,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0d,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x65, 0x70, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x8c,
	0x01, 0x0a, 0x19, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x69, 0x0a,
	0x1b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6b,
	0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0xca, 0x01, 0x0a, 0x16, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x61, 0x77, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x75, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6e, 0x75, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x74, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4c, 0x6f, 0x67, 0x52, 0x61, 0x77, 0x44, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x64, 0x42, 0x83, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42,
	0x0f, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_log_memory_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_log_memory_proto_rawDescData = file_tensorflow_core_framework_log_memory_proto_rawDesc
)

func file_tensorflow_core_framework_log_memory_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_log_memory_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_log_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_log_memory_proto_rawDescData)
	})
	return file_tensorflow_core_framework_log_memory_proto_rawDescData
}

var file_tensorflow_core_framework_log_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tensorflow_core_framework_log_memory_proto_goTypes = []interface{}{
	(*MemoryLogStep)(nil),                                 // 0: tensorflow.MemoryLogStep
	(*MemoryLogTensorAllocation)(nil),                     // 1: tensorflow.MemoryLogTensorAllocation
	(*MemoryLogTensorDeallocation)(nil),                   // 2: tensorflow.MemoryLogTensorDeallocation
	(*MemoryLogTensorOutput)(nil),                         // 3: tensorflow.MemoryLogTensorOutput
	(*MemoryLogRawAllocation)(nil),                        // 4: tensorflow.MemoryLogRawAllocation
	(*MemoryLogRawDeallocation)(nil),                      // 5: tensorflow.MemoryLogRawDeallocation
	(*tensor_description_go_proto.TensorDescription)(nil), // 6: tensorflow.TensorDescription
}
var file_tensorflow_core_framework_log_memory_proto_depIdxs = []int32{
	6, // 0: tensorflow.MemoryLogTensorAllocation.tensor:type_name -> tensorflow.TensorDescription
	6, // 1: tensorflow.MemoryLogTensorOutput.tensor:type_name -> tensorflow.TensorDescription
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_log_memory_proto_init() }
func file_tensorflow_core_framework_log_memory_proto_init() {
	if File_tensorflow_core_framework_log_memory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_log_memory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogStep); i {
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
		file_tensorflow_core_framework_log_memory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogTensorAllocation); i {
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
		file_tensorflow_core_framework_log_memory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogTensorDeallocation); i {
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
		file_tensorflow_core_framework_log_memory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogTensorOutput); i {
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
		file_tensorflow_core_framework_log_memory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogRawAllocation); i {
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
		file_tensorflow_core_framework_log_memory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryLogRawDeallocation); i {
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
			RawDescriptor: file_tensorflow_core_framework_log_memory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_log_memory_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_log_memory_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_log_memory_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_log_memory_proto = out.File
	file_tensorflow_core_framework_log_memory_proto_rawDesc = nil
	file_tensorflow_core_framework_log_memory_proto_goTypes = nil
	file_tensorflow_core_framework_log_memory_proto_depIdxs = nil
}
