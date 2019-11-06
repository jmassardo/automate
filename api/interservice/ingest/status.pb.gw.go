// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/interservice/ingest/status.proto

/*
Package ingest is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ingest

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_IngestStatus_GetHealth_0(ctx context.Context, marshaler runtime.Marshaler, client IngestStatusClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HealthRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetHealth(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_IngestStatus_GetHealth_0(ctx context.Context, marshaler runtime.Marshaler, server IngestStatusServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HealthRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetHealth(ctx, &protoReq)
	return msg, metadata, err

}

func request_IngestStatus_GetMigrationStatus_0(ctx context.Context, marshaler runtime.Marshaler, client IngestStatusClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MigrationStatusRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetMigrationStatus(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_IngestStatus_GetMigrationStatus_0(ctx context.Context, marshaler runtime.Marshaler, server IngestStatusServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MigrationStatusRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetMigrationStatus(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterIngestStatusHandlerServer registers the http handlers for service IngestStatus to "mux".
// UnaryRPC     :call IngestStatusServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterIngestStatusHandlerServer(ctx context.Context, mux *runtime.ServeMux, server IngestStatusServer) error {

	mux.Handle("GET", pattern_IngestStatus_GetHealth_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_IngestStatus_GetHealth_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IngestStatus_GetHealth_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_IngestStatus_GetMigrationStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_IngestStatus_GetMigrationStatus_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IngestStatus_GetMigrationStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterIngestStatusHandlerFromEndpoint is same as RegisterIngestStatusHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterIngestStatusHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterIngestStatusHandler(ctx, mux, conn)
}

// RegisterIngestStatusHandler registers the http handlers for service IngestStatus to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterIngestStatusHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterIngestStatusHandlerClient(ctx, mux, NewIngestStatusClient(conn))
}

// RegisterIngestStatusHandlerClient registers the http handlers for service IngestStatus
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "IngestStatusClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "IngestStatusClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "IngestStatusClient" to call the correct interceptors.
func RegisterIngestStatusHandlerClient(ctx context.Context, mux *runtime.ServeMux, client IngestStatusClient) error {

	mux.Handle("GET", pattern_IngestStatus_GetHealth_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IngestStatus_GetHealth_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IngestStatus_GetHealth_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_IngestStatus_GetMigrationStatus_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IngestStatus_GetMigrationStatus_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IngestStatus_GetMigrationStatus_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_IngestStatus_GetHealth_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"health"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_IngestStatus_GetMigrationStatus_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"migration"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_IngestStatus_GetHealth_0 = runtime.ForwardResponseMessage

	forward_IngestStatus_GetMigrationStatus_0 = runtime.ForwardResponseMessage
)
