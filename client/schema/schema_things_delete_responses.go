//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
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

// SchemaThingsDeleteReader is a Reader for the SchemaThingsDelete structure.
type SchemaThingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaThingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaThingsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSchemaThingsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSchemaThingsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaThingsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaThingsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaThingsDeleteOK creates a SchemaThingsDeleteOK with default headers values
func NewSchemaThingsDeleteOK() *SchemaThingsDeleteOK {
	return &SchemaThingsDeleteOK{}
}

/*SchemaThingsDeleteOK handles this case with default header values.

Removed the Thing class from the schema.
*/
type SchemaThingsDeleteOK struct {
}

func (o *SchemaThingsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] schemaThingsDeleteOK ", 200)
}

func (o *SchemaThingsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewSchemaThingsDeleteBadRequest creates a SchemaThingsDeleteBadRequest with default headers values
func NewSchemaThingsDeleteBadRequest() *SchemaThingsDeleteBadRequest {
	return &SchemaThingsDeleteBadRequest{}
}

/*SchemaThingsDeleteBadRequest handles this case with default header values.

Could not delete the Thing class.
*/
type SchemaThingsDeleteBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] schemaThingsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *SchemaThingsDeleteBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsDeleteUnauthorized creates a SchemaThingsDeleteUnauthorized with default headers values
func NewSchemaThingsDeleteUnauthorized() *SchemaThingsDeleteUnauthorized {
	return &SchemaThingsDeleteUnauthorized{}
}

/*SchemaThingsDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaThingsDeleteUnauthorized struct {
}

func (o *SchemaThingsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] schemaThingsDeleteUnauthorized ", 401)
}

func (o *SchemaThingsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewSchemaThingsDeleteForbidden creates a SchemaThingsDeleteForbidden with default headers values
func NewSchemaThingsDeleteForbidden() *SchemaThingsDeleteForbidden {
	return &SchemaThingsDeleteForbidden{}
}

/*SchemaThingsDeleteForbidden handles this case with default header values.

Forbidden
*/
type SchemaThingsDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] schemaThingsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *SchemaThingsDeleteForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsDeleteInternalServerError creates a SchemaThingsDeleteInternalServerError with default headers values
func NewSchemaThingsDeleteInternalServerError() *SchemaThingsDeleteInternalServerError {
	return &SchemaThingsDeleteInternalServerError{}
}

/*SchemaThingsDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaThingsDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}][%d] schemaThingsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaThingsDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
