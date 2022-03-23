package common

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `db:"uuid,primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
