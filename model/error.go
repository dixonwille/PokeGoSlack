package model

//ErrorMessage is for generic error messages sent back to server.
type ErrorMessage struct {
	Error string `json:"error"`
}

//NewErrorMessage is to create an ErrorMessage.
func NewErrorMessage(msg string) *ErrorMessage {
	return &ErrorMessage{
		Error: msg,
	}
}
