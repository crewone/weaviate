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

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// BackupsCreateStatusOKCode is the HTTP code returned for type BackupsCreateStatusOK
const BackupsCreateStatusOKCode int = 200

/*
BackupsCreateStatusOK Backup creation status successfully returned

swagger:response backupsCreateStatusOK
*/
type BackupsCreateStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.BackupCreateStatusResponse `json:"body,omitempty"`
}

// NewBackupsCreateStatusOK creates BackupsCreateStatusOK with default headers values
func NewBackupsCreateStatusOK() *BackupsCreateStatusOK {

	return &BackupsCreateStatusOK{}
}

// WithPayload adds the payload to the backups create status o k response
func (o *BackupsCreateStatusOK) WithPayload(payload *models.BackupCreateStatusResponse) *BackupsCreateStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backups create status o k response
func (o *BackupsCreateStatusOK) SetPayload(payload *models.BackupCreateStatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupsCreateStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupsCreateStatusUnauthorizedCode is the HTTP code returned for type BackupsCreateStatusUnauthorized
const BackupsCreateStatusUnauthorizedCode int = 401

/*
BackupsCreateStatusUnauthorized Unauthorized or invalid credentials.

swagger:response backupsCreateStatusUnauthorized
*/
type BackupsCreateStatusUnauthorized struct {
}

// NewBackupsCreateStatusUnauthorized creates BackupsCreateStatusUnauthorized with default headers values
func NewBackupsCreateStatusUnauthorized() *BackupsCreateStatusUnauthorized {

	return &BackupsCreateStatusUnauthorized{}
}

// WriteResponse to the client
func (o *BackupsCreateStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BackupsCreateStatusForbiddenCode is the HTTP code returned for type BackupsCreateStatusForbidden
const BackupsCreateStatusForbiddenCode int = 403

/*
BackupsCreateStatusForbidden Forbidden

swagger:response backupsCreateStatusForbidden
*/
type BackupsCreateStatusForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBackupsCreateStatusForbidden creates BackupsCreateStatusForbidden with default headers values
func NewBackupsCreateStatusForbidden() *BackupsCreateStatusForbidden {

	return &BackupsCreateStatusForbidden{}
}

// WithPayload adds the payload to the backups create status forbidden response
func (o *BackupsCreateStatusForbidden) WithPayload(payload *models.ErrorResponse) *BackupsCreateStatusForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backups create status forbidden response
func (o *BackupsCreateStatusForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupsCreateStatusForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupsCreateStatusNotFoundCode is the HTTP code returned for type BackupsCreateStatusNotFound
const BackupsCreateStatusNotFoundCode int = 404

/*
BackupsCreateStatusNotFound Not Found - Backup does not exist

swagger:response backupsCreateStatusNotFound
*/
type BackupsCreateStatusNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBackupsCreateStatusNotFound creates BackupsCreateStatusNotFound with default headers values
func NewBackupsCreateStatusNotFound() *BackupsCreateStatusNotFound {

	return &BackupsCreateStatusNotFound{}
}

// WithPayload adds the payload to the backups create status not found response
func (o *BackupsCreateStatusNotFound) WithPayload(payload *models.ErrorResponse) *BackupsCreateStatusNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backups create status not found response
func (o *BackupsCreateStatusNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupsCreateStatusNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupsCreateStatusUnprocessableEntityCode is the HTTP code returned for type BackupsCreateStatusUnprocessableEntity
const BackupsCreateStatusUnprocessableEntityCode int = 422

/*
BackupsCreateStatusUnprocessableEntity Invalid backup restoration status attempt.

swagger:response backupsCreateStatusUnprocessableEntity
*/
type BackupsCreateStatusUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBackupsCreateStatusUnprocessableEntity creates BackupsCreateStatusUnprocessableEntity with default headers values
func NewBackupsCreateStatusUnprocessableEntity() *BackupsCreateStatusUnprocessableEntity {

	return &BackupsCreateStatusUnprocessableEntity{}
}

// WithPayload adds the payload to the backups create status unprocessable entity response
func (o *BackupsCreateStatusUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BackupsCreateStatusUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backups create status unprocessable entity response
func (o *BackupsCreateStatusUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupsCreateStatusUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupsCreateStatusInternalServerErrorCode is the HTTP code returned for type BackupsCreateStatusInternalServerError
const BackupsCreateStatusInternalServerErrorCode int = 500

/*
BackupsCreateStatusInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response backupsCreateStatusInternalServerError
*/
type BackupsCreateStatusInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBackupsCreateStatusInternalServerError creates BackupsCreateStatusInternalServerError with default headers values
func NewBackupsCreateStatusInternalServerError() *BackupsCreateStatusInternalServerError {

	return &BackupsCreateStatusInternalServerError{}
}

// WithPayload adds the payload to the backups create status internal server error response
func (o *BackupsCreateStatusInternalServerError) WithPayload(payload *models.ErrorResponse) *BackupsCreateStatusInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backups create status internal server error response
func (o *BackupsCreateStatusInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupsCreateStatusInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
