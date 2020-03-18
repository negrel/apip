package main

import (
	"fmt"
	"os"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/buaazp/fasthttprouter"
	"github.com/negrel/apip/api"
	"github.com/rs/zerolog/log"
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
	var corsAllowedOrigins string = getenv("DOMAIN_NAME", "")

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   []string{corsAllowedOrigins, "*"},
		AllowedHeaders:   []string{},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: false,
	})

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

	log.Info().Msg("Starting HTTP server on port: " + port)
	fasthttp.ListenAndServe(addr, withCors.CorsMiddleware(r.Handler))
}
