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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateActionsListParams creates a new WeaviateActionsListParams object
// with the default values initialized.
func NewWeaviateActionsListParams() *WeaviateActionsListParams {
	var ()
	return &WeaviateActionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsListParamsWithTimeout creates a new WeaviateActionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsListParamsWithTimeout(timeout time.Duration) *WeaviateActionsListParams {
	var ()
	return &WeaviateActionsListParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsListParamsWithContext creates a new WeaviateActionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsListParamsWithContext(ctx context.Context) *WeaviateActionsListParams {
	var ()
	return &WeaviateActionsListParams{

		Context: ctx,
	}
}

// NewWeaviateActionsListParamsWithHTTPClient creates a new WeaviateActionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsListParamsWithHTTPClient(client *http.Client) *WeaviateActionsListParams {
	var ()
	return &WeaviateActionsListParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsListParams contains all the parameters to send to the API endpoint
for the weaviate actions list operation typically these are written to a http.Request
*/
type WeaviateActionsListParams struct {

	/*Limit
	  The maximum number of items to be returned per page. Default value is set in Weaviate config.

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions list params
func (o *WeaviateActionsListParams) WithTimeout(timeout time.Duration) *WeaviateActionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions list params
func (o *WeaviateActionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions list params
func (o *WeaviateActionsListParams) WithContext(ctx context.Context) *WeaviateActionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions list params
func (o *WeaviateActionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions list params
func (o *WeaviateActionsListParams) WithHTTPClient(client *http.Client) *WeaviateActionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions list params
func (o *WeaviateActionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the weaviate actions list params
func (o *WeaviateActionsListParams) WithLimit(limit *int64) *WeaviateActionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the weaviate actions list params
func (o *WeaviateActionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
