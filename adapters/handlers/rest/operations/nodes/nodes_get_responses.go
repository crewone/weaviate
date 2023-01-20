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

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// NodesGetOKCode is the HTTP code returned for type NodesGetOK
const NodesGetOKCode int = 200

/*
NodesGetOK Nodes status successfully returned

swagger:response nodesGetOK
*/
type NodesGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.NodesStatusResponse `json:"body,omitempty"`
}

// NewNodesGetOK creates NodesGetOK with default headers values
func NewNodesGetOK() *NodesGetOK {

	return &NodesGetOK{}
}

// WithPayload adds the payload to the nodes get o k response
func (o *NodesGetOK) WithPayload(payload *models.NodesStatusResponse) *NodesGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get o k response
func (o *NodesGetOK) SetPayload(payload *models.NodesStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodesGetUnauthorizedCode is the HTTP code returned for type NodesGetUnauthorized
const NodesGetUnauthorizedCode int = 401

/*
NodesGetUnauthorized Unauthorized or invalid credentials.

swagger:response nodesGetUnauthorized
*/
type NodesGetUnauthorized struct {
}

// NewNodesGetUnauthorized creates NodesGetUnauthorized with default headers values
func NewNodesGetUnauthorized() *NodesGetUnauthorized {

	return &NodesGetUnauthorized{}
}

// WriteResponse to the client
func (o *NodesGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// NodesGetForbiddenCode is the HTTP code returned for type NodesGetForbidden
const NodesGetForbiddenCode int = 403

/*
NodesGetForbidden Forbidden

swagger:response nodesGetForbidden
*/
type NodesGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodesGetForbidden creates NodesGetForbidden with default headers values
func NewNodesGetForbidden() *NodesGetForbidden {

	return &NodesGetForbidden{}
}

// WithPayload adds the payload to the nodes get forbidden response
func (o *NodesGetForbidden) WithPayload(payload *models.ErrorResponse) *NodesGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get forbidden response
func (o *NodesGetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodesGetNotFoundCode is the HTTP code returned for type NodesGetNotFound
const NodesGetNotFoundCode int = 404

/*
NodesGetNotFound Not Found - Backup does not exist

swagger:response nodesGetNotFound
*/
type NodesGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodesGetNotFound creates NodesGetNotFound with default headers values
func NewNodesGetNotFound() *NodesGetNotFound {

	return &NodesGetNotFound{}
}

// WithPayload adds the payload to the nodes get not found response
func (o *NodesGetNotFound) WithPayload(payload *models.ErrorResponse) *NodesGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get not found response
func (o *NodesGetNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodesGetUnprocessableEntityCode is the HTTP code returned for type NodesGetUnprocessableEntity
const NodesGetUnprocessableEntityCode int = 422

/*
NodesGetUnprocessableEntity Invalid backup restoration status attempt.

swagger:response nodesGetUnprocessableEntity
*/
type NodesGetUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodesGetUnprocessableEntity creates NodesGetUnprocessableEntity with default headers values
func NewNodesGetUnprocessableEntity() *NodesGetUnprocessableEntity {

	return &NodesGetUnprocessableEntity{}
}

// WithPayload adds the payload to the nodes get unprocessable entity response
func (o *NodesGetUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *NodesGetUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get unprocessable entity response
func (o *NodesGetUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodesGetInternalServerErrorCode is the HTTP code returned for type NodesGetInternalServerError
const NodesGetInternalServerErrorCode int = 500

/*
NodesGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response nodesGetInternalServerError
*/
type NodesGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodesGetInternalServerError creates NodesGetInternalServerError with default headers values
func NewNodesGetInternalServerError() *NodesGetInternalServerError {

	return &NodesGetInternalServerError{}
}

// WithPayload adds the payload to the nodes get internal server error response
func (o *NodesGetInternalServerError) WithPayload(payload *models.ErrorResponse) *NodesGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the nodes get internal server error response
func (o *NodesGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodesGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
