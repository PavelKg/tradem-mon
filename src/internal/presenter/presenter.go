package presenter

import (
	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
	"github.com/pavelkg/tradem-mon-api/internal/domain/service"
)

type Presenters struct {
	UserPresenter model.UserPresenter
}

func NewPresenters(handlers *service.Handlers) (*Presenters, error) {
	user := NewUserPresenter(handlers.Users)

	return &Presenters{

		UserPresenter: user,
	}, nil
}
