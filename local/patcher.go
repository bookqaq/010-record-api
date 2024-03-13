package local

import (
	"010-record-api/logger"
	"embed"
	"io/fs"
)

// webpage that will be embedded into executable after compile

//go:embed static/patcher
var patcherFS embed.FS

func MustGetPatcher() fs.FS {
	patcherSubFolder, err := fs.Sub(patcherFS, "static/patcher")
	if err != nil {
		logger.Error.Fatalln("set patcher service failed:", err)
	}
	return patcherSubFolder
}
