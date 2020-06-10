package www

import (
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
	ts "github.com/absinsekt/pnk/lib/templateset"
)

// Mount all subroutes
func Mount(t *ts.TemplateSet) func(string) fasthttp.RequestHandler {
	return func(path string) fasthttp.RequestHandler {
		switch path {
		case cfg.PathRoot:
			return mw.StrictMethods([]string{fasthttp.MethodGet}, buildIndexHandler(t))
		}

		return nil
	}
}

func buildIndexHandler(templateSet *ts.TemplateSet) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		templateSet.TestRender("index.html", ctx, nil)
	}
}
