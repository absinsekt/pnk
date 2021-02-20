package templateset

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"

	"github.com/absinsekt/pnk/lib/core"
	"github.com/absinsekt/pnk/lib/templateset/tags"
)

/*
 * TODO documentations of templates loader logic
 */

const (
	sharedPathFolder = "shared"
)

var (
	// Templates global templateset
	Templates *TemplateSet

	minifier = minify.New()
)

// InitTemplateSet creates templateSet instance
func InitTemplateSet() error {
	minifier.AddFunc("text/html", html.Minify)

	Templates = &TemplateSet{
		templateDir:   core.Config.TemplatePath,
		templateCache: map[string]*template.Template{},
	}

	Templates.ReloadTemplates(false)
	log.Debug(Templates)

	return nil
}

// IsExist checks if a given template exist and returns it
func (t *TemplateSet) IsExist(templateName string) (tmpl *template.Template, ok bool) {
	tmpl, ok = t.templateCache[templateName]

	return
}

// Render todo
func (t *TemplateSet) Render(templateName string, data map[string]interface{}) ([]byte, bool) {
	buf := bytes.Buffer{}
	timerStart := time.Now()

	if core.Config.Debug == true {
		t.ReloadTemplates(true)
	}

	if data == nil {
		data = map[string]interface{}{}
	}

	data["_"] = tags.Funcs

	if found, ok := t.IsExist(templateName); ok {
		tmpl := found.Lookup(templateName)

		if core.Config.Debug {
			if err := tmpl.Execute(&buf, data); err != nil {
				log.Error(err)
			}
		} else {
			mw := minifier.Writer("text/html", &buf)

			if err := tmpl.Execute(mw, data); err != nil {
				log.Error(err)
			}

			mw.Close()
		}

		log.Debugf(
			"Template %s rendered in %.2fms",
			templateName,
			time.Now().Sub(timerStart).Seconds()*1000,
		)
	} else {
		msg := fmt.Sprintf("Template `%s` not found", templateName)
		log.Debug(msg)

		if core.Config.Debug {
			buf.WriteString(msg)
		} else {
			return t.RenderError(http.StatusNotFound)
		}
	}

	return buf.Bytes(), true
}

// RenderError renders error page with status code set
func (t *TemplateSet) RenderError(status int) ([]byte, bool) {
	errorTemplate := fmt.Sprintf("errors/%d.html", status)

	return t.Render(errorTemplate, nil)
}

// ReloadTemplates reloads all templates from templateDir
func (t *TemplateSet) ReloadTemplates(quiet bool) {
	if !quiet {
		log.Info("Reloading templates")
	}

	tmpl, shared := t.searchTemplates()

	if err := t.loadTemplates(tmpl, shared); err != nil {
		log.Error(err)
	}
}

// Walk through templateDir and index .html-templates
// splitting them in two collections: ordinary templates and
// shared templates which can be reused by ordinaries
func (t *TemplateSet) searchTemplates() (templateFiles []string, sharedTemplateFiles []string) {
	err := filepath.Walk(t.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			if !isSharedTemplate(path) {
				templateFiles = append(templateFiles, path)
			} else {
				sharedTemplateFiles = append(sharedTemplateFiles, path)
			}
		}

		return nil
	})

	if err != nil {
		log.Error(err)
	}

	return
}

// Iterate templateFiles, attach shared templates to each of them,
// parse a final template group and append it to t's template cache
func (t *TemplateSet) loadTemplates(templateFiles []string, sharedTemplateFiles []string) error {
	for _, templateFile := range templateFiles {
		var finalTemplate *template.Template

		templateGroupFiles := append(sharedTemplateFiles, templateFile)

		for _, file := range templateGroupFiles {
			buf, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			templateContent := string(buf)
			templateName := buildTemplateName(file)

			if finalTemplate == nil {
				finalTemplate = template.New(templateName)
			}

			if finalTemplate.Name() == templateName {
				_, err = finalTemplate.Parse(string(templateContent))
			} else {
				_, err = finalTemplate.New(templateName).Parse(string(templateContent))
			}

			if err != nil {
				return err
			}
		}

		templateBundleName := buildTemplateName(templateFile)
		t.templateCache[templateBundleName] = finalTemplate
	}

	return nil
}

func splitPath(templatePath string) (string, string) {
	relPath, _ := filepath.Rel(core.Config.TemplatePath, templatePath)
	resultPath, templateName := filepath.Split(relPath)
	resultPath = strings.TrimPrefix(resultPath, "..")
	resultPath = strings.Trim(resultPath, string(os.PathSeparator))

	return resultPath, templateName
}

func isSharedTemplate(templatePath string) bool {
	pth, _ := splitPath(templatePath)
	return strings.HasPrefix(pth, sharedPathFolder)
}

func buildTemplateName(templatePath string) string {
	pth, templateName := splitPath(templatePath)

	if pth == "" || pth == sharedPathFolder {
		return templateName
	}

	return path.Join(pth, templateName)
}
