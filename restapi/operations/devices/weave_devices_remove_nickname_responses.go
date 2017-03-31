/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

/*WeaveDevicesRemoveNicknameOK Successful response

swagger:response weaveDevicesRemoveNicknameOK
*/
type WeaveDevicesRemoveNicknameOK struct {

	// In: body
	Payload *models.DevicesRemoveNicknameResponse `json:"body,omitempty"`
}

// NewWeaveDevicesRemoveNicknameOK creates WeaveDevicesRemoveNicknameOK with default headers values
func NewWeaveDevicesRemoveNicknameOK() *WeaveDevicesRemoveNicknameOK {
	return &WeaveDevicesRemoveNicknameOK{}
}

// WithPayload adds the payload to the weave devices remove nickname o k response
func (o *WeaveDevicesRemoveNicknameOK) WithPayload(payload *models.DevicesRemoveNicknameResponse) *WeaveDevicesRemoveNicknameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave devices remove nickname o k response
func (o *WeaveDevicesRemoveNicknameOK) SetPayload(payload *models.DevicesRemoveNicknameResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveDevicesRemoveNicknameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
