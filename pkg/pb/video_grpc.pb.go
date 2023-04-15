// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: video.proto

package pb

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

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoReply, error)
	GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosReply, error)
	LikeAction(ctx context.Context, in *LikeActionRequest, opts ...grpc.CallOption) (*LikeActionReply, error)
	GetLikedVideos(ctx context.Context, in *GetLikedVideosRequest, opts ...grpc.CallOption) (*GetLikedVideosReply, error)
	GetUserVideos(ctx context.Context, in *GetUserVideosRequest, opts ...grpc.CallOption) (*GetUserVideosReply, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoReply, error) {
	out := new(GetVideoReply)
	err := c.cc.Invoke(ctx, "/sixwaaaay.video.VideoService/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosReply, error) {
	out := new(GetVideosReply)
	err := c.cc.Invoke(ctx, "/sixwaaaay.video.VideoService/GetVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) LikeAction(ctx context.Context, in *LikeActionRequest, opts ...grpc.CallOption) (*LikeActionReply, error) {
	out := new(LikeActionReply)
	err := c.cc.Invoke(ctx, "/sixwaaaay.video.VideoService/LikeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetLikedVideos(ctx context.Context, in *GetLikedVideosRequest, opts ...grpc.CallOption) (*GetLikedVideosReply, error) {
	out := new(GetLikedVideosReply)
	err := c.cc.Invoke(ctx, "/sixwaaaay.video.VideoService/GetLikedVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) GetUserVideos(ctx context.Context, in *GetUserVideosRequest, opts ...grpc.CallOption) (*GetUserVideosReply, error) {
	out := new(GetUserVideosReply)
	err := c.cc.Invoke(ctx, "/sixwaaaay.video.VideoService/GetUserVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	GetVideo(context.Context, *GetVideoRequest) (*GetVideoReply, error)
	GetVideos(context.Context, *GetVideosRequest) (*GetVideosReply, error)
	LikeAction(context.Context, *LikeActionRequest) (*LikeActionReply, error)
	GetLikedVideos(context.Context, *GetLikedVideosRequest) (*GetLikedVideosReply, error)
	GetUserVideos(context.Context, *GetUserVideosRequest) (*GetUserVideosReply, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) GetVideo(context.Context, *GetVideoRequest) (*GetVideoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedVideoServiceServer) GetVideos(context.Context, *GetVideosRequest) (*GetVideosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideos not implemented")
}
func (UnimplementedVideoServiceServer) LikeAction(context.Context, *LikeActionRequest) (*LikeActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeAction not implemented")
}
func (UnimplementedVideoServiceServer) GetLikedVideos(context.Context, *GetLikedVideosRequest) (*GetLikedVideosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikedVideos not implemented")
}
func (UnimplementedVideoServiceServer) GetUserVideos(context.Context, *GetUserVideosRequest) (*GetUserVideosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserVideos not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixwaaaay.video.VideoService/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixwaaaay.video.VideoService/GetVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideos(ctx, req.(*GetVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_LikeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).LikeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixwaaaay.video.VideoService/LikeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).LikeAction(ctx, req.(*LikeActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetLikedVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikedVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetLikedVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixwaaaay.video.VideoService/GetLikedVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetLikedVideos(ctx, req.(*GetLikedVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_GetUserVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetUserVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sixwaaaay.video.VideoService/GetUserVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetUserVideos(ctx, req.(*GetUserVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sixwaaaay.video.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideo",
			Handler:    _VideoService_GetVideo_Handler,
		},
		{
			MethodName: "GetVideos",
			Handler:    _VideoService_GetVideos_Handler,
		},
		{
			MethodName: "LikeAction",
			Handler:    _VideoService_LikeAction_Handler,
		},
		{
			MethodName: "GetLikedVideos",
			Handler:    _VideoService_GetLikedVideos_Handler,
		},
		{
			MethodName: "GetUserVideos",
			Handler:    _VideoService_GetUserVideos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
