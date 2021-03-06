// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

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

// NewCreateNewCustomerGroupUsingPOST1Params creates a new CreateNewCustomerGroupUsingPOST1Params object
// with the default values initialized.
func NewCreateNewCustomerGroupUsingPOST1Params() *CreateNewCustomerGroupUsingPOST1Params {
	var ()
	return &CreateNewCustomerGroupUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNewCustomerGroupUsingPOST1ParamsWithTimeout creates a new CreateNewCustomerGroupUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNewCustomerGroupUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateNewCustomerGroupUsingPOST1Params {
	var ()
	return &CreateNewCustomerGroupUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateNewCustomerGroupUsingPOST1ParamsWithContext creates a new CreateNewCustomerGroupUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNewCustomerGroupUsingPOST1ParamsWithContext(ctx context.Context) *CreateNewCustomerGroupUsingPOST1Params {
	var ()
	return &CreateNewCustomerGroupUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateNewCustomerGroupUsingPOST1ParamsWithHTTPClient creates a new CreateNewCustomerGroupUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNewCustomerGroupUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateNewCustomerGroupUsingPOST1Params {
	var ()
	return &CreateNewCustomerGroupUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateNewCustomerGroupUsingPOST1Params contains all the parameters to send to the API endpoint
for the create new customer group using p o s t 1 operation typically these are written to a http.Request
*/
type CreateNewCustomerGroupUsingPOST1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*UserGroup
	  User group object with id and name.

	*/
	UserGroup *models.UserGroupWsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateNewCustomerGroupUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) WithContext(ctx context.Context) *CreateNewCustomerGroupUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateNewCustomerGroupUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) WithBaseSiteID(baseSiteID string) *CreateNewCustomerGroupUsingPOST1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithUserGroup adds the userGroup to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) WithUserGroup(userGroup *models.UserGroupWsDTO) *CreateNewCustomerGroupUsingPOST1Params {
	o.SetUserGroup(userGroup)
	return o
}

// SetUserGroup adds the userGroup to the create new customer group using p o s t 1 params
func (o *CreateNewCustomerGroupUsingPOST1Params) SetUserGroup(userGroup *models.UserGroupWsDTO) {
	o.UserGroup = userGroup
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNewCustomerGroupUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.UserGroup != nil {
		if err := r.SetBodyParam(o.UserGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
