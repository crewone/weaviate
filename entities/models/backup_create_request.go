//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupCreateRequest Request body for creating a backup of a set of classes
//
// swagger:model BackupCreateRequest
type BackupCreateRequest struct {

	// Custom configuration for the backup creation process
	Config interface{} `json:"config,omitempty"`

	// List of classes to exclude from the backup creation process
	Exclude []string `json:"exclude"`

	// The ID of the backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.
	ID string `json:"id,omitempty"`

	// List of classes to include in the backup creation process
	Include []string `json:"include"`
}

// Validate validates this backup create request
func (m *BackupCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup create request based on context it is used
func (m *BackupCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupCreateRequest) UnmarshalBinary(b []byte) error {
	var res BackupCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
