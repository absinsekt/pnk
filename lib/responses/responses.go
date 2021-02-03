package responses

import (
	"net/http"

	"github.com/absinsekt/pnk/lib/templateset"
	"github.com/valyala/fasthttp"
)

// OkResponse writes []byte data to ResponseWriter setting status 200 and content type
func OkResponse(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetContentType("text/html")
	ctx.Write(data.([]byte))
}

// ErrorResponse writes to ResponseWriter error with a corresponding template or serialized payload
func ErrorResponse(ctx *fasthttp.RequestCtx, status int) {
	contentType := ctx.Request.Header.ContentType()

	if string(contentType) == "application/json" {
		ErrorJSON(ctx, status)
	} else {
		ctx.SetStatusCode(status)
		ctx.SetContentType("text/html")
		data, _ := templateset.Templates.RenderError(status)
		ctx.Write(data)
	}
}
