// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceCommand device command
//
// swagger:model DeviceCommand
type DeviceCommand struct {

	// The request ID of the device command execution.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The payload for the device command.
	Request interface{} `json:"request,omitempty"`

	// The payload of the device command response.
	// Read Only: true
	Response interface{} `json:"response,omitempty"`

	// The status code of the device command response.
	// Read Only: true
	ResponseCode int64 `json:"responseCode,omitempty"`
}

// Validate validates this device command
func (m *DeviceCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceCommand) UnmarshalBinary(b []byte) error {
	var res DeviceCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}