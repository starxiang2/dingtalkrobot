package repo

// Msg see https://open.dingtalk.com/document/robots/custom-robot-access
type Msg interface {
	SetMsgType()
}
