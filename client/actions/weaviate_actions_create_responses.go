//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
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

// WeaviateActionsCreateReader is a Reader for the WeaviateActionsCreate structure.
type WeaviateActionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsCreateOK creates a WeaviateActionsCreateOK with default headers values
func NewWeaviateActionsCreateOK() *WeaviateActionsCreateOK {
	return &WeaviateActionsCreateOK{}
}

/*WeaviateActionsCreateOK handles this case with default header values.

Action created.
*/
type WeaviateActionsCreateOK struct {
	Payload *models.Action
}

func (o *WeaviateActionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateActionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateUnauthorized creates a WeaviateActionsCreateUnauthorized with default headers values
func NewWeaviateActionsCreateUnauthorized() *WeaviateActionsCreateUnauthorized {
	return &WeaviateActionsCreateUnauthorized{}
}

/*WeaviateActionsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsCreateUnauthorized struct {
}

func (o *WeaviateActionsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateUnauthorized ", 401)
}

func (o *WeaviateActionsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsCreateForbidden creates a WeaviateActionsCreateForbidden with default headers values
func NewWeaviateActionsCreateForbidden() *WeaviateActionsCreateForbidden {
	return &WeaviateActionsCreateForbidden{}
}

/*WeaviateActionsCreateForbidden handles this case with default header values.

Forbidden
*/
type WeaviateActionsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateActionsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateUnprocessableEntity creates a WeaviateActionsCreateUnprocessableEntity with default headers values
func NewWeaviateActionsCreateUnprocessableEntity() *WeaviateActionsCreateUnprocessableEntity {
	return &WeaviateActionsCreateUnprocessableEntity{}
}

/*WeaviateActionsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateActionsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateInternalServerError creates a WeaviateActionsCreateInternalServerError with default headers values
func NewWeaviateActionsCreateInternalServerError() *WeaviateActionsCreateInternalServerError {
	return &WeaviateActionsCreateInternalServerError{}
}

/*WeaviateActionsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
