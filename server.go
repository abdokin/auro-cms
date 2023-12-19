package main

import (
	"auro-cms/config"
	"auro-cms/views/pages"
	"github.com/labstack/echo/v4"
)

func main() {
	e := config.Init_server()
	// // Initialize handler
	h := config.Init_hanlders()
	// // Routes
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/login", h.LoginPage)
	e.GET("/register", h.RegisterPage)
	e.GET("/logout", h.Logout)
	e.GET("/", func(c echo.Context) error {
		return pages.WelcomePage(false).Render(c.Request().Context(), c.Response().Writer)
	})
	// dashboard
	e.GET("/dashboard/profile", h.Profile)
	e.GET("/", h.WelcomePage)
	e.GET("/dashboard", h.DashboardPage)
	e.GET("/dashboard/users", h.UsersPage)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
