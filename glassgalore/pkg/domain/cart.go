package domain

type Cart struct {
	ID     uint  `json:"id" gorm:"primarykey"`
	UserID uint  `json:"user_id" gorm:"not null"`
	Users  Users `json:"-"  gorm:"foreignkey:UserID"`
}

type LineItems struct {
	ID        uint     `json:"id" gorm:"primarykey"`
	CartID    uint     `json:"cart_id" gorm:"not null"`
	Cart      Cart     `json:"-" gorm:"foreignkey:CartID"`
	ProductID uint     `json:"product_id" gorm:"not null"`
	Products  Products `json:"-" gorm:"foreignkey:ProductID;constraint:OnDelete:CASCADE"`
	Quantity  int      `json:"quantity" gorm:"default:1"`
}
