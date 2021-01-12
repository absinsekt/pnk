package auth

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/core"
	"github.com/absinsekt/pnk/lib/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models"
)

const (
	// SessionNS session namespace
	SessionNS = "pnk_squat"
)

// BuildAuth factory for building middleware
func BuildAuth(staffOnly bool) func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			cookie := string(ctx.Request.Header.Cookie(SessionNS))
			session := models.SessionData{}

			if err := core.Config.SecureVault.Decode(SessionNS, cookie, &session); err != nil {
				ClearAuth(ctx)
				return
			}

			if session.SessionVersion != core.Config.SessionVersion {
				ClearAuth(ctx)
				return
			}

			if staffOnly && !session.IsStaff {
				ClearAuth(ctx)
				return
			}

			ctx.Response.Header.Add("Vary", "Cookie")
			ctx.SetUserValue(SessionNS, session)
			next(ctx)
		}
	}
}

// ClearAuth clears session and auth cookies
func ClearAuth(ctx *fasthttp.RequestCtx) {
	ctx.SetUserValue(SessionNS, nil)
	responses.ClearCookie(ctx, SessionNS)
	responses.ClearCookie(ctx, csrf.TokenCookieName)
	responses.ErrorResponse(ctx, fasthttp.StatusForbidden)
}
