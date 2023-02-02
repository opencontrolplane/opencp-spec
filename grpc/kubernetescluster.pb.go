// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: kubernetescluster.proto

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

type KubernetesCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	Metadata *v1.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// The spec of the KubernetesCluster.
	Spec *KubernetesClusterSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// The status of the KubernetesCluster.
	Status *KubernetesClusterStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *KubernetesCluster) Reset() {
	*x = KubernetesCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetescluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesCluster) ProtoMessage() {}

func (x *KubernetesCluster) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetescluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesCluster.ProtoReflect.Descriptor instead.
func (*KubernetesCluster) Descriptor() ([]byte, []int) {
	return file_kubernetescluster_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesCluster) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *KubernetesCluster) GetSpec() *KubernetesClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *KubernetesCluster) GetStatus() *KubernetesClusterStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type KubernetesClusterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pools       []*KubernetesClusterPool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	Version     string                   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Firewall    string                   `protobuf:"bytes,3,opt,name=firewall,proto3" json:"firewall,omitempty"`
	Cniplugin   string                   `protobuf:"bytes,4,opt,name=cniplugin,json=cni_plugin,proto3" json:"cniplugin,omitempty"`
	Clustertype string                   `protobuf:"bytes,5,opt,name=clustertype,json=cluster_type,proto3" json:"clustertype,omitempty"`
}

func (x *KubernetesClusterSpec) Reset() {
	*x = KubernetesClusterSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetescluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterSpec) ProtoMessage() {}

func (x *KubernetesClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetescluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterSpec.ProtoReflect.Descriptor instead.
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return file_kubernetescluster_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesClusterSpec) GetPools() []*KubernetesClusterPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *KubernetesClusterSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *KubernetesClusterSpec) GetFirewall() string {
	if x != nil {
		return x.Firewall
	}
	return ""
}

func (x *KubernetesClusterSpec) GetCniplugin() string {
	if x != nil {
		return x.Cniplugin
	}
	return ""
}

func (x *KubernetesClusterSpec) GetClustertype() string {
	if x != nil {
		return x.Clustertype
	}
	return ""
}

type KubernetesClusterPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size        string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Count       int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Autoscaling bool   `protobuf:"varint,4,opt,name=autoscaling,proto3" json:"autoscaling,omitempty"`
	Minsize     *int64 `protobuf:"varint,5,opt,name=minsize,json=min_size,proto3,oneof" json:"minsize,omitempty"`
	Maxsize     *int64 `protobuf:"varint,6,opt,name=maxsize,json=max_size,proto3,oneof" json:"maxsize,omitempty"`
}

func (x *KubernetesClusterPool) Reset() {
	*x = KubernetesClusterPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetescluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterPool) ProtoMessage() {}

func (x *KubernetesClusterPool) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetescluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterPool.ProtoReflect.Descriptor instead.
func (*KubernetesClusterPool) Descriptor() ([]byte, []int) {
	return file_kubernetescluster_proto_rawDescGZIP(), []int{2}
}

func (x *KubernetesClusterPool) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KubernetesClusterPool) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *KubernetesClusterPool) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *KubernetesClusterPool) GetAutoscaling() bool {
	if x != nil {
		return x.Autoscaling
	}
	return false
}

func (x *KubernetesClusterPool) GetMinsize() int64 {
	if x != nil && x.Minsize != nil {
		return *x.Minsize
	}
	return 0
}

func (x *KubernetesClusterPool) GetMaxsize() int64 {
	if x != nil && x.Maxsize != nil {
		return *x.Maxsize
	}
	return 0
}

type KubernetesClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Publicip string `protobuf:"bytes,3,opt,name=publicip,json=public_ip,proto3" json:"publicip,omitempty"`
}

func (x *KubernetesClusterStatus) Reset() {
	*x = KubernetesClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetescluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterStatus) ProtoMessage() {}

func (x *KubernetesClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetescluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterStatus.ProtoReflect.Descriptor instead.
func (*KubernetesClusterStatus) Descriptor() ([]byte, []int) {
	return file_kubernetescluster_proto_rawDescGZIP(), []int{3}
}

func (x *KubernetesClusterStatus) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *KubernetesClusterStatus) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *KubernetesClusterStatus) GetPublicip() string {
	if x != nil {
		return x.Publicip
	}
	return ""
}

// KubernetesClusterList is a list of KubernetesCluster.
type KubernetesClusterList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard list metadata.
	// +optional
	Metadata *v1.ListMeta `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	// Items is the list of VistualMachine.
	Items []*KubernetesCluster `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *KubernetesClusterList) Reset() {
	*x = KubernetesClusterList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubernetescluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesClusterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterList) ProtoMessage() {}

func (x *KubernetesClusterList) ProtoReflect() protoreflect.Message {
	mi := &file_kubernetescluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterList.ProtoReflect.Descriptor instead.
func (*KubernetesClusterList) Descriptor() ([]byte, []int) {
	return file_kubernetescluster_proto_rawDescGZIP(), []int{4}
}

func (x *KubernetesClusterList) GetMetadata() *v1.ListMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *KubernetesClusterList) GetItems() []*KubernetesCluster {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_kubernetescluster_proto protoreflect.FileDescriptor

var file_kubernetescluster_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x34, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x11, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xd8, 0x01, 0x0a, 0x15, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x47,
	0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x12, 0x1d, 0x0a,
	0x09, 0x63, 0x6e, 0x69, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6e, 0x69, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xcb, 0x01, 0x0a, 0x15, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x7a,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x78, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x68, 0x0a,
	0x17, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x22, 0xba, 0x01, 0x0a, 0x15, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x12, 0x43, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xbb, 0x04, 0x0a, 0x18, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x79, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x2d, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x2d, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x31,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2d,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x2d, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x62,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x2d, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x70, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubernetescluster_proto_rawDescOnce sync.Once
	file_kubernetescluster_proto_rawDescData = file_kubernetescluster_proto_rawDesc
)

