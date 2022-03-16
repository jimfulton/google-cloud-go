// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START networkconnectivity_v1alpha1_generated_HubService_ListHubs_sync]

package main

import (
	"context"

	networkconnectivity "cloud.google.com/go/networkconnectivity/apiv1alpha1"
	"google.golang.org/api/iterator"
	networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"
)

func main() {
	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &networkconnectivitypb.ListHubsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1#ListHubsRequest.
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

// [END networkconnectivity_v1alpha1_generated_HubService_ListHubs_sync]
