package msg

type Markdown struct {
	At `json:"at"`
	msgType
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}

func (m *Markdown) SetMsg(title, text string) {
	m.Markdown.Title = title
	m.Markdown.Text = text
}
func (l *Markdown) SetMsgType() {
	l.MsgType = "markdown"
}
func NewMarkdown() *Markdown {
	return &Markdown{}
}
