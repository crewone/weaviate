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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ActionsGetReader is a Reader for the ActionsGet structure.
type ActionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewActionsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewActionsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActionsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActionsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActionsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsGetOK creates a ActionsGetOK with default headers values
func NewActionsGetOK() *ActionsGetOK {
	return &ActionsGetOK{}
}

/*ActionsGetOK handles this case with default header values.

Successful response.
*/
type ActionsGetOK struct {
	Payload *models.Action
}

func (o *ActionsGetOK) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetOK  %+v", 200, o.Payload)
}

func (o *ActionsGetOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *ActionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsGetBadRequest creates a ActionsGetBadRequest with default headers values
func NewActionsGetBadRequest() *ActionsGetBadRequest {
	return &ActionsGetBadRequest{}
}

/*ActionsGetBadRequest handles this case with default header values.

Malformed request.
*/
type ActionsGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ActionsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ActionsGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsGetUnauthorized creates a ActionsGetUnauthorized with default headers values
func NewActionsGetUnauthorized() *ActionsGetUnauthorized {
	return &ActionsGetUnauthorized{}
}

/*ActionsGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsGetUnauthorized struct {
}

func (o *ActionsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetUnauthorized ", 401)
}

func (o *ActionsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsGetForbidden creates a ActionsGetForbidden with default headers values
func NewActionsGetForbidden() *ActionsGetForbidden {
	return &ActionsGetForbidden{}
}

/*ActionsGetForbidden handles this case with default header values.

Forbidden
*/
type ActionsGetForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetForbidden  %+v", 403, o.Payload)
}

func (o *ActionsGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsGetNotFound creates a ActionsGetNotFound with default headers values
func NewActionsGetNotFound() *ActionsGetNotFound {
	return &ActionsGetNotFound{}
}

/*ActionsGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ActionsGetNotFound struct {
}

func (o *ActionsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetNotFound ", 404)
}

func (o *ActionsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsGetInternalServerError creates a ActionsGetInternalServerError with default headers values
func NewActionsGetInternalServerError() *ActionsGetInternalServerError {
	return &ActionsGetInternalServerError{}
}

/*ActionsGetInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /actions/{id}][%d] actionsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
