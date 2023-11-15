// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: calculator.proto

package services

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
	Calculator_Hello_FullMethodName     = "/services.Calculator/Hello"
	Calculator_Fibonacci_FullMethodName = "/services.Calculator/Fibonacci"
	Calculator_Average_FullMethodName   = "/services.Calculator/Average"
	Calculator_Sum_FullMethodName       = "/services.Calculator/Sum"
)

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	// แบบ Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// แบบ Server Streaming
	// response แบบ streaming ตรง return จะใส streaming
	Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (Calculator_FibonacciClient, error)
	// แบบ Client Streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error)
	// แบบ Bi Direction Streaming
	Sum(ctx context.Context, opts ...grpc.CallOption) (Calculator_SumClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Calculator_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Fibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (Calculator_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], Calculator_Fibonacci_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Calculator_FibonacciClient interface {
	Recv() (*FibonacciResponse, error)
	grpc.ClientStream
}

type calculatorFibonacciClient struct {
	grpc.ClientStream
}

func (x *calculatorFibonacciClient) Recv() (*FibonacciResponse, error) {
	m := new(FibonacciResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Average(ctx context.Context, opts ...grpc.CallOption) (Calculator_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], Calculator_Average_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAverageClient{stream}
	return x, nil
}

type Calculator_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) Sum(ctx context.Context, opts ...grpc.CallOption) (Calculator_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[2], Calculator_Sum_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorSumClient{stream}
	return x, nil
}

type Calculator_SumClient interface {
	Send(*SumRequest) error
	Recv() (*SumResponse, error)
	grpc.ClientStream
}

type calculatorSumClient struct {
	grpc.ClientStream
}

func (x *calculatorSumClient) Send(m *SumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorSumClient) Recv() (*SumResponse, error) {
	m := new(SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	// แบบ Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// แบบ Server Streaming
	// response แบบ streaming ตรง return จะใส streaming
	Fibonacci(*FibonacciRequest, Calculator_FibonacciServer) error
	// แบบ Client Streaming
	Average(Calculator_AverageServer) error
	// แบบ Bi Direction Streaming
	Sum(Calculator_SumServer) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedCalculatorServer) Fibonacci(*FibonacciRequest, Calculator_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedCalculatorServer) Average(Calculator_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServer) Sum(Calculator_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calculator_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServer).Fibonacci(m, &calculatorFibonacciServer{stream})
}

type Calculator_FibonacciServer interface {
	Send(*FibonacciResponse) error
	grpc.ServerStream
}

type calculatorFibonacciServer struct {
	grpc.ServerStream
}

func (x *calculatorFibonacciServer) Send(m *FibonacciResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Calculator_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Average(&calculatorAverageServer{stream})
}

type Calculator_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Sum(&calculatorSumServer{stream})
}

type Calculator_SumServer interface {
	Send(*SumResponse) error
	Recv() (*SumRequest, error)
	grpc.ServerStream
}

type calculatorSumServer struct {
	grpc.ServerStream
}

func (x *calculatorSumServer) Send(m *SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorSumServer) Recv() (*SumRequest, error) {
	m := new(SumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Calculator_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fibonacci",
			Handler:       _Calculator_Fibonacci_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _Calculator_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Sum",
			Handler:       _Calculator_Sum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
