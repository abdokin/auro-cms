package handler

import (
	"auro-cms/model"
	"auro-cms/utils"
	"auro-cms/views"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) (err error) {
	// Bind
	user_request := &model.RegisterRequest{}
	if err = c.Bind(user_request); err != nil {
		return
	}
	// // Validate
	errorMap := h.Validator.Validate(user_request)
	if errorMap != nil {
		fmt.Println("error map :", errorMap)
		return views.RegisterPage(errorMap, *user_request).Render(c.Request().Context(), c.Response().Writer)
	}
	// Save user
	user := model.NewUser(user_request)
	if results := h.DB.Create(&user); results.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: results.Error}
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func (h *Handler) Logout(c echo.Context) (err error) {
	sess, _ := session.Get(AUTH_SESSION, c)
	// delete session
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false, // TODO: make in env
	}
	sess.Values[AUTH_KEY] = false
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *Handler) Login(c echo.Context) (err error) {
	user_request := &model.LoginRequest{}
	if err = c.Bind(user_request); err != nil {
		return views.LoginPage("Invalid email or password", *user_request).Render(c.Request().Context(), c.Response().Writer)
	}
	// // Validate
	errorMap := h.Validator.Validate(user_request)
	if errorMap != nil {
		user_request.Password = ""
		return views.LoginPage("Invalid email or password", *user_request).Render(c.Request().Context(), c.Response().Writer)
	}

	// Find user
	var user = model.User{}
	if result := h.DB.First(&user, "email =?", user_request.Email); result.Error != nil {
		user_request.Password = ""
		return views.LoginPage("Invalid email or password", *user_request).Render(c.Request().Context(), c.Response().Writer)
	}
	fmt.Println("user loaded", user)
	if !utils.ComparePasswords(user.Password, []byte(user_request.Password)) {
		return views.LoginPage("Invalid email or password", *user_request).Render(c.Request().Context(), c.Response().Writer)
	}
	// Sessions
	sess, err := session.Get(AUTH_SESSION, c)
	if err != nil {

		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Something went wrong"}
	}
	// Set user as authenticated
	sess.Values[AUTH_KEY] = true
	sess.Values["id"] = user.ID
	sess.Values["email"] = user.Email
	sess.Save(c.Request(), c.Response())
	fmt.Println("auth values :", sess.Values)

	return c.Redirect(http.StatusMovedPermanently, "/dashboard/profile")
}

func (h *Handler) Profile(c echo.Context) (err error) {
	user_id := userIDFromSession(c)
	user := model.User{}
	fmt.Println("user with id ", user_id)
	if result := h.DB.First(&user, user_id); result.Error != nil {
		fmt.Println("user not found")

		return h.Logout(c)
	}
	return views.ProfilePage(model.NewUserResponse(&user)).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) RegisterPage(c echo.Context) error {
	return views.RegisterPage(nil, model.RegisterRequest{}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) LoginPage(c echo.Context) error {
	return views.LoginPage("", model.LoginRequest{}).Render(c.Request().Context(), c.Response().Writer)
}
func userIDFromSession(c echo.Context) uint {
	sess, _ := session.Get(AUTH_SESSION, c)
	if userID, ok := sess.Values["id"].(uint); ok {
		return userID
	}
	return 0
}
