package controllers

import (
	"strings"

	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	mw "github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/www"
	"github.com/absinsekt/pnk/lib"
	"github.com/absinsekt/pnk/lib/responses"
	ts "github.com/absinsekt/pnk/lib/templateset"
)

// NewRouter creates root loader and mounts subrouters
func NewRouter() func(*fasthttp.RequestCtx) {
	templateSet, err := ts.NewTemplateSet(cfg.TemplatePath)
	lib.Check(err, true)

	if cfg.Debug {
		return buildDevRootHandler(templateSet)
	}

	return buildProductionRootHandler(templateSet)
}

func buildProductionRootHandler(templateSet *ts.TemplateSet) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		mwAuth := mw.BuildAuth(true)

		ctx.SetUserValue(ts.TemplateSetNS, templateSet)

		if path == cfg.PathRoot {
			www.Mount(path)(ctx)
			return

		} else if strings.HasPrefix(path, cfg.PathAPI) {
			mwAuth(api.Mount(strings.TrimPrefix(path, cfg.PathAPI)))(ctx)
			return

		} else if strings.HasPrefix(path, cfg.PathAdmin) {
			admin.Mount(strings.TrimPrefix(path, cfg.PathAdmin))(ctx)
			return

		} else {
			responses.DummyResponseHandler(ctx)
		}
	}
}

func buildDevRootHandler(templateSet *ts.TemplateSet) fasthttp.RequestHandler {
	fsHandlerOptions := &fasthttp.FS{
		Root:        "../www/dist",
		PathRewrite: fasthttp.NewPathSlashesStripper(1),
	}

	staticHandler := fsHandlerOptions.NewRequestHandler()
	productionHandler := buildProductionRootHandler(templateSet)

	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())

		if path == cfg.PathFaviconIco {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return
		}

		if strings.HasPrefix(path, cfg.PathDist) {
			staticHandler(ctx)
			return
		}

		productionHandler(ctx)
	}
}
