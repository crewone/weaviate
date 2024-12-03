//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// GetRolesForUserOKCode is the HTTP code returned for type GetRolesForUserOK
const GetRolesForUserOKCode int = 200

/*
GetRolesForUserOK Role assigned users

swagger:response getRolesForUserOK
*/
type GetRolesForUserOK struct {

	/*
	  In: Body
	*/
	Payload models.RolesListResponse `json:"body,omitempty"`
}

// NewGetRolesForUserOK creates GetRolesForUserOK with default headers values
func NewGetRolesForUserOK() *GetRolesForUserOK {

	return &GetRolesForUserOK{}
}

// WithPayload adds the payload to the get roles for user o k response
func (o *GetRolesForUserOK) WithPayload(payload models.RolesListResponse) *GetRolesForUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles for user o k response
func (o *GetRolesForUserOK) SetPayload(payload models.RolesListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesForUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.RolesListResponse{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetRolesForUserBadRequestCode is the HTTP code returned for type GetRolesForUserBadRequest
const GetRolesForUserBadRequestCode int = 400

/*
GetRolesForUserBadRequest Bad request

swagger:response getRolesForUserBadRequest
*/
type GetRolesForUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesForUserBadRequest creates GetRolesForUserBadRequest with default headers values
func NewGetRolesForUserBadRequest() *GetRolesForUserBadRequest {

	return &GetRolesForUserBadRequest{}
}

// WithPayload adds the payload to the get roles for user bad request response
func (o *GetRolesForUserBadRequest) WithPayload(payload *models.ErrorResponse) *GetRolesForUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles for user bad request response
func (o *GetRolesForUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesForUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesForUserUnauthorizedCode is the HTTP code returned for type GetRolesForUserUnauthorized
const GetRolesForUserUnauthorizedCode int = 401

/*
GetRolesForUserUnauthorized Unauthorized or invalid credentials.

swagger:response getRolesForUserUnauthorized
*/
type GetRolesForUserUnauthorized struct {
}

// NewGetRolesForUserUnauthorized creates GetRolesForUserUnauthorized with default headers values
func NewGetRolesForUserUnauthorized() *GetRolesForUserUnauthorized {

	return &GetRolesForUserUnauthorized{}
}

// WriteResponse to the client
func (o *GetRolesForUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetRolesForUserForbiddenCode is the HTTP code returned for type GetRolesForUserForbidden
const GetRolesForUserForbiddenCode int = 403

/*
GetRolesForUserForbidden Forbidden

swagger:response getRolesForUserForbidden
*/
type GetRolesForUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesForUserForbidden creates GetRolesForUserForbidden with default headers values
func NewGetRolesForUserForbidden() *GetRolesForUserForbidden {

	return &GetRolesForUserForbidden{}
}

// WithPayload adds the payload to the get roles for user forbidden response
func (o *GetRolesForUserForbidden) WithPayload(payload *models.ErrorResponse) *GetRolesForUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles for user forbidden response
func (o *GetRolesForUserForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesForUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRolesForUserNotFoundCode is the HTTP code returned for type GetRolesForUserNotFound
const GetRolesForUserNotFoundCode int = 404

/*
GetRolesForUserNotFound no role found for user

swagger:response getRolesForUserNotFound
*/
type GetRolesForUserNotFound struct {
}

// NewGetRolesForUserNotFound creates GetRolesForUserNotFound with default headers values
func NewGetRolesForUserNotFound() *GetRolesForUserNotFound {

	return &GetRolesForUserNotFound{}
}

// WriteResponse to the client
func (o *GetRolesForUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetRolesForUserInternalServerErrorCode is the HTTP code returned for type GetRolesForUserInternalServerError
const GetRolesForUserInternalServerErrorCode int = 500

/*
GetRolesForUserInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response getRolesForUserInternalServerError
*/
type GetRolesForUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetRolesForUserInternalServerError creates GetRolesForUserInternalServerError with default headers values
func NewGetRolesForUserInternalServerError() *GetRolesForUserInternalServerError {

	return &GetRolesForUserInternalServerError{}
}

// WithPayload adds the payload to the get roles for user internal server error response
func (o *GetRolesForUserInternalServerError) WithPayload(payload *models.ErrorResponse) *GetRolesForUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get roles for user internal server error response
func (o *GetRolesForUserInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRolesForUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
