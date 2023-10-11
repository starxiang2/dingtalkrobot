package dingtalkrobot

import (
	"github.com/starxiang2/dingtalkrobot/msg"
	"testing"
)

func TestActionCard(t *testing.T) {
	ding := New("", "")
	msgData := msg.NewActionCard()
	msgData.SetButtonType("1")

	text := `![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"`
	msgData.SetMsg("test title", text)
	//msgData.SetSingleButton("阅读全文", "https://promo.dealam.com")

	msgData.SetBtns([]msg.Btns{
		msg.Btns{Title: "阅读全文 test", ActionUrl: "https://promo.dealam.com"},
		msg.Btns{Title: "阅读全文 22222222", ActionUrl: "https://promo.dealam.com"},
		//msg.Btns{Title: "阅读全文 333333", ActionUrl: "https://www.dealam.com"},
	})
	t.Log(ding.SendMsg(msgData))
}
