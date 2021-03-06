// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CMSPageWsDTO c m s page ws d t o
// swagger:model CMSPageWsDTO
type CMSPageWsDTO struct {

	// content slots
	ContentSlots *ContentSlotListWsDTO `json:"contentSlots,omitempty"`

	// default page
	DefaultPage bool `json:"defaultPage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// template
	Template string `json:"template,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type code
	TypeCode string `json:"typeCode,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this c m s page ws d t o
func (m *CMSPageWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentSlots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CMSPageWsDTO) validateContentSlots(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentSlots) { // not required
		return nil
	}

	if m.ContentSlots != nil {
		if err := m.ContentSlots.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentSlots")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CMSPageWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CMSPageWsDTO) UnmarshalBinary(b []byte) error {
	var res CMSPageWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
