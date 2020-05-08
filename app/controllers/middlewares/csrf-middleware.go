package middlewares

import (
	"net/http"

	"github.com/absinsekt/pnk/configuration"

	"github.com/gorilla/csrf"
)

var csrfMiddleware = csrf.Protect([]byte(configuration.SessionEncryptionKey), csrf.Path("/"), csrf.Secure(false))

// CSRFMiddleware func
func CSRFMiddleware(next http.Handler) http.Handler {
	return csrfMiddleware(next)
}
