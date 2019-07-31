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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsCreateReader is a Reader for the ActionsCreate structure.
type ActionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewActionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewActionsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewActionsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewActionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewActionsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsCreateOK creates a ActionsCreateOK with default headers values
func NewActionsCreateOK() *ActionsCreateOK {
	return &ActionsCreateOK{}
}

/*ActionsCreateOK handles this case with default header values.

Action created.
*/
type ActionsCreateOK struct {
	Payload *models.Action
}

func (o *ActionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /actions][%d] actionsCreateOK  %+v", 200, o.Payload)
}

func (o *ActionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsCreateUnauthorized creates a ActionsCreateUnauthorized with default headers values
func NewActionsCreateUnauthorized() *ActionsCreateUnauthorized {
	return &ActionsCreateUnauthorized{}
}

/*ActionsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsCreateUnauthorized struct {
}

func (o *ActionsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions][%d] actionsCreateUnauthorized ", 401)
}

func (o *ActionsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsCreateForbidden creates a ActionsCreateForbidden with default headers values
func NewActionsCreateForbidden() *ActionsCreateForbidden {
	return &ActionsCreateForbidden{}
}

/*ActionsCreateForbidden handles this case with default header values.

Forbidden
*/
type ActionsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions][%d] actionsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ActionsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsCreateUnprocessableEntity creates a ActionsCreateUnprocessableEntity with default headers values
func NewActionsCreateUnprocessableEntity() *ActionsCreateUnprocessableEntity {
	return &ActionsCreateUnprocessableEntity{}
}

/*ActionsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ActionsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ActionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions][%d] actionsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ActionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsCreateInternalServerError creates a ActionsCreateInternalServerError with default headers values
func NewActionsCreateInternalServerError() *ActionsCreateInternalServerError {
	return &ActionsCreateInternalServerError{}
}

/*ActionsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions][%d] actionsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
