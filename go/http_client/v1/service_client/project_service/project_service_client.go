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

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ArchiveProject stops run
*/
func (a *Client) ArchiveProject(params *ArchiveProjectParams, authInfo runtime.ClientAuthInfoWriter) (*ArchiveProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewArchiveProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ArchiveProject",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &ArchiveProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ArchiveProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ArchiveProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BookmarkProject stops run
*/
func (a *Client) BookmarkProject(params *BookmarkProjectParams, authInfo runtime.ClientAuthInfoWriter) (*BookmarkProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBookmarkProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BookmarkProject",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/bookmark",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &BookmarkProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BookmarkProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BookmarkProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateProject gets run
*/
func (a *Client) CreateProject(params *CreateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateProject",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/projects/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteExperiment deletes runs
*/
func (a *Client) DeleteExperiment(params *DeleteExperimentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteExperimentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExperimentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteExperiment",
		Method:             "DELETE",
		PathPattern:        "/api/v1/{owner}/projecs/{project}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &DeleteExperimentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteExperimentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteExperiment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisableProjectCI resumes run
*/
func (a *Client) DisableProjectCI(params *DisableProjectCIParams, authInfo runtime.ClientAuthInfoWriter) (*DisableProjectCIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableProjectCIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DisableProjectCI",
		Method:             "DELETE",
		PathPattern:        "/api/v1/{owner}/{project}/unbookmark",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &DisableProjectCIReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableProjectCIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DisableProjectCI: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnableProjectCI restarts run
*/
func (a *Client) EnableProjectCI(params *EnableProjectCIParams, authInfo runtime.ClientAuthInfoWriter) (*EnableProjectCIOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableProjectCIParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnableProjectCI",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/ci",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &EnableProjectCIReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableProjectCIOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for EnableProjectCI: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetProject updates run
*/
func (a *Client) GetProject(params *GetProjectParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetProject",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/projects/{project}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListArchivedProjects creates new run
*/
func (a *Client) ListArchivedProjects(params *ListArchivedProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListArchivedProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArchivedProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListArchivedProjects",
		Method:             "GET",
		PathPattern:        "/api/v1/archives/{owner}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &ListArchivedProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListArchivedProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListArchivedProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBookmarkedProjects lists archived runs
*/
func (a *Client) ListBookmarkedProjects(params *ListBookmarkedProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListBookmarkedProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBookmarkedProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListBookmarkedProjects",
		Method:             "GET",
		PathPattern:        "/api/v1/bookmarks/{owner}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &ListBookmarkedProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBookmarkedProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListBookmarkedProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListProjectNames lists bookmarked runs
*/
func (a *Client) ListProjectNames(params *ListProjectNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListProjectNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListProjectNames",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/projects/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &ListProjectNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListProjectNames: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListProjects lists runs
*/
func (a *Client) ListProjects(params *ListProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListProjects",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/projects/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &ListProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchProject deletes run
*/
func (a *Client) PatchProject(params *PatchProjectParams, authInfo runtime.ClientAuthInfoWriter) (*PatchProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchProject",
		Method:             "PATCH",
		PathPattern:        "/api/v1/{owner}/projects/{project}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &PatchProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestoreExperiment stops runs
*/
func (a *Client) RestoreExperiment(params *RestoreExperimentParams, authInfo runtime.ClientAuthInfoWriter) (*RestoreExperimentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreExperimentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RestoreExperiment",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &RestoreExperimentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreExperimentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RestoreExperiment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateProject patches run
*/
func (a *Client) UpdateProject(params *UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateProject",
		Method:             "PUT",
		PathPattern:        "/api/v1/{owner}/projects/{project}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "ws", "wss"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}