package www

import (
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
	ts "github.com/absinsekt/pnk/lib/templateset"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	if path == cfg.PathRoot {
		return mw.Get(csrf.InjectToken(indexHandler))
	}

	return responses.DummyResponseHandler
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateSet := ctx.UserValue(ts.TemplateSetNS).(*ts.TemplateSet)

	templateSet.Render(ctx, "index.html", map[string]interface{}{
		csrf.TokenField: ctx.UserValue(csrf.TokenCookieName),
	})
}
