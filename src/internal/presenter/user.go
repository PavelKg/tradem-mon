package presenter

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/pavelkg/tradem-mon-api/internal/domain/model"
	"github.com/pavelkg/tradem-mon-api/pkg/utils"
)

type userPresenter struct {
	userService model.UserService
}

// GetById func returns user by id.
// @Description returns user by id.
// @Summary returns user by id
// @Tags Users

// @Accept json
// @Produce json
// @Param id path string true "User ID (login)"
// @Success 200
// @Failed 404 if user not find
// @Failed 500
// @Router /api/user/:id [get]
func (pr *userPresenter) GetById(ctx *fiber.Ctx) error {
	login := ctx.Params("id", "")
	user, err := pr.userService.GetUserById(login)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

// Get Returns a list of users.
// @Description Returns a list of users.
// @Summary get a list of users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Failed 500
// @Router /api/users [get]
func (pr *userPresenter) Get(ctx *fiber.Ctx) error {

	users, err := pr.userService.Get()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// Create creates a new user
// @Description Creates a new user
// @Summary creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param username body string true "Username"
// @Success 201
// @Failed 400 If body is incorrect or there is service err
// @Router /api/users/ [post]
func (pr *userPresenter) Create(ctx *fiber.Ctx) error {
	var userData model.User

	if err := ctx.BodyParser(&userData); err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err := pr.userService.Create(userData)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Return status 201 OK.
	return ctx.SendStatus(fiber.StatusCreated)
}

// Update updates user data
// @Description Updates user data
// @Summary updates user data
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User login"
// @Param email body string true "Email"
// @Param first_name body string true "First name"
// @Param last_name body string true "Last name"
// @Param role_id body int true "roles id"
// @Success 204
// @Failure 404
// @Router /api/users/{id} [put]
func (pr *userPresenter) Update(ctx *fiber.Ctx) error {
	var userData model.User

	login := ctx.Params("id", "")
	if login == "" {
		log.Println("missed parameter user ID")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := ctx.BodyParser(&userData); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err := pr.userService.Update(login, userData)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Return status 204 OK.
	return ctx.SendStatus(fiber.StatusNoContent)
}

// Delete  deletes a user
// @Description Deletes a user
// @Summary deletes a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID (login)"
// @Success 200
// @Failure 404 If user login was not found
// @Failure 500
// @Router /api/users/{id} [delete]
func (pr *userPresenter) Delete(ctx *fiber.Ctx) error {
	login := ctx.Params("id", "")
	if login == "" {
		log.Println("missed parameter user ID")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	err := pr.userService.Delete(login)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Return status 200 OK.
	return ctx.SendStatus(fiber.StatusOK)
}

// GetUserPersonalProps Returns user's personal properties [personal data, menu, permissions]
// @Description Returns user's personal properties.
// @Summary returns user's personal properties
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Router /api/auth/me [get]
func (pr *userPresenter) GetUserPersonalProps(c *fiber.Ctx) error {
	var meProps model.MeProperties

	login, err := utils.GetUserIdFromJwt(c)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}

	user, err := pr.userService.GetUserById(login)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	meProps.Name = user.Name
	meProps.UserName = user.UserName
	meProps.Email = user.Email
	if err != nil {
		log.Print(err)
	}

	return c.Status(200).JSON(meProps)
}

// LoginUser Returns user's JWT or forbidden
// @Description Returns user's JWT or forbidden.
// @Summary returns user's JWT or forbidden
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200
// @Router /api/auth/login [post]
func (pr *userPresenter) LoginUser(c *fiber.Ctx) error {
	var person = &model.UserAuthData{}
	if err := c.BodyParser(person); err != nil || person.Sub == "" || person.Pass == "" {
		log.Println("user auth data is incorrect")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	signedToken := ""
	//if err != nil {
	//	return c.SendStatus(fiber.StatusForbidden)
	//}
	secret := "ilsdahJHILUuygaiosb2345" //pr.appService.GetAppJwtSecret()
	signedToken, err := utils.JwtGenSignedToken(person.Sub, secret)
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusForbidden)
	}
	log.Println("{APP} user %s logged in", person.Sub)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": signedToken})
}

func NewUserPresenter(
	userService model.UserService,
) model.UserPresenter {

	return &userPresenter{
		userService: userService,
	}
}
