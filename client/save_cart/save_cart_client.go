// Code generated by go-swagger; DO NOT EDIT.

package save_cart

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new save cart API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for save cart API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CloneSaveCartUsingPOST explicitlies clones a cart

Explicitly clones a cart.
*/
func (a *Client) CloneSaveCartUsingPOST(params *CloneSaveCartUsingPOSTParams, authInfo runtime.ClientAuthInfoWriter) (*CloneSaveCartUsingPOSTOK, *CloneSaveCartUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloneSaveCartUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cloneSaveCartUsingPOST",
		Method:             "POST",
		PathPattern:        "/{baseSiteId}/users/{userId}/carts/{cartId}/clonesavedcart",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CloneSaveCartUsingPOSTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CloneSaveCartUsingPOSTOK:
		return value, nil, nil
	case *CloneSaveCartUsingPOSTCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
FlagForDeletionUsingPATCH flags a cart for deletion

Flags a cart for deletion (the cart doesn't have corresponding save cart attributes anymore). The cart is not actually deleted from the database. But with the removal of the saved cart attributes, this cart will be taken care of by the cart removal job just like any other cart.
*/
func (a *Client) FlagForDeletionUsingPATCH(params *FlagForDeletionUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*FlagForDeletionUsingPATCHOK, *FlagForDeletionUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlagForDeletionUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "flagForDeletionUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/{baseSiteId}/users/{userId}/carts/{cartId}/flagForDeletion",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FlagForDeletionUsingPATCHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FlagForDeletionUsingPATCHOK:
		return value, nil, nil
	case *FlagForDeletionUsingPATCHNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetSavedCartUsingGET gets a saved cart

Returns a saved cart for an authenticated user. The cart is identified using the "cartId" parameter.
*/
func (a *Client) GetSavedCartUsingGET(params *GetSavedCartUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetSavedCartUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSavedCartUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSavedCartUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/users/{userId}/carts/{cartId}/savedcart",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSavedCartUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSavedCartUsingGETOK), nil

}

/*
RestoreSavedCartUsingPATCH restores a saved cart

Restore a saved cart.
*/
func (a *Client) RestoreSavedCartUsingPATCH(params *RestoreSavedCartUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*RestoreSavedCartUsingPATCHOK, *RestoreSavedCartUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreSavedCartUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreSavedCartUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/{baseSiteId}/users/{userId}/carts/{cartId}/restoresavedcart",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RestoreSavedCartUsingPATCHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RestoreSavedCartUsingPATCHOK:
		return value, nil, nil
	case *RestoreSavedCartUsingPATCHNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
SaveCartUsingPATCH explicitlies saves a cart

Explicitly saves a cart.
*/
func (a *Client) SaveCartUsingPATCH(params *SaveCartUsingPATCHParams, authInfo runtime.ClientAuthInfoWriter) (*SaveCartUsingPATCHOK, *SaveCartUsingPATCHNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveCartUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveCartUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/{baseSiteId}/users/{userId}/carts/{cartId}/save",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveCartUsingPATCHReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SaveCartUsingPATCHOK:
		return value, nil, nil
	case *SaveCartUsingPATCHNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
