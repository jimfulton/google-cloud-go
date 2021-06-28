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

package translate

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newTranslationClientHook clientHook

// TranslationCallOptions contains the retry settings for each method of TranslationClient.
type TranslationCallOptions struct {
	TranslateText         []gax.CallOption
	DetectLanguage        []gax.CallOption
	GetSupportedLanguages []gax.CallOption
	BatchTranslateText    []gax.CallOption
	CreateGlossary        []gax.CallOption
	ListGlossaries        []gax.CallOption
	GetGlossary           []gax.CallOption
	DeleteGlossary        []gax.CallOption
}

func defaultTranslationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("translate.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("translate.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://translate.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTranslationCallOptions() *TranslationCallOptions {
	return &TranslationCallOptions{
		TranslateText:  []gax.CallOption{},
		DetectLanguage: []gax.CallOption{},
		GetSupportedLanguages: []gax.CallOption{
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
		BatchTranslateText: []gax.CallOption{},
		CreateGlossary:     []gax.CallOption{},
		ListGlossaries: []gax.CallOption{
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
		GetGlossary: []gax.CallOption{
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
		DeleteGlossary: []gax.CallOption{
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

// internalTranslationClient is an interface that defines the methods availaible from Cloud Translation API.
type internalTranslationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	TranslateText(context.Context, *translatepb.TranslateTextRequest, ...gax.CallOption) (*translatepb.TranslateTextResponse, error)
	DetectLanguage(context.Context, *translatepb.DetectLanguageRequest, ...gax.CallOption) (*translatepb.DetectLanguageResponse, error)
	GetSupportedLanguages(context.Context, *translatepb.GetSupportedLanguagesRequest, ...gax.CallOption) (*translatepb.SupportedLanguages, error)
	BatchTranslateText(context.Context, *translatepb.BatchTranslateTextRequest, ...gax.CallOption) (*BatchTranslateTextOperation, error)
	BatchTranslateTextOperation(name string) *BatchTranslateTextOperation
	CreateGlossary(context.Context, *translatepb.CreateGlossaryRequest, ...gax.CallOption) (*CreateGlossaryOperation, error)
	CreateGlossaryOperation(name string) *CreateGlossaryOperation
	ListGlossaries(context.Context, *translatepb.ListGlossariesRequest, ...gax.CallOption) *GlossaryIterator
	GetGlossary(context.Context, *translatepb.GetGlossaryRequest, ...gax.CallOption) (*translatepb.Glossary, error)
	DeleteGlossary(context.Context, *translatepb.DeleteGlossaryRequest, ...gax.CallOption) (*DeleteGlossaryOperation, error)
	DeleteGlossaryOperation(name string) *DeleteGlossaryOperation
}

// TranslationClient is a client for interacting with Cloud Translation API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides natural language translation operations.
type TranslationClient struct {
	// The internal transport-dependent client.
	internalClient internalTranslationClient

	// The call options for this service.
	CallOptions *TranslationCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TranslationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TranslationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TranslationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// TranslateText translates input text and returns translated text.
func (c *TranslationClient) TranslateText(ctx context.Context, req *translatepb.TranslateTextRequest, opts ...gax.CallOption) (*translatepb.TranslateTextResponse, error) {
	return c.internalClient.TranslateText(ctx, req, opts...)
}

// DetectLanguage detects the language of text within a request.
func (c *TranslationClient) DetectLanguage(ctx context.Context, req *translatepb.DetectLanguageRequest, opts ...gax.CallOption) (*translatepb.DetectLanguageResponse, error) {
	return c.internalClient.DetectLanguage(ctx, req, opts...)
}

// GetSupportedLanguages returns a list of supported languages for translation.
func (c *TranslationClient) GetSupportedLanguages(ctx context.Context, req *translatepb.GetSupportedLanguagesRequest, opts ...gax.CallOption) (*translatepb.SupportedLanguages, error) {
	return c.internalClient.GetSupportedLanguages(ctx, req, opts...)
}

// BatchTranslateText translates a large volume of text in asynchronous batch mode.
// This function provides real-time output as the inputs are being processed.
// If caller cancels a request, the partial results (for an input file, it’s
// all or nothing) may still be available on the specified output location.
//
// This call returns immediately and you can
// use google.longrunning.Operation.name (at http://google.longrunning.Operation.name) to poll the status of the call.
func (c *TranslationClient) BatchTranslateText(ctx context.Context, req *translatepb.BatchTranslateTextRequest, opts ...gax.CallOption) (*BatchTranslateTextOperation, error) {
	return c.internalClient.BatchTranslateText(ctx, req, opts...)
}

// BatchTranslateTextOperation returns a new BatchTranslateTextOperation from a given name.
// The name must be that of a previously created BatchTranslateTextOperation, possibly from a different process.
func (c *TranslationClient) BatchTranslateTextOperation(name string) *BatchTranslateTextOperation {
	return c.internalClient.BatchTranslateTextOperation(name)
}

// CreateGlossary creates a glossary and returns the long-running operation. Returns
// NOT_FOUND, if the project doesn’t exist.
func (c *TranslationClient) CreateGlossary(ctx context.Context, req *translatepb.CreateGlossaryRequest, opts ...gax.CallOption) (*CreateGlossaryOperation, error) {
	return c.internalClient.CreateGlossary(ctx, req, opts...)
}

// CreateGlossaryOperation returns a new CreateGlossaryOperation from a given name.
// The name must be that of a previously created CreateGlossaryOperation, possibly from a different process.
func (c *TranslationClient) CreateGlossaryOperation(name string) *CreateGlossaryOperation {
	return c.internalClient.CreateGlossaryOperation(name)
}

// ListGlossaries lists glossaries in a project. Returns NOT_FOUND, if the project doesn’t
// exist.
func (c *TranslationClient) ListGlossaries(ctx context.Context, req *translatepb.ListGlossariesRequest, opts ...gax.CallOption) *GlossaryIterator {
	return c.internalClient.ListGlossaries(ctx, req, opts...)
}

// GetGlossary gets a glossary. Returns NOT_FOUND, if the glossary doesn’t
// exist.
func (c *TranslationClient) GetGlossary(ctx context.Context, req *translatepb.GetGlossaryRequest, opts ...gax.CallOption) (*translatepb.Glossary, error) {
	return c.internalClient.GetGlossary(ctx, req, opts...)
}

// DeleteGlossary deletes a glossary, or cancels glossary construction
// if the glossary isn’t created yet.
// Returns NOT_FOUND, if the glossary doesn’t exist.
func (c *TranslationClient) DeleteGlossary(ctx context.Context, req *translatepb.DeleteGlossaryRequest, opts ...gax.CallOption) (*DeleteGlossaryOperation, error) {
	return c.internalClient.DeleteGlossary(ctx, req, opts...)
}

// DeleteGlossaryOperation returns a new DeleteGlossaryOperation from a given name.
// The name must be that of a previously created DeleteGlossaryOperation, possibly from a different process.
func (c *TranslationClient) DeleteGlossaryOperation(name string) *DeleteGlossaryOperation {
	return c.internalClient.DeleteGlossaryOperation(name)
}

// translationGRPCClient is a client for interacting with Cloud Translation API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type translationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TranslationClient
	CallOptions **TranslationCallOptions

	// The gRPC API client.
	translationClient translatepb.TranslationServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTranslationClient creates a new translation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides natural language translation operations.
func NewTranslationClient(ctx context.Context, opts ...option.ClientOption) (*TranslationClient, error) {
	clientOpts := defaultTranslationGRPCClientOptions()
	if newTranslationClientHook != nil {
		hookOpts, err := newTranslationClientHook(ctx, clientHookParams{})
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
	client := TranslationClient{CallOptions: defaultTranslationCallOptions()}

	c := &translationGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		translationClient: translatepb.NewTranslationServiceClient(connPool),
		CallOptions:       &client.CallOptions,
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
func (c *translationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *translationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *translationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *translationGRPCClient) TranslateText(ctx context.Context, req *translatepb.TranslateTextRequest, opts ...gax.CallOption) (*translatepb.TranslateTextResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TranslateText[0:len((*c.CallOptions).TranslateText):len((*c.CallOptions).TranslateText)], opts...)
	var resp *translatepb.TranslateTextResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.TranslateText(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *translationGRPCClient) DetectLanguage(ctx context.Context, req *translatepb.DetectLanguageRequest, opts ...gax.CallOption) (*translatepb.DetectLanguageResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DetectLanguage[0:len((*c.CallOptions).DetectLanguage):len((*c.CallOptions).DetectLanguage)], opts...)
	var resp *translatepb.DetectLanguageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.DetectLanguage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *translationGRPCClient) GetSupportedLanguages(ctx context.Context, req *translatepb.GetSupportedLanguagesRequest, opts ...gax.CallOption) (*translatepb.SupportedLanguages, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSupportedLanguages[0:len((*c.CallOptions).GetSupportedLanguages):len((*c.CallOptions).GetSupportedLanguages)], opts...)
	var resp *translatepb.SupportedLanguages
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.GetSupportedLanguages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *translationGRPCClient) BatchTranslateText(ctx context.Context, req *translatepb.BatchTranslateTextRequest, opts ...gax.CallOption) (*BatchTranslateTextOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchTranslateText[0:len((*c.CallOptions).BatchTranslateText):len((*c.CallOptions).BatchTranslateText)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.BatchTranslateText(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchTranslateTextOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *translationGRPCClient) CreateGlossary(ctx context.Context, req *translatepb.CreateGlossaryRequest, opts ...gax.CallOption) (*CreateGlossaryOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateGlossary[0:len((*c.CallOptions).CreateGlossary):len((*c.CallOptions).CreateGlossary)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.CreateGlossary(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateGlossaryOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *translationGRPCClient) ListGlossaries(ctx context.Context, req *translatepb.ListGlossariesRequest, opts ...gax.CallOption) *GlossaryIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListGlossaries[0:len((*c.CallOptions).ListGlossaries):len((*c.CallOptions).ListGlossaries)], opts...)
	it := &GlossaryIterator{}
	req = proto.Clone(req).(*translatepb.ListGlossariesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*translatepb.Glossary, string, error) {
		var resp *translatepb.ListGlossariesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.translationClient.ListGlossaries(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetGlossaries(), resp.GetNextPageToken(), nil
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

func (c *translationGRPCClient) GetGlossary(ctx context.Context, req *translatepb.GetGlossaryRequest, opts ...gax.CallOption) (*translatepb.Glossary, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGlossary[0:len((*c.CallOptions).GetGlossary):len((*c.CallOptions).GetGlossary)], opts...)
	var resp *translatepb.Glossary
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.GetGlossary(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *translationGRPCClient) DeleteGlossary(ctx context.Context, req *translatepb.DeleteGlossaryRequest, opts ...gax.CallOption) (*DeleteGlossaryOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteGlossary[0:len((*c.CallOptions).DeleteGlossary):len((*c.CallOptions).DeleteGlossary)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.translationClient.DeleteGlossary(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteGlossaryOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// BatchTranslateTextOperation manages a long-running operation from BatchTranslateText.
type BatchTranslateTextOperation struct {
	lro *longrunning.Operation
}

// BatchTranslateTextOperation returns a new BatchTranslateTextOperation from a given name.
// The name must be that of a previously created BatchTranslateTextOperation, possibly from a different process.
func (c *translationGRPCClient) BatchTranslateTextOperation(name string) *BatchTranslateTextOperation {
	return &BatchTranslateTextOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchTranslateTextOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*translatepb.BatchTranslateResponse, error) {
	var resp translatepb.BatchTranslateResponse
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
func (op *BatchTranslateTextOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*translatepb.BatchTranslateResponse, error) {
	var resp translatepb.BatchTranslateResponse
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
func (op *BatchTranslateTextOperation) Metadata() (*translatepb.BatchTranslateMetadata, error) {
	var meta translatepb.BatchTranslateMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchTranslateTextOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchTranslateTextOperation) Name() string {
	return op.lro.Name()
}

// CreateGlossaryOperation manages a long-running operation from CreateGlossary.
type CreateGlossaryOperation struct {
	lro *longrunning.Operation
}

// CreateGlossaryOperation returns a new CreateGlossaryOperation from a given name.
// The name must be that of a previously created CreateGlossaryOperation, possibly from a different process.
func (c *translationGRPCClient) CreateGlossaryOperation(name string) *CreateGlossaryOperation {
	return &CreateGlossaryOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateGlossaryOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*translatepb.Glossary, error) {
	var resp translatepb.Glossary
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
func (op *CreateGlossaryOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*translatepb.Glossary, error) {
	var resp translatepb.Glossary
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
func (op *CreateGlossaryOperation) Metadata() (*translatepb.CreateGlossaryMetadata, error) {
	var meta translatepb.CreateGlossaryMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateGlossaryOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateGlossaryOperation) Name() string {
	return op.lro.Name()
}

// DeleteGlossaryOperation manages a long-running operation from DeleteGlossary.
type DeleteGlossaryOperation struct {
	lro *longrunning.Operation
}

// DeleteGlossaryOperation returns a new DeleteGlossaryOperation from a given name.
// The name must be that of a previously created DeleteGlossaryOperation, possibly from a different process.
func (c *translationGRPCClient) DeleteGlossaryOperation(name string) *DeleteGlossaryOperation {
	return &DeleteGlossaryOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteGlossaryOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*translatepb.DeleteGlossaryResponse, error) {
	var resp translatepb.DeleteGlossaryResponse
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
func (op *DeleteGlossaryOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*translatepb.DeleteGlossaryResponse, error) {
	var resp translatepb.DeleteGlossaryResponse
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
func (op *DeleteGlossaryOperation) Metadata() (*translatepb.DeleteGlossaryMetadata, error) {
	var meta translatepb.DeleteGlossaryMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteGlossaryOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteGlossaryOperation) Name() string {
	return op.lro.Name()
}

// GlossaryIterator manages a stream of *translatepb.Glossary.
type GlossaryIterator struct {
	items    []*translatepb.Glossary
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
	InternalFetch func(pageSize int, pageToken string) (results []*translatepb.Glossary, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GlossaryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GlossaryIterator) Next() (*translatepb.Glossary, error) {
	var item *translatepb.Glossary
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GlossaryIterator) bufLen() int {
	return len(it.items)
}

func (it *GlossaryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
