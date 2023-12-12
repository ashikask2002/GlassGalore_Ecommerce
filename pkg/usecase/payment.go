package usecase

import (
	"GlassGalore/pkg/repository/interfaces"
	usecasee "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/razorpay/razorpay-go"
)

type paymentUseCaseImpl struct {
	orderRepository   interfaces.OrderRepository
	paymentRepository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.OrderRepository, payment interfaces.PaymentRepository) usecasee.PaymentUseCase {
	return &paymentUseCaseImpl{
		orderRepository:   repo,
		paymentRepository: payment,
	}

}

//make payment through razorpay

func (i *paymentUseCaseImpl) MakePaymentRazorPay(userID int, orderID int) (models.CombinedOrderDetails, string, error) {
	fmt.Println("dddddd", orderID)
	order, err := i.orderRepository.GetOrdersRazor(orderID)
	if err != nil {
		err = errors.New("error in getting order details  through orderID" + err.Error())
		return models.CombinedOrderDetails{}, "", err
	}

	client := razorpay.NewClient("rzp_test_5K9ErTOEvk0TLA", "OmjlL1D5Px5CqonGNzyzhFgM")
	fmt.Println("client.................", client)

	data := map[string]interface{}{
		"amount":   int(order.FinalPrice) * 100,
		"currency": "INR",
		// "recipt":   "some_receipt_id",
	}
	fmt.Println("data................", data)
	body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Println("returned")
		return models.CombinedOrderDetails{}, "", nil
	}
	razorPayOrderId := body["id"].(string)
	fmt.Println("qwertyuiopp", razorPayOrderId)

	err = i.paymentRepository.AddRazorPayDetails(orderID, razorPayOrderId)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	body2, err := i.orderRepository.GetDetailedOrderThroughId(int(order.OrderID))
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	fmt.Println("orderid", body2.OrderID)
	return body2, razorPayOrderId, nil
}

//verify payment through razorpay

func (i *paymentUseCaseImpl) SavePaymentDetails(paymenID, razorID, orderID string) error {

	status, err := i.paymentRepository.GetPaymentStatus(orderID)
	if err != nil {
		return err
	}

	fmt.Println("status", status)

	if !status {
		err = i.paymentRepository.UpdatePaymentDetails(razorID, paymenID)
		if err != nil {
			return err
		}

		err = i.paymentRepository.UpdatePaymentStatus(true, orderID)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("already paid")
}
