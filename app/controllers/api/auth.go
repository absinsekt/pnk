package api

import (
	"encoding/json"
	"net/http"

	"github.com/absinsekt/pnk/configuration"

	"github.com/go-pg/pg/v9"

	"github.com/absinsekt/pnk/models/user"
	"github.com/absinsekt/pnk/utils/responses"
	"github.com/absinsekt/pnk/utils/sessions"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type credentials struct {
	Login    string
	Password string
}

// MountAuth attach all entry points of file
func MountAuth(r *mux.Router) {
	r.Path("/login/").
		Methods("POST").
		HandlerFunc(handleLogin)
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	var creds credentials

	if err := json.NewDecoder(req.Body).Decode(&creds); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// try to authenticate user
	usr, err := user.Auth(creds.Login, creds.Password)
	if err != nil {
		if err == pg.ErrNoRows {
			res.WriteHeader(http.StatusNotFound)
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
		}

		return
	}

	// store usersession in cookie
	session, err := sessions.SessionStore.Get(req, configuration.SessionNS)
	if err != nil {
		log.Error(err)
	}

	session.Values["user"] = &user.SessionData{
		ID:       usr.ID,
		Username: usr.Username,
		Email:    usr.Email,
	}

	if err := session.Save(req, res); err != nil {
		log.Error(err)
		return
	}

	// done
	responses.WriteJSON(res, http.StatusAccepted, usr)
}
