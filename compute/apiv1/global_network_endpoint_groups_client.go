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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newGlobalNetworkEndpointGroupsClientHook clientHook

// GlobalNetworkEndpointGroupsCallOptions contains the retry settings for each method of GlobalNetworkEndpointGroupsClient.
type GlobalNetworkEndpointGroupsCallOptions struct {
	AttachNetworkEndpoints []gax.CallOption
	Delete                 []gax.CallOption
	DetachNetworkEndpoints []gax.CallOption
	Get                    []gax.CallOption
	Insert                 []gax.CallOption
	List                   []gax.CallOption
	ListNetworkEndpoints   []gax.CallOption
}

// internalGlobalNetworkEndpointGroupsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalGlobalNetworkEndpointGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AttachNetworkEndpoints(context.Context, *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.Operation, error)
	DetachNetworkEndpoints(context.Context, *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.NetworkEndpointGroup, error)
	Insert(context.Context, *computepb.InsertGlobalNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListGlobalNetworkEndpointGroupsRequest, ...gax.CallOption) (*computepb.NetworkEndpointGroupList, error)
	ListNetworkEndpoints(context.Context, *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, ...gax.CallOption) (*computepb.NetworkEndpointGroupsListNetworkEndpoints, error)
}

// GlobalNetworkEndpointGroupsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalNetworkEndpointGroups API.
type GlobalNetworkEndpointGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalNetworkEndpointGroupsClient

	// The call options for this service.
	CallOptions *GlobalNetworkEndpointGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalNetworkEndpointGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalNetworkEndpointGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GlobalNetworkEndpointGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AttachNetworkEndpoints attach a network endpoint to the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AttachNetworkEndpoints(ctx, req, opts...)
}

// Delete deletes the specified network endpoint group.Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *GlobalNetworkEndpointGroupsClient) Delete(ctx context.Context, req *computepb.DeleteGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DetachNetworkEndpoints detach the network endpoint from the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.DetachNetworkEndpoints(ctx, req, opts...)
}

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *GlobalNetworkEndpointGroupsClient) Get(ctx context.Context, req *computepb.GetGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *GlobalNetworkEndpointGroupsClient) Insert(ctx context.Context, req *computepb.InsertGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of network endpoint groups that are located in the specified project.
func (c *GlobalNetworkEndpointGroupsClient) List(ctx context.Context, req *computepb.ListGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroupList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *GlobalNetworkEndpointGroupsClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroupsListNetworkEndpoints, error) {
	return c.internalClient.ListNetworkEndpoints(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalNetworkEndpointGroupsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGlobalNetworkEndpointGroupsRESTClient creates a new global network endpoint groups rest client.
//
// The GlobalNetworkEndpointGroups API.
func NewGlobalNetworkEndpointGroupsRESTClient(ctx context.Context, opts ...option.ClientOption) (*GlobalNetworkEndpointGroupsClient, error) {
	clientOpts := append(defaultGlobalNetworkEndpointGroupsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &globalNetworkEndpointGroupsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &GlobalNetworkEndpointGroupsClient{internalClient: c, CallOptions: &GlobalNetworkEndpointGroupsCallOptions{}}, nil
}

func defaultGlobalNetworkEndpointGroupsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *globalNetworkEndpointGroupsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalNetworkEndpointGroupsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *globalNetworkEndpointGroupsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AttachNetworkEndpoints attach a network endpoint to the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetGlobalNetworkEndpointGroupsAttachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/attachNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified network endpoint group.Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *globalNetworkEndpointGroupsRESTClient) Delete(ctx context.Context, req *computepb.DeleteGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// DetachNetworkEndpoints detach the network endpoint from the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetGlobalNetworkEndpointGroupsDetachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/detachNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *globalNetworkEndpointGroupsRESTClient) Get(ctx context.Context, req *computepb.GetGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v", req.GetProject(), req.GetNetworkEndpointGroup())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.NetworkEndpointGroup{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *globalNetworkEndpointGroupsRESTClient) Insert(ctx context.Context, req *computepb.InsertGlobalNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNetworkEndpointGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List retrieves the list of network endpoint groups that are located in the specified project.
func (c *globalNetworkEndpointGroupsRESTClient) List(ctx context.Context, req *computepb.ListGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroupList, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.NetworkEndpointGroupList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *globalNetworkEndpointGroupsRESTClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsGlobalNetworkEndpointGroupsRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroupsListNetworkEndpoints, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/networkEndpointGroups/%v/listNetworkEndpoints", req.GetProject(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.NetworkEndpointGroupsListNetworkEndpoints{}

	return rsp, unm.Unmarshal(buf, rsp)
}
