package model

import "github.com/dzikrurrohmani/golang-echo-rest-api/internal/model/constant"

// Creating a one to many DB approach
type Order struct {
	// just note, gorm uses ID field as primary key by default
	ID            string               `gorm:"primaryKey" json:"id"`
	UserID        string               `gorm:"index" json:"user_id"`
	Status        constant.OrderStatus `json:"status"`
	ProductOrders []ProductOrder       `json:"product_orders"`
	ReferenceID   string               `gorm:"unique" json:"reference_id"`
}

type ProductOrder struct {
	ID         string
	OrderID    string
	OrderCode  string
	Quantity   int
	TotalPrice int64
	Status     constant.ProductOrderStatus
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	UserID        string                    `json:"-"`
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	OrderID string
	UserID  string
}
