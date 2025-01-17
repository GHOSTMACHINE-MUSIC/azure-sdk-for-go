//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProviderHubClient contains the methods for the ProviderHub group.
// Don't use this type directly, use NewProviderHubClient() instead.
type ProviderHubClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewProviderHubClient creates a new instance of ProviderHubClient with the specified values.
func NewProviderHubClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProviderHubClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ProviderHubClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckinManifest - Checkin the manifest.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProviderHubClient) CheckinManifest(ctx context.Context, providerNamespace string, checkinManifestParams CheckinManifestParams, options *ProviderHubCheckinManifestOptions) (ProviderHubCheckinManifestResponse, error) {
	req, err := client.checkinManifestCreateRequest(ctx, providerNamespace, checkinManifestParams, options)
	if err != nil {
		return ProviderHubCheckinManifestResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderHubCheckinManifestResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderHubCheckinManifestResponse{}, client.checkinManifestHandleError(resp)
	}
	return client.checkinManifestHandleResponse(resp)
}

// checkinManifestCreateRequest creates the CheckinManifest request.
func (client *ProviderHubClient) checkinManifestCreateRequest(ctx context.Context, providerNamespace string, checkinManifestParams CheckinManifestParams, options *ProviderHubCheckinManifestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/checkinManifest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, checkinManifestParams)
}

// checkinManifestHandleResponse handles the CheckinManifest response.
func (client *ProviderHubClient) checkinManifestHandleResponse(resp *http.Response) (ProviderHubCheckinManifestResponse, error) {
	result := ProviderHubCheckinManifestResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckinManifestInfo); err != nil {
		return ProviderHubCheckinManifestResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkinManifestHandleError handles the CheckinManifest error response.
func (client *ProviderHubClient) checkinManifestHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GenerateManifest - Generates the manifest for the given provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ProviderHubClient) GenerateManifest(ctx context.Context, providerNamespace string, options *ProviderHubGenerateManifestOptions) (ProviderHubGenerateManifestResponse, error) {
	req, err := client.generateManifestCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return ProviderHubGenerateManifestResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProviderHubGenerateManifestResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProviderHubGenerateManifestResponse{}, client.generateManifestHandleError(resp)
	}
	return client.generateManifestHandleResponse(resp)
}

// generateManifestCreateRequest creates the GenerateManifest request.
func (client *ProviderHubClient) generateManifestCreateRequest(ctx context.Context, providerNamespace string, options *ProviderHubGenerateManifestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/generateManifest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// generateManifestHandleResponse handles the GenerateManifest response.
func (client *ProviderHubClient) generateManifestHandleResponse(resp *http.Response) (ProviderHubGenerateManifestResponse, error) {
	result := ProviderHubGenerateManifestResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceProviderManifest); err != nil {
		return ProviderHubGenerateManifestResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// generateManifestHandleError handles the GenerateManifest error response.
func (client *ProviderHubClient) generateManifestHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
