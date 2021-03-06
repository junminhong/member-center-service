// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/v1/delivery/grpc/proto/member.proto

package proto

import (
	context "context"
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

type AtomicTokenAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AtomicToken string `protobuf:"bytes,1,opt,name=atomicToken,proto3" json:"atomicToken,omitempty"`
}

func (x *AtomicTokenAuthRequest) Reset() {
	*x = AtomicTokenAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_delivery_grpc_proto_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicTokenAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicTokenAuthRequest) ProtoMessage() {}

func (x *AtomicTokenAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delivery_grpc_proto_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicTokenAuthRequest.ProtoReflect.Descriptor instead.
func (*AtomicTokenAuthRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_delivery_grpc_proto_member_proto_rawDescGZIP(), []int{0}
}

func (x *AtomicTokenAuthRequest) GetAtomicToken() string {
	if x != nil {
		return x.AtomicToken
	}
	return ""
}

type AtomicTokenAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberUUID string `protobuf:"bytes,2,opt,name=memberUUID,proto3" json:"memberUUID,omitempty"`
}

func (x *AtomicTokenAuthResponse) Reset() {
	*x = AtomicTokenAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_delivery_grpc_proto_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AtomicTokenAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AtomicTokenAuthResponse) ProtoMessage() {}

func (x *AtomicTokenAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_delivery_grpc_proto_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AtomicTokenAuthResponse.ProtoReflect.Descriptor instead.
func (*AtomicTokenAuthResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_delivery_grpc_proto_member_proto_rawDescGZIP(), []int{1}
}

func (x *AtomicTokenAuthResponse) GetMemberUUID() string {
	if x != nil {
		return x.MemberUUID
	}
	return ""
}

var File_api_v1_delivery_grpc_proto_member_proto protoreflect.FileDescriptor

var file_api_v1_delivery_grpc_proto_member_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3a, 0x0a, 0x16, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a,
	0x17, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x32, 0x67, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_delivery_grpc_proto_member_proto_rawDescOnce sync.Once
	file_api_v1_delivery_grpc_proto_member_proto_rawDescData = file_api_v1_delivery_grpc_proto_member_proto_rawDesc
)

func file_api_v1_delivery_grpc_proto_member_proto_rawDescGZIP() []byte {
	file_api_v1_delivery_grpc_proto_member_proto_rawDescOnce.Do(func() {
		file_api_v1_delivery_grpc_proto_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_delivery_grpc_proto_member_proto_rawDescData)
	})
	return file_api_v1_delivery_grpc_proto_member_proto_rawDescData
}

var file_api_v1_delivery_grpc_proto_member_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_delivery_grpc_proto_member_proto_goTypes = []interface{}{
	(*AtomicTokenAuthRequest)(nil),  // 0: member.AtomicTokenAuthRequest
	(*AtomicTokenAuthResponse)(nil), // 1: member.AtomicTokenAuthResponse
}
var file_api_v1_delivery_grpc_proto_member_proto_depIdxs = []int32{
	0, // 0: member.MemberService.VerifyAtomicToken:input_type -> member.AtomicTokenAuthRequest
	1, // 1: member.MemberService.VerifyAtomicToken:output_type -> member.AtomicTokenAuthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_delivery_grpc_proto_member_proto_init() }
func file_api_v1_delivery_grpc_proto_member_proto_init() {
	if File_api_v1_delivery_grpc_proto_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_delivery_grpc_proto_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtomicTokenAuthRequest); i {
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
		file_api_v1_delivery_grpc_proto_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AtomicTokenAuthResponse); i {
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
			RawDescriptor: file_api_v1_delivery_grpc_proto_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_delivery_grpc_proto_member_proto_goTypes,
		DependencyIndexes: file_api_v1_delivery_grpc_proto_member_proto_depIdxs,
		MessageInfos:      file_api_v1_delivery_grpc_proto_member_proto_msgTypes,
	}.Build()
	File_api_v1_delivery_grpc_proto_member_proto = out.File
	file_api_v1_delivery_grpc_proto_member_proto_rawDesc = nil
	file_api_v1_delivery_grpc_proto_member_proto_goTypes = nil
	file_api_v1_delivery_grpc_proto_member_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemberServiceClient interface {
	VerifyAtomicToken(ctx context.Context, in *AtomicTokenAuthRequest, opts ...grpc.CallOption) (*AtomicTokenAuthResponse, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) VerifyAtomicToken(ctx context.Context, in *AtomicTokenAuthRequest, opts ...grpc.CallOption) (*AtomicTokenAuthResponse, error) {
	out := new(AtomicTokenAuthResponse)
	err := c.cc.Invoke(ctx, "/member.MemberService/VerifyAtomicToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
type MemberServiceServer interface {
	VerifyAtomicToken(context.Context, *AtomicTokenAuthRequest) (*AtomicTokenAuthResponse, error)
}

// UnimplementedMemberServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (*UnimplementedMemberServiceServer) VerifyAtomicToken(context.Context, *AtomicTokenAuthRequest) (*AtomicTokenAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAtomicToken not implemented")
}

func RegisterMemberServiceServer(s *grpc.Server, srv MemberServiceServer) {
	s.RegisterService(&_MemberService_serviceDesc, srv)
}

func _MemberService_VerifyAtomicToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicTokenAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).VerifyAtomicToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/member.MemberService/VerifyAtomicToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).VerifyAtomicToken(ctx, req.(*AtomicTokenAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "member.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyAtomicToken",
			Handler:    _MemberService_VerifyAtomicToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/delivery/grpc/proto/member.proto",
}
