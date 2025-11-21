package errortypes

import (
	"github.com/dropbox/godropbox/errors"
)

type ParseError struct {
	errors.DropboxError
}

type RequestError struct {
	errors.DropboxError
}
