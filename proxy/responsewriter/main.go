package responsewriter

import "net/http"

// this wrapped writer is only used to get StatusCode, which should be logged
type WrappedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewWrapped(w http.ResponseWriter) *WrappedResponseWriter {
	return &WrappedResponseWriter{w, -1}
}

func (wrw *WrappedResponseWriter) WriteHeader(code int) {
	wrw.StatusCode = code
	wrw.ResponseWriter.WriteHeader(code)
}
