package admin

import (
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	if path == cfg.PathRoot {
		return mw.Get(indexHandler)
	}

	return responses.DummyResponseHandler
}
