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
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/spf13/viper"
	v1 "polygon.am/core/api/v1"
	config "polygon.am/core/pkg"
)

func init() {
	// Loading the configuration in viper
	err := config.Load()
	if err != nil {
		log.Fatal("config error:", err)
	}
}

func main() {
	r := chi.NewRouter()

	// Only enabling request logging if it is specified
	// in the config.
	if viper.GetBool("polygon.general.logRequests") {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.GetHead)
	r.Use(middleware.NoCache)
	r.Use(middleware.Recoverer)

	// Only enabling the heartbeat route, if it is specified
	// in the config.
	if viper.GetBool("polygon.general.enableHeartbeatRoute") {
		r.Use(middleware.Heartbeat("/status"))
	}

	requestLimit := viper.GetInt("polygon.security.requests.max")
	windowLength := viper.GetInt("polygon.security.requests.threshold")
	r.Use(httprate.LimitAll(requestLimit, time.Duration(windowLength)))

	r.Group(func(r chi.Router) {
		r.Mount("/api/v1", v1.Router())
	})

	// Getting the connection address from the configuration and
	// attempting to bind to it.
	addr := viper.GetString("polygon.general.addr")
	log.Println("corexp started at http://" + addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}
