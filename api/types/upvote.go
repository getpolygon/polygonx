package types

import "polygon.am/core/api/types/common"

type Upvote struct {
	common.Model
	User   *User
	Post   *Post
	UserID *string
	PostID *string
}
