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

// NewSchemaDumpParams creates a new SchemaDumpParams object
// with the default values initialized.
func NewSchemaDumpParams() *SchemaDumpParams {

	return &SchemaDumpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaDumpParamsWithTimeout creates a new SchemaDumpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaDumpParamsWithTimeout(timeout time.Duration) *SchemaDumpParams {

	return &SchemaDumpParams{

		timeout: timeout,
	}
}

// NewSchemaDumpParamsWithContext creates a new SchemaDumpParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaDumpParamsWithContext(ctx context.Context) *SchemaDumpParams {

	return &SchemaDumpParams{

		Context: ctx,
	}
}

// NewSchemaDumpParamsWithHTTPClient creates a new SchemaDumpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaDumpParamsWithHTTPClient(client *http.Client) *SchemaDumpParams {

	return &SchemaDumpParams{
		HTTPClient: client,
	}
}

/*SchemaDumpParams contains all the parameters to send to the API endpoint
for the schema dump operation typically these are written to a http.Request
*/
type SchemaDumpParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema dump params
func (o *SchemaDumpParams) WithTimeout(timeout time.Duration) *SchemaDumpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema dump params
func (o *SchemaDumpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema dump params
func (o *SchemaDumpParams) WithContext(ctx context.Context) *SchemaDumpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema dump params
func (o *SchemaDumpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema dump params
func (o *SchemaDumpParams) WithHTTPClient(client *http.Client) *SchemaDumpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema dump params
func (o *SchemaDumpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaDumpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
