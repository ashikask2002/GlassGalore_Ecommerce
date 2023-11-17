package usecase

import (
	"GlassGalore/pkg/domain"
	helper_interface "GlassGalore/pkg/helper/interfaces"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"fmt"

	"errors"
)

type UserUseCase struct {
	userRepo interfaces.UserRepository
	helper   helper_interface.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, helper helper_interface.Helper) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
		helper:   helper,
	}
}

var IntenalError = "Internal server Error"
var ErrorHashingPassword = "Error In Hashing Password"

func (u *UserUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")

	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	hashedPassword, err := u.helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New(ErrorHashingPassword)
	}
	user.Password = hashedPassword

	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not add the user")
	}

	tokenString, err := u.helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users: userData,
		Token: tokenString,
	}, nil

}

func (u *UserUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")

	}
	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New(IntenalError)
	}
	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case (The same function can be reused in future)
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, errors.New(IntenalError)
	}

	err = u.helper.CompareHashAndPassword(user_details.Password, user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = int(user_details.Id)
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")

	}
	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (i *UserUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	details, err := i.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}
	fmt.Println("qqqqqqqqqqqq", details)
	return details, nil
}

func (i *UserUseCase) GetAddresses(id int) ([]domain.Address, error) {
	addresses, err := i.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil
}

func (i *UserUseCase) AddAddress(id int, address models.AddAddress) error {
	rslt := i.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := i.userRepo.AddAddress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}
	return nil

}
