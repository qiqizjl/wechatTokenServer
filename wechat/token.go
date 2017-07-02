package wechat

import (
	"encoding/json"
	"time"

	"fmt"

	"github.com/go-redis/redis"
)

type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

const accessTokenPre = "accessToken_"
const accessTokenURLPre = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&"

func (wechat *Wechat) GetAccessToken(refresh bool) (accessToken, error) {
	var (
		token accessToken
	)
	if refresh == false {
		info, err := client.Get(wechat.getAccessToken()).Result()
		if err != nil && err != redis.Nil {
			return token, err
		}
		json.Unmarshal([]byte(info), &token)
		if token.ExpiresIn != 0 && token.ExpiresIn <= time.Now().Unix() {
			token = accessToken{}
		}
	}
	if token.AccessToken == "" {
		err := httpGetJSON(accessTokenURLPre+"appid="+wechat.AppID+"&secret="+wechat.AppSecret, &token)
		if err != nil {
			return token, err
		}
		if token.Errcode != 0 {
			return token, fmt.Errorf(token.Errmsg)
		}
		exp := time.Duration(token.ExpiresIn + time.Now().Unix())
		token.ExpiresIn = token.ExpiresIn + time.Now().Unix()
		str, _ := json.Marshal(token)
		client.Set(wechat.getAccessToken(), string(str), exp*time.Second)
	}
	return token, nil
}

func (wechat *Wechat) getAccessToken() string {
	return accessTokenPre + wechat.AppID
}
