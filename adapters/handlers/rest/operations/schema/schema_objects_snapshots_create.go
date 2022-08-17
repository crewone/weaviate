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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsSnapshotsCreateHandlerFunc turns a function with the right signature into a schema objects snapshots create handler
type SchemaObjectsSnapshotsCreateHandlerFunc func(SchemaObjectsSnapshotsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SchemaObjectsSnapshotsCreateHandlerFunc) Handle(params SchemaObjectsSnapshotsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SchemaObjectsSnapshotsCreateHandler interface for that can handle valid schema objects snapshots create params
type SchemaObjectsSnapshotsCreateHandler interface {
	Handle(SchemaObjectsSnapshotsCreateParams, *models.Principal) middleware.Responder
}

// NewSchemaObjectsSnapshotsCreate creates a new http.Handler for the schema objects snapshots create operation
func NewSchemaObjectsSnapshotsCreate(ctx *middleware.Context, handler SchemaObjectsSnapshotsCreateHandler) *SchemaObjectsSnapshotsCreate {
	return &SchemaObjectsSnapshotsCreate{Context: ctx, Handler: handler}
}

/*SchemaObjectsSnapshotsCreate swagger:route POST /schema/{className}/snapshots/{storageName}/{id} schema schemaObjectsSnapshotsCreate

Starts a process of creation a snapshot for a class

*/
type SchemaObjectsSnapshotsCreate struct {
	Context *middleware.Context
	Handler SchemaObjectsSnapshotsCreateHandler
}

func (o *SchemaObjectsSnapshotsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSchemaObjectsSnapshotsCreateParams()

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
