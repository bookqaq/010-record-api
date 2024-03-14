package proxy

import (
	"net/http"

	"github.com/bookqaq/010-record-api/logger"
	"github.com/bookqaq/010-record-api/proxy/responsewriter"
)

type localHandle struct {
	local http.Handler
}

func (h *localHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer proxyPanicRecovery(w, r.Method, r.URL.Path) // panic recovery

	wrapw := responsewriter.NewWrapped(w) // response statusCode logger
	h.local.ServeHTTP(wrapw, r)           // serve locally

	logger.Warning.Println(wrapw.StatusCode, r.Method, r.URL.Path) // log
}
