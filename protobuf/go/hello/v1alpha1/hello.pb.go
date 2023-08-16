// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hello/v1alpha1/hello.proto

package hellov1alpha1

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

// Request message for the SayHello method.
type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the person to greet.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1alpha1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1alpha1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_v1alpha1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response message for the SayHello method.
type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The greeting.
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1alpha1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1alpha1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_hello_v1alpha1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

// Request message for the SayGoodbye method.
type SayGoodbyeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the person to greet.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayGoodbyeRequest) Reset() {
	*x = SayGoodbyeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1alpha1_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodbyeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodbyeRequest) ProtoMessage() {}

func (x *SayGoodbyeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1alpha1_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodbyeRequest.ProtoReflect.Descriptor instead.
func (*SayGoodbyeRequest) Descriptor() ([]byte, []int) {
	return file_hello_v1alpha1_hello_proto_rawDescGZIP(), []int{2}
}

func (x *SayGoodbyeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response message for the SayGoodbye method.
type SayGoodbyeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The goodbye.
	Goodbye string `protobuf:"bytes,1,opt,name=goodbye,proto3" json:"goodbye,omitempty"`
}

func (x *SayGoodbyeResponse) Reset() {
	*x = SayGoodbyeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1alpha1_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodbyeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodbyeResponse) ProtoMessage() {}

func (x *SayGoodbyeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1alpha1_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodbyeResponse.ProtoReflect.Descriptor instead.
func (*SayGoodbyeResponse) Descriptor() ([]byte, []int) {
	return file_hello_v1alpha1_hello_proto_rawDescGZIP(), []int{3}
}

func (x *SayGoodbyeResponse) GetGoodbye() string {
	if x != nil {
		return x.Goodbye
	}
	return ""
}

var File_hello_v1alpha1_hello_proto protoreflect.FileDescriptor

var file_hello_v1alpha1_hello_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x25, 0x0a, 0x0f,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12,
	0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x32, 0xb6, 0x01, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55,
	0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x12, 0x21, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc6, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x65, 0x62, 0x72, 0x6f, 0x77, 0x6e, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x0e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x0e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x1a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_v1alpha1_hello_proto_rawDescOnce sync.Once
	file_hello_v1alpha1_hello_proto_rawDescData = file_hello_v1alpha1_hello_proto_rawDesc
)

func file_hello_v1alpha1_hello_proto_rawDescGZIP() []byte {
	file_hello_v1alpha1_hello_proto_rawDescOnce.Do(func() {
		file_hello_v1alpha1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_v1alpha1_hello_proto_rawDescData)
	})
	return file_hello_v1alpha1_hello_proto_rawDescData
}

var file_hello_v1alpha1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hello_v1alpha1_hello_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),    // 0: hello.v1alpha1.SayHelloRequest
	(*SayHelloResponse)(nil),   // 1: hello.v1alpha1.SayHelloResponse
	(*SayGoodbyeRequest)(nil),  // 2: hello.v1alpha1.SayGoodbyeRequest
	(*SayGoodbyeResponse)(nil), // 3: hello.v1alpha1.SayGoodbyeResponse
}
var file_hello_v1alpha1_hello_proto_depIdxs = []int32{
	0, // 0: hello.v1alpha1.HelloService.SayHello:input_type -> hello.v1alpha1.SayHelloRequest
	2, // 1: hello.v1alpha1.HelloService.SayGoodbye:input_type -> hello.v1alpha1.SayGoodbyeRequest
	1, // 2: hello.v1alpha1.HelloService.SayHello:output_type -> hello.v1alpha1.SayHelloResponse
	3, // 3: hello.v1alpha1.HelloService.SayGoodbye:output_type -> hello.v1alpha1.SayGoodbyeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_v1alpha1_hello_proto_init() }
func file_hello_v1alpha1_hello_proto_init() {
	if File_hello_v1alpha1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_v1alpha1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_hello_v1alpha1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_hello_v1alpha1_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodbyeRequest); i {
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
		file_hello_v1alpha1_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodbyeResponse); i {
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
			RawDescriptor: file_hello_v1alpha1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_v1alpha1_hello_proto_goTypes,
		DependencyIndexes: file_hello_v1alpha1_hello_proto_depIdxs,
		MessageInfos:      file_hello_v1alpha1_hello_proto_msgTypes,
	}.Build()
	File_hello_v1alpha1_hello_proto = out.File
	file_hello_v1alpha1_hello_proto_rawDesc = nil
	file_hello_v1alpha1_hello_proto_goTypes = nil
	file_hello_v1alpha1_hello_proto_depIdxs = nil
}
