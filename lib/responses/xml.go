package responses

import (
	"encoding/xml"
	"net/http"

	"github.com/valyala/fasthttp"
)

const (
	// ContentTypeRss application/rss+xml
	ContentTypeRss = "application/rss+xml"
	// ContentTypeAtom application/atom+xml
	ContentTypeAtom = "application/atom+xml"
	// ContentTypeTextXML text/xml
	ContentTypeTextXML = "text/xml"
)

// SuccessXML writes []byte data to ResponseWriter setting status 200 and content type
func SuccessXML(ctx *fasthttp.RequestCtx, contentType string, data interface{}) {
	var (
		buf []byte
		err error
	)

	if buf, err = xml.Marshal(data); err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.SetContentType(contentType + "; charset=utf-8")
	ctx.Write([]byte(xml.Header))
	ctx.Write(buf)
}
