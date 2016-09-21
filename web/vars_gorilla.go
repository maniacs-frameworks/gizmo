// +build !go1.7

package web

import (
	"net/http"

	"github.com/gorilla/context"
)

// Vars is a helper function for accessing route
// parameters from any server.Router implementation. This is the equivalent
// of using `mux.Vars(r)` with the Gorilla mux.Router.
func Vars(r *http.Request) map[string]string {
	if rv := context.Get(r, varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}

// SetRouteVars will set the given value into into the request context
// with the shared 'vars' storage key.
func SetRouteVars(r *http.Request, val interface{}) *http.Request {
	if val != nil {
		context.Set(r, varsKey, val)
	}

	return r
}

type contextKey int

// key to set/retrieve URL params from a
// Gorilla request context.
const varsKey contextKey = 2
