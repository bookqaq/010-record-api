package local

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/bookqaq/010-record-api/config"
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
	vid := r.PathValue("vid") // vid is not nessary because upload is usually single-threaded (1.1.0+)
	operation := r.PathValue("operation")

	// updated in 1.1.0, use session_id only as key, there are no video upload
	// parallelization in a single session, so vid is not needed.
	key := session

	logger.Debug.Printf("MovieUploadManagement session: %s vid: %s operation: %s", session, vid, operation)

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

		// try to get the session info first (when Config.FeatureFileNameAddVideoOwner
		// is enabled, some session info will be set before this http request)
		res, ok := vsession.MapInfo.Load(key)
		if !ok {
			res = vsession.Info{}
		}
		info := res.(vsession.Info)

		// update session info with data sent to WebAPI2
		info.ShopName = body.EA3ShopName
		info.MD5Sum = body.MD5Sum
		info.MusicId = body.MusicId
		info.Timestamp = body.Timestamp
		// set 1p/2p player name
		info.VideoOwnerName = body.PlayerNames[0]
		if len(body.PlayerNames) > 1 {
			info.VideoOwnerName += "_" + body.PlayerNames[1]
		}

		// not using all value from request body.
		// also implement a md5-based video management, with little session info
		vsession.MapInfo.Store(key, info)

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

	// only generate filename with owner info if configuration is enabled
	if config.Config.FeatureXrpcIIDXMusicMovieInfo != nil && *config.Config.FeatureXrpcIIDXMusicMovieInfo {
		filename = info.ToFileNameWithOwner()
	}
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
