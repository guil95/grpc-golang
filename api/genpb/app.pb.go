// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: app.proto

package rickmorty

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CharactersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Characters []*Character `protobuf:"bytes,1,rep,name=characters,proto3" json:"characters,omitempty"`
}

func (x *CharactersResponse) Reset() {
	*x = CharactersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharactersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharactersResponse) ProtoMessage() {}

func (x *CharactersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharactersResponse.ProtoReflect.Descriptor instead.
func (*CharactersResponse) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *CharactersResponse) GetCharacters() []*Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species string `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{1}
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74,
	0x79, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x65, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x10, 0x52, 0x69, 0x63, 0x6b, 0x4d, 0x6f, 0x72, 0x74, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d,
	0x6f, 0x72, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42,
	0x0f, 0x5a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x69, 0x63, 0x6b, 0x6d, 0x6f, 0x72, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_proto_goTypes = []interface{}{
	(*CharactersResponse)(nil), // 0: api.rickmorty.CharactersResponse
	(*Character)(nil),          // 1: api.rickmorty.Character
	(*emptypb.Empty)(nil),      // 2: google.protobuf.Empty
}
var file_app_proto_depIdxs = []int32{
	1, // 0: api.rickmorty.CharactersResponse.characters:type_name -> api.rickmorty.Character
	2, // 1: api.rickmorty.RickMortyService.ListAllCharacters:input_type -> google.protobuf.Empty
	2, // 2: api.rickmorty.RickMortyService.ListAllCharactersStream:input_type -> google.protobuf.Empty
	0, // 3: api.rickmorty.RickMortyService.ListAllCharacters:output_type -> api.rickmorty.CharactersResponse
	0, // 4: api.rickmorty.RickMortyService.ListAllCharactersStream:output_type -> api.rickmorty.CharactersResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharactersResponse); i {
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
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RickMortyServiceClient is the client API for RickMortyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RickMortyServiceClient interface {
	ListAllCharacters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CharactersResponse, error)
	ListAllCharactersStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RickMortyService_ListAllCharactersStreamClient, error)
}

type rickMortyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRickMortyServiceClient(cc grpc.ClientConnInterface) RickMortyServiceClient {
	return &rickMortyServiceClient{cc}
}

func (c *rickMortyServiceClient) ListAllCharacters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CharactersResponse, error) {
	out := new(CharactersResponse)
	err := c.cc.Invoke(ctx, "/api.rickmorty.RickMortyService/ListAllCharacters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rickMortyServiceClient) ListAllCharactersStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RickMortyService_ListAllCharactersStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RickMortyService_serviceDesc.Streams[0], "/api.rickmorty.RickMortyService/ListAllCharactersStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rickMortyServiceListAllCharactersStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RickMortyService_ListAllCharactersStreamClient interface {
	Recv() (*CharactersResponse, error)
	grpc.ClientStream
}

type rickMortyServiceListAllCharactersStreamClient struct {
	grpc.ClientStream
}

func (x *rickMortyServiceListAllCharactersStreamClient) Recv() (*CharactersResponse, error) {
	m := new(CharactersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RickMortyServiceServer is the server API for RickMortyService service.
type RickMortyServiceServer interface {
	ListAllCharacters(context.Context, *emptypb.Empty) (*CharactersResponse, error)
	ListAllCharactersStream(*emptypb.Empty, RickMortyService_ListAllCharactersStreamServer) error
}

// UnimplementedRickMortyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRickMortyServiceServer struct {
}

func (*UnimplementedRickMortyServiceServer) ListAllCharacters(context.Context, *emptypb.Empty) (*CharactersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllCharacters not implemented")
}
func (*UnimplementedRickMortyServiceServer) ListAllCharactersStream(*emptypb.Empty, RickMortyService_ListAllCharactersStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllCharactersStream not implemented")
}

func RegisterRickMortyServiceServer(s *grpc.Server, srv RickMortyServiceServer) {
	s.RegisterService(&_RickMortyService_serviceDesc, srv)
}

func _RickMortyService_ListAllCharacters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RickMortyServiceServer).ListAllCharacters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.rickmorty.RickMortyService/ListAllCharacters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RickMortyServiceServer).ListAllCharacters(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RickMortyService_ListAllCharactersStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RickMortyServiceServer).ListAllCharactersStream(m, &rickMortyServiceListAllCharactersStreamServer{stream})
}

type RickMortyService_ListAllCharactersStreamServer interface {
	Send(*CharactersResponse) error
	grpc.ServerStream
}

type rickMortyServiceListAllCharactersStreamServer struct {
	grpc.ServerStream
}

func (x *rickMortyServiceListAllCharactersStreamServer) Send(m *CharactersResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RickMortyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.rickmorty.RickMortyService",
	HandlerType: (*RickMortyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllCharacters",
			Handler:    _RickMortyService_ListAllCharacters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllCharactersStream",
			Handler:       _RickMortyService_ListAllCharactersStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "app.proto",
}