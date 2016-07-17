package exception

import "github.com/dixonwille/PokeGoSlack/env"

//Exception is used for any error encountered in server
type Exception struct {
	Err  error
	Code int
	Msg  string
}

func (e *Exception) Error() string {
	msg := "Code " + string(e.Code)
	if e.Err != nil {
		msg += ": " + e.Err.Error()
	} else if e.Msg != "" {
		msg += ": " + e.Msg
	}
	return msg
}

//NewExceptionFromError creates an exception from an existing error
func NewExceptionFromError(err error, code int) *Exception {
	return &Exception{
		Err:  err,
		Code: code,
	}
}

//NewException creates an exception without an error
func NewException(msg string, code int) *Exception {
	return &Exception{
		Msg:  msg,
		Code: code,
	}
}

//NewExceptionAll creates and exception with all values
func NewExceptionAll(err error, msg string, code int) *Exception {
	return &Exception{
		Err:  err,
		Msg:  msg,
		Code: code,
	}
}

//IsException is used to see if the error is an Exception
func IsException(err error) (e *Exception, ok bool) {
	e, ok = err.(*Exception)
	return
}

//LogError will log the error to env.Logger
func (e *Exception) LogError() {
	env.Logger.Println(e.Error() + " " + e.Msg)
}
