package router

import (
	"encoding/json"

	"github.com/qiqizjl/wechatTokenServer/output"
	"github.com/qiqizjl/wechatTokenServer/wechat"
	"github.com/valyala/fasthttp"
)

type getTokenReqsut struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
	IsRefresh bool   `json:"isRefresh"`
}

type outputAccess struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

func GetToken(ctx *fasthttp.RequestCtx) {
	//解析json
	var req getTokenReqsut
	json.Unmarshal(ctx.PostBody(), &req)
	if req.AppID == "" && req.AppSecret == "" {
		ctx.WriteString(output.MakeReqParamsError())
		return
	}
	wechats := wechat.Wechat{
		AppID:     req.AppID,
		AppSecret: req.AppSecret,
	}
	token, err := wechats.GetAccessToken(req.IsRefresh)
	if err != nil {
		ctx.WriteString(output.MakeJson(output.Error, err.Error(), nil))
		return
	}

	ctx.WriteString(output.MakeSuccess(outputAccess{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
	}))
	return
}
