package responses

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type responseData struct {
	Items  []interface{} `json:"items"`
	Count  int64         `json:"count"`
	Offset int64         `json:"offset"`
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

	ctx.Response.SetStatusCode(status)

	if _, err := ctx.Response.BodyWriter().Write(raw); err != nil {
		return err
	}

	return nil
}

// SuccessJSON todo descr
func SuccessJSON(ctx *fasthttp.RequestCtx, status int, data interface{}) error {
	return writeJSON(ctx, status, &successResponse{
		Status:  "success",
		Message: "",
		Data: responseData{
			Items:  []interface{}{data},
			Count:  1,
			Offset: 0,
		},
	})
}

// ErrorJSON todo descr
func ErrorJSON(ctx *fasthttp.RequestCtx, status int) error {
	return writeJSON(ctx, status, &errorResponse{
		Status: "error",
	})
}
