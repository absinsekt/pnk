package lib

import (
	"fmt"

	"github.com/fasthttp/router"
)

// Mountable describes a routing group to be attached in the main pnk
type Mountable interface {
	Mount(*router.Router)
}

// NewRouter creates root loader and mounts subrouters
func NewRouter(mountPoints []Mountable) *router.Router {
	root := router.New()
	root.RedirectTrailingSlash = true
	root.RedirectFixedPath = false

	for _, mnt := range mountPoints {
		mnt.Mount(root)
	}

	fmt.Println(root.List())

	return root
}
