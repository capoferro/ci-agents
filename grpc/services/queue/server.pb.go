// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/services/queue/server.proto

package queue

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	types "github.com/tinyci/ci-agents/grpc/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Submission controls the submission of branches and pull requests. Some
// (noted) properties are not externally modifiable, so they will result in a
// noop if set.
type Submission struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Fork                 string   `protobuf:"bytes,2,opt,name=fork,proto3" json:"fork,omitempty"`
	Headsha              string   `protobuf:"bytes,3,opt,name=headsha,proto3" json:"headsha,omitempty"`
	Basesha              string   `protobuf:"bytes,4,opt,name=basesha,proto3" json:"basesha,omitempty"`
	SubmittedBy          string   `protobuf:"bytes,5,opt,name=submitted_by,json=submittedBy,proto3" json:"submitted_by,omitempty"`
	PullRequest          int64    `protobuf:"varint,6,opt,name=pull_request,json=pullRequest,proto3" json:"pull_request,omitempty"`
	All                  bool     `protobuf:"varint,7,opt,name=all,proto3" json:"all,omitempty"`
	Manual               bool     `protobuf:"varint,8,opt,name=manual,proto3" json:"manual,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Submission) Reset()         { *m = Submission{} }
func (m *Submission) String() string { return proto.CompactTextString(m) }
func (*Submission) ProtoMessage()    {}
func (*Submission) Descriptor() ([]byte, []int) {
	return fileDescriptor_18dba3f4b885f33f, []int{0}
}

func (m *Submission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Submission.Unmarshal(m, b)
}
func (m *Submission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Submission.Marshal(b, m, deterministic)
}
func (m *Submission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Submission.Merge(m, src)
}
func (m *Submission) XXX_Size() int {
	return xxx_messageInfo_Submission.Size(m)
}
func (m *Submission) XXX_DiscardUnknown() {
	xxx_messageInfo_Submission.DiscardUnknown(m)
}

var xxx_messageInfo_Submission proto.InternalMessageInfo

func (m *Submission) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Submission) GetFork() string {
	if m != nil {
		return m.Fork
	}
	return ""
}

func (m *Submission) GetHeadsha() string {
	if m != nil {
		return m.Headsha
	}
	return ""
}

func (m *Submission) GetBasesha() string {
	if m != nil {
		return m.Basesha
	}
	return ""
}

func (m *Submission) GetSubmittedBy() string {
	if m != nil {
		return m.SubmittedBy
	}
	return ""
}

func (m *Submission) GetPullRequest() int64 {
	if m != nil {
		return m.PullRequest
	}
	return 0
}

func (m *Submission) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *Submission) GetManual() bool {
	if m != nil {
		return m.Manual
	}
	return false
}

func init() {
	proto.RegisterType((*Submission)(nil), "queue.Submission")
}

func init() { proto.RegisterFile("grpc/services/queue/server.proto", fileDescriptor_18dba3f4b885f33f) }

var fileDescriptor_18dba3f4b885f33f = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb3, 0x4d, 0x93, 0x36, 0x6e, 0x2b, 0x15, 0x23, 0x55, 0x56, 0x4e, 0x21, 0xa7, 0x08,
	0x09, 0x5b, 0x10, 0x10, 0x9c, 0x0b, 0x08, 0xe5, 0x82, 0x20, 0xb9, 0x71, 0x89, 0xbc, 0x9b, 0xe9,
	0xc6, 0xc2, 0x6b, 0x6f, 0xd7, 0x63, 0xc4, 0x3e, 0x05, 0x2f, 0xc8, 0xc3, 0x20, 0xcf, 0x6e, 0xa1,
	0x3d, 0x10, 0xb8, 0xac, 0xfc, 0xff, 0xe3, 0xcf, 0xbf, 0x77, 0xc6, 0x6c, 0x56, 0x36, 0x75, 0xa1,
	0x02, 0x34, 0xdf, 0x4c, 0x01, 0x41, 0xdd, 0x46, 0x88, 0x40, 0x12, 0x1a, 0x59, 0x37, 0x1e, 0x3d,
	0x1f, 0x91, 0x37, 0x5d, 0x96, 0x06, 0xf7, 0x31, 0x97, 0x85, 0xaf, 0x54, 0xe9, 0xad, 0x76, 0xa5,
	0xa2, 0x7a, 0x1e, 0x6f, 0x54, 0x8d, 0x6d, 0x0d, 0x41, 0x41, 0x55, 0x63, 0xdb, 0x7d, 0x3b, 0x76,
	0xfa, 0xfa, 0x1e, 0x84, 0xc6, 0xb5, 0x85, 0x51, 0x85, 0x79, 0xa6, 0x4b, 0x70, 0x18, 0x14, 0x25,
	0x77, 0x24, 0x45, 0x6c, 0x0d, 0x42, 0xd5, 0x83, 0xea, 0x3f, 0x41, 0xb3, 0xeb, 0x80, 0xf9, 0xcf,
	0x8c, 0xb1, 0x4d, 0xcc, 0x2b, 0x13, 0x82, 0xf1, 0x8e, 0x5f, 0xb1, 0x71, 0xad, 0x1b, 0x70, 0x28,
	0xb2, 0x59, 0xb6, 0x98, 0xac, 0x7b, 0xc5, 0x39, 0x3b, 0xbe, 0xf1, 0xcd, 0x57, 0x71, 0x44, 0x2e,
	0xad, 0xb9, 0x60, 0x27, 0x7b, 0xd0, 0xbb, 0xb0, 0xd7, 0x62, 0x48, 0xf6, 0x9d, 0x4c, 0x95, 0x5c,
	0x07, 0x48, 0x95, 0xe3, 0xae, 0xd2, 0x4b, 0xfe, 0x84, 0x9d, 0x87, 0x94, 0x86, 0x08, 0xbb, 0x6d,
	0xde, 0x8a, 0x11, 0x95, 0xcf, 0x7e, 0x7b, 0xd7, 0x6d, 0xda, 0x52, 0x47, 0x6b, 0xb7, 0x0d, 0xdc,
	0x46, 0x08, 0x28, 0xc6, 0xb3, 0x6c, 0x31, 0x5c, 0x9f, 0x25, 0x6f, 0xdd, 0x59, 0xfc, 0x92, 0x0d,
	0xb5, 0xb5, 0xe2, 0x64, 0x96, 0x2d, 0x4e, 0xd7, 0x69, 0x99, 0xee, 0x5d, 0x69, 0x17, 0xb5, 0x15,
	0xa7, 0x64, 0xf6, 0xea, 0xc5, 0x8f, 0x23, 0x36, 0xfa, 0x9c, 0x9a, 0xc4, 0x5f, 0xb2, 0xc9, 0xa7,
	0x88, 0x1b, 0xd4, 0x18, 0x03, 0xbf, 0x90, 0xd4, 0x06, 0xd9, 0xc9, 0xe9, 0x95, 0x2c, 0xbd, 0x2f,
	0x2d, 0xc8, 0xbb, 0xc9, 0xc8, 0xf7, 0x69, 0x18, 0xf3, 0x01, 0x7f, 0xc3, 0x2e, 0x3e, 0xc2, 0x77,
	0xa4, 0x23, 0x56, 0x08, 0x15, 0x7f, 0xdc, 0x93, 0xe4, 0xf4, 0xf7, 0x99, 0x5e, 0xde, 0x37, 0xd3,
	0xb6, 0xf9, 0x80, 0xbf, 0x62, 0x63, 0xea, 0x2b, 0xf2, 0x47, 0x92, 0xc6, 0x24, 0xff, 0xb4, 0xf9,
	0x40, 0xe0, 0x92, 0x4d, 0x36, 0x80, 0x6f, 0xb5, 0x2b, 0xc0, 0xf2, 0xf3, 0xfe, 0xdc, 0x95, 0xc3,
	0xd5, 0xbb, 0x03, 0xd0, 0x53, 0x36, 0xf9, 0xf0, 0x17, 0xe8, 0xe1, 0x9f, 0xce, 0x07, 0xd7, 0xcf,
	0xbf, 0xfc, 0xfb, 0x8d, 0x3c, 0x7c, 0xd6, 0xf9, 0x98, 0x02, 0x97, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x51, 0xa8, 0x7e, 0xfb, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueClient interface {
	PutStatus(ctx context.Context, in *types.Status, opts ...grpc.CallOption) (*empty.Empty, error)
	NextQueueItem(ctx context.Context, in *types.QueueRequest, opts ...grpc.CallOption) (*types.QueueItem, error)
	Submit(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*empty.Empty, error)
	SetCancel(ctx context.Context, in *types.IntID, opts ...grpc.CallOption) (*empty.Empty, error)
	GetCancel(ctx context.Context, in *types.IntID, opts ...grpc.CallOption) (*types.Status, error)
}

type queueClient struct {
	cc *grpc.ClientConn
}

func NewQueueClient(cc *grpc.ClientConn) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) PutStatus(ctx context.Context, in *types.Status, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/queue.Queue/PutStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) NextQueueItem(ctx context.Context, in *types.QueueRequest, opts ...grpc.CallOption) (*types.QueueItem, error) {
	out := new(types.QueueItem)
	err := c.cc.Invoke(ctx, "/queue.Queue/NextQueueItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Submit(ctx context.Context, in *Submission, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/queue.Queue/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) SetCancel(ctx context.Context, in *types.IntID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/queue.Queue/SetCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) GetCancel(ctx context.Context, in *types.IntID, opts ...grpc.CallOption) (*types.Status, error) {
	out := new(types.Status)
	err := c.cc.Invoke(ctx, "/queue.Queue/GetCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServer is the server API for Queue service.
type QueueServer interface {
	PutStatus(context.Context, *types.Status) (*empty.Empty, error)
	NextQueueItem(context.Context, *types.QueueRequest) (*types.QueueItem, error)
	Submit(context.Context, *Submission) (*empty.Empty, error)
	SetCancel(context.Context, *types.IntID) (*empty.Empty, error)
	GetCancel(context.Context, *types.IntID) (*types.Status, error)
}

// UnimplementedQueueServer can be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (*UnimplementedQueueServer) PutStatus(ctx context.Context, req *types.Status) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStatus not implemented")
}
func (*UnimplementedQueueServer) NextQueueItem(ctx context.Context, req *types.QueueRequest) (*types.QueueItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextQueueItem not implemented")
}
func (*UnimplementedQueueServer) Submit(ctx context.Context, req *Submission) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (*UnimplementedQueueServer) SetCancel(ctx context.Context, req *types.IntID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCancel not implemented")
}
func (*UnimplementedQueueServer) GetCancel(ctx context.Context, req *types.IntID) (*types.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCancel not implemented")
}

func RegisterQueueServer(s *grpc.Server, srv QueueServer) {
	s.RegisterService(&_Queue_serviceDesc, srv)
}

func _Queue_PutStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).PutStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue.Queue/PutStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).PutStatus(ctx, req.(*types.Status))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_NextQueueItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).NextQueueItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue.Queue/NextQueueItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).NextQueueItem(ctx, req.(*types.QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Submission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue.Queue/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Submit(ctx, req.(*Submission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_SetCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.IntID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).SetCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue.Queue/SetCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).SetCancel(ctx, req.(*types.IntID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_GetCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.IntID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).GetCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queue.Queue/GetCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).GetCancel(ctx, req.(*types.IntID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Queue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "queue.Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutStatus",
			Handler:    _Queue_PutStatus_Handler,
		},
		{
			MethodName: "NextQueueItem",
			Handler:    _Queue_NextQueueItem_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _Queue_Submit_Handler,
		},
		{
			MethodName: "SetCancel",
			Handler:    _Queue_SetCancel_Handler,
		},
		{
			MethodName: "GetCancel",
			Handler:    _Queue_GetCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/services/queue/server.proto",
}