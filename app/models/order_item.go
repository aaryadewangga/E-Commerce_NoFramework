package models

type OrderItem struct {
	ID      string `gorm:"size:36;not bull;uniqueIndex;primary_key"`
	Order   Order
	OrderID string `gorm:"size:36;index"`
}
