package useradapter

import (
	"errors"
	"startup-backend/apps/user/userapp"
	"startup-backend/pkg/shared"

	"github.com/gofiber/fiber/v2"
)

func addNewUser(userapp *userapp.UserAppSvc) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		req := new(struct {
			Fullname string `json:"full_name"`
		})

		// c.FormValue("full_name")

		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(shared.NewRestErrorResponse(err))
		}

		if err := userapp.AddNewUser(req.Fullname); err != nil {
			if errors.Is(err, &shared.SystemError{}) {
				return c.Status(fiber.StatusInternalServerError).JSON(shared.NewRestErrorResponse(err))

			}

			return c.Status(fiber.StatusBadRequest).JSON(shared.NewRestErrorResponse(err))

		}

		return c.Status(fiber.StatusOK).JSON(shared.NewRestResponse("success bos", nil))

	}
}
