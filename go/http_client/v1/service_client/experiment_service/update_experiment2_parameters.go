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

package experiment_service

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

// NewUpdateExperiment2Params creates a new UpdateExperiment2Params object
// with the default values initialized.
func NewUpdateExperiment2Params() *UpdateExperiment2Params {
	var ()
	return &UpdateExperiment2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateExperiment2ParamsWithTimeout creates a new UpdateExperiment2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateExperiment2ParamsWithTimeout(timeout time.Duration) *UpdateExperiment2Params {
	var ()
	return &UpdateExperiment2Params{

		timeout: timeout,
	}
}

// NewUpdateExperiment2ParamsWithContext creates a new UpdateExperiment2Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateExperiment2ParamsWithContext(ctx context.Context) *UpdateExperiment2Params {
	var ()
	return &UpdateExperiment2Params{

		Context: ctx,
	}
}

// NewUpdateExperiment2ParamsWithHTTPClient creates a new UpdateExperiment2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateExperiment2ParamsWithHTTPClient(client *http.Client) *UpdateExperiment2Params {
	var ()
	return &UpdateExperiment2Params{
		HTTPClient: client,
	}
}

/*UpdateExperiment2Params contains all the parameters to send to the API endpoint
for the update experiment2 operation typically these are written to a http.Request
*/
type UpdateExperiment2Params struct {

	/*Body*/
	Body *service_model.V1ExperimentBodyRequest
	/*ExperimentID
	  Unique integer identifier

	*/
	ExperimentID string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the experiement will be assigned

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update experiment2 params
func (o *UpdateExperiment2Params) WithTimeout(timeout time.Duration) *UpdateExperiment2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update experiment2 params
func (o *UpdateExperiment2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update experiment2 params
func (o *UpdateExperiment2Params) WithContext(ctx context.Context) *UpdateExperiment2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update experiment2 params
func (o *UpdateExperiment2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update experiment2 params
func (o *UpdateExperiment2Params) WithHTTPClient(client *http.Client) *UpdateExperiment2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update experiment2 params
func (o *UpdateExperiment2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update experiment2 params
func (o *UpdateExperiment2Params) WithBody(body *service_model.V1ExperimentBodyRequest) *UpdateExperiment2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update experiment2 params
func (o *UpdateExperiment2Params) SetBody(body *service_model.V1ExperimentBodyRequest) {
	o.Body = body
}

// WithExperimentID adds the experimentID to the update experiment2 params
func (o *UpdateExperiment2Params) WithExperimentID(experimentID string) *UpdateExperiment2Params {
	o.SetExperimentID(experimentID)
	return o
}

// SetExperimentID adds the experimentId to the update experiment2 params
func (o *UpdateExperiment2Params) SetExperimentID(experimentID string) {
	o.ExperimentID = experimentID
}

// WithOwner adds the owner to the update experiment2 params
func (o *UpdateExperiment2Params) WithOwner(owner string) *UpdateExperiment2Params {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update experiment2 params
func (o *UpdateExperiment2Params) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the update experiment2 params
func (o *UpdateExperiment2Params) WithProject(project string) *UpdateExperiment2Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update experiment2 params
func (o *UpdateExperiment2Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateExperiment2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param experiment.id
	if err := r.SetPathParam("experiment.id", o.ExperimentID); err != nil {
		return err
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