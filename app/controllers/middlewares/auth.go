package middlewares

import (
	"net/http"

	"github.com/absinsekt/pnk/utils/responses"
	"github.com/absinsekt/pnk/utils/templateset"

	"github.com/absinsekt/pnk/configuration"

	"github.com/gorilla/context"

	mdl "github.com/absinsekt/pnk/models/user"
	"github.com/absinsekt/pnk/utils/sessions"
)

// GetAuthMiddleware returns AuthMiddleware. AuthMiddleware adds user info to request's context
func GetAuthMiddleware(ts *templateset.TemplateSet) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			s, err := sessions.SessionStore.Get(req, configuration.SessionNS)
			if err != nil {
				responses.ErrorResponse(res, req, http.StatusForbidden, ts)
				return
			}

			_user := s.Values["user"]
			user, ok := _user.(*mdl.SessionData)
			if !ok {
				responses.ErrorResponse(res, req, http.StatusForbidden, ts)
				return
			}

			context.Set(req, "user", user)
			next.ServeHTTP(res, req)
		})
	}
}
