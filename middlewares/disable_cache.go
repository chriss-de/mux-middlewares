package middlewares

import (
	"net/http"
)

// DisableCache adds http headers to disable caching in browser
func DisableCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "private, no-cache, no-store, max-age=0, no-transform")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("Vary", "*")

		next.ServeHTTP(w, r)
	})
}
