package interfaces

import "GlassGalore/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentRazorPay(userIDint, orderIDint int) (models.CombinedOrderDetails, string, error)
	SavePaymentDetails(paymentID, razorID, orderID string) error
}
