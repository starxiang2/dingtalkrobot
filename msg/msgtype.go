package msg

type msgType struct {
	MsgType string `json:"msgtype"`
}

func (m *Text) SetMsgType() {
	m.MsgType = "text"
}
