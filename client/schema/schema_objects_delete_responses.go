//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
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

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsDeleteReader is a Reader for the SchemaObjectsDelete structure.
type SchemaObjectsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSchemaObjectsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSchemaObjectsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaObjectsDeleteOK creates a SchemaObjectsDeleteOK with default headers values
func NewSchemaObjectsDeleteOK() *SchemaObjectsDeleteOK {
	return &SchemaObjectsDeleteOK{}
}

/*SchemaObjectsDeleteOK handles this case with default header values.

Removed the Object class from the schema.
*/
type SchemaObjectsDeleteOK struct {
}

func (o *SchemaObjectsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/{className}][%d] schemaObjectsDeleteOK ", 200)
}

func (o *SchemaObjectsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsDeleteBadRequest creates a SchemaObjectsDeleteBadRequest with default headers values
func NewSchemaObjectsDeleteBadRequest() *SchemaObjectsDeleteBadRequest {
	return &SchemaObjectsDeleteBadRequest{}
}

/*SchemaObjectsDeleteBadRequest handles this case with default header values.

Could not delete the Object class.
*/
type SchemaObjectsDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /schema/{className}][%d] schemaObjectsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SchemaObjectsDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsDeleteUnauthorized creates a SchemaObjectsDeleteUnauthorized with default headers values
func NewSchemaObjectsDeleteUnauthorized() *SchemaObjectsDeleteUnauthorized {
	return &SchemaObjectsDeleteUnauthorized{}
}

/*SchemaObjectsDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsDeleteUnauthorized struct {
}

func (o *SchemaObjectsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/{className}][%d] schemaObjectsDeleteUnauthorized ", 401)
}

func (o *SchemaObjectsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsDeleteForbidden creates a SchemaObjectsDeleteForbidden with default headers values
func NewSchemaObjectsDeleteForbidden() *SchemaObjectsDeleteForbidden {
	return &SchemaObjectsDeleteForbidden{}
}

/*SchemaObjectsDeleteForbidden handles this case with default header values.

Forbidden
*/
type SchemaObjectsDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schema/{className}][%d] schemaObjectsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsDeleteForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsDeleteInternalServerError creates a SchemaObjectsDeleteInternalServerError with default headers values
func NewSchemaObjectsDeleteInternalServerError() *SchemaObjectsDeleteInternalServerError {
	return &SchemaObjectsDeleteInternalServerError{}
}

/*SchemaObjectsDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schema/{className}][%d] schemaObjectsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
