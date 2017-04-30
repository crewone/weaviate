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
 package locations




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateLocationsInsertCreatedCode is the HTTP code returned for type WeaviateLocationsInsertCreated
const WeaviateLocationsInsertCreatedCode int = 201

/*WeaviateLocationsInsertCreated Successful created.

swagger:response weaviateLocationsInsertCreated
*/
type WeaviateLocationsInsertCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Location `json:"body,omitempty"`
}

// NewWeaviateLocationsInsertCreated creates WeaviateLocationsInsertCreated with default headers values
func NewWeaviateLocationsInsertCreated() *WeaviateLocationsInsertCreated {
	return &WeaviateLocationsInsertCreated{}
}

// WithPayload adds the payload to the weaviate locations insert created response
func (o *WeaviateLocationsInsertCreated) WithPayload(payload *models.Location) *WeaviateLocationsInsertCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate locations insert created response
func (o *WeaviateLocationsInsertCreated) SetPayload(payload *models.Location) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateLocationsInsertCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateLocationsInsertNotImplementedCode is the HTTP code returned for type WeaviateLocationsInsertNotImplemented
const WeaviateLocationsInsertNotImplementedCode int = 501

/*WeaviateLocationsInsertNotImplemented Not (yet) implemented.

swagger:response weaviateLocationsInsertNotImplemented
*/
type WeaviateLocationsInsertNotImplemented struct {
}

// NewWeaviateLocationsInsertNotImplemented creates WeaviateLocationsInsertNotImplemented with default headers values
func NewWeaviateLocationsInsertNotImplemented() *WeaviateLocationsInsertNotImplemented {
	return &WeaviateLocationsInsertNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateLocationsInsertNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
