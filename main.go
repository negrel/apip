package main

import (
	"fmt"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/negrel/apip/api"
	"github.com/valyala/fasthttp"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

func main() {
	var port string = getenv("PORT", "3000")
	var addr string = fmt.Sprintf(":%v", port)

	r := fasthttprouter.New()

	// Plain text api
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ip := api.GetIP(ctx)

		ctx.SetBodyString(ip.ToTXT())
		ctx.SetContentType("text/plain")
	})

	// JSON api
	r.GET("/json", func(ctx *fasthttp.RequestCtx) {
		ip := api.GetIP(ctx)

		ctx.SetBodyString(ip.ToJSON())
		ctx.SetContentType("application/json")
	})

	// XML api
	r.GET("/xml", func(ctx *fasthttp.RequestCtx) {
		ip := api.GetIP(ctx)

		ctx.SetBodyString(ip.ToXML())
		ctx.SetContentType("text/xml")
	})

	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(404)
	}

	r.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(405)
	}

	fasthttp.ListenAndServe(addr, r.Handler)
}
