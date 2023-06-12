package model

import (
	"github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"
)

type MenuItem struct {
	OrderCode string            `gorm:"primaryKey" json:"order_code"`
	Name      string            `json:"name"`
	Price     int64             `json:"price"`
	Type      constant.MenuType `json:"type"`
}
