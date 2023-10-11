package msg

type Text struct {
	At   `json:"at"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	msgType
}

func (t *Text) SetMsg(msg string) {
	t.Text.Content = msg
}

func NewText() *Text {
	return &Text{}
}
