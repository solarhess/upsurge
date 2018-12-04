// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// GetStockDataUsingGETReader is a Reader for the GetStockDataUsingGET structure.
type GetStockDataUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStockDataUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStockDataUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetStockDataUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetStockDataUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetStockDataUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStockDataUsingGETOK creates a GetStockDataUsingGETOK with default headers values
func NewGetStockDataUsingGETOK() *GetStockDataUsingGETOK {
	return &GetStockDataUsingGETOK{}
}

/*GetStockDataUsingGETOK handles this case with default header values.

OK
*/
type GetStockDataUsingGETOK struct {
	Payload *models.StockWsDTO
}

func (o *GetStockDataUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/{productCode}/stock/{storeName}][%d] getStockDataUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetStockDataUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StockWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStockDataUsingGETUnauthorized creates a GetStockDataUsingGETUnauthorized with default headers values
func NewGetStockDataUsingGETUnauthorized() *GetStockDataUsingGETUnauthorized {
	return &GetStockDataUsingGETUnauthorized{}
}

/*GetStockDataUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetStockDataUsingGETUnauthorized struct {
}

func (o *GetStockDataUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/{productCode}/stock/{storeName}][%d] getStockDataUsingGETUnauthorized ", 401)
}

func (o *GetStockDataUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStockDataUsingGETForbidden creates a GetStockDataUsingGETForbidden with default headers values
func NewGetStockDataUsingGETForbidden() *GetStockDataUsingGETForbidden {
	return &GetStockDataUsingGETForbidden{}
}

/*GetStockDataUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetStockDataUsingGETForbidden struct {
}

func (o *GetStockDataUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/{productCode}/stock/{storeName}][%d] getStockDataUsingGETForbidden ", 403)
}

func (o *GetStockDataUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStockDataUsingGETNotFound creates a GetStockDataUsingGETNotFound with default headers values
func NewGetStockDataUsingGETNotFound() *GetStockDataUsingGETNotFound {
	return &GetStockDataUsingGETNotFound{}
}

/*GetStockDataUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetStockDataUsingGETNotFound struct {
}

func (o *GetStockDataUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/{productCode}/stock/{storeName}][%d] getStockDataUsingGETNotFound ", 404)
}

func (o *GetStockDataUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}