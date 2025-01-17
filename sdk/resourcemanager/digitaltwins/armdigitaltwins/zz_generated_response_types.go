//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DigitalTwinsCheckNameAvailabilityResponse contains the response from method DigitalTwins.CheckNameAvailability.
type DigitalTwinsCheckNameAvailabilityResponse struct {
	DigitalTwinsCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsCheckNameAvailabilityResult contains the result from method DigitalTwins.CheckNameAvailability.
type DigitalTwinsCheckNameAvailabilityResult struct {
	CheckNameResult
}

// DigitalTwinsCreateOrUpdatePollerResponse contains the response from method DigitalTwins.CreateOrUpdate.
type DigitalTwinsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DigitalTwinsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DigitalTwinsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DigitalTwinsCreateOrUpdateResponse, error) {
	respType := DigitalTwinsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DigitalTwinsDescription)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DigitalTwinsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *DigitalTwinsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *DigitalTwinsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DigitalTwinsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &DigitalTwinsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DigitalTwinsCreateOrUpdateResponse contains the response from method DigitalTwins.CreateOrUpdate.
type DigitalTwinsCreateOrUpdateResponse struct {
	DigitalTwinsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsCreateOrUpdateResult contains the result from method DigitalTwins.CreateOrUpdate.
type DigitalTwinsCreateOrUpdateResult struct {
	DigitalTwinsDescription
}

// DigitalTwinsDeletePollerResponse contains the response from method DigitalTwins.Delete.
type DigitalTwinsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DigitalTwinsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DigitalTwinsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DigitalTwinsDeleteResponse, error) {
	respType := DigitalTwinsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DigitalTwinsDescription)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DigitalTwinsDeletePollerResponse from the provided client and resume token.
