// why the package name is called proxy, not handler or http or server? only time will tell...
package proxy

import (
	"net/http"

	"github.com/bookqaq/010-record-api/config"
	"github.com/bookqaq/010-record-api/local"
	"github.com/bookqaq/010-record-api/logger"
)

func Start() {
	local := local.New()

	handle := &localHandle{local: local}

	logger.Warning.Printf("starting 010 record api on %s ...\n", config.Config.ListenAddress)
	logger.Warning.Println(http.ListenAndServe(config.Config.ListenAddress, handle))
}
