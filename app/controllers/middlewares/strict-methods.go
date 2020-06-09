package middlewares

import "github.com/valyala/fasthttp"

// StrictMethods middleware to enable whitelisted methods only
func StrictMethods(methods []string, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		for _, method := range methods {
			if method == string(ctx.Method()) {
				next(ctx)
				return
			}
		}

		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	}
}
