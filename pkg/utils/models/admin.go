package models

type AdminLogin struct {
	Email    string `json:"email,omitempty"  validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `josn:"name"`
	Email string `json:"email"`
}
type NewPaymentMethod struct {
	PaymenMethod string `json:"payment_method"`
}
type Coupons struct {
	Coupon       string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate" gorm:"not null"`
	Valid        bool   `json:"valid" gorm:"default:true"`
}
type CompleteAdminDashboard struct {
	DashboardUser    DashBoardUser
	DashboardProduct DashBoardProduct
	DashboardOrder   DashboardOrder
	DashboardRevenue DashBoardRevenue
	DashboardAmount  DashBoardAmount
}

type DashBoardUser struct {
	TotalUser   int `json:"totaluser"`
	BlockedUser int `json:"blockuser"`
}
type DashBoardProduct struct {
	TotalProducts     int `json:"totalproduct"`
	OutOfStockProduct int `json:"outofstock"`
}
type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}

type DashBoardRevenue struct {
	ToadayRevenue float64
	MonthRevenue  float64
	YearRevenue   float64
}

type DashBoardAmount struct {
	CreditedAmount float64
	PendingAmounr  float64
}

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	ReturnedOrders  int
	CancelledOrders int
	TrendingProduct string
}
