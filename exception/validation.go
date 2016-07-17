package exception

import (
	"errors"
	"net/http"
)

var (
	errInvalidToken      = errors.New("This token did not come from Slack.")
	errParseSlackRequest = errors.New("Can not read the form Slack sent.")
	errCmdNotFound       = errors.New("The command you are looking for is not found. Please use `help` to see possible commands.")
	errOAuthAccessDenied = errors.New("PokeGo sees that you denied access. I am sorry you are not intrested.")
	errNotANumber        = errors.New("That was not a valid number.")
	errTeamNotFound      = errors.New("Could not find the team you are looking for.")
)

//NewInvalidTokenError creates a new invalid token error
func NewInvalidTokenError() *Exception {
	return NewExceptionFromError(errInvalidToken, http.StatusUnauthorized)
}

//NewParseSlackRequestError creates a new slack request error
func NewParseSlackRequestError() *Exception {
	return NewExceptionFromError(errParseSlackRequest, http.StatusBadRequest)
}

//NewCmdNotFoundError used when a command is not found
func NewCmdNotFoundError() *Exception {
	return NewExceptionFromError(errCmdNotFound, http.StatusBadRequest)
}

//NewOAuthAccessDeniedError is when a user cancels the OAuth handshake
func NewOAuthAccessDeniedError() *Exception {
	return NewExceptionFromError(errOAuthAccessDenied, http.StatusOK)
}

//NewNotANumberError is when trying to parse a string into a number
func NewNotANumberError() *Exception {
	return NewExceptionFromError(errNotANumber, http.StatusBadRequest)
}

//NewTeamNotFoundError is when trying to parse a string into a number
func NewTeamNotFoundError() *Exception {
	return NewExceptionFromError(errTeamNotFound, http.StatusBadRequest)
}

//IsInvalidTokenErr checks if err is an invalid token error
func IsInvalidTokenErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errInvalidToken.Error()
}

//IsParseSlackRequestErr checks if err is an slack request error
func IsParseSlackRequestErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errParseSlackRequest.Error()
}

//IsCmdNotFoundErr checks if err is a command not cound error
func IsCmdNotFoundErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errCmdNotFound.Error()
}

//IsOAuthAccessDeniedErr checks if err is a command not cound error
func IsOAuthAccessDeniedErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errOAuthAccessDenied.Error()
}

//IsNotANumberErr checks if err is not a number error
func IsNotANumberErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errNotANumber.Error()
}

//IsTeamNotFoundErr checks if err is a team not found error
func IsTeamNotFoundErr(err error) bool {
	e, ok := IsException(err)
	return ok && e.Err.Error() == errTeamNotFound.Error()
}
