package model

//NewErrorMessage is to create an ErrorMessage.
func NewErrorMessage(msg string) *Response {
	res := NewPrivateResponse("")
	att := NewAttachment("")
	att.Title = "Oops!"
	att.Text = msg
	att.Color = "danger"
	res.AddAttachments(*att)
	return res
}
