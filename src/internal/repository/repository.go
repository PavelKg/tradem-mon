package repository

import (
	"gorm.io/gorm"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
)

type Config struct {
	DB *gorm.DB
}

// Repositories ...
type Repositories struct {
	User model.UserRepository
}

// NewRepository ...
func NewRepository(config Config) (*Repositories, error) {
	user := NewUserRepository(config)
	return &Repositories{
		User: user,
	}, nil
}
