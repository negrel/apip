package api

import (
	"net"

	"github.com/valyala/fasthttp"
)

// Compare two IP address
func isEqual(ip, other net.IP) bool {
	// Compare lenght
	if len(ip) != len(other) {
		return false
	}

	// Compare each byte
	for i, len := 0, len(ip); i < len; i++ {
		if ip[i] != other[i] {
			return false
		}
	}

	return true
}

// GetIP handle the request and send back the public ip address.
func GetIP(ctx *fasthttp.RequestCtx) {
	var header *fasthttp.RequestHeader = &ctx.Request.Header

	// Standard forward header
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded
	var ip net.IP = net.ParseIP(string(header.Peek(fasthttp.HeaderForwarded)))
	if len(ip) != 0 {
		goto Send
	}

	// De-facto standard header
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	ip = net.ParseIP(string(header.Peek(fasthttp.HeaderXForwardedFor)))

	if len(ip) != 0 {
		goto Send
	}

	// No proxy
	ip = ctx.RemoteIP()

	// Valid ip address (different from 0.0.0.0 and ::0)
	if !ip.Equal(net.IPv4zero) && !ip.Equal(net.IPv6zero) {
		goto Send
	}

Send:
	ctx.Response.SetBodyString(ip.String())
}
