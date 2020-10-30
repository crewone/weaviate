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

// SchemaThingsPropertiesAddReader is a Reader for the SchemaThingsPropertiesAdd structure.
type SchemaThingsPropertiesAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaThingsPropertiesAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaThingsPropertiesAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaThingsPropertiesAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaThingsPropertiesAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSchemaThingsPropertiesAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaThingsPropertiesAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaThingsPropertiesAddOK creates a SchemaThingsPropertiesAddOK with default headers values
func NewSchemaThingsPropertiesAddOK() *SchemaThingsPropertiesAddOK {
	return &SchemaThingsPropertiesAddOK{}
}

/*SchemaThingsPropertiesAddOK handles this case with default header values.

Added the property.
*/
type SchemaThingsPropertiesAddOK struct {
	Payload *models.Property
}

func (o *SchemaThingsPropertiesAddOK) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] schemaThingsPropertiesAddOK  %+v", 200, o.Payload)
}

func (o *SchemaThingsPropertiesAddOK) GetPayload() *models.Property {
	return o.Payload
}

func (o *SchemaThingsPropertiesAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsPropertiesAddUnauthorized creates a SchemaThingsPropertiesAddUnauthorized with default headers values
func NewSchemaThingsPropertiesAddUnauthorized() *SchemaThingsPropertiesAddUnauthorized {
	return &SchemaThingsPropertiesAddUnauthorized{}
}

/*SchemaThingsPropertiesAddUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaThingsPropertiesAddUnauthorized struct {
}

func (o *SchemaThingsPropertiesAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] schemaThingsPropertiesAddUnauthorized ", 401)
}

func (o *SchemaThingsPropertiesAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewSchemaThingsPropertiesAddForbidden creates a SchemaThingsPropertiesAddForbidden with default headers values
func NewSchemaThingsPropertiesAddForbidden() *SchemaThingsPropertiesAddForbidden {
	return &SchemaThingsPropertiesAddForbidden{}
}

/*SchemaThingsPropertiesAddForbidden handles this case with default header values.

Forbidden
*/
type SchemaThingsPropertiesAddForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesAddForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] schemaThingsPropertiesAddForbidden  %+v", 403, o.Payload)
}

func (o *SchemaThingsPropertiesAddForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsPropertiesAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsPropertiesAddUnprocessableEntity creates a SchemaThingsPropertiesAddUnprocessableEntity with default headers values
func NewSchemaThingsPropertiesAddUnprocessableEntity() *SchemaThingsPropertiesAddUnprocessableEntity {
	return &SchemaThingsPropertiesAddUnprocessableEntity{}
}

/*SchemaThingsPropertiesAddUnprocessableEntity handles this case with default header values.

Invalid property.
*/
type SchemaThingsPropertiesAddUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] schemaThingsPropertiesAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaThingsPropertiesAddUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsPropertiesAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaThingsPropertiesAddInternalServerError creates a SchemaThingsPropertiesAddInternalServerError with default headers values
func NewSchemaThingsPropertiesAddInternalServerError() *SchemaThingsPropertiesAddInternalServerError {
	return &SchemaThingsPropertiesAddInternalServerError{}
}

/*SchemaThingsPropertiesAddInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaThingsPropertiesAddInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaThingsPropertiesAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/things/{className}/properties][%d] schemaThingsPropertiesAddInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaThingsPropertiesAddInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaThingsPropertiesAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
