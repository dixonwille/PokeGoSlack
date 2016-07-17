package exception

import (
	"errors"
	"net/http"
)

var (
	errInternal = errors.New("Something went wrong. Contact Maintainer.")
)

//NewInternalError creates a new internal error
func NewInternalError(msg string) *Exception {
	return NewExceptionAll(errInternal, msg, http.StatusInternalServerError)
}

//IsInternalErr is used to see if error is internal
func IsInternalErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errInternal.Error()
}
