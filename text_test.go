package main

import (
	"fmt"
	"github.com/starxiang2/dingtalkrobot/msg"
	"testing"
)

func TestText(t *testing.T) {
	ding := New("", "")
	msgData := msg.NewText()
	msgData.SetMsg("只是一个普通的测试")
	fmt.Println(ding.SendMsg(msgData))
}
