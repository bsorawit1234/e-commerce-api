package models

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	OrderItems  []OrderItem `json:"order_items"`
	TotalAmount uint        `json:"total_amount"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity_id"`
	Product   Product `json:"product"`
}
