package local

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bookqaq/010-record-api/logger"
	"github.com/bookqaq/010-record-api/utils"
	"github.com/bookqaq/010-record-api/vsession"
)

// handler GET /movie/server/status
func MovieServerStatus(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, nil)
}

// handler POST /movie/sessions/new
func MovieSessionNew(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error.Println("movie new session body error:", err)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"msg":    "invalid parameter",
		})
		return
	}

	var body requestMovieSessionNew
	if err := json.Unmarshal(requestBody, &body); err != nil {
		logger.Error.Println("movie new session body error:", err)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"msg":    "invalid parameter",
		})
		return
	}

	utils.ResponseJSON(w, http.StatusOK, map[string]string{
		"status":  "200",
		"session": mockSID,
	})
}

// handler POST /movie/sessions/{sid}/videos/{vid}/{operation}
func MovieUploadManagement(w http.ResponseWriter, r *http.Request) {
	session := r.PathValue("sid")
	vid := r.PathValue("vid")
	operation := r.PathValue("operation")

	key := vsession.GetKey(session, vid)

	switch operation {
	case constUploadStatusBegin:
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error.Println("movie upload begin body error:", err)
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{
				"status": "400",
				"msg":    "invalid parameter",
			})
			return
		}

		var body requestMovieSessionUploadBegin
		if err := json.Unmarshal(requestBody, &body); err != nil {
			logger.Error.Println("movie upload begin body error:", err)
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{
				"status": "400",
				"msg":    "invalid parameter",
			})
			return
		}

		// not using all value from request body.
		// also implement a md5-based video management, with little session info
		vsession.MapInfo.Store(key, vsession.Info{
			ShopName:  body.EA3ShopName,
			MD5Sum:    body.MD5Sum,
			MusicId:   body.MusicId,
			Timestamp: body.Timestamp,
		})

		// assign a video upload url path
		utils.ResponseJSON(w, http.StatusOK, map[string]string{
			"status": "200",
			"url":    vsession.GetNewUploadURL(key),
		})
	case constUploadStatusEnd:
		// dummy return
		vsession.MapInfo.Delete(key)
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

// handler PUT /movie-upload/{key}
func MovieUploadContext(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	// Content-Type verification
	contentType := r.Header.Get("Content-Type")
	if contentType != "video/mp4" {
		logger.Error.Println("movie: malformed Content-Type: ", contentType, key)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if strings.EqualFold(key, "") {
		logger.Error.Println("movie: bad key", key)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "invalid upload url destination",
		})
		return
	}

	// get key info
	res, ok := vsession.MapInfo.Load(key)
	if !ok {
		logger.Error.Println("movie: key not exist", key)
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "invalid upload url destination",
		})
		return
	}

	// use keyinfo to generate filename
	info := res.(vsession.Info)
	filename := info.ToFileName()
	logger.Warning.Println("receive video upload request:", filename)

	written, err := vsession.ReceiveUploadVideo(r.Body, filename)
	if err != nil {
		logger.Error.Println("movie upload: ", err)
		utils.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"status": "500",
			"msg":    err.Error(),
		})
		return
	}

	logger.Warning.Printf("movie %s done, %d bytes written", filename, written)
	w.WriteHeader(http.StatusOK)
}
