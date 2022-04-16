package web

import (
	"net/http"
	"time"

	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/settings"
	v1 "github.com/getpolygon/corexp/web/api/v1"
	"github.com/getpolygon/corexp/web/notfound"
	"github.com/getpolygon/corexp/web/wellknown"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

func New(p *postgres_codegen.Queries, s *settings.Settings) *chi.Mux {
	r := chi.NewRouter()

	// Instead of defaulting to the catchall route provided by chi and
	// net/http, we are redirecting the user to the page where we have
	// created a custom HTML page, with some information about current
	// instance.
	r.NotFound(http.RedirectHandler("/notfound", http.StatusTemporaryRedirect).ServeHTTP)

	r.Use(middleware.GetHead)
	r.Use(middleware.NoCache)

	// Heartbeat endpoint middleware useful to setting up a path like `/ping` that load balancers or
	// uptime testing external services can make a request before hitting any routes.
	r.Use(middleware.Heartbeat("/heartbeat"))

	// We are using the `httplimit` package for limiting all requests to the API. The maximum number
	// of consecutive requests is 100, and the threshold reset time is 1 minute.
	r.Use(httprate.LimitAll(100, 1*time.Minute))
	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Mount("/v1", v1.Router())
		})

		// `.well-known` routes will contain publicly accessible information
		// current Polygon instance, and will be used for health checks, and
		// version reporting, for usage with external tools and the ui.
		// more info: https://en.wikipedia.org/wiki/Well-known_URI#:~:text=A%20well%2Dknown%20URI%20is,well%2Dknown%20locations%20across%20servers.
		// 			  https://www.keycdn.com/support/well-known#:~:text=The%20well%2Dknown%20path%20prefix%20is%20essentially%20a%20place%20where,is%20defined%20as%20follows%3A%20%2F.
		r.Mount("/.well-known", wellknown.Router(p, s))

		// A custom not found route, where all the not found requests will
		// be redirected to. This will return an HTML response.
		r.Mount("/notfound", notfound.New())
	})

	return r
}
