package types

import "polygon.am/core/api/types/common"

type User struct {
	common.Model
	Email     string
	Username  string
	Password  string
	LastName  string
	FirstName string
}
