// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: map.proto

package grpc

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{0}
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Col        int32  `protobuf:"varint,2,opt,name=col,proto3" json:"col,omitempty"`
	Row        int32  `protobuf:"varint,3,opt,name=row,proto3" json:"row,omitempty"`
	Direction  string `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
	Production string `protobuf:"bytes,5,opt,name=production,proto3" json:"production,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Node) GetCol() int32 {
	if x != nil {
		return x.Col
	}
	return 0
}

func (x *Node) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *Node) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *Node) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuLoad     float32 `protobuf:"fixed32,1,opt,name=cpu_load,json=cpuLoad,proto3" json:"cpu_load,omitempty"`
	MemoryUsage int32   `protobuf:"varint,2,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	NetworkRx   int32   `protobuf:"varint,3,opt,name=network_rx,json=networkRx,proto3" json:"network_rx,omitempty"`
	NetworkTx   int32   `protobuf:"varint,4,opt,name=network_tx,json=networkTx,proto3" json:"network_tx,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetCpuLoad() float32 {
	if x != nil {
		return x.CpuLoad
	}
	return 0
}

func (x *Stats) GetMemoryUsage() int32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Stats) GetNetworkRx() int32 {
	if x != nil {
		return x.NetworkRx
	}
	return 0
}

func (x *Stats) GetNetworkTx() int32 {
	if x != nil {
		return x.NetworkTx
	}
	return 0
}

type NodeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node      *Node          `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Items     []*ItemCounter `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	NodeStats *Stats         `protobuf:"bytes,3,opt,name=node_stats,json=nodeStats,proto3" json:"node_stats,omitempty"`
}

func (x *NodeState) Reset() {
	*x = NodeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeState) ProtoMessage() {}

func (x *NodeState) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeState.ProtoReflect.Descriptor instead.
func (*NodeState) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{3}
}

func (x *NodeState) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *NodeState) GetItems() []*ItemCounter {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NodeState) GetNodeStats() *Stats {
	if x != nil {
		return x.NodeStats
	}
	return nil
}

type ItemCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ItemCounter) Reset() {
	*x = ItemCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCounter) ProtoMessage() {}

func (x *ItemCounter) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCounter.ProtoReflect.Descriptor instead.
func (*ItemCounter) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{4}
}

func (x *ItemCounter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ItemCounter) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id          string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Parents     []string `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	Ingredients []*Item  `protobuf:"bytes,4,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{5}
}

func (x *Item) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *Item) GetIngredients() []*Item {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{6}
}

func (x *ItemList) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type NodesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodesList) Reset() {
	*x = NodesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesList) ProtoMessage() {}

func (x *NodesList) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesList.ProtoReflect.Descriptor instead.
func (*NodesList) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{7}
}

func (x *NodesList) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_map_proto protoreflect.FileDescriptor

var file_map_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x7c, 0x0a, 0x04, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63, 0x70, 0x75, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x72, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x78, 0x22, 0x80,
	0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0x37, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x2c, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2c,
	0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2d, 0x0a, 0x09,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x38, 0x0a, 0x03, 0x4d,
	0x61, 0x70, 0x12, 0x31, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_map_proto_rawDescOnce sync.Once
	file_map_proto_rawDescData = file_map_proto_rawDesc
)

func file_map_proto_rawDescGZIP() []byte {
	file_map_proto_rawDescOnce.Do(func() {
		file_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_map_proto_rawDescData)
	})
	return file_map_proto_rawDescData
}

var file_map_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_map_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: grpc.Empty
	(*Node)(nil),        // 1: grpc.Node
	(*Stats)(nil),       // 2: grpc.Stats
	(*NodeState)(nil),   // 3: grpc.NodeState
	(*ItemCounter)(nil), // 4: grpc.ItemCounter
	(*Item)(nil),        // 5: grpc.Item
	(*ItemList)(nil),    // 6: grpc.ItemList
	(*NodesList)(nil),   // 7: grpc.NodesList
}
var file_map_proto_depIdxs = []int32{
	1, // 0: grpc.NodeState.node:type_name -> grpc.Node
	4, // 1: grpc.NodeState.items:type_name -> grpc.ItemCounter
	2, // 2: grpc.NodeState.node_stats:type_name -> grpc.Stats
	5, // 3: grpc.Item.ingredients:type_name -> grpc.Item
	5, // 4: grpc.ItemList.items:type_name -> grpc.Item
	1, // 5: grpc.NodesList.nodes:type_name -> grpc.Node
	3, // 6: grpc.Map.updateNodeState:input_type -> grpc.NodeState
	0, // 7: grpc.Map.updateNodeState:output_type -> grpc.Empty
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_map_proto_init() }
func file_map_proto_init() {
	if File_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_map_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_map_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeState); i {
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
		file_map_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCounter); i {
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
		file_map_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_map_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemList); i {
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
		file_map_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesList); i {
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
			RawDescriptor: file_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_map_proto_goTypes,
		DependencyIndexes: file_map_proto_depIdxs,
		MessageInfos:      file_map_proto_msgTypes,
	}.Build()
	File_map_proto = out.File
	file_map_proto_rawDesc = nil
	file_map_proto_goTypes = nil
	file_map_proto_depIdxs = nil
}
