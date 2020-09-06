// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: apis/k8s/v1/service.proto

package k8s

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/thalescpl-io/k8s-kms-plugin/apis/kms/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_apis_k8s_v1_service_proto protoreflect.FileDescriptor

var file_apis_k8s_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x68, 0x61,
	0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb6,
	0x02, 0x0a, 0x14, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74,
	0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73,
	0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74,
	0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x28, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74,
	0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2d,
	0x69, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apis_k8s_v1_service_proto_goTypes = []interface{}{
	(*v1.VersionRequest)(nil),  // 0: thalescpl.io.ekms.kms.v1.VersionRequest
	(*DecryptRequest)(nil),     // 1: thalescpl.io.ekms.k8s.v1.DecryptRequest
	(*EncryptRequest)(nil),     // 2: thalescpl.io.ekms.k8s.v1.EncryptRequest
	(*v1.VersionResponse)(nil), // 3: thalescpl.io.ekms.kms.v1.VersionResponse
	(*DecryptResponse)(nil),    // 4: thalescpl.io.ekms.k8s.v1.DecryptResponse
	(*EncryptResponse)(nil),    // 5: thalescpl.io.ekms.k8s.v1.EncryptResponse
}
var file_apis_k8s_v1_service_proto_depIdxs = []int32{
	0, // 0: thalescpl.io.ekms.k8s.v1.KeyManagementService.Version:input_type -> thalescpl.io.ekms.kms.v1.VersionRequest
	1, // 1: thalescpl.io.ekms.k8s.v1.KeyManagementService.Decrypt:input_type -> thalescpl.io.ekms.k8s.v1.DecryptRequest
	2, // 2: thalescpl.io.ekms.k8s.v1.KeyManagementService.Encrypt:input_type -> thalescpl.io.ekms.k8s.v1.EncryptRequest
	3, // 3: thalescpl.io.ekms.k8s.v1.KeyManagementService.Version:output_type -> thalescpl.io.ekms.kms.v1.VersionResponse
	4, // 4: thalescpl.io.ekms.k8s.v1.KeyManagementService.Decrypt:output_type -> thalescpl.io.ekms.k8s.v1.DecryptResponse
	5, // 5: thalescpl.io.ekms.k8s.v1.KeyManagementService.Encrypt:output_type -> thalescpl.io.ekms.k8s.v1.EncryptResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_k8s_v1_service_proto_init() }
func file_apis_k8s_v1_service_proto_init() {
	if File_apis_k8s_v1_service_proto != nil {
		return
	}
	file_apis_k8s_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_k8s_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_k8s_v1_service_proto_goTypes,
		DependencyIndexes: file_apis_k8s_v1_service_proto_depIdxs,
	}.Build()
	File_apis_k8s_v1_service_proto = out.File
	file_apis_k8s_v1_service_proto_rawDesc = nil
	file_apis_k8s_v1_service_proto_goTypes = nil
	file_apis_k8s_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyManagementServiceClient is the client API for KeyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyManagementServiceClient interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(ctx context.Context, in *v1.VersionRequest, opts ...grpc.CallOption) (*v1.VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type keyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyManagementServiceClient(cc grpc.ClientConnInterface) KeyManagementServiceClient {
	return &keyManagementServiceClient{cc}
}

func (c *keyManagementServiceClient) Version(ctx context.Context, in *v1.VersionRequest, opts ...grpc.CallOption) (*v1.VersionResponse, error) {
	out := new(v1.VersionResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyManagementServiceServer is the server API for KeyManagementService service.
type KeyManagementServiceServer interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(context.Context, *v1.VersionRequest) (*v1.VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
}

// UnimplementedKeyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyManagementServiceServer struct {
}

func (*UnimplementedKeyManagementServiceServer) Version(context.Context, *v1.VersionRequest) (*v1.VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}

func RegisterKeyManagementServiceServer(s *grpc.Server, srv KeyManagementServiceServer) {
	s.RegisterService(&_KeyManagementService_serviceDesc, srv)
}

func _KeyManagementService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Version(ctx, req.(*v1.VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.k8s.v1.KeyManagementService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thalescpl.io.ekms.k8s.v1.KeyManagementService",
	HandlerType: (*KeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KeyManagementService_Version_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _KeyManagementService_Decrypt_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _KeyManagementService_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/k8s/v1/service.proto",
}