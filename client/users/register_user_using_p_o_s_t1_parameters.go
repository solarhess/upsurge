// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
	"golang.org/x/net/context"
)

// NewRegisterUserUsingPOST1Params creates a new RegisterUserUsingPOST1Params object
// with the default values initialized.
func NewRegisterUserUsingPOST1Params() *RegisterUserUsingPOST1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RegisterUserUsingPOST1Params{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterUserUsingPOST1ParamsWithTimeout creates a new RegisterUserUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterUserUsingPOST1ParamsWithTimeout(timeout time.Duration) *RegisterUserUsingPOST1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RegisterUserUsingPOST1Params{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewRegisterUserUsingPOST1ParamsWithContext creates a new RegisterUserUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterUserUsingPOST1ParamsWithContext(ctx context.Context) *RegisterUserUsingPOST1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RegisterUserUsingPOST1Params{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewRegisterUserUsingPOST1ParamsWithHTTPClient creates a new RegisterUserUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterUserUsingPOST1ParamsWithHTTPClient(client *http.Client) *RegisterUserUsingPOST1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RegisterUserUsingPOST1Params{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*RegisterUserUsingPOST1Params contains all the parameters to send to the API endpoint
for the register user using p o s t 1 operation typically these are written to a http.Request
*/
type RegisterUserUsingPOST1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*User
	  User's object.

	*/
	User *models.UserSignUpWsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithTimeout(timeout time.Duration) *RegisterUserUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithContext(ctx context.Context) *RegisterUserUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithHTTPClient(client *http.Client) *RegisterUserUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithBaseSiteID(baseSiteID string) *RegisterUserUsingPOST1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithFields adds the fields to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithFields(fields *string) *RegisterUserUsingPOST1Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithUser adds the user to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) WithUser(user *models.UserSignUpWsDTO) *RegisterUserUsingPOST1Params {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the register user using p o s t 1 params
func (o *RegisterUserUsingPOST1Params) SetUser(user *models.UserSignUpWsDTO) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterUserUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.User != nil {
		if err := r.SetBodyParam(o.User); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
