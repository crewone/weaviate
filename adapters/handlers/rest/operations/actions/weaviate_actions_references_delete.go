/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateActionsReferencesDeleteHandlerFunc turns a function with the right signature into a weaviate actions references delete handler
type WeaviateActionsReferencesDeleteHandlerFunc func(WeaviateActionsReferencesDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsReferencesDeleteHandlerFunc) Handle(params WeaviateActionsReferencesDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsReferencesDeleteHandler interface for that can handle valid weaviate actions references delete params
type WeaviateActionsReferencesDeleteHandler interface {
	Handle(WeaviateActionsReferencesDeleteParams, *models.Principal) middleware.Responder
}

// NewWeaviateActionsReferencesDelete creates a new http.Handler for the weaviate actions references delete operation
func NewWeaviateActionsReferencesDelete(ctx *middleware.Context, handler WeaviateActionsReferencesDeleteHandler) *WeaviateActionsReferencesDelete {
	return &WeaviateActionsReferencesDelete{Context: ctx, Handler: handler}
}

/*WeaviateActionsReferencesDelete swagger:route DELETE /actions/{id}/references/{propertyName} actions weaviateActionsReferencesDelete

Delete the single reference that is given in the body from the list of references that this property has.

Delete the single reference that is given in the body from the list of references that this property has.

*/
type WeaviateActionsReferencesDelete struct {
	Context *middleware.Context
	Handler WeaviateActionsReferencesDeleteHandler
}

func (o *WeaviateActionsReferencesDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsReferencesDeleteParams()

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
