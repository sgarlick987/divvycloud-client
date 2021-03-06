// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostReader is a Reader for the PrototypeDomainOrganizationUpdateByOrganizationResourceIDPost structure.
type PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK creates a PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK with default headers values
func NewPrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK() *PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK {
	return &PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK{}
}

/*PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK handles this case with default header values.

PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK prototype domain organization update by organization resource Id post o k
*/
type PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK struct {
}

func (o *PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK) Error() string {
	return fmt.Sprintf("[POST /prototype/domain/organization/{organization_resource_id}/update][%d] prototypeDomainOrganizationUpdateByOrganizationResourceIdPostOK ", 200)
}

func (o *PrototypeDomainOrganizationUpdateByOrganizationResourceIDPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
