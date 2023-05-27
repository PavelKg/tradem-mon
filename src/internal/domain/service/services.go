package service

import (
	_ "github.com/gofiber/fiber/v2"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
	"github.com/pavelkg/tradem-mon-api/internal/repository"
)

// TODO Need modify MakeHeadersRowsObject input params for interface contract, not only for *sql.Rows,

type Handlers struct {
	Users model.UserService
}

func NewServices(repo *repository.Repositories) (*Handlers, error) {
	users := NewUserService(repo.User)

	handlers := &Handlers{
		Users: users,
	}

	return handlers, nil
}
