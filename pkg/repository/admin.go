package repository

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}

}

func (ad *adminRepository) LoginHandler(adminDtails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where username = ? ", adminDtails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return adminCompareDetails, nil
}

func (ad *adminRepository) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 5
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", 20, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil
}

func (ad *adminRepository) GetUserByID(id string) (domain.Users, error) {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does  not exist")
	}
	Query := fmt.Sprintf("select * from users where id = %d", user_id)
	var userDetails domain.Users

	if err := ad.DB.Raw(Query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}
	return userDetails, nil
}

func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {
	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *adminRepository) CheckIfPaymentMethodAlreadyExists(payment string) (bool, error) {
	var count int64
	err := i.DB.Raw("SELECT COUNT(*) FROM payment_methods WHERE payment_name = ?", payment).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *adminRepository) NewPaymentMethod(pay string) error {

	if err := i.DB.Exec("INSERT INTO payment_methods(payment_name) VALUES (?)", pay).Error; err != nil {
		return err
	}
	return nil
}

func (i *adminRepository) ListPaymentMethods() ([]domain.PaymentMethod, error) {
	var model []domain.PaymentMethod
	err := i.DB.Raw("SELECT * FROM payment_methods WHERE is_deleted = false").Scan(&model).Error
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return model, nil
}

func (i *adminRepository) DeletePaymentMethod(id int) error {
	err := i.DB.Exec("UPDATE payment_methods SET is_deleted =true WHERE id = $1 ", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *adminRepository) GetPaymentMethod() ([]models.PaymentMethodResponse, error) {
	var model []models.PaymentMethodResponse
	err := i.DB.Raw("select * from payment_methods").Scan(&model).Error
	if err != nil {
		return []models.PaymentMethodResponse{}, err
	}
	return model, nil
}

func (i *adminRepository) DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := i.DB.Raw("select count(*) from users where is_admin= 'false'").Scan(&userDetails.TotalUser).Error
	if err != nil {
		return models.DashBoardUser{}, err
	}

	err = i.DB.Raw("select count(*) from users where blocked = 'true'").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, err
	}
	return userDetails, nil

}

func (i *adminRepository) DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct

	err := i.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, err
	}

	err = i.DB.Raw("select * from products where stock <= 0").Scan(&productDetails.OutOfStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, err
	}
	return productDetails, nil
}

func (i *adminRepository) DashBoardOrder() (models.DashboardOrder, error) {
	var dashboardOrder models.DashboardOrder
	err := i.DB.Raw("select count(*) from orders where payment_status = 'PAID'").Scan(&dashboardOrder.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}

	err = i.DB.Raw("select count(*) from orders where order_status = 'PENDING'").Scan(&dashboardOrder.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}

	err = i.DB.Raw("select count(*) from orders where order_status = 'CANCELED'").Scan(&dashboardOrder.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}

	err = i.DB.Raw("select count(*) from orders").Scan(&dashboardOrder.TotalOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}
	err = i.DB.Raw("select sum(quantity) from order_items").Scan(&dashboardOrder.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}

	return dashboardOrder, nil
}

func (i *adminRepository) TotalRevenue() (models.DashBoardRevenue, error) {
	var revenueDetails models.DashBoardRevenue

	startTime := time.Now().AddDate(0, 0, -1)

	err := i.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.ToadayRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}

	startTime = time.Now().AddDate(0, -1, 1).UTC()
	err = i.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}
	startTime = time.Now().AddDate(-1, 1, 1).UTC()
	err = i.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'  and created_at >= ?", startTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}

	return revenueDetails, nil

}

func (i *adminRepository) AmountDetails() (models.DashBoardAmount, error) {
	var amountDetails models.DashBoardAmount
	err := i.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'PAID'").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashBoardAmount{}, err
	}

	err = i.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status= 'NOT PAID' and order_status = 'PENDING' or order_status = 'SHIPPED'").Scan(&amountDetails.PendingAmounr).Error
	if err != nil {
		return models.DashBoardAmount{}, err
	}
	return amountDetails, nil
}
