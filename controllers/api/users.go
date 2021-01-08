package api

import (
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models"
)

func handleUsersGet(ctx *fasthttp.RequestCtx) {
	users, _ := models.GetUsers()
	responses.SuccessJSON(ctx, fasthttp.StatusOK, users, len(users), 0)
}
