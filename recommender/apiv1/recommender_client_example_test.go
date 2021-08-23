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

package recommender_test

import (
	"context"

	recommender "cloud.google.com/go/recommender/apiv1"
	"google.golang.org/api/iterator"
	recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListInsights() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.ListInsightsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#ListInsightsRequest.
	}
	it := c.ListInsights(ctx, req)
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

func ExampleClient_GetInsight() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.GetInsightRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#GetInsightRequest.
	}
	resp, err := c.GetInsight(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkInsightAccepted() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.MarkInsightAcceptedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#MarkInsightAcceptedRequest.
	}
	resp, err := c.MarkInsightAccepted(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListRecommendations() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.ListRecommendationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#ListRecommendationsRequest.
	}
	it := c.ListRecommendations(ctx, req)
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

func ExampleClient_GetRecommendation() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.GetRecommendationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#GetRecommendationRequest.
	}
	resp, err := c.GetRecommendation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationClaimed() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.MarkRecommendationClaimedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#MarkRecommendationClaimedRequest.
	}
	resp, err := c.MarkRecommendationClaimed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationSucceeded() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.MarkRecommendationSucceededRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#MarkRecommendationSucceededRequest.
	}
	resp, err := c.MarkRecommendationSucceeded(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_MarkRecommendationFailed() {
	ctx := context.Background()
	c, err := recommender.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommenderpb.MarkRecommendationFailedRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommender/v1#MarkRecommendationFailedRequest.
	}
	resp, err := c.MarkRecommendationFailed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
