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

package retail_test

import (
	"context"

	retail "cloud.google.com/go/retail/apiv2"
	"google.golang.org/api/iterator"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
)

func ExampleNewCatalogClient() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCatalogClient_ListCatalogs() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.ListCatalogsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#ListCatalogsRequest.
	}
	it := c.ListCatalogs(ctx, req)
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

func ExampleCatalogClient_UpdateCatalog() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.UpdateCatalogRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#UpdateCatalogRequest.
	}
	resp, err := c.UpdateCatalog(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCatalogClient_SetDefaultBranch() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.SetDefaultBranchRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#SetDefaultBranchRequest.
	}
	err = c.SetDefaultBranch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCatalogClient_GetDefaultBranch() {
	ctx := context.Background()
	c, err := retail.NewCatalogClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &retailpb.GetDefaultBranchRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/retail/v2#GetDefaultBranchRequest.
	}
	resp, err := c.GetDefaultBranch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
