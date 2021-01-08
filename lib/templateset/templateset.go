package templateset

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/configuration"
	"github.com/absinsekt/pnk/lib/core"
)

// Templates main templateset
var Templates *TemplateSet

func init() {
	var err error

	Templates, err = NewTemplateSet(configuration.TemplatePath)
	core.Check(err, true)
}

// NewTemplateSet creates templateSet instance
func NewTemplateSet(templateDir string) (*TemplateSet, error) {
	templateSet := &TemplateSet{
		templateDir:   templateDir,
		templateCache: map[string]*template.Template{},
	}

	if err := templateSet.loadTemplates(); err != nil {
		return nil, err
	}

	return templateSet, nil
}

func (t *TemplateSet) loadTemplates() error {
	var shared []string
	var templates []string

	log.Info("Reloading templates")

	err := filepath.Walk(t.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			if strings.HasPrefix(info.Name(), "_") {
				shared = append(shared, path)
			} else {
				templates = append(templates, path)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	for _, tmpl := range templates {
		tmplName := path.Base(tmpl)
		grouped := append(shared, tmpl)

		t.templateCache[tmplName] = template.New(tmplName)
		t.templateCache[tmplName].ParseFiles(grouped...)
	}

	return nil
}

// Render todo
func (t *TemplateSet) Render(
	ctx *fasthttp.RequestCtx,
	templateName string,
	data map[string]interface{},
	funcs map[string]interface{},
) {
	timerStart := time.Now()

	if configuration.Debug == true {
		t.Reload()
	}

	found := t.templateCache[templateName]
	if found == nil {
		msg := fmt.Sprintf("Template `%s` nor found", templateName)
		log.Debug(msg)

		if configuration.Debug {
			ctx.WriteString(msg)
		}

		return
	}

	tmpl := found.
		Lookup(templateName).
		Funcs(funcs)

	ctx.SetContentType("text/html")

	if err := tmpl.Execute(ctx.Response.BodyWriter(), data); err != nil {
		log.Error(err)
	}

	log.Debugf(
		"Template %s rendered in %.2fms",
		templateName,
		time.Now().Sub(timerStart).Seconds()*1000,
	)
}

// Reload reloads all templates from templateDir
func (t *TemplateSet) Reload() error {
	return t.loadTemplates()
}
