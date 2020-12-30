package graphapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/services/graph/1.0/graph"
    "github.com/Azure/go-autorest/autorest"
)

        // GuestUserClientAPI contains the set of methods on the GuestUserClient type.
        type GuestUserClientAPI interface {
            Create(ctx context.Context, invitationContent graph.InvitationSended) (result graph.Invitation, err error)
            Delete(ctx context.Context, guestUserID string) (result autorest.Response, err error)
            Get(ctx context.Context, guestUserID string, selectParameter string) (result graph.GuestUser, err error)
            Update(ctx context.Context, guestUserProperties graph.GuestUser, guestUserID string) (result autorest.Response, err error)
        }

        var _ GuestUserClientAPI = (*graph.GuestUserClient)(nil)
