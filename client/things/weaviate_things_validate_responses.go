// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsValidateReader is a Reader for the WeaviateThingsValidate structure.
type WeaviateThingsValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsValidateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsValidateOK creates a WeaviateThingsValidateOK with default headers values
func NewWeaviateThingsValidateOK() *WeaviateThingsValidateOK {
	return &WeaviateThingsValidateOK{}
}

/*WeaviateThingsValidateOK handles this case with default header values.

Successful validated.
*/
type WeaviateThingsValidateOK struct {
}

func (o *WeaviateThingsValidateOK) Error() string {
	return fmt.Sprintf("[POST /things/validate][%d] weaviateThingsValidateOK ", 200)
}

func (o *WeaviateThingsValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsValidateUnauthorized creates a WeaviateThingsValidateUnauthorized with default headers values
func NewWeaviateThingsValidateUnauthorized() *WeaviateThingsValidateUnauthorized {
	return &WeaviateThingsValidateUnauthorized{}
}

/*WeaviateThingsValidateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsValidateUnauthorized struct {
}

func (o *WeaviateThingsValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /things/validate][%d] weaviateThingsValidateUnauthorized ", 401)
}

func (o *WeaviateThingsValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsValidateForbidden creates a WeaviateThingsValidateForbidden with default headers values
func NewWeaviateThingsValidateForbidden() *WeaviateThingsValidateForbidden {
	return &WeaviateThingsValidateForbidden{}
}

/*WeaviateThingsValidateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsValidateForbidden struct {
}

func (o *WeaviateThingsValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /things/validate][%d] weaviateThingsValidateForbidden ", 403)
}

func (o *WeaviateThingsValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsValidateUnprocessableEntity creates a WeaviateThingsValidateUnprocessableEntity with default headers values
func NewWeaviateThingsValidateUnprocessableEntity() *WeaviateThingsValidateUnprocessableEntity {
	return &WeaviateThingsValidateUnprocessableEntity{}
}

/*WeaviateThingsValidateUnprocessableEntity handles this case with default header values.

Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateThingsValidateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsValidateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /things/validate][%d] weaviateThingsValidateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsValidateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
