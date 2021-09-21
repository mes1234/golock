package locker

import (
	"github.com/google/uuid"
)

type Secret struct {
	Id       uuid.UUID // Identifier of secret
	Revision int       // Revision of locker when it was modified
	Active   bool      // Flag to indicate soft delete
	Content  []byte    //encrypted content of secret
}
