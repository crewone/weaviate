/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsPropertiesCreateOKCode is the HTTP code returned for type WeaviateActionsPropertiesCreateOK
const WeaviateActionsPropertiesCreateOKCode int = 200

/*WeaviateActionsPropertiesCreateOK Successfully added the reference.

swagger:response weaviateActionsPropertiesCreateOK
*/
type WeaviateActionsPropertiesCreateOK struct {
}

// NewWeaviateActionsPropertiesCreateOK creates WeaviateActionsPropertiesCreateOK with default headers values
func NewWeaviateActionsPropertiesCreateOK() *WeaviateActionsPropertiesCreateOK {

	return &WeaviateActionsPropertiesCreateOK{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateActionsPropertiesCreateUnauthorizedCode is the HTTP code returned for type WeaviateActionsPropertiesCreateUnauthorized
const WeaviateActionsPropertiesCreateUnauthorizedCode int = 401

/*WeaviateActionsPropertiesCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsPropertiesCreateUnauthorized
*/
type WeaviateActionsPropertiesCreateUnauthorized struct {
}

// NewWeaviateActionsPropertiesCreateUnauthorized creates WeaviateActionsPropertiesCreateUnauthorized with default headers values
func NewWeaviateActionsPropertiesCreateUnauthorized() *WeaviateActionsPropertiesCreateUnauthorized {

	return &WeaviateActionsPropertiesCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsPropertiesCreateForbiddenCode is the HTTP code returned for type WeaviateActionsPropertiesCreateForbidden
const WeaviateActionsPropertiesCreateForbiddenCode int = 403

/*WeaviateActionsPropertiesCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsPropertiesCreateForbidden
*/
type WeaviateActionsPropertiesCreateForbidden struct {
}

// NewWeaviateActionsPropertiesCreateForbidden creates WeaviateActionsPropertiesCreateForbidden with default headers values
func NewWeaviateActionsPropertiesCreateForbidden() *WeaviateActionsPropertiesCreateForbidden {

	return &WeaviateActionsPropertiesCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsPropertiesCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateActionsPropertiesCreateUnprocessableEntity
const WeaviateActionsPropertiesCreateUnprocessableEntityCode int = 422

/*WeaviateActionsPropertiesCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response weaviateActionsPropertiesCreateUnprocessableEntity
*/
type WeaviateActionsPropertiesCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPropertiesCreateUnprocessableEntity creates WeaviateActionsPropertiesCreateUnprocessableEntity with default headers values
func NewWeaviateActionsPropertiesCreateUnprocessableEntity() *WeaviateActionsPropertiesCreateUnprocessableEntity {

	return &WeaviateActionsPropertiesCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate actions properties create unprocessable entity response
func (o *WeaviateActionsPropertiesCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPropertiesCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions properties create unprocessable entity response
func (o *WeaviateActionsPropertiesCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsPropertiesCreateInternalServerErrorCode is the HTTP code returned for type WeaviateActionsPropertiesCreateInternalServerError
const WeaviateActionsPropertiesCreateInternalServerErrorCode int = 500

/*WeaviateActionsPropertiesCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsPropertiesCreateInternalServerError
*/
type WeaviateActionsPropertiesCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPropertiesCreateInternalServerError creates WeaviateActionsPropertiesCreateInternalServerError with default headers values
func NewWeaviateActionsPropertiesCreateInternalServerError() *WeaviateActionsPropertiesCreateInternalServerError {

	return &WeaviateActionsPropertiesCreateInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions properties create internal server error response
func (o *WeaviateActionsPropertiesCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPropertiesCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions properties create internal server error response
func (o *WeaviateActionsPropertiesCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
