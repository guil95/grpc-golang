// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/app.proto

package rickmorty

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RickMortyServiceClient is the client API for RickMortyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RickMortyServiceClient interface {
	ListAllCharacters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CharactersResponse, error)
	ListAllCharactersServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RickMortyService_ListAllCharactersServerStreamClient, error)
	ListByPageClientStream(ctx context.Context, opts ...grpc.CallOption) (RickMortyService_ListByPageClientStreamClient, error)
}

type rickMortyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRickMortyServiceClient(cc grpc.ClientConnInterface) RickMortyServiceClient {
	return &rickMortyServiceClient{cc}
}

func (c *rickMortyServiceClient) ListAllCharacters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CharactersResponse, error) {
	out := new(CharactersResponse)
	err := c.cc.Invoke(ctx, "/rickmorty.RickMortyService/ListAllCharacters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rickMortyServiceClient) ListAllCharactersServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RickMortyService_ListAllCharactersServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RickMortyService_ServiceDesc.Streams[0], "/rickmorty.RickMortyService/ListAllCharactersServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rickMortyServiceListAllCharactersServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RickMortyService_ListAllCharactersServerStreamClient interface {
	Recv() (*CharactersResponse, error)
	grpc.ClientStream
}

type rickMortyServiceListAllCharactersServerStreamClient struct {
	grpc.ClientStream
}

func (x *rickMortyServiceListAllCharactersServerStreamClient) Recv() (*CharactersResponse, error) {
	m := new(CharactersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rickMortyServiceClient) ListByPageClientStream(ctx context.Context, opts ...grpc.CallOption) (RickMortyService_ListByPageClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RickMortyService_ServiceDesc.Streams[1], "/rickmorty.RickMortyService/ListByPageClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rickMortyServiceListByPageClientStreamClient{stream}
	return x, nil
}

type RickMortyService_ListByPageClientStreamClient interface {
	Send(*PageRequest) error
	CloseAndRecv() (*CharactersResponse, error)
	grpc.ClientStream
}

type rickMortyServiceListByPageClientStreamClient struct {
	grpc.ClientStream
}

func (x *rickMortyServiceListByPageClientStreamClient) Send(m *PageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rickMortyServiceListByPageClientStreamClient) CloseAndRecv() (*CharactersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CharactersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RickMortyServiceServer is the server API for RickMortyService service.
// All implementations should embed UnimplementedRickMortyServiceServer
// for forward compatibility
type RickMortyServiceServer interface {
	ListAllCharacters(context.Context, *emptypb.Empty) (*CharactersResponse, error)
	ListAllCharactersServerStream(*emptypb.Empty, RickMortyService_ListAllCharactersServerStreamServer) error
	ListByPageClientStream(RickMortyService_ListByPageClientStreamServer) error
}

// UnimplementedRickMortyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRickMortyServiceServer struct {
}

func (UnimplementedRickMortyServiceServer) ListAllCharacters(context.Context, *emptypb.Empty) (*CharactersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllCharacters not implemented")
}
func (UnimplementedRickMortyServiceServer) ListAllCharactersServerStream(*emptypb.Empty, RickMortyService_ListAllCharactersServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllCharactersServerStream not implemented")
}
func (UnimplementedRickMortyServiceServer) ListByPageClientStream(RickMortyService_ListByPageClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListByPageClientStream not implemented")
}

// UnsafeRickMortyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RickMortyServiceServer will
// result in compilation errors.
type UnsafeRickMortyServiceServer interface {
	mustEmbedUnimplementedRickMortyServiceServer()
}

func RegisterRickMortyServiceServer(s grpc.ServiceRegistrar, srv RickMortyServiceServer) {
	s.RegisterService(&RickMortyService_ServiceDesc, srv)
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
		FullMethod: "/rickmorty.RickMortyService/ListAllCharacters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RickMortyServiceServer).ListAllCharacters(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RickMortyService_ListAllCharactersServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RickMortyServiceServer).ListAllCharactersServerStream(m, &rickMortyServiceListAllCharactersServerStreamServer{stream})
}

type RickMortyService_ListAllCharactersServerStreamServer interface {
	Send(*CharactersResponse) error
	grpc.ServerStream
}

type rickMortyServiceListAllCharactersServerStreamServer struct {
	grpc.ServerStream
}

func (x *rickMortyServiceListAllCharactersServerStreamServer) Send(m *CharactersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RickMortyService_ListByPageClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RickMortyServiceServer).ListByPageClientStream(&rickMortyServiceListByPageClientStreamServer{stream})
}

type RickMortyService_ListByPageClientStreamServer interface {
	SendAndClose(*CharactersResponse) error
	Recv() (*PageRequest, error)
	grpc.ServerStream
}

type rickMortyServiceListByPageClientStreamServer struct {
	grpc.ServerStream
}

func (x *rickMortyServiceListByPageClientStreamServer) SendAndClose(m *CharactersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rickMortyServiceListByPageClientStreamServer) Recv() (*PageRequest, error) {
	m := new(PageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RickMortyService_ServiceDesc is the grpc.ServiceDesc for RickMortyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RickMortyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rickmorty.RickMortyService",
	HandlerType: (*RickMortyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllCharacters",
			Handler:    _RickMortyService_ListAllCharacters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllCharactersServerStream",
			Handler:       _RickMortyService_ListAllCharactersServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListByPageClientStream",
			Handler:       _RickMortyService_ListByPageClientStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/app.proto",
}
