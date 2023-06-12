package rest

import (
	"net/http"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetMenu(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenu")
	defer span.End()

	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(ctx, menuType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][menu_handler][GetMenu] unable to get menu list")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
