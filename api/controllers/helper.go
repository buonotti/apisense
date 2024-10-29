package controllers

import (
	"github.com/buonotti/apisense/api/validation"
	"github.com/gofiber/fiber/v2"
)

// requestValid ensures the request can be parsed into the given type then validates the struct and fills in defaults
func requestValid(ctx *fiber.Ctx, request any) error {
	if err := ctx.BodyParser(request); err != nil {
		return err
	}
	err := validation.FillDefaults(request)
	if err != nil {
		return err
	}
	if err := validation.Validate(request); err != nil {
		return err
	}
	return nil
}
