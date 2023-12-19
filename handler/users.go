package handler

import (
	"auro-cms/model"
	"auro-cms/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UsersPage(c echo.Context) error {
	user, _ := CurrentUser(c, h.DB)
	var users []model.User
	result := h.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}
	return pages.UsersPage(model.NewUserResponse(user), users).Render(c.Request().Context(), c.Response().Writer)
}
