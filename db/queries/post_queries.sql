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

-- name: PostsCount :one
select count(*) from posts 
where deleted_at is null;

-- name: PostsList :many
select 
	p.id,
  p.title,
  p.content,
  p.created_by_id,
  u.name as created_by_name,
  p.comments_count,
  p.created_at,
  p.updated_at
from posts p
join users u on u.id = p.created_by_id
where p.deleted_at is null
order by p.created_at desc
limit $1
offset $2;

-- name: PostById :one
select 
	p.id,
  p.title,
  p.content,
  p.created_by_id,
  u.name as created_by_name,
  p.comments_count,
  p.created_at,
  p.updated_at
from posts p
join users u on u.id = p.created_by_id
where p.id = $1 and p.deleted_at is null
limit 1;