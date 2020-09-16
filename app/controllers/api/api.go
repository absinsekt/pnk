package api

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/api/users"
	"github.com/absinsekt/pnk/controllers/paths"
	"github.com/absinsekt/pnk/lib/middlewares"
	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	mwAuth := auth.BuildAuth(true)

	if path == paths.PathAPIAuth {
		return middlewares.Post(csrf.Protect(authHandler))

	} else if strings.HasPrefix(path, paths.PathAPIUsers) {
		return mwAuth(users.Mount(strings.TrimPrefix(path, paths.PathAPIUsers)))
	}

	return responses.DummyResponseHandler
}
