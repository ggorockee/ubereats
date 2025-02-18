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
		return c.Status(fiber.StatusBadRequest).JSON(JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}
	validate := validator.New()

	if err := validate.Struct(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return nil
}

func ErrorResponse(c *fiber.Ctx, arg ErrArg) error {
	return c.Status(arg.Code).JSON(JsonResponse{
		Error:   arg.IsError,
		Message: arg.Message,
		Data:    arg.Data,
	})
}

func SuccessResponse(c *fiber.Ctx, arg SuccessArg) error {
	return c.Status(fiber.StatusOK).JSON(JsonResponse{
		Error:   false,
		Message: arg.Message,
		Data:    arg.Data,
	})
}

type ErrArg struct {
	IsError bool
	Code    int
	Message string
	Data    any
}

type SuccessArg struct {
	Message string
	Data    any
}

type JsonResponse struct {
	Error   bool
	Message string
	Data    any
}
