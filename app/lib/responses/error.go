package responses

import (
	"fmt"

	"github.com/absinsekt/pnk/lib/templateset"
	"github.com/valyala/fasthttp"
)

// ErrorResponse writes to ResponseWriter error with a corresponding template or serialized payload
func ErrorResponse(ctx *fasthttp.RequestCtx, status int) {
	contentType := ctx.Request.Header.ContentType()

	if string(contentType) == "application/json" {
		ErrorJSON(ctx, status)
		return
	}

	errorTemplate := fmt.Sprintf("errors_%d.html", status)

	ctx.Response.SetStatusCode(status)
	templateset.Templates.Render(ctx, errorTemplate, nil)
}
