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




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateAdaptersInsertCreatedCode is the HTTP code returned for type WeaviateAdaptersInsertCreated
const WeaviateAdaptersInsertCreatedCode int = 201

/*WeaviateAdaptersInsertCreated Successful created.

swagger:response weaviateAdaptersInsertCreated
*/
type WeaviateAdaptersInsertCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Adapter `json:"body,omitempty"`
}

// NewWeaviateAdaptersInsertCreated creates WeaviateAdaptersInsertCreated with default headers values
func NewWeaviateAdaptersInsertCreated() *WeaviateAdaptersInsertCreated {
	return &WeaviateAdaptersInsertCreated{}
}

// WithPayload adds the payload to the weaviate adapters insert created response
func (o *WeaviateAdaptersInsertCreated) WithPayload(payload *models.Adapter) *WeaviateAdaptersInsertCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate adapters insert created response
func (o *WeaviateAdaptersInsertCreated) SetPayload(payload *models.Adapter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateAdaptersInsertCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateAdaptersInsertNotImplementedCode is the HTTP code returned for type WeaviateAdaptersInsertNotImplemented
const WeaviateAdaptersInsertNotImplementedCode int = 501

/*WeaviateAdaptersInsertNotImplemented Not (yet) implemented.

swagger:response weaviateAdaptersInsertNotImplemented
*/
type WeaviateAdaptersInsertNotImplemented struct {
}

// NewWeaviateAdaptersInsertNotImplemented creates WeaviateAdaptersInsertNotImplemented with default headers values
func NewWeaviateAdaptersInsertNotImplemented() *WeaviateAdaptersInsertNotImplemented {
	return &WeaviateAdaptersInsertNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateAdaptersInsertNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
