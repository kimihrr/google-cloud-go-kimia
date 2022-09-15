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

// [START analyticsadmin_v1alpha_generated_AnalyticsAdminService_BatchDeleteUserLinks_sync]

package main

import (
	"context"

	admin "cloud.google.com/go/analytics/admin/apiv1alpha"
	adminpb "google.golang.org/genproto/googleapis/analytics/admin/v1alpha"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := admin.NewAnalyticsAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &adminpb.BatchDeleteUserLinksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/analytics/admin/v1alpha#BatchDeleteUserLinksRequest.
	}
	err = c.BatchDeleteUserLinks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END analyticsadmin_v1alpha_generated_AnalyticsAdminService_BatchDeleteUserLinks_sync]
