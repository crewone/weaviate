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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaThingsPropertiesUpdateOKCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesUpdateOK
const WeaviateSchemaThingsPropertiesUpdateOKCode int = 200

/*WeaviateSchemaThingsPropertiesUpdateOK Changes applied.

swagger:response weaviateSchemaThingsPropertiesUpdateOK
*/
type WeaviateSchemaThingsPropertiesUpdateOK struct {
}

// NewWeaviateSchemaThingsPropertiesUpdateOK creates WeaviateSchemaThingsPropertiesUpdateOK with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateOK() *WeaviateSchemaThingsPropertiesUpdateOK {

	return &WeaviateSchemaThingsPropertiesUpdateOK{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateSchemaThingsPropertiesUpdateUnauthorizedCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesUpdateUnauthorized
const WeaviateSchemaThingsPropertiesUpdateUnauthorizedCode int = 401

/*WeaviateSchemaThingsPropertiesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaThingsPropertiesUpdateUnauthorized
*/
type WeaviateSchemaThingsPropertiesUpdateUnauthorized struct {
}

// NewWeaviateSchemaThingsPropertiesUpdateUnauthorized creates WeaviateSchemaThingsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateUnauthorized() *WeaviateSchemaThingsPropertiesUpdateUnauthorized {

	return &WeaviateSchemaThingsPropertiesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateSchemaThingsPropertiesUpdateForbiddenCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesUpdateForbidden
const WeaviateSchemaThingsPropertiesUpdateForbiddenCode int = 403

/*WeaviateSchemaThingsPropertiesUpdateForbidden Forbidden

swagger:response weaviateSchemaThingsPropertiesUpdateForbidden
*/
type WeaviateSchemaThingsPropertiesUpdateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaThingsPropertiesUpdateForbidden creates WeaviateSchemaThingsPropertiesUpdateForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateForbidden() *WeaviateSchemaThingsPropertiesUpdateForbidden {

	return &WeaviateSchemaThingsPropertiesUpdateForbidden{}
}

// WithPayload adds the payload to the weaviate schema things properties update forbidden response
func (o *WeaviateSchemaThingsPropertiesUpdateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaThingsPropertiesUpdateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema things properties update forbidden response
func (o *WeaviateSchemaThingsPropertiesUpdateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaThingsPropertiesUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity
const WeaviateSchemaThingsPropertiesUpdateUnprocessableEntityCode int = 422

/*WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity Invalid update.

swagger:response weaviateSchemaThingsPropertiesUpdateUnprocessableEntity
*/
type WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaThingsPropertiesUpdateUnprocessableEntity creates WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateUnprocessableEntity() *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity {

	return &WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate schema things properties update unprocessable entity response
func (o *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema things properties update unprocessable entity response
func (o *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaThingsPropertiesUpdateInternalServerErrorCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesUpdateInternalServerError
const WeaviateSchemaThingsPropertiesUpdateInternalServerErrorCode int = 500

/*WeaviateSchemaThingsPropertiesUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateSchemaThingsPropertiesUpdateInternalServerError
*/
type WeaviateSchemaThingsPropertiesUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaThingsPropertiesUpdateInternalServerError creates WeaviateSchemaThingsPropertiesUpdateInternalServerError with default headers values
func NewWeaviateSchemaThingsPropertiesUpdateInternalServerError() *WeaviateSchemaThingsPropertiesUpdateInternalServerError {

	return &WeaviateSchemaThingsPropertiesUpdateInternalServerError{}
}

// WithPayload adds the payload to the weaviate schema things properties update internal server error response
func (o *WeaviateSchemaThingsPropertiesUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaThingsPropertiesUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema things properties update internal server error response
func (o *WeaviateSchemaThingsPropertiesUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
