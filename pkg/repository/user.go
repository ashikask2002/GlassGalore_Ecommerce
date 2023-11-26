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

func (i *userDatabase) EditDetails(id int, user models.EditDetailsResponse) (models.EditDetailsResponse, error) {

	var body models.EditDetailsResponse

	args := []interface{}{}
	query := "update users set"

	if user.Email != "" {
		query += " email = $1,"

		args = append(args, user.Email)
	}

	if user.Name != "" {
		query += " name = $2,"
		args = append(args, user.Name)
	}

	if user.Phone != "" {
		query += " phone = $3,"

		args = append(args, user.Phone)
	}

	query = query[:len(query)-1] + " where id = $4"

	args = append(args, id)
	// fmt.Println(query, args)
	err := i.DB.Exec(query, args...).Error
	if err != nil {
		return models.EditDetailsResponse{}, err
	}
	query2 := "select * from users where id = ?"
	if err := i.DB.Raw(query2, id).Scan(&body).Error; err != nil {
		return models.EditDetailsResponse{}, err
	}

	return body, nil

}

// func (c *userDatabase) EditEmail(id int, email string) error {

// 	err := c.DB.Exec("update users set email =$1 where id = $2", email, id).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *userDatabase) EditPhone(id int, phone string) error {
// 	err := c.DB.Exec("update users set phone = $1 where id = $2", phone, id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c *userDatabase) ChangePassword(id int, password string) error {
	err := c.DB.Exec("UPDATE users SET password = $1 WHERE id = $2", password, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *userDatabase) GetPassword(id int) (string, error) {
	var userPassword string
	err := c.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil
}

func (i *userDatabase) GetCartID(id int) (int, error) {
	var cart_id int

	if err := i.DB.Raw("select id from carts where user_id = ?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}
	return cart_id, nil
}

func (i *userDatabase) GetProductsInCart(cart_id int) ([]int, error) {
	var cart_products []int
	if err := i.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}
	return cart_products, nil
}

func (i *userDatabase) FindProductNames(inventory_id int) (string, error) {
	var product_name string

	if err := i.DB.Raw("select product_name from inventories where id= ?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}
	return product_name, nil
}

func (i *userDatabase) FindCartQuantity(cart_id, inventory_id int) (int, error) {
	var quantity int

	if err := i.DB.Raw("select quantity from line_items where cart_id = $1 and inventory_id = $2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}
	return quantity, nil
}

func (i *userDatabase) FindPrice(inventory_id int) (float64, error) {
	var price float64

	if err := i.DB.Raw("select price from inventories where id = ?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}

func (i *userDatabase) FindCategory(inventory_id int) (int, error) {
	var category int

	if err := i.DB.Raw("select category_id from inventories where id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}

	return category, nil
}

func (i *userDatabase) FindStock(id int) (int, error) {
	var stock int
	if err := i.DB.Raw("select stock from inventories where id = ?", id).Scan(&stock).Error; err != nil {
		return 0, err
	}
	return stock, nil
}

func (i *userDatabase) RemoveFromCart(cart, inventory int) error {
	if err := i.DB.Exec(`DELETE FROM line_items WHERE cart_id = $1 AND inventory_id = $2`, cart, inventory).Error; err != nil {
		return err
	}
	return nil
}
func (i *userDatabase) UpdateQuantity(id, inv_id, qty int) error {

	if id <= 0 || inv_id <= 0 || qty <= 0 {
		return errors.New("negtive or zero values are not allowed")
	}
	if qty >= 0 {
		query := `update line_items set quantity = $1 where cart_id = $2 and inventory_id= $3`

		result := i.DB.Exec(query, qty, id, inv_id)
		{
			if result.Error != nil {
				return result.Error
			}

		}
	}
	return nil

}