func file_kubernetescluster_proto_rawDescGZIP() []byte {
	file_kubernetescluster_proto_rawDescOnce.Do(func() {
		file_kubernetescluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubernetescluster_proto_rawDescData)
	})
	return file_kubernetescluster_proto_rawDescData
}

var file_kubernetescluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kubernetescluster_proto_goTypes = []interface{}{
	(*KubernetesCluster)(nil),       // 0: kubernetescluster.v1alpha1.KubernetesCluster
	(*KubernetesClusterSpec)(nil),   // 1: kubernetescluster.v1alpha1.KubernetesClusterSpec
	(*KubernetesClusterPool)(nil),   // 2: kubernetescluster.v1alpha1.KubernetesClusterPool
	(*KubernetesClusterStatus)(nil), // 3: kubernetescluster.v1alpha1.KubernetesClusterStatus
	(*KubernetesClusterList)(nil),   // 4: kubernetescluster.v1alpha1.KubernetesClusterList
	(*v1.ObjectMeta)(nil),           // 5: k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	(*v1.ListMeta)(nil),             // 6: k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	(*FilterOptions)(nil),           // 7: options.FilterOptions
}
var file_kubernetescluster_proto_depIdxs = []int32{
	5,  // 0: kubernetescluster.v1alpha1.KubernetesCluster.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta
	1,  // 1: kubernetescluster.v1alpha1.KubernetesCluster.spec:type_name -> kubernetescluster.v1alpha1.KubernetesClusterSpec
	3,  // 2: kubernetescluster.v1alpha1.KubernetesCluster.status:type_name -> kubernetescluster.v1alpha1.KubernetesClusterStatus
	2,  // 3: kubernetescluster.v1alpha1.KubernetesClusterSpec.pools:type_name -> kubernetescluster.v1alpha1.KubernetesClusterPool
	6,  // 4: kubernetescluster.v1alpha1.KubernetesClusterList.metadata:type_name -> k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta
	0,  // 5: kubernetescluster.v1alpha1.KubernetesClusterList.items:type_name -> kubernetescluster.v1alpha1.KubernetesCluster
	0,  // 6: kubernetescluster.v1alpha1.KubernetesClusterService.CreateKubernetesCluster:input_type -> kubernetescluster.v1alpha1.KubernetesCluster
	7,  // 7: kubernetescluster.v1alpha1.KubernetesClusterService.GetKubernetesCluster:input_type -> options.FilterOptions
	7,  // 8: kubernetescluster.v1alpha1.KubernetesClusterService.ListKubernetesCluster:input_type -> options.FilterOptions
	0,  // 9: kubernetescluster.v1alpha1.KubernetesClusterService.UpdateKubernetesCluster:input_type -> kubernetescluster.v1alpha1.KubernetesCluster
	7,  // 10: kubernetescluster.v1alpha1.KubernetesClusterService.DeleteKubernetesCluster:input_type -> options.FilterOptions
	0,  // 11: kubernetescluster.v1alpha1.KubernetesClusterService.CreateKubernetesCluster:output_type -> kubernetescluster.v1alpha1.KubernetesCluster
	0,  // 12: kubernetescluster.v1alpha1.KubernetesClusterService.GetKubernetesCluster:output_type -> kubernetescluster.v1alpha1.KubernetesCluster
	4,  // 13: kubernetescluster.v1alpha1.KubernetesClusterService.ListKubernetesCluster:output_type -> kubernetescluster.v1alpha1.KubernetesClusterList
	0,  // 14: kubernetescluster.v1alpha1.KubernetesClusterService.UpdateKubernetesCluster:output_type -> kubernetescluster.v1alpha1.KubernetesCluster
	0,  // 15: kubernetescluster.v1alpha1.KubernetesClusterService.DeleteKubernetesCluster:output_type -> kubernetescluster.v1alpha1.KubernetesCluster
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_kubernetescluster_proto_init() }
func file_kubernetescluster_proto_init() {
	if File_kubernetescluster_proto != nil {
		return
	}
	file_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kubernetescluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesCluster); i {
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
		file_kubernetescluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterSpec); i {
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
		file_kubernetescluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterPool); i {
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
		file_kubernetescluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterStatus); i {
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
		file_kubernetescluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesClusterList); i {
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
	file_kubernetescluster_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_kubernetescluster_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_kubernetescluster_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kubernetescluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubernetescluster_proto_goTypes,
		DependencyIndexes: file_kubernetescluster_proto_depIdxs,
		MessageInfos:      file_kubernetescluster_proto_msgTypes,
	}.Build()
	File_kubernetescluster_proto = out.File
	file_kubernetescluster_proto_rawDesc = nil
	file_kubernetescluster_proto_goTypes = nil
	file_kubernetescluster_proto_depIdxs = nil
}
