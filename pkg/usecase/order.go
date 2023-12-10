package usecase

import (
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"strconv"
)

type orderUseCase struct {
	orderRepository  interfaces.OrderRepository
	userUseCase      services.UserUseCase
	walletRepository interfaces.WalletRepository
	couponRepository interfaces.CouponRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase, walletRepository interfaces.WalletRepository, couponRepository interfaces.CouponRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository:  repo,
		userUseCase:      userUseCase,
		walletRepository: walletRepository,
		couponRepository: couponRepository,
	}
}

func (i *orderUseCase) OrderItemsFromCart(userID int, addressID int, paymentID int, couponID int) error {
	// Retrieve the user's cart
	cart, err := i.userUseCase.GetCart(userID)
	if err != nil {
		return err
	}

	// Calculate the total price of the items in the cart
	var total float64
	for _, item := range cart.Data {
		if item.Quantity > 0 && item.Price > 0 {
			total += float64(item.Quantity) * float64(item.Price)
		}
	}

	couponvalid, err := i.couponRepository.CheckCouponValid(couponID)
	if err != nil {
		return err
	}
	if !couponvalid {
		return errors.New("this coupon is invalid")
	}

	coupon, err := i.couponRepository.FindCouponPrice(couponID)
	if err != nil {
		return err
	}

	totaldiscount := float64(coupon)

	total = total - totaldiscount

	// Place an order with the order repository
	orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, total)
	if err != nil {
		return err
	}

	// Add individual items from the cart to the order
	if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
		return err
	}
	//decrease the stock of product after order
	for _, v := range cart.Data {
		if err := i.orderRepository.ReduceStockAfterOrder(v.ProductName, v.Quantity); err != nil {
			return err
		}
	}

	// Remove the ordered items from the user's cart
	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}

	return nil
}

func (i *orderUseCase) GetOrders(orderid int) (models.AllItems, error) {

	orders, err := i.orderRepository.GetOrders(orderid)
	if err != nil {
		return models.AllItems{}, err
	}
	return orders, nil
}

func (i *orderUseCase) GerAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {
	allOrder, err := i.orderRepository.GetAllOrders(userId, page, pageSize)
	if err != nil {
		return []models.OrderDetails{}, err
	}
	return allOrder, nil
}

func (i *orderUseCase) CancelOrder(orderID int) error {
	// Check the order status
	orderStatus, err := i.orderRepository.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}
	if orderStatus == "RETURNED" {
		return errors.New("the order already returned")
	}

	if orderStatus == "CANCELED" {
		return errors.New("the order already canelled")
	}

	// orderProductDetails, err := i.orderRepository.GetProductDetailsFromOrder(orderID)
	// if err != nil{
	// 	return err
	// }

	// Check if the order can be canceled
	if orderStatus == "PENDING" || orderStatus == "SHIPPED" {

		paymentStatus, err := i.orderRepository.CheckPaymentStatusByID(orderID)
		if err != nil {
			return err
		}
		if paymentStatus != "PAID" {
			err = i.orderRepository.CancelOrder(orderID)
			if err != nil {
				return err
			}
		}
		if paymentStatus == "PAID" {
			UserID, err := i.orderRepository.FindUserID(orderID)
			if err != nil {
				return err
			}

			FinalPrice, err := i.orderRepository.FindFinalPrice(orderID)
			if err != nil {
				return err
			}

			//find if the user having the wallet

			walletID, err := i.walletRepository.FindWalletIDFromUserID(UserID)
			if err != nil {
				return err
			}

			//if not any wallet create a new wallet

			if walletID == 0 {
				walletID, err = i.walletRepository.CreateNewWallet(UserID)
				if err != nil {
					return err
				}
			}

			//add the price to the wallet
			if err := i.walletRepository.CreditToUserWallet(FinalPrice, walletID); err != nil {
				return err
			}
			// Attempt to cancel the order
			err = i.orderRepository.CancelOrderPaid(orderID)
			if err != nil {
				return err
			}

			// err = i.orderRepository.updateQuantityProduct(orderProductDetails)
			// if err != nil{
			// 	return err
			// }

		}

	}
	return nil

}

func (i *orderUseCase) GetAdminOrders(page int) ([]models.CombinedOrderDetails, error) {
	orderDetails, err := i.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil
}

func (i *orderUseCase) OrdersStatus(orderID string) error {
	orderId, sErr := strconv.Atoi(orderID)
	if sErr != nil {
		return sErr
	}
	status, err := i.orderRepository.CheckOrderStatusByID(orderId)
	if err != nil {
		return err
	}
	switch status {
	case "CANCELED", "RETURNED", "DELIVERED":
		return errors.New("cannot approve this order becouse it's in processed or cancelled state")
	case "PENDING":
		//for the admin approval change the PENDING to SHIPPED
		err := i.orderRepository.ChangeOrderStatus(orderID, "SHIPPED")
		if err != nil {
			return err
		}
	case "SHIPPED":
		shipmentStatus, err := i.orderRepository.GetShipmentStatus(orderID)
		if err != nil {
			return err
		}
		if shipmentStatus == "CANCELED" {
			return errors.New("cannot approve this orders becouse its cancelled")
		}

		//for admin approval, change SHIPPED to  DELIVERED
		err = i.orderRepository.ChangeOrderStatus(orderID, "DELIVERED")
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *orderUseCase) ReturnOrder(orderId int) error {
	shipmentStatus, err := i.orderRepository.GetOrderStatus(orderId)
	if err != nil {
		return err
	}
	if shipmentStatus == "RETURNED" {
		return errors.New("the order already returned")
	}

	if shipmentStatus != "DELIVERED" {
		return errors.New("user try to return the order that not even delivered")
	}

	if shipmentStatus == "DELIVERED" {

		UserID, err := i.orderRepository.FindUserID(orderId)
		if err != nil {
			return err
		}

		FinalPrice, err := i.orderRepository.FindFinalPrice(orderId)
		if err != nil {
			return err
		}

		//find if the user having the wallet

		walletID, err := i.walletRepository.FindWalletIDFromUserID(UserID)
		if err != nil {
			return err
		}

		//if not any wallet create a new wallet

		if walletID == 0 {
			walletID, err = i.walletRepository.CreateNewWallet(UserID)
			if err != nil {
				return err
			}
		}

		//add the price to the wallet
		if err := i.walletRepository.CreditToUserWallet(FinalPrice, walletID); err != nil {
			return err
		}

		if err := i.orderRepository.ReturnOrder("RETURNED", orderId); err != nil {
			return err
		}
		return nil
	}

	return errors.New("cannot return order")
}
