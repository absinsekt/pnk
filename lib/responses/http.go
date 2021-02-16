package responses

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// SuccessHTML writes []byte data to ResponseWriter setting status 200 and content type
func SuccessHTML(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetContentType("text/html")
	ctx.Write(data.([]byte))
}
