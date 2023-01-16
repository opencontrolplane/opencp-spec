// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: domain.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DomainStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DomainStatus) Reset() {
	*x = DomainStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainStatus) ProtoMessage() {}

func (x *DomainStatus) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainStatus.ProtoReflect.Descriptor instead.
func (*DomainStatus) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{0}
}

func (x *DomainStatus) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type Records struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Priority *int32 `protobuf:"varint,4,opt,name=Priority,proto3,oneof" json:"Priority,omitempty"`
	TTL      *int32 `protobuf:"varint,5,opt,name=TTL,proto3,oneof" json:"TTL,omitempty"`
}

func (x *Records) Reset() {
	*x = Records{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Records) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Records) ProtoMessage() {}

func (x *Records) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Records.ProtoReflect.Descriptor instead.
func (*Records) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{1}
}

func (x *Records) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Records) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Records) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Records) GetPriority() int32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *Records) GetTTL() int32 {
	if x != nil && x.TTL != nil {
		return *x.TTL
	}
	return 0
}

type DomainSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Records `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *DomainSpec) Reset() {
	*x = DomainSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainSpec) ProtoMessage() {}

func (x *DomainSpec) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainSpec.ProtoReflect.Descriptor instead.
func (*DomainSpec) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{2}
}

func (x *DomainSpec) GetRecords() []*Records {
	if x != nil {
		return x.Records
	}
	return nil
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// Spec defines the behavior of a domain.
	// +optional
	Spec *DomainSpec `protobuf:"bytes,4,opt,name=spec,proto3,oneof" json:"spec,omitempty"`
	// Status represents the current information about a domain. This data may be out of date.
	// +optional
	Status *DomainStatus `protobuf:"bytes,5,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{3}
}

func (x *Domain) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Domain) GetSpec() *DomainSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Domain) GetStatus() *DomainStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DomainList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,4,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// List of domains
	Items []*Domain `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DomainList) Reset() {
	*x = DomainList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainList) ProtoMessage() {}

func (x *DomainList) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainList.ProtoReflect.Descriptor instead.
func (*DomainList) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{4}
}

func (x *DomainList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DomainList) GetItems() []*Domain {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_domain_proto protoreflect.FileDescriptor

var file_domain_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x07, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x54, 0x54, 0x4c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x54, 0x54, 0x4c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x54, 0x54,
	0x4c, 0x22, 0x40, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x32, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x51,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x48, 0x01, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x9c, 0x02, 0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x17, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x17, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x1b, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x70, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_proto_rawDescOnce sync.Once
	file_domain_proto_rawDescData = file_domain_proto_rawDesc
)

func file_domain_proto_rawDescGZIP() []byte {
	file_domain_proto_rawDescOnce.Do(func() {
		file_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_proto_rawDescData)
	})
	return file_domain_proto_rawDescData
}

var file_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_domain_proto_goTypes = []interface{}{
	(*DomainStatus)(nil),  // 0: domain.v1alpha1.DomainStatus
	(*Records)(nil),       // 1: domain.v1alpha1.Records
	(*DomainSpec)(nil),    // 2: domain.v1alpha1.DomainSpec
	(*Domain)(nil),        // 3: domain.v1alpha1.Domain
	(*DomainList)(nil),    // 4: domain.v1alpha1.DomainList
	(*v1.ObjectMeta)(nil), // 5: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.ListMeta)(nil),   // 6: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	(*FilterOptions)(nil), // 7: options.FilterOptions
}
var file_domain_proto_depIdxs = []int32{
	1,  // 0: domain.v1alpha1.DomainSpec.records:type_name -> domain.v1alpha1.Records
	5,  // 1: domain.v1alpha1.Domain.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	2,  // 2: domain.v1alpha1.Domain.spec:type_name -> domain.v1alpha1.DomainSpec
	0,  // 3: domain.v1alpha1.Domain.status:type_name -> domain.v1alpha1.DomainStatus
	6,  // 4: domain.v1alpha1.DomainList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	3,  // 5: domain.v1alpha1.DomainList.items:type_name -> domain.v1alpha1.Domain
	3,  // 6: domain.v1alpha1.DomainService.CreateDomain:input_type -> domain.v1alpha1.Domain
	7,  // 7: domain.v1alpha1.DomainService.GetDomain:input_type -> options.FilterOptions
	7,  // 8: domain.v1alpha1.DomainService.DeleteDomain:input_type -> options.FilterOptions
	7,  // 9: domain.v1alpha1.DomainService.ListDomains:input_type -> options.FilterOptions
	3,  // 10: domain.v1alpha1.DomainService.CreateDomain:output_type -> domain.v1alpha1.Domain
	3,  // 11: domain.v1alpha1.DomainService.GetDomain:output_type -> domain.v1alpha1.Domain
	3,  // 12: domain.v1alpha1.DomainService.DeleteDomain:output_type -> domain.v1alpha1.Domain
	4,  // 13: domain.v1alpha1.DomainService.ListDomains:output_type -> domain.v1alpha1.DomainList
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_domain_proto_init() }
func file_domain_proto_init() {
	if File_domain_proto != nil {
		return
	}
	file_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainStatus); i {
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
		file_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Records); i {
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
		file_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainSpec); i {
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
		file_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainList); i {
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
	file_domain_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_domain_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_domain_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_proto_goTypes,
		DependencyIndexes: file_domain_proto_depIdxs,
		MessageInfos:      file_domain_proto_msgTypes,
	}.Build()
	File_domain_proto = out.File
	file_domain_proto_rawDesc = nil
	file_domain_proto_goTypes = nil
	file_domain_proto_depIdxs = nil
}