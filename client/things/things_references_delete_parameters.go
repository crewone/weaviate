//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// NewThingsReferencesDeleteParams creates a new ThingsReferencesDeleteParams object
// with the default values initialized.
func NewThingsReferencesDeleteParams() *ThingsReferencesDeleteParams {
	var ()
	return &ThingsReferencesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsReferencesDeleteParamsWithTimeout creates a new ThingsReferencesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsReferencesDeleteParamsWithTimeout(timeout time.Duration) *ThingsReferencesDeleteParams {
	var ()
	return &ThingsReferencesDeleteParams{

		timeout: timeout,
	}
}

// NewThingsReferencesDeleteParamsWithContext creates a new ThingsReferencesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsReferencesDeleteParamsWithContext(ctx context.Context) *ThingsReferencesDeleteParams {
	var ()
	return &ThingsReferencesDeleteParams{

		Context: ctx,
	}
}

// NewThingsReferencesDeleteParamsWithHTTPClient creates a new ThingsReferencesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsReferencesDeleteParamsWithHTTPClient(client *http.Client) *ThingsReferencesDeleteParams {
	var ()
	return &ThingsReferencesDeleteParams{
		HTTPClient: client,
	}
}

/*ThingsReferencesDeleteParams contains all the parameters to send to the API endpoint
for the things references delete operation typically these are written to a http.Request
*/
type ThingsReferencesDeleteParams struct {

	/*Body*/
	Body *models.SingleRef
	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Thing.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things references delete params
func (o *ThingsReferencesDeleteParams) WithTimeout(timeout time.Duration) *ThingsReferencesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things references delete params
func (o *ThingsReferencesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things references delete params
func (o *ThingsReferencesDeleteParams) WithContext(ctx context.Context) *ThingsReferencesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things references delete params
func (o *ThingsReferencesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things references delete params
func (o *ThingsReferencesDeleteParams) WithHTTPClient(client *http.Client) *ThingsReferencesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things references delete params
func (o *ThingsReferencesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the things references delete params
func (o *ThingsReferencesDeleteParams) WithBody(body *models.SingleRef) *ThingsReferencesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the things references delete params
func (o *ThingsReferencesDeleteParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithID adds the id to the things references delete params
func (o *ThingsReferencesDeleteParams) WithID(id strfmt.UUID) *ThingsReferencesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the things references delete params
func (o *ThingsReferencesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the things references delete params
func (o *ThingsReferencesDeleteParams) WithPropertyName(propertyName string) *ThingsReferencesDeleteParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the things references delete params
func (o *ThingsReferencesDeleteParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsReferencesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
