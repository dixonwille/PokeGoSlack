package exception

import (
	"errors"
	"fmt"

	"github.com/dixonwille/PokeGoSlack/env"
)

var (
	errInternal = errors.New("Something went wrong. Contact Maintainer.")
)

//InternalErr is used for internal error
type InternalErr struct {
	err  error
	code int
	msg  string
}

//NewInternalErr returns an internal error
func NewInternalErr(code int, msg string) *InternalErr {
	return &InternalErr{
		err:  errInternal,
		code: code,
		msg:  msg,
	}
}

func (err *InternalErr) Error() string {
	return fmt.Sprintf("%s Code: %d", err.err.Error(), err.code)
}

//LogError will log the error to env.Logger
func (err *InternalErr) LogError() {
	env.Logger.Println(fmt.Sprintf("Code %d: %s", err.code, err.msg))
}
