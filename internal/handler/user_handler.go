package handler

import (
	"context"
	"strconv"
	"time"

	"user-api/db/sqlc"
	"user-api/internal/models"
	"user-api/internal/repository"
	"user-api/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(queries *sqlc.Queries) *UserHandler {
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)

	return &UserHandler{
		service:  svc,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.UserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "validation failed"})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid dob"})
	}

	user, err := h.service.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Name: req.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(models.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.service.GetUser(context.Background(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  service.CalculateAge(user.Dob),
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := h.service.ListUsers(
		context.Background(),
		sqlc.ListUsersParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	response := make([]models.UserResponse, 0)

	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(user.Dob),
		})
	}

	return c.JSON(response)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	var req models.UserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "validation failed"})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid dob"})
	}

	user, err := h.service.UpdateUser(
		context.Background(),
		sqlc.UpdateUserParams{
			ID:   int32(id),
			Name: req.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(models.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	err = h.service.DeleteUser(context.Background(), int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}
