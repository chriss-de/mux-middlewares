package utilities

import (
	"net/http"
)

// NewWrappedResponseWriter creates a new ResponseWriter that wraps the original response writer
func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{ResponseWriter: w, code: http.StatusOK}
}

// WrappedResponseWriter definition of our responseWriter
type WrappedResponseWriter struct {
	http.ResponseWriter
	code  int
	bytes int64
}

func (b *WrappedResponseWriter) WriteHeader(code int) {
	b.code = code
	b.ResponseWriter.WriteHeader(code)
}

func (b *WrappedResponseWriter) Write(buf []byte) (int, error) {
	n, err := b.ResponseWriter.Write(buf)
	b.bytes += int64(n)
	return n, err
}

func (b *WrappedResponseWriter) Status() int {
	return b.code
}

func (b *WrappedResponseWriter) BytesWritten() int64 {
	return b.bytes
}

func (b *WrappedResponseWriter) Unwrap() http.ResponseWriter {
	return b.ResponseWriter
}
