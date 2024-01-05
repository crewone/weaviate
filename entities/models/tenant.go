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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tenant attributes representing a single tenant within weaviate
//
// swagger:model Tenant
type Tenant struct {

	// activity status of the tenant's shard. Optional for creating tenant (implicit `HOT`) and required for updating tenant. Allowed values are `HOT` - tenant is fully active, `WARM` - tenant is active, some restrictions are imposed (TBD; not supported yet), `COLD` - tenant is inactive; no actions can be performed on tenant, tenant's files are stored locally, `FROZEN` - as COLD, but files are stored on cloud storage (not supported yet)
	// Enum: [HOT WARM COLD FROZEN]
	ActivityStatus string `json:"activityStatus,omitempty"`

	// name of the tenant
	Name string `json:"name,omitempty"`
}

// Validate validates this tenant
func (m *Tenant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tenantTypeActivityStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HOT","WARM","COLD","FROZEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tenantTypeActivityStatusPropEnum = append(tenantTypeActivityStatusPropEnum, v)
	}
}

const (

	// TenantActivityStatusHOT captures enum value "HOT"
	TenantActivityStatusHOT string = "HOT"

	// TenantActivityStatusWARM captures enum value "WARM"
	TenantActivityStatusWARM string = "WARM"

	// TenantActivityStatusCOLD captures enum value "COLD"
	TenantActivityStatusCOLD string = "COLD"

	// TenantActivityStatusFROZEN captures enum value "FROZEN"
	TenantActivityStatusFROZEN string = "FROZEN"
)

// prop value enum
func (m *Tenant) validateActivityStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tenantTypeActivityStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Tenant) validateActivityStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ActivityStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateActivityStatusEnum("activityStatus", "body", m.ActivityStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tenant based on context it is used
func (m *Tenant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tenant) UnmarshalBinary(b []byte) error {
	var res Tenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
