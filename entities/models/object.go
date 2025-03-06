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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Object object
//
// swagger:model Object
type Object struct {

	// additional
	Additional AdditionalProperties `json:"additional,omitempty"`

	// Class of the Object, defined in the schema.
	Class string `json:"class,omitempty"`

	// (Response only) Timestamp of creation of this object in milliseconds since epoch UTC.
	CreationTimeUnix int64 `json:"creationTimeUnix,omitempty"`

	// ID of the Object.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// (Response only) Timestamp of the last object update in milliseconds since epoch UTC.
	LastUpdateTimeUnix int64 `json:"lastUpdateTimeUnix,omitempty"`

	// properties
	Properties PropertySchema `json:"properties,omitempty"`

	// Name of the Objects tenant.
	Tenant string `json:"tenant,omitempty"`

	// Tokens extracted from the text properties of the object.
	Tokens []string `json:"tokens"`

	// This field returns vectors associated with the Object. C11yVector, Vector or Vectors values are possible.
	Vector C11yVector `json:"vector,omitempty"`

	// vector weights
	VectorWeights VectorWeights `json:"vectorWeights,omitempty"`

	// This field returns vectors associated with the Object.
	Vectors Vectors `json:"vectors,omitempty"`
}

// Validate validates this object
func (m *Object) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditional(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Object) validateAdditional(formats strfmt.Registry) error {
	if swag.IsZero(m.Additional) { // not required
		return nil
	}

	if m.Additional != nil {
		if err := m.Additional.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additional")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("additional")
			}
			return err
		}
	}

	return nil
}

func (m *Object) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Object) validateVector(formats strfmt.Registry) error {
	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vector")
		}
		return err
	}

	return nil
}

func (m *Object) validateVectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Vectors) { // not required
		return nil
	}

	if m.Vectors != nil {
		if err := m.Vectors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vectors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vectors")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object based on the context it is used
func (m *Object) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditional(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Object) contextValidateAdditional(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Additional.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("additional")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("additional")
		}
		return err
	}

	return nil
}

func (m *Object) contextValidateVector(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Vector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vector")
		}
		return err
	}

	return nil
}

func (m *Object) contextValidateVectors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Vectors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vectors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vectors")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Object) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Object) UnmarshalBinary(b []byte) error {
	var res Object
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
