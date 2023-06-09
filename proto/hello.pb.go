// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/hello.proto

package protoSamples

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

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sample) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sample) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type SampleResponseAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sample []*Sample `protobuf:"bytes,1,rep,name=sample,proto3" json:"sample,omitempty"`
}

func (x *SampleResponseAll) Reset() {
	*x = SampleResponseAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResponseAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponseAll) ProtoMessage() {}

func (x *SampleResponseAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponseAll.ProtoReflect.Descriptor instead.
func (*SampleResponseAll) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SampleResponseAll) GetSample() []*Sample {
	if x != nil {
		return x.Sample
	}
	return nil
}

type SampleResponseByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sample *Sample `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
}

func (x *SampleResponseByID) Reset() {
	*x = SampleResponseByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResponseByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponseByID) ProtoMessage() {}

func (x *SampleResponseByID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponseByID.ProtoReflect.Descriptor instead.
func (*SampleResponseByID) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *SampleResponseByID) GetSample() *Sample {
	if x != nil {
		return x.Sample
	}
	return nil
}

type SampleRequestAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SampleRequestAll) Reset() {
	*x = SampleRequestAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequestAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequestAll) ProtoMessage() {}

func (x *SampleRequestAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequestAll.ProtoReflect.Descriptor instead.
func (*SampleRequestAll) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{3}
}

type SampleRequestByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SampleRequestByID) Reset() {
	*x = SampleRequestByID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequestByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequestByID) ProtoMessage() {}

func (x *SampleRequestByID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequestByID.ProtoReflect.Descriptor instead.
func (*SampleRequestByID) Descriptor() ([]byte, []int) {
	return file_proto_hello_proto_rawDescGZIP(), []int{4}
}

func (x *SampleRequestByID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_hello_proto protoreflect.FileDescriptor

var file_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x22, 0x44, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x41, 0x0a, 0x11, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb7, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x49, 0x44, 0x22,
	0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x65, 0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_proto_rawDescOnce sync.Once
	file_proto_hello_proto_rawDescData = file_proto_hello_proto_rawDesc
)

func file_proto_hello_proto_rawDescGZIP() []byte {
	file_proto_hello_proto_rawDescOnce.Do(func() {
		file_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_proto_rawDescData)
	})
	return file_proto_hello_proto_rawDescData
}

var file_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_hello_proto_goTypes = []interface{}{
	(*Sample)(nil),             // 0: protoSamples.Sample
	(*SampleResponseAll)(nil),  // 1: protoSamples.SampleResponseAll
	(*SampleResponseByID)(nil), // 2: protoSamples.SampleResponseByID
	(*SampleRequestAll)(nil),   // 3: protoSamples.SampleRequestAll
	(*SampleRequestByID)(nil),  // 4: protoSamples.SampleRequestByID
}
var file_proto_hello_proto_depIdxs = []int32{
	0, // 0: protoSamples.SampleResponseAll.sample:type_name -> protoSamples.Sample
	0, // 1: protoSamples.SampleResponseByID.sample:type_name -> protoSamples.Sample
	3, // 2: protoSamples.SampleService.GetSampleID:input_type -> protoSamples.SampleRequestAll
	4, // 3: protoSamples.SampleService.GetSampleByID:input_type -> protoSamples.SampleRequestByID
	1, // 4: protoSamples.SampleService.GetSampleID:output_type -> protoSamples.SampleResponseAll
	2, // 5: protoSamples.SampleService.GetSampleByID:output_type -> protoSamples.SampleResponseByID
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_hello_proto_init() }
func file_proto_hello_proto_init() {
	if File_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResponseAll); i {
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
		file_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResponseByID); i {
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
		file_proto_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequestAll); i {
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
		file_proto_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequestByID); i {
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
			RawDescriptor: file_proto_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_proto_goTypes,
		DependencyIndexes: file_proto_hello_proto_depIdxs,
		MessageInfos:      file_proto_hello_proto_msgTypes,
	}.Build()
	File_proto_hello_proto = out.File
	file_proto_hello_proto_rawDesc = nil
	file_proto_hello_proto_goTypes = nil
	file_proto_hello_proto_depIdxs = nil
}
