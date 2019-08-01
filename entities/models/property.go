//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Property property
// swagger:model Property
type Property struct {

	// The cardinality of this property. If you want to store more than one value in a property, set this to 'many'. Defaults to 'atMostOne'. Note that by default properties can be empty in Weaviate.
	// Enum: [atMostOne many]
	Cardinality *string `json:"cardinality,omitempty"`

	// Can be a reference to another type when it starts with a capital (for example Person), otherwise "string" or "int".
	DataType []string `json:"dataType"`

	// Description of the property.
	Description string `json:"description,omitempty"`

	// keywords
	Keywords Keywords `json:"keywords,omitempty"`

	// Name of the property as URI relative to the schema URL.
	Name string `json:"name,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardinality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeywords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var propertyTypeCardinalityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["atMostOne","many"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyTypeCardinalityPropEnum = append(propertyTypeCardinalityPropEnum, v)
	}
}

const (

	// PropertyCardinalityAtMostOne captures enum value "atMostOne"
	PropertyCardinalityAtMostOne string = "atMostOne"

	// PropertyCardinalityMany captures enum value "many"
	PropertyCardinalityMany string = "many"
)

// prop value enum
func (m *Property) validateCardinalityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, propertyTypeCardinalityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Property) validateCardinality(formats strfmt.Registry) error {

	if swag.IsZero(m.Cardinality) { // not required
		return nil
	}

	// value enum
	if err := m.validateCardinalityEnum("cardinality", "body", *m.Cardinality); err != nil {
		return err
	}

	return nil
}

func (m *Property) validateKeywords(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
