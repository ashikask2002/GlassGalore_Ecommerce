package domain

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}

type Inventories struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  int      `json:"category_id"`
	Category    Category `josn:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size" gorm:"size:3;default:'M';check:size IN ('S', 'M', 'L')"`
	Stock       int      `json:"stock"`
	Price       int      `json:"price"`
}
