/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsPropertiesUpdateReader is a Reader for the WeaviateActionsPropertiesUpdate structure.
type WeaviateActionsPropertiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsPropertiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsPropertiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsPropertiesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsPropertiesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsPropertiesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsPropertiesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsPropertiesUpdateOK creates a WeaviateActionsPropertiesUpdateOK with default headers values
func NewWeaviateActionsPropertiesUpdateOK() *WeaviateActionsPropertiesUpdateOK {
	return &WeaviateActionsPropertiesUpdateOK{}
}

/*WeaviateActionsPropertiesUpdateOK handles this case with default header values.

Successfully replaced all the references.
*/
type WeaviateActionsPropertiesUpdateOK struct {
}

func (o *WeaviateActionsPropertiesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}/references/{propertyName}][%d] weaviateActionsPropertiesUpdateOK ", 200)
}

func (o *WeaviateActionsPropertiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesUpdateUnauthorized creates a WeaviateActionsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateActionsPropertiesUpdateUnauthorized() *WeaviateActionsPropertiesUpdateUnauthorized {
	return &WeaviateActionsPropertiesUpdateUnauthorized{}
}

/*WeaviateActionsPropertiesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsPropertiesUpdateUnauthorized struct {
}

func (o *WeaviateActionsPropertiesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}/references/{propertyName}][%d] weaviateActionsPropertiesUpdateUnauthorized ", 401)
}

func (o *WeaviateActionsPropertiesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesUpdateForbidden creates a WeaviateActionsPropertiesUpdateForbidden with default headers values
func NewWeaviateActionsPropertiesUpdateForbidden() *WeaviateActionsPropertiesUpdateForbidden {
	return &WeaviateActionsPropertiesUpdateForbidden{}
}

/*WeaviateActionsPropertiesUpdateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateActionsPropertiesUpdateForbidden struct {
}

func (o *WeaviateActionsPropertiesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}/references/{propertyName}][%d] weaviateActionsPropertiesUpdateForbidden ", 403)
}

func (o *WeaviateActionsPropertiesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsPropertiesUpdateUnprocessableEntity creates a WeaviateActionsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateActionsPropertiesUpdateUnprocessableEntity() *WeaviateActionsPropertiesUpdateUnprocessableEntity {
	return &WeaviateActionsPropertiesUpdateUnprocessableEntity{}
}

/*WeaviateActionsPropertiesUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type WeaviateActionsPropertiesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsPropertiesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}/references/{propertyName}][%d] weaviateActionsPropertiesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsPropertiesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsPropertiesUpdateInternalServerError creates a WeaviateActionsPropertiesUpdateInternalServerError with default headers values
func NewWeaviateActionsPropertiesUpdateInternalServerError() *WeaviateActionsPropertiesUpdateInternalServerError {
	return &WeaviateActionsPropertiesUpdateInternalServerError{}
}

/*WeaviateActionsPropertiesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsPropertiesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsPropertiesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}/references/{propertyName}][%d] weaviateActionsPropertiesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsPropertiesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
