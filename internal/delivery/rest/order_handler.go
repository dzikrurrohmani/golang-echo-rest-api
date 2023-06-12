package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][Order] decode request data failed")

		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUsecase.Order(ctx, request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][Order] order failed")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetOrderInfo")
	defer span.End()

	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderInfo(ctx, model.GetOrderInfoRequest{
		OrderID: orderID,
		UserID:  userID,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][GetOrderInfo] unable to get order info")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
