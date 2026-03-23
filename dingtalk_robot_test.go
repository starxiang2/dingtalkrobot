package dingtalkrobot

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/starxiang2/dingtalkrobot/msg"
)

func newRobotForTest(t *testing.T, srv *httptest.Server, secret string) *DingtalkRobot {
	t.Helper()
	u := srv.URL + "/robot/send?access_token=test-token"
	return &DingtalkRobot{
		api:    u,
		secret: secret,
		client: srv.Client(),
	}
}

func TestSendMsg_SuccessAndSign(t *testing.T) {

	const secret = "test-secret"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/robot/send") {
			t.Errorf("path = %q", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("access_token") != "test-token" {
			t.Errorf("access_token = %q", q.Get("access_token"))
		}
		tsStr := q.Get("timestamp")
		ts, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			t.Fatalf("timestamp: %v", err)
		}
		if q.Get("sign") != computeSign(ts, secret) {
			http.Error(w, "bad sign", http.StatusUnauthorized)
			return
		}
		b, _ := io.ReadAll(r.Body)
		var payload map[string]interface{}
		if err := json.Unmarshal(b, &payload); err != nil {
			t.Errorf("body json: %v", err)
		}
		if payload["msgtype"] != "text" {
			t.Errorf("msgtype = %v", payload["msgtype"])
		}
		_, _ = w.Write([]byte(`{"errcode":0,"errmsg":"ok"}`))
	}))
	defer srv.Close()

	robot := newRobotForTest(t, srv, secret)
	m := msg.NewText()
	m.SetMsg("hello")
	if err := robot.SendMsg(m); err != nil {
		t.Fatal(err)
	}
}

func TestSendMsg_APIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		ts, _ := strconv.ParseInt(q.Get("timestamp"), 10, 64)
		if q.Get("sign") != computeSign(ts, "s") {
			http.Error(w, "sign", 401)
			return
		}
		_, _ = w.Write([]byte(`{"errcode":310000,"errmsg":"invalid"}`))
	}))
	defer srv.Close()

	robot := newRobotForTest(t, srv, "s")
	err := robot.SendMsg(msg.NewText())
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "invalid") {
		t.Fatalf("error = %v", err)
	}
}

func TestSendMsg_InvalidJSONResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`not json`))
	}))
	defer srv.Close()

	robot := newRobotForTest(t, srv, "")
	err := robot.SendMsg(msg.NewText())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewWithHTTPClient_OverridesTimeout(t *testing.T) {
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		q := r.URL.Query()
		ts, _ := strconv.ParseInt(q.Get("timestamp"), 10, 64)
		if q.Get("sign") != computeSign(ts, "") {
			http.Error(w, "sign", 401)
			return
		}
		_, _ = w.Write([]byte(`{"errcode":0,"errmsg":"ok"}`))
	}))
	defer slow.Close()

	u := slow.URL + "/robot/send?access_token=x"
	r := &DingtalkRobot{
		api:    u,
		secret: "",
		client: slow.Client(),
	}
	if err := r.SendMsg(msg.NewText()); err != nil {
		t.Fatal(err)
	}

	short := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		_, _ = w.Write([]byte(`{"errcode":0}`))
	}))
	defer short.Close()

	u2 := short.URL + "/robot/send?access_token=x"
	r2 := &DingtalkRobot{
		api:    u2,
		secret: "",
		client: &http.Client{Timeout: 50 * time.Millisecond},
	}
	if err := r2.SendMsg(msg.NewText()); err == nil {
		t.Fatal("expected timeout error")
	}
}
