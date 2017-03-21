package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// DevicesRemoveNicknameResponse devices remove nickname response
// swagger:model DevicesRemoveNicknameResponse
type DevicesRemoveNicknameResponse struct {

	// Nicknames of the device.
	Nicknames []string `json:"nicknames"`
}

// Validate validates this devices remove nickname response
func (m *DevicesRemoveNicknameResponse) Validate(formats strfmt.Registry) error {
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

func (m *DevicesRemoveNicknameResponse) validateNicknames(formats strfmt.Registry) error {

	if swag.IsZero(m.Nicknames) { // not required
		return nil
	}

	return nil
}
