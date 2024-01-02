//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// BackupsCreateHandlerFunc turns a function with the right signature into a backups create handler
type BackupsCreateHandlerFunc func(BackupsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupsCreateHandlerFunc) Handle(params BackupsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupsCreateHandler interface for that can handle valid backups create params
type BackupsCreateHandler interface {
	Handle(BackupsCreateParams, *models.Principal) middleware.Responder
}

// NewBackupsCreate creates a new http.Handler for the backups create operation
func NewBackupsCreate(ctx *middleware.Context, handler BackupsCreateHandler) *BackupsCreate {
	return &BackupsCreate{Context: ctx, Handler: handler}
}

/*
	BackupsCreate swagger:route POST /backups/{backend} backups backupsCreate

Starts a process of creating a backup for a set of classes
*/
type BackupsCreate struct {
	Context *middleware.Context
	Handler BackupsCreateHandler
}

func (o *BackupsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBackupsCreateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
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
