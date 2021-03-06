// Code generated by go-swagger; DO NOT EDIT.

package countries

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

// NewGetRegionsForCountryUsingGETParams creates a new GetRegionsForCountryUsingGETParams object
// with the default values initialized.
func NewGetRegionsForCountryUsingGETParams() *GetRegionsForCountryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetRegionsForCountryUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsForCountryUsingGETParamsWithTimeout creates a new GetRegionsForCountryUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionsForCountryUsingGETParamsWithTimeout(timeout time.Duration) *GetRegionsForCountryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetRegionsForCountryUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetRegionsForCountryUsingGETParamsWithContext creates a new GetRegionsForCountryUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionsForCountryUsingGETParamsWithContext(ctx context.Context) *GetRegionsForCountryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetRegionsForCountryUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetRegionsForCountryUsingGETParamsWithHTTPClient creates a new GetRegionsForCountryUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionsForCountryUsingGETParamsWithHTTPClient(client *http.Client) *GetRegionsForCountryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetRegionsForCountryUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetRegionsForCountryUsingGETParams contains all the parameters to send to the API endpoint
for the get regions for country using g e t operation typically these are written to a http.Request
*/
type GetRegionsForCountryUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CountyIsoCode
	  An ISO code for a country

	*/
	CountyIsoCode string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithTimeout(timeout time.Duration) *GetRegionsForCountryUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithContext(ctx context.Context) *GetRegionsForCountryUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithHTTPClient(client *http.Client) *GetRegionsForCountryUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithBaseSiteID(baseSiteID string) *GetRegionsForCountryUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCountyIsoCode adds the countyIsoCode to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithCountyIsoCode(countyIsoCode string) *GetRegionsForCountryUsingGETParams {
	o.SetCountyIsoCode(countyIsoCode)
	return o
}

// SetCountyIsoCode adds the countyIsoCode to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetCountyIsoCode(countyIsoCode string) {
	o.CountyIsoCode = countyIsoCode
}

// WithFields adds the fields to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) WithFields(fields *string) *GetRegionsForCountryUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get regions for country using g e t params
func (o *GetRegionsForCountryUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsForCountryUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param countyIsoCode
	if err := r.SetPathParam("countyIsoCode", o.CountyIsoCode); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
