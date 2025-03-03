-- name: UserCreate :one
insert into users (id, email, password, role)
values ($1, $2, $3, $4)
returning *;

-- name: UserListCount :one
select count(*) from users;

-- name: UserList :many
select * from users
limit $1
offset $2;

-- name: UserByID :one
select * from users
where id = $1
limit 1;

-- name: UserByEmail :one
select * from users
where email = $1
limit 1;

-- name: UserUpdate :one
update users
set role = $2, password = $3
where id = $1
returning *;

-- name: UserDelete :one
delete from users
where id = $1
returning *;