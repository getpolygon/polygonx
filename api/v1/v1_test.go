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
package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	v1 "polygon.am/core/api/v1"
)

func executeRequest(req *http.Request, router *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code: %d. Got %d\n", expected, actual)
	}
}

func TestV1StatusCodes(t *testing.T) {
	r := v1.Router()
	tests := []struct {
		Expected int
		Path     string
		Method   string
	}{
		{
			Method:   http.MethodGet,
			Path:     "/users/example",
			Expected: http.StatusUnauthorized,
		},
		{
			Method:   http.MethodGet,
			Path:     "/posts/12345678",
			Expected: http.StatusUnauthorized,
		},
		{
			Method:   http.MethodGet,
			Path:     "/auth/signin",
			Expected: http.StatusMethodNotAllowed,
		},
		{
			Method:   http.MethodGet,
			Path:     "/auth/signup",
			Expected: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.Method, tt.Path, nil)
		res := executeRequest(req, r)
		checkResponseCode(t, tt.Expected, res.Code)
	}
}
