package common

import "time"

type Model struct {
	ID        string `db:"uuid,primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
