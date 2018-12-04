// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Pagination Pagination info
// swagger:model pagination
type Pagination struct {

	// Number of elements on this page
	Count int32 `json:"count,omitempty"`

	// Indicates if there is next page
	HasNext bool `json:"hasNext,omitempty"`

	// Indicates if there is previous page
	HasPrevious bool `json:"hasPrevious,omitempty"`

	// Current page number
	Page int32 `json:"page,omitempty"`

	// Total number of elements
	TotalCount int64 `json:"totalCount,omitempty"`

	// Total number of pages
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}