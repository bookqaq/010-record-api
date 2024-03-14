package vsession

import (
	"fmt"
	"io"
	"os"

	"github.com/bookqaq/010-record-api/config"
)

const videoFileExtensionName = ".mp4"

func ReceiveUploadVideo(src io.Reader, fileName string) (int64, error) {
	finalFilename := fmt.Sprintf("%s/%s%s", config.Config.VideoSaveDirectory, fileName, videoFileExtensionName)

	fp, err := os.OpenFile(finalFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return -1, fmt.Errorf("failed to open file %s: %v", finalFilename, err)
	}
	defer fp.Close()

	written, err := io.Copy(fp, src)
	if err != nil {
		return -1, fmt.Errorf("failed to write file %s: %v", finalFilename, err)
	}

	return written, nil
}
