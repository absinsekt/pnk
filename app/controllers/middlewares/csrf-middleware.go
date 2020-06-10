package middlewares

import (
	"github.com/valyala/fasthttp"
)

// Protect middleware to enable whitelisted methods only
func Protect(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		cookie := &fasthttp.Cookie{}
		cookie.SetKey("csrf")
		cookie.
			ctx.Response.Header.SetCookie(cookie)
	}
}
