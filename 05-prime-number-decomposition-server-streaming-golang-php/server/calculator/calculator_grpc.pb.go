// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculator

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	GetPrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_GetPrimeNumberDecompositionClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) GetPrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_GetPrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/GetPrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetPrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_GetPrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServiceGetPrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetPrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations should embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	GetPrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_GetPrimeNumberDecompositionServer) error
}

// UnimplementedCalculatorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) GetPrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_GetPrimeNumberDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrimeNumberDecomposition not implemented")
}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_GetPrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).GetPrimeNumberDecomposition(m, &calculatorServiceGetPrimeNumberDecompositionServer{stream})
}

type CalculatorService_GetPrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServiceGetPrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetPrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrimeNumberDecomposition",
			Handler:       _CalculatorService_GetPrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
