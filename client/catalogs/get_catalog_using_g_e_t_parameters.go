// Code generated by go-swagger; DO NOT EDIT.

package catalogs

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

// NewGetCatalogUsingGETParams creates a new GetCatalogUsingGETParams object
// with the default values initialized.
func NewGetCatalogUsingGETParams() *GetCatalogUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCatalogUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogUsingGETParamsWithTimeout creates a new GetCatalogUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogUsingGETParamsWithTimeout(timeout time.Duration) *GetCatalogUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCatalogUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetCatalogUsingGETParamsWithContext creates a new GetCatalogUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogUsingGETParamsWithContext(ctx context.Context) *GetCatalogUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCatalogUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetCatalogUsingGETParamsWithHTTPClient creates a new GetCatalogUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogUsingGETParamsWithHTTPClient(client *http.Client) *GetCatalogUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCatalogUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetCatalogUsingGETParams contains all the parameters to send to the API endpoint
for the get catalog using g e t operation typically these are written to a http.Request
*/
type GetCatalogUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CatalogID
	  Catalog identifier

	*/
	CatalogID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithTimeout(timeout time.Duration) *GetCatalogUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithContext(ctx context.Context) *GetCatalogUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithHTTPClient(client *http.Client) *GetCatalogUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithBaseSiteID(baseSiteID string) *GetCatalogUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCatalogID adds the catalogID to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithCatalogID(catalogID string) *GetCatalogUsingGETParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithFields adds the fields to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) WithFields(fields *string) *GetCatalogUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get catalog using g e t params
func (o *GetCatalogUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param catalogId
	if err := r.SetPathParam("catalogId", o.CatalogID); err != nil {
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
