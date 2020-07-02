package users

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/paths"
	"github.com/absinsekt/pnk/lib/middlewares"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	if path == paths.PathRoot {
		return middlewares.Get(GetList)
	}

	return responses.DummyResponseHandler
}
