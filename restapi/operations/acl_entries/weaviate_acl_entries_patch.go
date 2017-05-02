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

// WeaviateACLEntriesPatchHandlerFunc turns a function with the right signature into a weaviate acl entries patch handler
type WeaviateACLEntriesPatchHandlerFunc func(WeaviateACLEntriesPatchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateACLEntriesPatchHandlerFunc) Handle(params WeaviateACLEntriesPatchParams) middleware.Responder {
	return fn(params)
}

// WeaviateACLEntriesPatchHandler interface for that can handle valid weaviate acl entries patch params
type WeaviateACLEntriesPatchHandler interface {
	Handle(WeaviateACLEntriesPatchParams) middleware.Responder
}

// NewWeaviateACLEntriesPatch creates a new http.Handler for the weaviate acl entries patch operation
func NewWeaviateACLEntriesPatch(ctx *middleware.Context, handler WeaviateACLEntriesPatchHandler) *WeaviateACLEntriesPatch {
	return &WeaviateACLEntriesPatch{Context: ctx, Handler: handler}
}

/*WeaviateACLEntriesPatch swagger:route PATCH /devices/{deviceId}/aclEntries/{aclEntryId} aclEntries weaviateAclEntriesPatch

Updates an ACL entry. This method supports patch semantics.

*/
type WeaviateACLEntriesPatch struct {
	Context *middleware.Context
	Handler WeaviateACLEntriesPatchHandler
}

func (o *WeaviateACLEntriesPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateACLEntriesPatchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
