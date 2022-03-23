package types

import "polygon.am/core/api/types/common"

type Comment struct {
	common.Model
	User    *User
	Post    *Post
	Content string
	UserID  *string
	PostID  *string
}
