// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListBookmarkedRunsParams creates a new ListBookmarkedRunsParams object
// with the default values initialized.
func NewListBookmarkedRunsParams() *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBookmarkedRunsParamsWithTimeout creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBookmarkedRunsParamsWithTimeout(timeout time.Duration) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		timeout: timeout,
	}
}

// NewListBookmarkedRunsParamsWithContext creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBookmarkedRunsParamsWithContext(ctx context.Context) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{

		Context: ctx,
	}
}

// NewListBookmarkedRunsParamsWithHTTPClient creates a new ListBookmarkedRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBookmarkedRunsParamsWithHTTPClient(client *http.Client) *ListBookmarkedRunsParams {
	var ()
	return &ListBookmarkedRunsParams{
		HTTPClient: client,
	}
}

/*ListBookmarkedRunsParams contains all the parameters to send to the API endpoint
for the list bookmarked runs operation typically these are written to a http.Request
*/
type ListBookmarkedRunsParams struct {

	/*User
	  Owner of the namespace

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithTimeout(timeout time.Duration) *ListBookmarkedRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithContext(ctx context.Context) *ListBookmarkedRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithHTTPClient(client *http.Client) *ListBookmarkedRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) WithUser(user string) *ListBookmarkedRunsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list bookmarked runs params
func (o *ListBookmarkedRunsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListBookmarkedRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}