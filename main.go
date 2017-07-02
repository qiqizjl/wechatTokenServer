package main

import (
	"fmt"

	"github.com/qiqizjl/wechatTokenServer/output"
	"github.com/qiqizjl/wechatTokenServer/router"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe("0.0.0.0:12345", routerHandler); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}

func routerHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case "POST":
		postRouter(ctx)
	default:
		notFound(ctx)
	}
}

func postRouter(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/getAccessToken":
		router.GetToken(ctx)
	case "/getJsApiTicket":
		router.GetTicket(ctx)
	default:
		notFound(ctx)
	}
}

//路由未找到
func notFound(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(404)
	ctx.SetBodyString(output.MakeNotFound())
}
