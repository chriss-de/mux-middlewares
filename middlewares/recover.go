package middlewares

import (
	"net/http"
)

// Recovery catches any panic so we don't crash
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)

	})
}
