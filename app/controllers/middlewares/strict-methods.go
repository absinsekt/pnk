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

func Get(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return StrictMethods([]string{fasthttp.MethodGet}, next)
}

func Post(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return StrictMethods([]string{fasthttp.MethodPost}, next)
}

func Put(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return StrictMethods([]string{fasthttp.MethodPut}, next)
}

func Delete(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return StrictMethods([]string{fasthttp.MethodDelete}, next)
}
