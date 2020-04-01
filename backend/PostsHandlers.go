package main

import "github.com/valyala/fasthttp"

func PostTestHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	_, _ = ctx.Write([]byte(`[{"text":"Всем привет!","num_likes":0,"num_comments":3,"num_shares":0,"ref_link":""},{"text":"Всем привет!","num_likes":0,"num_comments":3,"num_shares":0,"ref_link":""}]`))
}