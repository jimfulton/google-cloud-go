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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	structpb "github.com/golang/protobuf/ptypes/struct"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newAgentsClientHook clientHook

// AgentsCallOptions contains the retry settings for each method of AgentsClient.
type AgentsCallOptions struct {
	ListAgents               []gax.CallOption
	GetAgent                 []gax.CallOption
	CreateAgent              []gax.CallOption
	UpdateAgent              []gax.CallOption
	DeleteAgent              []gax.CallOption
	ExportAgent              []gax.CallOption
	RestoreAgent             []gax.CallOption
	ValidateAgent            []gax.CallOption
	GetAgentValidationResult []gax.CallOption
}

func defaultAgentsGRPCClientOptions() []option.ClientOption {
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

func defaultAgentsCallOptions() *AgentsCallOptions {
	return &AgentsCallOptions{
		ListAgents: []gax.CallOption{
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
		GetAgent: []gax.CallOption{
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
		CreateAgent: []gax.CallOption{
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
		UpdateAgent: []gax.CallOption{
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
		DeleteAgent: []gax.CallOption{
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
		ExportAgent: []gax.CallOption{
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
		RestoreAgent: []gax.CallOption{
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
		ValidateAgent: []gax.CallOption{
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
		GetAgentValidationResult: []gax.CallOption{
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

// internalAgentsClient is an interface that defines the methods availaible from Dialogflow API.
type internalAgentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAgents(context.Context, *cxpb.ListAgentsRequest, ...gax.CallOption) *AgentIterator
	GetAgent(context.Context, *cxpb.GetAgentRequest, ...gax.CallOption) (*cxpb.Agent, error)
	CreateAgent(context.Context, *cxpb.CreateAgentRequest, ...gax.CallOption) (*cxpb.Agent, error)
	UpdateAgent(context.Context, *cxpb.UpdateAgentRequest, ...gax.CallOption) (*cxpb.Agent, error)
	DeleteAgent(context.Context, *cxpb.DeleteAgentRequest, ...gax.CallOption) error
	ExportAgent(context.Context, *cxpb.ExportAgentRequest, ...gax.CallOption) (*ExportAgentOperation, error)
	ExportAgentOperation(name string) *ExportAgentOperation
	RestoreAgent(context.Context, *cxpb.RestoreAgentRequest, ...gax.CallOption) (*RestoreAgentOperation, error)
	RestoreAgentOperation(name string) *RestoreAgentOperation
	ValidateAgent(context.Context, *cxpb.ValidateAgentRequest, ...gax.CallOption) (*cxpb.AgentValidationResult, error)
	GetAgentValidationResult(context.Context, *cxpb.GetAgentValidationResultRequest, ...gax.CallOption) (*cxpb.AgentValidationResult, error)
}

// AgentsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Agents.
type AgentsClient struct {
	// The internal transport-dependent client.
	internalClient internalAgentsClient

	// The call options for this service.
	CallOptions *AgentsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AgentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AgentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AgentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAgents returns the list of all agents in the specified location.
func (c *AgentsClient) ListAgents(ctx context.Context, req *cxpb.ListAgentsRequest, opts ...gax.CallOption) *AgentIterator {
	return c.internalClient.ListAgents(ctx, req, opts...)
}

// GetAgent retrieves the specified agent.
func (c *AgentsClient) GetAgent(ctx context.Context, req *cxpb.GetAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	return c.internalClient.GetAgent(ctx, req, opts...)
}

// CreateAgent creates an agent in the specified location.
//
// Note: You should always train flows prior to sending them queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *AgentsClient) CreateAgent(ctx context.Context, req *cxpb.CreateAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	return c.internalClient.CreateAgent(ctx, req, opts...)
}

// UpdateAgent updates the specified agent.
//
// Note: You should always train flows prior to sending them queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *AgentsClient) UpdateAgent(ctx context.Context, req *cxpb.UpdateAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	return c.internalClient.UpdateAgent(ctx, req, opts...)
}

// DeleteAgent deletes the specified agent.
func (c *AgentsClient) DeleteAgent(ctx context.Context, req *cxpb.DeleteAgentRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAgent(ctx, req, opts...)
}

// ExportAgent exports the specified agent to a binary file.
//
// This method is a long-running
// operation (at https://cloud.google.com/dialogflow/cx/docs/how/long-running-operation).
// The returned Operation type has the following method-specific fields:
//
//   metadata: An empty Struct
//   message (at https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#struct)
//
//   response: ExportAgentResponse
func (c *AgentsClient) ExportAgent(ctx context.Context, req *cxpb.ExportAgentRequest, opts ...gax.CallOption) (*ExportAgentOperation, error) {
	return c.internalClient.ExportAgent(ctx, req, opts...)
}

// ExportAgentOperation returns a new ExportAgentOperation from a given name.
// The name must be that of a previously created ExportAgentOperation, possibly from a different process.
func (c *AgentsClient) ExportAgentOperation(name string) *ExportAgentOperation {
	return c.internalClient.ExportAgentOperation(name)
}

// RestoreAgent restores the specified agent from a binary file.
//
// Replaces the current agent with a new one. Note that all existing resources
// in agent (e.g. intents, entity types, flows) will be removed.
//
// This method is a long-running
// operation (at https://cloud.google.com/dialogflow/cx/docs/how/long-running-operation).
// The returned Operation type has the following method-specific fields:
//
//   metadata: An empty Struct
//   message (at https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#struct)
//
//   response: An Empty
//   message (at https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#empty)
//
// Note: You should always train flows prior to sending them queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *AgentsClient) RestoreAgent(ctx context.Context, req *cxpb.RestoreAgentRequest, opts ...gax.CallOption) (*RestoreAgentOperation, error) {
	return c.internalClient.RestoreAgent(ctx, req, opts...)
}

// RestoreAgentOperation returns a new RestoreAgentOperation from a given name.
// The name must be that of a previously created RestoreAgentOperation, possibly from a different process.
func (c *AgentsClient) RestoreAgentOperation(name string) *RestoreAgentOperation {
	return c.internalClient.RestoreAgentOperation(name)
}

// ValidateAgent validates the specified agent and creates or updates validation results.
// The agent in draft version is validated. Please call this API after the
// training is completed to get the complete validation results.
func (c *AgentsClient) ValidateAgent(ctx context.Context, req *cxpb.ValidateAgentRequest, opts ...gax.CallOption) (*cxpb.AgentValidationResult, error) {
	return c.internalClient.ValidateAgent(ctx, req, opts...)
}

// GetAgentValidationResult gets the latest agent validation result. Agent validation is performed
// when ValidateAgent is called.
func (c *AgentsClient) GetAgentValidationResult(ctx context.Context, req *cxpb.GetAgentValidationResultRequest, opts ...gax.CallOption) (*cxpb.AgentValidationResult, error) {
	return c.internalClient.GetAgentValidationResult(ctx, req, opts...)
}

// agentsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type agentsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AgentsClient
	CallOptions **AgentsCallOptions

	// The gRPC API client.
	agentsClient cxpb.AgentsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAgentsClient creates a new agents client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Agents.
func NewAgentsClient(ctx context.Context, opts ...option.ClientOption) (*AgentsClient, error) {
	clientOpts := defaultAgentsGRPCClientOptions()
	if newAgentsClientHook != nil {
		hookOpts, err := newAgentsClientHook(ctx, clientHookParams{})
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
	client := AgentsClient{CallOptions: defaultAgentsCallOptions()}

	c := &agentsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		agentsClient:     cxpb.NewAgentsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *agentsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *agentsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *agentsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *agentsGRPCClient) ListAgents(ctx context.Context, req *cxpb.ListAgentsRequest, opts ...gax.CallOption) *AgentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAgents[0:len((*c.CallOptions).ListAgents):len((*c.CallOptions).ListAgents)], opts...)
	it := &AgentIterator{}
	req = proto.Clone(req).(*cxpb.ListAgentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Agent, string, error) {
		resp := &cxpb.ListAgentsResponse{}
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
			resp, err = c.agentsClient.ListAgents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAgents(), resp.GetNextPageToken(), nil
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

func (c *agentsGRPCClient) GetAgent(ctx context.Context, req *cxpb.GetAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAgent[0:len((*c.CallOptions).GetAgent):len((*c.CallOptions).GetAgent)], opts...)
	var resp *cxpb.Agent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.GetAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *agentsGRPCClient) CreateAgent(ctx context.Context, req *cxpb.CreateAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAgent[0:len((*c.CallOptions).CreateAgent):len((*c.CallOptions).CreateAgent)], opts...)
	var resp *cxpb.Agent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.CreateAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *agentsGRPCClient) UpdateAgent(ctx context.Context, req *cxpb.UpdateAgentRequest, opts ...gax.CallOption) (*cxpb.Agent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "agent.name", url.QueryEscape(req.GetAgent().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAgent[0:len((*c.CallOptions).UpdateAgent):len((*c.CallOptions).UpdateAgent)], opts...)
	var resp *cxpb.Agent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.UpdateAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *agentsGRPCClient) DeleteAgent(ctx context.Context, req *cxpb.DeleteAgentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAgent[0:len((*c.CallOptions).DeleteAgent):len((*c.CallOptions).DeleteAgent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.agentsClient.DeleteAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *agentsGRPCClient) ExportAgent(ctx context.Context, req *cxpb.ExportAgentRequest, opts ...gax.CallOption) (*ExportAgentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExportAgent[0:len((*c.CallOptions).ExportAgent):len((*c.CallOptions).ExportAgent)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.ExportAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportAgentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *agentsGRPCClient) RestoreAgent(ctx context.Context, req *cxpb.RestoreAgentRequest, opts ...gax.CallOption) (*RestoreAgentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RestoreAgent[0:len((*c.CallOptions).RestoreAgent):len((*c.CallOptions).RestoreAgent)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.RestoreAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RestoreAgentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *agentsGRPCClient) ValidateAgent(ctx context.Context, req *cxpb.ValidateAgentRequest, opts ...gax.CallOption) (*cxpb.AgentValidationResult, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ValidateAgent[0:len((*c.CallOptions).ValidateAgent):len((*c.CallOptions).ValidateAgent)], opts...)
	var resp *cxpb.AgentValidationResult
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.ValidateAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *agentsGRPCClient) GetAgentValidationResult(ctx context.Context, req *cxpb.GetAgentValidationResultRequest, opts ...gax.CallOption) (*cxpb.AgentValidationResult, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAgentValidationResult[0:len((*c.CallOptions).GetAgentValidationResult):len((*c.CallOptions).GetAgentValidationResult)], opts...)
	var resp *cxpb.AgentValidationResult
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.agentsClient.GetAgentValidationResult(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExportAgentOperation manages a long-running operation from ExportAgent.
type ExportAgentOperation struct {
	lro *longrunning.Operation
}

// ExportAgentOperation returns a new ExportAgentOperation from a given name.
// The name must be that of a previously created ExportAgentOperation, possibly from a different process.
func (c *agentsGRPCClient) ExportAgentOperation(name string) *ExportAgentOperation {
	return &ExportAgentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportAgentOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.ExportAgentResponse, error) {
	var resp cxpb.ExportAgentResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportAgentOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.ExportAgentResponse, error) {
	var resp cxpb.ExportAgentResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportAgentOperation) Metadata() (*structpb.Struct, error) {
	var meta structpb.Struct
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportAgentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportAgentOperation) Name() string {
	return op.lro.Name()
}

// RestoreAgentOperation manages a long-running operation from RestoreAgent.
type RestoreAgentOperation struct {
	lro *longrunning.Operation
}

// RestoreAgentOperation returns a new RestoreAgentOperation from a given name.
// The name must be that of a previously created RestoreAgentOperation, possibly from a different process.
func (c *agentsGRPCClient) RestoreAgentOperation(name string) *RestoreAgentOperation {
	return &RestoreAgentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RestoreAgentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RestoreAgentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RestoreAgentOperation) Metadata() (*structpb.Struct, error) {
	var meta structpb.Struct
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RestoreAgentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RestoreAgentOperation) Name() string {
	return op.lro.Name()
}

// AgentIterator manages a stream of *cxpb.Agent.
type AgentIterator struct {
	items    []*cxpb.Agent
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Agent, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AgentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AgentIterator) Next() (*cxpb.Agent, error) {
	var item *cxpb.Agent
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AgentIterator) bufLen() int {
	return len(it.items)
}

func (it *AgentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
