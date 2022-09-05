//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// BackupsCreateReader is a Reader for the BackupsCreate structure.
type BackupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBackupsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewBackupsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupsCreateOK creates a BackupsCreateOK with default headers values
func NewBackupsCreateOK() *BackupsCreateOK {
	return &BackupsCreateOK{}
}

/*
BackupsCreateOK handles this case with default header values.

Backup create process successfully started.
*/
type BackupsCreateOK struct {
	Payload *models.BackupCreateResponse
}

func (o *BackupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /backups/{storageName}][%d] backupsCreateOK  %+v", 200, o.Payload)
}

func (o *BackupsCreateOK) GetPayload() *models.BackupCreateResponse {
	return o.Payload
}

func (o *BackupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateUnauthorized creates a BackupsCreateUnauthorized with default headers values
func NewBackupsCreateUnauthorized() *BackupsCreateUnauthorized {
	return &BackupsCreateUnauthorized{}
}

/*
BackupsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type BackupsCreateUnauthorized struct {
}

func (o *BackupsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /backups/{storageName}][%d] backupsCreateUnauthorized ", 401)
}

func (o *BackupsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBackupsCreateForbidden creates a BackupsCreateForbidden with default headers values
func NewBackupsCreateForbidden() *BackupsCreateForbidden {
	return &BackupsCreateForbidden{}
}

/*
BackupsCreateForbidden handles this case with default header values.

Forbidden
*/
type BackupsCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *BackupsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /backups/{storageName}][%d] backupsCreateForbidden  %+v", 403, o.Payload)
}

func (o *BackupsCreateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateUnprocessableEntity creates a BackupsCreateUnprocessableEntity with default headers values
func NewBackupsCreateUnprocessableEntity() *BackupsCreateUnprocessableEntity {
	return &BackupsCreateUnprocessableEntity{}
}

/*
BackupsCreateUnprocessableEntity handles this case with default header values.

Invalid backup creation attempt.
*/
type BackupsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *BackupsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /backups/{storageName}][%d] backupsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BackupsCreateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupsCreateInternalServerError creates a BackupsCreateInternalServerError with default headers values
func NewBackupsCreateInternalServerError() *BackupsCreateInternalServerError {
	return &BackupsCreateInternalServerError{}
}

/*
BackupsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type BackupsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *BackupsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /backups/{storageName}][%d] backupsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *BackupsCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BackupsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
