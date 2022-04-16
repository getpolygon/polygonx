-- BSD 3-Clause License

-- Copyright (c) 2021, Michael Grigoryan
-- All rights reserved.

-- Redistribution and use in source and binary forms, with or without
-- modification, are permitted provided that the following conditions are met:

-- 1. Redistributions of source code must retain the above copyright notice, this
--    list of conditions and the following disclaimer.

-- 2. Redistributions in binary form must reproduce the above copyright notice,
--    this list of conditions and the following disclaimer in the documentation
--    and/or other materials provided with the distribution.

-- 3. Neither the name of the copyright holder nor the names of its
--    contributors may be used to endorse or promote products derived from
--    this software without specific prior written permission.

-- THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
-- AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
-- IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
-- DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
-- FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
-- DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
-- SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
-- CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
-- OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
-- OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
-- name: GetPostByID :one
select * from "posts" where "id" = $1;

-- name: InsertPost :one
insert into "posts" ("user", "title", "content")
values ($1, $2, $3) returning *;

-- name: DeletePostByID :execresult
delete from "posts" where "id" = $1;

-- name: GetPostsOfUserAfterID :many
-- This query will fetch posts of a user by using their ID, starting 
-- from a cursor provided by the URL query of the HTTP request.
select
    "post"."id",
    "post"."title",
    "post"."updated",
    "post"."content",
    "post"."updated_at"
from "posts" "post"
left join
    "users" "user" on "post"."user" = "user"."id" 
where
    "user"."username" = $1
        and
    -- ULIDs are sortable, thus, we do not have to do any wizardry
    -- that would involve sorting UUIDs. more info at: https://github.com/ulid/spec
    "post"."id" > $2
group by "post"."id", "user"."id" limit 10;
