// Code generated by go-swagger; DO NOT EDIT.

package extended_carts

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

// NewConsolidatePickupLocationsUsingPOSTParams creates a new ConsolidatePickupLocationsUsingPOSTParams object
// with the default values initialized.
func NewConsolidatePickupLocationsUsingPOSTParams() *ConsolidatePickupLocationsUsingPOSTParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ConsolidatePickupLocationsUsingPOSTParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewConsolidatePickupLocationsUsingPOSTParamsWithTimeout creates a new ConsolidatePickupLocationsUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsolidatePickupLocationsUsingPOSTParamsWithTimeout(timeout time.Duration) *ConsolidatePickupLocationsUsingPOSTParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ConsolidatePickupLocationsUsingPOSTParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewConsolidatePickupLocationsUsingPOSTParamsWithContext creates a new ConsolidatePickupLocationsUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsolidatePickupLocationsUsingPOSTParamsWithContext(ctx context.Context) *ConsolidatePickupLocationsUsingPOSTParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ConsolidatePickupLocationsUsingPOSTParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewConsolidatePickupLocationsUsingPOSTParamsWithHTTPClient creates a new ConsolidatePickupLocationsUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsolidatePickupLocationsUsingPOSTParamsWithHTTPClient(client *http.Client) *ConsolidatePickupLocationsUsingPOSTParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ConsolidatePickupLocationsUsingPOSTParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*ConsolidatePickupLocationsUsingPOSTParams contains all the parameters to send to the API endpoint
for the consolidate pickup locations using p o s t operation typically these are written to a http.Request
*/
type ConsolidatePickupLocationsUsingPOSTParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*StoreName
	  The name of the store where items will be picked up

	*/
	StoreName string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithTimeout(timeout time.Duration) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithContext(ctx context.Context) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithHTTPClient(client *http.Client) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithBaseSiteID(baseSiteID string) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithCartID(cartID string) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithFields adds the fields to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithFields(fields *string) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStoreName adds the storeName to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithStoreName(storeName string) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetStoreName(storeName)
	return o
}

// SetStoreName adds the storeName to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetStoreName(storeName string) {
	o.StoreName = storeName
}

// WithUserID adds the userID to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) WithUserID(userID string) *ConsolidatePickupLocationsUsingPOSTParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the consolidate pickup locations using p o s t params
func (o *ConsolidatePickupLocationsUsingPOSTParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsolidatePickupLocationsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
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

	// query param storeName
	qrStoreName := o.StoreName
	qStoreName := qrStoreName
	if qStoreName != "" {
		if err := r.SetQueryParam("storeName", qStoreName); err != nil {
			return err
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