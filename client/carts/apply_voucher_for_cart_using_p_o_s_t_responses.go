// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// ApplyVoucherForCartUsingPOSTReader is a Reader for the ApplyVoucherForCartUsingPOST structure.
type ApplyVoucherForCartUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplyVoucherForCartUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewApplyVoucherForCartUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewApplyVoucherForCartUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewApplyVoucherForCartUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewApplyVoucherForCartUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewApplyVoucherForCartUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplyVoucherForCartUsingPOSTOK creates a ApplyVoucherForCartUsingPOSTOK with default headers values
func NewApplyVoucherForCartUsingPOSTOK() *ApplyVoucherForCartUsingPOSTOK {
	return &ApplyVoucherForCartUsingPOSTOK{}
}

/*ApplyVoucherForCartUsingPOSTOK handles this case with default header values.

OK
*/
type ApplyVoucherForCartUsingPOSTOK struct {
}

func (o *ApplyVoucherForCartUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/vouchers][%d] applyVoucherForCartUsingPOSTOK ", 200)
}

func (o *ApplyVoucherForCartUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyVoucherForCartUsingPOSTCreated creates a ApplyVoucherForCartUsingPOSTCreated with default headers values
func NewApplyVoucherForCartUsingPOSTCreated() *ApplyVoucherForCartUsingPOSTCreated {
	return &ApplyVoucherForCartUsingPOSTCreated{}
}

/*ApplyVoucherForCartUsingPOSTCreated handles this case with default header values.

Created
*/
type ApplyVoucherForCartUsingPOSTCreated struct {
}

func (o *ApplyVoucherForCartUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/vouchers][%d] applyVoucherForCartUsingPOSTCreated ", 201)
}

func (o *ApplyVoucherForCartUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyVoucherForCartUsingPOSTUnauthorized creates a ApplyVoucherForCartUsingPOSTUnauthorized with default headers values
func NewApplyVoucherForCartUsingPOSTUnauthorized() *ApplyVoucherForCartUsingPOSTUnauthorized {
	return &ApplyVoucherForCartUsingPOSTUnauthorized{}
}

/*ApplyVoucherForCartUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ApplyVoucherForCartUsingPOSTUnauthorized struct {
}

func (o *ApplyVoucherForCartUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/vouchers][%d] applyVoucherForCartUsingPOSTUnauthorized ", 401)
}

func (o *ApplyVoucherForCartUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyVoucherForCartUsingPOSTForbidden creates a ApplyVoucherForCartUsingPOSTForbidden with default headers values
func NewApplyVoucherForCartUsingPOSTForbidden() *ApplyVoucherForCartUsingPOSTForbidden {
	return &ApplyVoucherForCartUsingPOSTForbidden{}
}

/*ApplyVoucherForCartUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ApplyVoucherForCartUsingPOSTForbidden struct {
}

func (o *ApplyVoucherForCartUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/vouchers][%d] applyVoucherForCartUsingPOSTForbidden ", 403)
}

func (o *ApplyVoucherForCartUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApplyVoucherForCartUsingPOSTNotFound creates a ApplyVoucherForCartUsingPOSTNotFound with default headers values
func NewApplyVoucherForCartUsingPOSTNotFound() *ApplyVoucherForCartUsingPOSTNotFound {
	return &ApplyVoucherForCartUsingPOSTNotFound{}
}

/*ApplyVoucherForCartUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ApplyVoucherForCartUsingPOSTNotFound struct {
}

func (o *ApplyVoucherForCartUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/vouchers][%d] applyVoucherForCartUsingPOSTNotFound ", 404)
}

func (o *ApplyVoucherForCartUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
