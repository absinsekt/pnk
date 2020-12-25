package api

import (
	"github.com/fasthttp/router"

	"github.com/absinsekt/pnk/lib/middlewares/csrf"
)

// Mount all subroutes
func Mount(root *router.Router) {
	group := root.Group("/api")

	group.POST("/auth/", csrf.Protect(authHandler))
	group.GET("/logout/", csrf.Protect(logoutHandler))

	MountUsers(group)
}
