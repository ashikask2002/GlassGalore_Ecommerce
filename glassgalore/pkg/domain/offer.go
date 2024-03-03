package domain

import "time"

type CategoryOffer struct {
	ID            uint      `json:"id" gorm:"unique; not null"`
	CategoryID    uint      `json:"category_id"`
	Category      Category  `json:"-" gorm:"foreignkey:CategoryID"`
	OfferName     string    `json:"offer_name"`
	DiscountPrice int       `json:"discount_price"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}
