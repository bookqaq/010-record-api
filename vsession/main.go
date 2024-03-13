package vsession

import (
	"010-record-api/config"
	"fmt"
	"sync"
	"time"
)

// TODO: clean outdated records with trigger
var mapVsessionFilename *sync.Map = new(sync.Map)

func GetNewUploadURL(session, vid string) string {
	filename := fmt.Sprintf("%d.mp4", time.Now().Unix())
	mapVsessionFilename.Store(fmt.Sprintf("%s-%s", session, vid), filename)
	return fmt.Sprintf("http://%s/movie-upload/%s", config.Config.UploadServiceAddress, filename)
}
