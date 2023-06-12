package resto

import (
	"context"
	"errors"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/menu"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/order"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/user"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (m *restoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	res, err := m.menuRepo.GetMenuList(ctx, menuType)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][GetMenuList] unable to get menu list")

		return nil, err
	}

	return res, err
}

func (m *restoUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := m.menuRepo.GetMenu(ctx, orderProduct.OrderCode)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":        err,
				"order_code": orderProduct.OrderCode,
			}).Error("[usecase][resto][Order] unable to get menu")

			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  orderProduct.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrderData, err := m.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][Order] unable to create order data")

		return model.Order{}, err
	}

	return createdOrderData, nil
}

func (m *restoUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()

	orderData, err := m.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][GetOrderInfo] unable to get order info")

		return orderData, err
	}

	if orderData.UserID != request.UserID {
		logrus.WithFields(logrus.Fields{
			"request_user_id": request.UserID,
		}).Warn("[usecase][resto][GetOrderInfo] userID mismatch, not authorized")

		return model.Order{}, errors.New("unauthorized")
	}

	return orderData, nil
}

func (m *restoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	userRegistered, err := m.userRepo.CheckRegistered(ctx, request.Username)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][RegisterUser] unable to check registered")

		return model.User{}, err
	}
	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}

	userHash, err := m.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][RegisterUser] unable to generate user hash")

		return model.User{}, err
	}

	userData, err := m.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][RegisterUser] unable to register user")

		return model.User{}, err
	}

	return userData, nil
}

func (m *restoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()

	userData, err := m.userRepo.GetUserData(ctx, request.Username)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][Login] unable to get user data")

		return model.UserSession{}, err
	}

	verified, err := m.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][Login] unable to verify login")

		return model.UserSession{}, err
	}

	if !verified {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("[usecase][resto][Login] unverified login attempt")

		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := m.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][Login] unable to create user session")

		return model.UserSession{}, err
	}

	return userSession, nil
}

func (m *restoUsecase) CheckSession(ctx context.Context, sessionData model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	userID, err = m.userRepo.CheckSession(ctx, sessionData)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[usecase][resto][CheckSession] unable to check session")

		return "", err
	}

	return userID, err
}
