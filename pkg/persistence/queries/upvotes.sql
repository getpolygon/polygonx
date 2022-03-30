-- name: InsertPostUpvoteByID :one
insert into upvotes (user_id, post_id) 
values ($1, $2) returning *;

-- name: DeletePostUpvoteByID :exec
delete from upvotes where id = $1;
