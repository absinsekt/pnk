package middlewares

import (
	"net/http"

	"github.com/absinsekt/pnk/configuration"

	"github.com/gorilla/context"

	mdl "github.com/absinsekt/pnk/models/user"
	"github.com/absinsekt/pnk/utils/sessions"
)

// AuthMiddleware adds user info to request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		s, err := sessions.SessionStore.Get(req, configuration.SessionNS)
		if err != nil {
			res.WriteHeader(http.StatusForbidden)
			return
		}

		_user := s.Values["user"]
		user, ok := _user.(*mdl.SessionData)
		if !ok {
			res.WriteHeader(http.StatusForbidden)
			return
		}

		context.Set(req, "user", user)
		next.ServeHTTP(res, req)
	})
}
