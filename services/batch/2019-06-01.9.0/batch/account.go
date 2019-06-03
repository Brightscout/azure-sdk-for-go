package batch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// AccountClient is the a client for issuing REST requests to the Azure Batch service.
type AccountClient struct {
	BaseClient
}

// NewAccountClient creates an instance of the AccountClient client.
func NewAccountClient(batchURL string) AccountClient {
	return AccountClient{New(batchURL)}
}

// ListPoolNodeCounts gets the number of nodes in each state, grouped by pool.
// Parameters:
// filter - an OData $filter clause. For more information on constructing this filter, see
// https://docs.microsoft.com/en-us/rest/api/batchservice/odata-filters-in-batch.
// maxResults - the maximum number of items to return in the response.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client AccountClient) ListPoolNodeCounts(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result PoolNodeCountsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListPoolNodeCounts")
		defer func() {
			sc := -1
			if result.pnclr.Response.Response != nil {
				sc = result.pnclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(10), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "ListPoolNodeCounts", err.Error())
	}

	result.fn = client.listPoolNodeCountsNextResults
	req, err := client.ListPoolNodeCountsPreparer(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPoolNodeCountsSender(req)
	if err != nil {
		result.pnclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", resp, "Failure sending request")
		return
	}

	result.pnclr, err = client.ListPoolNodeCountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListPoolNodeCounts", resp, "Failure responding to request")
	}

	return
}

// ListPoolNodeCountsPreparer prepares the ListPoolNodeCounts request.
func (client AccountClient) ListPoolNodeCountsPreparer(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2019-06-01.9.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 10)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/nodecounts"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPoolNodeCountsSender sends the ListPoolNodeCounts request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListPoolNodeCountsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListPoolNodeCountsResponder handles the response to the ListPoolNodeCounts request. The method always
// closes the http.Response Body.
func (client AccountClient) ListPoolNodeCountsResponder(resp *http.Response) (result PoolNodeCountsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listPoolNodeCountsNextResults retrieves the next set of results, if any.
func (client AccountClient) listPoolNodeCountsNextResults(ctx context.Context, lastResults PoolNodeCountsListResult) (result PoolNodeCountsListResult, err error) {
	req, err := lastResults.poolNodeCountsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListPoolNodeCountsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListPoolNodeCountsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listPoolNodeCountsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListPoolNodeCountsComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListPoolNodeCountsComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result PoolNodeCountsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListPoolNodeCounts")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListPoolNodeCounts(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	return
}

// ListSupportedImages sends the list supported images request.
// Parameters:
// filter - an OData $filter clause. For more information on constructing this filter, see
// https://docs.microsoft.com/en-us/rest/api/batchservice/odata-filters-in-batch#list-support-images.
// maxResults - the maximum number of items to return in the response. A maximum of 1000 results will be
// returned.
// timeout - the maximum time that the server can spend processing the request, in seconds. The default is 30
// seconds.
// clientRequestID - the caller-generated request identity, in the form of a GUID with no decoration such as
// curly braces, e.g. 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0.
// returnClientRequestID - whether the server should return the client-request-id in the response.
// ocpDate - the time the request was issued. Client libraries typically set this to the current system clock
// time; set it explicitly if you are calling the REST API directly.
func (client AccountClient) ListSupportedImages(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result AccountListSupportedImagesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListSupportedImages")
		defer func() {
			sc := -1
			if result.alsir.Response.Response != nil {
				sc = result.alsir.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "ListSupportedImages", err.Error())
	}

	result.fn = client.listSupportedImagesNextResults
	req, err := client.ListSupportedImagesPreparer(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListSupportedImages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSupportedImagesSender(req)
	if err != nil {
		result.alsir.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListSupportedImages", resp, "Failure sending request")
		return
	}

	result.alsir, err = client.ListSupportedImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListSupportedImages", resp, "Failure responding to request")
	}

	return
}

// ListSupportedImagesPreparer prepares the ListSupportedImages request.
func (client AccountClient) ListSupportedImagesPreparer(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"batchUrl": client.BatchURL,
	}

	const APIVersion = "2019-06-01.9.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	} else {
		queryParameters["timeout"] = autorest.Encode("query", 30)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{batchUrl}", urlParameters),
		autorest.WithPath("/supportedimages"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(false)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSupportedImagesSender sends the ListSupportedImages request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListSupportedImagesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListSupportedImagesResponder handles the response to the ListSupportedImages request. The method always
// closes the http.Response Body.
func (client AccountClient) ListSupportedImagesResponder(resp *http.Response) (result AccountListSupportedImagesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listSupportedImagesNextResults retrieves the next set of results, if any.
func (client AccountClient) listSupportedImagesNextResults(ctx context.Context, lastResults AccountListSupportedImagesResult) (result AccountListSupportedImagesResult, err error) {
	req, err := lastResults.accountListSupportedImagesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listSupportedImagesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSupportedImagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listSupportedImagesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListSupportedImagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listSupportedImagesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListSupportedImagesComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListSupportedImagesComplete(ctx context.Context, filter string, maxResults *int32, timeout *int32, clientRequestID *uuid.UUID, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result AccountListSupportedImagesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListSupportedImages")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListSupportedImages(ctx, filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	return
}
