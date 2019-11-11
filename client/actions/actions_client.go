//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new actions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for actions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ActionsCreate creates actions between two things object and subject

Registers a new Action. Provided meta-data and schema values are validated.
*/
func (a *Client) ActionsCreate(params *ActionsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.create",
		Method:             "POST",
		PathPattern:        "/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsCreateOK), nil

}

/*
ActionsDelete deletes an action based on its UUID

Deletes an Action from the system.
*/
func (a *Client) ActionsDelete(params *ActionsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.delete",
		Method:             "DELETE",
		PathPattern:        "/actions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsDeleteNoContent), nil

}

/*
ActionsGet gets a specific action based on its UUID and a thing UUID also available as websocket bus

Lists Actions.
*/
func (a *Client) ActionsGet(params *ActionsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.get",
		Method:             "GET",
		PathPattern:        "/actions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsGetOK), nil

}

/*
ActionsList gets a list of actions

Lists all Actions in reverse order of creation, owned by the user that belongs to the used token.
*/
func (a *Client) ActionsList(params *ActionsListParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.list",
		Method:             "GET",
		PathPattern:        "/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsListOK), nil

}

/*
ActionsPatch updates an action based on its UUID using patch semantics

Updates an Action. This method supports json-merge style patch semantics (RFC 7396). Provided meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) ActionsPatch(params *ActionsPatchParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.patch",
		Method:             "PATCH",
		PathPattern:        "/actions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsPatchOK), nil

}

/*
ActionsReferencesCreate adds a single reference to a class property when cardinality is set to has many

Add a single reference to a class-property when cardinality is set to 'hasMany'.
*/
func (a *Client) ActionsReferencesCreate(params *ActionsReferencesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsReferencesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsReferencesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.references.create",
		Method:             "POST",
		PathPattern:        "/actions/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsReferencesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsReferencesCreateOK), nil

}

/*
ActionsReferencesDelete deletes the single reference that is given in the body from the list of references that this property has

Delete the single reference that is given in the body from the list of references that this property has.
*/
func (a *Client) ActionsReferencesDelete(params *ActionsReferencesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsReferencesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsReferencesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.references.delete",
		Method:             "DELETE",
		PathPattern:        "/actions/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsReferencesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsReferencesDeleteNoContent), nil

}

/*
ActionsReferencesUpdate replaces all references to a class property

Replace all references to a class-property.
*/
func (a *Client) ActionsReferencesUpdate(params *ActionsReferencesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsReferencesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsReferencesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.references.update",
		Method:             "PUT",
		PathPattern:        "/actions/{id}/references/{propertyName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsReferencesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsReferencesUpdateOK), nil

}

/*
ActionsUpdate updates an action based on its UUID

Updates an Action's data. Given meta-data and schema values are validated. LastUpdateTime is set to the time this function is called.
*/
func (a *Client) ActionsUpdate(params *ActionsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.update",
		Method:             "PUT",
		PathPattern:        "/actions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsUpdateOK), nil

}

/*
ActionsValidate validates an action based on a schema

Validate an Action's schema and meta-data. It has to be based on a schema, which is related to the given Action to be accepted by this validation.
*/
func (a *Client) ActionsValidate(params *ActionsValidateParams, authInfo runtime.ClientAuthInfoWriter) (*ActionsValidateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionsValidateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actions.validate",
		Method:             "POST",
		PathPattern:        "/actions/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionsValidateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionsValidateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
