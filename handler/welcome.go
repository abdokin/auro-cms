package handler

import (
	"auro-cms/model"
	"auro-cms/views/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) WelcomePage(c echo.Context) error {
	return pages.WelcomePage(IsLoggedIn(c)).Render(c.Request().Context(), c.Response().Writer)

}

func (h *Handler) DashboardPage(c echo.Context) error {
	user_id := userIDFromSession(c)
	user := model.User{}
	if result := h.DB.First(&user, user_id); result.Error != nil {
		return h.Logout(c)
	}
	return pages.DashboardPage(model.NewUserResponse(&user)).Render(c.Request().Context(), c.Response().Writer)
}
