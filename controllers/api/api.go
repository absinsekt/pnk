package api

import (
	"github.com/fasthttp/router"

	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/middlewares/csrf"
)

// Routes routing group mountable
type Routes struct{}

// Mount all subroutes
func (r *Routes) Mount(root *router.Router) {
	mwAuth := auth.BuildAuth(true)
	group := root.Group("/api")

	group.POST("/auth/", csrf.Protect(authHandler))
	group.GET("/logout/", csrf.Protect(logoutHandler))

	subUsers := group.Group("/users")
	subUsers.GET("/", mwAuth(handleUsersGet))
}
