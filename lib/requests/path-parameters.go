package requests

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

// GetString todo
func GetString(parameterName string, ctx *fasthttp.RequestCtx) (result string, ok bool) {
	if param := ctx.UserValue(parameterName); param != nil {
		result, ok = param.(string)
	}

	return
}

// GetBool todo
func GetBool(parameterName string, ctx *fasthttp.RequestCtx) (result bool, ok bool) {
	var (
		param string
		err   error
	)

	if param, ok = GetString(parameterName, ctx); ok {
		result, err = strconv.ParseBool(param)
		ok = err == nil
	}

	return
}

// GetInt64 todo
func GetInt64(parameterName string, ctx *fasthttp.RequestCtx) (result int64, ok bool) {
	var (
		param string
		err   error
	)

	if param, ok = GetString(parameterName, ctx); ok {
		result, err = strconv.ParseInt(param, 10, 64)
		ok = err == nil
	}

	return
}

// GetFloat64 todo
func GetFloat64(parameterName string, ctx *fasthttp.RequestCtx) (result float64, ok bool) {
	var (
		param string
		err   error
	)

	if param, ok = GetString(parameterName, ctx); ok {
		result, err = strconv.ParseFloat(param, 64)
		ok = err == nil
	}

	return
}
