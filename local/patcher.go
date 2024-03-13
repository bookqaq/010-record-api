package local

import (
	"010-record-api/logger"
	"embed"
	"io/fs"
)

// webpage that will be embedded into executable after compile.
// below comment declare the usage of go embed toolchain and should NOT be deleted
//
//go:embed static/patcher
var patcherFS embed.FS

// get patcher static file fs
func MustGetPatcher() fs.FS {
	// go's builtin static file http handle sum all the path, including fs directory path and
	// http url path. So in order to serve files in static/patcher/index.html(relative to this file)
	// and assign URL path /patcher/ to it, I have to leave folder "patcher" not being included in
	// fs.Sub() like the code below, and set handler URL path to "patcher/".
	// fuck this.
	// https://stackoverflow.com/questions/74969821/go-http-fileserver-gives-unexpected-404-error
	patcherSubFS, err := fs.Sub(patcherFS, "static")
	if err != nil {
		logger.Error.Fatalln("set patcher service failed:", err)
	}
	return patcherSubFS
}
