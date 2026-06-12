package sqlc

import (
	"context"
	"time"
)

type User struct {
	ID   int32
	Name string
	Dob  time.Time
}

const createUser = `
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob
`

type CreateUserParams struct {
	Name string
	Dob  time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Dob)

	var u User

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Dob,
	)

	return u, err
}

const getUser = `
SELECT id, name, dob
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)

	var u User

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Dob,
	)

	return u, err
}

const listUsers = `
SELECT id, name, dob
FROM users
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(
		ctx,
		listUsers,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Dob,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

const updateUser = `
UPDATE users
SET name = $2,
    dob = $3
WHERE id = $1
RETURNING id, name, dob
`

type UpdateUserParams struct {
	ID   int32
	Name string
	Dob  time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(
		ctx,
		updateUser,
		arg.ID,
		arg.Name,
		arg.Dob,
	)

	var u User

	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Dob,
	)

	return u, err
}

const deleteUser = `
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(
		ctx,
		deleteUser,
		id,
	)

	return err
}