package errortypes

import (
	"github.com/dropbox/godropbox/errors"
)

type RequestError struct {
	errors.DropboxError
}
