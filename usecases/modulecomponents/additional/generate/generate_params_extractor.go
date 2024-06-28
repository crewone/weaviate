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

package generate

import (
	"log"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/tailor-inc/graphql/language/ast"
)

var compile, _ = regexp.Compile(`{([\w\s]*?)}`)

func (p *GenerateProvider) parseGenerateArguments(args []*ast.Argument) *Params {
	out := &Params{}

	propertiesToExtract := make([]string, 0)

	for _, arg := range args {
		switch arg.Name.Value {
		case "singleResult":
			obj := arg.Value.(*ast.ObjectValue).Fields
			out.Prompt = &obj[0].Value.(*ast.StringValue).Value
			all := compile.FindAll([]byte(*out.Prompt), -1)
			for entry := range all {
				propName := string(all[entry])
				propName = strings.Trim(propName, "{")
				propName = strings.Trim(propName, "}")
				propertiesToExtract = append(propertiesToExtract, propName)
			}
		case "groupedResult":
			obj := arg.Value.(*ast.ObjectValue).Fields
			for _, field := range obj {
				switch field.Name.Value {
				case "task":
					out.Task = &field.Value.(*ast.StringValue).Value
				case "properties":
					inp := field.Value.GetValue().([]ast.Value)
					out.Properties = make([]string, len(inp))

					for i, value := range inp {
						out.Properties[i] = value.(*ast.StringValue).Value
					}
					propertiesToExtract = append(propertiesToExtract, out.Properties...)

				}
			}

		default:
			// ignore what we don't recognize
			log.Printf("Igonore not recognized value: %v", arg.Name.Value)
		}
	}

	out.PropertiesToExtract = propertiesToExtract

	return out
}

func GenerateForPrompt(textProperties map[string]string, prompt string) (string, error) {
	all := compile.FindAll([]byte(prompt), -1)
	for _, match := range all {
		originalProperty := string(match)
		replacedProperty := compile.FindStringSubmatch(originalProperty)[1]
		replacedProperty = strings.TrimSpace(replacedProperty)
		value := textProperties[replacedProperty]
		if value == "" {
			return "", errors.Errorf("Following property has empty value: '%v'. Make sure you spell the property name correctly, verify that the property exists and has a value", replacedProperty)
		}
		prompt = strings.ReplaceAll(prompt, originalProperty, value)
	}
	return prompt, nil
}
