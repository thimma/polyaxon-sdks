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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UnBookmarkBuildReader is a Reader for the UnBookmarkBuild structure.
type UnBookmarkBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnBookmarkBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnBookmarkBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUnBookmarkBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnBookmarkBuildOK creates a UnBookmarkBuildOK with default headers values
func NewUnBookmarkBuildOK() *UnBookmarkBuildOK {
	return &UnBookmarkBuildOK{}
}

/*UnBookmarkBuildOK handles this case with default header values.

A successful response.
*/
type UnBookmarkBuildOK struct {
	Payload interface{}
}

func (o *UnBookmarkBuildOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/builds/{id}/unbookmark][%d] unBookmarkBuildOK  %+v", 200, o.Payload)
}

func (o *UnBookmarkBuildOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UnBookmarkBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnBookmarkBuildNotFound creates a UnBookmarkBuildNotFound with default headers values
func NewUnBookmarkBuildNotFound() *UnBookmarkBuildNotFound {
	return &UnBookmarkBuildNotFound{}
}

/*UnBookmarkBuildNotFound handles this case with default header values.

Resource does not exist.
*/
type UnBookmarkBuildNotFound struct {
	Payload string
}

func (o *UnBookmarkBuildNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/builds/{id}/unbookmark][%d] unBookmarkBuildNotFound  %+v", 404, o.Payload)
}

func (o *UnBookmarkBuildNotFound) GetPayload() string {
	return o.Payload
}

func (o *UnBookmarkBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}