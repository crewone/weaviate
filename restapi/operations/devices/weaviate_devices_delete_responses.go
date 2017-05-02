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
 package devices




import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateDevicesDeleteNoContentCode is the HTTP code returned for type WeaviateDevicesDeleteNoContent
const WeaviateDevicesDeleteNoContentCode int = 204

/*WeaviateDevicesDeleteNoContent Successful deleted.

swagger:response weaviateDevicesDeleteNoContent
*/
type WeaviateDevicesDeleteNoContent struct {
}

// NewWeaviateDevicesDeleteNoContent creates WeaviateDevicesDeleteNoContent with default headers values
func NewWeaviateDevicesDeleteNoContent() *WeaviateDevicesDeleteNoContent {
	return &WeaviateDevicesDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateDevicesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// WeaviateDevicesDeleteNotImplementedCode is the HTTP code returned for type WeaviateDevicesDeleteNotImplemented
const WeaviateDevicesDeleteNotImplementedCode int = 501

/*WeaviateDevicesDeleteNotImplemented Not (yet) implemented.

swagger:response weaviateDevicesDeleteNotImplemented
*/
type WeaviateDevicesDeleteNotImplemented struct {
}

// NewWeaviateDevicesDeleteNotImplemented creates WeaviateDevicesDeleteNotImplemented with default headers values
func NewWeaviateDevicesDeleteNotImplemented() *WeaviateDevicesDeleteNotImplemented {
	return &WeaviateDevicesDeleteNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateDevicesDeleteNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
