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
 package model_manifests




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateModelManifestsUpdateOKCode is the HTTP code returned for type WeaviateModelManifestsUpdateOK
const WeaviateModelManifestsUpdateOKCode int = 200

/*WeaviateModelManifestsUpdateOK Successful updated.

swagger:response weaviateModelManifestsUpdateOK
*/
type WeaviateModelManifestsUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ModelManifest `json:"body,omitempty"`
}

// NewWeaviateModelManifestsUpdateOK creates WeaviateModelManifestsUpdateOK with default headers values
func NewWeaviateModelManifestsUpdateOK() *WeaviateModelManifestsUpdateOK {
	return &WeaviateModelManifestsUpdateOK{}
}

// WithPayload adds the payload to the weaviate model manifests update o k response
func (o *WeaviateModelManifestsUpdateOK) WithPayload(payload *models.ModelManifest) *WeaviateModelManifestsUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate model manifests update o k response
func (o *WeaviateModelManifestsUpdateOK) SetPayload(payload *models.ModelManifest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateModelManifestsUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateModelManifestsUpdateNotImplementedCode is the HTTP code returned for type WeaviateModelManifestsUpdateNotImplemented
const WeaviateModelManifestsUpdateNotImplementedCode int = 501

/*WeaviateModelManifestsUpdateNotImplemented Not (yet) implemented.

swagger:response weaviateModelManifestsUpdateNotImplemented
*/
type WeaviateModelManifestsUpdateNotImplemented struct {
}

// NewWeaviateModelManifestsUpdateNotImplemented creates WeaviateModelManifestsUpdateNotImplemented with default headers values
func NewWeaviateModelManifestsUpdateNotImplemented() *WeaviateModelManifestsUpdateNotImplemented {
	return &WeaviateModelManifestsUpdateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateModelManifestsUpdateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
