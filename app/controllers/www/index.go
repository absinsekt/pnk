package www

import (
	"log"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/controllers/middlewares"

	"github.com/absinsekt/pnk/utils/templateset"
	"github.com/gorilla/mux"
)

// Mount attach all entry points of file
func Mount(r *mux.Router, t *templateset.TemplateSet) {
	sub := r.PathPrefix("/").Subrouter()

	sub.Use(middlewares.CSRFMiddleware)

	sub.Path("/").Methods("GET").HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *templateset.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		templateSet.Render("index.html", res, req, nil)
	}
}

func TestGetHandlerIndex(templateSet *templateset.TemplateSet) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		log.Println("OLOLO!")
		log.Println(ctx)
		templateSet.TestRender("index.html", ctx, nil)
	}
}
