package templateset

import "html/template"

// TemplateSet proccessed templates container
type TemplateSet struct {
	templateDir   string
	templateCache map[string]*template.Template
}
