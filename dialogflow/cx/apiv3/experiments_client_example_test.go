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

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
)

func ExampleNewExperimentsClient() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleExperimentsClient_ListExperiments() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListExperimentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#ListExperimentsRequest.
	}
	it := c.ListExperiments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleExperimentsClient_GetExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.GetExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#GetExperimentRequest.
	}
	resp, err := c.GetExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleExperimentsClient_CreateExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.CreateExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#CreateExperimentRequest.
	}
	resp, err := c.CreateExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleExperimentsClient_UpdateExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.UpdateExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#UpdateExperimentRequest.
	}
	resp, err := c.UpdateExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleExperimentsClient_DeleteExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.DeleteExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#DeleteExperimentRequest.
	}
	err = c.DeleteExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleExperimentsClient_StartExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.StartExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#StartExperimentRequest.
	}
	resp, err := c.StartExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleExperimentsClient_StopExperiment() {
	ctx := context.Background()
	c, err := cx.NewExperimentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.StopExperimentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3#StopExperimentRequest.
	}
	resp, err := c.StopExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
