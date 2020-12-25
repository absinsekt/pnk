package responses

import (
	"time"

	cfg "github.com/absinsekt/pnk/lib/configuration"
	"github.com/valyala/fasthttp"
)

// Cookie simple cookie type
type Cookie struct {
	Name    string
	Value   string
	Path    string
	Expires time.Time
}

// NewCookie builds a cookie instance
func NewCookie(ctx *fasthttp.RequestCtx, cookie *Cookie) *fasthttp.Cookie {
	result := fasthttp.Cookie{}

	result.SetKey(cookie.Name)
	result.SetValue(cookie.Value)
	result.SetHTTPOnly(true)
	result.SetSecure(!cfg.Debug)

	if cookie.Path == "" {
		result.SetPath("/")
	} else {
		result.SetPath(cookie.Path)
	}

	if !cookie.Expires.IsZero() {
		result.SetExpire(cookie.Expires)
	}

	return &result
}

// SetRootCookie sets prebuilt cookie to response headers
func SetRootCookie(ctx *fasthttp.RequestCtx, name, value string, ttl time.Duration) {
	ctx.Response.Header.SetCookie(NewCookie(ctx, &Cookie{
		Name:    name,
		Value:   value,
		Expires: time.Now().UTC().Add(ttl),
	}))
}

// ClearCookie removes a cookie from client
func ClearCookie(ctx *fasthttp.RequestCtx, name string) {
	ctx.Response.Header.SetCookie(NewCookie(ctx, &Cookie{
		Name:    name,
		Value:   "",
		Expires: fasthttp.CookieExpireDelete,
	}))
}
