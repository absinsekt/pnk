package responses

import (
	"time"

	"github.com/valyala/fasthttp"
)

// SetCookie builds and sets a cookie to a response header
func SetCookie(ctx *fasthttp.RequestCtx, name, value, path string, ttl time.Duration) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey(name)
	cookie.SetPath(path)
	cookie.SetValue(value)
	cookie.SetExpire(time.Now().Add(ttl))

	ctx.Response.Header.SetCookie(&cookie)
}

func ClearCookie(ctx *fasthttp.RequestCtx, name string) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey(name)
	cookie.SetMaxAge(-1)

	ctx.Response.Header.SetCookie(&cookie)
}
