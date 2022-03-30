-- name: GetUserByUsername :one
select
    id,
    "name",
    username,
    created_at
from users where username = $1 limit 1; 

-- name: GetUserByID :one
select
    id,
    "name",
    username,
    created_at
from users where id = $1 limit 1;
