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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// GetUsersForRoleDeprecatedReader is a Reader for the GetUsersForRoleDeprecated structure.
type GetUsersForRoleDeprecatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersForRoleDeprecatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersForRoleDeprecatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersForRoleDeprecatedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersForRoleDeprecatedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForRoleDeprecatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersForRoleDeprecatedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersForRoleDeprecatedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersForRoleDeprecatedOK creates a GetUsersForRoleDeprecatedOK with default headers values
func NewGetUsersForRoleDeprecatedOK() *GetUsersForRoleDeprecatedOK {
	return &GetUsersForRoleDeprecatedOK{}
}

/*
GetUsersForRoleDeprecatedOK describes a response with status code 200, with default header values.

Users assigned to this role
*/
type GetUsersForRoleDeprecatedOK struct {
	Payload []string
}

// IsSuccess returns true when this get users for role deprecated o k response has a 2xx status code
func (o *GetUsersForRoleDeprecatedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users for role deprecated o k response has a 3xx status code
func (o *GetUsersForRoleDeprecatedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated o k response has a 4xx status code
func (o *GetUsersForRoleDeprecatedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users for role deprecated o k response has a 5xx status code
func (o *GetUsersForRoleDeprecatedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for role deprecated o k response a status code equal to that given
func (o *GetUsersForRoleDeprecatedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users for role deprecated o k response
func (o *GetUsersForRoleDeprecatedOK) Code() int {
	return 200
}

func (o *GetUsersForRoleDeprecatedOK) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetUsersForRoleDeprecatedOK) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedOK  %+v", 200, o.Payload)
}

func (o *GetUsersForRoleDeprecatedOK) GetPayload() []string {
	return o.Payload
}

func (o *GetUsersForRoleDeprecatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForRoleDeprecatedBadRequest creates a GetUsersForRoleDeprecatedBadRequest with default headers values
func NewGetUsersForRoleDeprecatedBadRequest() *GetUsersForRoleDeprecatedBadRequest {
	return &GetUsersForRoleDeprecatedBadRequest{}
}

/*
GetUsersForRoleDeprecatedBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUsersForRoleDeprecatedBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users for role deprecated bad request response has a 2xx status code
func (o *GetUsersForRoleDeprecatedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for role deprecated bad request response has a 3xx status code
func (o *GetUsersForRoleDeprecatedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated bad request response has a 4xx status code
func (o *GetUsersForRoleDeprecatedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for role deprecated bad request response has a 5xx status code
func (o *GetUsersForRoleDeprecatedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for role deprecated bad request response a status code equal to that given
func (o *GetUsersForRoleDeprecatedBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get users for role deprecated bad request response
func (o *GetUsersForRoleDeprecatedBadRequest) Code() int {
	return 400
}

func (o *GetUsersForRoleDeprecatedBadRequest) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersForRoleDeprecatedBadRequest) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersForRoleDeprecatedBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersForRoleDeprecatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForRoleDeprecatedUnauthorized creates a GetUsersForRoleDeprecatedUnauthorized with default headers values
func NewGetUsersForRoleDeprecatedUnauthorized() *GetUsersForRoleDeprecatedUnauthorized {
	return &GetUsersForRoleDeprecatedUnauthorized{}
}

/*
GetUsersForRoleDeprecatedUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GetUsersForRoleDeprecatedUnauthorized struct {
}

// IsSuccess returns true when this get users for role deprecated unauthorized response has a 2xx status code
func (o *GetUsersForRoleDeprecatedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for role deprecated unauthorized response has a 3xx status code
func (o *GetUsersForRoleDeprecatedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated unauthorized response has a 4xx status code
func (o *GetUsersForRoleDeprecatedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for role deprecated unauthorized response has a 5xx status code
func (o *GetUsersForRoleDeprecatedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for role deprecated unauthorized response a status code equal to that given
func (o *GetUsersForRoleDeprecatedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get users for role deprecated unauthorized response
func (o *GetUsersForRoleDeprecatedUnauthorized) Code() int {
	return 401
}

func (o *GetUsersForRoleDeprecatedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedUnauthorized ", 401)
}

func (o *GetUsersForRoleDeprecatedUnauthorized) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedUnauthorized ", 401)
}

func (o *GetUsersForRoleDeprecatedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForRoleDeprecatedForbidden creates a GetUsersForRoleDeprecatedForbidden with default headers values
func NewGetUsersForRoleDeprecatedForbidden() *GetUsersForRoleDeprecatedForbidden {
	return &GetUsersForRoleDeprecatedForbidden{}
}

/*
GetUsersForRoleDeprecatedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsersForRoleDeprecatedForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users for role deprecated forbidden response has a 2xx status code
func (o *GetUsersForRoleDeprecatedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for role deprecated forbidden response has a 3xx status code
func (o *GetUsersForRoleDeprecatedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated forbidden response has a 4xx status code
func (o *GetUsersForRoleDeprecatedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for role deprecated forbidden response has a 5xx status code
func (o *GetUsersForRoleDeprecatedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for role deprecated forbidden response a status code equal to that given
func (o *GetUsersForRoleDeprecatedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get users for role deprecated forbidden response
func (o *GetUsersForRoleDeprecatedForbidden) Code() int {
	return 403
}

func (o *GetUsersForRoleDeprecatedForbidden) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersForRoleDeprecatedForbidden) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersForRoleDeprecatedForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersForRoleDeprecatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForRoleDeprecatedNotFound creates a GetUsersForRoleDeprecatedNotFound with default headers values
func NewGetUsersForRoleDeprecatedNotFound() *GetUsersForRoleDeprecatedNotFound {
	return &GetUsersForRoleDeprecatedNotFound{}
}

/*
GetUsersForRoleDeprecatedNotFound describes a response with status code 404, with default header values.

no role found
*/
type GetUsersForRoleDeprecatedNotFound struct {
}

// IsSuccess returns true when this get users for role deprecated not found response has a 2xx status code
func (o *GetUsersForRoleDeprecatedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for role deprecated not found response has a 3xx status code
func (o *GetUsersForRoleDeprecatedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated not found response has a 4xx status code
func (o *GetUsersForRoleDeprecatedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for role deprecated not found response has a 5xx status code
func (o *GetUsersForRoleDeprecatedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for role deprecated not found response a status code equal to that given
func (o *GetUsersForRoleDeprecatedNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get users for role deprecated not found response
func (o *GetUsersForRoleDeprecatedNotFound) Code() int {
	return 404
}

func (o *GetUsersForRoleDeprecatedNotFound) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedNotFound ", 404)
}

func (o *GetUsersForRoleDeprecatedNotFound) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedNotFound ", 404)
}

func (o *GetUsersForRoleDeprecatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForRoleDeprecatedInternalServerError creates a GetUsersForRoleDeprecatedInternalServerError with default headers values
func NewGetUsersForRoleDeprecatedInternalServerError() *GetUsersForRoleDeprecatedInternalServerError {
	return &GetUsersForRoleDeprecatedInternalServerError{}
}

/*
GetUsersForRoleDeprecatedInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GetUsersForRoleDeprecatedInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get users for role deprecated internal server error response has a 2xx status code
func (o *GetUsersForRoleDeprecatedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for role deprecated internal server error response has a 3xx status code
func (o *GetUsersForRoleDeprecatedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for role deprecated internal server error response has a 4xx status code
func (o *GetUsersForRoleDeprecatedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users for role deprecated internal server error response has a 5xx status code
func (o *GetUsersForRoleDeprecatedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get users for role deprecated internal server error response a status code equal to that given
func (o *GetUsersForRoleDeprecatedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get users for role deprecated internal server error response
func (o *GetUsersForRoleDeprecatedInternalServerError) Code() int {
	return 500
}

func (o *GetUsersForRoleDeprecatedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersForRoleDeprecatedInternalServerError) String() string {
	return fmt.Sprintf("[GET /authz/roles/{id}/users][%d] getUsersForRoleDeprecatedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersForRoleDeprecatedInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetUsersForRoleDeprecatedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
