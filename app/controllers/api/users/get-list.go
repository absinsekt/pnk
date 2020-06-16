package users

import (
	"github.com/absinsekt/pnk/lib/responses"
	mdl "github.com/absinsekt/pnk/models/user"
	"github.com/valyala/fasthttp"
)

func GetList(ctx *fasthttp.RequestCtx) {
	users := mdl.GetList()
	responses.SuccessJSON(ctx, fasthttp.StatusOK, users, len(users), 0)
}
