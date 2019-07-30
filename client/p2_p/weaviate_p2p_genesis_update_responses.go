//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateP2pGenesisUpdateReader is a Reader for the WeaviateP2pGenesisUpdate structure.
type WeaviateP2pGenesisUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateP2pGenesisUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateP2pGenesisUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateP2pGenesisUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateP2pGenesisUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateP2pGenesisUpdateOK creates a WeaviateP2pGenesisUpdateOK with default headers values
func NewWeaviateP2pGenesisUpdateOK() *WeaviateP2pGenesisUpdateOK {
	return &WeaviateP2pGenesisUpdateOK{}
}

/*WeaviateP2pGenesisUpdateOK handles this case with default header values.

Alive and kicking!
*/
type WeaviateP2pGenesisUpdateOK struct {
}

func (o *WeaviateP2pGenesisUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] weaviateP2pGenesisUpdateOK ", 200)
}

func (o *WeaviateP2pGenesisUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateP2pGenesisUpdateUnauthorized creates a WeaviateP2pGenesisUpdateUnauthorized with default headers values
func NewWeaviateP2pGenesisUpdateUnauthorized() *WeaviateP2pGenesisUpdateUnauthorized {
	return &WeaviateP2pGenesisUpdateUnauthorized{}
}

/*WeaviateP2pGenesisUpdateUnauthorized handles this case with default header values.

Unauthorized update.
*/
type WeaviateP2pGenesisUpdateUnauthorized struct {
}

func (o *WeaviateP2pGenesisUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] weaviateP2pGenesisUpdateUnauthorized ", 401)
}

func (o *WeaviateP2pGenesisUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateP2pGenesisUpdateInternalServerError creates a WeaviateP2pGenesisUpdateInternalServerError with default headers values
func NewWeaviateP2pGenesisUpdateInternalServerError() *WeaviateP2pGenesisUpdateInternalServerError {
	return &WeaviateP2pGenesisUpdateInternalServerError{}
}

/*WeaviateP2pGenesisUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateP2pGenesisUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateP2pGenesisUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /p2p/genesis][%d] weaviateP2pGenesisUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateP2pGenesisUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
