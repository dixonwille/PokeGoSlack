package model

//Attachment is an attachment that can be included in a response
type Attachment struct {
	Fallback   string   `json:"fallback"`
	Color      string   `json:"color,omitempty"`
	Pretext    string   `json:"pretext,omitempty"`
	AuthorName string   `json:"author_name,omitempty"`
	AuthorLink string   `json:"author_link,omitempty"`
	AuthorIcon string   `json:"author_icon,omitempty"`
	Title      string   `json:"title,omitempty"`
	TitleLink  string   `json:"title_link,omitempty"`
	Text       string   `json:"text,omitempty"`
	Fields     []Field  `json:"fields,omitempty"`
	ImageURL   string   `json:"image_url,omitempty"`
	ThumbURL   string   `json:"thumb_url,omitempty"`
	Footer     string   `json:"footer,omitempty"`
	FooterIcon string   `json:"footer_icon,omitempty"`
	TimeStamp  int      `json:"ts,omitempty"`
	MrkDwn     []string `json:"mrkdwn_in,omitempty"`
}

//NewAttachment creates an attachment with a pretext
func NewAttachment(preText string) *Attachment {
	mrkdwn := []string{"pretext", "text", "fields"}
	return &Attachment{
		Pretext: preText,
		MrkDwn:  mrkdwn,
	}
}

//AddFields adds fields to the attachment
func (att *Attachment) AddFields(fields ...Field) {
	att.Fields = append(att.Fields, fields...)
}
