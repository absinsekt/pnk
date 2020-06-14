package admin

import (
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	mwAuth := mw.BuildAuth(true)

	if path == cfg.PathRoot {
		return mw.Get(mwAuth(indexHandler))

	} else if path == cfg.PathAdminAuth {
		return mw.Post(csrf.Protect(authHandler))
	}

	return responses.DummyResponseHandler
}
