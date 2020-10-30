//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ThingsReferencesUpdateOKCode is the HTTP code returned for type ThingsReferencesUpdateOK
const ThingsReferencesUpdateOKCode int = 200

/*ThingsReferencesUpdateOK Successfully replaced all the references (success is based on the behavior of the datastore).

swagger:response thingsReferencesUpdateOK
*/
type ThingsReferencesUpdateOK struct {
}

// NewThingsReferencesUpdateOK creates ThingsReferencesUpdateOK with default headers values
func NewThingsReferencesUpdateOK() *ThingsReferencesUpdateOK {
	return &ThingsReferencesUpdateOK{}
}

// WriteResponse to the client
func (o *ThingsReferencesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ThingsReferencesUpdateUnauthorizedCode is the HTTP code returned for type ThingsReferencesUpdateUnauthorized
const ThingsReferencesUpdateUnauthorizedCode int = 401

/*ThingsReferencesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response thingsReferencesUpdateUnauthorized
*/
type ThingsReferencesUpdateUnauthorized struct {
}

// NewThingsReferencesUpdateUnauthorized creates ThingsReferencesUpdateUnauthorized with default headers values
func NewThingsReferencesUpdateUnauthorized() *ThingsReferencesUpdateUnauthorized {
	return &ThingsReferencesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsReferencesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsReferencesUpdateForbiddenCode is the HTTP code returned for type ThingsReferencesUpdateForbidden
const ThingsReferencesUpdateForbiddenCode int = 403

/*ThingsReferencesUpdateForbidden Forbidden

swagger:response thingsReferencesUpdateForbidden
*/
type ThingsReferencesUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesUpdateForbidden creates ThingsReferencesUpdateForbidden with default headers values
func NewThingsReferencesUpdateForbidden() *ThingsReferencesUpdateForbidden {
	return &ThingsReferencesUpdateForbidden{}
}

// WithPayload adds the payload to the things references update forbidden response
func (o *ThingsReferencesUpdateForbidden) WithPayload(payload *models.ErrorResponse) *ThingsReferencesUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references update forbidden response
func (o *ThingsReferencesUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesUpdateUnprocessableEntityCode is the HTTP code returned for type ThingsReferencesUpdateUnprocessableEntity
const ThingsReferencesUpdateUnprocessableEntityCode int = 422

/*ThingsReferencesUpdateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response thingsReferencesUpdateUnprocessableEntity
*/
type ThingsReferencesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesUpdateUnprocessableEntity creates ThingsReferencesUpdateUnprocessableEntity with default headers values
func NewThingsReferencesUpdateUnprocessableEntity() *ThingsReferencesUpdateUnprocessableEntity {
	return &ThingsReferencesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the things references update unprocessable entity response
func (o *ThingsReferencesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ThingsReferencesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references update unprocessable entity response
func (o *ThingsReferencesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsReferencesUpdateInternalServerErrorCode is the HTTP code returned for type ThingsReferencesUpdateInternalServerError
const ThingsReferencesUpdateInternalServerErrorCode int = 500

/*ThingsReferencesUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsReferencesUpdateInternalServerError
*/
type ThingsReferencesUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsReferencesUpdateInternalServerError creates ThingsReferencesUpdateInternalServerError with default headers values
func NewThingsReferencesUpdateInternalServerError() *ThingsReferencesUpdateInternalServerError {
	return &ThingsReferencesUpdateInternalServerError{}
}

// WithPayload adds the payload to the things references update internal server error response
func (o *ThingsReferencesUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsReferencesUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things references update internal server error response
func (o *ThingsReferencesUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsReferencesUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
