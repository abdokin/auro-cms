package config

import (
	"auro-cms/handler"
	"auro-cms/middleware"
	"auro-cms/model"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init_db() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	return db
}

func Init_hanlders() *handler.Handler {
	valid := handler.CustomValidator{V: validator.New()}
	h := &handler.Handler{DB: init_db(), Validator: valid}
	return h
}

func Init_server() *echo.Echo {
	e := echo.New()
	// e.Logger.SetLevel(log.OFF)
	// e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.RequestID())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.WithAuth)

	// custome Validator
	// e.Validator = &CustomValidator{validator: validator.New()}
	return e

}
