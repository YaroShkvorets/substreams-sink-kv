// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: substreams/sink/kv/v1/read.proto

package kvv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/streamingfast/substreams-sink-kv/pb/substreams/sink/kv/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// KvName is the fully-qualified name of the Kv service.
	KvName = "substreams.sink.kv.v1.Kv"
)

// KvClient is a client for the substreams.sink.kv.v1.Kv service.
type KvClient interface {
	// Get returns the requested value as bytes if it exists, grpc_error: NOT_FOUND otherwise.
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	// GetMany returns the requested values as bytes if all of them exists, grpc_error: NOT_FOUND otherwise.
	GetMany(context.Context, *connect_go.Request[v1.GetManyRequest]) (*connect_go.Response[v1.GetManyResponse], error)
	// GetByPrefix returns the next _limit_ key/value pair that match the requested prefix if any exist, grpc_error: NOT_FOUND otherwise.
	GetByPrefix(context.Context, *connect_go.Request[v1.GetByPrefixRequest]) (*connect_go.Response[v1.GetByPrefixResponse], error)
	// Scan returns then next _limit_ key/value pairs starting lexicographically at the given key, grpc_error: NOT_FOUND otherwise.
	Scan(context.Context, *connect_go.Request[v1.ScanRequest]) (*connect_go.Response[v1.ScanResponse], error)
}

// NewKvClient constructs a client for the substreams.sink.kv.v1.Kv service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewKvClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) KvClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &kvClient{
		get: connect_go.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+"/substreams.sink.kv.v1.Kv/Get",
			opts...,
		),
		getMany: connect_go.NewClient[v1.GetManyRequest, v1.GetManyResponse](
			httpClient,
			baseURL+"/substreams.sink.kv.v1.Kv/GetMany",
			opts...,
		),
		getByPrefix: connect_go.NewClient[v1.GetByPrefixRequest, v1.GetByPrefixResponse](
			httpClient,
			baseURL+"/substreams.sink.kv.v1.Kv/GetByPrefix",
			opts...,
		),
		scan: connect_go.NewClient[v1.ScanRequest, v1.ScanResponse](
			httpClient,
			baseURL+"/substreams.sink.kv.v1.Kv/Scan",
			opts...,
		),
	}
}

// kvClient implements KvClient.
type kvClient struct {
	get         *connect_go.Client[v1.GetRequest, v1.GetResponse]
	getMany     *connect_go.Client[v1.GetManyRequest, v1.GetManyResponse]
	getByPrefix *connect_go.Client[v1.GetByPrefixRequest, v1.GetByPrefixResponse]
	scan        *connect_go.Client[v1.ScanRequest, v1.ScanResponse]
}

// Get calls substreams.sink.kv.v1.Kv.Get.
func (c *kvClient) Get(ctx context.Context, req *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// GetMany calls substreams.sink.kv.v1.Kv.GetMany.
func (c *kvClient) GetMany(ctx context.Context, req *connect_go.Request[v1.GetManyRequest]) (*connect_go.Response[v1.GetManyResponse], error) {
	return c.getMany.CallUnary(ctx, req)
}

// GetByPrefix calls substreams.sink.kv.v1.Kv.GetByPrefix.
func (c *kvClient) GetByPrefix(ctx context.Context, req *connect_go.Request[v1.GetByPrefixRequest]) (*connect_go.Response[v1.GetByPrefixResponse], error) {
	return c.getByPrefix.CallUnary(ctx, req)
}

// Scan calls substreams.sink.kv.v1.Kv.Scan.
func (c *kvClient) Scan(ctx context.Context, req *connect_go.Request[v1.ScanRequest]) (*connect_go.Response[v1.ScanResponse], error) {
	return c.scan.CallUnary(ctx, req)
}

// KvHandler is an implementation of the substreams.sink.kv.v1.Kv service.
type KvHandler interface {
	// Get returns the requested value as bytes if it exists, grpc_error: NOT_FOUND otherwise.
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
	// GetMany returns the requested values as bytes if all of them exists, grpc_error: NOT_FOUND otherwise.
	GetMany(context.Context, *connect_go.Request[v1.GetManyRequest]) (*connect_go.Response[v1.GetManyResponse], error)
	// GetByPrefix returns the next _limit_ key/value pair that match the requested prefix if any exist, grpc_error: NOT_FOUND otherwise.
	GetByPrefix(context.Context, *connect_go.Request[v1.GetByPrefixRequest]) (*connect_go.Response[v1.GetByPrefixResponse], error)
	// Scan returns then next _limit_ key/value pairs starting lexicographically at the given key, grpc_error: NOT_FOUND otherwise.
	Scan(context.Context, *connect_go.Request[v1.ScanRequest]) (*connect_go.Response[v1.ScanResponse], error)
}

// NewKvHandler builds an HTTP handler from the service implementation. It returns the path on which
// to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewKvHandler(svc KvHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/substreams.sink.kv.v1.Kv/Get", connect_go.NewUnaryHandler(
		"/substreams.sink.kv.v1.Kv/Get",
		svc.Get,
		opts...,
	))
	mux.Handle("/substreams.sink.kv.v1.Kv/GetMany", connect_go.NewUnaryHandler(
		"/substreams.sink.kv.v1.Kv/GetMany",
		svc.GetMany,
		opts...,
	))
	mux.Handle("/substreams.sink.kv.v1.Kv/GetByPrefix", connect_go.NewUnaryHandler(
		"/substreams.sink.kv.v1.Kv/GetByPrefix",
		svc.GetByPrefix,
		opts...,
	))
	mux.Handle("/substreams.sink.kv.v1.Kv/Scan", connect_go.NewUnaryHandler(
		"/substreams.sink.kv.v1.Kv/Scan",
		svc.Scan,
		opts...,
	))
	return "/substreams.sink.kv.v1.Kv/", mux
}

// UnimplementedKvHandler returns CodeUnimplemented from all methods.
type UnimplementedKvHandler struct{}

func (UnimplementedKvHandler) Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("substreams.sink.kv.v1.Kv.Get is not implemented"))
}

func (UnimplementedKvHandler) GetMany(context.Context, *connect_go.Request[v1.GetManyRequest]) (*connect_go.Response[v1.GetManyResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("substreams.sink.kv.v1.Kv.GetMany is not implemented"))
}

func (UnimplementedKvHandler) GetByPrefix(context.Context, *connect_go.Request[v1.GetByPrefixRequest]) (*connect_go.Response[v1.GetByPrefixResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("substreams.sink.kv.v1.Kv.GetByPrefix is not implemented"))
}

func (UnimplementedKvHandler) Scan(context.Context, *connect_go.Request[v1.ScanRequest]) (*connect_go.Response[v1.ScanResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("substreams.sink.kv.v1.Kv.Scan is not implemented"))
}
