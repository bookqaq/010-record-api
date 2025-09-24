package feature

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bookqaq/010-record-api/config"
	"github.com/bookqaq/010-record-api/logger"
	"github.com/bookqaq/010-record-api/utils"
	"github.com/bookqaq/010-record-api/vsession"
)

// these are data we need for xrpcIIDXMovieInfo
type requestXrpcIIDXMusicMovieInfo struct {
	SessionId string `json:"session_id"` // call.IIDX00Music.session_id : which session we want to update
	IIDXID    string `json:"iidxid"`     // call.IIDX00Music.iidxid     : iidxid of the video owner next

	// these are data we do not need but appears in the request. I'll write and
	// comment them, for we don't need these.
	// ClassId      string `json:"class_id"`
	// CompanyCode  string `json:"company_code"`
	// ConsumerCode string `json:"consumer_code"`
	// LocationId   string `json:"location_id"`
	// LocationName string `json:"location_name"`
	// Method       string `json:"method"`
	// MusicId      string `json:"music_id"`
	// ProcType     string `json:"proc_type"`
}

// handler POST /feature/xrpcIIDXMusicMovieInfo
// receive IIDX00music.movieinfo request data from a xrpc server, and update
// player info based on session id, including iidxid.
func FeatureXrpcIIDXMusicMovieInfo(w http.ResponseWriter, r *http.Request) {
	// only receive additional info when
	if config.Config.FeatureXrpcIIDXMusicMovieInfo != nil && *config.Config.FeatureXrpcIIDXMusicMovieInfo {
		// standard body read and parse process
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error.Println("feature xrpcIIDXMusicMovieInfo failed to read body:", err)
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"status": 400,
				"msg":    "can't read body",
			})
			return
		}
		logger.Debug.Printf("feature XrpcIIDXMusicMovieInfo receive body: %s", requestBody)

		var body requestXrpcIIDXMusicMovieInfo
		if err := json.Unmarshal(requestBody, &body); err != nil {
			logger.Error.Println("feature xrpcIIDXMusicMovieInfo failed to parse data:", err)
			utils.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"status": 400,
				"msg":    "data parse error",
			})
			return
		}

		// check session exists and get session info
		session, ok := vsession.MapInfo.Load(body.SessionId)
		if !ok {
			// create a new session info if not exists, because this call happens before begin-upload request.
			session = vsession.Info{}
		}
		info := session.(vsession.Info)

		// update session info with video owner's unique id
		info.VideoOwnerId = body.IIDXID

		// update vsession map. IIDX00music.movieinfo is called before the PUT request (before begin-upload
		// tbh). And this software is not intended to be run as a server for multiple client. Also upload
		// is single-threaded on one machine. So there will be no race condition. As a result, this update
		// is safe.
		vsession.MapInfo.Store(body.SessionId, info)
	}
	// done, return success
	utils.ResponseJSON(w, http.StatusOK, map[string]any{
		"status": 200,
		"msg":    "success",
	})
}
