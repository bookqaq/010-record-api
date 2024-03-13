package proxy

import (
	"010-record-api/logger"
	"010-record-api/proxy/responsewriter"
	"net/http"
)

type localHandle struct {
	local http.Handler
}

func (h *localHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrapw := responsewriter.NewWrapped(w)
	h.local.ServeHTTP(wrapw, r)
	logger.Warning.Println(wrapw.StatusCode, r.Method, r.URL.Path)
}
