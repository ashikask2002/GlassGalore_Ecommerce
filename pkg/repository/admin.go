package repository

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"

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

//function which will block and unblock the user

func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {
	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}
	return nil
}