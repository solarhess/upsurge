// Code generated by go-swagger; DO NOT EDIT.

package promotions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new promotions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for promotions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPromotionByCodeUsingGET gets a promotion based on code

Returns details of a single promotion specified by a promotion code. Requests pertaining to promotions have been developed for the previous version of promotions and vouchers and therefore some of them are currently not compatible with the new promotion engine.
*/
func (a *Client) GetPromotionByCodeUsingGET(params *GetPromotionByCodeUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPromotionByCodeUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPromotionByCodeUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPromotionByCodeUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/promotions/{code}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPromotionByCodeUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPromotionByCodeUsingGETOK), nil

}

/*
GetPromotionsUsingGET1 gets a list of promotions

Returns promotions defined for a current base site. Requests pertaining to promotions have been developed for the previous version of promotions and vouchers and therefore some of them are currently not compatible with the new promotion engine.
*/
func (a *Client) GetPromotionsUsingGET1(params *GetPromotionsUsingGET1Params, authInfo runtime.ClientAuthInfoWriter) (*GetPromotionsUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPromotionsUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPromotionsUsingGET_1",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/promotions",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPromotionsUsingGET1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPromotionsUsingGET1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
