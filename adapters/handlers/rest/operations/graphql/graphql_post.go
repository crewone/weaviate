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

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// GraphqlPostHandlerFunc turns a function with the right signature into a graphql post handler
type GraphqlPostHandlerFunc func(GraphqlPostParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GraphqlPostHandlerFunc) Handle(params GraphqlPostParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GraphqlPostHandler interface for that can handle valid graphql post params
type GraphqlPostHandler interface {
	Handle(GraphqlPostParams, *models.Principal) middleware.Responder
}

// NewGraphqlPost creates a new http.Handler for the graphql post operation
func NewGraphqlPost(ctx *middleware.Context, handler GraphqlPostHandler) *GraphqlPost {
	return &GraphqlPost{Context: ctx, Handler: handler}
}

/*GraphqlPost swagger:route POST /graphql graphql graphqlPost

Get a response based on GraphQL

Get an object based on GraphQL

*/
type GraphqlPost struct {
	Context *middleware.Context
	Handler GraphqlPostHandler
}

func (o *GraphqlPost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGraphqlPostParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
