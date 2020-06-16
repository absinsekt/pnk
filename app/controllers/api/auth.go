package api

import (
	"encoding/json"
	"time"

	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/models/user"
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
	usr, err := user.Auth(creds.Login, creds.Password)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		} else {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			log.Error(err)
		}

		return
	}

	sess := middlewares.SessionData{
		UserID:         usr.ID,
		IsActive:       usr.IsActive,
		IsStaff:        usr.IsStaff,
		SessionVersion: configuration.SessionVersion,
	}

	if encoded, err := configuration.SecureVault.Encode(middlewares.SessionNS, &sess); err == nil {
		responses.SetRootCookie(ctx, middlewares.SessionNS, encoded, 12*time.Hour)
	}

	responses.SuccessJSON(ctx, fasthttp.StatusAccepted, sess, 1, 0)
}
