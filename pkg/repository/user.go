package repository

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"

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

func (c *userDatabase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	fmt.Println("hhhhhhhhhh", id)
	var details models.UserDetailsResponse

	if err := c.DB.Raw("select id,name,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserDetailsResponse{}, errors.New("could not get the user details")
	}
	fmt.Println("hhhhhhhhhh", details)

	return details, nil
}

func (c *userDatabase) GetAddresses(id int) ([]domain.Address, error) {
	var addresses []domain.Address
	if err := c.DB.Raw("select * from addresses where user_id = ?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, errors.New("could not get the address")
	}
	return addresses, nil
}

func (c *userDatabase) CheckIfFirstAddress(id int) bool {
	var count int

	if err := c.DB.Raw("select count(*) from addresses where user_id = ?", id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *userDatabase) AddAddress(id int, address models.AddAddress, result bool) error {
	err := c.DB.Exec(`INSERT INTO addresses(user_id, name, house_name, street, city, state, phone, pin,"default")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Phone, address.Pin, result).Error
	if err != nil {
		return errors.New("could not add address")
	}
	return nil
}
