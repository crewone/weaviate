//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package vectorizer

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/semi-technologies/weaviate/entities/models"
)

// Vectorizer turns things and actions into vectors
type Vectorizer struct {
	client     client
	indexCheck IndexCheck
}

type ErrNoUsableWords struct {
	Err error
}

func (e ErrNoUsableWords) Error() string {
	return e.Err.Error()
}

func NewErrNoUsableWordsf(pattern string, args ...interface{}) ErrNoUsableWords {
	return ErrNoUsableWords{Err: fmt.Errorf(pattern, args...)}
}

type client interface {
	VectorForCorpi(ctx context.Context, corpi []string) ([]float32, error)
}

// IndexCheck returns whether a property of a class should be indexed
type IndexCheck interface {
	Indexed(className, property string) bool
	VectorizeClassName(className string) bool
	VectorizePropertyName(className, propertyName string) bool
}

// New from c11y client
func New(client client, indexCheck IndexCheck) *Vectorizer {
	return &Vectorizer{client, indexCheck}
}

func (v *Vectorizer) SetIndexChecker(ic IndexCheck) {
	v.indexCheck = ic
}

// Thing object to vector
func (v *Vectorizer) Thing(ctx context.Context, object *models.Thing) ([]float32, error) {
	return v.object(ctx, object.Class, object.Schema)
}

// Action object to vector
func (v *Vectorizer) Action(ctx context.Context, object *models.Action) ([]float32, error) {

	return v.object(ctx, object.Class, object.Schema)
}

func (v *Vectorizer) object(ctx context.Context, className string,
	schema interface{}) ([]float32, error) {
	var corpi []string

	if v.indexCheck.VectorizeClassName(className) {
		corpi = append(corpi, camelCaseToLower(className))
	}

	if schema != nil {
		for prop, value := range schema.(map[string]interface{}) {
			if !v.indexCheck.Indexed(className, prop) {
				continue
			}

			valueString, ok := value.(string)
			if ok {
				if v.indexCheck.VectorizePropertyName(className, prop) {
					// use prop and value
					corpi = append(corpi, strings.ToLower(
						fmt.Sprintf("%s %s", camelCaseToLower(prop), valueString)))
				} else {
					corpi = append(corpi, strings.ToLower(valueString))
				}
			}
		}
	}

	vector, err := v.client.VectorForCorpi(ctx, []string{strings.Join(corpi, " ")})
	if err != nil {
		switch err.(type) {
		case ErrNoUsableWords:
			return nil, fmt.Errorf("The object is invalid, as weaviate could not extract "+
				"any contextionary-valid words from it. This is the case when you do not "+
				"vectorize the class name or property names and not a single property's value "+
				"contains at least one contextionary-valid word. To fix this, you have two "+
				"options:\n\n1.) Make sure that the schema class name or the set properties are "+
				"a contextionary-valid term and include them in vectorization using the "+
				"'vectorizeClassName' or 'vectorizePropertyName' setting. In this case the vector position "+
				"will be composed of both the class/property names and the values for those fields. "+
				"Even if no property values are contextionary-valid, the overall word corpus is still valid "+
				"due to the contextionary-valid class/property names."+
				"\n\n2.) Alternatively, if you do not want to include schema class/property names "+
				"in vectorization, you must make sure that at least one text/string property contains "+
				"at least one contextionary-valid word."+
				"\n\nThe following words were extracted from your object: %v"+
				"\n\nOriginal error: %v", corpi, err)
		default:
			return nil, fmt.Errorf("vectorizing object with corpus '%+v': %v", corpi, err)
		}
	}

	return vector, nil
}

// Corpi takes any list of strings and builds a common vector for all of them
func (v *Vectorizer) Corpi(ctx context.Context, corpi []string,
) ([]float32, error) {

	for i, corpus := range corpi {
		corpi[i] = camelCaseToLower(corpus)
	}

	vector, err := v.client.VectorForCorpi(ctx, corpi)
	if err != nil {
		return nil, fmt.Errorf("vectorizing corpus '%+v': %v", corpi, err)
	}

	return vector, nil
}

func camelCaseToLower(in string) string {
	parts := camelcase.Split(in)
	var sb strings.Builder
	for i, part := range parts {
		if part == " " {
			continue
		}

		if i > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(strings.ToLower(part))
	}

	return sb.String()
}
