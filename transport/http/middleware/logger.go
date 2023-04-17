package middleware

import (
	"app/logger"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rid := c.Request().Header.Get(echo.HeaderXRequestID)

		ctx := logger.WithRqId(c.Request().Context(), rid)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
