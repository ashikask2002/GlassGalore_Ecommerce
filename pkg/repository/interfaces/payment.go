package interfaces

type PaymentRepository interface {
	AddRazorPayDetails(int, string) error
	GetPaymentStatus(orderId string) (bool, error)
	UpdatePaymentDetails(orderID, paymentID string) error
	UpdatePaymentStatus(status bool, orderID string) error 	
}
