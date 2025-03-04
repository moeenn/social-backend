-- name: PostCreate :one
insert into posts (id, title, content, created_by_id)
values ($1, $2, $3, $4)
returning *; 

-- name: PostUpdate :one
update posts
set title = $2, content = $3, updated_at = now()
where id = $1
returning *;

-- name: PostDelete :one
update posts
set deleted_at = now()
where id = $1
returning *;

-- name: PostsList :many
select * from posts
order by created_at
limit $1
offset $2;