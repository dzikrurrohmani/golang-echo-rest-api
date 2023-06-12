package main

import (
	"github.com/labstack/echo/v4"

	"github.com/dzikrurrohmani/golang-echo-rest-api/config"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/logger"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/menu"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/order"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/repository/user"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/tracing"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/usecase/resto"

	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/database"
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/delivery/rest"
)

func main() {
	config := config.NewConfig()

	logger.Init()
	tracing.Init("http://localhost:14268/api/traces")

	e := echo.New()

	db := database.GetDB(config.DataSourceName)

	menuRepo := menu.GetRepository(db)
	orderRepo := order.GetRepository(db)
	userRepo, err := user.GetRepository(
		db,
		config.Secret,
		config.Time,
		config.Memory,
		config.KeyLen,
		config.Parallelism,
		config.SignKey,
		config.AccessExp)
	if err != nil {
		panic(err)
	}

	restoUsecase := resto.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start((config.Url)))
}
