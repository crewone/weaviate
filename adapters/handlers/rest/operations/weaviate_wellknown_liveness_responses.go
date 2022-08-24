//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
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

// WeaviateWellknownLivenessOKCode is the HTTP code returned for type WeaviateWellknownLivenessOK
const WeaviateWellknownLivenessOKCode int = 200

/*
WeaviateWellknownLivenessOK The application is able to respond to HTTP requests

swagger:response weaviateWellknownLivenessOK
*/
type WeaviateWellknownLivenessOK struct{}

// NewWeaviateWellknownLivenessOK creates WeaviateWellknownLivenessOK with default headers values
func NewWeaviateWellknownLivenessOK() *WeaviateWellknownLivenessOK {
	return &WeaviateWellknownLivenessOK{}
}

// WriteResponse to the client
func (o *WeaviateWellknownLivenessOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
