package mixins

import "html/template"

// HTMLBody mixin to get raw html version of the body field
type HTMLBody struct {
	Body string
}

// HtmlBody todo
func (hb *HTMLBody) HtmlBody() template.HTML {
	return template.HTML(hb.Body)
}
