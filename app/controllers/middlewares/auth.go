package middlewares

import (
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
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
			session := &SessionData{}

			if err := cfg.SecureVault.Decode(SessionNS, cookie, &session); err != nil {
				responses.ErrorResponse(ctx, fasthttp.StatusForbidden)
				return
			}

			if session.SessionVersion != cfg.SessionVersion {
				ctx.SetUserValue(SessionNS, nil)
				responses.ClearCookie(ctx, SessionNS)
				responses.ErrorResponse(ctx, fasthttp.StatusForbidden)
				return
			}

			if staffOnly && !session.IsStaff {
				responses.ErrorResponse(ctx, fasthttp.StatusForbidden)
				return
			}

			ctx.SetUserValue(SessionNS, session)
			next(ctx)
		}
	}
}
