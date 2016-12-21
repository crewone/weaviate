package events


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/weaviate/models"
)

/*WeaveEventsListOK Successful response

swagger:response weaveEventsListOK
*/
type WeaveEventsListOK struct {

	// In: body
	Payload *models.EventsListResponse `json:"body,omitempty"`
}

// NewWeaveEventsListOK creates WeaveEventsListOK with default headers values
func NewWeaveEventsListOK() *WeaveEventsListOK {
	return &WeaveEventsListOK{}
}

// WithPayload adds the payload to the weave events list o k response
func (o *WeaveEventsListOK) WithPayload(payload *models.EventsListResponse) *WeaveEventsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave events list o k response
func (o *WeaveEventsListOK) SetPayload(payload *models.EventsListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveEventsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
