package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
}
