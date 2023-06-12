package constant

type OrderStatus string

const (
	OrderStatusProcessed OrderStatus = "processed"
	OrderStatusFinished  OrderStatus = "finished"
	OrderStatusFailed    OrderStatus = "failed"
)

type ProductOrderStatus string

const (
	ProductOrderStatusPreparing ProductOrderStatus = "preparing"
	ProductOrderStatusFinished  ProductOrderStatus = "finished"
)
