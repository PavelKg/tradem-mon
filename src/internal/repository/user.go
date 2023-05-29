package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
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

func (u user) Create(userData model.UserDto) error {
	type newUser struct {
		model.UserDto
	}
	res := u.db.Model(&newUser{}).Create(map[string]interface{}{
		"name":       userData.Name,
		"username":   userData.UserName,
		"email":      userData.Email,
		"password":   clause.Expr{SQL: "crypt(?, gen_salt('bf'))", Vars: []interface{}{userData.Password}},
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

func (u user) Update(email string, userData model.User) error {
	res := u.db.Model(model.User{}).
		Where(&model.User{Email: email}).
		Omit("email", "password").
		Updates(&userData)
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for update was't found", email))
		return fmt.Errorf("ErrUserNotFound")
	}

	return nil
}

func (u user) Delete(email string) error {
	res := u.db.Delete(&model.User{}, &model.User{Email: email})
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}
	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for delete was't found", email))
		return fmt.Errorf("ErrUserNotFound")
	}
	return nil
}

func (u user) SetPassword(email string, password string) error {
	if len(password) < 8 {
		return fmt.Errorf("ErrWrongPasswordFormat")
	}

	res := u.db.Exec("UPDATE public.users SET password = crypt(?, gen_salt('bf')) WHERE email = ?", password, email)

	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		log.Println(fmt.Sprintf("User %s for set password was't found", email))
	}
	return nil
}

// GetByID returns user props by user id (login)
func (u user) GetByID(email string) (model.User, error) {
	var user model.User
	res := u.db.Where(&model.User{Email: email}).Find(&user)
	if res.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

// NewUserRepository returns new UserRepository instance
func NewUserRepository(config Config) model.UserRepository {
	return &user{db: config.DB}
}
