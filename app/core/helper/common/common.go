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
