//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsCreateOKCode is the HTTP code returned for type SchemaObjectsCreateOK
const SchemaObjectsCreateOKCode int = 200

/*
SchemaObjectsCreateOK Added the new Object class to the schema.

swagger:response schemaObjectsCreateOK
*/
type SchemaObjectsCreateOK struct {
	/*
	  In: Body
	*/
	Payload *models.Class `json:"body,omitempty"`
}

// NewSchemaObjectsCreateOK creates SchemaObjectsCreateOK with default headers values
func NewSchemaObjectsCreateOK() *SchemaObjectsCreateOK {
	return &SchemaObjectsCreateOK{}
}

// WithPayload adds the payload to the schema objects create o k response
func (o *SchemaObjectsCreateOK) WithPayload(payload *models.Class) *SchemaObjectsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects create o k response
func (o *SchemaObjectsCreateOK) SetPayload(payload *models.Class) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsCreateUnauthorizedCode is the HTTP code returned for type SchemaObjectsCreateUnauthorized
const SchemaObjectsCreateUnauthorizedCode int = 401

/*
SchemaObjectsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaObjectsCreateUnauthorized
*/
type SchemaObjectsCreateUnauthorized struct{}

// NewSchemaObjectsCreateUnauthorized creates SchemaObjectsCreateUnauthorized with default headers values
func NewSchemaObjectsCreateUnauthorized() *SchemaObjectsCreateUnauthorized {
	return &SchemaObjectsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaObjectsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaObjectsCreateForbiddenCode is the HTTP code returned for type SchemaObjectsCreateForbidden
const SchemaObjectsCreateForbiddenCode int = 403

/*
SchemaObjectsCreateForbidden Forbidden

swagger:response schemaObjectsCreateForbidden
*/
type SchemaObjectsCreateForbidden struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsCreateForbidden creates SchemaObjectsCreateForbidden with default headers values
func NewSchemaObjectsCreateForbidden() *SchemaObjectsCreateForbidden {
	return &SchemaObjectsCreateForbidden{}
}

// WithPayload adds the payload to the schema objects create forbidden response
func (o *SchemaObjectsCreateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaObjectsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects create forbidden response
func (o *SchemaObjectsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsCreateUnprocessableEntityCode is the HTTP code returned for type SchemaObjectsCreateUnprocessableEntity
const SchemaObjectsCreateUnprocessableEntityCode int = 422

/*
SchemaObjectsCreateUnprocessableEntity Invalid Object class

swagger:response schemaObjectsCreateUnprocessableEntity
*/
type SchemaObjectsCreateUnprocessableEntity struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsCreateUnprocessableEntity creates SchemaObjectsCreateUnprocessableEntity with default headers values
func NewSchemaObjectsCreateUnprocessableEntity() *SchemaObjectsCreateUnprocessableEntity {
	return &SchemaObjectsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema objects create unprocessable entity response
func (o *SchemaObjectsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaObjectsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects create unprocessable entity response
func (o *SchemaObjectsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaObjectsCreateInternalServerErrorCode is the HTTP code returned for type SchemaObjectsCreateInternalServerError
const SchemaObjectsCreateInternalServerErrorCode int = 500

/*
SchemaObjectsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaObjectsCreateInternalServerError
*/
type SchemaObjectsCreateInternalServerError struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaObjectsCreateInternalServerError creates SchemaObjectsCreateInternalServerError with default headers values
func NewSchemaObjectsCreateInternalServerError() *SchemaObjectsCreateInternalServerError {
	return &SchemaObjectsCreateInternalServerError{}
}

// WithPayload adds the payload to the schema objects create internal server error response
func (o *SchemaObjectsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaObjectsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema objects create internal server error response
func (o *SchemaObjectsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaObjectsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
