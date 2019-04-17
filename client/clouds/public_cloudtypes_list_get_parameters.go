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

// NewPublicCloudtypesListGetParams creates a new PublicCloudtypesListGetParams object
// with the default values initialized.
func NewPublicCloudtypesListGetParams() *PublicCloudtypesListGetParams {
	var ()
	return &PublicCloudtypesListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicCloudtypesListGetParamsWithTimeout creates a new PublicCloudtypesListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicCloudtypesListGetParamsWithTimeout(timeout time.Duration) *PublicCloudtypesListGetParams {
	var ()
	return &PublicCloudtypesListGetParams{

		timeout: timeout,
	}
}

// NewPublicCloudtypesListGetParamsWithContext creates a new PublicCloudtypesListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicCloudtypesListGetParamsWithContext(ctx context.Context) *PublicCloudtypesListGetParams {
	var ()
	return &PublicCloudtypesListGetParams{

		Context: ctx,
	}
}

// NewPublicCloudtypesListGetParamsWithHTTPClient creates a new PublicCloudtypesListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicCloudtypesListGetParamsWithHTTPClient(client *http.Client) *PublicCloudtypesListGetParams {
	var ()
	return &PublicCloudtypesListGetParams{
		HTTPClient: client,
	}
}

/*PublicCloudtypesListGetParams contains all the parameters to send to the API endpoint
for the public cloudtypes list get operation typically these are written to a http.Request
*/
type PublicCloudtypesListGetParams struct {

	/*Content*/
	Content string
	/*XAuthToken*/
	XAuthToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) WithTimeout(timeout time.Duration) *PublicCloudtypesListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) WithContext(ctx context.Context) *PublicCloudtypesListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) WithHTTPClient(client *http.Client) *PublicCloudtypesListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContent adds the content to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) WithContent(content string) *PublicCloudtypesListGetParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) SetContent(content string) {
	o.Content = content
}

// WithXAuthToken adds the xAuthToken to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) WithXAuthToken(xAuthToken string) *PublicCloudtypesListGetParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the public cloudtypes list get params
func (o *PublicCloudtypesListGetParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *PublicCloudtypesListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content
	if err := r.SetHeaderParam("Content", o.Content); err != nil {
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
