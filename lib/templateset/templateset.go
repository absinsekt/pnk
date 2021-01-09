package templateset

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
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

	if err := Templates.loadTemplates(); err != nil {
		return err
	}

	if configuration.Debug {
		log.Info(Templates)
	}

	return nil
}

// Render todo
func (t *TemplateSet) Render(ctx *fasthttp.RequestCtx, templateName string, data map[string]interface{}) {
	timerStart := time.Now()

	if configuration.Debug == true {
		t.Reload()
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

// Reload reloads all templates from templateDir
func (t *TemplateSet) Reload() error {
	return t.loadTemplates()
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
			if isSharedTemplate(path) {
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
		tmplName := buildTemplateName(tmpl)
		grouped := append(shared, tmpl)

		log.Info(tmplName, grouped)

		fullTemplate, err := template.
			New(tmplName).
			Parse(concatContent(grouped))

		if err != nil {
			log.Error(err)
			return err
		}

		t.templateCache[tmplName] = fullTemplate
	}

	return nil
}

func concatContent(files []string) string {
	var (
		buf      = map[string][]byte{}
		result   = ""
		lastName = ""
	)

	for _, fileName := range files {
		data, err := ioutil.ReadFile(fileName)

		if err != nil {
			return ""
		}

		lastName = buildTemplateName(fileName)
		buf[lastName] = data
	}

	result = string(buf[lastName])

	for {
		found := reTemplates.FindStringSubmatch(result)

		if len(found) == 0 {
			break
		}

		result = strings.ReplaceAll(result, found[0], string(buf[found[1]]))
	}

	log.Info(result)

	return result
}

func buildTemplateName(fullPath string) string {
	var union = []string{}

	pathLevels, fileName := splitPath(fullPath)
	if pathLevels[0] == sharedPathFolder {
		pathLevels = pathLevels[1:]
	}

	for _, pth := range pathLevels {
		if pth != "" {
			union = append(union, pth)
		}
	}

	union = append(union, fileName)

	return strings.Join(union, string(os.PathSeparator))
}

func splitPath(fullPath string) ([]string, string) {
	relPath, _ := filepath.Rel(configuration.TemplatePath, fullPath)
	pathOnly, fileName := filepath.Split(relPath)
	pathLevels := strings.Split(pathOnly, string(os.PathSeparator))

	if pathLevels[0] == ".." {
		pathLevels = pathLevels[1:]
	}

	return pathLevels, fileName
}

func isSharedTemplate(fullPath string) bool {
	pathLevels, _ := splitPath(fullPath)

	return pathLevels[0] == sharedPathFolder
}
