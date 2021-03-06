// Code generated by go-swagger; DO NOT EDIT.

package merchant_callback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new merchant callback API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for merchant callback API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
HandleMerchantCallbackUsingPOST verifies the decision of the merchant

Verifies the decision of the merchant.

Note, the “Try it out” button is not enabled for this method (always returns an error) because the Merchant Callback Controller handles parameters differently, depending on which payment provider is used. For more information about this controller, please refer to the “acceleratorwebservicesaddon AddOn” documentation on help.hybris.com.
*/
func (a *Client) HandleMerchantCallbackUsingPOST(params *HandleMerchantCallbackUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*HandleMerchantCallbackUsingPOSTOK, *HandleMerchantCallbackUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHandleMerchantCallbackUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "handleMerchantCallbackUsingPOST",
		Method:             "POST",
		PathPattern:        "/{baseSiteId}/integration/merchant_callback",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HandleMerchantCallbackUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *HandleMerchantCallbackUsingPOSTOK:
		return value, nil, nil
	case *HandleMerchantCallbackUsingPOSTCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
HandleMerchantCallbackUsingPOST1 verifies the decision of the merchant for a cart

Verifies the decision of the merchant for a specified cart, and stores information of the PaymentSubscriptionResult for the cart.

Note, the “Try it out” button is not enabled for this method (always returns an error) because the Merchant Callback Controller handles parameters differently, depending on which payment provider is used. For more information about this controller, please refer to the “acceleratorwebservicesaddon AddOn” documentation on help.hybris.com.
*/
func (a *Client) HandleMerchantCallbackUsingPOST1(params *HandleMerchantCallbackUsingPOST1Params, authInfo runtime.ClientAuthInfoWriter) (*HandleMerchantCallbackUsingPOST1OK, *HandleMerchantCallbackUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHandleMerchantCallbackUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "handleMerchantCallbackUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/{baseSiteId}/integration/users/{userId}/carts/{cartId}/payment/sop/response",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HandleMerchantCallbackUsingPOST1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *HandleMerchantCallbackUsingPOST1OK:
		return value, nil, nil
	case *HandleMerchantCallbackUsingPOST1Created:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
