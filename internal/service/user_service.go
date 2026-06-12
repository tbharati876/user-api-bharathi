package service

import (
	"context"

	"user-api/db/sqlc"
	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return s.repo.CreateUser(ctx, arg)
}

func (s *UserService) GetUser(ctx context.Context, id int32) (sqlc.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	return s.repo.ListUsers(ctx, arg)
}

func (s *UserService) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	return s.repo.UpdateUser(ctx, arg)
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}