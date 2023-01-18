//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// BatchObjectsDeleteHandlerFunc turns a function with the right signature into a batch objects delete handler
type BatchObjectsDeleteHandlerFunc func(BatchObjectsDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BatchObjectsDeleteHandlerFunc) Handle(params BatchObjectsDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BatchObjectsDeleteHandler interface for that can handle valid batch objects delete params
type BatchObjectsDeleteHandler interface {
	Handle(BatchObjectsDeleteParams, *models.Principal) middleware.Responder
}

// NewBatchObjectsDelete creates a new http.Handler for the batch objects delete operation
func NewBatchObjectsDelete(ctx *middleware.Context, handler BatchObjectsDeleteHandler) *BatchObjectsDelete {
	return &BatchObjectsDelete{Context: ctx, Handler: handler}
}

/*
BatchObjectsDelete swagger:route DELETE /batch/objects batch objects batchObjectsDelete

Deletes Objects based on a match filter as a batch.

Delete Objects in bulk that match a certain filter.
*/
type BatchObjectsDelete struct {
	Context *middleware.Context
	Handler BatchObjectsDeleteHandler
}

func (o *BatchObjectsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBatchObjectsDeleteParams()

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
