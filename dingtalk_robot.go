package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/starxiang2/dingtalkrobot/msg/repo"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type DingtalkRobot struct {
	api       string
	secret    string
	timestamp int64
}

type dingtalkResponse struct {
	Code int    `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (d *DingtalkRobot) SendMsg(msg repo.Msg) error {
	api, _ := url.Parse(d.api)

	q := api.Query()
	q.Set("sign", d.sign())
	q.Set("timestamp", strconv.FormatInt(d.timestamp, 10))

	api.RawQuery = q.Encode()
	msg.SetMsgType()
	msgString, err := json.Marshal(msg)

	if err != nil {
		return errors.New("json 转换失败")
	}

	post, err := http.Post(api.String(), "application/json", bytes.NewReader(msgString))

	if err != nil {
		return err
	}

	resp := dingtalkResponse{}
	err = json.NewDecoder(post.Body).Decode(&resp)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}
	return nil
}

func (d *DingtalkRobot) sign() string {
	d.timestamp = time.Now().UnixMilli()
	signString := fmt.Sprintf("%d\n%s", d.timestamp, d.secret)

	mac := hmac.New(sha256.New, []byte(d.secret))
	mac.Write([]byte(signString))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
func New(token, secret string) *DingtalkRobot {
	api := fmt.Sprintf(`https://oapi.dingtalk.com/robot/send?access_token=%s`, strings.Trim(token, ""))
	return &DingtalkRobot{secret: secret, api: api}
}
