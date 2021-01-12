package api

import (
	"encoding/json"
	"time"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/core"
	"github.com/absinsekt/pnk/lib/middlewares/auth"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models"
)

type credentials struct {
	Login    string
	Password string
}

func authHandler(ctx *fasthttp.RequestCtx) {
	creds := credentials{}

	if err := json.Unmarshal(ctx.PostBody(), &creds); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// try to authenticate user
	usr, err := models.Auth(creds.Login, creds.Password)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		} else {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			log.Error(err)
		}

		return
	}

	sess := models.SessionData{
		ID:             usr.ID,
		Username:       usr.Username,
		Email:          usr.Email,
		IsStaff:        usr.IsStaff,
		SessionVersion: core.Config.SessionVersion,
	}

	if encoded, err := core.Config.SecureVault.Encode(auth.SessionNS, &sess); err == nil {
		responses.SetRootCookie(ctx, auth.SessionNS, encoded, 12*time.Hour)
	}

	responses.SuccessJSON(ctx, fasthttp.StatusAccepted, &sess, 1, 0)
}

func logoutHandler(ctx *fasthttp.RequestCtx) {
	auth.ClearAuth(ctx)
	ctx.Redirect("/", fasthttp.StatusFound)
}
