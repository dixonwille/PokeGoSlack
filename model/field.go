package model

//Field are the fields inside an Attachment
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

//NewField creates a new field to add to attachment
func NewField(title, value string, short bool) *Field {
	return &Field{
		Title: title,
		Value: value,
		Short: short,
	}
}
