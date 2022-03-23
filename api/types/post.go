package types

import (
	"database/sql"

	"polygon.am/core/api/types/common"
)

type Post struct {
	common.Model
	User    *User
	Title   string
	UserID  *string
	Content sql.NullString
}
