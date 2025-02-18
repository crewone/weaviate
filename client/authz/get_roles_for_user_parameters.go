//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

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

// NewGetRolesForUserParams creates a new GetRolesForUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRolesForUserParams() *GetRolesForUserParams {
	return &GetRolesForUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolesForUserParamsWithTimeout creates a new GetRolesForUserParams object
// with the ability to set a timeout on a request.
func NewGetRolesForUserParamsWithTimeout(timeout time.Duration) *GetRolesForUserParams {
	return &GetRolesForUserParams{
		timeout: timeout,
	}
}

// NewGetRolesForUserParamsWithContext creates a new GetRolesForUserParams object
// with the ability to set a context for a request.
func NewGetRolesForUserParamsWithContext(ctx context.Context) *GetRolesForUserParams {
	return &GetRolesForUserParams{
		Context: ctx,
	}
}

// NewGetRolesForUserParamsWithHTTPClient creates a new GetRolesForUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRolesForUserParamsWithHTTPClient(client *http.Client) *GetRolesForUserParams {
	return &GetRolesForUserParams{
		HTTPClient: client,
	}
}

/*
GetRolesForUserParams contains all the parameters to send to the API endpoint

	for the get roles for user operation.

	Typically these are written to a http.Request.
*/
type GetRolesForUserParams struct {

	/* ID.

	   The name of the user.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get roles for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesForUserParams) WithDefaults() *GetRolesForUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get roles for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolesForUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get roles for user params
func (o *GetRolesForUserParams) WithTimeout(timeout time.Duration) *GetRolesForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get roles for user params
func (o *GetRolesForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get roles for user params
func (o *GetRolesForUserParams) WithContext(ctx context.Context) *GetRolesForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get roles for user params
func (o *GetRolesForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get roles for user params
func (o *GetRolesForUserParams) WithHTTPClient(client *http.Client) *GetRolesForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get roles for user params
func (o *GetRolesForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get roles for user params
func (o *GetRolesForUserParams) WithID(id string) *GetRolesForUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get roles for user params
func (o *GetRolesForUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolesForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
