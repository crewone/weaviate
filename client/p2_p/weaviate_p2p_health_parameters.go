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

package p2_p

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

// NewWeaviateP2pHealthParams creates a new WeaviateP2pHealthParams object
// with the default values initialized.
func NewWeaviateP2pHealthParams() *WeaviateP2pHealthParams {

	return &WeaviateP2pHealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateP2pHealthParamsWithTimeout creates a new WeaviateP2pHealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateP2pHealthParamsWithTimeout(timeout time.Duration) *WeaviateP2pHealthParams {

	return &WeaviateP2pHealthParams{

		timeout: timeout,
	}
}

// NewWeaviateP2pHealthParamsWithContext creates a new WeaviateP2pHealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateP2pHealthParamsWithContext(ctx context.Context) *WeaviateP2pHealthParams {

	return &WeaviateP2pHealthParams{

		Context: ctx,
	}
}

// NewWeaviateP2pHealthParamsWithHTTPClient creates a new WeaviateP2pHealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateP2pHealthParamsWithHTTPClient(client *http.Client) *WeaviateP2pHealthParams {

	return &WeaviateP2pHealthParams{
		HTTPClient: client,
	}
}

/*WeaviateP2pHealthParams contains all the parameters to send to the API endpoint
for the weaviate p2p health operation typically these are written to a http.Request
*/
type WeaviateP2pHealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) WithTimeout(timeout time.Duration) *WeaviateP2pHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) WithContext(ctx context.Context) *WeaviateP2pHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) WithHTTPClient(client *http.Client) *WeaviateP2pHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate p2p health params
func (o *WeaviateP2pHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateP2pHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
