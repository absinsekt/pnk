package api

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/api/users"
	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/middlewares/csrf"
	"github.com/absinsekt/pnk/controllers/paths"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	mwAuth := middlewares.BuildAuth(true)

	if path == paths.PathAPIAuth {
		return middlewares.Post(csrf.Protect(authHandler))

	} else if strings.HasPrefix(path, paths.PathAPIUsers) {
		return mwAuth(users.Mount(strings.TrimPrefix(path, paths.PathAPIUsers)))
	}

	return responses.DummyResponseHandler
}
