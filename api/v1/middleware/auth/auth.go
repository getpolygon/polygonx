package middleware

import (
	"log"
	"os"

	"github.com/go-chi/jwtauth"
)

// A global JWT authentication strategy which will be used as
// a chi jwtauth.Verifier middleware value to find, verify and
// validate the
var Strategy *jwtauth.JWTAuth

func init() {
	// Getting the signing key from an environment variable and panicing
	// if it was not specified or was empty.
	pkey := os.Getenv("POLYGON_CORE_CONFIG_JWTPKEY")
	if pkey == "" {
		log.Fatal("jwt signing key cannot be empty. consider updating `POLYGON_CORE_CONFIG_JWTPKEY` environment variable")
	}

	// Creating a new JWT strategy using the private key
	// from the environment variableand assigning it to
	// a global variable.
	Strategy = jwtauth.New("HS256", []byte(pkey), nil)
}
