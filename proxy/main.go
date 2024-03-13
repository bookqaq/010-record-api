// Why the package name is called proxy? only time will tell...
package proxy

import (
	"010-record-api/config"
	"010-record-api/local"
	"010-record-api/logger"
	"net/http"
)

func Start() {
	local := local.New()
	logger.Warning.Printf("starting 010 record api on %s ...\n", config.Config.ListenAddress)
	logger.Warning.Println(http.ListenAndServe(config.Config.ListenAddress, local))
}
