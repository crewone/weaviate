//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BatchingThingsCreateReader is a Reader for the BatchingThingsCreate structure.
type BatchingThingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchingThingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchingThingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBatchingThingsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchingThingsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBatchingThingsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchingThingsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBatchingThingsCreateOK creates a BatchingThingsCreateOK with default headers values
func NewBatchingThingsCreateOK() *BatchingThingsCreateOK {
	return &BatchingThingsCreateOK{}
}

/*BatchingThingsCreateOK handles this case with default header values.

Request succeeded, see response body to get detailed information about each batched item.
*/
type BatchingThingsCreateOK struct {
	Payload []*models.ThingsGetResponse
}

func (o *BatchingThingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] batchingThingsCreateOK  %+v", 200, o.Payload)
}

func (o *BatchingThingsCreateOK) GetPayload() []*models.ThingsGetResponse {
	return o.Payload
}

func (o *BatchingThingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchingThingsCreateUnauthorized creates a BatchingThingsCreateUnauthorized with default headers values
func NewBatchingThingsCreateUnauthorized() *BatchingThingsCreateUnauthorized {
	return &BatchingThingsCreateUnauthorized{}
}

/*BatchingThingsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type BatchingThingsCreateUnauthorized struct {
}

func (o *BatchingThingsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] batchingThingsCreateUnauthorized ", 401)
}

func (o *BatchingThingsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchingThingsCreateForbidden creates a BatchingThingsCreateForbidden with default headers values
func NewBatchingThingsCreateForbidden() *BatchingThingsCreateForbidden {
	return &BatchingThingsCreateForbidden{}
}

/*BatchingThingsCreateForbidden handles this case with default header values.

Forbidden
*/
type BatchingThingsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *BatchingThingsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] batchingThingsCreateForbidden  %+v", 403, o.Payload)
}

func (o *BatchingThingsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchingThingsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchingThingsCreateUnprocessableEntity creates a BatchingThingsCreateUnprocessableEntity with default headers values
func NewBatchingThingsCreateUnprocessableEntity() *BatchingThingsCreateUnprocessableEntity {
	return &BatchingThingsCreateUnprocessableEntity{}
}

/*BatchingThingsCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type BatchingThingsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *BatchingThingsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] batchingThingsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BatchingThingsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchingThingsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchingThingsCreateInternalServerError creates a BatchingThingsCreateInternalServerError with default headers values
func NewBatchingThingsCreateInternalServerError() *BatchingThingsCreateInternalServerError {
	return &BatchingThingsCreateInternalServerError{}
}

/*BatchingThingsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type BatchingThingsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *BatchingThingsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batching/things][%d] batchingThingsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchingThingsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BatchingThingsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*BatchingThingsCreateBody batching things create body
swagger:model BatchingThingsCreateBody
*/
type BatchingThingsCreateBody struct {

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields"`

	// things
	Things []*models.Thing `json:"things"`
}

// Validate validates this batching things create body
func (o *BatchingThingsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var batchingThingsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","class","schema","id","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		batchingThingsCreateBodyFieldsItemsEnum = append(batchingThingsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *BatchingThingsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, batchingThingsCreateBodyFieldsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *BatchingThingsCreateBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *BatchingThingsCreateBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	for i := 0; i < len(o.Things); i++ {
		if swag.IsZero(o.Things[i]) { // not required
			continue
		}

		if o.Things[i] != nil {
			if err := o.Things[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "things" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BatchingThingsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BatchingThingsCreateBody) UnmarshalBinary(b []byte) error {
	var res BatchingThingsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
