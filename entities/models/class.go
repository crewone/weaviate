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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Class class
//
// swagger:model Class
type Class struct {

	// Name of the class as URI relative to the schema URL.
	Class string `json:"class,omitempty"`

	// Description of the class.
	Description string `json:"description,omitempty"`

	// keywords
	Keywords Keywords `json:"keywords,omitempty"`

	// The properties of the class.
	Properties []*Property `json:"properties"`

	// Set this to true if the object vector should include the class name in calculating the overall vector position
	VectorizeClassName *bool `json:"vectorizeClassName,omitempty"`
}

// Validate validates this class
func (m *Class) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Class) validateKeywords(formats strfmt.Registry) error {
	if swag.IsZero(m.Keywords) { // not required
		return nil
	}

	if err := m.Keywords.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("keywords")
		}
		return err
	}

	return nil
}

func (m *Class) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for i := 0; i < len(m.Properties); i++ {
		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {
			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Class) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Class) UnmarshalBinary(b []byte) error {
	var res Class
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
