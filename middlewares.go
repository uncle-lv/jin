package jin

import (
	"jin/log"
)

func DefaultLogger() HandlerFunc {
	return func(context *Context) {
		context.Next()
		log.Infof("[%d] - %s", context.StatusCode, context.Req.RequestURI)
	}
}
