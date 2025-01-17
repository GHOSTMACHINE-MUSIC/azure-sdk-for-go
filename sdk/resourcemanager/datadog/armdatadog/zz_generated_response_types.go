//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// MarketplaceAgreementsCreateOrUpdateResponse contains the response from method MarketplaceAgreements.CreateOrUpdate.
type MarketplaceAgreementsCreateOrUpdateResponse struct {
	MarketplaceAgreementsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceAgreementsCreateOrUpdateResult contains the result from method MarketplaceAgreements.CreateOrUpdate.
type MarketplaceAgreementsCreateOrUpdateResult struct {
	DatadogAgreementResource
}

// MarketplaceAgreementsListResponse contains the response from method MarketplaceAgreements.List.
type MarketplaceAgreementsListResponse struct {
	MarketplaceAgreementsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceAgreementsListResult contains the result from method MarketplaceAgreements.List.
type MarketplaceAgreementsListResult struct {
	DatadogAgreementResourceListResponse
}

// MonitorsCreatePollerResponse contains the response from method Monitors.Create.
type MonitorsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsCreateResponse, error) {
	respType := MonitorsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DatadogMonitorResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsCreatePollerResponse from the provided client and resume token.
func (l *MonitorsCreatePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &MonitorsCreatePoller{
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

// MonitorsCreateResponse contains the response from method Monitors.Create.
type MonitorsCreateResponse struct {
	MonitorsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsCreateResult contains the result from method Monitors.Create.
type MonitorsCreateResult struct {
	DatadogMonitorResource
}

// MonitorsDeletePollerResponse contains the response from method Monitors.Delete.
type MonitorsDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsDeleteResponse, error) {
	respType := MonitorsDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsDeletePollerResponse from the provided client and resume token.
func (l *MonitorsDeletePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &MonitorsDeletePoller{
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

// MonitorsDeleteResponse contains the response from method Monitors.Delete.
type MonitorsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsGetDefaultKeyResponse contains the response from method Monitors.GetDefaultKey.
type MonitorsGetDefaultKeyResponse struct {
	MonitorsGetDefaultKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsGetDefaultKeyResult contains the result from method Monitors.GetDefaultKey.
type MonitorsGetDefaultKeyResult struct {
	DatadogAPIKey
}

// MonitorsGetResponse contains the response from method Monitors.Get.
type MonitorsGetResponse struct {
	MonitorsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsGetResult contains the result from method Monitors.Get.
type MonitorsGetResult struct {
	DatadogMonitorResource
}

// MonitorsListAPIKeysResponse contains the response from method Monitors.ListAPIKeys.
type MonitorsListAPIKeysResponse struct {
	MonitorsListAPIKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListAPIKeysResult contains the result from method Monitors.ListAPIKeys.
type MonitorsListAPIKeysResult struct {
	DatadogAPIKeyListResponse
}

// MonitorsListByResourceGroupResponse contains the response from method Monitors.ListByResourceGroup.
type MonitorsListByResourceGroupResponse struct {
	MonitorsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListByResourceGroupResult contains the result from method Monitors.ListByResourceGroup.
type MonitorsListByResourceGroupResult struct {
	DatadogMonitorResourceListResponse
}

// MonitorsListHostsResponse contains the response from method Monitors.ListHosts.
type MonitorsListHostsResponse struct {
	MonitorsListHostsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListHostsResult contains the result from method Monitors.ListHosts.
type MonitorsListHostsResult struct {
	DatadogHostListResponse
}

// MonitorsListLinkedResourcesResponse contains the response from method Monitors.ListLinkedResources.
type MonitorsListLinkedResourcesResponse struct {
	MonitorsListLinkedResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListLinkedResourcesResult contains the result from method Monitors.ListLinkedResources.
type MonitorsListLinkedResourcesResult struct {
	LinkedResourceListResponse
}

// MonitorsListMonitoredResourcesResponse contains the response from method Monitors.ListMonitoredResources.
type MonitorsListMonitoredResourcesResponse struct {
	MonitorsListMonitoredResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListMonitoredResourcesResult contains the result from method Monitors.ListMonitoredResources.
type MonitorsListMonitoredResourcesResult struct {
	MonitoredResourceListResponse
}

// MonitorsListResponse contains the response from method Monitors.List.
type MonitorsListResponse struct {
	MonitorsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsListResult contains the result from method Monitors.List.
type MonitorsListResult struct {
	DatadogMonitorResourceListResponse
}

// MonitorsRefreshSetPasswordLinkResponse contains the response from method Monitors.RefreshSetPasswordLink.
type MonitorsRefreshSetPasswordLinkResponse struct {
	MonitorsRefreshSetPasswordLinkResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsRefreshSetPasswordLinkResult contains the result from method Monitors.RefreshSetPasswordLink.
type MonitorsRefreshSetPasswordLinkResult struct {
	DatadogSetPasswordLink
}

// MonitorsSetDefaultKeyResponse contains the response from method Monitors.SetDefaultKey.
type MonitorsSetDefaultKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsUpdatePollerResponse contains the response from method Monitors.Update.
type MonitorsUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsUpdateResponse, error) {
	respType := MonitorsUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DatadogMonitorResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsUpdatePollerResponse from the provided client and resume token.
func (l *MonitorsUpdatePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Update", token, client.pl, client.updateHandleError)
	if err != nil {
		return err
	}
	poller := &MonitorsUpdatePoller{
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

// MonitorsUpdateResponse contains the response from method Monitors.Update.
type MonitorsUpdateResponse struct {
	MonitorsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsUpdateResult contains the result from method Monitors.Update.
type MonitorsUpdateResult struct {
	DatadogMonitorResource
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

// SingleSignOnConfigurationsCreateOrUpdatePollerResponse contains the response from method SingleSignOnConfigurations.CreateOrUpdate.
type SingleSignOnConfigurationsCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SingleSignOnConfigurationsCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SingleSignOnConfigurationsCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SingleSignOnConfigurationsCreateOrUpdateResponse, error) {
	respType := SingleSignOnConfigurationsCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DatadogSingleSignOnResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SingleSignOnConfigurationsCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *SingleSignOnConfigurationsCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *SingleSignOnConfigurationsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SingleSignOnConfigurationsClient.CreateOrUpdate", token, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return err
	}
	poller := &SingleSignOnConfigurationsCreateOrUpdatePoller{
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

// SingleSignOnConfigurationsCreateOrUpdateResponse contains the response from method SingleSignOnConfigurations.CreateOrUpdate.
type SingleSignOnConfigurationsCreateOrUpdateResponse struct {
	SingleSignOnConfigurationsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnConfigurationsCreateOrUpdateResult contains the result from method SingleSignOnConfigurations.CreateOrUpdate.
type SingleSignOnConfigurationsCreateOrUpdateResult struct {
	DatadogSingleSignOnResource
}

// SingleSignOnConfigurationsGetResponse contains the response from method SingleSignOnConfigurations.Get.
type SingleSignOnConfigurationsGetResponse struct {
	SingleSignOnConfigurationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnConfigurationsGetResult contains the result from method SingleSignOnConfigurations.Get.
type SingleSignOnConfigurationsGetResult struct {
	DatadogSingleSignOnResource
}

// SingleSignOnConfigurationsListResponse contains the response from method SingleSignOnConfigurations.List.
type SingleSignOnConfigurationsListResponse struct {
	SingleSignOnConfigurationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnConfigurationsListResult contains the result from method SingleSignOnConfigurations.List.
type SingleSignOnConfigurationsListResult struct {
	DatadogSingleSignOnResourceListResponse
}

// TagRulesCreateOrUpdateResponse contains the response from method TagRules.CreateOrUpdate.
type TagRulesCreateOrUpdateResponse struct {
	TagRulesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesCreateOrUpdateResult contains the result from method TagRules.CreateOrUpdate.
type TagRulesCreateOrUpdateResult struct {
	MonitoringTagRules
}

// TagRulesGetResponse contains the response from method TagRules.Get.
type TagRulesGetResponse struct {
	TagRulesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesGetResult contains the result from method TagRules.Get.
type TagRulesGetResult struct {
	MonitoringTagRules
}

// TagRulesListResponse contains the response from method TagRules.List.
type TagRulesListResponse struct {
	TagRulesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesListResult contains the result from method TagRules.List.
type TagRulesListResult struct {
	MonitoringTagRulesListResponse
}
