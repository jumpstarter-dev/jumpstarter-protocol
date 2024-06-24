// Copyright 2024 The Jumpstarter Authors

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: jumpstarter/v1/rendezvous.proto

package jumpstarterv1

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

type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{0}
}

type ListenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListenResponse) Reset() {
	*x = ListenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenResponse) ProtoMessage() {}

func (x *ListenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenResponse.ProtoReflect.Descriptor instead.
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{1}
}

func (x *ListenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *DialRequest) Reset() {
	*x = DialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialRequest) ProtoMessage() {}

func (x *DialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialRequest.ProtoReflect.Descriptor instead.
func (*DialRequest) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{2}
}

func (x *DialRequest) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

type DialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DialResponse) Reset() {
	*x = DialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialResponse) ProtoMessage() {}

func (x *DialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialResponse.ProtoReflect.Descriptor instead.
func (*DialResponse) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{3}
}

func (x *DialResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{4}
}

func (x *StreamRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jumpstarter_v1_rendezvous_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_jumpstarter_v1_rendezvous_proto_rawDescGZIP(), []int{5}
}

func (x *StreamResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_jumpstarter_v1_rendezvous_proto protoreflect.FileDescriptor

var file_jumpstarter_v1_rendezvous_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x26, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x0b, 0x44, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0x24, 0x0a, 0x0c, 0x44,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2a, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xa1, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6e,
	0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x04, 0x44, 0x69, 0x61,
	0x6c, 0x12, 0x1b, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5c, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0xcc, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x75, 0x6d, 0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x4a, 0x75, 0x6d, 0x70, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x4a, 0x75, 0x6d, 0x70,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x4a, 0x75, 0x6d,
	0x70, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4a, 0x75, 0x6d, 0x70, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_jumpstarter_v1_rendezvous_proto_rawDescOnce sync.Once
	file_jumpstarter_v1_rendezvous_proto_rawDescData = file_jumpstarter_v1_rendezvous_proto_rawDesc
)

func file_jumpstarter_v1_rendezvous_proto_rawDescGZIP() []byte {
	file_jumpstarter_v1_rendezvous_proto_rawDescOnce.Do(func() {
		file_jumpstarter_v1_rendezvous_proto_rawDescData = protoimpl.X.CompressGZIP(file_jumpstarter_v1_rendezvous_proto_rawDescData)
	})
	return file_jumpstarter_v1_rendezvous_proto_rawDescData
}

var file_jumpstarter_v1_rendezvous_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_jumpstarter_v1_rendezvous_proto_goTypes = []any{
	(*ListenRequest)(nil),  // 0: jumpstarter.v1.ListenRequest
	(*ListenResponse)(nil), // 1: jumpstarter.v1.ListenResponse
	(*DialRequest)(nil),    // 2: jumpstarter.v1.DialRequest
	(*DialResponse)(nil),   // 3: jumpstarter.v1.DialResponse
	(*StreamRequest)(nil),  // 4: jumpstarter.v1.StreamRequest
	(*StreamResponse)(nil), // 5: jumpstarter.v1.StreamResponse
}
var file_jumpstarter_v1_rendezvous_proto_depIdxs = []int32{
	0, // 0: jumpstarter.v1.RendezvousService.Listen:input_type -> jumpstarter.v1.ListenRequest
	2, // 1: jumpstarter.v1.RendezvousService.Dial:input_type -> jumpstarter.v1.DialRequest
	4, // 2: jumpstarter.v1.StreamService.Stream:input_type -> jumpstarter.v1.StreamRequest
	1, // 3: jumpstarter.v1.RendezvousService.Listen:output_type -> jumpstarter.v1.ListenResponse
	3, // 4: jumpstarter.v1.RendezvousService.Dial:output_type -> jumpstarter.v1.DialResponse
	5, // 5: jumpstarter.v1.StreamService.Stream:output_type -> jumpstarter.v1.StreamResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jumpstarter_v1_rendezvous_proto_init() }
func file_jumpstarter_v1_rendezvous_proto_init() {
	if File_jumpstarter_v1_rendezvous_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jumpstarter_v1_rendezvous_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListenRequest); i {
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
		file_jumpstarter_v1_rendezvous_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListenResponse); i {
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
		file_jumpstarter_v1_rendezvous_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DialRequest); i {
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
		file_jumpstarter_v1_rendezvous_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DialResponse); i {
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
		file_jumpstarter_v1_rendezvous_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StreamRequest); i {
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
		file_jumpstarter_v1_rendezvous_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*StreamResponse); i {
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
			RawDescriptor: file_jumpstarter_v1_rendezvous_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_jumpstarter_v1_rendezvous_proto_goTypes,
		DependencyIndexes: file_jumpstarter_v1_rendezvous_proto_depIdxs,
		MessageInfos:      file_jumpstarter_v1_rendezvous_proto_msgTypes,
	}.Build()
	File_jumpstarter_v1_rendezvous_proto = out.File
	file_jumpstarter_v1_rendezvous_proto_rawDesc = nil
	file_jumpstarter_v1_rendezvous_proto_goTypes = nil
	file_jumpstarter_v1_rendezvous_proto_depIdxs = nil
}