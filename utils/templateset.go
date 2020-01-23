package utils

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// TemplateSet proccessed templates container
type TemplateSet struct {
	templateCache *template.Template
}

// NewTemplateSet creates templateSet instance
func NewTemplateSet(templateDir string) (*TemplateSet, error) {
	var templates []string

	templateSet := &TemplateSet{}

	log.Debugf("Loading templates")

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(info.Name(), ".html") {
			templates = append(templates, path)

			log.Debugf("Found template %s", path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	templateSet.templateCache = template.New("templateCache")
	templateSet.templateCache.ParseFiles(templates...)

	return templateSet, nil
}

// Render writes rendered html to an io.Writer
func (t *TemplateSet) Render(templateName string, w io.Writer, ctx interface{}) {
	timerStart := time.Now()

	template := t.templateCache.Lookup(templateName)
	template.Execute(w, ctx)

	log.Debugf(
		"Template %s rendered in %.2fms",
		templateName,
		time.Now().Sub(timerStart).Seconds()*1000,
	)
}
