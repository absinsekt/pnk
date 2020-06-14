package core

import "github.com/valyala/fasthttp"

type Subrouter func(path string) fasthttp.RequestHandler
