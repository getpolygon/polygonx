package util_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"polygon.am/core/pkg/util"
)

// A function, the error of which can be ignored, since we
// are getting the string directly. This function is used
// only in this test.
func nonthrowing() (string, error) {
	return "Hello World", nil
}

// This function is going to return a single error, instead
// of returning a single output value
func throwing() (string, error) {
	return "", errors.New("Error 1")
}

func TestAssumeNoError(t *testing.T) {
	tests := []struct {
		mustfail bool
		name     string
		want     any
		executor func() (string, error)
	}{
		{executor: nonthrowing,
			mustfail: false,
			want:     "Hello World",
			name:     "AssumeNoError() should return a string but no error"},
		{executor: throwing,
			want:     "",
			mustfail: true,
			name:     "AssumeNoError() should panic with single error being present",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mustfail {
				require.Panics(t, func() { util.AssumeNoError(tt.executor()) })
			} else {
				if got := util.AssumeNoError(tt.executor()); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AssumeNoError() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
