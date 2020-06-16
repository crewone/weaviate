//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batching

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
)

// NewBatchingThingsCreateParams creates a new BatchingThingsCreateParams object
// with the default values initialized.
func NewBatchingThingsCreateParams() *BatchingThingsCreateParams {
	var ()
	return &BatchingThingsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBatchingThingsCreateParamsWithTimeout creates a new BatchingThingsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBatchingThingsCreateParamsWithTimeout(timeout time.Duration) *BatchingThingsCreateParams {
	var ()
	return &BatchingThingsCreateParams{

		timeout: timeout,
	}
}

// NewBatchingThingsCreateParamsWithContext creates a new BatchingThingsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBatchingThingsCreateParamsWithContext(ctx context.Context) *BatchingThingsCreateParams {
	var ()
	return &BatchingThingsCreateParams{

		Context: ctx,
	}
}

// NewBatchingThingsCreateParamsWithHTTPClient creates a new BatchingThingsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBatchingThingsCreateParamsWithHTTPClient(client *http.Client) *BatchingThingsCreateParams {
	var ()
	return &BatchingThingsCreateParams{
		HTTPClient: client,
	}
}

/*BatchingThingsCreateParams contains all the parameters to send to the API endpoint
for the batching things create operation typically these are written to a http.Request
*/
type BatchingThingsCreateParams struct {

	/*Body*/
	Body BatchingThingsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the batching things create params
func (o *BatchingThingsCreateParams) WithTimeout(timeout time.Duration) *BatchingThingsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the batching things create params
func (o *BatchingThingsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the batching things create params
func (o *BatchingThingsCreateParams) WithContext(ctx context.Context) *BatchingThingsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the batching things create params
func (o *BatchingThingsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the batching things create params
func (o *BatchingThingsCreateParams) WithHTTPClient(client *http.Client) *BatchingThingsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the batching things create params
func (o *BatchingThingsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the batching things create params
func (o *BatchingThingsCreateParams) WithBody(body BatchingThingsCreateBody) *BatchingThingsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the batching things create params
func (o *BatchingThingsCreateParams) SetBody(body BatchingThingsCreateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BatchingThingsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
