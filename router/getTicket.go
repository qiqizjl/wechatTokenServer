package router

import (
	"encoding/json"

	"github.com/qiqizjl/wechatTokenServer/output"
	"github.com/qiqizjl/wechatTokenServer/wechat"
	"github.com/valyala/fasthttp"
)

type getTicketReqsut struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
	IsRefresh bool   `json:"isRefresh"`
}

type outputTicket struct {
	JsAPITicket string `json:"jsApiTicket`
	ExpiresIn   int64  `json:"expiresIn"`
}

func GetTicket(ctx *fasthttp.RequestCtx) {
	//解析json
	var req getTicketReqsut
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
	ctx.WriteString(output.MakeSuccess(outputTicket{
		JsAPITicket: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
	}))
	return
}
