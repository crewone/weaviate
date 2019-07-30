//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaActionsCreateReader is a Reader for the WeaviateSchemaActionsCreate structure.
type WeaviateSchemaActionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaActionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaActionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaActionsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaActionsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaActionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaActionsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaActionsCreateOK creates a WeaviateSchemaActionsCreateOK with default headers values
func NewWeaviateSchemaActionsCreateOK() *WeaviateSchemaActionsCreateOK {
	return &WeaviateSchemaActionsCreateOK{}
}

/*WeaviateSchemaActionsCreateOK handles this case with default header values.

Added the new Action class to the ontology.
*/
type WeaviateSchemaActionsCreateOK struct {
	Payload *models.Class
}

func (o *WeaviateSchemaActionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /schema/actions][%d] weaviateSchemaActionsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateSchemaActionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Class)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsCreateUnauthorized creates a WeaviateSchemaActionsCreateUnauthorized with default headers values
func NewWeaviateSchemaActionsCreateUnauthorized() *WeaviateSchemaActionsCreateUnauthorized {
	return &WeaviateSchemaActionsCreateUnauthorized{}
}

/*WeaviateSchemaActionsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaActionsCreateUnauthorized struct {
}

func (o *WeaviateSchemaActionsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/actions][%d] weaviateSchemaActionsCreateUnauthorized ", 401)
}

func (o *WeaviateSchemaActionsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsCreateForbidden creates a WeaviateSchemaActionsCreateForbidden with default headers values
func NewWeaviateSchemaActionsCreateForbidden() *WeaviateSchemaActionsCreateForbidden {
	return &WeaviateSchemaActionsCreateForbidden{}
}

/*WeaviateSchemaActionsCreateForbidden handles this case with default header values.

Forbidden
*/
type WeaviateSchemaActionsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/actions][%d] weaviateSchemaActionsCreateForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateSchemaActionsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsCreateUnprocessableEntity creates a WeaviateSchemaActionsCreateUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsCreateUnprocessableEntity() *WeaviateSchemaActionsCreateUnprocessableEntity {
	return &WeaviateSchemaActionsCreateUnprocessableEntity{}
}

/*WeaviateSchemaActionsCreateUnprocessableEntity handles this case with default header values.

Invalid Action class
*/
type WeaviateSchemaActionsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/actions][%d] weaviateSchemaActionsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaActionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsCreateInternalServerError creates a WeaviateSchemaActionsCreateInternalServerError with default headers values
func NewWeaviateSchemaActionsCreateInternalServerError() *WeaviateSchemaActionsCreateInternalServerError {
	return &WeaviateSchemaActionsCreateInternalServerError{}
}

/*WeaviateSchemaActionsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaActionsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/actions][%d] weaviateSchemaActionsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaActionsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
