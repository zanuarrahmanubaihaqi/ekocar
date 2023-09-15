package middleware

import (
	"eko-car/domain/shared/constant"
	"eko-car/domain/shared/context"
	Error "eko-car/domain/shared/error"
	"eko-car/domain/shared/response"
	"eko-car/infrastructure/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthValidations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.CreateContext()
		ctx = context.SetValueToContext(ctx, c)

		authToken := c.Get(constant.AUTHORIZATION)
		tokenString := strings.ReplaceAll(authToken, constant.BEARER, "")
		_, err := jwt.JWTChecking(tokenString)
		if err != nil {
			err = Error.New(constant.ErrAuth, constant.ErrAuth, err)
			return response.ResponseErrorWithContext(ctx, err)
		}

		c.Next()
		return nil
	}
}
