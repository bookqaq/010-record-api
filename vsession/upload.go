package vsession

import (
	"010-record-api/config"
	"fmt"
	"io"
	"os"
)

func ReceiveUploadVideo(src io.Reader, fileName string) (int64, error) {
	finalFilename := fmt.Sprintf("%s/%s", config.Config.VideoSaveDirectory, fileName)

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
