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

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// NewStopBuildsParams creates a new StopBuildsParams object
// with the default values initialized.
func NewStopBuildsParams() *StopBuildsParams {
	var ()
	return &StopBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopBuildsParamsWithTimeout creates a new StopBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopBuildsParamsWithTimeout(timeout time.Duration) *StopBuildsParams {
	var ()
	return &StopBuildsParams{

		timeout: timeout,
	}
}

// NewStopBuildsParamsWithContext creates a new StopBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopBuildsParamsWithContext(ctx context.Context) *StopBuildsParams {
	var ()
	return &StopBuildsParams{

		Context: ctx,
	}
}

// NewStopBuildsParamsWithHTTPClient creates a new StopBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopBuildsParamsWithHTTPClient(client *http.Client) *StopBuildsParams {
	var ()
	return &StopBuildsParams{
		HTTPClient: client,
	}
}

/*StopBuildsParams contains all the parameters to send to the API endpoint
for the stop builds operation typically these are written to a http.Request
*/
type StopBuildsParams struct {

	/*Body*/
	Body *service_model.V1ProjectBodyRequest
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop builds params
func (o *StopBuildsParams) WithTimeout(timeout time.Duration) *StopBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop builds params
func (o *StopBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop builds params
func (o *StopBuildsParams) WithContext(ctx context.Context) *StopBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop builds params
func (o *StopBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop builds params
func (o *StopBuildsParams) WithHTTPClient(client *http.Client) *StopBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop builds params
func (o *StopBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stop builds params
func (o *StopBuildsParams) WithBody(body *service_model.V1ProjectBodyRequest) *StopBuildsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop builds params
func (o *StopBuildsParams) SetBody(body *service_model.V1ProjectBodyRequest) {
	o.Body = body
}

// WithOwner adds the owner to the stop builds params
func (o *StopBuildsParams) WithOwner(owner string) *StopBuildsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the stop builds params
func (o *StopBuildsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the stop builds params
func (o *StopBuildsParams) WithProject(project string) *StopBuildsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the stop builds params
func (o *StopBuildsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *StopBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}