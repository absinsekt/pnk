package responses

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type responseData struct {
	Items  interface{} `json:"items"`
	Count  int         `json:"count"`
	Offset int         `json:"offset"`
}

type successResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    responseData `json:"data"`
}

type errorResponse struct {
	Status string `json:"status"`
}

func writeJSON(ctx *fasthttp.RequestCtx, status int, data interface{}) error {
	raw, err := json.Marshal(data)

	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}

	ctx.SetStatusCode(status)
	ctx.SetContentType("application/json")
	ctx.Write(raw)

	return nil
}

// SuccessJSON todo descr
func SuccessJSON(ctx *fasthttp.RequestCtx, status int, data interface{}, count int, offset int) error {
	return writeJSON(ctx, status, &successResponse{
		Status:  "success",
		Message: "",
		Data: responseData{
			Items:  data,
			Count:  count,
			Offset: offset,
		},
	})
}

// ErrorJSON todo descr
func ErrorJSON(ctx *fasthttp.RequestCtx, status int) error {
	return writeJSON(ctx, status, &errorResponse{
		Status: "error",
	})
}
