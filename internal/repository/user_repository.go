package repository

import (
	"context"

	"user-api/db/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *UserRepository) GetUser(ctx context.Context, id int32) (sqlc.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx, arg)
}

func (r *UserRepository) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
