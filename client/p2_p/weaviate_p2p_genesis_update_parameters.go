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

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// NewWeaviateP2pGenesisUpdateParams creates a new WeaviateP2pGenesisUpdateParams object
// with the default values initialized.
func NewWeaviateP2pGenesisUpdateParams() *WeaviateP2pGenesisUpdateParams {
	var ()
	return &WeaviateP2pGenesisUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateP2pGenesisUpdateParamsWithTimeout creates a new WeaviateP2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateP2pGenesisUpdateParamsWithTimeout(timeout time.Duration) *WeaviateP2pGenesisUpdateParams {
	var ()
	return &WeaviateP2pGenesisUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateP2pGenesisUpdateParamsWithContext creates a new WeaviateP2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateP2pGenesisUpdateParamsWithContext(ctx context.Context) *WeaviateP2pGenesisUpdateParams {
	var ()
	return &WeaviateP2pGenesisUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateP2pGenesisUpdateParamsWithHTTPClient creates a new WeaviateP2pGenesisUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateP2pGenesisUpdateParamsWithHTTPClient(client *http.Client) *WeaviateP2pGenesisUpdateParams {
	var ()
	return &WeaviateP2pGenesisUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateP2pGenesisUpdateParams contains all the parameters to send to the API endpoint
for the weaviate p2p genesis update operation typically these are written to a http.Request
*/
type WeaviateP2pGenesisUpdateParams struct {

	/*Peers*/
	Peers models.PeerUpdateList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) WithTimeout(timeout time.Duration) *WeaviateP2pGenesisUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) WithContext(ctx context.Context) *WeaviateP2pGenesisUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) WithHTTPClient(client *http.Client) *WeaviateP2pGenesisUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeers adds the peers to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) WithPeers(peers models.PeerUpdateList) *WeaviateP2pGenesisUpdateParams {
	o.SetPeers(peers)
	return o
}

// SetPeers adds the peers to the weaviate p2p genesis update params
func (o *WeaviateP2pGenesisUpdateParams) SetPeers(peers models.PeerUpdateList) {
	o.Peers = peers
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateP2pGenesisUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Peers != nil {
		if err := r.SetBodyParam(o.Peers); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
