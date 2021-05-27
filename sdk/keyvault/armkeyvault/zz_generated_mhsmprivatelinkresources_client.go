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
	"strings"
)

// MHSMPrivateLinkResourcesClient contains the methods for the MHSMPrivateLinkResources group.
// Don't use this type directly, use NewMHSMPrivateLinkResourcesClient() instead.
type MHSMPrivateLinkResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewMHSMPrivateLinkResourcesClient creates a new instance of MHSMPrivateLinkResourcesClient with the specified values.
func NewMHSMPrivateLinkResourcesClient(con *armcore.Connection, subscriptionID string) *MHSMPrivateLinkResourcesClient {
	return &MHSMPrivateLinkResourcesClient{con: con, subscriptionID: subscriptionID}
}

// ListByMHSMResource - Gets the private link resources supported for the managed hsm pool.
// If the operation fails it returns the *CloudError error type.
func (client *MHSMPrivateLinkResourcesClient) ListByMHSMResource(ctx context.Context, resourceGroupName string, name string, options *MHSMPrivateLinkResourcesListByMHSMResourceOptions) (MHSMPrivateLinkResourceListResultResponse, error) {
	req, err := client.listByMHSMResourceCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return MHSMPrivateLinkResourceListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MHSMPrivateLinkResourceListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MHSMPrivateLinkResourceListResultResponse{}, client.listByMHSMResourceHandleError(resp)
	}
	return client.listByMHSMResourceHandleResponse(resp)
}

// listByMHSMResourceCreateRequest creates the ListByMHSMResource request.
func (client *MHSMPrivateLinkResourcesClient) listByMHSMResourceCreateRequest(ctx context.Context, resourceGroupName string, name string, options *MHSMPrivateLinkResourcesListByMHSMResourceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// listByMHSMResourceHandleResponse handles the ListByMHSMResource response.
func (client *MHSMPrivateLinkResourcesClient) listByMHSMResourceHandleResponse(resp *azcore.Response) (MHSMPrivateLinkResourceListResultResponse, error) {
	var val *MHSMPrivateLinkResourceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MHSMPrivateLinkResourceListResultResponse{}, err
	}
	return MHSMPrivateLinkResourceListResultResponse{RawResponse: resp.Response, MHSMPrivateLinkResourceListResult: val}, nil
}

// listByMHSMResourceHandleError handles the ListByMHSMResource error response.
func (client *MHSMPrivateLinkResourcesClient) listByMHSMResourceHandleError(resp *azcore.Response) error {
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