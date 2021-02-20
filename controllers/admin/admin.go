package admin

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/templateset"
)

// Routes routing group mountable
type Routes struct{}

// Mount all subroutes
func (r *Routes) Mount(root *router.Router) {
	mwAuth := auth.BuildAuth(true)
	group := root.Group("/squat")

	group.GET("/", mwAuth(csrf.InjectToken(r.indexHandler)))
}

func (r *Routes) indexHandler(ctx *fasthttp.RequestCtx) {
	if data, ok := templateset.Templates.Render("admin/index.html", map[string]interface{}{
		csrf.TokenField: ctx.UserValue(csrf.TokenCookieName),
	}); ok {
		responses.SuccessHTML(ctx, data)
	}
}
