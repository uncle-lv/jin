package jin

import (
	"jin/log"
)

func DefaultLogger() HandlerFunc {
	return func(context *Context) {
		context.Next()
		defaultLoggerPrintf(context)
	}
}

func defaultLoggerPrintf(context *Context) {
	switch context.StatusCode {
	case 200, 201, 202, 203, 204, 205, 206:
		log.Infof("\033[32m[%d]\033[0m  %s", context.StatusCode, context.Req.RequestURI)

	case 300, 301, 302, 303, 304, 305, 307:
		log.Infof("\033[33m[%d]\033[0m  %s", context.StatusCode, context.Req.RequestURI)

	case 400, 401, 403, 404, 405, 406, 407, 408, 410, 411, 412, 413, 414, 415, 416, 417,
		500, 501, 502, 503, 504, 505:
		log.Infof("\033[31m[%d]\033[0m  %s", context.StatusCode, context.Req.RequestURI)
	}
}
