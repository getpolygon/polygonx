-- name: GetPostByID :one
select * from "posts" where "id" = $1;

-- name: InsertPost :one
insert into "posts" ("user", "title", "content")
values ($1, $2, $3) returning *;
