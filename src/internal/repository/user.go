package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
	"github.com/pavelkg/tradem-mon-api/pkg/utils"
)

// user ...
type user struct {
	db *gorm.DB
}

func (u user) Get() ([]model.User, error) {
	users := make([]model.User, 0)
	res := u.db.Order("login asc").Find(&users)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return users, nil
}

func (u user) Create(userData model.User) error {
	type newUser struct {
		model.User
		Password string
	}
	res := u.db.Model(&newUser{}).Create(map[string]interface{}{
		"login":      userData.Login,
		"email":      userData.Email,
		"password":   clause.Expr{SQL: "crypt(?, gen_salt('bf'))", Vars: []interface{}{utils.GenerateRandomString(20)}},
		"created_at": clause.Expr{SQL: "now()"},
		"updated_at": clause.Expr{SQL: "now()"},
	})
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %v was't added", userData))
		return fmt.Errorf("ErrUserNotFound")
	}

	return nil
}

func (u user) Update(login string, userData model.User) error {
	res := u.db.Model(model.User{}).
		Where(&model.User{Login: login}).
		Omit("login", "password").
		Updates(&userData)
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for update was't found", login))
		return fmt.Errorf("ErrUserNotFound")
	}

	return nil
}

func (u user) Delete(login string) error {
	res := u.db.Delete(&model.User{}, &model.User{Login: login})
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for delete was't found", login))
		return fmt.Errorf("ErrUserNotFound")
	}
	return nil
}

func (u user) SetPassword(login string, password string) error {
	if len(password) < 8 {
		return fmt.Errorf("ErrWrongPasswordFormat")
	}

	res := u.db.Exec("UPDATE public.users SET password = crypt(?, gen_salt('bf')) WHERE login = ?", password, login)

	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for set password was't found", login))
	}
	return nil
}

// GetByID returns user props by user id (login)
func (u user) GetByID(login string) (model.User, error) {
	var user model.User
	res := u.db.Where(&model.User{Login: login}).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

// NewUserRepository returns new UserRepository instance
func NewUserRepository(config Config) model.UserRepository {
	return &user{db: config.DB}
}
