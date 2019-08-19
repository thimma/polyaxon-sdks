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

// BookmarkJobReader is a Reader for the BookmarkJob structure.
type BookmarkJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewBookmarkJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBookmarkJobOK creates a BookmarkJobOK with default headers values
func NewBookmarkJobOK() *BookmarkJobOK {
	return &BookmarkJobOK{}
}

/*BookmarkJobOK handles this case with default header values.

A successful response.
*/
type BookmarkJobOK struct {
	Payload interface{}
}

func (o *BookmarkJobOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/bookmark][%d] bookmarkJobOK  %+v", 200, o.Payload)
}

func (o *BookmarkJobOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkJobNotFound creates a BookmarkJobNotFound with default headers values
func NewBookmarkJobNotFound() *BookmarkJobNotFound {
	return &BookmarkJobNotFound{}
}

/*BookmarkJobNotFound handles this case with default header values.

Resource does not exist.
*/
type BookmarkJobNotFound struct {
	Payload string
}

func (o *BookmarkJobNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/jobs/{id}/bookmark][%d] bookmarkJobNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkJobNotFound) GetPayload() string {
	return o.Payload
}

func (o *BookmarkJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}