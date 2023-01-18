//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupRestoreRequest Request body for restoring a backup for a set of classes
//
// swagger:model BackupRestoreRequest
type BackupRestoreRequest struct {

	// Custom configuration for the backup restoration process
	Config interface{} `json:"config,omitempty"`

	// List of classes to exclude from the backup restoration process
	Exclude []string `json:"exclude"`

	// List of classes to include in the backup restoration process
	Include []string `json:"include"`
}

// Validate validates this backup restore request
func (m *BackupRestoreRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRestoreRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestoreRequest) UnmarshalBinary(b []byte) error {
	var res BackupRestoreRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
