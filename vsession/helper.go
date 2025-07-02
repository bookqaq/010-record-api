package vsession

import (
	"fmt"

	"github.com/bookqaq/010-record-api/config"
)

// wrapper function to get an upload url
func GetNewUploadURL(key string) string {
	return fmt.Sprintf("http://%s/movie-upload/%s", config.Config.UploadServiceAddress, key)
}

// GetKey returns a key for an upload session.
//
// Deprecated: it's not necessary for session key to contain video upload counter vid.
func GetKey(session, id string) string {
	return fmt.Sprintf("%s-%s", session, id)
}
