// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gproto

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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*Response, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Response, error)
	Notifications(ctx context.Context, in *RegisterNotifications, opts ...grpc.CallOption) (Game_NotificationsClient, error)
	Movement(ctx context.Context, in *MovementRequest, opts ...grpc.CallOption) (*Response, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/gproto.Game/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/gproto.Game/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Notifications(ctx context.Context, in *RegisterNotifications, opts ...grpc.CallOption) (Game_NotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[0], "/gproto.Game/Notifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Game_NotificationsClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type gameNotificationsClient struct {
	grpc.ClientStream
}

func (x *gameNotificationsClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameClient) Movement(ctx context.Context, in *MovementRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/gproto.Game/Movement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	NewGame(context.Context, *NewGameRequest) (*Response, error)
	Join(context.Context, *JoinRequest) (*Response, error)
	Notifications(*RegisterNotifications, Game_NotificationsServer) error
	Movement(context.Context, *MovementRequest) (*Response, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) NewGame(context.Context, *NewGameRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (UnimplementedGameServer) Join(context.Context, *JoinRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGameServer) Notifications(*RegisterNotifications, Game_NotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method Notifications not implemented")
}
func (UnimplementedGameServer) Movement(context.Context, *MovementRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Movement not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.Game/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.Game/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Notifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterNotifications)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServer).Notifications(m, &gameNotificationsServer{stream})
}

type Game_NotificationsServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type gameNotificationsServer struct {
	grpc.ServerStream
}

func (x *gameNotificationsServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Game_Movement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Movement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gproto.Game/Movement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Movement(ctx, req.(*MovementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gproto.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _Game_NewGame_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Game_Join_Handler,
		},
		{
			MethodName: "Movement",
			Handler:    _Game_Movement_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notifications",
			Handler:       _Game_Notifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server/game.proto",
}
