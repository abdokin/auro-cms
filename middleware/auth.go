package middleware

import (
	"auro-cms/handler"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		// Skip authentication for login and register routes
		if path == "/login" || path == "/register" {
			fmt.Println("Skipping authentication for:", path)
			return next(c)
		}
		// Check if the session value exists and is authenticated
		if !handler.IsLoggedIn(c) {
			return c.Redirect(http.StatusNotFound, "/login")
		}
		fmt.Println("User authenticated. Proceeding to next handler.")
		return next(c)
	}
}
