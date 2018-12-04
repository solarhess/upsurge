// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Error Error message
// swagger:model error
type Error struct {

	// Error code
	ErrorCode string `json:"errorCode,omitempty"`

	// exception message
	ExceptionMessage string `json:"exceptionMessage,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// Descriptive, human readable error message.
	Message string `json:"message,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// Additional classification specific for each error type e.g. 'noStock'.
	Reason string `json:"reason,omitempty"`

	// Identifier of the related object e.g. '1'.
	Subject string `json:"subject,omitempty"`

	// Type of the object related to the error e.g. 'entry'.
	SubjectType string `json:"subjectType,omitempty"`

	// Type of the error e.g. 'LowStockError'.
	Type string `json:"type,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}