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

package servicecontrol_test

import (
	"context"

	servicecontrol "cloud.google.com/go/servicecontrol/apiv1"
	servicecontrolpb "google.golang.org/genproto/googleapis/api/servicecontrol/v1"
)

func ExampleNewServiceControllerClient() {
	ctx := context.Background()
	c, err := servicecontrol.NewServiceControllerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleServiceControllerClient_Check() {
	ctx := context.Background()
	c, err := servicecontrol.NewServiceControllerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicecontrolpb.CheckRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/api/servicecontrol/v1#CheckRequest.
	}
	resp, err := c.Check(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceControllerClient_Report() {
	ctx := context.Background()
	c, err := servicecontrol.NewServiceControllerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicecontrolpb.ReportRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/api/servicecontrol/v1#ReportRequest.
	}
	resp, err := c.Report(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
