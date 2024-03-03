package repository

import (
	"GlassGalore/pkg/repository/interfaces"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &PaymentRepositoryImpl{
		DB: DB,
	}
}

//add payment details

func (i *PaymentRepositoryImpl) AddRazorPayDetails(orderId int, razorpayID string) error {
	query := `insert into payments (order_id,razer_id) values($1,$2)`
	if err := i.DB.Exec(query, orderId, razorpayID).Error; err != nil {
		err := errors.New("error in inserting values to razor pay table" + err.Error())
		return err
	}
	return nil
}

func (i *PaymentRepositoryImpl) GetPaymentStatus(orderID string) (bool, error) {
	var paymentStatus string
	err := i.DB.Raw("select payment_status from orders where id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return false, err
	}

	//check if payment status is paid
	isPaid := paymentStatus == "PAID"
	fmt.Println("is payment status is PAID", isPaid)
	return isPaid, nil
}

func (i *PaymentRepositoryImpl) UpdatePaymentDetails(orderID, paymentID string) error {
	if err := i.DB.Exec("update payments set payment = $1 where razer_id = $2", paymentID, orderID).Error; err != nil {
		err = errors.New("error in updating the razerpay table" + err.Error())
		return err
	}
	return nil
}

func (i *PaymentRepositoryImpl) UpdatePaymentStatus(status bool, orderId string) error {
	var paymentStatus string

	if status {
		paymentStatus = "PAID"
	} else {
		paymentStatus = "NOT PAID"
	}
	if err := i.DB.Exec("update orders set payment_status = $1, order_status = 'SHIPPED' where id = $2", paymentStatus, orderId).Error; err != nil {
		err = errors.New("error in update order payment status" + err.Error())
		return err
	}
	return nil
}
