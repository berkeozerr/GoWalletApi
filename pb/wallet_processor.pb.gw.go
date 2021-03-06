// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: wallet_processor.proto

/*
Package pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_BitcoinWalletService_GenerateRandomMnemonic_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_BitcoinWalletService_GenerateRandomMnemonic_0(ctx context.Context, marshaler runtime.Marshaler, client BitcoinWalletServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RandomMnemonicRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BitcoinWalletService_GenerateRandomMnemonic_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GenerateRandomMnemonic(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BitcoinWalletService_GenerateRandomMnemonic_0(ctx context.Context, marshaler runtime.Marshaler, server BitcoinWalletServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RandomMnemonicRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_BitcoinWalletService_GenerateRandomMnemonic_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GenerateRandomMnemonic(ctx, &protoReq)
	return msg, metadata, err

}

func request_BitcoinWalletService_GenerateHDSegWitAddress_0(ctx context.Context, marshaler runtime.Marshaler, client BitcoinWalletServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateHDSegWitAddressRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GenerateHDSegWitAddress(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BitcoinWalletService_GenerateHDSegWitAddress_0(ctx context.Context, marshaler runtime.Marshaler, server BitcoinWalletServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateHDSegWitAddressRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GenerateHDSegWitAddress(ctx, &protoReq)
	return msg, metadata, err

}

func request_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(ctx context.Context, marshaler runtime.Marshaler, client BitcoinWalletServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateMultiSigP2SHAddressRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GenerateMultiSigP2SHAddress(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(ctx context.Context, marshaler runtime.Marshaler, server BitcoinWalletServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GenerateMultiSigP2SHAddressRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GenerateMultiSigP2SHAddress(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBitcoinWalletServiceHandlerServer registers the http handlers for service BitcoinWalletService to "mux".
// UnaryRPC     :call BitcoinWalletServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBitcoinWalletServiceHandlerFromEndpoint instead.
func RegisterBitcoinWalletServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BitcoinWalletServiceServer) error {

	mux.Handle("GET", pattern_BitcoinWalletService_GenerateRandomMnemonic_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateRandomMnemonic", runtime.WithHTTPPathPattern("/v1/wallet/random_mnemonic"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BitcoinWalletService_GenerateRandomMnemonic_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateRandomMnemonic_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BitcoinWalletService_GenerateHDSegWitAddress_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateHDSegWitAddress", runtime.WithHTTPPathPattern("/v1/wallet/generate_hd_seg_with"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BitcoinWalletService_GenerateHDSegWitAddress_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateHDSegWitAddress_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BitcoinWalletService_GenerateMultiSigP2SHAddress_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateMultiSigP2SHAddress", runtime.WithHTTPPathPattern("/v1/wallet/generate_multi_sig_p2sh"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBitcoinWalletServiceHandlerFromEndpoint is same as RegisterBitcoinWalletServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBitcoinWalletServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBitcoinWalletServiceHandler(ctx, mux, conn)
}

// RegisterBitcoinWalletServiceHandler registers the http handlers for service BitcoinWalletService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBitcoinWalletServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBitcoinWalletServiceHandlerClient(ctx, mux, NewBitcoinWalletServiceClient(conn))
}

// RegisterBitcoinWalletServiceHandlerClient registers the http handlers for service BitcoinWalletService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BitcoinWalletServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BitcoinWalletServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BitcoinWalletServiceClient" to call the correct interceptors.
func RegisterBitcoinWalletServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BitcoinWalletServiceClient) error {

	mux.Handle("GET", pattern_BitcoinWalletService_GenerateRandomMnemonic_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateRandomMnemonic", runtime.WithHTTPPathPattern("/v1/wallet/random_mnemonic"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BitcoinWalletService_GenerateRandomMnemonic_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateRandomMnemonic_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BitcoinWalletService_GenerateHDSegWitAddress_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateHDSegWitAddress", runtime.WithHTTPPathPattern("/v1/wallet/generate_hd_seg_with"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BitcoinWalletService_GenerateHDSegWitAddress_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateHDSegWitAddress_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BitcoinWalletService_GenerateMultiSigP2SHAddress_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/berkeozerr.GoWallet.BitcoinWalletService/GenerateMultiSigP2SHAddress", runtime.WithHTTPPathPattern("/v1/wallet/generate_multi_sig_p2sh"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BitcoinWalletService_GenerateMultiSigP2SHAddress_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BitcoinWalletService_GenerateRandomMnemonic_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "wallet", "random_mnemonic"}, ""))

	pattern_BitcoinWalletService_GenerateHDSegWitAddress_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "wallet", "generate_hd_seg_with"}, ""))

	pattern_BitcoinWalletService_GenerateMultiSigP2SHAddress_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "wallet", "generate_multi_sig_p2sh"}, ""))
)

var (
	forward_BitcoinWalletService_GenerateRandomMnemonic_0 = runtime.ForwardResponseMessage

	forward_BitcoinWalletService_GenerateHDSegWitAddress_0 = runtime.ForwardResponseMessage

	forward_BitcoinWalletService_GenerateMultiSigP2SHAddress_0 = runtime.ForwardResponseMessage
)
