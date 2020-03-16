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

	r.GET("/", api.GetIP)

	fasthttp.ListenAndServe(addr, r.Handler)
}
