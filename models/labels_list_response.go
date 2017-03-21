package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// LabelsListResponse labels list response
// swagger:model LabelsListResponse
type LabelsListResponse struct {

	// Identifies what kind of resource this is. Value: the fixed string "weave#labelsListResponse".
	Kind *string `json:"kind,omitempty"`

	// The list of labels.
	Labels []*Label `json:"labels"`

	// page info
	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// token pagination
	TokenPagination *TokenPagination `json:"tokenPagination,omitempty"`
}

// Validate validates this labels list response
func (m *LabelsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePageInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTokenPagination(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelsListResponse) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {

		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {

			if err := m.Labels[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *LabelsListResponse) validatePageInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PageInfo) { // not required
		return nil
	}

	if m.PageInfo != nil {

		if err := m.PageInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *LabelsListResponse) validateTokenPagination(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenPagination) { // not required
		return nil
	}

	if m.TokenPagination != nil {

		if err := m.TokenPagination.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
