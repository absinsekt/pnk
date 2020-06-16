package users

import (
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models/user"
	"github.com/valyala/fasthttp"
)

func GetList(ctx *fasthttp.RequestCtx) {
	users, _ := user.GetList()
	responses.SuccessJSON(ctx, fasthttp.StatusOK, users, len(users), 0)
}
