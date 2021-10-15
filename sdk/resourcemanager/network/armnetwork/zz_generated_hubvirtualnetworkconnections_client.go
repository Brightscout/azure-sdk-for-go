//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// HubVirtualNetworkConnectionsClient contains the methods for the HubVirtualNetworkConnections group.
// Don't use this type directly, use NewHubVirtualNetworkConnectionsClient() instead.
type HubVirtualNetworkConnectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewHubVirtualNetworkConnectionsClient creates a new instance of HubVirtualNetworkConnectionsClient with the specified values.
func NewHubVirtualNetworkConnectionsClient(con *arm.Connection, subscriptionID string) *HubVirtualNetworkConnectionsClient {
	return &HubVirtualNetworkConnectionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a hub virtual network connection if it doesn't exist else updates the existing one.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, options *HubVirtualNetworkConnectionsBeginCreateOrUpdateOptions) (HubVirtualNetworkConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, connectionName, hubVirtualNetworkConnectionParameters, options)
	if err != nil {
		return HubVirtualNetworkConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := HubVirtualNetworkConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("HubVirtualNetworkConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return HubVirtualNetworkConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &HubVirtualNetworkConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a hub virtual network connection if it doesn't exist else updates the existing one.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, options *HubVirtualNetworkConnectionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, hubVirtualNetworkConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HubVirtualNetworkConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, options *HubVirtualNetworkConnectionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, hubVirtualNetworkConnectionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *HubVirtualNetworkConnectionsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes a HubVirtualNetworkConnection.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsBeginDeleteOptions) (HubVirtualNetworkConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return HubVirtualNetworkConnectionsDeletePollerResponse{}, err
	}
	result := HubVirtualNetworkConnectionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("HubVirtualNetworkConnectionsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return HubVirtualNetworkConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &HubVirtualNetworkConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a HubVirtualNetworkConnection.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
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
func (client *HubVirtualNetworkConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *HubVirtualNetworkConnectionsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieves the details of a HubVirtualNetworkConnection.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsGetOptions) (HubVirtualNetworkConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return HubVirtualNetworkConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubVirtualNetworkConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HubVirtualNetworkConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HubVirtualNetworkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubVirtualNetworkConnectionsClient) getHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsGetResponse, error) {
	result := HubVirtualNetworkConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubVirtualNetworkConnection); err != nil {
		return HubVirtualNetworkConnectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *HubVirtualNetworkConnectionsClient) getHandleError(resp *http.Response) error {
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

// List - Retrieves the details of all HubVirtualNetworkConnections.
// If the operation fails it returns the *CloudError error type.
func (client *HubVirtualNetworkConnectionsClient) List(resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsListOptions) *HubVirtualNetworkConnectionsListPager {
	return &HubVirtualNetworkConnectionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		advancer: func(ctx context.Context, resp HubVirtualNetworkConnectionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListHubVirtualNetworkConnectionsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *HubVirtualNetworkConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubVirtualNetworkConnectionsClient) listHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsListResponse, error) {
	result := HubVirtualNetworkConnectionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListHubVirtualNetworkConnectionsResult); err != nil {
		return HubVirtualNetworkConnectionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *HubVirtualNetworkConnectionsClient) listHandleError(resp *http.Response) error {
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