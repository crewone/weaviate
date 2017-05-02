/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package models




import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LocalAccessInfo local access info
// swagger:model LocalAccessInfo
type LocalAccessInfo struct {

	// local access entry
	LocalAccessEntry *LocalAccessEntry `json:"localAccessEntry,omitempty"`

	// Time in milliseconds since unix epoch of when the local auth token was minted.
	LocalAuthTokenMintTimeMs int64 `json:"localAuthTokenMintTimeMs,omitempty"`

	// Relative time left of token after API call.
	LocalAuthTokenTimeLeftMs int64 `json:"localAuthTokenTimeLeftMs,omitempty"`

	// Time in milliseconds of hold long the token is valid after minting.
	LocalAuthTokenTTLTimeMs int64 `json:"localAuthTokenTtlTimeMs,omitempty"`
}

// Validate validates this local access info
func (m *LocalAccessInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocalAccessEntry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocalAccessInfo) validateLocalAccessEntry(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalAccessEntry) { // not required
		return nil
	}

	if m.LocalAccessEntry != nil {

		if err := m.LocalAccessEntry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localAccessEntry")
			}
			return err
		}
	}

	return nil
}
