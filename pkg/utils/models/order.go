package models

type PaymentMethodResponse struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}
