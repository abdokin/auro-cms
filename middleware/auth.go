package middleware

import (
	"auro-cms/handler"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)


func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the session using the store
		sess, err := session.Get(handler.AUTH_SESSION, c)
		fmt.Println("sessions" , sess.Name())
		if err != nil {
			// Handle the error, maybe return an error response or log it
			fmt.Println("Error getting session:", err)
			return err
		}
		// Get the request path to check if it's a login or register route
		path := c.Request().URL.Path

		// Skip authentication for login and register routes
		if path == "/login" || path == "/register" {
			fmt.Println("Skipping authentication for:", path)
			return next(c)
		}

		// Check if the session value exists and is authenticated
		authValue, ok := sess.Values[handler.AUTH_KEY]
		fmt.Println("auth values :",sess.Values);
		if !ok || authValue != true {
			// Redirect to login if not authenticated
			fmt.Println("User not authenticated. Redirecting to login.")
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		// If authenticated and not login or register, proceed to the next handler
		fmt.Println("User authenticated. Proceeding to next handler.")
		return next(c)
	}
}
