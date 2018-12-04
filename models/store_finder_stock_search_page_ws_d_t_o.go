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

// StoreFinderStockSearchPageWsDTO store finder stock search page ws d t o
// swagger:model StoreFinderStockSearchPageWsDTO
type StoreFinderStockSearchPageWsDTO struct {

	// bound east longitude
	BoundEastLongitude float64 `json:"boundEastLongitude,omitempty"`

	// bound north latitude
	BoundNorthLatitude float64 `json:"boundNorthLatitude,omitempty"`

	// bound south latitude
	BoundSouthLatitude float64 `json:"boundSouthLatitude,omitempty"`

	// bound west longitude
	BoundWestLongitude float64 `json:"boundWestLongitude,omitempty"`

	// location text
	LocationText string `json:"locationText,omitempty"`

	// pagination
	Pagination *PaginationWsDTO `json:"pagination,omitempty"`

	// product
	Product *ProductWsDTO `json:"product,omitempty"`

	// sorts
	Sorts []*SortWsDTO `json:"sorts"`

	// source latitude
	SourceLatitude float64 `json:"sourceLatitude,omitempty"`

	// source longitude
	SourceLongitude float64 `json:"sourceLongitude,omitempty"`

	// stores
	Stores []*PointOfServiceStockWsDTO `json:"stores"`
}

// Validate validates this store finder stock search page ws d t o
func (m *StoreFinderStockSearchPageWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoreFinderStockSearchPageWsDTO) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *StoreFinderStockSearchPageWsDTO) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *StoreFinderStockSearchPageWsDTO) validateSorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Sorts) { // not required
		return nil
	}

	for i := 0; i < len(m.Sorts); i++ {
		if swag.IsZero(m.Sorts[i]) { // not required
			continue
		}

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StoreFinderStockSearchPageWsDTO) validateStores(formats strfmt.Registry) error {

	if swag.IsZero(m.Stores) { // not required
		return nil
	}

	for i := 0; i < len(m.Stores); i++ {
		if swag.IsZero(m.Stores[i]) { // not required
			continue
		}

		if m.Stores[i] != nil {
			if err := m.Stores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoreFinderStockSearchPageWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoreFinderStockSearchPageWsDTO) UnmarshalBinary(b []byte) error {
	var res StoreFinderStockSearchPageWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}