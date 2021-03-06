// Code generated by go-swagger; DO NOT EDIT.

package extended_carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// ConsolidatePickupLocationsUsingPOSTReader is a Reader for the ConsolidatePickupLocationsUsingPOST structure.
type ConsolidatePickupLocationsUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConsolidatePickupLocationsUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConsolidatePickupLocationsUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewConsolidatePickupLocationsUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewConsolidatePickupLocationsUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConsolidatePickupLocationsUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConsolidatePickupLocationsUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConsolidatePickupLocationsUsingPOSTOK creates a ConsolidatePickupLocationsUsingPOSTOK with default headers values
func NewConsolidatePickupLocationsUsingPOSTOK() *ConsolidatePickupLocationsUsingPOSTOK {
	return &ConsolidatePickupLocationsUsingPOSTOK{}
}

/*ConsolidatePickupLocationsUsingPOSTOK handles this case with default header values.

OK
*/
type ConsolidatePickupLocationsUsingPOSTOK struct {
	Payload *models.CartModificationListWsDTO
}

func (o *ConsolidatePickupLocationsUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/consolidate][%d] consolidatePickupLocationsUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ConsolidatePickupLocationsUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CartModificationListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConsolidatePickupLocationsUsingPOSTCreated creates a ConsolidatePickupLocationsUsingPOSTCreated with default headers values
func NewConsolidatePickupLocationsUsingPOSTCreated() *ConsolidatePickupLocationsUsingPOSTCreated {
	return &ConsolidatePickupLocationsUsingPOSTCreated{}
}

/*ConsolidatePickupLocationsUsingPOSTCreated handles this case with default header values.

Created
*/
type ConsolidatePickupLocationsUsingPOSTCreated struct {
}

func (o *ConsolidatePickupLocationsUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/consolidate][%d] consolidatePickupLocationsUsingPOSTCreated ", 201)
}

func (o *ConsolidatePickupLocationsUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsolidatePickupLocationsUsingPOSTUnauthorized creates a ConsolidatePickupLocationsUsingPOSTUnauthorized with default headers values
func NewConsolidatePickupLocationsUsingPOSTUnauthorized() *ConsolidatePickupLocationsUsingPOSTUnauthorized {
	return &ConsolidatePickupLocationsUsingPOSTUnauthorized{}
}

/*ConsolidatePickupLocationsUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ConsolidatePickupLocationsUsingPOSTUnauthorized struct {
}

func (o *ConsolidatePickupLocationsUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/consolidate][%d] consolidatePickupLocationsUsingPOSTUnauthorized ", 401)
}

func (o *ConsolidatePickupLocationsUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsolidatePickupLocationsUsingPOSTForbidden creates a ConsolidatePickupLocationsUsingPOSTForbidden with default headers values
func NewConsolidatePickupLocationsUsingPOSTForbidden() *ConsolidatePickupLocationsUsingPOSTForbidden {
	return &ConsolidatePickupLocationsUsingPOSTForbidden{}
}

/*ConsolidatePickupLocationsUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ConsolidatePickupLocationsUsingPOSTForbidden struct {
}

func (o *ConsolidatePickupLocationsUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/consolidate][%d] consolidatePickupLocationsUsingPOSTForbidden ", 403)
}

func (o *ConsolidatePickupLocationsUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConsolidatePickupLocationsUsingPOSTNotFound creates a ConsolidatePickupLocationsUsingPOSTNotFound with default headers values
func NewConsolidatePickupLocationsUsingPOSTNotFound() *ConsolidatePickupLocationsUsingPOSTNotFound {
	return &ConsolidatePickupLocationsUsingPOSTNotFound{}
}

/*ConsolidatePickupLocationsUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ConsolidatePickupLocationsUsingPOSTNotFound struct {
}

func (o *ConsolidatePickupLocationsUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/users/{userId}/carts/{cartId}/consolidate][%d] consolidatePickupLocationsUsingPOSTNotFound ", 404)
}

func (o *ConsolidatePickupLocationsUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
