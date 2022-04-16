// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package postgres_codegen

import (
	"context"
	"time"
)

const checkUserExistsByID = `-- name: CheckUserExistsByID :one
select 1::boolean from "users" where "id" = $1 LIMIT 1
`

func (q *Queries) CheckUserExistsByID(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUserExistsByID, id)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const getActiveUsersCountHalfYear = `-- name: GetActiveUsersCountHalfYear :one
select count("id") from "users" where datediff("active_at", current_timestamp) between 0 and 180
`

func (q *Queries) GetActiveUsersCountHalfYear(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getActiveUsersCountHalfYear)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getActiveUsersCountMonth = `-- name: GetActiveUsersCountMonth :one
select count("id") from "users" where datediff("active_at", current_timestamp) between 0 and 30
`

func (q *Queries) GetActiveUsersCountMonth(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getActiveUsersCountMonth)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFullUserByEmail = `-- name: GetFullUserByEmail :one
select id, name, email, password, username, online_at, created_at from "users" where "email" = $1 LIMIT 1
`

// This query should only be used for retrieving private
// user information.
func (q *Queries) GetFullUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getFullUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Username,
		&i.OnlineAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPartialUserByUsername = `-- name: GetPartialUserByUsername :one






select
    "id",
    "name",
    "username",
    "created_at"
from "users" where "username" = $1 limit 1
`

type GetPartialUserByUsernameRow struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

// BSD 3-Clause License
// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// This query should only be used when retrieving public
// user information.
func (q *Queries) GetPartialUserByUsername(ctx context.Context, username string) (GetPartialUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getPartialUserByUsername, username)
	var i GetPartialUserByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const getTotalUsersCount = `-- name: GetTotalUsersCount :one
select count("id") from "users"
`

func (q *Queries) GetTotalUsersCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalUsersCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const insertUser = `-- name: InsertUser :one
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
    "created_at"
`

type InsertUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type InsertUserRow struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

// This query is only used whenever a new user is created
// and must not be used anywhere else.
func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Name,
		arg.Email,
		arg.Username,
		arg.Password,
	)
	var i InsertUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}