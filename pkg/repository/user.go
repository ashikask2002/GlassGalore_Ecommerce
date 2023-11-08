package repository

import (
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (c *userDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var UserDetails models.UserDetailsResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone) VALUES (?, ?, ?, ?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&UserDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return UserDetails, nil

}

func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int
	if err := c.DB.Raw("select count(*) from users where email= ?", email).Scan(&count).Error; err != nil {
		return false
	}
	//count greater than means user already exist
	return count > 0
}

func (cr *userDatabase) UserBlockStatus(email string) (bool, error) {
	var isBlocked bool
	err := cr.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}
	return isBlocked, nil
}

func (c *userDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {

	var user_details models.UserSignInResponse
	err := c.DB.Raw(`SELECT * FROM users where email = ? and blocked = false`, user.Email).Scan(&user_details).Error
	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}

	return user_details, nil
}
