package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) RegisterUser(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "RegisterUser")
	defer span.End()

	var request model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][user_handler][RegisterUser] unable to decode request")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userData, err := h.restoUsecase.RegisterUser(ctx, request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][user_handler][RegisterUser] unable to register user")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userData,
	})
}

func (h *handler) Login(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Login")
	defer span.End()

	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][user_handler][Login] unable to decode request")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
	}

	sessionData, err := h.restoUsecase.Login(ctx, request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][user_handler][Login] unable to login user")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": sessionData,
	})
}
