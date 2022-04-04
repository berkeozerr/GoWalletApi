// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// RandomMnemonicGeneratorClient is the client API for RandomMnemonicGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomMnemonicGeneratorClient interface {
	GenerateRandomMnemonic(ctx context.Context, in *RandomMnemonicRequest, opts ...grpc.CallOption) (*RandomMnemonicResponse, error)
	GenerateHDSegWitAddress(ctx context.Context, in *GenerateHDSegWitAddressRequest, opts ...grpc.CallOption) (*GenerateHDSegWitAddressResponse, error)
	GenerateMultiSigP2SHAddress(ctx context.Context, in *GenerateMultiSigP2SHAddressRequest, opts ...grpc.CallOption) (*GenerateMultiSigP2SHAddressResponse, error)
}

type randomMnemonicGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomMnemonicGeneratorClient(cc grpc.ClientConnInterface) RandomMnemonicGeneratorClient {
	return &randomMnemonicGeneratorClient{cc}
}

func (c *randomMnemonicGeneratorClient) GenerateRandomMnemonic(ctx context.Context, in *RandomMnemonicRequest, opts ...grpc.CallOption) (*RandomMnemonicResponse, error) {
	out := new(RandomMnemonicResponse)
	err := c.cc.Invoke(ctx, "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateRandomMnemonic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomMnemonicGeneratorClient) GenerateHDSegWitAddress(ctx context.Context, in *GenerateHDSegWitAddressRequest, opts ...grpc.CallOption) (*GenerateHDSegWitAddressResponse, error) {
	out := new(GenerateHDSegWitAddressResponse)
	err := c.cc.Invoke(ctx, "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateHDSegWitAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomMnemonicGeneratorClient) GenerateMultiSigP2SHAddress(ctx context.Context, in *GenerateMultiSigP2SHAddressRequest, opts ...grpc.CallOption) (*GenerateMultiSigP2SHAddressResponse, error) {
	out := new(GenerateMultiSigP2SHAddressResponse)
	err := c.cc.Invoke(ctx, "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateMultiSigP2SHAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomMnemonicGeneratorServer is the server API for RandomMnemonicGenerator service.
// All implementations must embed UnimplementedRandomMnemonicGeneratorServer
// for forward compatibility
type RandomMnemonicGeneratorServer interface {
	GenerateRandomMnemonic(context.Context, *RandomMnemonicRequest) (*RandomMnemonicResponse, error)
	GenerateHDSegWitAddress(context.Context, *GenerateHDSegWitAddressRequest) (*GenerateHDSegWitAddressResponse, error)
	GenerateMultiSigP2SHAddress(context.Context, *GenerateMultiSigP2SHAddressRequest) (*GenerateMultiSigP2SHAddressResponse, error)
	mustEmbedUnimplementedRandomMnemonicGeneratorServer()
}

// UnimplementedRandomMnemonicGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedRandomMnemonicGeneratorServer struct {
}

func (UnimplementedRandomMnemonicGeneratorServer) GenerateRandomMnemonic(context.Context, *RandomMnemonicRequest) (*RandomMnemonicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRandomMnemonic not implemented")
}
func (UnimplementedRandomMnemonicGeneratorServer) GenerateHDSegWitAddress(context.Context, *GenerateHDSegWitAddressRequest) (*GenerateHDSegWitAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHDSegWitAddress not implemented")
}
func (UnimplementedRandomMnemonicGeneratorServer) GenerateMultiSigP2SHAddress(context.Context, *GenerateMultiSigP2SHAddressRequest) (*GenerateMultiSigP2SHAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMultiSigP2SHAddress not implemented")
}
func (UnimplementedRandomMnemonicGeneratorServer) mustEmbedUnimplementedRandomMnemonicGeneratorServer() {
}

// UnsafeRandomMnemonicGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomMnemonicGeneratorServer will
// result in compilation errors.
type UnsafeRandomMnemonicGeneratorServer interface {
	mustEmbedUnimplementedRandomMnemonicGeneratorServer()
}

func RegisterRandomMnemonicGeneratorServer(s grpc.ServiceRegistrar, srv RandomMnemonicGeneratorServer) {
	s.RegisterService(&RandomMnemonicGenerator_ServiceDesc, srv)
}

func _RandomMnemonicGenerator_GenerateRandomMnemonic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomMnemonicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomMnemonicGeneratorServer).GenerateRandomMnemonic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateRandomMnemonic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomMnemonicGeneratorServer).GenerateRandomMnemonic(ctx, req.(*RandomMnemonicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomMnemonicGenerator_GenerateHDSegWitAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHDSegWitAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomMnemonicGeneratorServer).GenerateHDSegWitAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateHDSegWitAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomMnemonicGeneratorServer).GenerateHDSegWitAddress(ctx, req.(*GenerateHDSegWitAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomMnemonicGenerator_GenerateMultiSigP2SHAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMultiSigP2SHAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomMnemonicGeneratorServer).GenerateMultiSigP2SHAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berkeozerr.GoWallet.RandomMnemonicGenerator/GenerateMultiSigP2SHAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomMnemonicGeneratorServer).GenerateMultiSigP2SHAddress(ctx, req.(*GenerateMultiSigP2SHAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RandomMnemonicGenerator_ServiceDesc is the grpc.ServiceDesc for RandomMnemonicGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomMnemonicGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "berkeozerr.GoWallet.RandomMnemonicGenerator",
	HandlerType: (*RandomMnemonicGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateRandomMnemonic",
			Handler:    _RandomMnemonicGenerator_GenerateRandomMnemonic_Handler,
		},
		{
			MethodName: "GenerateHDSegWitAddress",
			Handler:    _RandomMnemonicGenerator_GenerateHDSegWitAddress_Handler,
		},
		{
			MethodName: "GenerateMultiSigP2SHAddress",
			Handler:    _RandomMnemonicGenerator_GenerateMultiSigP2SHAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_processor.proto",
}
