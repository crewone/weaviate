//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
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

// ObjectsClassPutOKCode is the HTTP code returned for type ObjectsClassPutOK
const ObjectsClassPutOKCode int = 200

/*
ObjectsClassPutOK Successfully received.

swagger:response objectsClassPutOK
*/
type ObjectsClassPutOK struct {

	/*
	  In: Body
	*/
	Payload *models.Object `json:"body,omitempty"`
}

// NewObjectsClassPutOK creates ObjectsClassPutOK with default headers values
func NewObjectsClassPutOK() *ObjectsClassPutOK {

	return &ObjectsClassPutOK{}
}

// WithPayload adds the payload to the objects class put o k response
func (o *ObjectsClassPutOK) WithPayload(payload *models.Object) *ObjectsClassPutOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class put o k response
func (o *ObjectsClassPutOK) SetPayload(payload *models.Object) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassPutUnauthorizedCode is the HTTP code returned for type ObjectsClassPutUnauthorized
const ObjectsClassPutUnauthorizedCode int = 401

/*
ObjectsClassPutUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassPutUnauthorized
*/
type ObjectsClassPutUnauthorized struct {
}

// NewObjectsClassPutUnauthorized creates ObjectsClassPutUnauthorized with default headers values
func NewObjectsClassPutUnauthorized() *ObjectsClassPutUnauthorized {

	return &ObjectsClassPutUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassPutUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassPutForbiddenCode is the HTTP code returned for type ObjectsClassPutForbidden
const ObjectsClassPutForbiddenCode int = 403

/*
ObjectsClassPutForbidden Forbidden

swagger:response objectsClassPutForbidden
*/
type ObjectsClassPutForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPutForbidden creates ObjectsClassPutForbidden with default headers values
func NewObjectsClassPutForbidden() *ObjectsClassPutForbidden {

	return &ObjectsClassPutForbidden{}
}

// WithPayload adds the payload to the objects class put forbidden response
func (o *ObjectsClassPutForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassPutForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class put forbidden response
func (o *ObjectsClassPutForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPutForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassPutNotFoundCode is the HTTP code returned for type ObjectsClassPutNotFound
const ObjectsClassPutNotFoundCode int = 404

/*
ObjectsClassPutNotFound Successful query result but no resource was found.

swagger:response objectsClassPutNotFound
*/
type ObjectsClassPutNotFound struct {
}

// NewObjectsClassPutNotFound creates ObjectsClassPutNotFound with default headers values
func NewObjectsClassPutNotFound() *ObjectsClassPutNotFound {

	return &ObjectsClassPutNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassPutNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassPutUnprocessableEntityCode is the HTTP code returned for type ObjectsClassPutUnprocessableEntity
const ObjectsClassPutUnprocessableEntityCode int = 422

/*
ObjectsClassPutUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response objectsClassPutUnprocessableEntity
*/
type ObjectsClassPutUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPutUnprocessableEntity creates ObjectsClassPutUnprocessableEntity with default headers values
func NewObjectsClassPutUnprocessableEntity() *ObjectsClassPutUnprocessableEntity {

	return &ObjectsClassPutUnprocessableEntity{}
}

// WithPayload adds the payload to the objects class put unprocessable entity response
func (o *ObjectsClassPutUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsClassPutUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class put unprocessable entity response
func (o *ObjectsClassPutUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPutUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassPutInternalServerErrorCode is the HTTP code returned for type ObjectsClassPutInternalServerError
const ObjectsClassPutInternalServerErrorCode int = 500

/*
ObjectsClassPutInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassPutInternalServerError
*/
type ObjectsClassPutInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassPutInternalServerError creates ObjectsClassPutInternalServerError with default headers values
func NewObjectsClassPutInternalServerError() *ObjectsClassPutInternalServerError {

	return &ObjectsClassPutInternalServerError{}
}

// WithPayload adds the payload to the objects class put internal server error response
func (o *ObjectsClassPutInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassPutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class put internal server error response
func (o *ObjectsClassPutInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassPutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
