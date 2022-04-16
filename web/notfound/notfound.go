package notfound

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/go-chi/render"
)

// Caching the contents of the notfound HTML template to
// memory for faster response times and less I/O activity.
var content string

const DefaultNotFoundTemplatePath string = "templates/notfound.html"

// This function will return a new not found handler for usage
// with catchall request endpoints. It will render a nice HTML
// page instead of the default `not found` plain text.
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if content == "" {
			path, err := filepath.Abs(DefaultNotFoundTemplatePath)
			html, err := ioutil.ReadFile(path)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			content = string(html)
		}

		render.Status(r, http.StatusNotFound)
		render.HTML(w, r, content)
	}
}
