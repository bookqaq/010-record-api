package local

import (
	"net/http"
	"strings"

	"github.com/bookqaq/010-record-api/logger"
	"github.com/bookqaq/010-record-api/utils"
	"github.com/bookqaq/010-record-api/vsession"
)

// handler GET /movie/server/status
func MovieServerStatus(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, map[string]any{})
}

// handler POST /movie/sessions/new
func MovieSessionNew(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, map[string]string{
		"status":  "200",
		"session": mockSID,
	})
}

// handler POST /movie/sessions/{sid}/videos/{vid}/{operation}
func MovieUploadManagement(w http.ResponseWriter, r *http.Request) {
	session := r.PathValue("session")
	vid := r.PathValue("vid")
	operation := r.PathValue("operation")

	switch operation {
	case constUploadStatusBegin:
		// assign a video upload url path
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"status": "200",
			"url":    vsession.GetNewUploadURL(session, vid),
		})
	case constUploadStatusEnd:
		// dummy return
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"status":  "200",
			"session": mockSID,
		})
	default:
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "invalid movie upload state",
		})
	}
}

// handler PUT /movie-upload/{filename}
func MovieUploadContext(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	if strings.EqualFold(filename, "") {
		logger.Error.Println("movie: bad filename", filename)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "invalid upload url destination",
		})
		return
	}

	written, err := vsession.ReceiveUploadVideo(r.Body, filename)
	if err != nil {
		logger.Info.Println("movie upload: ", err)
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"status": "500",
			"msg":    err.Error(),
		})
		return
	}

	logger.Info.Printf("movie %s: %d bytes written", filename, written)
	w.WriteHeader(http.StatusOK)
}
