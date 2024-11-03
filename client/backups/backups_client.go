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

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BackupsCancel(params *BackupsCancelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCancelNoContent, error)

	BackupsCreate(params *BackupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCreateOK, error)

	BackupsCreateStatus(params *BackupsCreateStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCreateStatusOK, error)

	BackupsList(params *BackupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsListOK, error)

	BackupsRestore(params *BackupsRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsRestoreOK, error)

	BackupsRestoreStatus(params *BackupsRestoreStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsRestoreStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BackupsCancel Cancel created backup with specified ID
*/
func (a *Client) BackupsCancel(params *BackupsCancelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCancelNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsCancelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.cancel",
		Method:             "DELETE",
		PathPattern:        "/backups/{backend}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsCancelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsCancelNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.cancel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupsCreate Start creating a backup for a set of collections. <br/><br/>Notes: <br/>- Weaviate uses gzip compression by default. <br/>- Weaviate stays usable while a backup process is ongoing.
*/
func (a *Client) BackupsCreate(params *BackupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.create",
		Method:             "POST",
		PathPattern:        "/backups/{backend}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupsCreateStatus Returns status of backup creation attempt for a set of collections. <br/><br/>All client implementations have a `wait for completion` option which will poll the backup status in the background and only return once the backup has completed (successfully or unsuccessfully). If you set the `wait for completion` option to false, you can also check the status yourself using this endpoint.
*/
func (a *Client) BackupsCreateStatus(params *BackupsCreateStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsCreateStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsCreateStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.create.status",
		Method:             "GET",
		PathPattern:        "/backups/{backend}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsCreateStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsCreateStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.create.status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupsList [Coming soon] List all backups in progress not implemented yet.
*/
func (a *Client) BackupsList(params *BackupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.list",
		Method:             "GET",
		PathPattern:        "/backups/{backend}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupsRestore Starts a process of restoring a backup for a set of collections. <br/><br/>Any backup can be restored to any machine, as long as the number of nodes between source and target are identical.<br/><br/>Requrements:<br/><br/>- None of the collections to be restored already exist on the target restoration node(s).<br/>- The node names of the backed-up collections' must match those of the target restoration node(s).
*/
func (a *Client) BackupsRestore(params *BackupsRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.restore",
		Method:             "POST",
		PathPattern:        "/backups/{backend}/{id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsRestoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.restore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupsRestoreStatus Returns status of a backup restoration attempt for a set of classes. <br/><br/>All client implementations have a `wait for completion` option which will poll the backup status in the background and only return once the backup has completed (successfully or unsuccessfully). If you set the `wait for completion` option to false, you can also check the status yourself using the this endpoint.
*/
func (a *Client) BackupsRestoreStatus(params *BackupsRestoreStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupsRestoreStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupsRestoreStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "backups.restore.status",
		Method:             "GET",
		PathPattern:        "/backups/{backend}/{id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BackupsRestoreStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BackupsRestoreStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for backups.restore.status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
