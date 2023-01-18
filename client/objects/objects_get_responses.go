//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsGetReader is a Reader for the ObjectsGet structure.
type ObjectsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewObjectsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewObjectsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsGetOK creates a ObjectsGetOK with default headers values
func NewObjectsGetOK() *ObjectsGetOK {
	return &ObjectsGetOK{}
}

/*
ObjectsGetOK handles this case with default header values.

Successful response.
*/
type ObjectsGetOK struct {
	Payload *models.Object
}

func (o *ObjectsGetOK) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetOK  %+v", 200, o.Payload)
}

func (o *ObjectsGetOK) GetPayload() *models.Object {
	return o.Payload
}

func (o *ObjectsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Object)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetBadRequest creates a ObjectsGetBadRequest with default headers values
func NewObjectsGetBadRequest() *ObjectsGetBadRequest {
	return &ObjectsGetBadRequest{}
}

/*
ObjectsGetBadRequest handles this case with default header values.

Malformed request.
*/
type ObjectsGetBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ObjectsGetBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetUnauthorized creates a ObjectsGetUnauthorized with default headers values
func NewObjectsGetUnauthorized() *ObjectsGetUnauthorized {
	return &ObjectsGetUnauthorized{}
}

/*
ObjectsGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsGetUnauthorized struct {
}

func (o *ObjectsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetUnauthorized ", 401)
}

func (o *ObjectsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsGetForbidden creates a ObjectsGetForbidden with default headers values
func NewObjectsGetForbidden() *ObjectsGetForbidden {
	return &ObjectsGetForbidden{}
}

/*
ObjectsGetForbidden handles this case with default header values.

Forbidden
*/
type ObjectsGetForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsGetNotFound creates a ObjectsGetNotFound with default headers values
func NewObjectsGetNotFound() *ObjectsGetNotFound {
	return &ObjectsGetNotFound{}
}

/*
ObjectsGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ObjectsGetNotFound struct {
}

func (o *ObjectsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetNotFound ", 404)
}

func (o *ObjectsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsGetInternalServerError creates a ObjectsGetInternalServerError with default headers values
func NewObjectsGetInternalServerError() *ObjectsGetInternalServerError {
	return &ObjectsGetInternalServerError{}
}

/*
ObjectsGetInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /objects/{id}][%d] objectsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
