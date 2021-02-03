package www

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/middlewares/csrf"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/templateset"
)

// Routes routing group mountable
type Routes struct{}

// Mount all subroutes
func (r *Routes) Mount(root *router.Router) {
	group := root.Group("")

	group.GET("/", csrf.InjectToken(r.indexHandler))
	group.ANY("/", responses.DummyResponseHandler)
}

func (r *Routes) indexHandler(ctx *fasthttp.RequestCtx) {
	if data, ok := templateset.Templates.Render("index.html", map[string]interface{}{
		csrf.TokenField: ctx.UserValue(csrf.TokenCookieName),
	}); ok {
		responses.OkResponse(ctx, data)
	}
}
