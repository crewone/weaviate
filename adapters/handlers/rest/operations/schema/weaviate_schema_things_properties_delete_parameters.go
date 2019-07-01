/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
 * LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
 * CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateSchemaThingsPropertiesDeleteParams creates a new WeaviateSchemaThingsPropertiesDeleteParams object
// no default values defined in spec.
func NewWeaviateSchemaThingsPropertiesDeleteParams() WeaviateSchemaThingsPropertiesDeleteParams {

	return WeaviateSchemaThingsPropertiesDeleteParams{}
}

// WeaviateSchemaThingsPropertiesDeleteParams contains all the bound params for the weaviate schema things properties delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.schema.things.properties.delete
type WeaviateSchemaThingsPropertiesDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClassName string
	/*
	  Required: true
	  In: path
	*/
	PropertyName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateSchemaThingsPropertiesDeleteParams() beforehand.
func (o *WeaviateSchemaThingsPropertiesDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClassName, rhkClassName, _ := route.Params.GetOK("className")
	if err := o.bindClassName(rClassName, rhkClassName, route.Formats); err != nil {
		res = append(res, err)
	}

	rPropertyName, rhkPropertyName, _ := route.Params.GetOK("propertyName")
	if err := o.bindPropertyName(rPropertyName, rhkPropertyName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClassName binds and validates parameter ClassName from path.
func (o *WeaviateSchemaThingsPropertiesDeleteParams) bindClassName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ClassName = raw

	return nil
}

// bindPropertyName binds and validates parameter PropertyName from path.
func (o *WeaviateSchemaThingsPropertiesDeleteParams) bindPropertyName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PropertyName = raw

	return nil
}
