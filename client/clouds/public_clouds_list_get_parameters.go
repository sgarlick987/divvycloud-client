// Code generated by go-swagger; DO NOT EDIT.

package clouds

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

// NewPublicCloudsListGetParams creates a new PublicCloudsListGetParams object
// with the default values initialized.
func NewPublicCloudsListGetParams() *PublicCloudsListGetParams {
	var ()
	return &PublicCloudsListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicCloudsListGetParamsWithTimeout creates a new PublicCloudsListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicCloudsListGetParamsWithTimeout(timeout time.Duration) *PublicCloudsListGetParams {
	var ()
	return &PublicCloudsListGetParams{

		timeout: timeout,
	}
}

// NewPublicCloudsListGetParamsWithContext creates a new PublicCloudsListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicCloudsListGetParamsWithContext(ctx context.Context) *PublicCloudsListGetParams {
	var ()
	return &PublicCloudsListGetParams{

		Context: ctx,
	}
}

// NewPublicCloudsListGetParamsWithHTTPClient creates a new PublicCloudsListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicCloudsListGetParamsWithHTTPClient(client *http.Client) *PublicCloudsListGetParams {
	var ()
	return &PublicCloudsListGetParams{
		HTTPClient: client,
	}
}

/*PublicCloudsListGetParams contains all the parameters to send to the API endpoint
for the public clouds list get operation typically these are written to a http.Request
*/
type PublicCloudsListGetParams struct {

	/*ContentType*/
	ContentType string
	/*XAuthToken*/
	XAuthToken string
	/*CloudTypeID*/
	CloudTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public clouds list get params
func (o *PublicCloudsListGetParams) WithTimeout(timeout time.Duration) *PublicCloudsListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public clouds list get params
func (o *PublicCloudsListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public clouds list get params
func (o *PublicCloudsListGetParams) WithContext(ctx context.Context) *PublicCloudsListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public clouds list get params
func (o *PublicCloudsListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public clouds list get params
func (o *PublicCloudsListGetParams) WithHTTPClient(client *http.Client) *PublicCloudsListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public clouds list get params
func (o *PublicCloudsListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the public clouds list get params
func (o *PublicCloudsListGetParams) WithContentType(contentType string) *PublicCloudsListGetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the public clouds list get params
func (o *PublicCloudsListGetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithXAuthToken adds the xAuthToken to the public clouds list get params
func (o *PublicCloudsListGetParams) WithXAuthToken(xAuthToken string) *PublicCloudsListGetParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the public clouds list get params
func (o *PublicCloudsListGetParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithCloudTypeID adds the cloudTypeID to the public clouds list get params
func (o *PublicCloudsListGetParams) WithCloudTypeID(cloudTypeID string) *PublicCloudsListGetParams {
	o.SetCloudTypeID(cloudTypeID)
	return o
}

// SetCloudTypeID adds the cloudTypeId to the public clouds list get params
func (o *PublicCloudsListGetParams) SetCloudTypeID(cloudTypeID string) {
	o.CloudTypeID = cloudTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicCloudsListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param cloud_type_id
	if err := r.SetHeaderParam("cloud_type_id", o.CloudTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
