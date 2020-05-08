package responses

import (
	"fmt"
	"net/http"

	"github.com/absinsekt/pnk/utils/templateset"
)

// ErrorResponse writes to ResponseWriter error with a corresponding template or serialized payload
func ErrorResponse(res http.ResponseWriter, req *http.Request, status int, templateSet *templateset.TemplateSet) {
	contentType := req.Header.Get("Content-Type")

	if contentType == "application/json" {
		WriteJSON(res, status, map[string]interface{}{"status": "error", "todo": "move to json error builder"})
		return
	}

	errorTemplate := fmt.Sprintf("errors_%d.html", status)

	res.WriteHeader(status)
	templateSet.Render(errorTemplate, res, req, nil)
}
