package persistence_test

import (
	"testing"

	"polygon.am/core/pkg/persistence"
)

func TestConnect(t *testing.T) {
	if err := persistence.Connect(); err != nil {
		t.Fatal(err)
	}
}
