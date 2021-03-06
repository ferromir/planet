// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("planet/query.proto", fileDescriptor_2b55d081339b1f1e) }

var fileDescriptor_2b55d081339b1f1e = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0x6e, 0x83, 0x30,
	0x10, 0x86, 0x61, 0x68, 0x2b, 0x31, 0x32, 0x22, 0xe4, 0x07, 0xa8, 0x54, 0xae, 0xb4, 0x6f, 0xd0,
	0xb1, 0x5b, 0xd7, 0x6e, 0x67, 0x72, 0x71, 0x2c, 0x81, 0xcf, 0xb1, 0x4d, 0x14, 0xde, 0x22, 0x8f,
	0x95, 0x91, 0x31, 0x63, 0x04, 0x2f, 0x12, 0x05, 0xc3, 0x74, 0x37, 0x7c, 0xff, 0xa7, 0x2f, 0xcb,
	0x6d, 0x8b, 0x86, 0x02, 0x1c, 0x7b, 0x72, 0x43, 0x65, 0x1d, 0x07, 0xce, 0xcb, 0x3d, 0x39, 0x83,
	0x66, 0xc7, 0x8e, 0x3b, 0x72, 0x5c, 0x45, 0x64, 0x3d, 0x45, 0xa9, 0x98, 0x55, 0x4b, 0x80, 0x56,
	0x03, 0x1a, 0xc3, 0x01, 0x83, 0x66, 0xe3, 0xe3, 0xb6, 0x78, 0x6f, 0xd8, 0x77, 0xec, 0x41, 0xa2,
	0xa7, 0x28, 0x85, 0x53, 0x2d, 0x29, 0x60, 0x0d, 0x16, 0x95, 0x36, 0x0b, 0x1c, 0xd9, 0xaf, 0xb7,
	0xec, 0xe5, 0xef, 0x49, 0xfc, 0xfc, 0x5e, 0x27, 0x91, 0x8e, 0x93, 0x48, 0xef, 0x93, 0x48, 0x2f,
	0xb3, 0x48, 0xc6, 0x59, 0x24, 0xb7, 0x59, 0x24, 0xff, 0x9f, 0x4a, 0x87, 0x43, 0x2f, 0xab, 0x86,
	0x3b, 0xd8, 0xaa, 0x3e, 0x62, 0x16, 0xac, 0xe5, 0xe7, 0xed, 0x09, 0x83, 0x25, 0x2f, 0x5f, 0x17,
	0xf7, 0xf7, 0x23, 0x00, 0x00, 0xff, 0xff, 0xad, 0xca, 0x76, 0xb0, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fernandoromero.planet.planet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "planet/query.proto",
}
