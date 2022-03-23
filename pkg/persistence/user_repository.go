package persistence

import "polygon.am/core/api/types"

type UserRepository interface {
	Find(user *types.User, id string) error
}
