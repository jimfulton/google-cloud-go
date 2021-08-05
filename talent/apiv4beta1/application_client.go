// Copyright 2021 Google LLC
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

package talent

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newApplicationClientHook clientHook

// ApplicationCallOptions contains the retry settings for each method of ApplicationClient.
type ApplicationCallOptions struct {
	CreateApplication []gax.CallOption
	GetApplication    []gax.CallOption
	UpdateApplication []gax.CallOption
	DeleteApplication []gax.CallOption
	ListApplications  []gax.CallOption
}

func defaultApplicationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://jobs.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultApplicationCallOptions() *ApplicationCallOptions {
	return &ApplicationCallOptions{
		CreateApplication: []gax.CallOption{},
		GetApplication: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateApplication: []gax.CallOption{},
		DeleteApplication: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListApplications: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalApplicationClient is an interface that defines the methods availaible from Cloud Talent Solution API.
type internalApplicationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateApplication(context.Context, *talentpb.CreateApplicationRequest, ...gax.CallOption) (*talentpb.Application, error)
	GetApplication(context.Context, *talentpb.GetApplicationRequest, ...gax.CallOption) (*talentpb.Application, error)
	UpdateApplication(context.Context, *talentpb.UpdateApplicationRequest, ...gax.CallOption) (*talentpb.Application, error)
	DeleteApplication(context.Context, *talentpb.DeleteApplicationRequest, ...gax.CallOption) error
	ListApplications(context.Context, *talentpb.ListApplicationsRequest, ...gax.CallOption) *ApplicationIterator
}

// ApplicationClient is a client for interacting with Cloud Talent Solution API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that handles application management, including CRUD and
// enumeration.
type ApplicationClient struct {
	// The internal transport-dependent client.
	internalClient internalApplicationClient

	// The call options for this service.
	CallOptions *ApplicationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ApplicationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ApplicationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ApplicationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateApplication creates a new application entity.
func (c *ApplicationClient) CreateApplication(ctx context.Context, req *talentpb.CreateApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	return c.internalClient.CreateApplication(ctx, req, opts...)
}

// GetApplication retrieves specified application.
func (c *ApplicationClient) GetApplication(ctx context.Context, req *talentpb.GetApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	return c.internalClient.GetApplication(ctx, req, opts...)
}

// UpdateApplication updates specified application.
func (c *ApplicationClient) UpdateApplication(ctx context.Context, req *talentpb.UpdateApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	return c.internalClient.UpdateApplication(ctx, req, opts...)
}

// DeleteApplication deletes specified application.
func (c *ApplicationClient) DeleteApplication(ctx context.Context, req *talentpb.DeleteApplicationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteApplication(ctx, req, opts...)
}

// ListApplications lists all applications associated with the profile.
func (c *ApplicationClient) ListApplications(ctx context.Context, req *talentpb.ListApplicationsRequest, opts ...gax.CallOption) *ApplicationIterator {
	return c.internalClient.ListApplications(ctx, req, opts...)
}

// applicationGRPCClient is a client for interacting with Cloud Talent Solution API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type applicationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ApplicationClient
	CallOptions **ApplicationCallOptions

	// The gRPC API client.
	applicationClient talentpb.ApplicationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewApplicationClient creates a new application service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that handles application management, including CRUD and
// enumeration.
func NewApplicationClient(ctx context.Context, opts ...option.ClientOption) (*ApplicationClient, error) {
	clientOpts := defaultApplicationGRPCClientOptions()
	if newApplicationClientHook != nil {
		hookOpts, err := newApplicationClientHook(ctx, clientHookParams{})
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
	client := ApplicationClient{CallOptions: defaultApplicationCallOptions()}

	c := &applicationGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		applicationClient: talentpb.NewApplicationServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *applicationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *applicationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *applicationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *applicationGRPCClient) CreateApplication(ctx context.Context, req *talentpb.CreateApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateApplication[0:len((*c.CallOptions).CreateApplication):len((*c.CallOptions).CreateApplication)], opts...)
	var resp *talentpb.Application
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationClient.CreateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *applicationGRPCClient) GetApplication(ctx context.Context, req *talentpb.GetApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetApplication[0:len((*c.CallOptions).GetApplication):len((*c.CallOptions).GetApplication)], opts...)
	var resp *talentpb.Application
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationClient.GetApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *applicationGRPCClient) UpdateApplication(ctx context.Context, req *talentpb.UpdateApplicationRequest, opts ...gax.CallOption) (*talentpb.Application, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "application.name", url.QueryEscape(req.GetApplication().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateApplication[0:len((*c.CallOptions).UpdateApplication):len((*c.CallOptions).UpdateApplication)], opts...)
	var resp *talentpb.Application
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationClient.UpdateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *applicationGRPCClient) DeleteApplication(ctx context.Context, req *talentpb.DeleteApplicationRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteApplication[0:len((*c.CallOptions).DeleteApplication):len((*c.CallOptions).DeleteApplication)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.applicationClient.DeleteApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *applicationGRPCClient) ListApplications(ctx context.Context, req *talentpb.ListApplicationsRequest, opts ...gax.CallOption) *ApplicationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListApplications[0:len((*c.CallOptions).ListApplications):len((*c.CallOptions).ListApplications)], opts...)
	it := &ApplicationIterator{}
	req = proto.Clone(req).(*talentpb.ListApplicationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Application, string, error) {
		resp := &talentpb.ListApplicationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.applicationClient.ListApplications(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetApplications(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ApplicationIterator manages a stream of *talentpb.Application.
type ApplicationIterator struct {
	items    []*talentpb.Application
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Application, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ApplicationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ApplicationIterator) Next() (*talentpb.Application, error) {
	var item *talentpb.Application
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ApplicationIterator) bufLen() int {
	return len(it.items)
}

func (it *ApplicationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
