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


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateDevicesListHandlerFunc turns a function with the right signature into a weaviate devices list handler
type WeaviateDevicesListHandlerFunc func(WeaviateDevicesListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateDevicesListHandlerFunc) Handle(params WeaviateDevicesListParams) middleware.Responder {
	return fn(params)
}

// WeaviateDevicesListHandler interface for that can handle valid weaviate devices list params
type WeaviateDevicesListHandler interface {
	Handle(WeaviateDevicesListParams) middleware.Responder
}

// NewWeaviateDevicesList creates a new http.Handler for the weaviate devices list operation
func NewWeaviateDevicesList(ctx *middleware.Context, handler WeaviateDevicesListHandler) *WeaviateDevicesList {
	return &WeaviateDevicesList{Context: ctx, Handler: handler}
}

/*WeaviateDevicesList swagger:route GET /devices devices weaviateDevicesList

Lists all devices user has access to.

*/
type WeaviateDevicesList struct {
	Context *middleware.Context
	Handler WeaviateDevicesListHandler
}

func (o *WeaviateDevicesList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateDevicesListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
