package persistence_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/go-playground/assert/v2"
	"polygon.am/core/pkg/persistence"
	"polygon.am/core/pkg/persistence/codegen"
)

var TestUser = codegen.InsertUserParams{
	Name:     gofakeit.Name(),
	Email:    gofakeit.Email(),
	Username: gofakeit.Username(),
	Password: gofakeit.Password(true, true, true, true, true, 8),
}

func TestInsertUser(t *testing.T) {
	_, err := persistence.Queries.InsertUser(context.Background(), TestUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	ctx, username := context.Background(), TestUser.Username
	if user, err := persistence.Queries.GetUserByUsername(ctx, username); err != nil {
		t.Fatal(err)
	} else {
		assert.NotEqual(t, user, nil)
	}
}

func TestDeleteUserByEmail(t *testing.T) {
	ctx, email := context.Background(), TestUser.Email
	if err := persistence.Queries.DeleteUserByEmail(ctx, email); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserByUsername(t *testing.T) {
	ctx, username := context.Background(), TestUser.Username
	if err := persistence.Queries.DeleteUserByUsername(ctx, username); err != nil {
		t.Fatal(err)
	}
}
