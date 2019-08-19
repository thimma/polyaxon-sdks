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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/service_model"
)

// GetExperimentCodeRefReader is a Reader for the GetExperimentCodeRef structure.
type GetExperimentCodeRefReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExperimentCodeRefReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExperimentCodeRefOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetExperimentCodeRefNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExperimentCodeRefOK creates a GetExperimentCodeRefOK with default headers values
func NewGetExperimentCodeRefOK() *GetExperimentCodeRefOK {
	return &GetExperimentCodeRefOK{}
}

/*GetExperimentCodeRefOK handles this case with default header values.

A successful response.
*/
type GetExperimentCodeRefOK struct {
	Payload *service_model.V1CodeReference
}

func (o *GetExperimentCodeRefOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/experiments/{id}/coderef][%d] getExperimentCodeRefOK  %+v", 200, o.Payload)
}

func (o *GetExperimentCodeRefOK) GetPayload() *service_model.V1CodeReference {
	return o.Payload
}

func (o *GetExperimentCodeRefOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1CodeReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExperimentCodeRefNotFound creates a GetExperimentCodeRefNotFound with default headers values
func NewGetExperimentCodeRefNotFound() *GetExperimentCodeRefNotFound {
	return &GetExperimentCodeRefNotFound{}
}

/*GetExperimentCodeRefNotFound handles this case with default header values.

Resource does not exist.
*/
type GetExperimentCodeRefNotFound struct {
	Payload string
}

func (o *GetExperimentCodeRefNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/experiments/{id}/coderef][%d] getExperimentCodeRefNotFound  %+v", 404, o.Payload)
}

func (o *GetExperimentCodeRefNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetExperimentCodeRefNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}