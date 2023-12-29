package models

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	CouponID        int `json:"coupon_id"`
}

type AddProducts struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Discription string  `json:"discription"`
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
	Name        string  `json:"name"`
	Discription string  `json:"discription"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	Size        string  `json:"size"`
}

type Products struct {
	ID          uint   `json:"id"`
	CategoryID  int    `json:"category_id"`
	ProductName string `json:"product_name"`
	Size        string `json:"size"`
	// Stock       int     `json:"stock"`
	Price         float64  `json:"price"`
	DiscountPrice float64  `json:"discount_price"`
	Rating        float64  `josn:"rating"`
	Image         string `json:"image"`
}

type AddToCart struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `josnL:"quantity"`
}

type SearchItems struct {
	ProductName string `json:"product_name"`
}
type ProductUserResponse struct {
	ID         uint `json:"id"`
	CategoryID int  `json:"category_id"`
	//Category    string `json:"category" gorm:"unique;not null"`
	ProductName string `json:"productname"`

	Size       string  `json:"size"`
	Price      float64 `json:"price"`
	OfferPrice float64 `json:"offerprice"`
}

type Search struct {
	Search string `json:"search"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}
type ItemDetails struct {
	ProductName string  `json:"product_name"`
	FinalPrice  float64 `json:"final_price"`
	Price       float64 `json:"price" `
	Total       float64 `json:"total_price"`
	Quantity    int     `json:"quantity"`
}
