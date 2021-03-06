// Code generated by go-swagger; DO NOT EDIT.

package address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewGetAddressUsingGETParams creates a new GetAddressUsingGETParams object
// with the default values initialized.
func NewGetAddressUsingGETParams() *GetAddressUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetAddressUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddressUsingGETParamsWithTimeout creates a new GetAddressUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddressUsingGETParamsWithTimeout(timeout time.Duration) *GetAddressUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetAddressUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetAddressUsingGETParamsWithContext creates a new GetAddressUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddressUsingGETParamsWithContext(ctx context.Context) *GetAddressUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetAddressUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetAddressUsingGETParamsWithHTTPClient creates a new GetAddressUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddressUsingGETParamsWithHTTPClient(client *http.Client) *GetAddressUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetAddressUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetAddressUsingGETParams contains all the parameters to send to the API endpoint
for the get address using g e t operation typically these are written to a http.Request
*/
type GetAddressUsingGETParams struct {

	/*AddressID
	  Address identifier.

	*/
	AddressID string
	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get address using g e t params
func (o *GetAddressUsingGETParams) WithTimeout(timeout time.Duration) *GetAddressUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get address using g e t params
func (o *GetAddressUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get address using g e t params
func (o *GetAddressUsingGETParams) WithContext(ctx context.Context) *GetAddressUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get address using g e t params
func (o *GetAddressUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get address using g e t params
func (o *GetAddressUsingGETParams) WithHTTPClient(client *http.Client) *GetAddressUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get address using g e t params
func (o *GetAddressUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the get address using g e t params
func (o *GetAddressUsingGETParams) WithAddressID(addressID string) *GetAddressUsingGETParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the get address using g e t params
func (o *GetAddressUsingGETParams) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WithBaseSiteID adds the baseSiteID to the get address using g e t params
func (o *GetAddressUsingGETParams) WithBaseSiteID(baseSiteID string) *GetAddressUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get address using g e t params
func (o *GetAddressUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithFields adds the fields to the get address using g e t params
func (o *GetAddressUsingGETParams) WithFields(fields *string) *GetAddressUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get address using g e t params
func (o *GetAddressUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get address using g e t params
func (o *GetAddressUsingGETParams) WithUserID(userID string) *GetAddressUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get address using g e t params
func (o *GetAddressUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddressUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addressId
	if err := r.SetPathParam("addressId", o.AddressID); err != nil {
		return err
	}

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
