package main

import (
	"auro-cms/handler"
	"auro-cms/model"
	"auro-cms/views"
	"net/http"
	"io"
	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Template struct {
	component templ.Component
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
	  // Optionally, you could return the error to give each route more control over the status code
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
  }

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.component.Render(c.Request().Context(), w)
}
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/register" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))
	// custome Validator 
	e.Validator = &CustomValidator{validator: validator.New()}


	// Database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// // Initialize handler
	h := &handler.Handler{DB: db}

	// // Routes
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/", func(c echo.Context) error {
		return views.WelcomePage().Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/login", func(c echo.Context) error {
		return views.LoginPage().Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/register", func(c echo.Context) error {
		return views.RegisterPage().Render(c.Request().Context(), c.Response().Writer)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
