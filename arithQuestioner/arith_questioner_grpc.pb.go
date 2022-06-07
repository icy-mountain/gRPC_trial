// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: arithQuestioner/arith_questioner.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArithQuestionerClient is the client API for ArithQuestioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithQuestionerClient interface {
	QuestionChat(ctx context.Context, opts ...grpc.CallOption) (ArithQuestioner_QuestionChatClient, error)
}

type arithQuestionerClient struct {
	cc grpc.ClientConnInterface
}

func NewArithQuestionerClient(cc grpc.ClientConnInterface) ArithQuestionerClient {
	return &arithQuestionerClient{cc}
}

func (c *arithQuestionerClient) QuestionChat(ctx context.Context, opts ...grpc.CallOption) (ArithQuestioner_QuestionChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArithQuestioner_ServiceDesc.Streams[0], "/arithquestioner.ArithQuestioner/QuestionChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &arithQuestionerQuestionChatClient{stream}
	return x, nil
}

type ArithQuestioner_QuestionChatClient interface {
	Send(*QuestionMessage) error
	Recv() (*QuestionMessage, error)
	grpc.ClientStream
}

type arithQuestionerQuestionChatClient struct {
	grpc.ClientStream
}

func (x *arithQuestionerQuestionChatClient) Send(m *QuestionMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *arithQuestionerQuestionChatClient) Recv() (*QuestionMessage, error) {
	m := new(QuestionMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArithQuestionerServer is the server API for ArithQuestioner service.
// All implementations must embed UnimplementedArithQuestionerServer
// for forward compatibility
type ArithQuestionerServer interface {
	QuestionChat(ArithQuestioner_QuestionChatServer) error
	mustEmbedUnimplementedArithQuestionerServer()
}

// UnimplementedArithQuestionerServer must be embedded to have forward compatible implementations.
type UnimplementedArithQuestionerServer struct {
}

func (UnimplementedArithQuestionerServer) QuestionChat(ArithQuestioner_QuestionChatServer) error {
	return status.Errorf(codes.Unimplemented, "method QuestionChat not implemented")
}
func (UnimplementedArithQuestionerServer) mustEmbedUnimplementedArithQuestionerServer() {}

// UnsafeArithQuestionerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithQuestionerServer will
// result in compilation errors.
type UnsafeArithQuestionerServer interface {
	mustEmbedUnimplementedArithQuestionerServer()
}

func RegisterArithQuestionerServer(s grpc.ServiceRegistrar, srv ArithQuestionerServer) {
	s.RegisterService(&ArithQuestioner_ServiceDesc, srv)
}

func _ArithQuestioner_QuestionChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ArithQuestionerServer).QuestionChat(&arithQuestionerQuestionChatServer{stream})
}

type ArithQuestioner_QuestionChatServer interface {
	Send(*QuestionMessage) error
	Recv() (*QuestionMessage, error)
	grpc.ServerStream
}

type arithQuestionerQuestionChatServer struct {
	grpc.ServerStream
}

func (x *arithQuestionerQuestionChatServer) Send(m *QuestionMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *arithQuestionerQuestionChatServer) Recv() (*QuestionMessage, error) {
	m := new(QuestionMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArithQuestioner_ServiceDesc is the grpc.ServiceDesc for ArithQuestioner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithQuestioner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arithquestioner.ArithQuestioner",
	HandlerType: (*ArithQuestionerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QuestionChat",
			Handler:       _ArithQuestioner_QuestionChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "arithQuestioner/arith_questioner.proto",
}