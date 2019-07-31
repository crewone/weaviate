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

package schema

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

// NewSchemaActionsUpdateParams creates a new SchemaActionsUpdateParams object
// with the default values initialized.
func NewSchemaActionsUpdateParams() *SchemaActionsUpdateParams {
	var ()
	return &SchemaActionsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaActionsUpdateParamsWithTimeout creates a new SchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaActionsUpdateParamsWithTimeout(timeout time.Duration) *SchemaActionsUpdateParams {
	var ()
	return &SchemaActionsUpdateParams{

		timeout: timeout,
	}
}

// NewSchemaActionsUpdateParamsWithContext creates a new SchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaActionsUpdateParamsWithContext(ctx context.Context) *SchemaActionsUpdateParams {
	var ()
	return &SchemaActionsUpdateParams{

		Context: ctx,
	}
}

// NewSchemaActionsUpdateParamsWithHTTPClient creates a new SchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaActionsUpdateParamsWithHTTPClient(client *http.Client) *SchemaActionsUpdateParams {
	var ()
	return &SchemaActionsUpdateParams{
		HTTPClient: client,
	}
}

/*SchemaActionsUpdateParams contains all the parameters to send to the API endpoint
for the schema actions update operation typically these are written to a http.Request
*/
type SchemaActionsUpdateParams struct {

	/*Body*/
	Body *models.Class
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema actions update params
func (o *SchemaActionsUpdateParams) WithTimeout(timeout time.Duration) *SchemaActionsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema actions update params
func (o *SchemaActionsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema actions update params
func (o *SchemaActionsUpdateParams) WithContext(ctx context.Context) *SchemaActionsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema actions update params
func (o *SchemaActionsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema actions update params
func (o *SchemaActionsUpdateParams) WithHTTPClient(client *http.Client) *SchemaActionsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema actions update params
func (o *SchemaActionsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema actions update params
func (o *SchemaActionsUpdateParams) WithBody(body *models.Class) *SchemaActionsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema actions update params
func (o *SchemaActionsUpdateParams) SetBody(body *models.Class) {
	o.Body = body
}

// WithClassName adds the className to the schema actions update params
func (o *SchemaActionsUpdateParams) WithClassName(className string) *SchemaActionsUpdateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema actions update params
func (o *SchemaActionsUpdateParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaActionsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
