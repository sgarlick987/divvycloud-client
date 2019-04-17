// Code generated by go-swagger; DO NOT EDIT.

package clouds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sgarlick987/godivvycloud/models"
)

// PublicCloudsListGetReader is a Reader for the PublicCloudsListGet structure.
type PublicCloudsListGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicCloudsListGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPublicCloudsListGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPublicCloudsListGetOK creates a PublicCloudsListGetOK with default headers values
func NewPublicCloudsListGetOK() *PublicCloudsListGetOK {
	return &PublicCloudsListGetOK{
		AccessControlAllowHeaders: "Origin, X-Requested-With, Content-Type, Accept",
		AccessControlAllowMethods: "POST, GET, OPTIONS",
		AccessControlAllowOrigin:  "*",
		ContentLength:             "263",
		ContentType:               "application/json",
		Date:                      "Wed, 25 Oct 2017 14:14:14 GMT",
		Server:                    "waitress",
	}
}

/*PublicCloudsListGetOK handles this case with default header values.

PublicCloudsListGetOK public clouds list get o k
*/
type PublicCloudsListGetOK struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	ContentLength string

	ContentType string

	Date string

	Server string

	Payload *models.ListClouds
}

func (o *PublicCloudsListGetOK) Error() string {
	return fmt.Sprintf("[GET /public/clouds/list][%d] publicCloudsListGetOK  %+v", 200, o.Payload)
}

func (o *PublicCloudsListGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header access-control-allow-headers
	o.AccessControlAllowHeaders = response.GetHeader("access-control-allow-headers")

	// response header access-control-allow-methods
	o.AccessControlAllowMethods = response.GetHeader("access-control-allow-methods")

	// response header access-control-allow-origin
	o.AccessControlAllowOrigin = response.GetHeader("access-control-allow-origin")

	// response header content-length
	o.ContentLength = response.GetHeader("content-length")

	// response header content-type
	o.ContentType = response.GetHeader("content-type")

	// response header date
	o.Date = response.GetHeader("date")

	// response header server
	o.Server = response.GetHeader("server")

	o.Payload = new(models.ListClouds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
