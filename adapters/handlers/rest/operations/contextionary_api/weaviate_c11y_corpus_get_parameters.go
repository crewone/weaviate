/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewWeaviateC11yCorpusGetParams creates a new WeaviateC11yCorpusGetParams object
// no default values defined in spec.
func NewWeaviateC11yCorpusGetParams() WeaviateC11yCorpusGetParams {

	return WeaviateC11yCorpusGetParams{}
}

// WeaviateC11yCorpusGetParams contains all the bound params for the weaviate c11y corpus get operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.c11y.corpus.get
type WeaviateC11yCorpusGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A text corpus
	  Required: true
	  In: body
	*/
	Corpus WeaviateC11yCorpusGetBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateC11yCorpusGetParams() beforehand.
func (o *WeaviateC11yCorpusGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body WeaviateC11yCorpusGetBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("corpus", "body"))
			} else {
				res = append(res, errors.NewParseError("corpus", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Corpus = body
			}
		}
	} else {
		res = append(res, errors.Required("corpus", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
