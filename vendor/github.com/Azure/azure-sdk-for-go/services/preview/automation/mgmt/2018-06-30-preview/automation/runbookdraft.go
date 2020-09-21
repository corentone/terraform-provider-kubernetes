package automation

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"io"
	"net/http"
)

// RunbookDraftClient is the automation Client
type RunbookDraftClient struct {
	BaseClient
}

// NewRunbookDraftClient creates an instance of the RunbookDraftClient client.
func NewRunbookDraftClient(subscriptionID string) RunbookDraftClient {
	return NewRunbookDraftClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRunbookDraftClientWithBaseURI creates an instance of the RunbookDraftClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRunbookDraftClientWithBaseURI(baseURI string, subscriptionID string) RunbookDraftClient {
	return RunbookDraftClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the runbook draft identified by runbook name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
func (client RunbookDraftClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result RunbookDraft, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunbookDraftClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.RunbookDraftClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RunbookDraftClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetResponder(resp *http.Response) (result RunbookDraft, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetContent retrieve the content of runbook draft identified by runbook name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
func (client RunbookDraftClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunbookDraftClient.GetContent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.RunbookDraftClient", "GetContent", err.Error())
	}

	req, err := client.GetContentPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure responding to request")
	}

	return
}

// GetContentPreparer prepares the GetContent request.
func (client RunbookDraftClient) GetContentPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContentSender sends the GetContent request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetContentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetContentResponder handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// ReplaceContent replaces the runbook draft content.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
// runbookContent - the runbook draft content.
func (client RunbookDraftClient) ReplaceContent(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, runbookContent io.ReadCloser) (result RunbookDraftReplaceContentFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunbookDraftClient.ReplaceContent")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.RunbookDraftClient", "ReplaceContent", err.Error())
	}

	req, err := client.ReplaceContentPreparer(ctx, resourceGroupName, automationAccountName, runbookName, runbookContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "ReplaceContent", nil, "Failure preparing request")
		return
	}

	result, err = client.ReplaceContentSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "ReplaceContent", result.Response(), "Failure sending request")
		return
	}

	return
}

// ReplaceContentPreparer prepares the ReplaceContent request.
func (client RunbookDraftClient) ReplaceContentPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, runbookContent io.ReadCloser) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("text/powershell"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithFile(runbookContent),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReplaceContentSender sends the ReplaceContent request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) ReplaceContentSender(req *http.Request) (future RunbookDraftReplaceContentFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ReplaceContentResponder handles the response to the ReplaceContent request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) ReplaceContentResponder(resp *http.Response) (result ReadCloser, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UndoEdit undo draft edit to last known published state identified by runbook name.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// runbookName - the runbook name.
func (client RunbookDraftClient) UndoEdit(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (result RunbookDraftUndoEditResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunbookDraftClient.UndoEdit")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.RunbookDraftClient", "UndoEdit", err.Error())
	}

	req, err := client.UndoEditPreparer(ctx, resourceGroupName, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", nil, "Failure preparing request")
		return
	}

	resp, err := client.UndoEditSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure sending request")
		return
	}

	result, err = client.UndoEditResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure responding to request")
	}

	return
}

// UndoEditPreparer prepares the UndoEdit request.
func (client RunbookDraftClient) UndoEditPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/undoEdit", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UndoEditSender sends the UndoEdit request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) UndoEditSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UndoEditResponder handles the response to the UndoEdit request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) UndoEditResponder(resp *http.Response) (result RunbookDraftUndoEditResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
