package wechat

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type jsAPIticket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

const jsAPITicketPre = "jsAPITicket_"
const jsAPITicketURLPre = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&"

func (wechat *Wechat) GetJsAPITicket(refresh bool) (jsAPIticket, error) {
	var (
		token jsAPIticket
	)
	if refresh == false {
		info, err := client.Get(wechat.getAccessToken()).Result()
		if err != nil && err != redis.Nil {
			return token, err
		}
		json.Unmarshal([]byte(info), &token)
		if token.ExpiresIn != 0 && token.ExpiresIn <= time.Now().Unix() {
			token = jsAPIticket{}
		}
	}
	if token.Ticket == "" {
		access, err := wechat.GetAccessToken(false)
		if err != nil {
			return token, err
		}
		err = httpGetJSON(jsAPITicketURLPre+"access_token"+access.AccessToken, &token)
		if err != nil {
			return token, err
		}
		if token.Errcode == 40001 {
			access, err := wechat.GetAccessToken(true)
			if err != nil {
				return token, err
			}
			err = httpGetJSON(jsAPITicketURLPre+"access_token"+access.AccessToken, &token)
			if err != nil {
				return token, err
			}
		}
		if token.Errcode != 0 {
			return token, fmt.Errorf(token.Errmsg)
		}
		exp := time.Duration(token.ExpiresIn + time.Now().Unix())
		token.ExpiresIn = token.ExpiresIn + time.Now().Unix()
		str, _ := json.Marshal(token)
		err = client.Set(wechat.getJsAPITicket(), string(str), exp*time.Second).Err()
		fmt.Println(str)
	}
	return token, nil
}

func (wechat *Wechat) getJsAPITicket() string {
	return jsAPITicketPre + wechat.AppID
}
