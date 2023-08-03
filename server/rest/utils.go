package rest

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func allowMethods(next fasthttp.RequestHandler, allows ...string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		method := string(ctx.Method())
		if contains(method, allows) {
			next(ctx)
			return
		}

		msg := fmt.Sprintf("%s is not allowed", method)
		ctx.Error(msg, fasthttp.StatusMethodNotAllowed)
	}
}

func contains(method string, allows []string) bool {
	for _, m := range allows {
		if m == method {
			return true
		}
	}
	return false
}
