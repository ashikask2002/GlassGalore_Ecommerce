package models

type PaymentMethodResponse struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}

type OrderDetails struct {
	OrderID         int     `json:"order_id" gorm:"column:order_id"`
	AddressID       int     `json:"address_id" gorm:"column:address_id"`
	PaymentMethodID int     `json:"payment_method_id" gorm:"column:payment_method_id"`
	Price           float64 `json:"price" gorm:"column:price"`
	OrderStatus     string  `json:"order_status" gorm:"column:order_status"`
	PaymentStatus   string  `json:"payment_status" gorm:"column:payment_status"`
}
type CombinedOrderDetails struct {
	OrderID       string  `json:"order_id"`
	FinalPrice    float64 `json:"final_price"`
	OrderStatus   string  `json:"order_status" gorm:"column:order_status"`
	PaymentStatus string  `json:"payment_status"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	Street        string  `json:"street" validate:"required"`
	City          string  `json:"city" validate:"required"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
}

type OrderPay struct {
	UserID uint `json:"user_id" `

	AddressID uint `json:"address_id" `

	PaymentMethodID uint    `json:"paymentmethod_id"`
	FinalPrice      float64 `json:"price"`
	OrderStatus     string  `json:"order_status" `
	PaymentStatus   string  `json:"payment_status"`
}

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type AllItems struct {
	OrderPay  OrderPay
	OrderItem []OrderItem
}
type OrderPayOnly struct {
	OrderID    uint    `json:"order_id" `
	FinalPrice float64 `json:"final_price"`
}

type OrderProducts struct {
	ProductIs string `json:"id"`
	Stock     int    `json:"stock"`
}
