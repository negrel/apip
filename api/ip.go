package api

import (
	"net"

	"github.com/valyala/fasthttp"
)

func getIPVersion(ip net.IP) string {
	if ip.To4() == nil {
		return "6"
	}

	return "4"
}

// GetIP handle the request and send back the public ip address.
func GetIP(ctx *fasthttp.RequestCtx) Response {
	var header *fasthttp.RequestHeader = &ctx.Request.Header

	// Standard forward header
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded
	var ip net.IP = net.ParseIP(string(header.Peek(fasthttp.HeaderForwarded)))
	if len(ip) != 0 {
		Log.Info().Msg("IP Found via standard Forwarded header field.")
		goto ipFound
	}

	// De-facto standard header
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	ip = net.ParseIP(string(header.Peek(fasthttp.HeaderXForwardedFor)))

	if len(ip) != 0 {
		Log.Info().Msg("IP Found via non-standard X-Forwarded-For header field.")
		goto ipFound
	}

	// No proxy
	ip = ctx.RemoteIP()

	// Valid ip address (different from 0.0.0.0 and ::0)
	if !ip.Equal(net.IPv4zero) && !ip.Equal(net.IPv6zero) {
		Log.Info().Msg("IP Found using remote ip. (no proxy)")
		goto ipFound
	}

	// ---------------------------------------------------
	// IP NOT FOUND
	// ---------------------------------------------------

	Log.Error().Msg("IP NOT FOUND")

	ctx.SetStatusCode(500)

	return response{
		"code":    "500",
		"message": "Server failed to identify your ip address.",
	}

	// ---------------------------------------------------
	// IP FOUND
	// ---------------------------------------------------

ipFound:
	res := response{
		"ip":      ip.String(),
		"version": getIPVersion(ip),
	}

	Log.Info().Str("ip", res["ip"]).Str("version", res["version"]).Msg("")

	return res
}
