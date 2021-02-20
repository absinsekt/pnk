package lib

import (
	"net/http"

	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/responses"
)

// Mountable describes a routing group to be attached in the main pnk
type Mountable interface {
	Mount(*router.Router)
}

// NewRouter creates root loader and mounts subrouters
func NewRouter(mountPoints []Mountable) *router.Router {
	log.Info("Mounting controllers")

	root := router.New()
	root.RedirectTrailingSlash = true
	root.RedirectFixedPath = false
	root.NotFound = notFoundHandler

	for _, mnt := range mountPoints {
		mnt.Mount(root)
	}

	log.Debug(root.List())

	return root
}

func notFoundHandler(ctx *fasthttp.RequestCtx) {
	responses.ErrorResponse(ctx, http.StatusNotFound)
}
