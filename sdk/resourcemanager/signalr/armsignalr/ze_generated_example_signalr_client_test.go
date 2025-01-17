//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsignalr_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_CheckNameAvailability.json
func ExampleSignalRClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	_, err = client.CheckNameAvailability(ctx,
		"<location>",
		armsignalr.NameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_ListBySubscription.json
func ExampleSignalRClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("SignalRResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_ListByResourceGroup.json
func ExampleSignalRClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("SignalRResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Get.json
func ExampleSignalRClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SignalRResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_CreateOrUpdate.json
func ExampleSignalRClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.SignalRResource{
			TrackedResource: armsignalr.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
			Identity: &armsignalr.ManagedIdentity{
				Type: armsignalr.ManagedIdentityTypeSystemAssigned.ToPtr(),
			},
			Kind: armsignalr.ServiceKindSignalR.ToPtr(),
			Properties: &armsignalr.SignalRProperties{
				Cors: &armsignalr.SignalRCorsSettings{
					AllowedOrigins: []*string{
						to.StringPtr("https://foo.com"),
						to.StringPtr("https://bar.com")},
				},
				DisableAADAuth:   to.BoolPtr(false),
				DisableLocalAuth: to.BoolPtr(false),
				Features: []*armsignalr.SignalRFeature{
					{
						Flag:       armsignalr.FeatureFlagsServiceMode.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableConnectivityLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableMessagingLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableLiveTrace.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					}},
				NetworkACLs: &armsignalr.SignalRNetworkACLs{
					DefaultAction: armsignalr.ACLActionDeny.ToPtr(),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							NetworkACL: armsignalr.NetworkACL{
								Allow: []*armsignalr.SignalRRequestType{
									armsignalr.SignalRRequestTypeServerConnection.ToPtr()},
							},
							Name: to.StringPtr("<name>"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							armsignalr.SignalRRequestTypeClientConnection.ToPtr()},
					},
				},
				PublicNetworkAccess: to.StringPtr("<public-network-access>"),
				TLS: &armsignalr.SignalRTLSSettings{
					ClientCertEnabled: to.BoolPtr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: armsignalr.UpstreamAuthTypeManagedIdentity.ToPtr(),
								ManagedIdentity: &armsignalr.ManagedIdentitySettings{
									Resource: to.StringPtr("<resource>"),
								},
							},
							CategoryPattern: to.StringPtr("<category-pattern>"),
							EventPattern:    to.StringPtr("<event-pattern>"),
							HubPattern:      to.StringPtr("<hub-pattern>"),
							URLTemplate:     to.StringPtr("<urltemplate>"),
						}},
				},
			},
			SKU: &armsignalr.ResourceSKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(1),
				Tier:     armsignalr.SignalRSKUTierStandard.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SignalRResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Delete.json
func ExampleSignalRClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Update.json
func ExampleSignalRClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.SignalRResource{
			TrackedResource: armsignalr.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key1": to.StringPtr("value1"),
				},
			},
			Identity: &armsignalr.ManagedIdentity{
				Type: armsignalr.ManagedIdentityTypeSystemAssigned.ToPtr(),
			},
			Kind: armsignalr.ServiceKindSignalR.ToPtr(),
			Properties: &armsignalr.SignalRProperties{
				Cors: &armsignalr.SignalRCorsSettings{
					AllowedOrigins: []*string{
						to.StringPtr("https://foo.com"),
						to.StringPtr("https://bar.com")},
				},
				DisableAADAuth:   to.BoolPtr(false),
				DisableLocalAuth: to.BoolPtr(false),
				Features: []*armsignalr.SignalRFeature{
					{
						Flag:       armsignalr.FeatureFlagsServiceMode.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableConnectivityLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableMessagingLogs.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					},
					{
						Flag:       armsignalr.FeatureFlagsEnableLiveTrace.ToPtr(),
						Properties: map[string]*string{},
						Value:      to.StringPtr("<value>"),
					}},
				NetworkACLs: &armsignalr.SignalRNetworkACLs{
					DefaultAction: armsignalr.ACLActionDeny.ToPtr(),
					PrivateEndpoints: []*armsignalr.PrivateEndpointACL{
						{
							NetworkACL: armsignalr.NetworkACL{
								Allow: []*armsignalr.SignalRRequestType{
									armsignalr.SignalRRequestTypeServerConnection.ToPtr()},
							},
							Name: to.StringPtr("<name>"),
						}},
					PublicNetwork: &armsignalr.NetworkACL{
						Allow: []*armsignalr.SignalRRequestType{
							armsignalr.SignalRRequestTypeClientConnection.ToPtr()},
					},
				},
				PublicNetworkAccess: to.StringPtr("<public-network-access>"),
				TLS: &armsignalr.SignalRTLSSettings{
					ClientCertEnabled: to.BoolPtr(false),
				},
				Upstream: &armsignalr.ServerlessUpstreamSettings{
					Templates: []*armsignalr.UpstreamTemplate{
						{
							Auth: &armsignalr.UpstreamAuthSettings{
								Type: armsignalr.UpstreamAuthTypeManagedIdentity.ToPtr(),
								ManagedIdentity: &armsignalr.ManagedIdentitySettings{
									Resource: to.StringPtr("<resource>"),
								},
							},
							CategoryPattern: to.StringPtr("<category-pattern>"),
							EventPattern:    to.StringPtr("<event-pattern>"),
							HubPattern:      to.StringPtr("<hub-pattern>"),
							URLTemplate:     to.StringPtr("<urltemplate>"),
						}},
				},
			},
			SKU: &armsignalr.ResourceSKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int32Ptr(1),
				Tier:     armsignalr.SignalRSKUTierStandard.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SignalRResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_ListKeys.json
func ExampleSignalRClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	_, err = client.ListKeys(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_RegenerateKey.json
func ExampleSignalRClient_BeginRegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRegenerateKey(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armsignalr.RegenerateKeyParameters{
			KeyType: armsignalr.KeyTypePrimary.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_Restart.json
func ExampleSignalRClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_ListSkus.json
func ExampleSignalRClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsignalr.NewSignalRClient("<subscription-id>", cred, nil)
	_, err = client.ListSKUs(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
