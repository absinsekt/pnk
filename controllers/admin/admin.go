package admin

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/templateset"
)

// Mount all subroutes
func Mount(root *router.Router) {
	mwAuth := auth.BuildAuth(true)
	group := root.Group("/squat")

	group.GET("/", mwAuth(csrf.InjectToken(indexHandler)))
	group.ANY("/", responses.DummyResponseHandler)
}

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateset.Templates.Render(ctx, "admin_index.html", map[string]interface{}{
		csrf.TokenField: ctx.UserValue(csrf.TokenCookieName),
	})
}