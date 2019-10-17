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

package projects_v1

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

// NewListBookmarkedProjectsParams creates a new ListBookmarkedProjectsParams object
// with the default values initialized.
func NewListBookmarkedProjectsParams() *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBookmarkedProjectsParamsWithTimeout creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBookmarkedProjectsParamsWithTimeout(timeout time.Duration) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		timeout: timeout,
	}
}

// NewListBookmarkedProjectsParamsWithContext creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBookmarkedProjectsParamsWithContext(ctx context.Context) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{

		Context: ctx,
	}
}

// NewListBookmarkedProjectsParamsWithHTTPClient creates a new ListBookmarkedProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBookmarkedProjectsParamsWithHTTPClient(client *http.Client) *ListBookmarkedProjectsParams {
	var ()
	return &ListBookmarkedProjectsParams{
		HTTPClient: client,
	}
}

/*ListBookmarkedProjectsParams contains all the parameters to send to the API endpoint
for the list bookmarked projects operation typically these are written to a http.Request
*/
type ListBookmarkedProjectsParams struct {

	/*User
	  Owner of the namespace

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithTimeout(timeout time.Duration) *ListBookmarkedProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithContext(ctx context.Context) *ListBookmarkedProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithHTTPClient(client *http.Client) *ListBookmarkedProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) WithUser(user string) *ListBookmarkedProjectsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the list bookmarked projects params
func (o *ListBookmarkedProjectsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *ListBookmarkedProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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