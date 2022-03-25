// BSD 3-Clause License

// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
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
