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
package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sony/sonyflake"
)

const TemporaryUserSignUpExpirationTime = 20 * time.Minute

type TemporaryUserSignUpPayload struct {
	Name     string
	Email    string
	Username string
}

// This will temporarily persist user's registration details to Redis, for
// 20 minutes, and after that, if the account was not activated with the
// token, the token will auto-remove.
func TemporaryUserSignUp(ctx context.Context, r *redis.Client, p TemporaryUserSignUpPayload) (string, error) {
	// Generating a random token.
	genToken, err := sonyflake.NewSonyflake(sonyflake.Settings{}).NextID()
	if err != nil {
		return "", err
	}

	tkn := fmt.Sprint(genToken)
	// Persisting the information in Redis.
	if err := r.Set(ctx, fmt.Sprint(tkn), p, TemporaryUserSignUpExpirationTime).Err(); err != nil {
		return "", err
	}

	return tkn, nil
}
