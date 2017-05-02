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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
// swagger:model Event
type Event struct {

	// command patch
	CommandPatch *EventCommandPatch `json:"commandPatch,omitempty"`

	// New device connection state (if connectivity change event).
	ConnectionStatus string `json:"connectionStatus,omitempty"`

	// The device that was affected by this event.
	DeviceID string `json:"deviceId,omitempty"`

	// ID of the event.
	ID string `json:"id,omitempty"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#event".
	Kind *string `json:"kind,omitempty"`

	// state patch
	StatePatch JSONObject `json:"statePatch,omitempty"`

	// Time the event was generated in milliseconds since epoch UTC.
	TimeMs int64 `json:"timeMs,omitempty"`

	// Type of the event.
	Type string `json:"type,omitempty"`

	// User that caused the event (if applicable).
	UserEmail string `json:"userEmail,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandPatch(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateCommandPatch(formats strfmt.Registry) error {

	if swag.IsZero(m.CommandPatch) { // not required
		return nil
	}

	if m.CommandPatch != nil {

		if err := m.CommandPatch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commandPatch")
			}
			return err
		}
	}

	return nil
}

var eventTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["adapterDeactivated","commandCancelled","commandCreated","commandDeleted","commandExpired","commandUpdated","deviceAclUpdated","deviceConnectivityChange","deviceCreated","deviceDeleted","deviceLocationUpdated","deviceTransferred","deviceUpdated","deviceUseTimeUpdated","deviceUserAclCreated","deviceUserAclDeleted","deviceUserAclUpdated","eventsDeleted","eventsRecordingDisabled","eventsRecordingEnabled","locationCreated","locationDeleted","locationMemberAdded","locationMemberRemoved","locationUpdated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeTypePropEnum = append(eventTypeTypePropEnum, v)
	}
}

const (
	// EventTypeAdapterDeactivated captures enum value "adapterDeactivated"
	EventTypeAdapterDeactivated string = "adapterDeactivated"
	// EventTypeCommandCancelled captures enum value "commandCancelled"
	EventTypeCommandCancelled string = "commandCancelled"
	// EventTypeCommandCreated captures enum value "commandCreated"
	EventTypeCommandCreated string = "commandCreated"
	// EventTypeCommandDeleted captures enum value "commandDeleted"
	EventTypeCommandDeleted string = "commandDeleted"
	// EventTypeCommandExpired captures enum value "commandExpired"
	EventTypeCommandExpired string = "commandExpired"
	// EventTypeCommandUpdated captures enum value "commandUpdated"
	EventTypeCommandUpdated string = "commandUpdated"
	// EventTypeDeviceACLUpdated captures enum value "deviceAclUpdated"
	EventTypeDeviceACLUpdated string = "deviceAclUpdated"
	// EventTypeDeviceConnectivityChange captures enum value "deviceConnectivityChange"
	EventTypeDeviceConnectivityChange string = "deviceConnectivityChange"
	// EventTypeDeviceCreated captures enum value "deviceCreated"
	EventTypeDeviceCreated string = "deviceCreated"
	// EventTypeDeviceDeleted captures enum value "deviceDeleted"
	EventTypeDeviceDeleted string = "deviceDeleted"
	// EventTypeDeviceLocationUpdated captures enum value "deviceLocationUpdated"
	EventTypeDeviceLocationUpdated string = "deviceLocationUpdated"
	// EventTypeDeviceTransferred captures enum value "deviceTransferred"
	EventTypeDeviceTransferred string = "deviceTransferred"
	// EventTypeDeviceUpdated captures enum value "deviceUpdated"
	EventTypeDeviceUpdated string = "deviceUpdated"
	// EventTypeDeviceUseTimeUpdated captures enum value "deviceUseTimeUpdated"
	EventTypeDeviceUseTimeUpdated string = "deviceUseTimeUpdated"
	// EventTypeDeviceUserACLCreated captures enum value "deviceUserAclCreated"
	EventTypeDeviceUserACLCreated string = "deviceUserAclCreated"
	// EventTypeDeviceUserACLDeleted captures enum value "deviceUserAclDeleted"
	EventTypeDeviceUserACLDeleted string = "deviceUserAclDeleted"
	// EventTypeDeviceUserACLUpdated captures enum value "deviceUserAclUpdated"
	EventTypeDeviceUserACLUpdated string = "deviceUserAclUpdated"
	// EventTypeEventsDeleted captures enum value "eventsDeleted"
	EventTypeEventsDeleted string = "eventsDeleted"
	// EventTypeEventsRecordingDisabled captures enum value "eventsRecordingDisabled"
	EventTypeEventsRecordingDisabled string = "eventsRecordingDisabled"
	// EventTypeEventsRecordingEnabled captures enum value "eventsRecordingEnabled"
	EventTypeEventsRecordingEnabled string = "eventsRecordingEnabled"
	// EventTypeLocationCreated captures enum value "locationCreated"
	EventTypeLocationCreated string = "locationCreated"
	// EventTypeLocationDeleted captures enum value "locationDeleted"
	EventTypeLocationDeleted string = "locationDeleted"
	// EventTypeLocationMemberAdded captures enum value "locationMemberAdded"
	EventTypeLocationMemberAdded string = "locationMemberAdded"
	// EventTypeLocationMemberRemoved captures enum value "locationMemberRemoved"
	EventTypeLocationMemberRemoved string = "locationMemberRemoved"
	// EventTypeLocationUpdated captures enum value "locationUpdated"
	EventTypeLocationUpdated string = "locationUpdated"
)

// prop value enum
func (m *Event) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, eventTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Event) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// EventCommandPatch Command-related changes (if applicable).
// swagger:model EventCommandPatch
type EventCommandPatch struct {

	// ID of the affected command.
	CommandID string `json:"commandId,omitempty"`

	// New command state.
	State string `json:"state,omitempty"`
}

// Validate validates this event command patch
func (m *EventCommandPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
