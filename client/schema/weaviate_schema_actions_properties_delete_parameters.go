/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

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
)

// NewWeaviateSchemaActionsPropertiesDeleteParams creates a new WeaviateSchemaActionsPropertiesDeleteParams object
// with the default values initialized.
func NewWeaviateSchemaActionsPropertiesDeleteParams() *WeaviateSchemaActionsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaActionsPropertiesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaActionsPropertiesDeleteParamsWithTimeout creates a new WeaviateSchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaActionsPropertiesDeleteParamsWithTimeout(timeout time.Duration) *WeaviateSchemaActionsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaActionsPropertiesDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaActionsPropertiesDeleteParamsWithContext creates a new WeaviateSchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaActionsPropertiesDeleteParamsWithContext(ctx context.Context) *WeaviateSchemaActionsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaActionsPropertiesDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaActionsPropertiesDeleteParamsWithHTTPClient creates a new WeaviateSchemaActionsPropertiesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaActionsPropertiesDeleteParamsWithHTTPClient(client *http.Client) *WeaviateSchemaActionsPropertiesDeleteParams {
	var ()
	return &WeaviateSchemaActionsPropertiesDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaActionsPropertiesDeleteParams contains all the parameters to send to the API endpoint
for the weaviate schema actions properties delete operation typically these are written to a http.Request
*/
type WeaviateSchemaActionsPropertiesDeleteParams struct {

	/*ClassName*/
	ClassName string
	/*PropertyName*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WithTimeout(timeout time.Duration) *WeaviateSchemaActionsPropertiesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WithContext(ctx context.Context) *WeaviateSchemaActionsPropertiesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WithHTTPClient(client *http.Client) *WeaviateSchemaActionsPropertiesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WithClassName(className string) *WeaviateSchemaActionsPropertiesDeleteParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) SetClassName(className string) {
	o.ClassName = className
}

// WithPropertyName adds the propertyName to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WithPropertyName(propertyName string) *WeaviateSchemaActionsPropertiesDeleteParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate schema actions properties delete params
func (o *WeaviateSchemaActionsPropertiesDeleteParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaActionsPropertiesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
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
