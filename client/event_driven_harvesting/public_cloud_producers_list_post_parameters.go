// Code generated by go-swagger; DO NOT EDIT.

package event_driven_harvesting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPublicCloudProducersListPostParams creates a new PublicCloudProducersListPostParams object
// with the default values initialized.
func NewPublicCloudProducersListPostParams() *PublicCloudProducersListPostParams {
	var ()
	return &PublicCloudProducersListPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicCloudProducersListPostParamsWithTimeout creates a new PublicCloudProducersListPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicCloudProducersListPostParamsWithTimeout(timeout time.Duration) *PublicCloudProducersListPostParams {
	var ()
	return &PublicCloudProducersListPostParams{

		timeout: timeout,
	}
}

// NewPublicCloudProducersListPostParamsWithContext creates a new PublicCloudProducersListPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicCloudProducersListPostParamsWithContext(ctx context.Context) *PublicCloudProducersListPostParams {
	var ()
	return &PublicCloudProducersListPostParams{

		Context: ctx,
	}
}

// NewPublicCloudProducersListPostParamsWithHTTPClient creates a new PublicCloudProducersListPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicCloudProducersListPostParamsWithHTTPClient(client *http.Client) *PublicCloudProducersListPostParams {
	var ()
	return &PublicCloudProducersListPostParams{
		HTTPClient: client,
	}
}

/*PublicCloudProducersListPostParams contains all the parameters to send to the API endpoint
for the public cloud producers list post operation typically these are written to a http.Request
*/
type PublicCloudProducersListPostParams struct {

	/*ContentType*/
	ContentType string
	/*XAuthToken*/
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) WithTimeout(timeout time.Duration) *PublicCloudProducersListPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) WithContext(ctx context.Context) *PublicCloudProducersListPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) WithHTTPClient(client *http.Client) *PublicCloudProducersListPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) WithContentType(contentType string) *PublicCloudProducersListPostParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithXAuthToken adds the xAuthToken to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) WithXAuthToken(xAuthToken string) *PublicCloudProducersListPostParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the public cloud producers list post params
func (o *PublicCloudProducersListPostParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *PublicCloudProducersListPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
