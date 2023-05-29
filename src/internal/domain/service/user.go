package service

import (
	"fmt"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
)

type userService struct {
	userRepo model.UserRepository
}

// Get returns list of user
func (u userService) Get() ([]model.User, error) {
	users, err := u.userRepo.Get()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create send request to create a new user
func (u userService) Create(userData model.UserDto) error {
	return u.userRepo.Create(userData)
}

// Update sends request to update a user
func (u userService) Update(login string, userData model.User) error {
	return u.userRepo.Update(login, userData)
}

// Delete sends request to delete a user
func (u userService) Delete(login string) error {
	return u.userRepo.Delete(login)
}

// GetUserById returns user by ID
func (u userService) GetUserById(login string) (model.User, error) {
	var user model.User
	user, err := u.userRepo.GetByID(login)

	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

// GetUserPersonalProps returns user personal props
func (u userService) GetUserPersonalProps(login string) (model.MeProperties, error) {
	return model.MeProperties{}, nil
}

func NewUserService(userRepo model.UserRepository) model.UserService {
	return &userService{userRepo: userRepo}
}
