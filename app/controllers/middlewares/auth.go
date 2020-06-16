package middlewares

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
)

const (
	// SessionNS session namespace
	SessionNS = "pnk_squat"
)

// SessionData holds info about current user
type SessionData struct {
	UserID         int64
	IsActive       bool
	IsStaff        bool
	SessionVersion string
}

// BuildAuth factory for building middleware
func BuildAuth(staffOnly bool) func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			cookie := string(ctx.Request.Header.Cookie(SessionNS))
			session := SessionData{}

			if err := configuration.SecureVault.Decode(SessionNS, cookie, &session); err != nil {
				deny(ctx)
				return
			}

			if session.SessionVersion != configuration.SessionVersion {
				deny(ctx)
				return
			}

			if staffOnly && !session.IsStaff {
				deny(ctx)
				return
			}

			ctx.Response.Header.Add("Vary", "Cookie")
			ctx.SetUserValue(SessionNS, session)
			next(ctx)
		}
	}
}

func deny(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue(SessionNS, nil)
	responses.ClearCookie(ctx, SessionNS)
	responses.ClearCookie(ctx, csrf.TokenCookieName)
	responses.ErrorResponse(ctx, fasthttp.StatusForbidden)
}
