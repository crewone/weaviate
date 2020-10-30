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

package contextionary_api

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

// NewC11yWordsParams creates a new C11yWordsParams object
// with the default values initialized.
func NewC11yWordsParams() *C11yWordsParams {
	var ()
	return &C11yWordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewC11yWordsParamsWithTimeout creates a new C11yWordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewC11yWordsParamsWithTimeout(timeout time.Duration) *C11yWordsParams {
	var ()
	return &C11yWordsParams{

		timeout: timeout,
	}
}

// NewC11yWordsParamsWithContext creates a new C11yWordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewC11yWordsParamsWithContext(ctx context.Context) *C11yWordsParams {
	var ()
	return &C11yWordsParams{

		Context: ctx,
	}
}

// NewC11yWordsParamsWithHTTPClient creates a new C11yWordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewC11yWordsParamsWithHTTPClient(client *http.Client) *C11yWordsParams {
	var ()
	return &C11yWordsParams{
		HTTPClient: client,
	}
}

/*C11yWordsParams contains all the parameters to send to the API endpoint
for the c11y words operation typically these are written to a http.Request
*/
type C11yWordsParams struct {

	/*Words
	  CamelCase list of words to validate.

	*/
	Words string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the c11y words params
func (o *C11yWordsParams) WithTimeout(timeout time.Duration) *C11yWordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the c11y words params
func (o *C11yWordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the c11y words params
func (o *C11yWordsParams) WithContext(ctx context.Context) *C11yWordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the c11y words params
func (o *C11yWordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the c11y words params
func (o *C11yWordsParams) WithHTTPClient(client *http.Client) *C11yWordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the c11y words params
func (o *C11yWordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWords adds the words to the c11y words params
func (o *C11yWordsParams) WithWords(words string) *C11yWordsParams {
	o.SetWords(words)
	return o
}

// SetWords adds the words to the c11y words params
func (o *C11yWordsParams) SetWords(words string) {
	o.Words = words
}

// WriteToRequest writes these params to a swagger request
func (o *C11yWordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param words
	if err := r.SetPathParam("words", o.Words); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
