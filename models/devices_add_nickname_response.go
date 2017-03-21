package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// DevicesAddNicknameResponse devices add nickname response
// swagger:model DevicesAddNicknameResponse
type DevicesAddNicknameResponse struct {

	// Nicknames of the device.
	Nicknames []string `json:"nicknames"`
}

// Validate validates this devices add nickname response
func (m *DevicesAddNicknameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNicknames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicesAddNicknameResponse) validateNicknames(formats strfmt.Registry) error {

	if swag.IsZero(m.Nicknames) { // not required
		return nil
	}

	return nil
}
