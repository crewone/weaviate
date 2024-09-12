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

// BackupsListHandlerFunc turns a function with the right signature into a backups list handler
type BackupsListHandlerFunc func(BackupsListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupsListHandlerFunc) Handle(params BackupsListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupsListHandler interface for that can handle valid backups list params
type BackupsListHandler interface {
	Handle(BackupsListParams, *models.Principal) middleware.Responder
}

// NewBackupsList creates a new http.Handler for the backups list operation
func NewBackupsList(ctx *middleware.Context, handler BackupsListHandler) *BackupsList {
	return &BackupsList{Context: ctx, Handler: handler}
}

/*
	BackupsList swagger:route GET /backups/{backend} backups backupsList

List all backups in progress
*/
type BackupsList struct {
	Context *middleware.Context
	Handler BackupsListHandler
}

func (o *BackupsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBackupsListParams()
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
