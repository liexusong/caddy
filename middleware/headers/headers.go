// Package headers provides middleware that appends headers to
// requests based on a set of configuration rules that define
// which routes receive which headers.
package headers

import (
	"net/http"

	"github.com/mholt/caddy/middleware"
)

// Headers is middleware that adds headers to the responses
// for requests matching a certain path.
type Headers struct {
	Next  http.HandlerFunc
	Rules []HeaderRule
}

// ServeHTTP implements the http.Handler interface and serves the requests,
// adding headers to the response according to the configured rules.
func (h Headers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rule := range h.Rules {
		if middleware.Path(r.URL.Path).Matches(rule.Url) {
			for _, header := range rule.Headers {
				w.Header().Set(header.Name, header.Value)
			}
		}
	}
	h.Next(w, r)
}

type (
	// HeaderRule groups a slice of HTTP headers by a URL pattern.
	// TODO: use http.Header type instead??
	HeaderRule struct {
		Url     string
		Headers []Header
	}

	// Header represents a single HTTP header, simply a name and value.
	Header struct {
		Name  string
		Value string
	}
)