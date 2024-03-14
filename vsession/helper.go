package vsession

import (
	"fmt"

	"github.com/bookqaq/010-record-api/config"
)

// wrapper function to get an upload url
func GetNewUploadURL(key string) string {
	return fmt.Sprintf("http://%s/movie-upload/%s", config.Config.UploadServiceAddress, key)
}

func GetKey(session, id string) string {
	return fmt.Sprintf("%s-%s", session, id)
}
