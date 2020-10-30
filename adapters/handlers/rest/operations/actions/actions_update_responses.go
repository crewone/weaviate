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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ActionsUpdateOKCode is the HTTP code returned for type ActionsUpdateOK
const ActionsUpdateOKCode int = 200

/*ActionsUpdateOK Successfully received.

swagger:response actionsUpdateOK
*/
type ActionsUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Action `json:"body,omitempty"`
}

// NewActionsUpdateOK creates ActionsUpdateOK with default headers values
func NewActionsUpdateOK() *ActionsUpdateOK {
	return &ActionsUpdateOK{}
}

// WithPayload adds the payload to the actions update o k response
func (o *ActionsUpdateOK) WithPayload(payload *models.Action) *ActionsUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions update o k response
func (o *ActionsUpdateOK) SetPayload(payload *models.Action) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsUpdateUnauthorizedCode is the HTTP code returned for type ActionsUpdateUnauthorized
const ActionsUpdateUnauthorizedCode int = 401

/*ActionsUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response actionsUpdateUnauthorized
*/
type ActionsUpdateUnauthorized struct {
}

// NewActionsUpdateUnauthorized creates ActionsUpdateUnauthorized with default headers values
func NewActionsUpdateUnauthorized() *ActionsUpdateUnauthorized {
	return &ActionsUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *ActionsUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActionsUpdateForbiddenCode is the HTTP code returned for type ActionsUpdateForbidden
const ActionsUpdateForbiddenCode int = 403

/*ActionsUpdateForbidden Forbidden

swagger:response actionsUpdateForbidden
*/
type ActionsUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsUpdateForbidden creates ActionsUpdateForbidden with default headers values
func NewActionsUpdateForbidden() *ActionsUpdateForbidden {
	return &ActionsUpdateForbidden{}
}

// WithPayload adds the payload to the actions update forbidden response
func (o *ActionsUpdateForbidden) WithPayload(payload *models.ErrorResponse) *ActionsUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions update forbidden response
func (o *ActionsUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsUpdateNotFoundCode is the HTTP code returned for type ActionsUpdateNotFound
const ActionsUpdateNotFoundCode int = 404

/*ActionsUpdateNotFound Successful query result but no resource was found.

swagger:response actionsUpdateNotFound
*/
type ActionsUpdateNotFound struct {
}

// NewActionsUpdateNotFound creates ActionsUpdateNotFound with default headers values
func NewActionsUpdateNotFound() *ActionsUpdateNotFound {
	return &ActionsUpdateNotFound{}
}

// WriteResponse to the client
func (o *ActionsUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ActionsUpdateUnprocessableEntityCode is the HTTP code returned for type ActionsUpdateUnprocessableEntity
const ActionsUpdateUnprocessableEntityCode int = 422

/*ActionsUpdateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response actionsUpdateUnprocessableEntity
*/
type ActionsUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsUpdateUnprocessableEntity creates ActionsUpdateUnprocessableEntity with default headers values
func NewActionsUpdateUnprocessableEntity() *ActionsUpdateUnprocessableEntity {
	return &ActionsUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the actions update unprocessable entity response
func (o *ActionsUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ActionsUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions update unprocessable entity response
func (o *ActionsUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsUpdateInternalServerErrorCode is the HTTP code returned for type ActionsUpdateInternalServerError
const ActionsUpdateInternalServerErrorCode int = 500

/*ActionsUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response actionsUpdateInternalServerError
*/
type ActionsUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsUpdateInternalServerError creates ActionsUpdateInternalServerError with default headers values
func NewActionsUpdateInternalServerError() *ActionsUpdateInternalServerError {
	return &ActionsUpdateInternalServerError{}
}

// WithPayload adds the payload to the actions update internal server error response
func (o *ActionsUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *ActionsUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions update internal server error response
func (o *ActionsUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
