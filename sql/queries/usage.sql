-- name: GetFullUsageStats :one
select
    count(distinct "users"."id") as "users_count",
    count(distinct "posts"."id") as "posts_count",
    count(distinct "comments"."id") as "comments_count",
    (
        select
            count(distinct "id")
        from
            "users"
        where
            date_part('day', "online_at" :: date) - date_part('day', current_timestamp :: date) between 0
            and 30
    ) as "active_users_month",
    (
        select
            count(distinct "id")
        from
            "users"
        where
            date_part('day', "online_at" :: date) - date_part('day', current_timestamp :: date) between 0
            and 180
    ) as "active_users_half_year"
from
    "users",
    "posts",
    "comments";