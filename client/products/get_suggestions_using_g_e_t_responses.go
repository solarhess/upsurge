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

// GetSuggestionsUsingGETReader is a Reader for the GetSuggestionsUsingGET structure.
type GetSuggestionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSuggestionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSuggestionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSuggestionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSuggestionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSuggestionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSuggestionsUsingGETOK creates a GetSuggestionsUsingGETOK with default headers values
func NewGetSuggestionsUsingGETOK() *GetSuggestionsUsingGETOK {
	return &GetSuggestionsUsingGETOK{}
}

/*GetSuggestionsUsingGETOK handles this case with default header values.

OK
*/
type GetSuggestionsUsingGETOK struct {
	Payload *models.SuggestionListWsDTO
}

func (o *GetSuggestionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/suggestions][%d] getSuggestionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetSuggestionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuggestionListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSuggestionsUsingGETUnauthorized creates a GetSuggestionsUsingGETUnauthorized with default headers values
func NewGetSuggestionsUsingGETUnauthorized() *GetSuggestionsUsingGETUnauthorized {
	return &GetSuggestionsUsingGETUnauthorized{}
}

/*GetSuggestionsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSuggestionsUsingGETUnauthorized struct {
}

func (o *GetSuggestionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/suggestions][%d] getSuggestionsUsingGETUnauthorized ", 401)
}

func (o *GetSuggestionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuggestionsUsingGETForbidden creates a GetSuggestionsUsingGETForbidden with default headers values
func NewGetSuggestionsUsingGETForbidden() *GetSuggestionsUsingGETForbidden {
	return &GetSuggestionsUsingGETForbidden{}
}

/*GetSuggestionsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetSuggestionsUsingGETForbidden struct {
}

func (o *GetSuggestionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/suggestions][%d] getSuggestionsUsingGETForbidden ", 403)
}

func (o *GetSuggestionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSuggestionsUsingGETNotFound creates a GetSuggestionsUsingGETNotFound with default headers values
func NewGetSuggestionsUsingGETNotFound() *GetSuggestionsUsingGETNotFound {
	return &GetSuggestionsUsingGETNotFound{}
}

/*GetSuggestionsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetSuggestionsUsingGETNotFound struct {
}

func (o *GetSuggestionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/products/suggestions][%d] getSuggestionsUsingGETNotFound ", 404)
}

func (o *GetSuggestionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}