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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateThingsReferencesUpdateOKCode is the HTTP code returned for type WeaviateThingsReferencesUpdateOK
const WeaviateThingsReferencesUpdateOKCode int = 200

/*WeaviateThingsReferencesUpdateOK Successfully replaced all the references (success is based on the behavior of the datastore).

swagger:response weaviateThingsReferencesUpdateOK
*/
type WeaviateThingsReferencesUpdateOK struct {
}

// NewWeaviateThingsReferencesUpdateOK creates WeaviateThingsReferencesUpdateOK with default headers values
func NewWeaviateThingsReferencesUpdateOK() *WeaviateThingsReferencesUpdateOK {

	return &WeaviateThingsReferencesUpdateOK{}
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateThingsReferencesUpdateUnauthorizedCode is the HTTP code returned for type WeaviateThingsReferencesUpdateUnauthorized
const WeaviateThingsReferencesUpdateUnauthorizedCode int = 401

/*WeaviateThingsReferencesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsReferencesUpdateUnauthorized
*/
type WeaviateThingsReferencesUpdateUnauthorized struct {
}

// NewWeaviateThingsReferencesUpdateUnauthorized creates WeaviateThingsReferencesUpdateUnauthorized with default headers values
func NewWeaviateThingsReferencesUpdateUnauthorized() *WeaviateThingsReferencesUpdateUnauthorized {

	return &WeaviateThingsReferencesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsReferencesUpdateForbiddenCode is the HTTP code returned for type WeaviateThingsReferencesUpdateForbidden
const WeaviateThingsReferencesUpdateForbiddenCode int = 403

/*WeaviateThingsReferencesUpdateForbidden Forbidden

swagger:response weaviateThingsReferencesUpdateForbidden
*/
type WeaviateThingsReferencesUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsReferencesUpdateForbidden creates WeaviateThingsReferencesUpdateForbidden with default headers values
func NewWeaviateThingsReferencesUpdateForbidden() *WeaviateThingsReferencesUpdateForbidden {

	return &WeaviateThingsReferencesUpdateForbidden{}
}

// WithPayload adds the payload to the weaviate things references update forbidden response
func (o *WeaviateThingsReferencesUpdateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateThingsReferencesUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things references update forbidden response
func (o *WeaviateThingsReferencesUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsReferencesUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateThingsReferencesUpdateUnprocessableEntity
const WeaviateThingsReferencesUpdateUnprocessableEntityCode int = 422

/*WeaviateThingsReferencesUpdateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response weaviateThingsReferencesUpdateUnprocessableEntity
*/
type WeaviateThingsReferencesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsReferencesUpdateUnprocessableEntity creates WeaviateThingsReferencesUpdateUnprocessableEntity with default headers values
func NewWeaviateThingsReferencesUpdateUnprocessableEntity() *WeaviateThingsReferencesUpdateUnprocessableEntity {

	return &WeaviateThingsReferencesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate things references update unprocessable entity response
func (o *WeaviateThingsReferencesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateThingsReferencesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things references update unprocessable entity response
func (o *WeaviateThingsReferencesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsReferencesUpdateInternalServerErrorCode is the HTTP code returned for type WeaviateThingsReferencesUpdateInternalServerError
const WeaviateThingsReferencesUpdateInternalServerErrorCode int = 500

/*WeaviateThingsReferencesUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateThingsReferencesUpdateInternalServerError
*/
type WeaviateThingsReferencesUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsReferencesUpdateInternalServerError creates WeaviateThingsReferencesUpdateInternalServerError with default headers values
func NewWeaviateThingsReferencesUpdateInternalServerError() *WeaviateThingsReferencesUpdateInternalServerError {

	return &WeaviateThingsReferencesUpdateInternalServerError{}
}

// WithPayload adds the payload to the weaviate things references update internal server error response
func (o *WeaviateThingsReferencesUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateThingsReferencesUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things references update internal server error response
func (o *WeaviateThingsReferencesUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
