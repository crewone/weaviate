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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateSchemaThingsPropertiesAddOKCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesAddOK
const WeaviateSchemaThingsPropertiesAddOKCode int = 200

/*WeaviateSchemaThingsPropertiesAddOK Added the property

swagger:response weaviateSchemaThingsPropertiesAddOK
*/
type WeaviateSchemaThingsPropertiesAddOK struct {
}

// NewWeaviateSchemaThingsPropertiesAddOK creates WeaviateSchemaThingsPropertiesAddOK with default headers values
func NewWeaviateSchemaThingsPropertiesAddOK() *WeaviateSchemaThingsPropertiesAddOK {

	return &WeaviateSchemaThingsPropertiesAddOK{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesAddOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WeaviateSchemaThingsPropertiesAddUnauthorizedCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesAddUnauthorized
const WeaviateSchemaThingsPropertiesAddUnauthorizedCode int = 401

/*WeaviateSchemaThingsPropertiesAddUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateSchemaThingsPropertiesAddUnauthorized
*/
type WeaviateSchemaThingsPropertiesAddUnauthorized struct {
}

// NewWeaviateSchemaThingsPropertiesAddUnauthorized creates WeaviateSchemaThingsPropertiesAddUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesAddUnauthorized() *WeaviateSchemaThingsPropertiesAddUnauthorized {

	return &WeaviateSchemaThingsPropertiesAddUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesAddUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateSchemaThingsPropertiesAddForbiddenCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesAddForbidden
const WeaviateSchemaThingsPropertiesAddForbiddenCode int = 403

/*WeaviateSchemaThingsPropertiesAddForbidden Could not find the Thing class

swagger:response weaviateSchemaThingsPropertiesAddForbidden
*/
type WeaviateSchemaThingsPropertiesAddForbidden struct {
}

// NewWeaviateSchemaThingsPropertiesAddForbidden creates WeaviateSchemaThingsPropertiesAddForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesAddForbidden() *WeaviateSchemaThingsPropertiesAddForbidden {

	return &WeaviateSchemaThingsPropertiesAddForbidden{}
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesAddForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateSchemaThingsPropertiesAddUnprocessableEntityCode is the HTTP code returned for type WeaviateSchemaThingsPropertiesAddUnprocessableEntity
const WeaviateSchemaThingsPropertiesAddUnprocessableEntityCode int = 422

/*WeaviateSchemaThingsPropertiesAddUnprocessableEntity Invalid property

swagger:response weaviateSchemaThingsPropertiesAddUnprocessableEntity
*/
type WeaviateSchemaThingsPropertiesAddUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateSchemaThingsPropertiesAddUnprocessableEntity creates WeaviateSchemaThingsPropertiesAddUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsPropertiesAddUnprocessableEntity() *WeaviateSchemaThingsPropertiesAddUnprocessableEntity {

	return &WeaviateSchemaThingsPropertiesAddUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate schema things properties add unprocessable entity response
func (o *WeaviateSchemaThingsPropertiesAddUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateSchemaThingsPropertiesAddUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate schema things properties add unprocessable entity response
func (o *WeaviateSchemaThingsPropertiesAddUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateSchemaThingsPropertiesAddUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
