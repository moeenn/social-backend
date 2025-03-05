// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_queries.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const userByEmail = `-- name: UserByEmail :one
select id, email, password, name, role, created_at, deleted_at from users
where email = $1
limit 1
`

func (q *Queries) UserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, userByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const userByID = `-- name: UserByID :one
select id, email, password, name, role, created_at, deleted_at from users
where id = $1 and deleted_at is null
limit 1
`

func (q *Queries) UserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, userByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const userCreate = `-- name: UserCreate :one
insert into users (id, email, password, name, role)
values ($1, $2, $3, $4, $5)
returning id, email, password, name, role, created_at, deleted_at
`

type UserCreateParams struct {
	ID       uuid.UUID
	Email    string
	Password string
	Name     string
	Role     string
}

func (q *Queries) UserCreate(ctx context.Context, arg UserCreateParams) (User, error) {
	row := q.db.QueryRow(ctx, userCreate,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.Name,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const userDelete = `-- name: UserDelete :exec
update users
set deleted_at = now()
where id = $1
`

func (q *Queries) UserDelete(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, userDelete, id)
	return err
}

const userList = `-- name: UserList :many
select id, email, password, name, role, created_at, deleted_at from users
where deleted_at is null
limit $1
offset $2
`

type UserListParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) UserList(ctx context.Context, arg UserListParams) ([]User, error) {
	rows, err := q.db.Query(ctx, userList, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.Role,
			&i.CreatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userListCount = `-- name: UserListCount :one
select count(*) from users
where deleted_at is null
`

func (q *Queries) UserListCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, userListCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const userUpdate = `-- name: UserUpdate :one
update users
set password = $2, name = $3, role = $4
where id = $1
returning id, email, password, name, role, created_at, deleted_at
`

type UserUpdateParams struct {
	ID       uuid.UUID
	Password string
	Name     string
	Role     string
}

func (q *Queries) UserUpdate(ctx context.Context, arg UserUpdateParams) (User, error) {
	row := q.db.QueryRow(ctx, userUpdate,
		arg.ID,
		arg.Password,
		arg.Name,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
