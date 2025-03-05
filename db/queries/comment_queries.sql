-- name: CommentAdd :one
insert into comments (id, content, post_id, created_by_id, hierarchy_id, parent_comment_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: CommentAddReply :one
with parent_comment_hierarchy_id as (
  select hierarchy_id from comments
  where id = $5
)
insert into comments (id, content, post_id, created_by_id, hierarchy_id, parent_comment_id)
values ($1, $2, $3, $4, (select hierarchy_id from parent_comment_hierarchy_id), $5)
returning *;


-- name: CommentUpdate :one
update comments
set content = $2, updated_at = now()
where id = $1
returning *;

-- name: CommentDelete :one
update comments
set deleted_at = now()
where id = $1
returning *;

-- name: CommentLike :one
with increment_like_count as (
  update comments
  set likes_count = likes_count + 1
  where id = $1
  returning *
)
insert into comment_likes (comment_id, user_id)
values ($1, $2)
returning (select * from increment_like_count);

-- name: CommentsByPostId :many
select * from comments
where post_id = $1
order by created_at
limit $2
offset $3;