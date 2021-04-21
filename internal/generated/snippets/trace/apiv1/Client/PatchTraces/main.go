// Copyright 2021 Google LLC
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

// [START cloudtrace_generated_trace_apiv1_Client_PatchTraces]

package main

import (
	"context"

	trace "cloud.google.com/go/trace/apiv1"
	cloudtracepb "google.golang.org/genproto/googleapis/devtools/cloudtrace/v1"
)

func main() {
	ctx := context.Background()
	c, err := trace.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cloudtracepb.PatchTracesRequest{
		// TODO: Fill request struct fields.
	}
	err = c.PatchTraces(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END cloudtrace_generated_trace_apiv1_Client_PatchTraces]