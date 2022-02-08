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

package cx

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
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newWebhooksClientHook clientHook

// WebhooksCallOptions contains the retry settings for each method of WebhooksClient.
type WebhooksCallOptions struct {
	ListWebhooks  []gax.CallOption
	GetWebhook    []gax.CallOption
	CreateWebhook []gax.CallOption
	UpdateWebhook []gax.CallOption
	DeleteWebhook []gax.CallOption
}

func defaultWebhooksGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultWebhooksCallOptions() *WebhooksCallOptions {
	return &WebhooksCallOptions{
		ListWebhooks: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetWebhook: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateWebhook: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateWebhook: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteWebhook: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
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

// internalWebhooksClient is an interface that defines the methods availaible from Dialogflow API.
type internalWebhooksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListWebhooks(context.Context, *cxpb.ListWebhooksRequest, ...gax.CallOption) *WebhookIterator
	GetWebhook(context.Context, *cxpb.GetWebhookRequest, ...gax.CallOption) (*cxpb.Webhook, error)
	CreateWebhook(context.Context, *cxpb.CreateWebhookRequest, ...gax.CallOption) (*cxpb.Webhook, error)
	UpdateWebhook(context.Context, *cxpb.UpdateWebhookRequest, ...gax.CallOption) (*cxpb.Webhook, error)
	DeleteWebhook(context.Context, *cxpb.DeleteWebhookRequest, ...gax.CallOption) error
}

// WebhooksClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Webhooks.
type WebhooksClient struct {
	// The internal transport-dependent client.
	internalClient internalWebhooksClient

	// The call options for this service.
	CallOptions *WebhooksCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *WebhooksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *WebhooksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *WebhooksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListWebhooks returns the list of all webhooks in the specified agent.
func (c *WebhooksClient) ListWebhooks(ctx context.Context, req *cxpb.ListWebhooksRequest, opts ...gax.CallOption) *WebhookIterator {
	return c.internalClient.ListWebhooks(ctx, req, opts...)
}

// GetWebhook retrieves the specified webhook.
func (c *WebhooksClient) GetWebhook(ctx context.Context, req *cxpb.GetWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	return c.internalClient.GetWebhook(ctx, req, opts...)
}

// CreateWebhook creates a webhook in the specified agent.
func (c *WebhooksClient) CreateWebhook(ctx context.Context, req *cxpb.CreateWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	return c.internalClient.CreateWebhook(ctx, req, opts...)
}

// UpdateWebhook updates the specified webhook.
func (c *WebhooksClient) UpdateWebhook(ctx context.Context, req *cxpb.UpdateWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	return c.internalClient.UpdateWebhook(ctx, req, opts...)
}

// DeleteWebhook deletes the specified webhook.
func (c *WebhooksClient) DeleteWebhook(ctx context.Context, req *cxpb.DeleteWebhookRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteWebhook(ctx, req, opts...)
}

// webhooksGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type webhooksGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing WebhooksClient
	CallOptions **WebhooksCallOptions

	// The gRPC API client.
	webhooksClient cxpb.WebhooksClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewWebhooksClient creates a new webhooks client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Webhooks.
func NewWebhooksClient(ctx context.Context, opts ...option.ClientOption) (*WebhooksClient, error) {
	clientOpts := defaultWebhooksGRPCClientOptions()
	if newWebhooksClientHook != nil {
		hookOpts, err := newWebhooksClientHook(ctx, clientHookParams{})
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
	client := WebhooksClient{CallOptions: defaultWebhooksCallOptions()}

	c := &webhooksGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		webhooksClient:   cxpb.NewWebhooksClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *webhooksGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *webhooksGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *webhooksGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *webhooksGRPCClient) ListWebhooks(ctx context.Context, req *cxpb.ListWebhooksRequest, opts ...gax.CallOption) *WebhookIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListWebhooks[0:len((*c.CallOptions).ListWebhooks):len((*c.CallOptions).ListWebhooks)], opts...)
	it := &WebhookIterator{}
	req = proto.Clone(req).(*cxpb.ListWebhooksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Webhook, string, error) {
		resp := &cxpb.ListWebhooksResponse{}
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
			resp, err = c.webhooksClient.ListWebhooks(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetWebhooks(), resp.GetNextPageToken(), nil
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

func (c *webhooksGRPCClient) GetWebhook(ctx context.Context, req *cxpb.GetWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetWebhook[0:len((*c.CallOptions).GetWebhook):len((*c.CallOptions).GetWebhook)], opts...)
	var resp *cxpb.Webhook
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webhooksClient.GetWebhook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *webhooksGRPCClient) CreateWebhook(ctx context.Context, req *cxpb.CreateWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateWebhook[0:len((*c.CallOptions).CreateWebhook):len((*c.CallOptions).CreateWebhook)], opts...)
	var resp *cxpb.Webhook
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webhooksClient.CreateWebhook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *webhooksGRPCClient) UpdateWebhook(ctx context.Context, req *cxpb.UpdateWebhookRequest, opts ...gax.CallOption) (*cxpb.Webhook, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "webhook.name", url.QueryEscape(req.GetWebhook().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateWebhook[0:len((*c.CallOptions).UpdateWebhook):len((*c.CallOptions).UpdateWebhook)], opts...)
	var resp *cxpb.Webhook
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webhooksClient.UpdateWebhook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *webhooksGRPCClient) DeleteWebhook(ctx context.Context, req *cxpb.DeleteWebhookRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteWebhook[0:len((*c.CallOptions).DeleteWebhook):len((*c.CallOptions).DeleteWebhook)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.webhooksClient.DeleteWebhook(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// WebhookIterator manages a stream of *cxpb.Webhook.
type WebhookIterator struct {
	items    []*cxpb.Webhook
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Webhook, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WebhookIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WebhookIterator) Next() (*cxpb.Webhook, error) {
	var item *cxpb.Webhook
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WebhookIterator) bufLen() int {
	return len(it.items)
}

func (it *WebhookIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
