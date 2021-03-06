// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VariantMatrixElementWsDTO variant matrix element ws d t o
// swagger:model VariantMatrixElementWsDTO
type VariantMatrixElementWsDTO struct {

	// elements
	Elements []*VariantMatrixElementWsDTO `json:"elements"`

	// is leaf
	IsLeaf bool `json:"isLeaf,omitempty"`

	// parent variant category
	ParentVariantCategory *VariantCategoryWsDTO `json:"parentVariantCategory,omitempty"`

	// variant option
	VariantOption *VariantOptionWsDTO `json:"variantOption,omitempty"`

	// variant value category
	VariantValueCategory *VariantValueCategoryWsDTO `json:"variantValueCategory,omitempty"`
}

// Validate validates this variant matrix element ws d t o
func (m *VariantMatrixElementWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentVariantCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariantOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariantValueCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariantMatrixElementWsDTO) validateElements(formats strfmt.Registry) error {

	if swag.IsZero(m.Elements) { // not required
		return nil
	}

	for i := 0; i < len(m.Elements); i++ {
		if swag.IsZero(m.Elements[i]) { // not required
			continue
		}

		if m.Elements[i] != nil {
			if err := m.Elements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VariantMatrixElementWsDTO) validateParentVariantCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentVariantCategory) { // not required
		return nil
	}

	if m.ParentVariantCategory != nil {
		if err := m.ParentVariantCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentVariantCategory")
			}
			return err
		}
	}

	return nil
}

func (m *VariantMatrixElementWsDTO) validateVariantOption(formats strfmt.Registry) error {

	if swag.IsZero(m.VariantOption) { // not required
		return nil
	}

	if m.VariantOption != nil {
		if err := m.VariantOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variantOption")
			}
			return err
		}
	}

	return nil
}

func (m *VariantMatrixElementWsDTO) validateVariantValueCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.VariantValueCategory) { // not required
		return nil
	}

	if m.VariantValueCategory != nil {
		if err := m.VariantValueCategory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variantValueCategory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariantMatrixElementWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariantMatrixElementWsDTO) UnmarshalBinary(b []byte) error {
	var res VariantMatrixElementWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
