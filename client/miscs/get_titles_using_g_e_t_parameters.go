// Code generated by go-swagger; DO NOT EDIT.

package miscs

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

// NewGetTitlesUsingGETParams creates a new GetTitlesUsingGETParams object
// with the default values initialized.
func NewGetTitlesUsingGETParams() *GetTitlesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetTitlesUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTitlesUsingGETParamsWithTimeout creates a new GetTitlesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTitlesUsingGETParamsWithTimeout(timeout time.Duration) *GetTitlesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetTitlesUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetTitlesUsingGETParamsWithContext creates a new GetTitlesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTitlesUsingGETParamsWithContext(ctx context.Context) *GetTitlesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetTitlesUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetTitlesUsingGETParamsWithHTTPClient creates a new GetTitlesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTitlesUsingGETParamsWithHTTPClient(client *http.Client) *GetTitlesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetTitlesUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetTitlesUsingGETParams contains all the parameters to send to the API endpoint
for the get titles using g e t operation typically these are written to a http.Request
*/
type GetTitlesUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get titles using g e t params
func (o *GetTitlesUsingGETParams) WithTimeout(timeout time.Duration) *GetTitlesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get titles using g e t params
func (o *GetTitlesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get titles using g e t params
func (o *GetTitlesUsingGETParams) WithContext(ctx context.Context) *GetTitlesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get titles using g e t params
func (o *GetTitlesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get titles using g e t params
func (o *GetTitlesUsingGETParams) WithHTTPClient(client *http.Client) *GetTitlesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get titles using g e t params
func (o *GetTitlesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get titles using g e t params
func (o *GetTitlesUsingGETParams) WithBaseSiteID(baseSiteID string) *GetTitlesUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get titles using g e t params
func (o *GetTitlesUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithFields adds the fields to the get titles using g e t params
func (o *GetTitlesUsingGETParams) WithFields(fields *string) *GetTitlesUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get titles using g e t params
func (o *GetTitlesUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetTitlesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}