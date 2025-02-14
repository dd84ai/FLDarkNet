package web

import (
	"fmt"
	"html"
	"net/http"

	"github.com/darklab8/fl-darkstat/darkcore/core_types"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
)

const URLPing core_types.Url = "GET /ping"

func NewEndpointPing(w *Web) *registry.Endpoint {
	return &registry.Endpoint{
		Url: URLPing,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "pong at %q", html.EscapeString(r.URL.Path))
		},
	}
}
