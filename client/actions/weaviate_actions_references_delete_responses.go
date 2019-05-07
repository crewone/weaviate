/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

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

// WeaviateActionsReferencesDeleteReader is a Reader for the WeaviateActionsReferencesDelete structure.
type WeaviateActionsReferencesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsReferencesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateActionsReferencesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsReferencesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsReferencesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionsReferencesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsReferencesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsReferencesDeleteNoContent creates a WeaviateActionsReferencesDeleteNoContent with default headers values
func NewWeaviateActionsReferencesDeleteNoContent() *WeaviateActionsReferencesDeleteNoContent {
	return &WeaviateActionsReferencesDeleteNoContent{}
}

/*WeaviateActionsReferencesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type WeaviateActionsReferencesDeleteNoContent struct {
}

func (o *WeaviateActionsReferencesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesDeleteNoContent ", 204)
}

func (o *WeaviateActionsReferencesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesDeleteUnauthorized creates a WeaviateActionsReferencesDeleteUnauthorized with default headers values
func NewWeaviateActionsReferencesDeleteUnauthorized() *WeaviateActionsReferencesDeleteUnauthorized {
	return &WeaviateActionsReferencesDeleteUnauthorized{}
}

/*WeaviateActionsReferencesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsReferencesDeleteUnauthorized struct {
}

func (o *WeaviateActionsReferencesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesDeleteUnauthorized ", 401)
}

func (o *WeaviateActionsReferencesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesDeleteForbidden creates a WeaviateActionsReferencesDeleteForbidden with default headers values
func NewWeaviateActionsReferencesDeleteForbidden() *WeaviateActionsReferencesDeleteForbidden {
	return &WeaviateActionsReferencesDeleteForbidden{}
}

/*WeaviateActionsReferencesDeleteForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateActionsReferencesDeleteForbidden struct {
}

func (o *WeaviateActionsReferencesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesDeleteForbidden ", 403)
}

func (o *WeaviateActionsReferencesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesDeleteNotFound creates a WeaviateActionsReferencesDeleteNotFound with default headers values
func NewWeaviateActionsReferencesDeleteNotFound() *WeaviateActionsReferencesDeleteNotFound {
	return &WeaviateActionsReferencesDeleteNotFound{}
}

/*WeaviateActionsReferencesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionsReferencesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *WeaviateActionsReferencesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsReferencesDeleteInternalServerError creates a WeaviateActionsReferencesDeleteInternalServerError with default headers values
func NewWeaviateActionsReferencesDeleteInternalServerError() *WeaviateActionsReferencesDeleteInternalServerError {
	return &WeaviateActionsReferencesDeleteInternalServerError{}
}

/*WeaviateActionsReferencesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsReferencesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsReferencesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
