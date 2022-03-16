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

package networkconnectivity_test

import (
	"context"

	networkconnectivity "cloud.google.com/go/networkconnectivity/apiv1"
	"google.golang.org/api/iterator"
	networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1"
)

func ExampleNewHubClient() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleHubClient_ListHubs() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.ListHubsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#ListHubsRequest.
	}
	it := c.ListHubs(ctx, req)
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

func ExampleHubClient_GetHub() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.GetHubRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#GetHubRequest.
	}
	resp, err := c.GetHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_CreateHub() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.CreateHubRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#CreateHubRequest.
	}
	op, err := c.CreateHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_UpdateHub() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.UpdateHubRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#UpdateHubRequest.
	}
	op, err := c.UpdateHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_DeleteHub() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.DeleteHubRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#DeleteHubRequest.
	}
	op, err := c.DeleteHub(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleHubClient_ListSpokes() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.ListSpokesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#ListSpokesRequest.
	}
	it := c.ListSpokes(ctx, req)
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

func ExampleHubClient_GetSpoke() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.GetSpokeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#GetSpokeRequest.
	}
	resp, err := c.GetSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_CreateSpoke() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.CreateSpokeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#CreateSpokeRequest.
	}
	op, err := c.CreateSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_UpdateSpoke() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.UpdateSpokeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#UpdateSpokeRequest.
	}
	op, err := c.UpdateSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHubClient_DeleteSpoke() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.DeleteSpokeRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1#DeleteSpokeRequest.
	}
	op, err := c.DeleteSpoke(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
