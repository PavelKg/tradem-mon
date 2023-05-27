package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	Login      string `json:"login"`
	Email      string `json:"email"`
	gorm.Model `json:"-"`
}

type UserAuthData struct {
	Sub  string `json:"username" validate:"required,min=3,max=32"`
	Pass string `json:"password" validate:"required,min=3,max=32"`
}

type UserContent struct {
	Sub  string
	Role string
}

// MeProperties is an app user personal props
type MeProperties struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	GetByID(login string) (User, error)
	Get() ([]User, error)
	Create(u User) error
	Update(login string, user User) error
	Delete(login string) error
}

// UserService represent the user's service contract
type UserService interface {
	GetUserById(login string) (User, error)
	GetUserPersonalProps(login string) (MeProperties, error)
	Get() ([]User, error)
	Create(u User) error
	Update(login string, user User) error
	Delete(login string) error
}

// UserPresenter represent the user's presenter contract
type UserPresenter interface {
	LoginUser(*fiber.Ctx) error
	GetUserPersonalProps(*fiber.Ctx) error
	Get(*fiber.Ctx) error
	GetById(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
}

// TableName set DB table name for model User
func (User) TableName() string {
	return "users"
}
