package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models"
)

// MountUsers all subroutes
func MountUsers(root *router.Group) {
	mwAuth := auth.BuildAuth(true)
	group := root.Group("/users")

	group.GET("/", mwAuth(handleUsersGet))
}

func handleUsersGet(ctx *fasthttp.RequestCtx) {
	users, _ := models.GetUsers()
	responses.SuccessJSON(ctx, fasthttp.StatusOK, users, len(users), 0)
}
