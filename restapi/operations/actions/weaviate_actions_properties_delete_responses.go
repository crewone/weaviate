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

// WeaviateActionsPropertiesDeleteNoContentCode is the HTTP code returned for type WeaviateActionsPropertiesDeleteNoContent
const WeaviateActionsPropertiesDeleteNoContentCode int = 204

/*WeaviateActionsPropertiesDeleteNoContent Successfully deleted.

swagger:response weaviateActionsPropertiesDeleteNoContent
*/
type WeaviateActionsPropertiesDeleteNoContent struct {
}

// NewWeaviateActionsPropertiesDeleteNoContent creates WeaviateActionsPropertiesDeleteNoContent with default headers values
func NewWeaviateActionsPropertiesDeleteNoContent() *WeaviateActionsPropertiesDeleteNoContent {

	return &WeaviateActionsPropertiesDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// WeaviateActionsPropertiesDeleteUnauthorizedCode is the HTTP code returned for type WeaviateActionsPropertiesDeleteUnauthorized
const WeaviateActionsPropertiesDeleteUnauthorizedCode int = 401

/*WeaviateActionsPropertiesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsPropertiesDeleteUnauthorized
*/
type WeaviateActionsPropertiesDeleteUnauthorized struct {
}

// NewWeaviateActionsPropertiesDeleteUnauthorized creates WeaviateActionsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateActionsPropertiesDeleteUnauthorized() *WeaviateActionsPropertiesDeleteUnauthorized {

	return &WeaviateActionsPropertiesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsPropertiesDeleteForbiddenCode is the HTTP code returned for type WeaviateActionsPropertiesDeleteForbidden
const WeaviateActionsPropertiesDeleteForbiddenCode int = 403

/*WeaviateActionsPropertiesDeleteForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsPropertiesDeleteForbidden
*/
type WeaviateActionsPropertiesDeleteForbidden struct {
}

// NewWeaviateActionsPropertiesDeleteForbidden creates WeaviateActionsPropertiesDeleteForbidden with default headers values
func NewWeaviateActionsPropertiesDeleteForbidden() *WeaviateActionsPropertiesDeleteForbidden {

	return &WeaviateActionsPropertiesDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsPropertiesDeleteNotFoundCode is the HTTP code returned for type WeaviateActionsPropertiesDeleteNotFound
const WeaviateActionsPropertiesDeleteNotFoundCode int = 404

/*WeaviateActionsPropertiesDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateActionsPropertiesDeleteNotFound
*/
type WeaviateActionsPropertiesDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPropertiesDeleteNotFound creates WeaviateActionsPropertiesDeleteNotFound with default headers values
func NewWeaviateActionsPropertiesDeleteNotFound() *WeaviateActionsPropertiesDeleteNotFound {

	return &WeaviateActionsPropertiesDeleteNotFound{}
}

// WithPayload adds the payload to the weaviate actions properties delete not found response
func (o *WeaviateActionsPropertiesDeleteNotFound) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPropertiesDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions properties delete not found response
func (o *WeaviateActionsPropertiesDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsPropertiesDeleteInternalServerErrorCode is the HTTP code returned for type WeaviateActionsPropertiesDeleteInternalServerError
const WeaviateActionsPropertiesDeleteInternalServerErrorCode int = 500

/*WeaviateActionsPropertiesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateActionsPropertiesDeleteInternalServerError
*/
type WeaviateActionsPropertiesDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPropertiesDeleteInternalServerError creates WeaviateActionsPropertiesDeleteInternalServerError with default headers values
func NewWeaviateActionsPropertiesDeleteInternalServerError() *WeaviateActionsPropertiesDeleteInternalServerError {

	return &WeaviateActionsPropertiesDeleteInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions properties delete internal server error response
func (o *WeaviateActionsPropertiesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPropertiesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions properties delete internal server error response
func (o *WeaviateActionsPropertiesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
