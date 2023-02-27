//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// SchemaObjectsUpdateReader is a Reader for the SchemaObjectsUpdate structure.
type SchemaObjectsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaObjectsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSchemaObjectsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSchemaObjectsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSchemaObjectsUpdateOK creates a SchemaObjectsUpdateOK with default headers values
func NewSchemaObjectsUpdateOK() *SchemaObjectsUpdateOK {
	return &SchemaObjectsUpdateOK{}
}

/*
SchemaObjectsUpdateOK describes a response with status code 200, with default header values.

Class was updated successfully
*/
type SchemaObjectsUpdateOK struct {
	Payload *models.Class
}

// IsSuccess returns true when this schema objects update o k response has a 2xx status code
func (o *SchemaObjectsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schema objects update o k response has a 3xx status code
func (o *SchemaObjectsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update o k response has a 4xx status code
func (o *SchemaObjectsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects update o k response has a 5xx status code
func (o *SchemaObjectsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects update o k response a status code equal to that given
func (o *SchemaObjectsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schema objects update o k response
func (o *SchemaObjectsUpdateOK) Code() int {
	return 200
}

func (o *SchemaObjectsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsUpdateOK) GetPayload() *models.Class {
	return o.Payload
}

func (o *SchemaObjectsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Class)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsUpdateUnauthorized creates a SchemaObjectsUpdateUnauthorized with default headers values
func NewSchemaObjectsUpdateUnauthorized() *SchemaObjectsUpdateUnauthorized {
	return &SchemaObjectsUpdateUnauthorized{}
}

/*
SchemaObjectsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsUpdateUnauthorized struct {
}

// IsSuccess returns true when this schema objects update unauthorized response has a 2xx status code
func (o *SchemaObjectsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects update unauthorized response has a 3xx status code
func (o *SchemaObjectsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update unauthorized response has a 4xx status code
func (o *SchemaObjectsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects update unauthorized response has a 5xx status code
func (o *SchemaObjectsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects update unauthorized response a status code equal to that given
func (o *SchemaObjectsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the schema objects update unauthorized response
func (o *SchemaObjectsUpdateUnauthorized) Code() int {
	return 401
}

func (o *SchemaObjectsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateUnauthorized ", 401)
}

func (o *SchemaObjectsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateUnauthorized ", 401)
}

func (o *SchemaObjectsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsUpdateForbidden creates a SchemaObjectsUpdateForbidden with default headers values
func NewSchemaObjectsUpdateForbidden() *SchemaObjectsUpdateForbidden {
	return &SchemaObjectsUpdateForbidden{}
}

/*
SchemaObjectsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SchemaObjectsUpdateForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects update forbidden response has a 2xx status code
func (o *SchemaObjectsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects update forbidden response has a 3xx status code
func (o *SchemaObjectsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update forbidden response has a 4xx status code
func (o *SchemaObjectsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects update forbidden response has a 5xx status code
func (o *SchemaObjectsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects update forbidden response a status code equal to that given
func (o *SchemaObjectsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the schema objects update forbidden response
func (o *SchemaObjectsUpdateForbidden) Code() int {
	return 403
}

func (o *SchemaObjectsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsUpdateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsUpdateNotFound creates a SchemaObjectsUpdateNotFound with default headers values
func NewSchemaObjectsUpdateNotFound() *SchemaObjectsUpdateNotFound {
	return &SchemaObjectsUpdateNotFound{}
}

/*
SchemaObjectsUpdateNotFound describes a response with status code 404, with default header values.

Class to be updated does not exist
*/
type SchemaObjectsUpdateNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects update not found response has a 2xx status code
func (o *SchemaObjectsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects update not found response has a 3xx status code
func (o *SchemaObjectsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update not found response has a 4xx status code
func (o *SchemaObjectsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects update not found response has a 5xx status code
func (o *SchemaObjectsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects update not found response a status code equal to that given
func (o *SchemaObjectsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the schema objects update not found response
func (o *SchemaObjectsUpdateNotFound) Code() int {
	return 404
}

func (o *SchemaObjectsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SchemaObjectsUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SchemaObjectsUpdateNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsUpdateUnprocessableEntity creates a SchemaObjectsUpdateUnprocessableEntity with default headers values
func NewSchemaObjectsUpdateUnprocessableEntity() *SchemaObjectsUpdateUnprocessableEntity {
	return &SchemaObjectsUpdateUnprocessableEntity{}
}

/*
SchemaObjectsUpdateUnprocessableEntity describes a response with status code 422, with default header values.

Invalid update attempt
*/
type SchemaObjectsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects update unprocessable entity response has a 2xx status code
func (o *SchemaObjectsUpdateUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects update unprocessable entity response has a 3xx status code
func (o *SchemaObjectsUpdateUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update unprocessable entity response has a 4xx status code
func (o *SchemaObjectsUpdateUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this schema objects update unprocessable entity response has a 5xx status code
func (o *SchemaObjectsUpdateUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this schema objects update unprocessable entity response a status code equal to that given
func (o *SchemaObjectsUpdateUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the schema objects update unprocessable entity response
func (o *SchemaObjectsUpdateUnprocessableEntity) Code() int {
	return 422
}

func (o *SchemaObjectsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaObjectsUpdateUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaObjectsUpdateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsUpdateInternalServerError creates a SchemaObjectsUpdateInternalServerError with default headers values
func NewSchemaObjectsUpdateInternalServerError() *SchemaObjectsUpdateInternalServerError {
	return &SchemaObjectsUpdateInternalServerError{}
}

/*
SchemaObjectsUpdateInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this schema objects update internal server error response has a 2xx status code
func (o *SchemaObjectsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schema objects update internal server error response has a 3xx status code
func (o *SchemaObjectsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schema objects update internal server error response has a 4xx status code
func (o *SchemaObjectsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this schema objects update internal server error response has a 5xx status code
func (o *SchemaObjectsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this schema objects update internal server error response a status code equal to that given
func (o *SchemaObjectsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the schema objects update internal server error response
func (o *SchemaObjectsUpdateInternalServerError) Code() int {
	return 500
}

func (o *SchemaObjectsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /schema/{className}][%d] schemaObjectsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
