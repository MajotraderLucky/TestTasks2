// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.0
// source: books.proto

package protobuf

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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time        *MyTimestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetTime() *MyTimestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Activity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type MyTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *MyTimestamp) Reset() {
	*x = MyTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyTimestamp) ProtoMessage() {}

func (x *MyTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyTimestamp.ProtoReflect.Descriptor instead.
func (*MyTimestamp) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{1}
}

func (x *MyTimestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *MyTimestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{2}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_books_proto_rawDescGZIP(), []int{3}
}

func (x *PingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_books_proto protoreflect.FileDescriptor

var file_books_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x0b,
	0x4d, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0x40, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_books_proto_rawDescOnce sync.Once
	file_books_proto_rawDescData = file_books_proto_rawDesc
)

func file_books_proto_rawDescGZIP() []byte {
	file_books_proto_rawDescOnce.Do(func() {
		file_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_books_proto_rawDescData)
	})
	return file_books_proto_rawDescData
}

var file_books_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_books_proto_goTypes = []interface{}{
	(*Activity)(nil),     // 0: api.v1.Activity
	(*MyTimestamp)(nil),  // 1: api.v1.MyTimestamp
	(*PingRequest)(nil),  // 2: api.v1.PingRequest
	(*PingResponse)(nil), // 3: api.v1.PingResponse
}
var file_books_proto_depIdxs = []int32{
	1, // 0: api.v1.Activity.time:type_name -> api.v1.MyTimestamp
	2, // 1: api.v1.PingService.Ping:input_type -> api.v1.PingRequest
	3, // 2: api.v1.PingService.Ping:output_type -> api.v1.PingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_books_proto_init() }
func file_books_proto_init() {
	if File_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyTimestamp); i {
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
		file_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_books_proto_goTypes,
		DependencyIndexes: file_books_proto_depIdxs,
		MessageInfos:      file_books_proto_msgTypes,
	}.Build()
	File_books_proto = out.File
	file_books_proto_rawDesc = nil
	file_books_proto_goTypes = nil
	file_books_proto_depIdxs = nil
}
