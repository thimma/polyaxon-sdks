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

package build_service

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

// NewListArchivedBuildsParams creates a new ListArchivedBuildsParams object
// with the default values initialized.
func NewListArchivedBuildsParams() *ListArchivedBuildsParams {
	var ()
	return &ListArchivedBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListArchivedBuildsParamsWithTimeout creates a new ListArchivedBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListArchivedBuildsParamsWithTimeout(timeout time.Duration) *ListArchivedBuildsParams {
	var ()
	return &ListArchivedBuildsParams{

		timeout: timeout,
	}
}

// NewListArchivedBuildsParamsWithContext creates a new ListArchivedBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListArchivedBuildsParamsWithContext(ctx context.Context) *ListArchivedBuildsParams {
	var ()
	return &ListArchivedBuildsParams{

		Context: ctx,
	}
}

// NewListArchivedBuildsParamsWithHTTPClient creates a new ListArchivedBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListArchivedBuildsParamsWithHTTPClient(client *http.Client) *ListArchivedBuildsParams {
	var ()
	return &ListArchivedBuildsParams{
		HTTPClient: client,
	}
}

/*ListArchivedBuildsParams contains all the parameters to send to the API endpoint
for the list archived builds operation typically these are written to a http.Request
*/
type ListArchivedBuildsParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list archived builds params
func (o *ListArchivedBuildsParams) WithTimeout(timeout time.Duration) *ListArchivedBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list archived builds params
func (o *ListArchivedBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list archived builds params
func (o *ListArchivedBuildsParams) WithContext(ctx context.Context) *ListArchivedBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list archived builds params
func (o *ListArchivedBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list archived builds params
func (o *ListArchivedBuildsParams) WithHTTPClient(client *http.Client) *ListArchivedBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list archived builds params
func (o *ListArchivedBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the list archived builds params
func (o *ListArchivedBuildsParams) WithOwner(owner string) *ListArchivedBuildsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list archived builds params
func (o *ListArchivedBuildsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *ListArchivedBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}