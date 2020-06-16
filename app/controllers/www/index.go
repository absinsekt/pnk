package www

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/middlewares/csrf"
	"github.com/absinsekt/pnk/controllers/paths"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/templateset"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	if path == paths.PathRoot {
		return middlewares.Get(csrf.InjectToken(indexHandler))
	}

	return responses.DummyResponseHandler
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateSet := ctx.UserValue(templateset.TemplateSetNS).(*templateset.TemplateSet)

	templateSet.Render(ctx, "index.html", map[string]interface{}{
		csrf.TokenField: ctx.UserValue(csrf.TokenCookieName),
	})
}
