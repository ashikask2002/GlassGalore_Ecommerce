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

type AddProducts struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type ProductResponse struct {
	ProductID int `json:"id"`
	Stock     int `json:"stock"`
}

type ProductUpdate struct {
	Productid int `json:"product_id"`
	Stock     int `json:"stock"`
}

type EditProductDetails struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
	Size       string  `json:"size"`
}

type Products struct {
	ID          uint   `json:"id"`
	CategoryID  int    `json:"category_id"`
	ProductName string `json:"product_name"`
	Size        string `json:"size"`
	// Stock       int     `json:"stock"`
	Price float64 `json:"price"`
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}
