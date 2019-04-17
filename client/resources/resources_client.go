// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PublicResourceDetailByResourceIDGet details resource

Get all of the details about a resource including its dependencies.
*/
func (a *Client) PublicResourceDetailByResourceIDGet(params *PublicResourceDetailByResourceIDGetParams) (*PublicResourceDetailByResourceIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicResourceDetailByResourceIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PublicResourceDetailByResourceIdGet",
		Method:             "GET",
		PathPattern:        "/public/resource/{resource_id}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PublicResourceDetailByResourceIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PublicResourceDetailByResourceIDGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}