//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

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
	"strconv"
	"strings"
)

// ContainersClient contains the methods for the Containers group.
// Don't use this type directly, use NewContainersClient() instead.
type ContainersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewContainersClient creates a new instance of ContainersClient with the specified values.
func NewContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContainersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ContainersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Attach - Attach to the output stream of a specific container instance in a specified resource group and container group.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) Attach(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersAttachOptions) (ContainersAttachResponse, error) {
	req, err := client.attachCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, options)
	if err != nil {
		return ContainersAttachResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersAttachResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersAttachResponse{}, client.attachHandleError(resp)
	}
	return client.attachHandleResponse(resp)
}

// attachCreateRequest creates the Attach request.
func (client *ContainersClient) attachCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersAttachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/attach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// attachHandleResponse handles the Attach response.
func (client *ContainersClient) attachHandleResponse(resp *http.Response) (ContainersAttachResponse, error) {
	result := ContainersAttachResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAttachResponse); err != nil {
		return ContainersAttachResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// attachHandleError handles the Attach error response.
func (client *ContainersClient) attachHandleError(resp *http.Response) error {
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

// ExecuteCommand - Executes a command for a specific container instance in a specified resource group and container group.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) ExecuteCommand(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest, options *ContainersExecuteCommandOptions) (ContainersExecuteCommandResponse, error) {
	req, err := client.executeCommandCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, containerExecRequest, options)
	if err != nil {
		return ContainersExecuteCommandResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersExecuteCommandResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersExecuteCommandResponse{}, client.executeCommandHandleError(resp)
	}
	return client.executeCommandHandleResponse(resp)
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *ContainersClient) executeCommandCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest, options *ContainersExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/exec"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerExecRequest)
}

// executeCommandHandleResponse handles the ExecuteCommand response.
func (client *ContainersClient) executeCommandHandleResponse(resp *http.Response) (ContainersExecuteCommandResponse, error) {
	result := ContainersExecuteCommandResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerExecResponse); err != nil {
		return ContainersExecuteCommandResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// executeCommandHandleError handles the ExecuteCommand error response.
func (client *ContainersClient) executeCommandHandleError(resp *http.Response) error {
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

// ListLogs - Get the logs for a specified container instance in a specified resource group and container group.
// If the operation fails it returns the *CloudError error type.
func (client *ContainersClient) ListLogs(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersListLogsOptions) (ContainersListLogsResponse, error) {
	req, err := client.listLogsCreateRequest(ctx, resourceGroupName, containerGroupName, containerName, options)
	if err != nil {
		return ContainersListLogsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainersListLogsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainersListLogsResponse{}, client.listLogsHandleError(resp)
	}
	return client.listLogsHandleResponse(resp)
}

// listLogsCreateRequest creates the ListLogs request.
func (client *ContainersClient) listLogsCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, options *ContainersListLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	if options != nil && options.Tail != nil {
		reqQP.Set("tail", strconv.FormatInt(int64(*options.Tail), 10))
	}
	if options != nil && options.Timestamps != nil {
		reqQP.Set("timestamps", strconv.FormatBool(*options.Timestamps))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listLogsHandleResponse handles the ListLogs response.
func (client *ContainersClient) listLogsHandleResponse(resp *http.Response) (ContainersListLogsResponse, error) {
	result := ContainersListLogsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Logs); err != nil {
		return ContainersListLogsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listLogsHandleError handles the ListLogs error response.
func (client *ContainersClient) listLogsHandleError(resp *http.Response) error {
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
