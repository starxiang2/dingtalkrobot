package dingtalkrobot

import (
	"encoding/json"
	"testing"

	"github.com/starxiang2/dingtalkrobot/msg"
)

func TestTextMessageJSON(t *testing.T) {
	m := msg.NewText()
	m.SetMsg("只是一个普通的测试")
	m.SetMsgType()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out["msgtype"] != "text" {
		t.Fatalf("msgtype = %v", out["msgtype"])
	}
	text, ok := out["text"].(map[string]interface{})
	if !ok {
		t.Fatalf("text field: %T", out["text"])
	}
	if text["content"] != "只是一个普通的测试" {
		t.Fatalf("content = %v", text["content"])
	}
}

func TestMarkdownMessageJSON(t *testing.T) {
	m := msg.NewMarkdown()
	m.SetMsg("Markdown title", "#### line")
	m.SetMsgType()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var out struct {
		MsgType  string `json:"msgtype"`
		Markdown struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		} `json:"markdown"`
	}
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out.MsgType != "markdown" {
		t.Fatalf("msgtype = %q", out.MsgType)
	}
	clinet := New("f62d40748653b0f5f13f4eb5409d2361ac481f8624e680cae127c6a7f7c6b1ab", "SEC75bc615d78be55b2c358154086a656265656747819b52d6c5fb20330d477638e")
	clinet.SendMsg(m)
	if out.Markdown.Title != "Markdown title" || out.Markdown.Text != "#### line" {
		t.Fatalf("markdown = %+v", out.Markdown)
	}
}

func TestActionCardMessageJSON(t *testing.T) {
	m := msg.NewActionCard()
	m.SetButtonType("1")
	m.SetMsg("test title", "body")
	m.SetBtns([]msg.Btns{
		{Title: "A", ActionUrl: "https://example.com"},
		{Title: "B", ActionUrl: "https://example.com"},
	})
	m.SetMsgType()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}

	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out["msgtype"] != "actionCard" {
		t.Fatalf("msgtype = %v", out["msgtype"])
	}
}

func TestFeedCardMessageJSON(t *testing.T) {
	m := msg.NewFeedCardLink()
	m.SetFeedCardLinks([]msg.FeedCardLink{
		{Title: "T1", MessageUrl: "https://www.baidu.com", PicUrl: "http://captain-export.captainbi.com/news/2025-09-12/e34f785906b8527f41971b2a9a00ec8120250912101905882.jpg"},
	})
	m.SetMsgType()
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out["msgtype"] != "feedCard" {
		t.Fatalf("msgtype = %v", out["msgtype"])
	}
}
