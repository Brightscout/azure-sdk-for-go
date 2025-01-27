//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationAccountSchemasClient contains the methods for the IntegrationAccountSchemas group.
// Don't use this type directly, use NewIntegrationAccountSchemasClient() instead.
type IntegrationAccountSchemasClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationAccountSchemasClient creates a new instance of IntegrationAccountSchemasClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationAccountSchemasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountSchemasClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountSchemasClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// schemaName - The integration account schema name.
// schema - The integration account schema.
// options - IntegrationAccountSchemasClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountSchemasClient.CreateOrUpdate
// method.
func (client *IntegrationAccountSchemasClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema, options *IntegrationAccountSchemasClientCreateOrUpdateOptions) (IntegrationAccountSchemasClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, schema, options)
	if err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountSchemasClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, schema IntegrationAccountSchema, options *IntegrationAccountSchemasClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, schema)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountSchemasClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountSchemasClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchema); err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// schemaName - The integration account schema name.
// options - IntegrationAccountSchemasClientDeleteOptions contains the optional parameters for the IntegrationAccountSchemasClient.Delete
// method.
func (client *IntegrationAccountSchemasClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientDeleteOptions) (IntegrationAccountSchemasClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, options)
	if err != nil {
		return IntegrationAccountSchemasClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountSchemasClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationAccountSchemasClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountSchemasClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// schemaName - The integration account schema name.
// options - IntegrationAccountSchemasClientGetOptions contains the optional parameters for the IntegrationAccountSchemasClient.Get
// method.
func (client *IntegrationAccountSchemasClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientGetOptions) (IntegrationAccountSchemasClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, options)
	if err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountSchemasClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountSchemasClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountSchemasClient) getHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientGetResponse, error) {
	result := IntegrationAccountSchemasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchema); err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account schemas.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// options - IntegrationAccountSchemasClientListOptions contains the optional parameters for the IntegrationAccountSchemasClient.List
// method.
func (client *IntegrationAccountSchemasClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountSchemasClientListOptions) *runtime.Pager[IntegrationAccountSchemasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountSchemasClientListResponse]{
		More: func(page IntegrationAccountSchemasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountSchemasClientListResponse) (IntegrationAccountSchemasClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IntegrationAccountSchemasClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationAccountSchemasClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationAccountSchemasClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountSchemasClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountSchemasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountSchemasClient) listHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientListResponse, error) {
	result := IntegrationAccountSchemasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchemaListResult); err != nil {
		return IntegrationAccountSchemasClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// schemaName - The integration account schema name.
// options - IntegrationAccountSchemasClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountSchemasClient.ListContentCallbackURL
// method.
func (client *IntegrationAccountSchemasClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountSchemasClientListContentCallbackURLOptions) (IntegrationAccountSchemasClientListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, listContentCallbackURL, options)
	if err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountSchemasClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountSchemasClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, listContentCallbackURL)
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountSchemasClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountSchemasClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
