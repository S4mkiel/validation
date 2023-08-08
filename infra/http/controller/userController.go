package controller

import (
	"github.com/S4mkiel/validation/domain/entity"
	service "github.com/S4mkiel/validation/domain/service/user"
	"github.com/S4mkiel/validation/infra/http/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserController struct {
	logger      *zap.SugaredLogger
	userService *service.UserService
}

func NewUserController(logger *zap.SugaredLogger, userService *service.UserService) *UserController {
	return &UserController{logger: logger, userService: userService}
}

func (c UserController) RegisterRoutes(app fiber.Router) {
	pa := app.Group("/user")
	pa.Post("/create", c.Create)
}

func (c UserController) Create(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	type Payload struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"required"`
	}

	var request Payload

	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	// validator request
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		if err.Error() == "Key: 'Payload.Email' Error:Field validation for 'Email' failed on the 'email' tag" {
			failureResponse.Error = ErrWithoutEmail.Error()
		} else if err.Error() == "Key: 'Payload.Name' Error:Field validation for 'Name' failed on the 'required' tag" {
			failureResponse.Error = ErrWithoutName.Error()
		} else {
			failureResponse.Error = "Invalid request"
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	u := entity.User{
		Email: request.Email,
		Name:  request.Name,
	}

	user, err := c.userService.Create(ctx.Context(), u)
	if err != nil {
		failureResponse.Error = ErrWithoutEmail.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = user

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
