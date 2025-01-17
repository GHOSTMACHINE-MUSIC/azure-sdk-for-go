//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredata

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

// SQLServersClient contains the methods for the SQLServers group.
// Don't use this type directly, use NewSQLServersClient() instead.
type SQLServersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLServersClient creates a new instance of SQLServersClient with the specified values.
func NewSQLServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLServersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SQLServersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a SQL Server.
// If the operation fails it returns the *CloudError error type.
func (client *SQLServersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, parameters SQLServer, options *SQLServersCreateOrUpdateOptions) (SQLServersCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlServerRegistrationName, sqlServerName, parameters, options)
	if err != nil {
		return SQLServersCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLServersCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLServersCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLServersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, parameters SQLServer, options *SQLServersCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerRegistrationName == "" {
		return nil, errors.New("parameter sqlServerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerRegistrationName}", url.PathEscape(sqlServerRegistrationName))
	if sqlServerName == "" {
		return nil, errors.New("parameter sqlServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerName}", url.PathEscape(sqlServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-24-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SQLServersClient) createOrUpdateHandleResponse(resp *http.Response) (SQLServersCreateOrUpdateResponse, error) {
	result := SQLServersCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServer); err != nil {
		return SQLServersCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SQLServersClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a SQL Server.
// If the operation fails it returns the *CloudError error type.
func (client *SQLServersClient) Delete(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, options *SQLServersDeleteOptions) (SQLServersDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlServerRegistrationName, sqlServerName, options)
	if err != nil {
		return SQLServersDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLServersDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SQLServersDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SQLServersDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLServersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, options *SQLServersDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerRegistrationName == "" {
		return nil, errors.New("parameter sqlServerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerRegistrationName}", url.PathEscape(sqlServerRegistrationName))
	if sqlServerName == "" {
		return nil, errors.New("parameter sqlServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerName}", url.PathEscape(sqlServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-24-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SQLServersClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a SQL Server.
// If the operation fails it returns the *CloudError error type.
func (client *SQLServersClient) Get(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, options *SQLServersGetOptions) (SQLServersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlServerRegistrationName, sqlServerName, options)
	if err != nil {
		return SQLServersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLServersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLServersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, sqlServerName string, options *SQLServersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerRegistrationName == "" {
		return nil, errors.New("parameter sqlServerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerRegistrationName}", url.PathEscape(sqlServerRegistrationName))
	if sqlServerName == "" {
		return nil, errors.New("parameter sqlServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerName}", url.PathEscape(sqlServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-07-24-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLServersClient) getHandleResponse(resp *http.Response) (SQLServersGetResponse, error) {
	result := SQLServersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServer); err != nil {
		return SQLServersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLServersClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all SQL Servers in a SQL Server Registration.
// If the operation fails it returns the *CloudError error type.
func (client *SQLServersClient) ListByResourceGroup(resourceGroupName string, sqlServerRegistrationName string, options *SQLServersListByResourceGroupOptions) *SQLServersListByResourceGroupPager {
	return &SQLServersListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, sqlServerRegistrationName, options)
		},
		advancer: func(ctx context.Context, resp SQLServersListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLServerListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLServersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, sqlServerRegistrationName string, options *SQLServersListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerRegistrationName == "" {
		return nil, errors.New("parameter sqlServerRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerRegistrationName}", url.PathEscape(sqlServerRegistrationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2019-07-24-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLServersClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLServersListByResourceGroupResponse, error) {
	result := SQLServersListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerListResult); err != nil {
		return SQLServersListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SQLServersClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
