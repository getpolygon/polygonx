package users

import "github.com/gofiber/fiber/v2"

// This endpoint is dedicated to fetching users' public
// information by using their usernames.
func GetUserByUsername(c *fiber.Ctx) error {
	return nil
}

// This endpoint is dedicated to fetching users' public
// information by using their IDs. This is useful when
// there are API consumers that would need to identify
// the user even when their username changes.
func GetUserById(c *fiber.Ctx) error {
	return nil
}

// This endpoint is dedicated to closing users' accounts
// and deleting all the information that is related to
// their profiles.
func CloseUserAccount(c *fiber.Ctx) error {
	return nil
}

// This endpoint is dedicated to updating users' public
// and private information by using their IDs.
func UpdateUserById(c *fiber.Ctx) error {
	return nil
}
