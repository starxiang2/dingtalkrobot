package msg

type Link struct {
	Link struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	} `json:"link"`
	msgType
}

func (l *Link) SetMsgType() {
	l.MsgType = "link"
}

func (l *Link) SetMsg(title, text, msgUrl string) {
	l.Link.Title = title
	l.Link.Text = text
	l.Link.MessageUrl = msgUrl
}
func (l *Link) SetImage(imgUrl string) {
	l.Link.PicUrl = imgUrl
}
func NewLink() *Link {
	return &Link{}
}
