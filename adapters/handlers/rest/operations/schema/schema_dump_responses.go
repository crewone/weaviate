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

// SchemaDumpOKCode is the HTTP code returned for type SchemaDumpOK
const SchemaDumpOKCode int = 200

/*SchemaDumpOK Successfully dumped the database schema.

swagger:response schemaDumpOK
*/
type SchemaDumpOK struct {

	/*
	  In: Body
	*/
	Payload *SchemaDumpOKBody `json:"body,omitempty"`
}

// NewSchemaDumpOK creates SchemaDumpOK with default headers values
func NewSchemaDumpOK() *SchemaDumpOK {

	return &SchemaDumpOK{}
}

// WithPayload adds the payload to the schema dump o k response
func (o *SchemaDumpOK) WithPayload(payload *SchemaDumpOKBody) *SchemaDumpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema dump o k response
func (o *SchemaDumpOK) SetPayload(payload *SchemaDumpOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaDumpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaDumpUnauthorizedCode is the HTTP code returned for type SchemaDumpUnauthorized
const SchemaDumpUnauthorizedCode int = 401

/*SchemaDumpUnauthorized Unauthorized or invalid credentials.

swagger:response schemaDumpUnauthorized
*/
type SchemaDumpUnauthorized struct {
}

// NewSchemaDumpUnauthorized creates SchemaDumpUnauthorized with default headers values
func NewSchemaDumpUnauthorized() *SchemaDumpUnauthorized {

	return &SchemaDumpUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaDumpUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaDumpForbiddenCode is the HTTP code returned for type SchemaDumpForbidden
const SchemaDumpForbiddenCode int = 403

/*SchemaDumpForbidden Forbidden

swagger:response schemaDumpForbidden
*/
type SchemaDumpForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaDumpForbidden creates SchemaDumpForbidden with default headers values
func NewSchemaDumpForbidden() *SchemaDumpForbidden {

	return &SchemaDumpForbidden{}
}

// WithPayload adds the payload to the schema dump forbidden response
func (o *SchemaDumpForbidden) WithPayload(payload *models.ErrorResponse) *SchemaDumpForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema dump forbidden response
func (o *SchemaDumpForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaDumpForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaDumpInternalServerErrorCode is the HTTP code returned for type SchemaDumpInternalServerError
const SchemaDumpInternalServerErrorCode int = 500

/*SchemaDumpInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaDumpInternalServerError
*/
type SchemaDumpInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaDumpInternalServerError creates SchemaDumpInternalServerError with default headers values
func NewSchemaDumpInternalServerError() *SchemaDumpInternalServerError {

	return &SchemaDumpInternalServerError{}
}

// WithPayload adds the payload to the schema dump internal server error response
func (o *SchemaDumpInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaDumpInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema dump internal server error response
func (o *SchemaDumpInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaDumpInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
