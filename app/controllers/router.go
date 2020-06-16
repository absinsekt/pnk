package controllers

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/controllers/paths"
	"github.com/absinsekt/pnk/controllers/www"
	"github.com/absinsekt/pnk/lib"
	"github.com/absinsekt/pnk/lib/responses"
	"github.com/absinsekt/pnk/lib/templateset"
)

// NewRouter creates root loader and mounts subrouters
func NewRouter() func(*fasthttp.RequestCtx) {
	templateSet, err := templateset.NewTemplateSet(configuration.TemplatePath)
	lib.Check(err, true)

	if configuration.Debug {
		return buildDevRootHandler(templateSet)
	}

	return buildProductionRootHandler(templateSet)
}

func buildProductionRootHandler(templateSet *templateset.TemplateSet) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		mwAuth := middlewares.BuildAuth(true)

		ctx.SetUserValue(templateset.TemplateSetNS, templateSet)

		if path == paths.PathRoot {
			www.Mount(path)(ctx)
			return

		} else if strings.HasPrefix(path, paths.PathAPI) {
			api.Mount(strings.TrimPrefix(path, paths.PathAPI))(ctx)
			return

		} else if strings.HasPrefix(path, paths.PathAdmin) {
			mwAuth(admin.Mount(strings.TrimPrefix(path, paths.PathAdmin)))(ctx)
			return

		} else {
			responses.DummyResponseHandler(ctx)
		}
	}
}

func buildDevRootHandler(templateSet *templateset.TemplateSet) fasthttp.RequestHandler {
	fsHandlerOptions := &fasthttp.FS{
		Root:        "../www/dist",
		PathRewrite: fasthttp.NewPathSlashesStripper(1),
	}

	staticHandler := fsHandlerOptions.NewRequestHandler()
	productionHandler := buildProductionRootHandler(templateSet)

	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())

		if path == paths.PathFaviconIco {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return
		}

		if strings.HasPrefix(path, paths.PathDist) {
			staticHandler(ctx)
			return
		}

		productionHandler(ctx)
	}
}
