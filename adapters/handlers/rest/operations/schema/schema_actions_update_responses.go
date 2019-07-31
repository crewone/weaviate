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
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsUpdateOKCode is the HTTP code returned for type SchemaActionsUpdateOK
const SchemaActionsUpdateOKCode int = 200

/*SchemaActionsUpdateOK Changes applied.

swagger:response schemaActionsUpdateOK
*/
type SchemaActionsUpdateOK struct {
}

// NewSchemaActionsUpdateOK creates SchemaActionsUpdateOK with default headers values
func NewSchemaActionsUpdateOK() *SchemaActionsUpdateOK {

	return &SchemaActionsUpdateOK{}
}

// WriteResponse to the client
func (o *SchemaActionsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SchemaActionsUpdateUnauthorizedCode is the HTTP code returned for type SchemaActionsUpdateUnauthorized
const SchemaActionsUpdateUnauthorizedCode int = 401

/*SchemaActionsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response schemaActionsUpdateUnauthorized
*/
type SchemaActionsUpdateUnauthorized struct {
}

// NewSchemaActionsUpdateUnauthorized creates SchemaActionsUpdateUnauthorized with default headers values
func NewSchemaActionsUpdateUnauthorized() *SchemaActionsUpdateUnauthorized {

	return &SchemaActionsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaActionsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaActionsUpdateForbiddenCode is the HTTP code returned for type SchemaActionsUpdateForbidden
const SchemaActionsUpdateForbiddenCode int = 403

/*SchemaActionsUpdateForbidden Forbidden

swagger:response schemaActionsUpdateForbidden
*/
type SchemaActionsUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsUpdateForbidden creates SchemaActionsUpdateForbidden with default headers values
func NewSchemaActionsUpdateForbidden() *SchemaActionsUpdateForbidden {

	return &SchemaActionsUpdateForbidden{}
}

// WithPayload adds the payload to the schema actions update forbidden response
func (o *SchemaActionsUpdateForbidden) WithPayload(payload *models.ErrorResponse) *SchemaActionsUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions update forbidden response
func (o *SchemaActionsUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsUpdateUnprocessableEntityCode is the HTTP code returned for type SchemaActionsUpdateUnprocessableEntity
const SchemaActionsUpdateUnprocessableEntityCode int = 422

/*SchemaActionsUpdateUnprocessableEntity Invalid update.

swagger:response schemaActionsUpdateUnprocessableEntity
*/
type SchemaActionsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsUpdateUnprocessableEntity creates SchemaActionsUpdateUnprocessableEntity with default headers values
func NewSchemaActionsUpdateUnprocessableEntity() *SchemaActionsUpdateUnprocessableEntity {

	return &SchemaActionsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the schema actions update unprocessable entity response
func (o *SchemaActionsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *SchemaActionsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions update unprocessable entity response
func (o *SchemaActionsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsUpdateInternalServerErrorCode is the HTTP code returned for type SchemaActionsUpdateInternalServerError
const SchemaActionsUpdateInternalServerErrorCode int = 500

/*SchemaActionsUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaActionsUpdateInternalServerError
*/
type SchemaActionsUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsUpdateInternalServerError creates SchemaActionsUpdateInternalServerError with default headers values
func NewSchemaActionsUpdateInternalServerError() *SchemaActionsUpdateInternalServerError {

	return &SchemaActionsUpdateInternalServerError{}
}

// WithPayload adds the payload to the schema actions update internal server error response
func (o *SchemaActionsUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaActionsUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions update internal server error response
func (o *SchemaActionsUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
