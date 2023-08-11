// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: feed_server.proto

package pbfs

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

const (
	Upstream_Handshake_FullMethodName       = "/pbfs.Upstream/Handshake"
	Upstream_Watch_FullMethodName           = "/pbfs.Upstream/Watch"
	Upstream_Messaging_FullMethodName       = "/pbfs.Upstream/Messaging"
	Upstream_PullAppFileMeta_FullMethodName = "/pbfs.Upstream/PullAppFileMeta"
	Upstream_GetDownloadURL_FullMethodName  = "/pbfs.Upstream/GetDownloadURL"
)

// UpstreamClient is the client API for Upstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpstreamClient interface {
	// APIs for sidecar.
	Handshake(ctx context.Context, in *HandshakeMessage, opts ...grpc.CallOption) (*HandshakeResp, error)
	Watch(ctx context.Context, in *SideWatchMeta, opts ...grpc.CallOption) (Upstream_WatchClient, error)
	Messaging(ctx context.Context, in *MessagingMeta, opts ...grpc.CallOption) (*MessagingResp, error)
	PullAppFileMeta(ctx context.Context, in *PullAppFileMetaReq, opts ...grpc.CallOption) (*PullAppFileMetaResp, error)
	GetDownloadURL(ctx context.Context, in *GetDownloadURLReq, opts ...grpc.CallOption) (*GetDownloadURLResp, error)
}

type upstreamClient struct {
	cc grpc.ClientConnInterface
}

func NewUpstreamClient(cc grpc.ClientConnInterface) UpstreamClient {
	return &upstreamClient{cc}
}

func (c *upstreamClient) Handshake(ctx context.Context, in *HandshakeMessage, opts ...grpc.CallOption) (*HandshakeResp, error) {
	out := new(HandshakeResp)
	err := c.cc.Invoke(ctx, Upstream_Handshake_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) Watch(ctx context.Context, in *SideWatchMeta, opts ...grpc.CallOption) (Upstream_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Upstream_ServiceDesc.Streams[0], Upstream_Watch_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &upstreamWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Upstream_WatchClient interface {
	Recv() (*FeedWatchMessage, error)
	grpc.ClientStream
}

type upstreamWatchClient struct {
	grpc.ClientStream
}

func (x *upstreamWatchClient) Recv() (*FeedWatchMessage, error) {
	m := new(FeedWatchMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *upstreamClient) Messaging(ctx context.Context, in *MessagingMeta, opts ...grpc.CallOption) (*MessagingResp, error) {
	out := new(MessagingResp)
	err := c.cc.Invoke(ctx, Upstream_Messaging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) PullAppFileMeta(ctx context.Context, in *PullAppFileMetaReq, opts ...grpc.CallOption) (*PullAppFileMetaResp, error) {
	out := new(PullAppFileMetaResp)
	err := c.cc.Invoke(ctx, Upstream_PullAppFileMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamClient) GetDownloadURL(ctx context.Context, in *GetDownloadURLReq, opts ...grpc.CallOption) (*GetDownloadURLResp, error) {
	out := new(GetDownloadURLResp)
	err := c.cc.Invoke(ctx, Upstream_GetDownloadURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpstreamServer is the server API for Upstream service.
// All implementations should embed UnimplementedUpstreamServer
// for forward compatibility
type UpstreamServer interface {
	// APIs for sidecar.
	Handshake(context.Context, *HandshakeMessage) (*HandshakeResp, error)
	Watch(*SideWatchMeta, Upstream_WatchServer) error
	Messaging(context.Context, *MessagingMeta) (*MessagingResp, error)
	PullAppFileMeta(context.Context, *PullAppFileMetaReq) (*PullAppFileMetaResp, error)
	GetDownloadURL(context.Context, *GetDownloadURLReq) (*GetDownloadURLResp, error)
}

// UnimplementedUpstreamServer should be embedded to have forward compatible implementations.
type UnimplementedUpstreamServer struct {
}

func (UnimplementedUpstreamServer) Handshake(context.Context, *HandshakeMessage) (*HandshakeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedUpstreamServer) Watch(*SideWatchMeta, Upstream_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedUpstreamServer) Messaging(context.Context, *MessagingMeta) (*MessagingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Messaging not implemented")
}
func (UnimplementedUpstreamServer) PullAppFileMeta(context.Context, *PullAppFileMetaReq) (*PullAppFileMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullAppFileMeta not implemented")
}
func (UnimplementedUpstreamServer) GetDownloadURL(context.Context, *GetDownloadURLReq) (*GetDownloadURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadURL not implemented")
}

// UnsafeUpstreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpstreamServer will
// result in compilation errors.
type UnsafeUpstreamServer interface {
	mustEmbedUnimplementedUpstreamServer()
}

func RegisterUpstreamServer(s grpc.ServiceRegistrar, srv UpstreamServer) {
	s.RegisterService(&Upstream_ServiceDesc, srv)
}

func _Upstream_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_Handshake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).Handshake(ctx, req.(*HandshakeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SideWatchMeta)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpstreamServer).Watch(m, &upstreamWatchServer{stream})
}

type Upstream_WatchServer interface {
	Send(*FeedWatchMessage) error
	grpc.ServerStream
}

type upstreamWatchServer struct {
	grpc.ServerStream
}

func (x *upstreamWatchServer) Send(m *FeedWatchMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Upstream_Messaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagingMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).Messaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_Messaging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).Messaging(ctx, req.(*MessagingMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_PullAppFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullAppFileMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).PullAppFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_PullAppFileMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).PullAppFileMeta(ctx, req.(*PullAppFileMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Upstream_GetDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamServer).GetDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Upstream_GetDownloadURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamServer).GetDownloadURL(ctx, req.(*GetDownloadURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Upstream_ServiceDesc is the grpc.ServiceDesc for Upstream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upstream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbfs.Upstream",
	HandlerType: (*UpstreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Upstream_Handshake_Handler,
		},
		{
			MethodName: "Messaging",
			Handler:    _Upstream_Messaging_Handler,
		},
		{
			MethodName: "PullAppFileMeta",
			Handler:    _Upstream_PullAppFileMeta_Handler,
		},
		{
			MethodName: "GetDownloadURL",
			Handler:    _Upstream_GetDownloadURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Upstream_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "feed_server.proto",
}