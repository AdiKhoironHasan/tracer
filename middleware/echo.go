package middleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func EchoMiddleware(serviceName string) echo.MiddlewareFunc {
	return otelecho.Middleware(serviceName)
}
