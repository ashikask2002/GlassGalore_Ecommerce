package models

type CategorytOfferResp struct {
	CategoryID    uint   `json:"category_id" binding:"required"`
	OfferName     string `json:"offer_name" binding:"required"`
	DiscountPrice int    `json:"discount_price" binding:"required"`
}
