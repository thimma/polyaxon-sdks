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

// NewCreateBuildStatusParams creates a new CreateBuildStatusParams object
// with the default values initialized.
func NewCreateBuildStatusParams() *CreateBuildStatusParams {
	var ()
	return &CreateBuildStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBuildStatusParamsWithTimeout creates a new CreateBuildStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBuildStatusParamsWithTimeout(timeout time.Duration) *CreateBuildStatusParams {
	var ()
	return &CreateBuildStatusParams{

		timeout: timeout,
	}
}

// NewCreateBuildStatusParamsWithContext creates a new CreateBuildStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBuildStatusParamsWithContext(ctx context.Context) *CreateBuildStatusParams {
	var ()
	return &CreateBuildStatusParams{

		Context: ctx,
	}
}

// NewCreateBuildStatusParamsWithHTTPClient creates a new CreateBuildStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBuildStatusParamsWithHTTPClient(client *http.Client) *CreateBuildStatusParams {
	var ()
	return &CreateBuildStatusParams{
		HTTPClient: client,
	}
}

/*CreateBuildStatusParams contains all the parameters to send to the API endpoint
for the create build status operation typically these are written to a http.Request
*/
type CreateBuildStatusParams struct {

	/*Body*/
	Body *service_model.V1OwnedEntityIDRequest
	/*ID
	  Unique integer identifier of the entity

	*/
	ID string
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

// WithTimeout adds the timeout to the create build status params
func (o *CreateBuildStatusParams) WithTimeout(timeout time.Duration) *CreateBuildStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create build status params
func (o *CreateBuildStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create build status params
func (o *CreateBuildStatusParams) WithContext(ctx context.Context) *CreateBuildStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create build status params
func (o *CreateBuildStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create build status params
func (o *CreateBuildStatusParams) WithHTTPClient(client *http.Client) *CreateBuildStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create build status params
func (o *CreateBuildStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create build status params
func (o *CreateBuildStatusParams) WithBody(body *service_model.V1OwnedEntityIDRequest) *CreateBuildStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create build status params
func (o *CreateBuildStatusParams) SetBody(body *service_model.V1OwnedEntityIDRequest) {
	o.Body = body
}

// WithID adds the id to the create build status params
func (o *CreateBuildStatusParams) WithID(id string) *CreateBuildStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create build status params
func (o *CreateBuildStatusParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the create build status params
func (o *CreateBuildStatusParams) WithOwner(owner string) *CreateBuildStatusParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create build status params
func (o *CreateBuildStatusParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the create build status params
func (o *CreateBuildStatusParams) WithProject(project string) *CreateBuildStatusParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create build status params
func (o *CreateBuildStatusParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBuildStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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