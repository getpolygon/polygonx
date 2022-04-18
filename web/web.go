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
package web

import (
	"net/http"
	"time"

	"github.com/getpolygon/corexp/internal/deps"
	"github.com/getpolygon/corexp/internal/settings"
	v1 "github.com/getpolygon/corexp/web/api/v1"
	"github.com/getpolygon/corexp/web/notfound"
	"github.com/getpolygon/corexp/web/wellknown"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func New(deps *deps.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.GetHead)
	r.Use(middleware.NoCache)

	// Heartbeat endpoint middleware useful to setting up a path like `/ping` that load balancers or
	// uptime testing external services can make a request before hitting any routes.
	r.Use(middleware.Heartbeat("/heartbeat"))

	// We are using the `httplimit` package for limiting all requests to the API. The maximum number
	// of consecutive requests is 100, and the threshold reset time is 1 minute.
	r.Use(httprate.LimitAll(100, 1*time.Minute))

	if *deps.Settings.Logging == settings.LoggingEnvDevelopment || deps.Settings.Logging == nil {
		// Enabling HTTP logging only during development.
		r.Use(middleware.Logger)
	}

	// Supporting multi-version API routes, for backwards-compatibility
	// and for general accessibility in applications.
	r.Route("/api", func(r chi.Router) {
		// Mounting version 1 API route to the main router.
		r.Mount("/v1", v1.Router(deps))
	})

	// `.well-known` routes will contain publicly accessible information
	// current Polygon instance, and will be used for health checks, and
	// version reporting, for usage with external tools and the ui.
	// more info: https://en.wikipedia.org/wiki/Well-known_URI#:~:text=A%20well%2Dknown%20URI%20is,well%2Dknown%20locations%20across%20servers.
	// 			  https://www.keycdn.com/support/well-known#:~:text=The%20well%2Dknown%20path%20prefix%20is%20essentially%20a%20place%20where,is%20defined%20as%20follows%3A%20%2F.
	r.Mount("/.well-known", wellknown.Router(deps))

	// A custom not found route, where all the not found requests will
	// be redirected to. This will return an HTML response.
	r.Mount("/notfound", notfound.New())

	// Instead of defaulting to the catchall route provided by chi and
	// net/http, we are redirecting the user to the page where we have
	// created a custom HTML page, with some information about current
	// instance.
	r.NotFound(http.RedirectHandler("/notfound", http.StatusTemporaryRedirect).ServeHTTP)

	return r
}
