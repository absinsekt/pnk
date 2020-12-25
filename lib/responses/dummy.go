package responses

import (
	"github.com/valyala/fasthttp"
)

func DummyResponseHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusTeapot)
}
