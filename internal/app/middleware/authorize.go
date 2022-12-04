package middleware

import (
	"konntent-service-template/pkg/claimer"
	"konntent-service-template/pkg/middlewarepkg"
	"konntent-service-template/pkg/utils"
	"konntent-service-template/pkg/utils/userutil"

	"github.com/gofiber/fiber/v2"
)

func Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwt := c.Context().Value(utils.Claimer).(claimer.Claimer)
		jwtJson, valid := jwt.IsValid(c.UserContext(), middlewarepkg.GetAuthorizationHeader(c))

		if !valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		userutil.PutAuthModel(c, jwtJson)

		return c.Next()
	}
}
