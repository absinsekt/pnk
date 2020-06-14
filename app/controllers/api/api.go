package api

// import (
// 	"github.com/valyala/fasthttp"

// 	cfg "github.com/absinsekt/pnk/configuration"
// 	mw "github.com/absinsekt/pnk/controllers/middlewares"
// 	"github.com/absinsekt/pnk/lib/core"
// )

// // Mount all subroutes
// func Mount() core.Subrouter {
// 	return func(path string) fasthttp.RequestHandler {
// 		switch path {
// 		case cfg.PathAPIAuth:
// 			return mw.StrictMethods([]string{fasthttp.MethodPost}, buildAuthHandler())
// 		}

// 		return nil
// 	}
// 	// sub := r.PathPrefix("/api").Subrouter()

// 	// sub.Use(middlewares.CSRFMiddleware)

// 	// sub.Path("/login/").Methods("POST").HandlerFunc(handleLogin)
// }
