package handler

import (
	"auro-cms/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) (err error) {
	// Bind
	user_request := &model.UserRequest{}
	if err = c.Bind(user_request); err != nil {
		return
	}

	// // Validate
	// if err = c.Validate(user_request); err != nil {
	// 	return
	// }

	// Save user
	user := model.User{
		Email:    user_request.Email,
		Password: user_request.Password,
	}
	if results:= h.DB.Create(&user); results.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: results.Error}
	}
	user.Password = ""

	return c.JSON(http.StatusCreated, user_request)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	user_request := &model.UserRequest{}

	if err = c.Bind(user_request); err != nil {
		return
	}

	// Find user
	var user = model.User{Email: user_request.Email}
	if result := h.DB.First(&user); result.Error != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	user.Token, err = token.SignedString([]byte(Key))

	if result:= h.DB.Save(&user); result.Error != nil {
		// return http error 
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: result.Error}
	}
	if err != nil {
		return err
	}

	user.Password = "" // Don't send password
	return c.JSON(http.StatusOK, user)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
