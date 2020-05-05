package templateset

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/absinsekt/pnk/configuration"

	log "github.com/sirupsen/logrus"
)

// TemplateSet proccessed templates container
type TemplateSet struct {
	templateDir   string
	templateCache map[string]*template.Template
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

// Render writes rendered html to an io.Writer
func (t *TemplateSet) Render(templateName string, w io.Writer, ctx interface{}) {
	timerStart := time.Now()

	if configuration.Debug == true {
		t.Reload()
	}

	found := t.templateCache[templateName]
	if found == nil {
		if configuration.Debug {
			msg := fmt.Sprintf("Template `%s` nor found", templateName)

			log.Errorf(msg)
			w.Write([]byte(msg))
		}

		return
	}

	tmpl := found.Lookup(templateName)

	if err := tmpl.Execute(w, ctx); err != nil {
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
