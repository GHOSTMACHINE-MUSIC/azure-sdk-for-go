// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// LoggerClient contains the methods for the Logger group.
// Don't use this type directly, use NewLoggerClient() instead.
type LoggerClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoggerClient creates a new instance of LoggerClient with the specified values.
func NewLoggerClient(con *armcore.Connection, subscriptionID string) *LoggerClient {
	return &LoggerClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or Updates a logger.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, parameters LoggerContract, options *LoggerCreateOrUpdateOptions) (LoggerCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, loggerID, parameters, options)
	if err != nil {
		return LoggerCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoggerCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return LoggerCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoggerClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, parameters LoggerContract, options *LoggerCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if loggerID == "" {
		return nil, errors.New("parameter loggerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerId}", url.PathEscape(loggerID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LoggerClient) createOrUpdateHandleResponse(resp *azcore.Response) (LoggerCreateOrUpdateResponse, error) {
	result := LoggerCreateOrUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.LoggerContract); err != nil {
		return LoggerCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LoggerClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes the specified logger.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, ifMatch string, options *LoggerDeleteOptions) (LoggerDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, loggerID, ifMatch, options)
	if err != nil {
		return LoggerDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoggerDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return LoggerDeleteResponse{}, client.deleteHandleError(resp)
	}
	return LoggerDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoggerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, ifMatch string, options *LoggerDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if loggerID == "" {
		return nil, errors.New("parameter loggerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerId}", url.PathEscape(loggerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LoggerClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the details of the logger specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) Get(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, options *LoggerGetOptions) (LoggerGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, loggerID, options)
	if err != nil {
		return LoggerGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoggerGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoggerGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoggerClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, options *LoggerGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if loggerID == "" {
		return nil, errors.New("parameter loggerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerId}", url.PathEscape(loggerID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoggerClient) getHandleResponse(resp *azcore.Response) (LoggerGetResponse, error) {
	result := LoggerGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.LoggerContract); err != nil {
		return LoggerGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LoggerClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Gets the entity state (Etag) version of the logger specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, options *LoggerGetEntityTagOptions) (LoggerGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, loggerID, options)
	if err != nil {
		return LoggerGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoggerGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *LoggerClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, options *LoggerGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if loggerID == "" {
		return nil, errors.New("parameter loggerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerId}", url.PathEscape(loggerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *LoggerClient) getEntityTagHandleResponse(resp *azcore.Response) (LoggerGetEntityTagResponse, error) {
	result := LoggerGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists a collection of loggers in the specified service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) ListByService(resourceGroupName string, serviceName string, options *LoggerListByServiceOptions) LoggerListByServicePager {
	return &loggerListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp LoggerListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoggerCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *LoggerClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *LoggerListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *LoggerClient) listByServiceHandleResponse(resp *azcore.Response) (LoggerListByServiceResponse, error) {
	result := LoggerListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LoggerCollection); err != nil {
		return LoggerListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *LoggerClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates an existing logger.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LoggerClient) Update(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, ifMatch string, parameters LoggerUpdateContract, options *LoggerUpdateOptions) (LoggerUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, loggerID, ifMatch, parameters, options)
	if err != nil {
		return LoggerUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoggerUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoggerUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *LoggerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, loggerID string, ifMatch string, parameters LoggerUpdateContract, options *LoggerUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if loggerID == "" {
		return nil, errors.New("parameter loggerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loggerId}", url.PathEscape(loggerID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *LoggerClient) updateHandleResponse(resp *azcore.Response) (LoggerUpdateResponse, error) {
	result := LoggerUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.LoggerContract); err != nil {
		return LoggerUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *LoggerClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}