// why the package name is called proxy, not handler or http or server? only time will tell...
package proxy

import (
	"010-record-api/config"
	"010-record-api/local"
	"010-record-api/logger"
	"net/http"
)

func Start() {
	local := local.New()

	handle := &localHandle{local: local}

	logger.Warning.Printf("starting 010 record api on %s ...\n", config.Config.ListenAddress)
	logger.Warning.Println(http.ListenAndServe(config.Config.ListenAddress, handle))
}
