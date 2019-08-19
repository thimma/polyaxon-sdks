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

// NewDeleteBuildParams creates a new DeleteBuildParams object
// with the default values initialized.
func NewDeleteBuildParams() *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildParamsWithTimeout creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBuildParamsWithTimeout(timeout time.Duration) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		timeout: timeout,
	}
}

// NewDeleteBuildParamsWithContext creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBuildParamsWithContext(ctx context.Context) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{

		Context: ctx,
	}
}

// NewDeleteBuildParamsWithHTTPClient creates a new DeleteBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBuildParamsWithHTTPClient(client *http.Client) *DeleteBuildParams {
	var ()
	return &DeleteBuildParams{
		HTTPClient: client,
	}
}

/*DeleteBuildParams contains all the parameters to send to the API endpoint
for the delete build operation typically these are written to a http.Request
*/
type DeleteBuildParams struct {

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

// WithTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) WithTimeout(timeout time.Duration) *DeleteBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete build params
func (o *DeleteBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete build params
func (o *DeleteBuildParams) WithContext(ctx context.Context) *DeleteBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete build params
func (o *DeleteBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) WithHTTPClient(client *http.Client) *DeleteBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete build params
func (o *DeleteBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete build params
func (o *DeleteBuildParams) WithID(id string) *DeleteBuildParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete build params
func (o *DeleteBuildParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the delete build params
func (o *DeleteBuildParams) WithOwner(owner string) *DeleteBuildParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete build params
func (o *DeleteBuildParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the delete build params
func (o *DeleteBuildParams) WithProject(project string) *DeleteBuildParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete build params
func (o *DeleteBuildParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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