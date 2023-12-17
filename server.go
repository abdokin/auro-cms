package main

import (
	"auro-cms/config"
	"auro-cms/model"
	"auro-cms/views"

	"github.com/labstack/echo/v4"
)

func main() {
	e := config.Init_server()
	// // Initialize handler
	h := config.Init_hanlders()
	// // Routes
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/profile", h.Profile)
	e.POST("/profile", h.Profile)

	e.GET("/", func(c echo.Context) error {
		return views.WelcomePage().Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/login", func(c echo.Context) error {
		return views.LoginPage("", model.LoginRequest{}).Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/register", func(c echo.Context) error {
		return views.RegisterPage(nil, model.RegisterRequest{}).Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/logout", h.Logout)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
