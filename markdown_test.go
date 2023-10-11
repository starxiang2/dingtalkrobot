package dingtalkrobot

import (
	"github.com/starxiang2/dingtalkrobot/msg"
	"log"
	"testing"
)

func TestMarkdown(t *testing.T) {
	ding := New("", "")

	text := "#### 杭州天气 @150XXXXXXXX \n > 9度，西北风1级，空气良89，相对温度73%\n > ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n > ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
	msgData := msg.NewMarkdown()

	msgData.SetMsg("Markdown title", text)
	err := ding.SendMsg(msgData)
	if err != nil {
		log.Fatal(err.Error())
	}
}
