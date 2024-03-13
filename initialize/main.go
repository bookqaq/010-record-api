package initialize

import (
	"010-record-api/config"
	"010-record-api/logger"
	"os"
)

func CreateRequiredDirectory() {
	info, err := os.Stat(config.Config.VideoSaveDirectory)
	if os.IsNotExist(err) {
		if err := os.Mkdir(config.Config.VideoSaveDirectory, 0755); err != nil {
			logger.Error.Fatal("create video directory failed: ", err)
		}
		return
	}

	if !info.IsDir() {
		logger.Error.Fatalf("configured video directory (%s) exists, and it's not a directory, please recreate or use another directory...", config.Config.VideoSaveDirectory)
		os.Exit(1)
	}
}
