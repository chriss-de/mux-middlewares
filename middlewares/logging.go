package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/chriss-de/mux-middlewares/utilities"
)

// Logging calls slog to log the http request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wrap the http.ResponseWriter so we can access the http.Status in the defer()
		wrw := utilities.NewWrappedResponseWriter(w, r.ProtoMajor)

		reqStart := time.Now()

		// is called when the handler func is done
		defer func() {
			slog.Info("access",
				slog.String("request_id", r.Header.Get("X-Request-Id")),
				slog.String("remote_ip", r.RemoteAddr),
				slog.String("url", r.URL.Path),
				slog.String("proto", r.Proto),
				slog.String("method", r.Method),
				slog.Int("status", wrw.Status()),
				slog.String("latency", time.Since(reqStart).String()),
				slog.Int("size", wrw.BytesWritten()),
				slog.String("user_agent", r.UserAgent()),
			)
		}()

		next.ServeHTTP(wrw, r)
	})
}
