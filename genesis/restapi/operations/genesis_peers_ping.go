//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GenesisPeersPingHandlerFunc turns a function with the right signature into a genesis peers ping handler
type GenesisPeersPingHandlerFunc func(GenesisPeersPingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GenesisPeersPingHandlerFunc) Handle(params GenesisPeersPingParams) middleware.Responder {
	return fn(params)
}

// GenesisPeersPingHandler interface for that can handle valid genesis peers ping params
type GenesisPeersPingHandler interface {
	Handle(GenesisPeersPingParams) middleware.Responder
}

// NewGenesisPeersPing creates a new http.Handler for the genesis peers ping operation
func NewGenesisPeersPing(ctx *middleware.Context, handler GenesisPeersPingHandler) *GenesisPeersPing {
	return &GenesisPeersPing{Context: ctx, Handler: handler}
}

/*GenesisPeersPing swagger:route POST /peers/{peerId}/ping genesisPeersPing

Ping the Genesis server, to make mark the peer as alive and udpate schema info

*/
type GenesisPeersPing struct {
	Context *middleware.Context
	Handler GenesisPeersPingHandler
}

func (o *GenesisPeersPing) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGenesisPeersPingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
