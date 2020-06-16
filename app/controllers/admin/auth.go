package admin

import (
	"encoding/json"
	"time"

	"github.com/go-pg/pg"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
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

	sess := mw.SessionData{
		UserID:         usr.ID,
		IsActive:       usr.IsActive,
		IsStaff:        usr.IsStaff,
		SessionVersion: cfg.SessionVersion,
	}

	if encoded, err := cfg.SecureVault.Encode(mw.SessionNS, &sess); err == nil {
		responses.SetRootCookie(ctx, mw.SessionNS, encoded, 12*time.Hour)
	}

	responses.SuccessJSON(ctx, fasthttp.StatusAccepted, sess, 1, 0)
}
