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

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// NewStopRunsParams creates a new StopRunsParams object
// with the default values initialized.
func NewStopRunsParams() *StopRunsParams {
	var ()
	return &StopRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopRunsParamsWithTimeout creates a new StopRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopRunsParamsWithTimeout(timeout time.Duration) *StopRunsParams {
	var ()
	return &StopRunsParams{

		timeout: timeout,
	}
}

// NewStopRunsParamsWithContext creates a new StopRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopRunsParamsWithContext(ctx context.Context) *StopRunsParams {
	var ()
	return &StopRunsParams{

		Context: ctx,
	}
}

// NewStopRunsParamsWithHTTPClient creates a new StopRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopRunsParamsWithHTTPClient(client *http.Client) *StopRunsParams {
	var ()
	return &StopRunsParams{
		HTTPClient: client,
	}
}

/*StopRunsParams contains all the parameters to send to the API endpoint
for the stop runs operation typically these are written to a http.Request
*/
type StopRunsParams struct {

	/*Body
	  Uuids of the entities

	*/
	Body *service_model.V1Uuids
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

// WithTimeout adds the timeout to the stop runs params
func (o *StopRunsParams) WithTimeout(timeout time.Duration) *StopRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop runs params
func (o *StopRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop runs params
func (o *StopRunsParams) WithContext(ctx context.Context) *StopRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop runs params
func (o *StopRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop runs params
func (o *StopRunsParams) WithHTTPClient(client *http.Client) *StopRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop runs params
func (o *StopRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stop runs params
func (o *StopRunsParams) WithBody(body *service_model.V1Uuids) *StopRunsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stop runs params
func (o *StopRunsParams) SetBody(body *service_model.V1Uuids) {
	o.Body = body
}

// WithOwner adds the owner to the stop runs params
func (o *StopRunsParams) WithOwner(owner string) *StopRunsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the stop runs params
func (o *StopRunsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the stop runs params
func (o *StopRunsParams) WithProject(project string) *StopRunsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the stop runs params
func (o *StopRunsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *StopRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
