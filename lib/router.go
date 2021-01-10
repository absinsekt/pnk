package lib

import (
	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
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

	for _, mnt := range mountPoints {
		mnt.Mount(root)
	}

	log.Debug(root.List())

	return root
}
