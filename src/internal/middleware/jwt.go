package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"

	"github.com/pavelkg/tradem-mon-api/pkg/utils"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt

func JWTProtected(key string) func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	//pem := "" // p.GetPubKey()
	//rsa, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))

	//if err != nil {
	//	fmt.Println(err)
	//}

	config := jwtMiddleware.Config{
		SigningMethod:  "HS256",
		SigningKey:     []byte(key),
		ContextKey:     utils.JwtContextKey, // used in private routes
		ErrorHandler:   func(c *fiber.Ctx, err error) error { return jwtError(c, err) },
		SuccessHandler: mixJwtToCtx,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	log.Println("{JWT} %s", err.Error())
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
}

func mixJwtToCtx(c *fiber.Ctx) error {
	return c.Next()
}
