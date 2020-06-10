package controllers

import (
	"strings"

	"github.com/valyala/fasthttp"

	cfg "github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers/www"
	"github.com/absinsekt/pnk/lib"
	ts "github.com/absinsekt/pnk/lib/templateset"
)

// NewRouter creates root loader and mounts subrouters
func NewRouter() func(*fasthttp.RequestCtx) {
	templateSet, err := ts.NewTemplateSet(cfg.TemplatePath)
	utils.Check(err, true)

	if cfg.Debug {
		return buildDevRootHandler(templateSet)
	}

	return buildProductionRootHandler(templateSet)
}

func buildProductionRootHandler(templateSet *ts.TemplateSet) fasthttp.RequestHandler {
	// create groups here
	groupWWW := www.Mount(templateSet)

	return func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())

		switch path {
		case cfg.PathRoot:
			groupWWW(path)(ctx)
			// groupAPI(strings.TrimPrefix(path, cfgPathAPI))

		default:
			ctx.SetStatusCode(fasthttp.StatusTeapot)
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
