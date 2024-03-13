package local

import (
	"010-record-api/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() http.Handler {
	local := gin.New()
	local.Use(gin.RecoveryWithWriter(logger.GetWriter()))

	groupMovie := local.Group("/movie") // record api group
	initRouterGroupMovie(groupMovie)

	local.StaticFS("/patcher", http.FS(MustGetPatcher()))
	// tell user way to access patcher
	fmt.Printf("\tto access embedded patcher, go http://your_ip:your_port/patcher/\n\n")

	// not setting this handler into group /movie
	local.PUT("/movie-upload/:filename", MovieUploadContext)

	return local
}

func initRouterGroupMovie(group *gin.RouterGroup) {
	group.GET("/server/status", MovieServerStatus)
	group.POST("/sessions/new", MovieSessionNew)
	group.POST("/sessions/:sid/videos/:vid/:operation", MovieUploadManagement)
}
