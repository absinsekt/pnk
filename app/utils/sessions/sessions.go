package sessions

import (
	"encoding/gob"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/models/user"

	"github.com/gorilla/sessions"
)

// SessionStore main session store
var SessionStore *sessions.CookieStore

func init() {
	SessionStore = sessions.NewCookieStore(
		[]byte(configuration.SessionAuthKey),
		[]byte(configuration.SessionEncryptionKey))

	SessionStore.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   30 * 60,
	}

	gob.Register(&user.SessionData{})
}
