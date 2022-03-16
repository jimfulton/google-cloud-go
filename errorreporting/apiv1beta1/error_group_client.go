// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package errorreporting

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newErrorGroupClientHook clientHook

// ErrorGroupCallOptions contains the retry settings for each method of ErrorGroupClient.
type ErrorGroupCallOptions struct {
	GetGroup    []gax.CallOption
	UpdateGroup []gax.CallOption
}

func defaultErrorGroupGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("clouderrorreporting.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("clouderrorreporting.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultErrorGroupCallOptions() *ErrorGroupCallOptions {
	return &ErrorGroupCallOptions{
		GetGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalErrorGroupClient is an interface that defines the methods availaible from Error Reporting API.
type internalErrorGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetGroup(context.Context, *clouderrorreportingpb.GetGroupRequest, ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error)
	UpdateGroup(context.Context, *clouderrorreportingpb.UpdateGroupRequest, ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error)
}

// ErrorGroupClient is a client for interacting with Error Reporting API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for retrieving and updating individual error groups.
type ErrorGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalErrorGroupClient

	// The call options for this service.
	CallOptions *ErrorGroupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ErrorGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ErrorGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ErrorGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetGroup get the specified group.
func (c *ErrorGroupClient) GetGroup(ctx context.Context, req *clouderrorreportingpb.GetGroupRequest, opts ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error) {
	return c.internalClient.GetGroup(ctx, req, opts...)
}

// UpdateGroup replace the data for the specified group.
// Fails if the group does not exist.
func (c *ErrorGroupClient) UpdateGroup(ctx context.Context, req *clouderrorreportingpb.UpdateGroupRequest, opts ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error) {
	return c.internalClient.UpdateGroup(ctx, req, opts...)
}

// errorGroupGRPCClient is a client for interacting with Error Reporting API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type errorGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ErrorGroupClient
	CallOptions **ErrorGroupCallOptions

	// The gRPC API client.
	errorGroupClient clouderrorreportingpb.ErrorGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewErrorGroupClient creates a new error group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for retrieving and updating individual error groups.
func NewErrorGroupClient(ctx context.Context, opts ...option.ClientOption) (*ErrorGroupClient, error) {
	clientOpts := defaultErrorGroupGRPCClientOptions()
	if newErrorGroupClientHook != nil {
		hookOpts, err := newErrorGroupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ErrorGroupClient{CallOptions: defaultErrorGroupCallOptions()}

	c := &errorGroupGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		errorGroupClient: clouderrorreportingpb.NewErrorGroupServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *errorGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *errorGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *errorGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *errorGroupGRPCClient) GetGroup(ctx context.Context, req *clouderrorreportingpb.GetGroupRequest, opts ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group_name", url.QueryEscape(req.GetGroupName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGroup[0:len((*c.CallOptions).GetGroup):len((*c.CallOptions).GetGroup)], opts...)
	var resp *clouderrorreportingpb.ErrorGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorGroupClient.GetGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *errorGroupGRPCClient) UpdateGroup(ctx context.Context, req *clouderrorreportingpb.UpdateGroupRequest, opts ...gax.CallOption) (*clouderrorreportingpb.ErrorGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "group.name", url.QueryEscape(req.GetGroup().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateGroup[0:len((*c.CallOptions).UpdateGroup):len((*c.CallOptions).UpdateGroup)], opts...)
	var resp *clouderrorreportingpb.ErrorGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorGroupClient.UpdateGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
