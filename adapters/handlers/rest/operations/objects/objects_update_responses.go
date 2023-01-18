//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsUpdateOKCode is the HTTP code returned for type ObjectsUpdateOK
const ObjectsUpdateOKCode int = 200

/*
ObjectsUpdateOK Successfully received.

swagger:response objectsUpdateOK
*/
type ObjectsUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Object `json:"body,omitempty"`
}

// NewObjectsUpdateOK creates ObjectsUpdateOK with default headers values
func NewObjectsUpdateOK() *ObjectsUpdateOK {

	return &ObjectsUpdateOK{}
}

// WithPayload adds the payload to the objects update o k response
func (o *ObjectsUpdateOK) WithPayload(payload *models.Object) *ObjectsUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects update o k response
func (o *ObjectsUpdateOK) SetPayload(payload *models.Object) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsUpdateUnauthorizedCode is the HTTP code returned for type ObjectsUpdateUnauthorized
const ObjectsUpdateUnauthorizedCode int = 401

/*
ObjectsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response objectsUpdateUnauthorized
*/
type ObjectsUpdateUnauthorized struct {
}

// NewObjectsUpdateUnauthorized creates ObjectsUpdateUnauthorized with default headers values
func NewObjectsUpdateUnauthorized() *ObjectsUpdateUnauthorized {

	return &ObjectsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsUpdateForbiddenCode is the HTTP code returned for type ObjectsUpdateForbidden
const ObjectsUpdateForbiddenCode int = 403

/*
ObjectsUpdateForbidden Forbidden

swagger:response objectsUpdateForbidden
*/
type ObjectsUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsUpdateForbidden creates ObjectsUpdateForbidden with default headers values
func NewObjectsUpdateForbidden() *ObjectsUpdateForbidden {

	return &ObjectsUpdateForbidden{}
}

// WithPayload adds the payload to the objects update forbidden response
func (o *ObjectsUpdateForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects update forbidden response
func (o *ObjectsUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsUpdateNotFoundCode is the HTTP code returned for type ObjectsUpdateNotFound
const ObjectsUpdateNotFoundCode int = 404

/*
ObjectsUpdateNotFound Successful query result but no resource was found.

swagger:response objectsUpdateNotFound
*/
type ObjectsUpdateNotFound struct {
}

// NewObjectsUpdateNotFound creates ObjectsUpdateNotFound with default headers values
func NewObjectsUpdateNotFound() *ObjectsUpdateNotFound {

	return &ObjectsUpdateNotFound{}
}

// WriteResponse to the client
func (o *ObjectsUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsUpdateUnprocessableEntityCode is the HTTP code returned for type ObjectsUpdateUnprocessableEntity
const ObjectsUpdateUnprocessableEntityCode int = 422

/*
ObjectsUpdateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response objectsUpdateUnprocessableEntity
*/
type ObjectsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsUpdateUnprocessableEntity creates ObjectsUpdateUnprocessableEntity with default headers values
func NewObjectsUpdateUnprocessableEntity() *ObjectsUpdateUnprocessableEntity {

	return &ObjectsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the objects update unprocessable entity response
func (o *ObjectsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects update unprocessable entity response
func (o *ObjectsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsUpdateInternalServerErrorCode is the HTTP code returned for type ObjectsUpdateInternalServerError
const ObjectsUpdateInternalServerErrorCode int = 500

/*
ObjectsUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsUpdateInternalServerError
*/
type ObjectsUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsUpdateInternalServerError creates ObjectsUpdateInternalServerError with default headers values
func NewObjectsUpdateInternalServerError() *ObjectsUpdateInternalServerError {

	return &ObjectsUpdateInternalServerError{}
}

// WithPayload adds the payload to the objects update internal server error response
func (o *ObjectsUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects update internal server error response
func (o *ObjectsUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
