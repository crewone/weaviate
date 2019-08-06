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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// BatchingReferencesCreateOKCode is the HTTP code returned for type BatchingReferencesCreateOK
const BatchingReferencesCreateOKCode int = 200

/*BatchingReferencesCreateOK Request Successful. Warning: A successful request does not guarantuee that every batched reference was successfully created. Inspect the response body to see which references succeeded and which failed.

swagger:response batchingReferencesCreateOK
*/
type BatchingReferencesCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.BatchReferenceResponse `json:"body,omitempty"`
}

// NewBatchingReferencesCreateOK creates BatchingReferencesCreateOK with default headers values
func NewBatchingReferencesCreateOK() *BatchingReferencesCreateOK {

	return &BatchingReferencesCreateOK{}
}

// WithPayload adds the payload to the batching references create o k response
func (o *BatchingReferencesCreateOK) WithPayload(payload []*models.BatchReferenceResponse) *BatchingReferencesCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching references create o k response
func (o *BatchingReferencesCreateOK) SetPayload(payload []*models.BatchReferenceResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingReferencesCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.BatchReferenceResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// BatchingReferencesCreateUnauthorizedCode is the HTTP code returned for type BatchingReferencesCreateUnauthorized
const BatchingReferencesCreateUnauthorizedCode int = 401

/*BatchingReferencesCreateUnauthorized Unauthorized or invalid credentials.

swagger:response batchingReferencesCreateUnauthorized
*/
type BatchingReferencesCreateUnauthorized struct {
}

// NewBatchingReferencesCreateUnauthorized creates BatchingReferencesCreateUnauthorized with default headers values
func NewBatchingReferencesCreateUnauthorized() *BatchingReferencesCreateUnauthorized {

	return &BatchingReferencesCreateUnauthorized{}
}

// WriteResponse to the client
func (o *BatchingReferencesCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BatchingReferencesCreateForbiddenCode is the HTTP code returned for type BatchingReferencesCreateForbidden
const BatchingReferencesCreateForbiddenCode int = 403

/*BatchingReferencesCreateForbidden Forbidden

swagger:response batchingReferencesCreateForbidden
*/
type BatchingReferencesCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingReferencesCreateForbidden creates BatchingReferencesCreateForbidden with default headers values
func NewBatchingReferencesCreateForbidden() *BatchingReferencesCreateForbidden {

	return &BatchingReferencesCreateForbidden{}
}

// WithPayload adds the payload to the batching references create forbidden response
func (o *BatchingReferencesCreateForbidden) WithPayload(payload *models.ErrorResponse) *BatchingReferencesCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching references create forbidden response
func (o *BatchingReferencesCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingReferencesCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchingReferencesCreateUnprocessableEntityCode is the HTTP code returned for type BatchingReferencesCreateUnprocessableEntity
const BatchingReferencesCreateUnprocessableEntityCode int = 422

/*BatchingReferencesCreateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response batchingReferencesCreateUnprocessableEntity
*/
type BatchingReferencesCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingReferencesCreateUnprocessableEntity creates BatchingReferencesCreateUnprocessableEntity with default headers values
func NewBatchingReferencesCreateUnprocessableEntity() *BatchingReferencesCreateUnprocessableEntity {

	return &BatchingReferencesCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the batching references create unprocessable entity response
func (o *BatchingReferencesCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BatchingReferencesCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching references create unprocessable entity response
func (o *BatchingReferencesCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingReferencesCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchingReferencesCreateInternalServerErrorCode is the HTTP code returned for type BatchingReferencesCreateInternalServerError
const BatchingReferencesCreateInternalServerErrorCode int = 500

/*BatchingReferencesCreateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response batchingReferencesCreateInternalServerError
*/
type BatchingReferencesCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchingReferencesCreateInternalServerError creates BatchingReferencesCreateInternalServerError with default headers values
func NewBatchingReferencesCreateInternalServerError() *BatchingReferencesCreateInternalServerError {

	return &BatchingReferencesCreateInternalServerError{}
}

// WithPayload adds the payload to the batching references create internal server error response
func (o *BatchingReferencesCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *BatchingReferencesCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batching references create internal server error response
func (o *BatchingReferencesCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchingReferencesCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
