package model

//ResType is what types of responses we can have.
type ResType string

const (
	//Channel response type is used to post the command and response to everyone.
	//Just include this if you have to wait till response but still want to show command.
	Channel ResType = "in_channel"
	//Issuer response type is used to post the response back only to the user that issued the command.
	Issuer ResType = "ephemeral"
)

//Response is what is returned to slack
type Response struct {
	ResponseType ResType      `json:"response_type,omitempty"`
	Text         string       `json:"text,omitempty"`
	MrkDwn       bool         `json:"mrkdwn"`
	Attachments  []Attachment `json:"attachments,omitempty"`
}

//NewPrivateResponse only returns response to issuer.
func NewPrivateResponse(msg string) *Response {
	return &Response{
		ResponseType: Issuer,
		Text:         msg,
		MrkDwn:       true,
	}
}

//NewPublicResponse return response to channel.
func NewPublicResponse(msg string) *Response {
	return &Response{
		ResponseType: Channel,
		Text:         msg,
		MrkDwn:       true,
	}
}

//RespondLater is if you want to respond back within 30 minutes
func RespondLater(withCommand bool) *Response {
	if withCommand {
		return &Response{
			ResponseType: Channel,
		}
	}
	return &Response{}
}

//AddAttachments adds attachments to the response
func (res *Response) AddAttachments(attachments ...Attachment) {
	res.Attachments = append(res.Attachments, attachments...)
}
