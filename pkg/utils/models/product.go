package models

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	// CouponID        int `json:"coupon_id"`
}
