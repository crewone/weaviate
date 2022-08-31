//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

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

// NewBackupsRestoreParams creates a new BackupsRestoreParams object
// with the default values initialized.
func NewBackupsRestoreParams() *BackupsRestoreParams {
	var ()
	return &BackupsRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupsRestoreParamsWithTimeout creates a new BackupsRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupsRestoreParamsWithTimeout(timeout time.Duration) *BackupsRestoreParams {
	var ()
	return &BackupsRestoreParams{

		timeout: timeout,
	}
}

// NewBackupsRestoreParamsWithContext creates a new BackupsRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupsRestoreParamsWithContext(ctx context.Context) *BackupsRestoreParams {
	var ()
	return &BackupsRestoreParams{

		Context: ctx,
	}
}

// NewBackupsRestoreParamsWithHTTPClient creates a new BackupsRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupsRestoreParamsWithHTTPClient(client *http.Client) *BackupsRestoreParams {
	var ()
	return &BackupsRestoreParams{
		HTTPClient: client,
	}
}

/*BackupsRestoreParams contains all the parameters to send to the API endpoint
for the backups restore operation typically these are written to a http.Request
*/
type BackupsRestoreParams struct {

	/*Body*/
	Body *models.BackupRestoreRequest
	/*ID
	  The ID of a backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.

	*/
	ID string
	/*StorageName
	  Storage name e.g. filesystem, gcs, s3.

	*/
	StorageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backups restore params
func (o *BackupsRestoreParams) WithTimeout(timeout time.Duration) *BackupsRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backups restore params
func (o *BackupsRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backups restore params
func (o *BackupsRestoreParams) WithContext(ctx context.Context) *BackupsRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backups restore params
func (o *BackupsRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backups restore params
func (o *BackupsRestoreParams) WithHTTPClient(client *http.Client) *BackupsRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backups restore params
func (o *BackupsRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the backups restore params
func (o *BackupsRestoreParams) WithBody(body *models.BackupRestoreRequest) *BackupsRestoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backups restore params
func (o *BackupsRestoreParams) SetBody(body *models.BackupRestoreRequest) {
	o.Body = body
}

// WithID adds the id to the backups restore params
func (o *BackupsRestoreParams) WithID(id string) *BackupsRestoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the backups restore params
func (o *BackupsRestoreParams) SetID(id string) {
	o.ID = id
}

// WithStorageName adds the storageName to the backups restore params
func (o *BackupsRestoreParams) WithStorageName(storageName string) *BackupsRestoreParams {
	o.SetStorageName(storageName)
	return o
}

// SetStorageName adds the storageName to the backups restore params
func (o *BackupsRestoreParams) SetStorageName(storageName string) {
	o.StorageName = storageName
}

// WriteToRequest writes these params to a swagger request
func (o *BackupsRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param storageName
	if err := r.SetPathParam("storageName", o.StorageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
