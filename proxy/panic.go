package proxy

import (
	"net/http"

	"github.com/bookqaq/010-record-api/logger"
)

// go's panic and recovery, reveice panic, log and send StatusInternalServerError.
// see https://go.dev/blog/defer-panic-and-recover
func proxyPanicRecovery(w http.ResponseWriter, method, path string) {
	if r := recover(); r != nil {
		logger.Error.Printf("%s %s http service panic recovered: %s\n", method, path, r)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
