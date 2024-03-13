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
	wrapw := responsewriter.NewWrapped(w)
	h.local.ServeHTTP(wrapw, r)
	logger.Warning.Println(wrapw.StatusCode, r.Method, r.URL.Path)
}
