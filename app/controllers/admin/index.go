package admin

import (
	"github.com/valyala/fasthttp"

	ts "github.com/absinsekt/pnk/lib/templateset"
)

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateSet := ctx.UserValue(ts.TemplateSetNS).(*ts.TemplateSet)
	templateSet.Render(ctx, "admin_index.html", nil)
}
