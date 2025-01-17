//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// SchedulesClient contains the methods for the Schedules group.
// Don't use this type directly, use NewSchedulesClient() instead.
type SchedulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSchedulesClient creates a new instance of SchedulesClient with the specified values.
func NewSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SchedulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SchedulesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Operation to create or update a lab schedule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body Schedule, options *SchedulesCreateOrUpdateOptions) (SchedulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, scheduleName, body, options)
	if err != nil {
		return SchedulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SchedulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body Schedule, options *SchedulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (SchedulesCreateOrUpdateResponse, error) {
	result := SchedulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SchedulesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Operation to delete a schedule resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesBeginDeleteOptions) (SchedulesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, labName, scheduleName, options)
	if err != nil {
		return SchedulesDeletePollerResponse{}, err
	}
	result := SchedulesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SchedulesClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SchedulesDeletePollerResponse{}, err
	}
	result.Poller = &SchedulesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a schedule resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, scheduleName, options)
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
func (client *SchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SchedulesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Returns the properties of a lab Schedule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesGetOptions) (SchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, scheduleName, options)
	if err != nil {
		return SchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, options *SchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SchedulesClient) getHandleResponse(resp *http.Response) (SchedulesGetResponse, error) {
	result := SchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SchedulesClient) getHandleError(resp *http.Response) error {
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

// ListByLab - Returns a list of all schedules for a lab.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) ListByLab(resourceGroupName string, labName string, options *SchedulesListByLabOptions) *SchedulesListByLabPager {
	return &SchedulesListByLabPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLabCreateRequest(ctx, resourceGroupName, labName, options)
		},
		advancer: func(ctx context.Context, resp SchedulesListByLabResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PagedSchedules.NextLink)
		},
	}
}

// listByLabCreateRequest creates the ListByLab request.
func (client *SchedulesClient) listByLabCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *SchedulesListByLabOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLabHandleResponse handles the ListByLab response.
func (client *SchedulesClient) listByLabHandleResponse(resp *http.Response) (SchedulesListByLabResponse, error) {
	result := SchedulesListByLabResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedSchedules); err != nil {
		return SchedulesListByLabResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByLabHandleError handles the ListByLab error response.
func (client *SchedulesClient) listByLabHandleError(resp *http.Response) error {
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

// Update - Operation to update a lab schedule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body ScheduleUpdate, options *SchedulesUpdateOptions) (SchedulesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, scheduleName, body, options)
	if err != nil {
		return SchedulesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, scheduleName string, body ScheduleUpdate, options *SchedulesUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/schedules/{scheduleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if scheduleName == "" {
		return nil, errors.New("parameter scheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleName}", url.PathEscape(scheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *SchedulesClient) updateHandleResponse(resp *http.Response) (SchedulesUpdateResponse, error) {
	result := SchedulesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *SchedulesClient) updateHandleError(resp *http.Response) error {
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
