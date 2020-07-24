// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: currentLoop.proto

package currentLoop

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CurrentLoopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash      string  `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Timestamp string  `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     int32   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	Voltage   float32 `protobuf:"fixed32,4,opt,name=voltage,proto3" json:"voltage,omitempty"`
	Current   float32 `protobuf:"fixed32,5,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *CurrentLoopRequest) Reset() {
	*x = CurrentLoopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currentLoop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentLoopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentLoopRequest) ProtoMessage() {}

func (x *CurrentLoopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currentLoop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentLoopRequest.ProtoReflect.Descriptor instead.
func (*CurrentLoopRequest) Descriptor() ([]byte, []int) {
	return file_currentLoop_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentLoopRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *CurrentLoopRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CurrentLoopRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CurrentLoopRequest) GetVoltage() float32 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

func (x *CurrentLoopRequest) GetCurrent() float32 {
	if x != nil {
		return x.Current
	}
	return 0
}

type CurrentLoopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CurrentLoopResponse) Reset() {
	*x = CurrentLoopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currentLoop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentLoopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentLoopResponse) ProtoMessage() {}

func (x *CurrentLoopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currentLoop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentLoopResponse.ProtoReflect.Descriptor instead.
func (*CurrentLoopResponse) Descriptor() ([]byte, []int) {
	return file_currentLoop_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentLoopResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_currentLoop_proto protoreflect.FileDescriptor

var file_currentLoop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x32, 0x4f, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6f, 0x70,
	0x12, 0x40, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x12, 0x13, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x64,
	0x68, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x6f, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currentLoop_proto_rawDescOnce sync.Once
	file_currentLoop_proto_rawDescData = file_currentLoop_proto_rawDesc
)

func file_currentLoop_proto_rawDescGZIP() []byte {
	file_currentLoop_proto_rawDescOnce.Do(func() {
		file_currentLoop_proto_rawDescData = protoimpl.X.CompressGZIP(file_currentLoop_proto_rawDescData)
	})
	return file_currentLoop_proto_rawDescData
}

var file_currentLoop_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_currentLoop_proto_goTypes = []interface{}{
	(*CurrentLoopRequest)(nil),  // 0: CurrentLoopRequest
	(*CurrentLoopResponse)(nil), // 1: CurrentLoopResponse
}
var file_currentLoop_proto_depIdxs = []int32{
	0, // 0: CurrentLoop.SendTelemetry:input_type -> CurrentLoopRequest
	1, // 1: CurrentLoop.SendTelemetry:output_type -> CurrentLoopResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_currentLoop_proto_init() }
func file_currentLoop_proto_init() {
	if File_currentLoop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currentLoop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentLoopRequest); i {
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
		file_currentLoop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentLoopResponse); i {
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
			RawDescriptor: file_currentLoop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currentLoop_proto_goTypes,
		DependencyIndexes: file_currentLoop_proto_depIdxs,
		MessageInfos:      file_currentLoop_proto_msgTypes,
	}.Build()
	File_currentLoop_proto = out.File
	file_currentLoop_proto_rawDesc = nil
	file_currentLoop_proto_goTypes = nil
	file_currentLoop_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrentLoopClient is the client API for CurrentLoop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrentLoopClient interface {
	SendTelemetry(ctx context.Context, opts ...grpc.CallOption) (CurrentLoop_SendTelemetryClient, error)
}

type currentLoopClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrentLoopClient(cc grpc.ClientConnInterface) CurrentLoopClient {
	return &currentLoopClient{cc}
}

func (c *currentLoopClient) SendTelemetry(ctx context.Context, opts ...grpc.CallOption) (CurrentLoop_SendTelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CurrentLoop_serviceDesc.Streams[0], "/CurrentLoop/SendTelemetry", opts...)
	if err != nil {
		return nil, err
	}
	x := &currentLoopSendTelemetryClient{stream}
	return x, nil
}

type CurrentLoop_SendTelemetryClient interface {
	Send(*CurrentLoopRequest) error
	Recv() (*CurrentLoopResponse, error)
	grpc.ClientStream
}

type currentLoopSendTelemetryClient struct {
	grpc.ClientStream
}

func (x *currentLoopSendTelemetryClient) Send(m *CurrentLoopRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *currentLoopSendTelemetryClient) Recv() (*CurrentLoopResponse, error) {
	m := new(CurrentLoopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrentLoopServer is the server API for CurrentLoop service.
type CurrentLoopServer interface {
	SendTelemetry(CurrentLoop_SendTelemetryServer) error
}

// UnimplementedCurrentLoopServer can be embedded to have forward compatible implementations.
type UnimplementedCurrentLoopServer struct {
}

func (*UnimplementedCurrentLoopServer) SendTelemetry(CurrentLoop_SendTelemetryServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTelemetry not implemented")
}

func RegisterCurrentLoopServer(s *grpc.Server, srv CurrentLoopServer) {
	s.RegisterService(&_CurrentLoop_serviceDesc, srv)
}

func _CurrentLoop_SendTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CurrentLoopServer).SendTelemetry(&currentLoopSendTelemetryServer{stream})
}

type CurrentLoop_SendTelemetryServer interface {
	Send(*CurrentLoopResponse) error
	Recv() (*CurrentLoopRequest, error)
	grpc.ServerStream
}

type currentLoopSendTelemetryServer struct {
	grpc.ServerStream
}

func (x *currentLoopSendTelemetryServer) Send(m *CurrentLoopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *currentLoopSendTelemetryServer) Recv() (*CurrentLoopRequest, error) {
	m := new(CurrentLoopRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CurrentLoop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CurrentLoop",
	HandlerType: (*CurrentLoopServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendTelemetry",
			Handler:       _CurrentLoop_SendTelemetry_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "currentLoop.proto",
}
