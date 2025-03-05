-- name: UserCreate :one
insert into users (id, email, password, name, role)
values ($1, $2, $3, $4, $5)
returning *;

-- name: UserListCount :one
select count(*) from users
where deleted_at is null;

-- name: UserList :many
select * from users
where deleted_at is null
limit $1
offset $2;

-- name: UserByID :one
select * from users
where id = $1 and deleted_at is null
limit 1;

-- name: UserByEmail :one
select * from users
where email = $1
limit 1;

-- name: UserUpdate :one
update users
set password = $2, name = $3, role = $4
where id = $1
returning *;

-- name: UserDelete :exec
update users
set deleted_at = now()
where id = $1;
