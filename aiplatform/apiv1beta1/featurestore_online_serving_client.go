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

package aiplatform

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
	aiplatformpb "google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newFeaturestoreOnlineServingClientHook clientHook

// FeaturestoreOnlineServingCallOptions contains the retry settings for each method of FeaturestoreOnlineServingClient.
type FeaturestoreOnlineServingCallOptions struct {
	ReadFeatureValues          []gax.CallOption
	StreamingReadFeatureValues []gax.CallOption
}

func defaultFeaturestoreOnlineServingGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFeaturestoreOnlineServingCallOptions() *FeaturestoreOnlineServingCallOptions {
	return &FeaturestoreOnlineServingCallOptions{
		ReadFeatureValues:          []gax.CallOption{},
		StreamingReadFeatureValues: []gax.CallOption{},
	}
}

// internalFeaturestoreOnlineServingClient is an interface that defines the methods availaible from Vertex AI API.
type internalFeaturestoreOnlineServingClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReadFeatureValues(context.Context, *aiplatformpb.ReadFeatureValuesRequest, ...gax.CallOption) (*aiplatformpb.ReadFeatureValuesResponse, error)
	StreamingReadFeatureValues(context.Context, *aiplatformpb.StreamingReadFeatureValuesRequest, ...gax.CallOption) (aiplatformpb.FeaturestoreOnlineServingService_StreamingReadFeatureValuesClient, error)
}

// FeaturestoreOnlineServingClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for serving online feature values.
type FeaturestoreOnlineServingClient struct {
	// The internal transport-dependent client.
	internalClient internalFeaturestoreOnlineServingClient

	// The call options for this service.
	CallOptions *FeaturestoreOnlineServingCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeaturestoreOnlineServingClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeaturestoreOnlineServingClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeaturestoreOnlineServingClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReadFeatureValues reads Feature values of a specific entity of an EntityType. For reading
// feature values of multiple entities of an EntityType, please use
// StreamingReadFeatureValues.
func (c *FeaturestoreOnlineServingClient) ReadFeatureValues(ctx context.Context, req *aiplatformpb.ReadFeatureValuesRequest, opts ...gax.CallOption) (*aiplatformpb.ReadFeatureValuesResponse, error) {
	return c.internalClient.ReadFeatureValues(ctx, req, opts...)
}

// StreamingReadFeatureValues reads Feature values for multiple entities. Depending on their size, data
// for different entities may be broken
// up across multiple responses.
func (c *FeaturestoreOnlineServingClient) StreamingReadFeatureValues(ctx context.Context, req *aiplatformpb.StreamingReadFeatureValuesRequest, opts ...gax.CallOption) (aiplatformpb.FeaturestoreOnlineServingService_StreamingReadFeatureValuesClient, error) {
	return c.internalClient.StreamingReadFeatureValues(ctx, req, opts...)
}

// featurestoreOnlineServingGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type featurestoreOnlineServingGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeaturestoreOnlineServingClient
	CallOptions **FeaturestoreOnlineServingCallOptions

	// The gRPC API client.
	featurestoreOnlineServingClient aiplatformpb.FeaturestoreOnlineServingServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeaturestoreOnlineServingClient creates a new featurestore online serving service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for serving online feature values.
func NewFeaturestoreOnlineServingClient(ctx context.Context, opts ...option.ClientOption) (*FeaturestoreOnlineServingClient, error) {
	clientOpts := defaultFeaturestoreOnlineServingGRPCClientOptions()
	if newFeaturestoreOnlineServingClientHook != nil {
		hookOpts, err := newFeaturestoreOnlineServingClientHook(ctx, clientHookParams{})
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
	client := FeaturestoreOnlineServingClient{CallOptions: defaultFeaturestoreOnlineServingCallOptions()}

	c := &featurestoreOnlineServingGRPCClient{
		connPool:                        connPool,
		disableDeadlines:                disableDeadlines,
		featurestoreOnlineServingClient: aiplatformpb.NewFeaturestoreOnlineServingServiceClient(connPool),
		CallOptions:                     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *featurestoreOnlineServingGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *featurestoreOnlineServingGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *featurestoreOnlineServingGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *featurestoreOnlineServingGRPCClient) ReadFeatureValues(ctx context.Context, req *aiplatformpb.ReadFeatureValuesRequest, opts ...gax.CallOption) (*aiplatformpb.ReadFeatureValuesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "entity_type", url.QueryEscape(req.GetEntityType())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReadFeatureValues[0:len((*c.CallOptions).ReadFeatureValues):len((*c.CallOptions).ReadFeatureValues)], opts...)
	var resp *aiplatformpb.ReadFeatureValuesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.featurestoreOnlineServingClient.ReadFeatureValues(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *featurestoreOnlineServingGRPCClient) StreamingReadFeatureValues(ctx context.Context, req *aiplatformpb.StreamingReadFeatureValuesRequest, opts ...gax.CallOption) (aiplatformpb.FeaturestoreOnlineServingService_StreamingReadFeatureValuesClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "entity_type", url.QueryEscape(req.GetEntityType())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	var resp aiplatformpb.FeaturestoreOnlineServingService_StreamingReadFeatureValuesClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.featurestoreOnlineServingClient.StreamingReadFeatureValues(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
