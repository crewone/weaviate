//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/weaviate/weaviate/entities/models"
)

// AssignRoleReader is a Reader for the AssignRole structure.
type AssignRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssignRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssignRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignRoleOK creates a AssignRoleOK with default headers values
func NewAssignRoleOK() *AssignRoleOK {
	return &AssignRoleOK{}
}

/*
AssignRoleOK describes a response with status code 200, with default header values.

Role assigned successfully
*/
type AssignRoleOK struct {
}

// IsSuccess returns true when this assign role o k response has a 2xx status code
func (o *AssignRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign role o k response has a 3xx status code
func (o *AssignRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role o k response has a 4xx status code
func (o *AssignRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign role o k response has a 5xx status code
func (o *AssignRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role o k response a status code equal to that given
func (o *AssignRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign role o k response
func (o *AssignRoleOK) Code() int {
	return 200
}

func (o *AssignRoleOK) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleOK ", 200)
}

func (o *AssignRoleOK) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleOK ", 200)
}

func (o *AssignRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRoleBadRequest creates a AssignRoleBadRequest with default headers values
func NewAssignRoleBadRequest() *AssignRoleBadRequest {
	return &AssignRoleBadRequest{}
}

/*
AssignRoleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AssignRoleBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role bad request response has a 2xx status code
func (o *AssignRoleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role bad request response has a 3xx status code
func (o *AssignRoleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role bad request response has a 4xx status code
func (o *AssignRoleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role bad request response has a 5xx status code
func (o *AssignRoleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role bad request response a status code equal to that given
func (o *AssignRoleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the assign role bad request response
func (o *AssignRoleBadRequest) Code() int {
	return 400
}

func (o *AssignRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleBadRequest  %+v", 400, o.Payload)
}

func (o *AssignRoleBadRequest) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleBadRequest  %+v", 400, o.Payload)
}

func (o *AssignRoleBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRoleUnauthorized creates a AssignRoleUnauthorized with default headers values
func NewAssignRoleUnauthorized() *AssignRoleUnauthorized {
	return &AssignRoleUnauthorized{}
}

/*
AssignRoleUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type AssignRoleUnauthorized struct {
}

// IsSuccess returns true when this assign role unauthorized response has a 2xx status code
func (o *AssignRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role unauthorized response has a 3xx status code
func (o *AssignRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role unauthorized response has a 4xx status code
func (o *AssignRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role unauthorized response has a 5xx status code
func (o *AssignRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role unauthorized response a status code equal to that given
func (o *AssignRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the assign role unauthorized response
func (o *AssignRoleUnauthorized) Code() int {
	return 401
}

func (o *AssignRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleUnauthorized ", 401)
}

func (o *AssignRoleUnauthorized) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleUnauthorized ", 401)
}

func (o *AssignRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRoleForbidden creates a AssignRoleForbidden with default headers values
func NewAssignRoleForbidden() *AssignRoleForbidden {
	return &AssignRoleForbidden{}
}

/*
AssignRoleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AssignRoleForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role forbidden response has a 2xx status code
func (o *AssignRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role forbidden response has a 3xx status code
func (o *AssignRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role forbidden response has a 4xx status code
func (o *AssignRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role forbidden response has a 5xx status code
func (o *AssignRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role forbidden response a status code equal to that given
func (o *AssignRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the assign role forbidden response
func (o *AssignRoleForbidden) Code() int {
	return 403
}

func (o *AssignRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleForbidden  %+v", 403, o.Payload)
}

func (o *AssignRoleForbidden) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleForbidden  %+v", 403, o.Payload)
}

func (o *AssignRoleForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRoleNotFound creates a AssignRoleNotFound with default headers values
func NewAssignRoleNotFound() *AssignRoleNotFound {
	return &AssignRoleNotFound{}
}

/*
AssignRoleNotFound describes a response with status code 404, with default header values.

role or user is not found.
*/
type AssignRoleNotFound struct {
}

// IsSuccess returns true when this assign role not found response has a 2xx status code
func (o *AssignRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role not found response has a 3xx status code
func (o *AssignRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role not found response has a 4xx status code
func (o *AssignRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role not found response has a 5xx status code
func (o *AssignRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role not found response a status code equal to that given
func (o *AssignRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the assign role not found response
func (o *AssignRoleNotFound) Code() int {
	return 404
}

func (o *AssignRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleNotFound ", 404)
}

func (o *AssignRoleNotFound) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleNotFound ", 404)
}

func (o *AssignRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRoleInternalServerError creates a AssignRoleInternalServerError with default headers values
func NewAssignRoleInternalServerError() *AssignRoleInternalServerError {
	return &AssignRoleInternalServerError{}
}

/*
AssignRoleInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type AssignRoleInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role internal server error response has a 2xx status code
func (o *AssignRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role internal server error response has a 3xx status code
func (o *AssignRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role internal server error response has a 4xx status code
func (o *AssignRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign role internal server error response has a 5xx status code
func (o *AssignRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assign role internal server error response a status code equal to that given
func (o *AssignRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the assign role internal server error response
func (o *AssignRoleInternalServerError) Code() int {
	return 500
}

func (o *AssignRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignRoleInternalServerError) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignRoleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AssignRoleBody assign role body
swagger:model AssignRoleBody
*/
type AssignRoleBody struct {

	// the roles that assigned to user
	Roles []string `json:"roles"`
}

// Validate validates this assign role body
func (o *AssignRoleBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this assign role body based on context it is used
func (o *AssignRoleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignRoleBody) UnmarshalBinary(b []byte) error {
	var res AssignRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
