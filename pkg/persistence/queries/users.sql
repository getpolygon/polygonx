-- name: GetUserByUsername :one
select * from "users" where "username" = $1 limit 1; 

-- name: GetUserByEmail :one
select * from "users" where "email" = $1 LIMIT 1;

-- name: InsertUser :one
insert into "users" ("name", "username", "email", "password") 
values ($1, $2, $3, $4) returning *;

-- name: DeleteUserByUsername :exec
delete from "users" where "username" = $1;

-- name: DeleteUserByEmail :exec
delete from "users" where "email" = $1;
