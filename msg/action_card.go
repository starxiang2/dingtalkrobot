package msg

type Btns struct {
	Title     string `json:"title"`
	ActionUrl string `json:"actionURL"`
}

type ActionCard struct {
	msgType
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`           // markdown格式的消息。
		BtnOrientation string `json:"btnOrientation"` // 0：按钮竖直排列 1：按钮横向排列
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
		Btns           []Btns `json:"btns"`
	} `json:"actionCard"`
}

func (a *ActionCard) SetButtonType(but string) {
	a.ActionCard.BtnOrientation = but
}

func (a *ActionCard) SetSingleButton(title, url string) {
	a.ActionCard.SingleTitle = title
	a.ActionCard.SingleURL = url
}

func (a *ActionCard) SetBtns(btns []Btns) {
	a.ActionCard.Btns = btns
}

func (a *ActionCard) SetMsgType() {
	a.MsgType = "actionCard"
}

func (a *ActionCard) SetMsg(title, text string) {
	a.ActionCard.Title = title
	a.ActionCard.Text = text
}

func NewActionCard() *ActionCard {
	return &ActionCard{}
}
