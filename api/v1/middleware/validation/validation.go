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
package validation

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
	"github.com/hashicorp/go-multierror"
)

var decoder *schema.Decoder = schema.NewDecoder()
var validate *validator.Validate = validator.New()

// This function will take care of parsing the HTTP form sent by the
// request and optionally, will return an error if something fails.
func handleParseForm(r *http.Request) error {
	return r.ParseForm()
}

// This function will attempt to serialize contents of the parsed HTTP form
// into our desired struct.
func handleDecodeBody[T any](r *http.Request, dst *T) error {
	return decoder.Decode(dst, r.PostForm)
}

// This function will take care of validating the HTTP form by using
// a predefined schema struct.
func handleValidateBody[T any](r *http.Request, dst *T) error {
	return validate.StructCtx(r.Context(), dst)
}

// This function will take care of the whole process for parsing, decoding
// and validating parts from the HTTP request.
func ValidateRequest[T any](r *http.Request, dst *T) error {
	// Initializing multierror and appending all errors to it repsectively
	var err *multierror.Error
	// Pushing errors from all internal operations to multierror
	multierror.Append(err, handleParseForm(r), handleDecodeBody(r, dst), handleValidateBody(r, dst))
	// Joining all errors and returning the result as one error
	return err.Unwrap()
}
