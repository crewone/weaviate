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
 package acl_entries


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateACLEntriesListHandlerFunc turns a function with the right signature into a weaviate acl entries list handler
type WeaviateACLEntriesListHandlerFunc func(WeaviateACLEntriesListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateACLEntriesListHandlerFunc) Handle(params WeaviateACLEntriesListParams) middleware.Responder {
	return fn(params)
}

// WeaviateACLEntriesListHandler interface for that can handle valid weaviate acl entries list params
type WeaviateACLEntriesListHandler interface {
	Handle(WeaviateACLEntriesListParams) middleware.Responder
}

// NewWeaviateACLEntriesList creates a new http.Handler for the weaviate acl entries list operation
func NewWeaviateACLEntriesList(ctx *middleware.Context, handler WeaviateACLEntriesListHandler) *WeaviateACLEntriesList {
	return &WeaviateACLEntriesList{Context: ctx, Handler: handler}
}

/*WeaviateACLEntriesList swagger:route GET /devices/{deviceId}/aclEntries aclEntries weaviateAclEntriesList

Lists ACL entries.

*/
type WeaviateACLEntriesList struct {
	Context *middleware.Context
	Handler WeaviateACLEntriesListHandler
}

func (o *WeaviateACLEntriesList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateACLEntriesListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
