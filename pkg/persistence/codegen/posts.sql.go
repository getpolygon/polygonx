// Code generated by sqlc. DO NOT EDIT.
// source: posts.sql

package codegen

import (
	"context"
	"database/sql"
)

const deletePostByID = `-- name: DeletePostByID :exec
delete from posts where id = $1
`

func (q *Queries) DeletePostByID(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePostByID, id)
	return err
}

const getPostByID = `-- name: GetPostByID :one
select id, "user", title, content, updated, updated_at from "posts" where "id" = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id string) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.User,
		&i.Title,
		&i.Content,
		&i.Updated,
		&i.UpdatedAt,
	)
	return i, err
}

const insertPost = `-- name: InsertPost :one
insert into "posts" ("user", "title", "content")
values ($1, $2, $3) returning id, "user", title, content, updated, updated_at
`

type InsertPostParams struct {
	User    string         `json:"user"`
	Title   string         `json:"title"`
	Content sql.NullString `json:"content"`
}

func (q *Queries) InsertPost(ctx context.Context, arg InsertPostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, insertPost, arg.User, arg.Title, arg.Content)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.User,
		&i.Title,
		&i.Content,
		&i.Updated,
		&i.UpdatedAt,
	)
	return i, err
}
