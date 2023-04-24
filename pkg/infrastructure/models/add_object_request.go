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

// AddObjectRequest add object request
//
// swagger:model AddObjectRequest
type AddObjectRequest struct {

	// object
	// Required: true
	Object *string `json:"object"`
}

// Validate validates this add object request
func (m *AddObjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddObjectRequest) validateObject(formats strfmt.Registry) error {

	if err := validate.Required("object", "body", m.Object); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add object request based on context it is used
func (m *AddObjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddObjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddObjectRequest) UnmarshalBinary(b []byte) error {
	var res AddObjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}