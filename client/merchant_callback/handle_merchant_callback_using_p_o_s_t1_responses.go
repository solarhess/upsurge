// Code generated by go-swagger; DO NOT EDIT.

package merchant_callback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// HandleMerchantCallbackUsingPOST1Reader is a Reader for the HandleMerchantCallbackUsingPOST1 structure.
type HandleMerchantCallbackUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleMerchantCallbackUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHandleMerchantCallbackUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewHandleMerchantCallbackUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewHandleMerchantCallbackUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewHandleMerchantCallbackUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewHandleMerchantCallbackUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHandleMerchantCallbackUsingPOST1OK creates a HandleMerchantCallbackUsingPOST1OK with default headers values
func NewHandleMerchantCallbackUsingPOST1OK() *HandleMerchantCallbackUsingPOST1OK {
	return &HandleMerchantCallbackUsingPOST1OK{}
}

/*HandleMerchantCallbackUsingPOST1OK handles this case with default header values.

OK
*/
type HandleMerchantCallbackUsingPOST1OK struct {
}

func (o *HandleMerchantCallbackUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response][%d] handleMerchantCallbackUsingPOST1OK ", 200)
}

func (o *HandleMerchantCallbackUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleMerchantCallbackUsingPOST1Created creates a HandleMerchantCallbackUsingPOST1Created with default headers values
func NewHandleMerchantCallbackUsingPOST1Created() *HandleMerchantCallbackUsingPOST1Created {
	return &HandleMerchantCallbackUsingPOST1Created{}
}

/*HandleMerchantCallbackUsingPOST1Created handles this case with default header values.

Created
*/
type HandleMerchantCallbackUsingPOST1Created struct {
}

func (o *HandleMerchantCallbackUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response][%d] handleMerchantCallbackUsingPOST1Created ", 201)
}

func (o *HandleMerchantCallbackUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleMerchantCallbackUsingPOST1Unauthorized creates a HandleMerchantCallbackUsingPOST1Unauthorized with default headers values
func NewHandleMerchantCallbackUsingPOST1Unauthorized() *HandleMerchantCallbackUsingPOST1Unauthorized {
	return &HandleMerchantCallbackUsingPOST1Unauthorized{}
}

/*HandleMerchantCallbackUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type HandleMerchantCallbackUsingPOST1Unauthorized struct {
}

func (o *HandleMerchantCallbackUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response][%d] handleMerchantCallbackUsingPOST1Unauthorized ", 401)
}

func (o *HandleMerchantCallbackUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleMerchantCallbackUsingPOST1Forbidden creates a HandleMerchantCallbackUsingPOST1Forbidden with default headers values
func NewHandleMerchantCallbackUsingPOST1Forbidden() *HandleMerchantCallbackUsingPOST1Forbidden {
	return &HandleMerchantCallbackUsingPOST1Forbidden{}
}

/*HandleMerchantCallbackUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type HandleMerchantCallbackUsingPOST1Forbidden struct {
}

func (o *HandleMerchantCallbackUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response][%d] handleMerchantCallbackUsingPOST1Forbidden ", 403)
}

func (o *HandleMerchantCallbackUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleMerchantCallbackUsingPOST1NotFound creates a HandleMerchantCallbackUsingPOST1NotFound with default headers values
func NewHandleMerchantCallbackUsingPOST1NotFound() *HandleMerchantCallbackUsingPOST1NotFound {
	return &HandleMerchantCallbackUsingPOST1NotFound{}
}

/*HandleMerchantCallbackUsingPOST1NotFound handles this case with default header values.

Not Found
*/
type HandleMerchantCallbackUsingPOST1NotFound struct {
}

func (o *HandleMerchantCallbackUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response][%d] handleMerchantCallbackUsingPOST1NotFound ", 404)
}

func (o *HandleMerchantCallbackUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
