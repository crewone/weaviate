/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaActionsCreateOKCode is the HTTP code returned for type WeaviateSchemaActionsCreateOK
const WeaviateSchemaActionsCreateOKCode int = 200

/*WeaviateSchemaActionsCreateOK Added the new Action class to the ontology.

swagger:response weaviateSchemaActionsCreateOK
*/
type WeaviateSchemaActionsCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.SemanticSchemaClass `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsCreateOK creates WeaviateSchemaActionsCreateOK with default headers values
func NewWeaviateSchemaActionsCreateOK() *WeaviateSchemaActionsCreateOK {

	return &WeaviateSchemaActionsCreateOK{}
}

// WithPayload adds the payload to the weaviate schema actions create o k response
func (o *WeaviateSchemaActionsCreateOK) WithPayload(payload *models.SemanticSchemaClass) *WeaviateSchemaActionsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions create o k response
func (o *WeaviateSchemaActionsCreateOK) SetPayload(payload *models.SemanticSchemaClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsCreateUnauthorizedCode is the HTTP code returned for type WeaviateSchemaActionsCreateUnauthorized
const WeaviateSchemaActionsCreateUnauthorizedCode int = 401

/*WeaviateSchemaActionsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaActionsCreateUnauthorized
*/
type WeaviateSchemaActionsCreateUnauthorized struct {
}

// NewWeaviateSchemaActionsCreateUnauthorized creates WeaviateSchemaActionsCreateUnauthorized with default headers values
func NewWeaviateSchemaActionsCreateUnauthorized() *WeaviateSchemaActionsCreateUnauthorized {

	return &WeaviateSchemaActionsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateSchemaActionsCreateForbiddenCode is the HTTP code returned for type WeaviateSchemaActionsCreateForbidden
const WeaviateSchemaActionsCreateForbiddenCode int = 403

/*WeaviateSchemaActionsCreateForbidden Forbidden

swagger:response weaviateSchemaActionsCreateForbidden
*/
type WeaviateSchemaActionsCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsCreateForbidden creates WeaviateSchemaActionsCreateForbidden with default headers values
func NewWeaviateSchemaActionsCreateForbidden() *WeaviateSchemaActionsCreateForbidden {

	return &WeaviateSchemaActionsCreateForbidden{}
}

// WithPayload adds the payload to the weaviate schema actions create forbidden response
func (o *WeaviateSchemaActionsCreateForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions create forbidden response
func (o *WeaviateSchemaActionsCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateSchemaActionsCreateUnprocessableEntity
const WeaviateSchemaActionsCreateUnprocessableEntityCode int = 422

/*WeaviateSchemaActionsCreateUnprocessableEntity Invalid Action class

swagger:response weaviateSchemaActionsCreateUnprocessableEntity
*/
type WeaviateSchemaActionsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsCreateUnprocessableEntity creates WeaviateSchemaActionsCreateUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsCreateUnprocessableEntity() *WeaviateSchemaActionsCreateUnprocessableEntity {

	return &WeaviateSchemaActionsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate schema actions create unprocessable entity response
func (o *WeaviateSchemaActionsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions create unprocessable entity response
func (o *WeaviateSchemaActionsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateSchemaActionsCreateInternalServerErrorCode is the HTTP code returned for type WeaviateSchemaActionsCreateInternalServerError
const WeaviateSchemaActionsCreateInternalServerErrorCode int = 500

/*WeaviateSchemaActionsCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateSchemaActionsCreateInternalServerError
*/
type WeaviateSchemaActionsCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaActionsCreateInternalServerError creates WeaviateSchemaActionsCreateInternalServerError with default headers values
func NewWeaviateSchemaActionsCreateInternalServerError() *WeaviateSchemaActionsCreateInternalServerError {

	return &WeaviateSchemaActionsCreateInternalServerError{}
}

// WithPayload adds the payload to the weaviate schema actions create internal server error response
func (o *WeaviateSchemaActionsCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaActionsCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema actions create internal server error response
func (o *WeaviateSchemaActionsCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaActionsCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
