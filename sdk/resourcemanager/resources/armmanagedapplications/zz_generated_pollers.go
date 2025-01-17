//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ApplicationDefinitionsCreateOrUpdateByIDPoller provides polling facilities until the operation reaches a terminal state.
type ApplicationDefinitionsCreateOrUpdateByIDPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationDefinitionsCreateOrUpdateByIDPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationDefinitionsCreateOrUpdateByIDPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationDefinitionsCreateOrUpdateByIDResponse will be returned.
func (p *ApplicationDefinitionsCreateOrUpdateByIDPoller) FinalResponse(ctx context.Context) (ApplicationDefinitionsCreateOrUpdateByIDResponse, error) {
	respType := ApplicationDefinitionsCreateOrUpdateByIDResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ApplicationDefinition)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdateByIDResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationDefinitionsCreateOrUpdateByIDPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationDefinitionsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ApplicationDefinitionsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationDefinitionsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationDefinitionsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationDefinitionsCreateOrUpdateResponse will be returned.
func (p *ApplicationDefinitionsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ApplicationDefinitionsCreateOrUpdateResponse, error) {
	respType := ApplicationDefinitionsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ApplicationDefinition)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationDefinitionsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationDefinitionsDeleteByIDPoller provides polling facilities until the operation reaches a terminal state.
type ApplicationDefinitionsDeleteByIDPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationDefinitionsDeleteByIDPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationDefinitionsDeleteByIDPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationDefinitionsDeleteByIDResponse will be returned.
func (p *ApplicationDefinitionsDeleteByIDPoller) FinalResponse(ctx context.Context) (ApplicationDefinitionsDeleteByIDResponse, error) {
	respType := ApplicationDefinitionsDeleteByIDResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ApplicationDefinitionsDeleteByIDResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationDefinitionsDeleteByIDPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationDefinitionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type ApplicationDefinitionsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationDefinitionsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationDefinitionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationDefinitionsDeleteResponse will be returned.
func (p *ApplicationDefinitionsDeletePoller) FinalResponse(ctx context.Context) (ApplicationDefinitionsDeleteResponse, error) {
	respType := ApplicationDefinitionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ApplicationDefinitionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationDefinitionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationsCreateOrUpdateByIDPoller provides polling facilities until the operation reaches a terminal state.
type ApplicationsCreateOrUpdateByIDPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationsCreateOrUpdateByIDPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationsCreateOrUpdateByIDPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationsCreateOrUpdateByIDResponse will be returned.
func (p *ApplicationsCreateOrUpdateByIDPoller) FinalResponse(ctx context.Context) (ApplicationsCreateOrUpdateByIDResponse, error) {
	respType := ApplicationsCreateOrUpdateByIDResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Application)
	if err != nil {
		return ApplicationsCreateOrUpdateByIDResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationsCreateOrUpdateByIDPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ApplicationsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationsCreateOrUpdateResponse will be returned.
func (p *ApplicationsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ApplicationsCreateOrUpdateResponse, error) {
	respType := ApplicationsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Application)
	if err != nil {
		return ApplicationsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationsDeleteByIDPoller provides polling facilities until the operation reaches a terminal state.
type ApplicationsDeleteByIDPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationsDeleteByIDPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationsDeleteByIDPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationsDeleteByIDResponse will be returned.
func (p *ApplicationsDeleteByIDPoller) FinalResponse(ctx context.Context) (ApplicationsDeleteByIDResponse, error) {
	respType := ApplicationsDeleteByIDResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ApplicationsDeleteByIDResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationsDeleteByIDPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ApplicationsDeletePoller provides polling facilities until the operation reaches a terminal state.
type ApplicationsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ApplicationsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ApplicationsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ApplicationsDeleteResponse will be returned.
func (p *ApplicationsDeletePoller) FinalResponse(ctx context.Context) (ApplicationsDeleteResponse, error) {
	respType := ApplicationsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ApplicationsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ApplicationsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
