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

// WeaviateLocationsGetOKCode is the HTTP code returned for type WeaviateLocationsGetOK
const WeaviateLocationsGetOKCode int = 200

/*WeaviateLocationsGetOK Successful response.

swagger:response weaviateLocationsGetOK
*/
type WeaviateLocationsGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Location `json:"body,omitempty"`
}

// NewWeaviateLocationsGetOK creates WeaviateLocationsGetOK with default headers values
func NewWeaviateLocationsGetOK() *WeaviateLocationsGetOK {
	return &WeaviateLocationsGetOK{}
}

// WithPayload adds the payload to the weaviate locations get o k response
func (o *WeaviateLocationsGetOK) WithPayload(payload *models.Location) *WeaviateLocationsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate locations get o k response
func (o *WeaviateLocationsGetOK) SetPayload(payload *models.Location) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateLocationsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateLocationsGetNotImplementedCode is the HTTP code returned for type WeaviateLocationsGetNotImplemented
const WeaviateLocationsGetNotImplementedCode int = 501

/*WeaviateLocationsGetNotImplemented Not (yet) implemented.

swagger:response weaviateLocationsGetNotImplemented
*/
type WeaviateLocationsGetNotImplemented struct {
}

// NewWeaviateLocationsGetNotImplemented creates WeaviateLocationsGetNotImplemented with default headers values
func NewWeaviateLocationsGetNotImplemented() *WeaviateLocationsGetNotImplemented {
	return &WeaviateLocationsGetNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateLocationsGetNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
