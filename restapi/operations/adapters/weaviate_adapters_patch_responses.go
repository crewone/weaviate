/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package adapters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateAdaptersPatchOKCode is the HTTP code returned for type WeaviateAdaptersPatchOK
const WeaviateAdaptersPatchOKCode int = 200

/*WeaviateAdaptersPatchOK Successful updated.

swagger:response weaviateAdaptersPatchOK
*/
type WeaviateAdaptersPatchOK struct {

	/*
	  In: Body
	*/
	Payload *models.Adapter `json:"body,omitempty"`
}

// NewWeaviateAdaptersPatchOK creates WeaviateAdaptersPatchOK with default headers values
func NewWeaviateAdaptersPatchOK() *WeaviateAdaptersPatchOK {
	return &WeaviateAdaptersPatchOK{}
}

// WithPayload adds the payload to the weaviate adapters patch o k response
func (o *WeaviateAdaptersPatchOK) WithPayload(payload *models.Adapter) *WeaviateAdaptersPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate adapters patch o k response
func (o *WeaviateAdaptersPatchOK) SetPayload(payload *models.Adapter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateAdaptersPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateAdaptersPatchNotImplementedCode is the HTTP code returned for type WeaviateAdaptersPatchNotImplemented
const WeaviateAdaptersPatchNotImplementedCode int = 501

/*WeaviateAdaptersPatchNotImplemented Not (yet) implemented.

swagger:response weaviateAdaptersPatchNotImplemented
*/
type WeaviateAdaptersPatchNotImplemented struct {
}

// NewWeaviateAdaptersPatchNotImplemented creates WeaviateAdaptersPatchNotImplemented with default headers values
func NewWeaviateAdaptersPatchNotImplemented() *WeaviateAdaptersPatchNotImplemented {
	return &WeaviateAdaptersPatchNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateAdaptersPatchNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
