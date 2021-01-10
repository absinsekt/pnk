package templateset

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib/configuration"
)

/*
 * TODO documentations of templates loader logic
 */

const (
	sharedPathFolder = "shared"
)

var (
	// Templates main templateset
	Templates *TemplateSet

	reTemplates = regexp.MustCompile(`{{\s*template\s*"([a-zA-Z./]+)"\s*\.*\s*}}`)
)

// InitTemplateSet creates templateSet instance
func InitTemplateSet() error {
	Templates = &TemplateSet{
		templateDir:   configuration.TemplatePath,
		templateCache: map[string]*template.Template{},
	}

	Templates.ReloadTemplates()

	if configuration.Debug {
		log.Info(Templates)
	}

	return nil
}

// Render todo
func (t *TemplateSet) Render(ctx *fasthttp.RequestCtx, templateName string, data map[string]interface{}) {
	timerStart := time.Now()

	if configuration.Debug == true {
		t.ReloadTemplates()
	}

	found := t.templateCache[templateName]
	if found == nil {
		msg := fmt.Sprintf("Template `%s` not found", templateName)
		log.Debug(msg)

		if configuration.Debug {
			ctx.WriteString(msg)
		}

		return
	}

	tmpl := found.Lookup(templateName)

	ctx.SetContentType("text/html")
	log.Info(tmpl)
	if err := tmpl.Execute(ctx.Response.BodyWriter(), data); err != nil {
		log.Error(err)
	}

	log.Debugf(
		"Template %s rendered in %.2fms",
		templateName,
		time.Now().Sub(timerStart).Seconds()*1000,
	)
}

// ReloadTemplates reloads all templates from templateDir
func (t *TemplateSet) ReloadTemplates() {
	log.Info("Reloading templates")

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
	relPath, _ := filepath.Rel(configuration.TemplatePath, templatePath)
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
