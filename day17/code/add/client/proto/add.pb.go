// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/add.proto

package proto

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

// Op 枚举
type Op int32

const (
	Op__   Op = 0
	Op_ADD Op = 1
	Op_SUB Op = 2
)

// Enum value maps for Op.
var (
	Op_name = map[int32]string{
		0: "_",
		1: "ADD",
		2: "SUB",
	}
	Op_value = map[string]int32{
		"_":   0,
		"ADD": 1,
		"SUB": 2,
	}
)

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}

func (x Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Op) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_add_proto_enumTypes[0].Descriptor()
}

func (Op) Type() protoreflect.EnumType {
	return &file_proto_add_proto_enumTypes[0]
}

func (x Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Op.Descriptor instead.
func (Op) EnumDescriptor() ([]byte, []int) {
	return file_proto_add_proto_rawDescGZIP(), []int{0}
}

// Req 请求服务的参数
type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X  int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y  int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Op Op    `protobuf:"varint,3,opt,name=op,proto3,enum=proto.Op" json:"op,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_proto_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_proto_add_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Req) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Req) GetOp() Op {
	if x != nil {
		return x.Op
	}
	return Op__
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_proto_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_proto_add_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_proto_add_proto protoreflect.FileDescriptor

var file_proto_add_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x12, 0x19, 0x0a, 0x02, 0x6f,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x17, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x2a,
	0x1d, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x42, 0x10, 0x02, 0x32, 0x25,
	0x0a, 0x03, 0x63, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x02, 0x44, 0x6f, 0x12, 0x0a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x64, 0x64, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_add_proto_rawDescOnce sync.Once
	file_proto_add_proto_rawDescData = file_proto_add_proto_rawDesc
)

func file_proto_add_proto_rawDescGZIP() []byte {
	file_proto_add_proto_rawDescOnce.Do(func() {
		file_proto_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_add_proto_rawDescData)
	})
	return file_proto_add_proto_rawDescData
}

var file_proto_add_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_add_proto_goTypes = []interface{}{
	(Op)(0),     // 0: proto.Op
	(*Req)(nil), // 1: proto.Req
	(*Res)(nil), // 2: proto.Res
}
var file_proto_add_proto_depIdxs = []int32{
	0, // 0: proto.Req.op:type_name -> proto.Op
	1, // 1: proto.cal.Do:input_type -> proto.Req
	2, // 2: proto.cal.Do:output_type -> proto.Res
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_add_proto_init() }
func file_proto_add_proto_init() {
	if File_proto_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_proto_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_proto_add_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_add_proto_goTypes,
		DependencyIndexes: file_proto_add_proto_depIdxs,
		EnumInfos:         file_proto_add_proto_enumTypes,
		MessageInfos:      file_proto_add_proto_msgTypes,
	}.Build()
	File_proto_add_proto = out.File
	file_proto_add_proto_rawDesc = nil
	file_proto_add_proto_goTypes = nil
	file_proto_add_proto_depIdxs = nil
}
