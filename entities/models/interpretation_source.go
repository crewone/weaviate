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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterpretationSource This additional property contains additional info about the how the class was vectorized
//
// swagger:model InterpretationSource
type InterpretationSource struct {

	// concept
	Concept string `json:"concept,omitempty"`

	// occurrence
	Occurrence uint64 `json:"occurrence,omitempty"`

	// weight
	Weight float64 `json:"weight,omitempty"`
}

// Validate validates this interpretation source
func (m *InterpretationSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterpretationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterpretationSource) UnmarshalBinary(b []byte) error {
	var res InterpretationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
