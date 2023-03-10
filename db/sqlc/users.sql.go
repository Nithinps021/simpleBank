// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hash_password,
  full_name,
  emial_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING username, hash_password, full_name, emial_id, password_last_changed, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	FullName     string `json:"full_name"`
	EmialID      string `json:"emial_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashPassword,
		arg.FullName,
		arg.EmialID,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashPassword,
		&i.FullName,
		&i.EmialID,
		&i.PasswordLastChanged,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hash_password, full_name, emial_id, password_last_changed, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashPassword,
		&i.FullName,
		&i.EmialID,
		&i.PasswordLastChanged,
		&i.CreatedAt,
	)
	return i, err
}
