// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: mine.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Col  int32  `protobuf:"varint,2,opt,name=col,proto3" json:"col,omitempty"`
	Row  int32  `protobuf:"varint,3,opt,name=row,proto3" json:"row,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_mine_proto_msgTypes[0]
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
	return file_mine_proto_rawDescGZIP(), []int{0}
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

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_mine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_mine_proto_rawDescGZIP(), []int{1}
}

func (x *URL) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *URL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type MapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes  []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	UrlMap []*URL  `protobuf:"bytes,2,rep,name=url_map,json=urlMap,proto3" json:"url_map,omitempty"`
}

func (x *MapRequest) Reset() {
	*x = MapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRequest) ProtoMessage() {}

func (x *MapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRequest.ProtoReflect.Descriptor instead.
func (*MapRequest) Descriptor() ([]byte, []int) {
	return file_mine_proto_rawDescGZIP(), []int{2}
}

func (x *MapRequest) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *MapRequest) GetUrlMap() []*URL {
	if x != nil {
		return x.UrlMap
	}
	return nil
}

type MapReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MapReply) Reset() {
	*x = MapReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapReply) ProtoMessage() {}

func (x *MapReply) ProtoReflect() protoreflect.Message {
	mi := &file_mine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapReply.ProtoReflect.Descriptor instead.
func (*MapReply) Descriptor() ([]byte, []int) {
	return file_mine_proto_rawDescGZIP(), []int{3}
}

var File_mine_proto protoreflect.FileDescriptor

var file_mine_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x3e, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x6f, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x6f, 0x77, 0x22, 0x2b, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22,
	0x52, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x07, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x06, 0x75, 0x72, 0x6c,
	0x4d, 0x61, 0x70, 0x22, 0x0a, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0x39, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6d, 0x65, 0x73,
	0x67, 0x69, 0x74, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mine_proto_rawDescOnce sync.Once
	file_mine_proto_rawDescData = file_mine_proto_rawDesc
)

func file_mine_proto_rawDescGZIP() []byte {
	file_mine_proto_rawDescOnce.Do(func() {
		file_mine_proto_rawDescData = protoimpl.X.CompressGZIP(file_mine_proto_rawDescData)
	})
	return file_mine_proto_rawDescData
}

var file_mine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mine_proto_goTypes = []interface{}{
	(*Node)(nil),       // 0: grpc.Node
	(*URL)(nil),        // 1: grpc.URL
	(*MapRequest)(nil), // 2: grpc.MapRequest
	(*MapReply)(nil),   // 3: grpc.MapReply
}
var file_mine_proto_depIdxs = []int32{
	0, // 0: grpc.MapRequest.nodes:type_name -> grpc.Node
	1, // 1: grpc.MapRequest.url_map:type_name -> grpc.URL
	2, // 2: grpc.Mapper.updateMap:input_type -> grpc.MapRequest
	3, // 3: grpc.Mapper.updateMap:output_type -> grpc.MapReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mine_proto_init() }
func file_mine_proto_init() {
	if File_mine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_mine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_mine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRequest); i {
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
		file_mine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapReply); i {
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
			RawDescriptor: file_mine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mine_proto_goTypes,
		DependencyIndexes: file_mine_proto_depIdxs,
		MessageInfos:      file_mine_proto_msgTypes,
	}.Build()
	File_mine_proto = out.File
	file_mine_proto_rawDesc = nil
	file_mine_proto_goTypes = nil
	file_mine_proto_depIdxs = nil
}
