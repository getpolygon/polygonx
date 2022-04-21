package cache

import (
	"context"

	"github.com/getpolygon/corexp/internal/settings"
	"github.com/go-redis/redis/v8"
)

// This function will attempt to create a new Redis instance
// with the provided settings.
func New(s *settings.Settings) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr: s.Redis,
	})

	// After creating a new Redis client, we are testing the
	// connection, by pinging the instance and checking if
	// the request was executed successfully.
	if err := c.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return c, nil
}
