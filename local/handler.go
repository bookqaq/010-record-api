package local

import (
	"010-record-api/logger"
	"010-record-api/vsession"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// handler GET /movie/server/status
func MovieServerStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

// handler POST /movie/sessions/new
func MovieSessionNew(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "200",
		"session": mockSID,
	})
}

// handler POST /movie/sessions/:sid/videos/:vid/:operation
func MovieUploadManagement(ctx *gin.Context) {
	session := ctx.Param("session")
	vid := ctx.Param("vid")
	operation := ctx.Param("operation")

	switch operation {
	case constUploadStatusBegin:
		// assign a video upload url path
		ctx.JSON(http.StatusOK, map[string]string{
			"status": "200",
			"url":    vsession.GetNewUploadURL(session, vid),
		})
	case constUploadStatusEnd:
		// dummy return
		ctx.JSON(http.StatusOK, map[string]string{
			"status":  "200",
			"session": mockSID,
		})
	default:
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "test error msg",
		})
	}
}

// handler PUT /movie-upload/:filename
func MovieUploadContext(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if strings.EqualFold(filename, "") {
		logger.Error.Println("movie: bad filename", filename)
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": "400",
			"msg":    "invalid upload url destination",
		})
		return
	}

	written, err := vsession.ReceiveUploadVideo(ctx.Request.Body, filename)
	if err != nil {
		logger.Info.Println("movie upload: ", err)
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"status": "500",
			"msg":    err.Error(),
		})
		return
	}

	logger.Info.Printf("movie %s: %d bytes written", filename, written)
	ctx.Status(http.StatusOK)
}
