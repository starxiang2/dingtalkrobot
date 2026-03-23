package dingtalkrobot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/starxiang2/dingtalkrobot/msg/repo"
)

const defaultSendTimeout = 30 * time.Second

type DingtalkRobot struct {
	api    string
	secret string
	client *http.Client
}

type dingtalkResponse struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (d *DingtalkRobot) SendMsg(msg repo.Msg) error {
	apiURL, err := url.Parse(d.api)
	if err != nil {
		return fmt.Errorf("parse api url: %w", err)
	}

	ts := time.Now().UnixMilli()
	q := apiURL.Query()
	q.Set("sign", computeSign(ts, d.secret))
	q.Set("timestamp", strconv.FormatInt(ts, 10))
	apiURL.RawQuery = q.Encode()

	msg.SetMsgType()
	msgString, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	client := d.client
	if client == nil {
		client = http.DefaultClient
	}

	req, err := http.NewRequest(http.MethodPost, apiURL.String(), bytes.NewReader(msgString))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	resp := dingtalkResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	if resp.Code != 0 {
		return fmt.Errorf("dingtalk api: %s (code %d)", resp.Msg, resp.Code)
	}
	return nil
}

func computeSign(ts int64, secret string) string {
	signString := fmt.Sprintf("%d\n%s", ts, secret)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signString))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func New(token, secret string) *DingtalkRobot {
	api := fmt.Sprintf(`https://oapi.dingtalk.com/robot/send?access_token=%s`, strings.TrimSpace(token))
	return &DingtalkRobot{
		secret: secret,
		api:    api,
		client: &http.Client{Timeout: defaultSendTimeout},
	}
}

func NewWithHTTPClient(token, secret string, client *http.Client) *DingtalkRobot {
	r := New(token, secret)
	if client != nil {
		r.client = client
	}
	return r
}
