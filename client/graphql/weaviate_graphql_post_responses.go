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

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateGraphqlPostReader is a Reader for the WeaviateGraphqlPost structure.
type WeaviateGraphqlPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateGraphqlPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateGraphqlPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateGraphqlPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateGraphqlPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateGraphqlPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateGraphqlPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateGraphqlPostOK creates a WeaviateGraphqlPostOK with default headers values
func NewWeaviateGraphqlPostOK() *WeaviateGraphqlPostOK {
	return &WeaviateGraphqlPostOK{}
}

/*WeaviateGraphqlPostOK handles this case with default header values.

Successful query (with select).
*/
type WeaviateGraphqlPostOK struct {
	Payload *models.GraphQLResponse
}

func (o *WeaviateGraphqlPostOK) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] weaviateGraphqlPostOK  %+v", 200, o.Payload)
}

func (o *WeaviateGraphqlPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphQLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateGraphqlPostUnauthorized creates a WeaviateGraphqlPostUnauthorized with default headers values
func NewWeaviateGraphqlPostUnauthorized() *WeaviateGraphqlPostUnauthorized {
	return &WeaviateGraphqlPostUnauthorized{}
}

/*WeaviateGraphqlPostUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateGraphqlPostUnauthorized struct {
}

func (o *WeaviateGraphqlPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] weaviateGraphqlPostUnauthorized ", 401)
}

func (o *WeaviateGraphqlPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateGraphqlPostForbidden creates a WeaviateGraphqlPostForbidden with default headers values
func NewWeaviateGraphqlPostForbidden() *WeaviateGraphqlPostForbidden {
	return &WeaviateGraphqlPostForbidden{}
}

/*WeaviateGraphqlPostForbidden handles this case with default header values.

Forbidden
*/
type WeaviateGraphqlPostForbidden struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateGraphqlPostForbidden) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] weaviateGraphqlPostForbidden  %+v", 403, o.Payload)
}

func (o *WeaviateGraphqlPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateGraphqlPostUnprocessableEntity creates a WeaviateGraphqlPostUnprocessableEntity with default headers values
func NewWeaviateGraphqlPostUnprocessableEntity() *WeaviateGraphqlPostUnprocessableEntity {
	return &WeaviateGraphqlPostUnprocessableEntity{}
}

/*WeaviateGraphqlPostUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateGraphqlPostUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateGraphqlPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] weaviateGraphqlPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateGraphqlPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateGraphqlPostInternalServerError creates a WeaviateGraphqlPostInternalServerError with default headers values
func NewWeaviateGraphqlPostInternalServerError() *WeaviateGraphqlPostInternalServerError {
	return &WeaviateGraphqlPostInternalServerError{}
}

/*WeaviateGraphqlPostInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateGraphqlPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateGraphqlPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /graphql][%d] weaviateGraphqlPostInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateGraphqlPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
