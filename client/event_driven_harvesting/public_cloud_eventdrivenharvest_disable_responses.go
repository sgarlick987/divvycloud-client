// Code generated by go-swagger; DO NOT EDIT.

package event_driven_harvesting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PublicCloudEventdrivenharvestDisableReader is a Reader for the PublicCloudEventdrivenharvestDisable structure.
type PublicCloudEventdrivenharvestDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicCloudEventdrivenharvestDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPublicCloudEventdrivenharvestDisableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPublicCloudEventdrivenharvestDisableOK creates a PublicCloudEventdrivenharvestDisableOK with default headers values
func NewPublicCloudEventdrivenharvestDisableOK() *PublicCloudEventdrivenharvestDisableOK {
	return &PublicCloudEventdrivenharvestDisableOK{}
}

/*PublicCloudEventdrivenharvestDisableOK handles this case with default header values.

PublicCloudEventdrivenharvestDisableOK public cloud eventdrivenharvest disable o k
*/
type PublicCloudEventdrivenharvestDisableOK struct {
}

func (o *PublicCloudEventdrivenharvestDisableOK) Error() string {
	return fmt.Sprintf("[DELETE /public/cloud/eventdrivenharvest][%d] publicCloudEventdrivenharvestDisableOK ", 200)
}

func (o *PublicCloudEventdrivenharvestDisableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
