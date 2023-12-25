package domain

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}

type Products struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  int      `json:"category_id"`
	Category    Category `josn:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Discription string   `json:"discription"`
	Size        string   `json:"size" gorm:"size:3;default:'M';check:size IN ('S', 'M', 'L')"`
	Stock       int      `json:"stock"`
	Price       int      `json:"price"`
}

type Rating struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	Rating     float64 `json:"rating"`
}
