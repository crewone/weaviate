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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateThingsReferencesDeleteReader is a Reader for the WeaviateThingsReferencesDelete structure.
type WeaviateThingsReferencesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsReferencesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateThingsReferencesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsReferencesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsReferencesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsReferencesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsReferencesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsReferencesDeleteNoContent creates a WeaviateThingsReferencesDeleteNoContent with default headers values
func NewWeaviateThingsReferencesDeleteNoContent() *WeaviateThingsReferencesDeleteNoContent {
	return &WeaviateThingsReferencesDeleteNoContent{}
}

/*WeaviateThingsReferencesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type WeaviateThingsReferencesDeleteNoContent struct {
}

func (o *WeaviateThingsReferencesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesDeleteNoContent ", 204)
}

func (o *WeaviateThingsReferencesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesDeleteUnauthorized creates a WeaviateThingsReferencesDeleteUnauthorized with default headers values
func NewWeaviateThingsReferencesDeleteUnauthorized() *WeaviateThingsReferencesDeleteUnauthorized {
	return &WeaviateThingsReferencesDeleteUnauthorized{}
}

/*WeaviateThingsReferencesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsReferencesDeleteUnauthorized struct {
}

func (o *WeaviateThingsReferencesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesDeleteUnauthorized ", 401)
}

func (o *WeaviateThingsReferencesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesDeleteForbidden creates a WeaviateThingsReferencesDeleteForbidden with default headers values
func NewWeaviateThingsReferencesDeleteForbidden() *WeaviateThingsReferencesDeleteForbidden {
	return &WeaviateThingsReferencesDeleteForbidden{}
}

/*WeaviateThingsReferencesDeleteForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateThingsReferencesDeleteForbidden struct {
}

func (o *WeaviateThingsReferencesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesDeleteForbidden ", 403)
}

func (o *WeaviateThingsReferencesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesDeleteNotFound creates a WeaviateThingsReferencesDeleteNotFound with default headers values
func NewWeaviateThingsReferencesDeleteNotFound() *WeaviateThingsReferencesDeleteNotFound {
	return &WeaviateThingsReferencesDeleteNotFound{}
}

/*WeaviateThingsReferencesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsReferencesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsReferencesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *WeaviateThingsReferencesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsReferencesDeleteInternalServerError creates a WeaviateThingsReferencesDeleteInternalServerError with default headers values
func NewWeaviateThingsReferencesDeleteInternalServerError() *WeaviateThingsReferencesDeleteInternalServerError {
	return &WeaviateThingsReferencesDeleteInternalServerError{}
}

/*WeaviateThingsReferencesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsReferencesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsReferencesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsReferencesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
