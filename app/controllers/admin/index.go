package admin

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/templateset"
)

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateSet := ctx.UserValue(templateset.TemplateSetNS).(*templateset.TemplateSet)
	templateSet.Render(ctx, "admin_index.html", nil)
}
