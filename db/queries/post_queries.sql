-- name: PostCreate :one
insert into posts (id, title, content, created_by_id)
values ($1, $2, $3, $4)
returning *; 

-- name: PostUpdate :one
update posts
set title = $2, content = $3, updated_at = now()
where id = $1
returning *;

-- name: PostDelete :exec
update posts
set deleted_at = now()
where id = $1;

-- name: PostsList :many
select * from posts
where deleted_at is null
order by created_at
limit $1
offset $2;

-- TODO: get post by id