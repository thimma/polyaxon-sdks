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

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteJobReader is a Reader for the DeleteJob structure.
type DeleteJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteJobOK creates a DeleteJobOK with default headers values
func NewDeleteJobOK() *DeleteJobOK {
	return &DeleteJobOK{}
}

/*DeleteJobOK handles this case with default header values.

A successful response.
*/
type DeleteJobOK struct {
	Payload interface{}
}

func (o *DeleteJobOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/jobs/{id}][%d] deleteJobOK  %+v", 200, o.Payload)
}

func (o *DeleteJobOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobNotFound creates a DeleteJobNotFound with default headers values
func NewDeleteJobNotFound() *DeleteJobNotFound {
	return &DeleteJobNotFound{}
}

/*DeleteJobNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteJobNotFound struct {
	Payload string
}

func (o *DeleteJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/jobs/{id}][%d] deleteJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteJobNotFound) GetPayload() string {
	return o.Payload
}

func (o *DeleteJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}