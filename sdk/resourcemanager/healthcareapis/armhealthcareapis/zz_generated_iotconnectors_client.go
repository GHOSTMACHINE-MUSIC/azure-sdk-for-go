//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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

// IotConnectorsClient contains the methods for the IotConnectors group.
// Don't use this type directly, use NewIotConnectorsClient() instead.
type IotConnectorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIotConnectorsClient creates a new instance of IotConnectorsClient with the specified values.
func NewIotConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IotConnectorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IotConnectorsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsBeginCreateOrUpdateOptions) (IotConnectorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
	if err != nil {
		return IotConnectorsCreateOrUpdatePollerResponse{}, err
	}
	result := IotConnectorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IotConnectorsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return IotConnectorsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IotConnectorsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IotConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iotConnector)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IotConnectorsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes an IoT Connector.
// If the operation fails it returns the *Error error type.
func (client *IotConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsBeginDeleteOptions) (IotConnectorsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
	if err != nil {
		return IotConnectorsDeletePollerResponse{}, err
	}
	result := IotConnectorsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IotConnectorsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return IotConnectorsDeletePollerResponse{}, err
	}
	result.Poller = &IotConnectorsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an IoT Connector.
// If the operation fails it returns the *Error error type.
func (client *IotConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IotConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IotConnectorsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the properties of the specified IoT Connector.
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsGetOptions) (IotConnectorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, options)
	if err != nil {
		return IotConnectorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotConnectorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotConnectorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotConnectorsClient) getHandleResponse(resp *http.Response) (IotConnectorsGetResponse, error) {
	result := IotConnectorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnector); err != nil {
		return IotConnectorsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IotConnectorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByWorkspace - Lists all IoT Connectors for the given workspace
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *IotConnectorsListByWorkspaceOptions) *IotConnectorsListByWorkspacePager {
	return &IotConnectorsListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp IotConnectorsListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IotConnectorCollection.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *IotConnectorsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IotConnectorsListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *IotConnectorsClient) listByWorkspaceHandleResponse(resp *http.Response) (IotConnectorsListByWorkspaceResponse, error) {
	result := IotConnectorsListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnectorCollection); err != nil {
		return IotConnectorsListByWorkspaceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *IotConnectorsClient) listByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Patch an IoT Connector.
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) BeginUpdate(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsBeginUpdateOptions) (IotConnectorsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
	if err != nil {
		return IotConnectorsUpdatePollerResponse{}, err
	}
	result := IotConnectorsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IotConnectorsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return IotConnectorsUpdatePollerResponse{}, err
	}
	result.Poller = &IotConnectorsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch an IoT Connector.
// If the operation fails it returns the *ErrorDetails error type.
func (client *IotConnectorsClient) update(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *IotConnectorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, iotConnectorPatchResource)
}

// updateHandleError handles the Update error response.
func (client *IotConnectorsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDetails{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
