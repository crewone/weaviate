/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// NewWeaviateBatchingReferencesCreateParams creates a new WeaviateBatchingReferencesCreateParams object
// with the default values initialized.
func NewWeaviateBatchingReferencesCreateParams() *WeaviateBatchingReferencesCreateParams {
	var ()
	return &WeaviateBatchingReferencesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateBatchingReferencesCreateParamsWithTimeout creates a new WeaviateBatchingReferencesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateBatchingReferencesCreateParamsWithTimeout(timeout time.Duration) *WeaviateBatchingReferencesCreateParams {
	var ()
	return &WeaviateBatchingReferencesCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateBatchingReferencesCreateParamsWithContext creates a new WeaviateBatchingReferencesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateBatchingReferencesCreateParamsWithContext(ctx context.Context) *WeaviateBatchingReferencesCreateParams {
	var ()
	return &WeaviateBatchingReferencesCreateParams{

		Context: ctx,
	}
}

// NewWeaviateBatchingReferencesCreateParamsWithHTTPClient creates a new WeaviateBatchingReferencesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateBatchingReferencesCreateParamsWithHTTPClient(client *http.Client) *WeaviateBatchingReferencesCreateParams {
	var ()
	return &WeaviateBatchingReferencesCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateBatchingReferencesCreateParams contains all the parameters to send to the API endpoint
for the weaviate batching references create operation typically these are written to a http.Request
*/
type WeaviateBatchingReferencesCreateParams struct {

	/*Body
	  A list of references to be batched. The ideal size depends on the used database connector. Please see the documentation of the used connector for help

	*/
	Body []*models.BatchReference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) WithTimeout(timeout time.Duration) *WeaviateBatchingReferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) WithContext(ctx context.Context) *WeaviateBatchingReferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) WithHTTPClient(client *http.Client) *WeaviateBatchingReferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) WithBody(body []*models.BatchReference) *WeaviateBatchingReferencesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate batching references create params
func (o *WeaviateBatchingReferencesCreateParams) SetBody(body []*models.BatchReference) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateBatchingReferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
