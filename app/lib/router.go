package lib

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	"github.com/absinsekt/pnk/controllers/www"
)

// NewRouter creates root loader and mounts subrouters
func NewRouter() func(*fasthttp.RequestCtx) {
	root := router.New()
	root.RedirectTrailingSlash = true
	root.RedirectFixedPath = false

	admin.Mount(root)
	www.Mount(root)
	api.Mount(root)

	fmt.Println(root.List())

	return root.Handler
}
