// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sgarlick987/godivvycloud/models"
)

// PublicResourceDetailByResourceIDGetReader is a Reader for the PublicResourceDetailByResourceIDGet structure.
type PublicResourceDetailByResourceIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicResourceDetailByResourceIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPublicResourceDetailByResourceIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPublicResourceDetailByResourceIDGetOK creates a PublicResourceDetailByResourceIDGetOK with default headers values
func NewPublicResourceDetailByResourceIDGetOK() *PublicResourceDetailByResourceIDGetOK {
	return &PublicResourceDetailByResourceIDGetOK{
		AccessControlAllowHeaders: "Origin, X-Requested-With, Content-Type, Accept",
		AccessControlAllowMethods: "POST, GET, OPTIONS",
		AccessControlAllowOrigin:  "*",
		CacheControl:              "private, no-cache, no-store, no-transform, must-revalidate, max-age=0",
		ContentEncoding:           "gzip",
		ContentLength:             "399",
		Date:                      "Wed, 25 Oct 2017 16:58:21 GMT",
		LastModified:              "2017-10-25 12:58:21.494332",
		Server:                    "waitress",
		Vary:                      "Accept-Encoding",
	}
}

/*PublicResourceDetailByResourceIDGetOK handles this case with default header values.

PublicResourceDetailByResourceIDGetOK public resource detail by resource Id get o k
*/
type PublicResourceDetailByResourceIDGetOK struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	CacheControl string

	ContentEncoding string

	ContentLength string

	Date string

	LastModified string

	Server string

	Vary string

	Payload *models.DetailResource
}

func (o *PublicResourceDetailByResourceIDGetOK) Error() string {
	return fmt.Sprintf("[GET /public/resource/{resource_id}/detail][%d] publicResourceDetailByResourceIdGetOK  %+v", 200, o.Payload)
}

func (o *PublicResourceDetailByResourceIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Encoding
	o.ContentEncoding = response.GetHeader("Content-Encoding")

	// response header Content-Length
	o.ContentLength = response.GetHeader("Content-Length")

	// response header Date
	o.Date = response.GetHeader("Date")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header Server
	o.Server = response.GetHeader("Server")

	// response header Vary
	o.Vary = response.GetHeader("Vary")

	o.Payload = new(models.DetailResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
