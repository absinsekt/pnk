package users

import (
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models"
	"github.com/valyala/fasthttp"
)

func GetList(ctx *fasthttp.RequestCtx) {
	users, _ := models.GetUsers()
	responses.SuccessJSON(ctx, fasthttp.StatusOK, users, len(users), 0)
}
