package rooms


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*WeaveRoomsDeleteOK Successful response

swagger:response weaveRoomsDeleteOK
*/
type WeaveRoomsDeleteOK struct {
}

// NewWeaveRoomsDeleteOK creates WeaveRoomsDeleteOK with default headers values
func NewWeaveRoomsDeleteOK() *WeaveRoomsDeleteOK {
	return &WeaveRoomsDeleteOK{}
}

// WriteResponse to the client
func (o *WeaveRoomsDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}
