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


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateLocationsUpdateHandlerFunc turns a function with the right signature into a weaviate locations update handler
type WeaviateLocationsUpdateHandlerFunc func(WeaviateLocationsUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateLocationsUpdateHandlerFunc) Handle(params WeaviateLocationsUpdateParams) middleware.Responder {
	return fn(params)
}

// WeaviateLocationsUpdateHandler interface for that can handle valid weaviate locations update params
type WeaviateLocationsUpdateHandler interface {
	Handle(WeaviateLocationsUpdateParams) middleware.Responder
}

// NewWeaviateLocationsUpdate creates a new http.Handler for the weaviate locations update operation
func NewWeaviateLocationsUpdate(ctx *middleware.Context, handler WeaviateLocationsUpdateHandler) *WeaviateLocationsUpdate {
	return &WeaviateLocationsUpdate{Context: ctx, Handler: handler}
}

/*WeaviateLocationsUpdate swagger:route PUT /locations/{locationId} locations weaviateLocationsUpdate

Updates an location.

*/
type WeaviateLocationsUpdate struct {
	Context *middleware.Context
	Handler WeaviateLocationsUpdateHandler
}

func (o *WeaviateLocationsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateLocationsUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
