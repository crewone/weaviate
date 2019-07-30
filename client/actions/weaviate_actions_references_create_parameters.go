//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewWeaviateActionsReferencesCreateParams creates a new WeaviateActionsReferencesCreateParams object
// with the default values initialized.
func NewWeaviateActionsReferencesCreateParams() *WeaviateActionsReferencesCreateParams {
	var ()
	return &WeaviateActionsReferencesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsReferencesCreateParamsWithTimeout creates a new WeaviateActionsReferencesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsReferencesCreateParamsWithTimeout(timeout time.Duration) *WeaviateActionsReferencesCreateParams {
	var ()
	return &WeaviateActionsReferencesCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsReferencesCreateParamsWithContext creates a new WeaviateActionsReferencesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsReferencesCreateParamsWithContext(ctx context.Context) *WeaviateActionsReferencesCreateParams {
	var ()
	return &WeaviateActionsReferencesCreateParams{

		Context: ctx,
	}
}

// NewWeaviateActionsReferencesCreateParamsWithHTTPClient creates a new WeaviateActionsReferencesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsReferencesCreateParamsWithHTTPClient(client *http.Client) *WeaviateActionsReferencesCreateParams {
	var ()
	return &WeaviateActionsReferencesCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsReferencesCreateParams contains all the parameters to send to the API endpoint
for the weaviate actions references create operation typically these are written to a http.Request
*/
type WeaviateActionsReferencesCreateParams struct {

	/*Body*/
	Body *models.SingleRef
	/*ID
	  Unique ID of the Action.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Action.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithTimeout(timeout time.Duration) *WeaviateActionsReferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithContext(ctx context.Context) *WeaviateActionsReferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithHTTPClient(client *http.Client) *WeaviateActionsReferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithBody(body *models.SingleRef) *WeaviateActionsReferencesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithID adds the id to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithID(id strfmt.UUID) *WeaviateActionsReferencesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) WithPropertyName(propertyName string) *WeaviateActionsReferencesCreateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate actions references create params
func (o *WeaviateActionsReferencesCreateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsReferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
