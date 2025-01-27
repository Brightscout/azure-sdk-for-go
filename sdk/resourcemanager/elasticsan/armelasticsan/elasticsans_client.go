//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan

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
	"strings"
)

// ElasticSansClient contains the methods for the ElasticSans group.
// Don't use this type directly, use NewElasticSansClient() instead.
type ElasticSansClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewElasticSansClient creates a new instance of ElasticSansClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewElasticSansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ElasticSansClient, error) {
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
	client := &ElasticSansClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create ElasticSan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// parameters - Elastic San object.
// options - ElasticSansClientBeginCreateOptions contains the optional parameters for the ElasticSansClient.BeginCreate method.
func (client *ElasticSansClient) BeginCreate(ctx context.Context, resourceGroupName string, elasticSanName string, parameters ElasticSan, options *ElasticSansClientBeginCreateOptions) (*runtime.Poller[ElasticSansClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, elasticSanName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ElasticSansClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ElasticSansClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create ElasticSan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
func (client *ElasticSansClient) create(ctx context.Context, resourceGroupName string, elasticSanName string, parameters ElasticSan, options *ElasticSansClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, elasticSanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ElasticSansClient) createCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, parameters ElasticSan, options *ElasticSansClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a Elastic San.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// options - ElasticSansClientBeginDeleteOptions contains the optional parameters for the ElasticSansClient.BeginDelete method.
func (client *ElasticSansClient) BeginDelete(ctx context.Context, resourceGroupName string, elasticSanName string, options *ElasticSansClientBeginDeleteOptions) (*runtime.Poller[ElasticSansClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, elasticSanName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ElasticSansClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ElasticSansClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Elastic San.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
func (client *ElasticSansClient) deleteOperation(ctx context.Context, resourceGroupName string, elasticSanName string, options *ElasticSansClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, elasticSanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ElasticSansClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, options *ElasticSansClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ElasticSan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// options - ElasticSansClientGetOptions contains the optional parameters for the ElasticSansClient.Get method.
func (client *ElasticSansClient) Get(ctx context.Context, resourceGroupName string, elasticSanName string, options *ElasticSansClientGetOptions) (ElasticSansClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, elasticSanName, options)
	if err != nil {
		return ElasticSansClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ElasticSansClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ElasticSansClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ElasticSansClient) getCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, options *ElasticSansClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ElasticSansClient) getHandleResponse(resp *http.Response) (ElasticSansClientGetResponse, error) {
	result := ElasticSansClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticSan); err != nil {
		return ElasticSansClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of ElasticSan in a resource group.
// Generated from API version 2021-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ElasticSansClientListByResourceGroupOptions contains the optional parameters for the ElasticSansClient.ListByResourceGroup
// method.
func (client *ElasticSansClient) NewListByResourceGroupPager(resourceGroupName string, options *ElasticSansClientListByResourceGroupOptions) *runtime.Pager[ElasticSansClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ElasticSansClientListByResourceGroupResponse]{
		More: func(page ElasticSansClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ElasticSansClientListByResourceGroupResponse) (ElasticSansClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ElasticSansClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticSansClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticSansClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ElasticSansClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ElasticSansClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ElasticSansClient) listByResourceGroupHandleResponse(resp *http.Response) (ElasticSansClientListByResourceGroupResponse, error) {
	result := ElasticSansClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return ElasticSansClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of ElasticSans in a subscription
// Generated from API version 2021-11-20-preview
// options - ElasticSansClientListBySubscriptionOptions contains the optional parameters for the ElasticSansClient.ListBySubscription
// method.
func (client *ElasticSansClient) NewListBySubscriptionPager(options *ElasticSansClientListBySubscriptionOptions) *runtime.Pager[ElasticSansClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ElasticSansClientListBySubscriptionResponse]{
		More: func(page ElasticSansClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ElasticSansClientListBySubscriptionResponse) (ElasticSansClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ElasticSansClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticSansClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticSansClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ElasticSansClient) listBySubscriptionCreateRequest(ctx context.Context, options *ElasticSansClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ElasticSan/elasticSans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ElasticSansClient) listBySubscriptionHandleResponse(resp *http.Response) (ElasticSansClientListBySubscriptionResponse, error) {
	result := ElasticSansClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return ElasticSansClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Elastic San.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// elasticSanName - The name of the ElasticSan.
// parameters - Elastic San object.
// options - ElasticSansClientBeginUpdateOptions contains the optional parameters for the ElasticSansClient.BeginUpdate method.
func (client *ElasticSansClient) BeginUpdate(ctx context.Context, resourceGroupName string, elasticSanName string, parameters Update, options *ElasticSansClientBeginUpdateOptions) (*runtime.Poller[ElasticSansClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, elasticSanName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ElasticSansClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ElasticSansClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a Elastic San.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-20-preview
func (client *ElasticSansClient) update(ctx context.Context, resourceGroupName string, elasticSanName string, parameters Update, options *ElasticSansClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, elasticSanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ElasticSansClient) updateCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, parameters Update, options *ElasticSansClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-20-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
