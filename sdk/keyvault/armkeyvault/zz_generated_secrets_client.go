// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SecretsClient contains the methods for the Secrets group.
// Don't use this type directly, use NewSecretsClient() instead.
type SecretsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSecretsClient creates a new instance of SecretsClient with the specified values.
func NewSecretsClient(con *armcore.Connection, subscriptionID string) *SecretsClient {
	return &SecretsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a secret in a key vault in the specified subscription. NOTE: This API is intended for internal use in ARM deployments.
// Users should use the data-plane REST service for interaction
// with vault secrets.
// If the operation fails it returns the *CloudError error type.
func (client *SecretsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters, options *SecretsCreateOrUpdateOptions) (SecretResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, secretName, parameters, options)
	if err != nil {
		return SecretResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecretResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return SecretResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecretsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretCreateOrUpdateParameters, options *SecretsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SecretsClient) createOrUpdateHandleResponse(resp *azcore.Response) (SecretResponse, error) {
	var val *Secret
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecretResponse{}, err
	}
	return SecretResponse{RawResponse: resp.Response, Secret: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecretsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified secret. NOTE: This API is intended for internal use in ARM deployments. Users should use the data-plane REST service for interaction
// with vault secrets.
// If the operation fails it returns the *CloudError error type.
func (client *SecretsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, secretName string, options *SecretsGetOptions) (SecretResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, secretName, options)
	if err != nil {
		return SecretResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecretResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecretResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecretsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, options *SecretsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecretsClient) getHandleResponse(resp *azcore.Response) (SecretResponse, error) {
	var val *Secret
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecretResponse{}, err
	}
	return SecretResponse{RawResponse: resp.Response, Secret: val}, nil
}

// getHandleError handles the Get error response.
func (client *SecretsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - The List operation gets information about the secrets in a vault. NOTE: This API is intended for internal use in ARM deployments. Users should
// use the data-plane REST service for interaction with
// vault secrets.
// If the operation fails it returns the *CloudError error type.
func (client *SecretsClient) List(resourceGroupName string, vaultName string, options *SecretsListOptions) SecretListResultPager {
	return &secretListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp SecretListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SecretListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *SecretsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *SecretsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecretsClient) listHandleResponse(resp *azcore.Response) (SecretListResultResponse, error) {
	var val *SecretListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecretListResultResponse{}, err
	}
	return SecretListResultResponse{RawResponse: resp.Response, SecretListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *SecretsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Update a secret in the specified subscription. NOTE: This API is intended for internal use in ARM deployments. Users should use the data-plane
// REST service for interaction with vault secrets.
// If the operation fails it returns the *CloudError error type.
func (client *SecretsClient) Update(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters, options *SecretsUpdateOptions) (SecretResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, secretName, parameters, options)
	if err != nil {
		return SecretResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecretResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return SecretResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SecretsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters SecretPatchParameters, options *SecretsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if secretName == "" {
		return nil, errors.New("parameter secretName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretName}", url.PathEscape(secretName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *SecretsClient) updateHandleResponse(resp *azcore.Response) (SecretResponse, error) {
	var val *Secret
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecretResponse{}, err
	}
	return SecretResponse{RawResponse: resp.Response, Secret: val}, nil
}

// updateHandleError handles the Update error response.
func (client *SecretsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}