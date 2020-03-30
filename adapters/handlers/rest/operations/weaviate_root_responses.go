//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateRootOKCode is the HTTP code returned for type WeaviateRootOK
const WeaviateRootOKCode int = 200

/*WeaviateRootOK Weaviate is alive and ready to serve content

swagger:response weaviateRootOK
*/
type WeaviateRootOK struct {

	/*
	  In: Body
	*/
	Payload *WeaviateRootOKBody `json:"body,omitempty"`
}

// NewWeaviateRootOK creates WeaviateRootOK with default headers values
func NewWeaviateRootOK() *WeaviateRootOK {

	return &WeaviateRootOK{}
}

// WithPayload adds the payload to the weaviate root o k response
func (o *WeaviateRootOK) WithPayload(payload *WeaviateRootOKBody) *WeaviateRootOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate root o k response
func (o *WeaviateRootOK) SetPayload(payload *WeaviateRootOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateRootOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
