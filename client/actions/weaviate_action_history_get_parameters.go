// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateActionHistoryGetParams creates a new WeaviateActionHistoryGetParams object
// with the default values initialized.
func NewWeaviateActionHistoryGetParams() *WeaviateActionHistoryGetParams {
	var ()
	return &WeaviateActionHistoryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionHistoryGetParamsWithTimeout creates a new WeaviateActionHistoryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionHistoryGetParamsWithTimeout(timeout time.Duration) *WeaviateActionHistoryGetParams {
	var ()
	return &WeaviateActionHistoryGetParams{

		timeout: timeout,
	}
}

// NewWeaviateActionHistoryGetParamsWithContext creates a new WeaviateActionHistoryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionHistoryGetParamsWithContext(ctx context.Context) *WeaviateActionHistoryGetParams {
	var ()
	return &WeaviateActionHistoryGetParams{

		Context: ctx,
	}
}

// NewWeaviateActionHistoryGetParamsWithHTTPClient creates a new WeaviateActionHistoryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionHistoryGetParamsWithHTTPClient(client *http.Client) *WeaviateActionHistoryGetParams {
	var ()
	return &WeaviateActionHistoryGetParams{
		HTTPClient: client,
	}
}

/*WeaviateActionHistoryGetParams contains all the parameters to send to the API endpoint
for the weaviate action history get operation typically these are written to a http.Request
*/
type WeaviateActionHistoryGetParams struct {

	/*ActionID
	  Unique ID of the Action.

	*/
	ActionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) WithTimeout(timeout time.Duration) *WeaviateActionHistoryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) WithContext(ctx context.Context) *WeaviateActionHistoryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) WithHTTPClient(client *http.Client) *WeaviateActionHistoryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) WithActionID(actionID strfmt.UUID) *WeaviateActionHistoryGetParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate action history get params
func (o *WeaviateActionHistoryGetParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionHistoryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
