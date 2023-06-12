package rest

import (
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/usecase/resto"
)

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restoUsecase: restoUsecase,
	}
}
