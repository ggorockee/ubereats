package common

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

func DecodeStructure(from any, to any) error {

	config := mapstructure.DecoderConfig{
		Result:           to,
		WeaklyTypedInput: true,
		TagName:          "json",
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}

	if err := decoder.Decode(from); err != nil {
		return fmt.Errorf("failed to decode structure: %w", err)
	}

	return nil
}

func RequestParserAndValidate(c *fiber.Ctx, requestBody any) error {

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CoreResponse{
			Message: err.Error(),
		})
	}
	validate := validator.New()
	// userDto.RegisterCustomValidations(validate)

	if err := validate.Struct(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CoreResponse{
			Message: err.Error(),
		})
	}

	return nil
}

func ValidateStruct(input any, ctx ...*fiber.Ctx) error {
	options := struct{ context *fiber.Ctx }{}

	switch {
	case len(ctx) > 0:
		options.context = ctx[0]
	}

	validate := validator.New()

	// validate.RegisterValidation("follow_status", func(fl validator.FieldLevel) bool {
	// 	// fl.Field().Interface().(string)
	// 	return fl.Field().String() == "pending" ||
	// 		fl.Field().String() == "approved" ||
	// 		fl.Field().String() == "blocked"
	// })

	validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "client" ||
			fl.Field().String() == "owner" ||
			fl.Field().String() == "delivery"
	})

	if c := options.context; c != nil {
		if err := validate.Struct(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				CoreResponse{
					Message: err.Error(),
				},
			)
		}

		return c.Status(fiber.StatusOK).JSON(
			CoreResponse{
				Ok: true,
			},
		)
	}

	// context가 없음
	return validate.Struct(input)

}