func (l *DigitalTwinsDeletePollerResponse) Resume(ctx context.Context, client *DigitalTwinsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DigitalTwinsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &DigitalTwinsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DigitalTwinsDeleteResponse contains the response from method DigitalTwins.Delete.
type DigitalTwinsDeleteResponse struct {
	DigitalTwinsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsDeleteResult contains the result from method DigitalTwins.Delete.
type DigitalTwinsDeleteResult struct {
	DigitalTwinsDescription
}

// DigitalTwinsEndpointCreateOrUpdatePollerResponse contains the response from method DigitalTwinsEndpoint.CreateOrUpdate.
type DigitalTwinsEndpointCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DigitalTwinsEndpointCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DigitalTwinsEndpointCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DigitalTwinsEndpointCreateOrUpdateResponse, error) {
	respType := DigitalTwinsEndpointCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DigitalTwinsEndpointResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DigitalTwinsEndpointCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *DigitalTwinsEndpointCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *DigitalTwinsEndpointClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DigitalTwinsEndpointClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &DigitalTwinsEndpointCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DigitalTwinsEndpointCreateOrUpdateResponse contains the response from method DigitalTwinsEndpoint.CreateOrUpdate.
type DigitalTwinsEndpointCreateOrUpdateResponse struct {
	DigitalTwinsEndpointCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsEndpointCreateOrUpdateResult contains the result from method DigitalTwinsEndpoint.CreateOrUpdate.
type DigitalTwinsEndpointCreateOrUpdateResult struct {
	DigitalTwinsEndpointResource
}

// DigitalTwinsEndpointDeletePollerResponse contains the response from method DigitalTwinsEndpoint.Delete.
type DigitalTwinsEndpointDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DigitalTwinsEndpointDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DigitalTwinsEndpointDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DigitalTwinsEndpointDeleteResponse, error) {
	respType := DigitalTwinsEndpointDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DigitalTwinsEndpointResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DigitalTwinsEndpointDeletePollerResponse from the provided client and resume token.
func (l *DigitalTwinsEndpointDeletePollerResponse) Resume(ctx context.Context, client *DigitalTwinsEndpointClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DigitalTwinsEndpointClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &DigitalTwinsEndpointDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DigitalTwinsEndpointDeleteResponse contains the response from method DigitalTwinsEndpoint.Delete.
type DigitalTwinsEndpointDeleteResponse struct {
	DigitalTwinsEndpointDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsEndpointDeleteResult contains the result from method DigitalTwinsEndpoint.Delete.
type DigitalTwinsEndpointDeleteResult struct {
	DigitalTwinsEndpointResource
}

// DigitalTwinsEndpointGetResponse contains the response from method DigitalTwinsEndpoint.Get.
type DigitalTwinsEndpointGetResponse struct {
	DigitalTwinsEndpointGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsEndpointGetResult contains the result from method DigitalTwinsEndpoint.Get.
type DigitalTwinsEndpointGetResult struct {
	DigitalTwinsEndpointResource
}

// DigitalTwinsEndpointListResponse contains the response from method DigitalTwinsEndpoint.List.
type DigitalTwinsEndpointListResponse struct {
	DigitalTwinsEndpointListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsEndpointListResult contains the result from method DigitalTwinsEndpoint.List.
type DigitalTwinsEndpointListResult struct {
	DigitalTwinsEndpointResourceListResult
}

// DigitalTwinsGetResponse contains the response from method DigitalTwins.Get.
type DigitalTwinsGetResponse struct {
	DigitalTwinsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsGetResult contains the result from method DigitalTwins.Get.
type DigitalTwinsGetResult struct {
	DigitalTwinsDescription
}

// DigitalTwinsListByResourceGroupResponse contains the response from method DigitalTwins.ListByResourceGroup.
type DigitalTwinsListByResourceGroupResponse struct {
	DigitalTwinsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsListByResourceGroupResult contains the result from method DigitalTwins.ListByResourceGroup.
type DigitalTwinsListByResourceGroupResult struct {
	DigitalTwinsDescriptionListResult
}

// DigitalTwinsListResponse contains the response from method DigitalTwins.List.
type DigitalTwinsListResponse struct {
	DigitalTwinsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsListResult contains the result from method DigitalTwins.List.
type DigitalTwinsListResult struct {
	DigitalTwinsDescriptionListResult
}

// DigitalTwinsUpdatePollerResponse contains the response from method DigitalTwins.Update.
type DigitalTwinsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DigitalTwinsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DigitalTwinsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DigitalTwinsUpdateResponse, error) {
	respType := DigitalTwinsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DigitalTwinsDescription)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DigitalTwinsUpdatePollerResponse from the provided client and resume token.
func (l *DigitalTwinsUpdatePollerResponse) Resume(ctx context.Context, client *DigitalTwinsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DigitalTwinsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &DigitalTwinsUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DigitalTwinsUpdateResponse contains the response from method DigitalTwins.Update.
type DigitalTwinsUpdateResponse struct {
	DigitalTwinsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DigitalTwinsUpdateResult contains the result from method DigitalTwins.Update.
type DigitalTwinsUpdateResult struct {
	DigitalTwinsDescription
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// PrivateEndpointConnectionsCreateOrUpdatePollerResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsCreateOrUpdateResponse, error) {
	respType := PrivateEndpointConnectionsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsCreateOrUpdateResponse contains the response from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResponse struct {
	PrivateEndpointConnectionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsCreateOrUpdateResult contains the result from method PrivateEndpointConnections.CreateOrUpdate.
type PrivateEndpointConnectionsCreateOrUpdateResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsDeletePollerResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PrivateEndpointConnectionsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l PrivateEndpointConnectionsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a PrivateEndpointConnectionsDeletePollerResponse from the provided client and resume token.
func (l *PrivateEndpointConnectionsDeletePollerResponse) Resume(ctx context.Context, client *PrivateEndpointConnectionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PrivateEndpointConnectionsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &PrivateEndpointConnectionsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PrivateEndpointConnectionsDeleteResponse contains the response from method PrivateEndpointConnections.Delete.
type PrivateEndpointConnectionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResponse contains the response from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResponse struct {
	PrivateEndpointConnectionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsGetResult contains the result from method PrivateEndpointConnections.Get.
type PrivateEndpointConnectionsGetResult struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsListResponse contains the response from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResponse struct {
	PrivateEndpointConnectionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionsListResult contains the result from method PrivateEndpointConnections.List.
type PrivateEndpointConnectionsListResult struct {
	PrivateEndpointConnectionsResponse
}

// PrivateLinkResourcesGetResponse contains the response from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResponse struct {
	PrivateLinkResourcesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesGetResult contains the result from method PrivateLinkResources.Get.
type PrivateLinkResourcesGetResult struct {
	GroupIDInformation
}

// PrivateLinkResourcesListResponse contains the response from method PrivateLinkResources.List.
type PrivateLinkResourcesListResponse struct {
	PrivateLinkResourcesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesListResult contains the result from method PrivateLinkResources.List.
type PrivateLinkResourcesListResult struct {
	GroupIDInformationResponse
}
