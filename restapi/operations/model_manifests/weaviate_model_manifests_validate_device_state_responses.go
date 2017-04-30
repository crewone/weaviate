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

// WeaviateModelManifestsValidateDeviceStateCreatedCode is the HTTP code returned for type WeaviateModelManifestsValidateDeviceStateCreated
const WeaviateModelManifestsValidateDeviceStateCreatedCode int = 201

/*WeaviateModelManifestsValidateDeviceStateCreated Successful created.

swagger:response weaviateModelManifestsValidateDeviceStateCreated
*/
type WeaviateModelManifestsValidateDeviceStateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ModelManifestsValidateDeviceStateResponse `json:"body,omitempty"`
}

// NewWeaviateModelManifestsValidateDeviceStateCreated creates WeaviateModelManifestsValidateDeviceStateCreated with default headers values
func NewWeaviateModelManifestsValidateDeviceStateCreated() *WeaviateModelManifestsValidateDeviceStateCreated {
	return &WeaviateModelManifestsValidateDeviceStateCreated{}
}

// WithPayload adds the payload to the weaviate model manifests validate device state created response
func (o *WeaviateModelManifestsValidateDeviceStateCreated) WithPayload(payload *models.ModelManifestsValidateDeviceStateResponse) *WeaviateModelManifestsValidateDeviceStateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate model manifests validate device state created response
func (o *WeaviateModelManifestsValidateDeviceStateCreated) SetPayload(payload *models.ModelManifestsValidateDeviceStateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateModelManifestsValidateDeviceStateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateModelManifestsValidateDeviceStateNotImplementedCode is the HTTP code returned for type WeaviateModelManifestsValidateDeviceStateNotImplemented
const WeaviateModelManifestsValidateDeviceStateNotImplementedCode int = 501

/*WeaviateModelManifestsValidateDeviceStateNotImplemented Not (yet) implemented.

swagger:response weaviateModelManifestsValidateDeviceStateNotImplemented
*/
type WeaviateModelManifestsValidateDeviceStateNotImplemented struct {
}

// NewWeaviateModelManifestsValidateDeviceStateNotImplemented creates WeaviateModelManifestsValidateDeviceStateNotImplemented with default headers values
func NewWeaviateModelManifestsValidateDeviceStateNotImplemented() *WeaviateModelManifestsValidateDeviceStateNotImplemented {
	return &WeaviateModelManifestsValidateDeviceStateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateModelManifestsValidateDeviceStateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
