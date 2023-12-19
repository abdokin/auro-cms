package handler

import (
	"auro-cms/model"
	"auro-cms/views/pages"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Profile(c echo.Context) (err error) {

	user, _ := CurrentUser(c, h.DB)
	return pages.ProfilePage(model.NewUserResponse(user)).Render(c.Request().Context(), c.Response().Writer)
}
