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

// WeaviateActionsPropertiesUpdateOKCode is the HTTP code returned for type WeaviateActionsPropertiesUpdateOK
const WeaviateActionsPropertiesUpdateOKCode int = 200

/*WeaviateActionsPropertiesUpdateOK Successfully replaced all the refferences.

swagger:response weaviateActionsPropertiesUpdateOK
*/
type WeaviateActionsPropertiesUpdateOK struct {
}

// NewWeaviateActionsPropertiesUpdateOK creates WeaviateActionsPropertiesUpdateOK with default headers values
func NewWeaviateActionsPropertiesUpdateOK() *WeaviateActionsPropertiesUpdateOK {

	return &WeaviateActionsPropertiesUpdateOK{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateActionsPropertiesUpdateUnauthorizedCode is the HTTP code returned for type WeaviateActionsPropertiesUpdateUnauthorized
const WeaviateActionsPropertiesUpdateUnauthorizedCode int = 401

/*WeaviateActionsPropertiesUpdateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsPropertiesUpdateUnauthorized
*/
type WeaviateActionsPropertiesUpdateUnauthorized struct {
}

// NewWeaviateActionsPropertiesUpdateUnauthorized creates WeaviateActionsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateActionsPropertiesUpdateUnauthorized() *WeaviateActionsPropertiesUpdateUnauthorized {

	return &WeaviateActionsPropertiesUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsPropertiesUpdateForbiddenCode is the HTTP code returned for type WeaviateActionsPropertiesUpdateForbidden
const WeaviateActionsPropertiesUpdateForbiddenCode int = 403

/*WeaviateActionsPropertiesUpdateForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsPropertiesUpdateForbidden
*/
type WeaviateActionsPropertiesUpdateForbidden struct {
}

// NewWeaviateActionsPropertiesUpdateForbidden creates WeaviateActionsPropertiesUpdateForbidden with default headers values
func NewWeaviateActionsPropertiesUpdateForbidden() *WeaviateActionsPropertiesUpdateForbidden {

	return &WeaviateActionsPropertiesUpdateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesUpdateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsPropertiesUpdateUnprocessableEntityCode is the HTTP code returned for type WeaviateActionsPropertiesUpdateUnprocessableEntity
const WeaviateActionsPropertiesUpdateUnprocessableEntityCode int = 422

/*WeaviateActionsPropertiesUpdateUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?

swagger:response weaviateActionsPropertiesUpdateUnprocessableEntity
*/
type WeaviateActionsPropertiesUpdateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsPropertiesUpdateUnprocessableEntity creates WeaviateActionsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateActionsPropertiesUpdateUnprocessableEntity() *WeaviateActionsPropertiesUpdateUnprocessableEntity {

	return &WeaviateActionsPropertiesUpdateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate actions properties update unprocessable entity response
func (o *WeaviateActionsPropertiesUpdateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionsPropertiesUpdateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions properties update unprocessable entity response
func (o *WeaviateActionsPropertiesUpdateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsPropertiesUpdateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
