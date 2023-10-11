package main

import (
	"fmt"
	"github.com/starxiang2/dingtalkrobot/msg"
	"testing"
)

func TestFeedCard(t *testing.T) {
	ding := New("", "")
	msgData := msg.NewFeedCardLink()
	msgData.SetFeedCardLinks([]msg.FeedCardLink{
		msg.FeedCardLink{
			Title:      "Title",
			MessageUrl: "https://www.dealam.com",
			PicUrl:     "https://img.dealam.com/2023/10/d1d4b84b0f595ed5f3f23ac869a4c5ae.jpg",
		},

		msg.FeedCardLink{
			Title:      "Title2",
			MessageUrl: "https://www.dealam.com",
			PicUrl:     "https://img.dealam.com/2023/10/5c6b8535c9c4783d676df77219d93a9c.jpg",
		},
		msg.FeedCardLink{
			Title:      "Title3",
			MessageUrl: "https://www.dealam.com",
			PicUrl:     "https://img.dealam.com/2022/12d57a70b13494476d5a930d44076deaac.jpg",
		},
	})
	fmt.Println(ding.SendMsg(msgData))

}
