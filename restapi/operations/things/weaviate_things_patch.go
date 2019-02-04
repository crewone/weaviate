/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateThingsPatchHandlerFunc turns a function with the right signature into a weaviate things patch handler
type WeaviateThingsPatchHandlerFunc func(context.Context, WeaviateThingsPatchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsPatchHandlerFunc) Handle(ctx context.Context, params WeaviateThingsPatchParams) middleware.Responder {
	return fn(ctx, params)
}

// WeaviateThingsPatchHandler interface for that can handle valid weaviate things patch params
type WeaviateThingsPatchHandler interface {
	Handle(context.Context, WeaviateThingsPatchParams) middleware.Responder
}

// NewWeaviateThingsPatch creates a new http.Handler for the weaviate things patch operation
func NewWeaviateThingsPatch(ctx *middleware.Context, handler WeaviateThingsPatchHandler) *WeaviateThingsPatch {
	return &WeaviateThingsPatch{Context: ctx, Handler: handler}
}

/*WeaviateThingsPatch swagger:route PATCH /things/{thingId} things weaviateThingsPatch

Update a Thing based on its UUID (using patch semantics) related to this key.

Updates a Thing's data. This method supports patch semantics. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.

*/
type WeaviateThingsPatch struct {
	Context *middleware.Context
	Handler WeaviateThingsPatchHandler
}

func (o *WeaviateThingsPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsPatchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(r.Context(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
