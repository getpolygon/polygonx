-- name: InsertPostUpvoteByUser :one
insert into upvotes ("user", "post") 
values ($1, $2) returning *;

-- name: DeletePostUpvoteOfUser :exec
delete from upvotes where ("user", "post") = ($1, $2);
