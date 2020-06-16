package api

import (
	"strings"

	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/api/users"
	"github.com/absinsekt/pnk/lib/responses"
)

// Mount all subroutes
func Mount(path string) fasthttp.RequestHandler {
	if strings.HasPrefix(path, cfg.PathAPIUsers) {
		return users.Mount(strings.TrimPrefix(path, cfg.PathAPIUsers))
	}

	return responses.DummyResponseHandler
}
