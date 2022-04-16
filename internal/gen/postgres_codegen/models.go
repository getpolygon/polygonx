// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package postgres_codegen

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	Post      string    `json:"post"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        string         `json:"id"`
	User      string         `json:"user"`
	Title     string         `json:"title"`
	Content   sql.NullString `json:"content"`
	Updated   bool           `json:"updated"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Upvote struct {
	Post      string    `json:"post"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	OnlineAt  time.Time `json:"online_at"`
	CreatedAt time.Time `json:"created_at"`
}