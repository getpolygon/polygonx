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
-- name: GetPartialUserByUsername :one
-- This query should only be used when retrieving public
-- user information.
select
    "id",
    "name",
    "username",
    "created_at"
from "users" where "username" = $1 limit 1; 

-- name: GetFullUserByEmail :one
-- This query should only be used for retrieving private
-- user information.
select * from "users" where "email" = $1 LIMIT 1;

-- name: CheckUserExistsByID :one
select 1::boolean from "users" where "id" = $1 LIMIT 1;

-- name: InsertUser :one
-- This query is only used whenever a new user is created
-- and must not be used anywhere else.
insert into "users" (
    "name", 
    "email", 
    "username", 
    "password"
) values ($1, $2, $3, $4) returning 
    "id", 
    "name",
    "email", 
    "username",
    "created_at";
