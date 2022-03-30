-- name: GetPostByID :one
select
    p.id,
    p.title,
    p.content,
    p.created_at
from posts p inner join (
    select id, "name", username from users u
) u ON p.user_id = u.id
where p.id = $1 limit 1;

-- name: GetPostsOfUserByUsername :many
select
    p.id,
    p.title,
    p.content,
    p.created_at
from posts p inner join (
    select id, "name", username from users u
) u ON p.user_id = u.id
where u.username = $1;
