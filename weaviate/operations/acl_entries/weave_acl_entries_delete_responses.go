package acl_entries


// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*WeaveACLEntriesDeleteOK Successful response

swagger:response weaveAclEntriesDeleteOK
*/
type WeaveACLEntriesDeleteOK struct {
}

// NewWeaveACLEntriesDeleteOK creates WeaveACLEntriesDeleteOK with default headers values
func NewWeaveACLEntriesDeleteOK() *WeaveACLEntriesDeleteOK {
	return &WeaveACLEntriesDeleteOK{}
}

// WriteResponse to the client
func (o *WeaveACLEntriesDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}
