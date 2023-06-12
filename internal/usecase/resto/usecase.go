package resto

import (
	"context"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockRestoUsecase -destination=../../mocks/resto_usecase_mock.go -source=usecase.go

type Usecase interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error)
	Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, sessionData model.UserSession) (userID string, err error)
}
